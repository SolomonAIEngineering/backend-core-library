// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream

import (
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/stretchr/testify/assert"

	"github.com/GetStream/stream-go2/v7"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	type args struct {
		client  *stream.Client
		logger  *zap.Logger
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name: "pass - test valid client creation",
			args: args{
				options: []Option{
					WithLogger(testClient.logger),
					WithKey(generateRandomString(10)),
					WithSecret(generateRandomString(10)),
					WithInstrumentationClient(&instrumentation.Client{}),
				},
			},
			wantErr: false,
			want:    testClient,
		},
		{
			name: "fail - test invalid client creation",
			args: args{
				options: []Option{
					WithKey(generateRandomString(10)),
					WithSecret(generateRandomString(10)),
					WithInstrumentationClient(&instrumentation.Client{}),
				},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "fail - test invalid client",
			args: args{
				options: []Option{
					WithKey(generateRandomString(10)),
					WithInstrumentationClient(&instrumentation.Client{}),
				},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "fail - test invalid logger",
			args: args{
				options: []Option{
					WithKey(generateRandomString(10)),
					WithInstrumentationClient(&instrumentation.Client{}),
				},
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestClient_GetFlatFeedFromFeedID(t *testing.T) {
	userId := generateRandomString(10)
	flatFeed := generateFlatFeed(t, userId)
	flatFeedId := flatFeed.ID()

	type args struct {
		feedID *string
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		want    *stream.FlatFeed
		wantErr bool
	}{
		{
			name: "pass - get a valid flat feed",
			f:    testClient,
			args: args{
				feedID: &flatFeedId,
			},
			want:    flatFeed,
			wantErr: false,
		},
		{
			name: "fail - invalid feed id",
			f:    testClient,
			args: args{
				feedID: nil,
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetFlatFeedFromFeedID(tt.args.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetFlatFeedFromFeedID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetFlatFeedFromFeedID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNotificationFeedFromFeedID(t *testing.T) {
	userId := generateRandomString(10)
	notificationFeed := generateNotificationFeed(t, userId)
	notificationFeedId := notificationFeed.ID()

	type args struct {
		feedID *string
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		want    *stream.NotificationFeed
		wantErr bool
	}{
		{
			name: "pass - get a valid flat feed",
			f:    testClient,
			args: args{
				feedID: &notificationFeedId,
			},
			want:    notificationFeed,
			wantErr: false,
		},
		{
			name: "fail - invalid feed id",
			f:    testClient,
			args: args{
				feedID: nil,
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetNotificationFeedFromFeedID(tt.args.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNotificationFeedFromFeedID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetNotificationFeedFromFeedID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_validateFeedID(t *testing.T) {
	id := generateRandomString(10)
	type args struct {
		feedID *string
	}
	tests := []struct {
		name    string
		f       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - validate feed id",
			f:    testClient,
			args: args{
				feedID: &id,
			},
			wantErr: false,
		},
		{
			name: "fail - invalid feed id",
			f:    testClient,
			args: args{
				feedID: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.validateFeedID(tt.args.feedID); (err != nil) != tt.wantErr {
				t.Errorf("Client.validateFeedID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateFlatFeed(t *testing.T) {
	feedType := PersonalFeed
	feedSlug := generateRandomString(10)
	flatFeed, err := testClient.client.FlatFeed(feedType.String(), feedSlug)
	assert.NoError(t, err)

	type args struct {
		feedType FeedType
		feedSlug string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *stream.FlatFeed
		wantErr bool
	}{
		{
			name: "pass - create flat feed",
			c:    testClient,
			args: args{
				feedType: feedType,
				feedSlug: feedSlug,
			},
			wantErr: false,
			want:    flatFeed,
		},
		{
			name: "fail - create an invalid flat feed",
			c:    testClient,
			args: args{
				feedType: feedType,
				feedSlug: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail - create feedslug flat feed",
			c:    testClient,
			args: args{
				feedType: "",
				feedSlug: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateFlatFeed(tt.args.feedType, tt.args.feedSlug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateFlatFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateFlatFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateAggregateFeed(t *testing.T) {
	feedType := PersonalFeed
	feedSlug := generateRandomString(10)
	aggregatedFeed, err := testClient.client.AggregatedFeed(feedType.String(), feedSlug)
	assert.NoError(t, err)

	type args struct {
		feedType FeedType
		feedSlug string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *stream.AggregatedFeed
		wantErr bool
	}{
		{
			name: "pass - create aggregated feed",
			c:    testClient,
			args: args{
				feedType: feedType,
				feedSlug: feedSlug,
			},
			wantErr: false,
			want:    aggregatedFeed,
		},
		{
			name: "fail - create an invalid flat feed",
			c:    testClient,
			args: args{
				feedType: "",
				feedSlug: feedSlug,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail - create feedslug flat feed",
			c:    testClient,
			args: args{
				feedType: "",
				feedSlug: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateAggregateFeed(tt.args.feedType, tt.args.feedSlug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateAggregateFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateAggregateFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateNotificationFeed(t *testing.T) {
	feedType := PersonalFeed
	feedSlug := generateRandomString(10)
	notificationFeed, err := testClient.client.NotificationFeed(feedType.String(), feedSlug)
	assert.NoError(t, err)

	type args struct {
		feedType FeedType
		feedSlug string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *stream.NotificationFeed
		wantErr bool
	}{
		{
			name: "pass - create aggregated feed",
			c:    testClient,
			args: args{
				feedType: feedType,
				feedSlug: feedSlug,
			},
			wantErr: false,
			want:    notificationFeed,
		},
		{
			name: "fail - create an invalid flat feed",
			c:    testClient,
			args: args{
				feedType: "",
				feedSlug: feedSlug,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail - create feedslug flat feed",
			c:    testClient,
			args: args{
				feedType: "",
				feedSlug: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateNotificationFeed(tt.args.feedType, tt.args.feedSlug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateNotificationFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateNotificationFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
