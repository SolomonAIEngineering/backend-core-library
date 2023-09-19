// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream // import "github.com/SolomonAIEngineering/backend-core-library/third-party/getstream"

import (
	"context"

	"github.com/GetStream/stream-go2/v7"
)

type ActivityResponse struct {
}

// This is a method defined on a struct type `Client`. The method is named
// `SendFollowRequestNotification` and it takes three parameters: a context object `ctx`, a string
// `notificationFeedID`, and a pointer to a `FollowRequestActivity` struct `params`. The method returns
// an error.
func (f *Client) SendFollowRequestNotification(ctx context.Context, notificationFeedID string, params *FollowRequestActivity) (*stream.AddActivityResponse, error) {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.send_follow_request_notification")
	defer span.End()

	notificationFeed, err := f.GetNotificationFeedFromFeedID(&notificationFeedID)
	if err != nil {
		return nil, err
	}

	activity := &stream.Activity{
		Actor:     params.SourceActor,
		Verb:      params.ActionName,
		Object:    params.FollowRequestRecordID,
		ForeignID: params.ActivityForeignID,
		Time:      stream.Time{Time: params.Time},
	}

	return notificationFeed.AddActivity(ctx, *activity)
}
