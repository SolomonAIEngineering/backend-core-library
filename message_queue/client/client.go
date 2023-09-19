// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package client // import "github.com/SolomonAIEngineering/backend-core-library/message_queue/client"

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// The `Client` struct defines a client for interacting with the Amazon Simple Queue Service (SQS) API,
// with fields for the SQS API interface, timeout durations, configuration settings, and AWS
// credentials.
// @property SqsClient - a field in the `Client` struct that holds an interface for the SQS API. This
// allows for the use of a mock SQS client for testing purposes.
// @property Timeout - The timeout duration for operations performed by the SQS client. This timeout
// duration specifies the maximum amount of time that the client will wait for an operation to complete
// before timing out.
// @property ReadTimeout - `ReadTimeout` is a field in the `Client` struct that holds the timeout
// duration for read operations performed by the SQS client. This timeout duration specifies the
// maximum amount of time that the client will wait for a read operation to complete before timing out.
// @property WriteTimeout - The timeout duration for write operations performed by the SQS client. This
// timeout duration specifies the maximum amount of time that the client will wait for a write
// operation to complete before timing out.
// @property Config - A pointer to a struct that holds configuration settings for the SQS client, such
// as the maximum number of messages to receive, the visibility timeout, and the wait time for long
// polling.
// @property AwsConfig - A field in the `Client` struct that holds the AWS configuration for the SQS
// client. It is used to configure the AWS credentials and region for the SQS client. The `AwsConfig`
// struct is defined elsewhere in the codebase and contains fields for these settings.
type Client struct {
	// `SqsClient sqsiface.SQSAPI` is defining a field in the `Client` struct that holds an interface for
	// the SQS API. This allows for the use of a mock SQS client for testing purposes.
	SqsClient sqsiface.SQSAPI
	// Timeout is the timeout for the client.
	// `Timeout time.Duration` is a field in the `Client` struct that holds the timeout duration for
	// operations performed by the SQS client. This timeout duration specifies the maximum amount of time
	// that the client will wait for an operation to complete before timing out.
	Timeout time.Duration
	// `ReadTimeout time.Duration` is a field in the `Client` struct that holds the timeout duration for
	// read operations performed by the SQS client. This timeout duration specifies the maximum amount of
	// time that the client will wait for a read operation to complete before timing out.
	ReadTimeout time.Duration
	// `WriteTimeout time.Duration` is a field in the `Client` struct that holds the timeout duration for
	// write operations performed by the SQS client. This timeout duration specifies the maximum amount of
	// time that the client will wait for a write operation to complete before timing out.
	WriteTimeout time.Duration
	// `Config *Config` is a field in the `Client` struct that holds the configuration for the SQS client.
	// It is used to configure various settings such as the maximum number of messages to receive,
	// the visibility timeout, and the wait time for long polling. The `Config` struct is defined
	// elsewhere in the codebase and contains fields for these settings.
	Config *Config

	// `AwsConfig    *AwsConfig` is a field in the `Client` struct that holds the AWS configuration for
	// the SQS client. It is used to configure the AWS credentials and region for the SQS client.
	AwsConfig *AwsConfig
}

// The MessageClientInterface is an interface that defines methods for sending, receiving, and deleting
// messages from an SQS queue.
// @property Send - A method that sends a message to an SQS queue. It takes a context and a
// `SendRequest` struct as input, and returns a string and an error. The string returned is the message
// ID of the sent message, and the error returned indicates whether there was an error sending the
// message.
// @property SendMessage - A method of the MessageClientInterface that is used to send a message to an
// SQS queue using the `sqs.SendMessageInput` struct as input. It takes a context and a
// `SendMessageInput` struct as input, and returns a string and an error. The string returned is the
// message ID
// @property Receive - A method of the MessageClientInterface interface that is used to receive
// messages from an SQS queue. It takes a context and a queue URL as input, and returns a slice of
// `Message` structs and an error. The `Message` struct contains information about the received
// message, such as the message
// @property {error} Delete - The `Delete` method is used to delete a message from an SQS queue. It
// takes a context, a queue URL, and a receipt handle as input, and returns an error indicating whether
// there was an error deleting the message. The receipt handle is a unique identifier for the message
// that was received
type MessageClientInterface interface {
	// The `Send` method is used to send a message to an SQS queue. It takes a context and a `SendRequest`
	// struct as input, and returns a string and an error. The string returned is the message ID of the
	// sent message, and the error returned indicates whether there was an error sending the message.
	Send(ctx context.Context, req *SendRequest) (string, error)
	// The `SendMessage` method is used to send a message to an SQS queue using the `sqs.SendMessageInput`
	// struct as input. It takes a context and a `SendMessageInput` struct as input, and returns a string
	// and an error. The string returned is the message ID of the sent message, and the error returned
	// indicates whether there was an error sending the message.
	SendMessage(ctx context.Context, msg *sqs.SendMessageInput) (*string, error)
	// The `Receive` method is used to receive messages from an SQS queue. It takes a context and a queue
	// URL as input, and returns a slice of `Message` structs and an error. The `Message` struct contains
	// information about the received message, such as the message ID, receipt handle, and message body.
	// The error returned indicates whether there was an error receiving the messages.
	Receive(ctx context.Context, queueURL string) ([]*Message, error)
	// The `Delete` method is used to delete a message from an SQS queue. It takes a context, a queue URL,
	// and a receipt handle as input, and returns an error indicating whether there was an error deleting
	// the message. The receipt handle is a unique identifier for the message that was received from the
	// queue, and is used to identify the message to be deleted.
	Delete(ctx context.Context, queueURL, rcvHandle string) error
}

// checks if the `Client` struct
// implements the `MessageClientInterface` interface. It creates a variable of type
// `MessageClientInterface` and assigns it to a pointer to a `Client` struct that is `nil`. This
// statement will cause a compilation error if the `Client` struct does not implement all the methods
// defined in the `MessageClientInterface` interface. It is a way to ensure that the `Client` struct
// satisfies the interface at compile time.
var _ MessageClientInterface = (*Client)(nil)

// NewClient creates a new client with optional configuration options.
func NewClient(opts ...Option) (*Client, error) {
	if len(opts) == 0 {
		return nil, fmt.Errorf("no options provided")
	}

	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}
