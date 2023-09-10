package redis

type DatastoreTxName string

// String is a method defined on the `DatastoreTxName` type. It
// returns a string representation of the `DatastoreTxName` value. This method is used to convert the
// `DatastoreTxName` value to a string when it needs to be printed or displayed.
func (d DatastoreTxName) String() string {
	return string(d)
}

const (
	// Defining a constant named `RedisWriteToCacheTxn` of type `DatastoreTxName` and assigning it the
	// value `"txn.redis.write-to-cache"`. This constant is used as a transaction name or identifier for
	// write operations on a Redis cache in a larger codebase.
	RedisWriteToCacheTxn DatastoreTxName = "txn.redis.write-to-cache"
	// `RedisReadFromCacheTxn` of type
	// `DatastoreTxName` and assigning it the value `"txn.redis.read-from-cache"`. This constant is
	// used as a transaction name or identifier for read operations on a Redis cache in a larger codebase.
	RedisReadFromCacheTxn DatastoreTxName = "txn.redis.read-from-cache"
	// Defining a constant named `RedisReadManyFromCacheTxn` of type `DatastoreTxName` and assigning it the
	// value `"txn.redis.read-many-from-cache"`. This constant is used as a transaction name or identifier
	// for read operations on multiple keys from a Redis cache in a larger codebase.
	RedisReadManyFromCacheTxn DatastoreTxName = "txn.redis.read-many-from-cache"
	// Defining a constant named `RedisDeleteFromCacheTxn` of type `DatastoreTxName` and assigning it the
	// value `"txn.redis.delete-from-cache"`. This constant is used as a transaction name or identifier for
	// delete operations on a Redis cache in a larger codebase.
	RedisDeleteFromCacheTxn DatastoreTxName = "txn.redis.delete-from-cache"
)
