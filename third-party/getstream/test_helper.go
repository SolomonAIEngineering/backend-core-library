package getstream

import (
	"testing"
	"time"

	"github.com/GetStream/stream-go2/v7"
	"github.com/stretchr/testify/require"
)

func generateRandomUserIds(numUserIdsToGenerate int) []string {
	userIds := make([]string, 0, numUserIdsToGenerate)
	for i := 0; i < numUserIdsToGenerate; i++ {
		userIds = append(userIds, generateRandomString(10))
	}

	return userIds
}

// generateManyFlatFeeds  generates a set of flat feeds for target feed
func generateManyFlatFeeds(t *testing.T, userIds []string) []stream.Feed {
	flatFeeds := make([]stream.Feed, 0, len(userIds))
	for _, id := range userIds {
		// create flat feed
		singularFlatFeed, err := newFlatFeedWithUserID(testClient.client, id)
		require.NoError(t, err)

		flatFeeds = append(flatFeeds, singularFlatFeed)
	}

	return flatFeeds
}

// generateManyAggregatedFeeds generates a set of aggregated feeds for target feed
func generateManyAggregatedFeeds(t *testing.T, userIds []string) []*stream.AggregatedFeed {
	aggregatedFeeds := make([]*stream.AggregatedFeed, 0, len(userIds))
	for _, id := range userIds {
		// create aggregated flat feed
		singularAggregatedFeed, err := newAggregatedFeedWithUserID(testClient.client, id)
		require.NoError(t, err)

		aggregatedFeeds = append(aggregatedFeeds, singularAggregatedFeed)
	}

	return aggregatedFeeds
}

// generates a random flat feed used for testing
func generateFlatFeed(t *testing.T, userId string) *stream.FlatFeed {
	flatfeed, err := newFlatFeedWithUserID(testClient.client, userId)
	require.NoError(t, err)

	return flatfeed
}

// generateNotificationFeed generates a notification flat feed used for testing
func generateNotificationFeed(t *testing.T, userId string) *stream.NotificationFeed {
	notificationFeed, err := newNotificationFeedWithUserID(testClient.client, userId)
	require.NoError(t, err)

	return notificationFeed
}

// generateRandomActivities generates a random set of activities
func generateRandomActivities(numActivities int) []*stream.Activity {
	activitySet := make([]*stream.Activity, 0, numActivities)
	for i := 0; i < numActivities; i++ {
		activitySet = append(activitySet, generateRandomActivity())
	}

	return activitySet
}

// generates a random activity
func generateRandomActivity() *stream.Activity {
	return &stream.Activity{
		ID:        generateRandomString(20),
		Actor:     generateRandomString(20),
		Verb:      generateRandomString(20),
		Object:    generateRandomString(20),
		ForeignID: generateRandomString(20),
		Target:    generateRandomString(20),
		Time: stream.Time{
			Time: time.Now(),
		},
		Origin: generateRandomString(20),
		To: []string{
			generateRandomString(10),
			generateRandomString(10),
			generateRandomString(10),
		},
		Score: float64(generateRandomId(10, 100)),
		Extra: map[string]interface{}{
			"test-detail-1": generateRandomString(10),
			"test-detail-2": generateRandomString(10),
			"test-detail-3": generateRandomString(10),
			"test-detail-4": generateRandomString(10),
			"test-detail-5": generateRandomString(10),
			"test-detail-6": generateRandomString(10),
		},
	}
}
