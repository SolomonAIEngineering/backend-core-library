package redis

import (
	"fmt"
	"net/url"
	"time"

	corelib "github.com/SimifiniiCTO/simfiny-core-lib"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

// startCachePool starts a Redis connection pool for caching. It creates a Redis pool with a maximum of
// 3 idle connections and an idle timeout of 240 seconds. It also sets a function to test the
// connection on borrow by sending a PING command to the Redis server. Additionally, it sets the
// version of the service in Redis with an expiry time of one minute and schedules a periodic update of
// the version using a ticker. The function takes a ticker and a stop channel as arguments to control
// the periodic version updates.
func (c *Client) startCachePool(ticker *time.Ticker, stopCh <-chan struct{}) {
	c.pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        c.getCacheConn,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// set <hostname>=<version> with an expiry time of one minute
	setVersion := func() {
		conn := c.pool.Get()
		if _, err := conn.Do("SET", c.serviceName, corelib.Version(), "EX", 60); err != nil {
			c.Logger.Warn("cache server is offline", zap.Error(err), zap.String("server", c.URI))
		}
		_ = conn.Close()
	}

	// set version on a schedule
	go func() {
		setVersion()
		for {
			select {
			case <-stopCh:
				return
			case <-ticker.C:
				setVersion()
			}
		}
	}()
}

// getCacheConn returns a Redis connection and an error. It parses the Redis URI, extracts the username and password
// (if present), and uses them to create a Redis connection with the `redis.Dial` function. The
// connection is returned to the caller along with any errors that occurred during the process.
func (c *Client) getCacheConn() (redis.Conn, error) {
	redisUrl, err := url.Parse(c.URI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %v", err)
	}

	var opts []redis.DialOption
	if user := redisUrl.User; user != nil {
		opts = append(opts, redis.DialUsername(user.Username()))
		if password, ok := user.Password(); ok {
			opts = append(opts, redis.DialPassword(password))
		}
	}

	if c.tlsEnabled {
		opts = append(opts, redis.DialUseTLS(true))
	}

	return redis.Dial("tcp", redisUrl.Host, opts...)
}
