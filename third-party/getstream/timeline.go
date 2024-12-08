// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream // import "github.com/SolomonAIEngineering/backend-core-library/third-party/getstream"

import (
	"context"
	"fmt"

	"github.com/GetStream/stream-go2/v7"
)

// This is a method defined on the `Client` struct that retrieves a timeline of activities from a flat
// feed identified by the `feedID` parameter. It returns a slice of `stream.Activity` objects and an
// error if there was a problem retrieving the activities. The method takes a `context.Context` object
// as the first parameter to allow for cancellation or timeout of the request.
func (f *Client) GetTimeline(ctx context.Context, feedID *string) (*stream.FlatFeedResponse, error) {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.get_timeline")
	defer span.End()

	if feedID == nil {
		return nil, fmt.Errorf("invalid input argument. feedId: %v", feedID)
	}

	feed, err := f.GetFlatFeedFromFeedID(feedID)
	if err != nil {
		return nil, err
	}

	res, err := feed.GetActivities(ctx, stream.WithActivitiesLimit(200))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (f *Client) GetTimelineNextPage(ctx context.Context, feedID *string, opts []stream.GetActivitiesOption) (*stream.FlatFeedResponse, error) {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.get_timeline")
	defer span.End()

	if feedID == nil {
		return nil, fmt.Errorf("invalid input argument. feedId: %v", feedID)
	}

	feed, err := f.GetFlatFeedFromFeedID(feedID)
	if err != nil {
		return nil, err
	}

	res, err := feed.GetActivities(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// `func (f *Client) GetNotificationTimeline(ctx context.Context, feedID *string)
// ([]stream.NotificationFeedResult, error)` is a method defined on the `Client` struct that retrieves
// a timeline of notification activities from a notification feed identified by the `feedID` parameter.
// It returns a slice of `stream.NotificationFeedResult` objects and an error if there was a problem
// retrieving the activities. The method takes a `context.Context` object as the first parameter to
// allow for cancellation or timeout of the request.
func (f *Client) GetNotificationTimeline(ctx context.Context, feedID *string, opts []stream.GetActivitiesOption) (*stream.NotificationFeedResponse, error) {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.get_notification_timeline")
	defer span.End()

	if feedID == nil {
		return nil, fmt.Errorf("invalid input argument. feedId: %v", feedID)
	}

	feed, err := f.GetNotificationFeedFromFeedID(feedID)
	if err != nil {
		return nil, err
	}

	return feed.GetActivities(ctx, opts...)
}
