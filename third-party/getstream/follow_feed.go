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

// This is a method defined on a struct type `Client`. The method is called `FollowFeed` and it takes
// in four arguments: a context, two strings `sourceFeedID` and `targetFeedID`, and a slice of
// `stream.FollowFeedOption` options. The method returns an error.
func (f *Client) FollowFeed(ctx context.Context, sourceFeedID, targetFeedID string, opts []stream.FollowFeedOption) (*stream.BaseResponse, error) {
	txn := f.instrumentationClient.GetTraceFromContext(ctx)
	span := f.instrumentationClient.StartSegment(txn, "getstream.follow_feed")
	defer span.End()

	if sourceFeedID == "" || targetFeedID == "" {
		return nil, fmt.Errorf("invalid input argument. sourceFeedID: %s, targetFeedID: %s", sourceFeedID, targetFeedID)
	}

	sourceFeed, err := f.GetFlatFeedFromFeedID(&sourceFeedID)
	if err != nil {
		return nil, err
	}

	targetFeed, err := f.GetFlatFeedFromFeedID(&targetFeedID)
	if err != nil {
		return nil, err
	}

	return sourceFeed.Follow(ctx, targetFeed, opts...)
}
