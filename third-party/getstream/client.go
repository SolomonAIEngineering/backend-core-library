package getstream

import (
	"context"

	"github.com/GetStream/stream-go2/v7"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"go.uber.org/zap"
)

// This is defining a struct type called `Client` which will be used to create instances of the
// GetStream client. The struct contains fields for the GetStream API key, secret, a pointer to a
// `stream.Client` instance, a logger, and an instrumentation client.
type Client struct {
	// `getstreamKey` is a field in the `Client` struct that stores the GetStream API key as a string. It
	// is used to authenticate requests to the GetStream API.
	getstreamKey string
	// `getstreamSecret` is a field in the `Client` struct that stores the GetStream API secret as a
	// string. It is used to authenticate requests to the GetStream API along with the API key.
	getstreamSecret string
	// `client                *stream.Client` is a field in the `Client` struct that stores a pointer to a
	// `stream.Client` instance. This instance is used to make requests to the GetStream API.
	client *stream.Client
	// `logger *zap.Logger` is a field in the `Client` struct that stores a pointer to a `zap.Logger`
	// instance. This instance is used for logging messages related to the GetStream client. `zap` is a
	// popular logging library in the Go programming language.
	logger *zap.Logger
	// `instrumentationClient *instrumentation.Client` is a field in the `Client` struct that stores a
	// pointer to an `instrumentation.Client` instance. This instance is used for instrumenting the
	// GetStream client, which means collecting and analyzing data related to the performance and behavior
	// of the client. This can help identify issues and optimize the client's performance.
	instrumentationClient *instrumentation.Client
}

// `type IClient interface` is defining an interface called `IClient` which specifies a set of methods
// that a type must implement in order to be considered a valid implementation of the interface. In
// this case, the `IClient` interface specifies a set of methods that the `Client` struct must
// implement in order to be considered a valid implementation of the interface. This allows other parts
// of the code to interact with the `Client` struct through the `IClient` interface, which can make the
// code more modular and easier to test.
type IClient interface {
	// `GetFlatFeedFromFeedID(feedID *string) (*stream.FlatFeed, error)` is a method defined in the
	// `IClient` interface and implemented in the `Client` struct. It takes a pointer to a string `feedID`
	// as input and returns a pointer to a `stream.FlatFeed` and an error. This method is used to retrieve
	// a flat feed from the GetStream API based on the provided `feedID`. A flat feed is a feed that
	// contains all activities in chronological order, without any grouping or aggregation. The method
	// returns a pointer to the retrieved `stream.FlatFeed` object and an error if there was an issue
	// retrieving the feed.
	GetFlatFeedFromFeedID(feedID *string) (*stream.FlatFeed, error)
	// `GetNotificationFeedFromFeedID(feedID *string) (*stream.NotificationFeed, error)` is a method
	// defined in the `IClient` interface and implemented in the `Client` struct. It takes a pointer to a
	// string `feedID` as input and returns a pointer to a `stream.NotificationFeed` and an error. This
	// method is used to retrieve a notification feed from the GetStream API based on the provided
	// `feedID`. A notification feed is a feed that contains activities related to notifications, such as
	// when a user is mentioned or receives a direct message. The method returns a pointer to the retrieved
	// `stream.NotificationFeed` object and an error if there was an issue retrieving the feed.
	GetNotificationFeedFromFeedID(feedID *string) (*stream.NotificationFeed, error)
	// `CreateFlatFeed` is a method defined in the `IClient` interface and implemented in the `Client`
	// struct. It takes a `FeedType` and a `feedSlug` as input and returns a pointer to a `stream.FlatFeed`
	// and an error. This method is used to create a new flat feed on the GetStream API with the specified
	// `feedType` and `feedSlug`. A flat feed is a feed that contains all activities in chronological
	// order, without any grouping or aggregation. The method returns a pointer to the newly created
	// `stream.FlatFeed` object and an error if there was an issue creating the feed.
	CreateFlatFeed(feedType FeedType, feedSlug string) (*stream.FlatFeed, error)
	// `CreateAggregateFeed` is a method defined in the `IClient` interface and implemented in the `Client`
	// struct. It takes a `FeedType` and a `feedSlug` as input and returns a pointer to a
	// `stream.AggregatedFeed` and an error. This method is used to create a new aggregated feed on the
	// GetStream API with the specified `feedType` and `feedSlug`. An aggregated feed is a feed that groups
	// activities based on certain criteria, such as time or user, and presents them in a summarized
	// format. The method returns a pointer to the newly created `stream.AggregatedFeed` object and an
	// error if there was an issue creating the feed.
	CreateAggregateFeed(feedType FeedType, feedSlug string) (*stream.AggregatedFeed, error)
	// `CreateNotificationFeed` is a method defined in the `IClient` interface and implemented in the
	// `Client` struct. It takes a `FeedType` and a `feedSlug` as input and returns a pointer to a
	// `stream.NotificationFeed` and an error. This method is used to create a new notification feed on the
	// GetStream API with the specified `feedType` and `feedSlug`. A notification feed is a feed that
	// contains activities related to notifications, such as when a user is mentioned or receives a direct
	// message. The method returns a pointer to the newly created `stream.NotificationFeed` object and an
	// error if there was an issue creating the feed.
	CreateNotificationFeed(feedType FeedType, feedSlug string) (*stream.NotificationFeed, error)
	// `AddActivity` is a method defined in the `IClient` interface and implemented in the `Client` struct.
	// It takes a context, a pointer to a string `feedID`, and a pointer to a `stream.Activity` as input
	// and returns a pointer to a `stream.AddActivityResponse` and an error. This method is used to add a
	// new activity to the feed specified by the `feedID`. The `stream.Activity` parameter contains the
	// data for the new activity. The method returns a pointer to a `stream.AddActivityResponse` object and
	// an error if there was an issue adding the activity.
	AddActivity(ctx context.Context, feedID *string, activity *stream.Activity) (*stream.AddActivityResponse, error)
	// The `DeleteActivity` method is defined in the `IClient` interface and implemented in the `Client`
	// struct. It takes a context, a pointer to a string `feedID`, and a pointer to a string
	// `activityForeignID` as input and returns a pointer to a `stream.RemoveActivityResponse` and an
	// error. This method is used to remove an activity from the feed specified by the `feedID` and
	// `activityForeignID`. The `activityForeignID` parameter is a unique identifier for the activity that
	// was assigned by the client when the activity was added to the feed. The method returns a pointer to
	// a `stream.RemoveActivityResponse` object and an error if there was an issue removing the activity.
	DeleteActivity(ctx context.Context, feedID *string, activityForeignID *string) (*stream.RemoveActivityResponse, error)
	// `AddActivityToManyFeeds` is a method defined in the `IClient` interface and implemented in the
	// `Client` struct. It takes a context, a pointer to a `stream.Activity`, and a variable number of
	// `stream.Feed` objects as input and returns an error. This method is used to add a new activity to
	// multiple feeds at once. The `stream.Activity` parameter contains the data for the new activity, and
	// the `stream.Feed` parameters specify the feeds to which the activity should be added. The method
	// returns an error if there was an issue adding the activity to any of the specified feeds.
	AddActivityToManyFeeds(ctx context.Context, activity *stream.Activity, feeds ...stream.Feed) error
	// `FollowFeed` is a method defined in the `IClient` interface and implemented in the `Client` struct.
	// It takes a context, a source feed ID, a target feed ID, and an optional list of
	// `stream.FollowFeedOption` objects as input. This method is used to follow a target feed from a
	// source feed. The `sourceFeedID` parameter specifies the ID of the feed that will follow the target
	// feed, and the `targetFeedID` parameter specifies the ID of the feed that will be followed. The
	// `opts` parameter is an optional list of `stream.FollowFeedOption` objects that can be used to
	// specify additional options for the follow operation, such as whether to copy existing activities
	// from the target feed to the source feed. The method returns a pointer to a `stream.BaseResponse`
	// object and an error if there was an issue following the target feed.
	FollowFeed(ctx context.Context, sourceFeedID, targetFeedID string, opts []stream.FollowFeedOption) (*stream.BaseResponse, error)
	// `SendFollowRequestNotification` is a method defined in the `IClient` interface and implemented in
	// the `Client` struct. It takes a context, a string `notificationFeedID`, and a pointer to a
	// `FollowRequestActivity` object as input and returns a pointer to a `stream.AddActivityResponse`
	// object and an error. This method is used to send a follow request notification to the specified
	// `notificationFeedID`. The `FollowRequestActivity` parameter contains the data for the follow request
	// activity. The method returns a pointer to a `stream.AddActivityResponse` object and an error if
	// there was an issue sending the notification.
	SendFollowRequestNotification(ctx context.Context, notificationFeedID string, params *FollowRequestActivity) (*stream.AddActivityResponse, error)
	// `GetTimeline` is a method defined in the `IClient` interface and implemented in the `Client` struct.
	// It takes a context and a pointer to a string `feedID` as input and returns a pointer to a
	// `stream.FlatFeedResponse` and an error. This method is used to retrieve a timeline feed from the
	// GetStream API based on the provided `feedID`. A timeline feed is a feed that contains activities
	// from the user's followers in reverse chronological order. The method returns a pointer to the
	// retrieved `stream.FlatFeedResponse` object and an error if there was an issue retrieving the feed.
	GetTimeline(ctx context.Context, feedID *string) (*stream.FlatFeedResponse, error)
	// `GetNotificationTimeline` is a method defined in the `IClient` interface and implemented in the
	// `Client` struct. It takes a context and a pointer to a string `feedID` as input and returns a
	// pointer to a `stream.NotificationFeedResponse` and an error. This method is used to retrieve a
	// notification timeline feed from the GetStream API based on the provided `feedID`. A notification
	// timeline feed is a feed that contains activities related to notifications, such as when a user is
	// mentioned or receives a direct message, from the user's followers in reverse chronological order.
	// The method returns a pointer to the retrieved `stream.NotificationFeedResponse` object and an error
	// if there was an issue retrieving the feed.
	GetNotificationTimeline(ctx context.Context, feedID *string, opts []stream.GetActivitiesOption) (*stream.NotificationFeedResponse, error)

	// The above code is defining a method called `GetTimelineWithOptions` in the Go programming language.
	// This method takes in a context object, a pointer to a string representing a feed ID, and an array of
	// options for getting activities. It returns a `FlatFeedResponse` object and an error. This method is
	// likely used to retrieve a timeline of activities from a feed with the specified ID and options.
	GetTimelineNextPage(ctx context.Context, feedID *string, opts []stream.GetActivitiesOption) (*stream.FlatFeedResponse, error)
}

var _ IClient = (*Client)(nil)

// The function creates a new client with optional configuration options.
func New(opts ...Option) (*Client, error) {
	c := &Client{}
	c.ApplyOptions(opts...)
	if err := c.Validate(); err != nil {
		return nil, err
	}

	streamClient, err := stream.New(c.getstreamKey, c.getstreamSecret)
	if err != nil {
		return nil, err
	}

	c.client = streamClient
	return c, nil
}
