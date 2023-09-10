// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream

import (
	"context"
	"fmt"

	"github.com/GetStream/stream-go2/v7"
)

// This is a method of the `Client` struct that adds an activity to a feed specified by `feedID`. It
// takes in a `context.Context` object, a pointer to a string `feedID`, and a pointer to a
// `stream.Activity` object. It returns an error if there is an issue with the input arguments or if
// there is an error adding the activity to the feed. The method first checks if the input arguments
// are valid, gets the flat feed from the feed ID, adds the activity to the feed, and logs the creation
// of the activity.
func (f *Client) AddActivity(ctx context.Context, feedID *string, activity *stream.Activity) (*stream.AddActivityResponse, error) {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.add_activity")
	defer span.End()

	// TODO: instrument this
	if activity == nil {
		return nil, fmt.Errorf("invalid input argument. activity: %v", activity)
	}

	if feedID == nil {
		return nil, fmt.Errorf("invalid input argument. feedId: %v", feedID)
	}

	feed, err := f.GetFlatFeedFromFeedID(feedID)
	if err != nil {
		return nil, err
	}

	return feed.AddActivity(ctx, *activity)
}

// The `DeleteActivity` function is a method of the `Client` struct that deletes an activity from a
// feed specified by `feedID` and `activityForeignID`. It takes in a `context.Context` object, a
// pointer to a string `feedID`, and a pointer to a string `activityForeignID`. It returns an error if
// there is an issue with the input arguments or if there is an error deleting the activity from the
// feed. The method first checks if the input arguments are valid, gets the flat feed from the feed ID,
// removes the activity from the feed using the foreign ID, and logs the removal of the activity.
func (f *Client) DeleteActivity(ctx context.Context, feedID *string, activityForeignID *string) (*stream.RemoveActivityResponse, error) {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.delete_activity")
	defer span.End()

	if activityForeignID == nil {
		err := fmt.Errorf("invalid input argument. postID: %d", activityForeignID)
		return nil, err
	}

	if feedID == nil {
		err := fmt.Errorf("invalid input argument. feedId: %v", feedID)
		return nil, err
	}

	feed, err := f.GetFlatFeedFromFeedID(feedID)
	if err != nil {
		return nil, err
	}

	return feed.RemoveActivityByForeignID(ctx, *activityForeignID)
}

// `func (f *Client) AddActivityToManyFeeds(ctx context.Context, activity *stream.Activity, feeds
// ...stream.Feed) error` is a method of the `Client` struct that adds an activity to multiple feeds
// specified by `feeds`. It takes in a `context.Context` object, a pointer to a `stream.Activity`
// object, and a variadic argument of type `stream.Feed`. It returns an error if there is an issue with
// the input arguments or if there is an error adding the activity to the feeds. The method first
// checks if the input arguments are valid, and then adds the activity to all the specified feeds using
// the `AddToMany` method of the `stream.Client` object.
func (f *Client) AddActivityToManyFeeds(ctx context.Context, activity *stream.Activity, feeds ...stream.Feed) error {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.add_activity_to_many_feeds")
	defer span.End()

	if activity == nil {
		return fmt.Errorf("invalid activity object. activity cannot be nil")
	}

	if len(feeds) == 0 {
		return fmt.Errorf("invalid feed set. target feeds cannot be empty")
	}

	return f.client.AddToMany(ctx, *activity, feeds...)
}
