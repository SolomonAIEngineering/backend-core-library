// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package getstream // import "github.com/SolomonAIEngineering/backend-core-library/third-party/getstream"

import (
	"fmt"
	"strings"

	"github.com/GetStream/stream-go2/v7"
)

type FeedType string

const (
	PersonalFeed     FeedType = "personal"
	TimelineFeed     FeedType = "timeline"
	NotificationFeed FeedType = "notification"
)

func (f FeedType) String() string {
	return string(f)
}

// `func (f *Client) GetFlatFeedFromFeedID(feedID *string) (*stream.FlatFeed, error)` is a method of
// the `Client` struct that takes a pointer to a string as input and returns a pointer to a
// `stream.FlatFeed` object and an error. It is used to retrieve a `FlatFeed` object from the GetStream
// API based on a given feed ID. The method first validates the feed ID using the `validateFeedID`
// method, and then splits the feed ID into its constituent parts (feed type and feed slug) using the
// `strings.Split` function. Finally, it calls the `FlatFeed` method of the `stream.Client` instance
// stored in the `Client` struct, passing in the feed type and feed slug as arguments, and returns the
// resulting `FlatFeed` object and any error that occurred during the API call.
func (f *Client) GetFlatFeedFromFeedID(feedID *string) (*stream.FlatFeed, error) {
	if err := f.validateFeedID(feedID); err != nil {
		return nil, err
	}

	details := strings.Split(*feedID, ":")
	return f.client.FlatFeed(details[0], details[1])
}

// `func (f *Client) GetNotificationFeedFromFeedID(feedID *string) (*stream.NotificationFeed, error)`
// is a method of the `Client` struct that takes a pointer to a string as input and returns a pointer
// to a `stream.NotificationFeed` object and an error. It is used to retrieve a `NotificationFeed`
// object from the GetStream API based on a given feed ID. The method first validates the feed ID using
// the `validateFeedID` method, and then splits the feed ID into its constituent parts (feed type and
// feed slug) using the `strings.Split` function. Finally, it calls the `NotificationFeed` method of
// the `stream.Client` instance stored in the `Client` struct, passing in the feed type and feed slug
// as arguments, and returns the resulting `NotificationFeed` object and any error that occurred during
// the API call.
func (f *Client) GetNotificationFeedFromFeedID(feedID *string) (*stream.NotificationFeed, error) {
	if err := f.validateFeedID(feedID); err != nil {
		return nil, err
	}

	details := strings.Split(*feedID, ":")
	return f.client.NotificationFeed(details[0], details[1])
}

// `func (f *Client) validateFeedID(feedID *string) error` is a method of the `Client` struct that
// takes a pointer to a string as input and returns an error. It is used to validate the feed ID passed
// as input to other methods of the `Client` struct. The method checks if the `Client` instance and the
// feed ID are not nil, and returns an error if either of them is nil. If both are not nil, the method
// returns nil, indicating that the feed ID is valid.
func (f *Client) validateFeedID(feedID *string) error {
	if f.client == nil || feedID == nil {
		return fmt.Errorf("invalid input argument. client: %v, feedId :%v", f.client, feedID)
	}
	return nil
}

// `func (c *Client) CreateFlatFeed(feedType, feedSlug string) (*stream.FlatFeed, error)` is a method
// of the `Client` struct that creates a new `FlatFeed` object in the GetStream API based on the given
// feed type and feed slug. It takes the feed type and feed slug as input arguments and returns a
// pointer to the newly created `FlatFeed` object and any error that occurred during the API call. The
// method first checks if the input arguments are valid, and then calls the `FlatFeed` method of the
// `stream.Client` instance stored in the `Client` struct, passing in the feed type and feed slug as
// arguments.
func (c *Client) CreateFlatFeed(feedType FeedType, feedSlug string) (*stream.FlatFeed, error) {
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.FlatFeed(feedType.String(), feedSlug)
}

// `func (c *Client) CreateAggregateFeed(feedType, feedSlug string) (*stream.AggregatedFeed, error)` is
// a method of the `Client` struct that creates a new `AggregatedFeed` object in the GetStream API
// based on the given feed type and feed slug. It takes the feed type and feed slug as input arguments
// and returns a pointer to the newly created `AggregatedFeed` object and any error that occurred
// during the API call. The method first checks if the input arguments are valid, and then calls the
// `AggregatedFeed` method of the `stream.Client` instance stored in the `Client` struct, passing in
// the feed type and feed slug as arguments.
func (c *Client) CreateAggregateFeed(feedType FeedType, feedSlug string) (*stream.AggregatedFeed, error) {
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.AggregatedFeed(feedType.String(), feedSlug)
}

// `func (c *Client) CreateNotificationFeed(feedType, feedSlug string) (*stream.NotificationFeed,
// error)` is a method of the `Client` struct that creates a new `NotificationFeed` object in the
// GetStream API based on the given feed type and feed slug. It takes the feed type and feed slug as
// input arguments and returns a pointer to the newly created `NotificationFeed` object and any error
// that occurred during the API call. The method first checks if the input arguments are valid, and
// then calls the `NotificationFeed` method of the `stream.Client` instance stored in the `Client`
// struct, passing in the feed type and feed slug as arguments.
func (c *Client) CreateNotificationFeed(feedType FeedType, feedSlug string) (*stream.NotificationFeed, error) {
	if feedSlug == "" || feedType == "" {
		return nil, fmt.Errorf("invalid input argument. feedSlug: %v, feedType: %v", feedSlug, feedType)
	}

	return c.client.NotificationFeed(feedType.String(), feedSlug)
}
