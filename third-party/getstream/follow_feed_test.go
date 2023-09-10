// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream

import (
	"context"
	"testing"

	"github.com/GetStream/stream-go2/v7"
)

func TestClient_FollowFeed(t *testing.T) {
	sourceFlatFeed := generateFlatFeed(t, generateRandomString(5))
	targetFlatFeed := generateFlatFeed(t, generateRandomString(5))

	type args struct {
		ctx               context.Context
		sourceFeedID      string
		targetFeedID      string
		activityCopyLimit []stream.FollowFeedOption
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - follow feed",
			args: args{
				ctx:          context.Background(),
				sourceFeedID: sourceFlatFeed.ID(),
				targetFeedID: targetFlatFeed.ID(),
				activityCopyLimit: []stream.FollowFeedOption{
					stream.WithFollowFeedActivityCopyLimit(50),
				},
			},
			f:       testClient,
			wantErr: false,
		},
		{
			name: "fail - follow feed",
			args: args{
				ctx:               context.Background(),
				sourceFeedID:      "",
				targetFeedID:      "",
				activityCopyLimit: []stream.FollowFeedOption{},
			},
			f:       testClient,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.f.FollowFeed(tt.args.ctx, tt.args.sourceFeedID, tt.args.targetFeedID, tt.args.activityCopyLimit); (err != nil) != tt.wantErr {
				t.Errorf("Client.FollowFeed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
