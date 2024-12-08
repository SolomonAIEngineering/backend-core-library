// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/GetStream/stream-go2/v7"
)

func TestClient_AddActivity(t *testing.T) {
	userId := generateRandomString(10)
	flatFeed := generateFlatFeed(t, userId)
	feedId := flatFeed.ID()
	activity := &stream.Activity{
		ID:        "123",
		Actor:     "Test User",
		Verb:      "Disklike",
		Object:    "ice-cream",
		ForeignID: "xxx-111-xxx",
		Target:    "",
		Time:      stream.Time{},
		Origin:    "",
		To:        []string{"flat:456"},
		Score:     100,
		Extra:     nil,
	}
	type args struct {
		ctx      context.Context
		feedID   *string
		activity *stream.Activity
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - test add activity",
			f:    testClient,
			args: args{
				ctx:      context.Background(),
				feedID:   &feedId,
				activity: activity,
			},
			wantErr: false,
		},
		{
			name: "fail - test invalid activity",
			f:    testClient,
			args: args{
				ctx:      context.Background(),
				feedID:   &feedId,
				activity: nil,
			},
			wantErr: true,
		},
		{
			name: "fail - test invalid feed id",
			f:    testClient,
			args: args{
				ctx:      context.Background(),
				feedID:   nil,
				activity: activity,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.f.AddActivity(tt.args.ctx, tt.args.feedID, tt.args.activity); (err != nil) != tt.wantErr {
				t.Errorf("Client.AddActivity() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				body := `{"actor":"Test User","foreign_id":"xxx-111-xxx","id":"123","object":"ice-cream","score":100,"to":["flat:456"],"verb":"Disklike"}`
				details := strings.Split(*tt.args.feedID, ":")
				testRequest(
					t,
					mockRequestor.req,
					http.MethodPost,
					fmt.Sprintf("https://api.stream-io-api.com/api/v1.0/feed/flat/%s/?api_key=key", details[1]),
					body)
			}
		})
	}
}

func TestClient_DeleteActivity(t *testing.T) {
	userId := generateRandomString(10)
	flatFeed := generateFlatFeed(t, userId)
	feedId := flatFeed.ID()
	activityForeignId := generateRandomString(10)
	randomId := generateRandomString(10)

	type args struct {
		ctx               context.Context
		feedID            *string
		activityForeignID *string
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - delete a valid activity",
			f:    testClient,
			args: args{
				ctx:               context.Background(),
				feedID:            &feedId,
				activityForeignID: &activityForeignId,
			},
			wantErr: false,
		},
		{
			name: "fail - delete an activity with no feedId",
			f:    testClient,
			args: args{
				ctx:               context.Background(),
				feedID:            nil,
				activityForeignID: &activityForeignId,
			},
			wantErr: true,
		},
		{
			name: "fail - delete with no foreign id",
			f:    testClient,
			args: args{
				ctx:               context.Background(),
				feedID:            &feedId,
				activityForeignID: nil,
			},
			wantErr: true,
		},
		{
			name: "fail - delete an account with invalid feedId",
			f:    testClient,
			args: args{
				ctx: context.Background(),
				// generate random feed
				feedID:            &randomId,
				activityForeignID: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.f.DeleteActivity(tt.args.ctx, tt.args.feedID, tt.args.activityForeignID); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteActivity() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				details := strings.Split(*tt.args.feedID, ":")
				uri := fmt.Sprintf("https://api.stream-io-api.com/api/v1.0/feed/flat/%s/%s/?api_key=key&foreign_id=1", details[1], *tt.args.activityForeignID)
				testRequest(t, mockRequestor.req, http.MethodDelete, uri, "")
			}
		})
	}
}

func TestClient_AddActivityToManyFeeds(t *testing.T) {
	activity := &stream.Activity{
		ID:        "123",
		Actor:     "bob",
		Verb:      "like",
		Object:    "cake",
		ForeignID: "xxx-111-xxx",
		Target:    "",
		Time:      stream.Time{},
		Origin:    "",
		Score:     100,
		Extra:     nil,
	}

	feeds := []stream.Feed{
		generateFlatFeed(t, "123"),
		generateFlatFeed(t, "456"),
		generateFlatFeed(t, "789"),
	}

	type args struct {
		ctx      context.Context
		activity *stream.Activity
		feeds    []stream.Feed
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - add activity to many feeds",
			args: args{
				ctx:      context.Background(),
				activity: activity,
				feeds:    feeds,
			},
			f:       testClient,
			wantErr: false,
		},
		{
			name: "fail - add invalid activity to many feeds",
			args: args{
				ctx:      context.Background(),
				activity: nil,
				feeds:    feeds,
			},
			f:       testClient,
			wantErr: true,
		},
		{
			name: "fail - add valid activity to empty feeds",
			args: args{
				ctx:      context.Background(),
				activity: activity,
				feeds:    nil,
			},
			f:       testClient,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.AddActivityToManyFeeds(tt.args.ctx, tt.args.activity, tt.args.feeds...); (err != nil) != tt.wantErr {
				t.Errorf("Client.AddActivityToManyFeeds() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				body := `{"activity":{"actor":"bob","foreign_id":"xxx-111-xxx","id":"123","object":"cake","score":100,"verb":"like"},"feeds":["flat:123","flat:456","flat:789"]}`
				testRequest(t, mockRequestor.req, http.MethodPost, "https://api.stream-io-api.com/api/v1.0/feed/add_to_many/?api_key=key", body)
			}
		})
	}
}
