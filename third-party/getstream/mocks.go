package getstream

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/GetStream/stream-go2/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var (
	testClient    *Client
	mockRequestor *mockRequester
)

func init() {
	rand.Seed(time.Now().UnixNano())
	testClient, mockRequestor = newClient()
}

type mockRequester struct {
	req  *http.Request
	resp string
}

func (m *mockRequester) Do(req *http.Request) (*http.Response, error) {
	m.req = req
	body := "{}"
	if m.resp != "" {
		body = m.resp
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
	}, nil
}

// newClient creates a new test client for the various unit test in this package
func newClient() (*Client, *mockRequester) {
	requester := &mockRequester{}
	c, _ := stream.New("key", "secret", stream.WithHTTPRequester(requester))
	return &Client{
		client: c,
		logger: zap.L(),
	}, requester
}

func testRequest(t *testing.T, req *http.Request, method, url, body string) {
	assert.Equal(t, url, req.URL.String())
	assert.Equal(t, method, req.Method)
	if req.Method == http.MethodPost {
		reqBody, err := io.ReadAll(req.Body)
		require.NoError(t, err)
		assert.JSONEq(t, body, string(reqBody))
	}
	headers := req.Header
	if headers.Get("X-API-Key") == "" {
		assert.NotEmpty(t, headers.Get("Stream-Auth-Type"))
		assert.NotEmpty(t, headers.Get("Authorization"))
	}
}

func getTime(t time.Time) stream.Time {
	st, _ := time.Parse(stream.TimeLayout, t.Truncate(time.Second).Format(stream.TimeLayout))
	return stream.Time{Time: st}
}

func newFlatFeedWithUserID(c *stream.Client, userID string) (*stream.FlatFeed, error) {
	return c.FlatFeed("flat", userID)
}

func newAggregatedFeedWithUserID(c *stream.Client, userID string) (*stream.AggregatedFeed, error) {
	return c.AggregatedFeed("aggregated", userID)
}

func newNotificationFeedWithUserID(c *stream.Client, userID string) (*stream.NotificationFeed, error) {
	return c.NotificationFeed("notification", userID)
}
