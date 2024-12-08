// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package client // import "github.com/SolomonAIEngineering/backend-core-library/message_queue/client"

// The SendRequest type represents a message to be sent to a queue with a specified URL, body, and
// attributes.
// @property {string} QueueURL - QueueURL is a string property that represents the URL of the Amazon
// SQS queue to which the message will be sent.
// @property {string} Body - The message body that will be sent to the specified queue.
// @property {[]Attribute} Attributes - Attributes is a slice of Attribute structs that can be used to
// add custom metadata to the message being sent to the queue. These attributes can be used to provide
// additional information about the message, such as message type, priority, or any other relevant
// information. The Attribute struct typically contains a Name and a Value
type SendRequest struct {
	// `QueueURL string` is defining a string property named `QueueURL` for the `SendRequest` struct. This
	// property represents the URL of the Amazon SQS queue to which the message will be sent.
	QueueURL string
	// The `Body` property is defining a string property named `Body` for the `SendRequest` struct. This
	// property represents the message body that will be sent to the specified queue.
	Body string
	// The `Attributes []Attribute` field in the `SendRequest` struct is defining a slice of `Attribute`
	// structs that can be used to add custom metadata to the message being sent to the queue. These
	// attributes can be used to provide additional information about the message, such as message type,
	// priority, or any other relevant information. The `Attribute` struct typically contains a `Key` and a
	// `Value`, and an optional `Type` field.
	Attributes []Attribute
}

// The Attribute type represents an HTML attribute with a key, value, and type.
// @property {string} Key - A string representing the name of the attribute.
// @property {string} Value - The "Value" property is a string type that represents the value of an
// attribute. In the context of the "Attribute" struct, it is used to store the value of an HTML
// attribute.
// @property {string} Type - The "Type" property is a string that represents the type of the attribute.
// It could be "text", "number", "date", "boolean", or any other data type that is relevant to the
// attribute's purpose.
type Attribute struct {
	// The `Key` field is defining a string property named `Key` for the `Attribute` struct. This property
	// represents the name of the attribute.
	Key string
	// The `Value` field is defining a string property named `Value` for the `Attribute` struct. This
	// property represents the value of an HTML attribute. In the context of the `Attribute` struct, it is
	// used to store the value of an attribute that can be used to provide additional information about the
	// message being sent to the queue.
	Value string
	// The `Type` field in the `Attribute` struct is defining a string property named `Type` that
	// represents the data type of the attribute's value. It could be "text", "number", "date", "boolean",
	// or any other data type that is relevant to the attribute's purpose. This field is optional and can
	// be used to provide additional information about the attribute.
	Type string
}

// The above code defines a Go struct type called "Message" with four fields: ID, ReceiptHandle, Body,
// and Attributes.
// @property {string} ID - The unique identifier for the message.
// @property {string} ReceiptHandle - ReceiptHandle is a unique identifier for a message that has been
// received from an Amazon Simple Queue Service (SQS) queue. It is used to acknowledge the receipt of a
// message and to delete it from the queue.
// @property {string} Body - The message body, which contains the content of the message being sent or
// received.
// @property Attributes - Attributes is a map of key-value pairs that provide additional information
// about the message. These attributes can include metadata such as message timestamps, message
// grouping IDs, and message deduplication IDs. The specific attributes available depend on the
// messaging service being used.
type Message struct {
	// The `ID` field is defining a string property named `ID` for the `Message` struct. This property
	// represents the unique identifier for the message.
	ID string
	// The `ReceiptHandle` field in the `Message` struct is defining a string property named
	// `ReceiptHandle` that represents a unique identifier for a message that has been received from an
	// Amazon Simple Queue Service (SQS) queue. It is used to acknowledge the receipt of a message and to
	// delete it from the queue.
	ReceiptHandle string
	// The `Body` field is defining a string property named `Body` for the `Message` struct. This property
	// represents the message body, which contains the content of the message being sent or received.
	Body string
	// The `Attributes` field in the `Message` struct is defining a map of key-value pairs that provide
	// additional information about the message. These attributes can include metadata such as message
	// timestamps, message grouping IDs, and message deduplication IDs. The specific attributes available
	// depend on the messaging service being used. The `map[string]string` type indicates that the keys and
	// values in the map are both strings.
	Attributes map[string]string
}
