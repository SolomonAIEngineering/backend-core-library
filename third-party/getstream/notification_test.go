// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream

import (
	"context"
	"testing"
	"time"
)

func TestClient_SendFollowRequestNotification(t *testing.T) {
	userId := generateRandomString(10)
	notificationFeed := generateNotificationFeed(t, userId)

	type args struct {
		ctx                context.Context
		notificationFeedID string
		params             *FollowRequestActivity
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - send proper notification",
			f:    testClient,
			args: args{
				ctx:                context.Background(),
				notificationFeedID: notificationFeed.ID(),
				params: &FollowRequestActivity{
					SourceActor:           generateRandomString(10),
					ActionName:            generateRandomString(10),
					FollowRequestRecordID: generateRandomString(5),
					Time:                  time.Time{},
					ActivityForeignID:     generateRandomString(5),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.f.SendFollowRequestNotification(tt.args.ctx, tt.args.notificationFeedID, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Client.SendFollowRequestNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
