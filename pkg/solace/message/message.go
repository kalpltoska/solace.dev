// pubsubplus-go-client
//
// Copyright 2021-2023 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message

import (
	"time"

	"solace.dev/go/messaging/pkg/solace/message/sdt"
)

// Message represents the common functionality between an Inbound and Outbound message.
type Message interface {
	// Message is of type Disposable. If the message has any underlying resources from the backing
	// messaging implementation, these can be freed using Disposable.Dispose(). If Disposable.Dispose()
	// is not called on the Message, finalizers attempt to clean up any underlying resources.
	// Calling Dispose() is optional, but is recommended because it improves garbage collection performance.
	// After Disposable.Dispose() is called, Message is considered unusable and all any subsequent function calls
	// return nil or empty data.
	Disposable

	// GetProperties returns a map of user properties.
	GetProperties() sdt.Map

	// GetProperty returns the user property at the given key, and a boolean indicating if its present.
	// Will return nil if not found. Property may be present and set to nil.
	GetProperty(key string) (value sdt.Data, ok bool)

	// HasProperty return whether a user property is present in Message.
	HasProperty(key string) bool

	// GetPayloadAsBytes attempts to get the payload of the message as a byte array.
	// This function return bytes containing the byte array and an ok flag indicating if it was
	// successful. If the content is not accessible in byte array form, an empty slice is
	// returned and the ok flag is false.
	GetPayloadAsBytes() (bytes []byte, ok bool)

	// GetPayloadAsString attempts to get the payload of the message as a string.
	// This function returns a string containing the data stored in the message and an ok flag
	// indicating if it was successful. If the content is not accessible in string form,
	// an empty string is returned and the ok flag is false.
	GetPayloadAsString() (str string, ok bool)

	// GetPayloadAsMap attempts to get the payload of the message as an SDTMap.
	// This function returns a SDTMap instance containing the data stored in the message and
	// an ok indicating if it was success. If the content is not accessible in SDTMap
	// form, sdtMap is nil and ok is false.
	GetPayloadAsMap() (sdtMap sdt.Map, ok bool)

	// GetPayloadAsStream attempts to get the payload of the message as an SDTStream.
	// This function returns a SDTStream instance containing the data stored in the message and
	// an ok indicating if it was success. If the content is not accessible in SDTStream
	// form, sdtStream is nil and ok is false.
	GetPayloadAsStream() (sdtStream sdt.Stream, ok bool)

	// GetCorrelationID returns the correlation ID of the message.
	// If not present, the id argument is an empty string and ok is false.
	GetCorrelationID() (id string, ok bool)

	// GetExpiration returns the expiration time of the message.
	// The expiration time is UTC time of when the message was discarded or
	// moved to the Dead Message Queue by the broker.
	// A value of zero (as determined by time.isZero()) indicates that the
	// message never expires. The default value is zero.
	GetExpiration() time.Time

	// GetSequenceNumber returns the sequence number of the message.
	// Sequence numbers may be set by the publisher applications or
	// automatically generated by the publisher APIs. The sequence number
	// is carried in the Message metadata in addition to the payload, and
	// can be retrieved by consumer applications. Returns a positive
	// sequenceNumber if set, otherwise ok is false if priority is not set.
	GetSequenceNumber() (sequenceNumber int64, ok bool)

	// GetPriority returns the priority value. Valid priorities range from
	// 0 (lowest) to 255 (highest). Returns the priority, otherwise ok is false if priority is not set.
	GetPriority() (priority int, ok bool)

	// GetHTTPContentType returns the HTTP content-type set on the message.
	// If not set, returns an empty string and ok is false.
	GetHTTPContentType() (contentType string, ok bool)

	// GetHTTPContentEncoding returns the HTTP content encoding set on the message.
	// If not set, returns an empty string and ok is false.
	GetHTTPContentEncoding() (contentEncoding string, ok bool)

	// GetApplicationMessageID returns the Application Message ID of the message.
	// This value is used by applications only and is passed through the API unmodified.
	// If not set, returns an empty string and ok is false.
	GetApplicationMessageID() (applicationMessageID string, ok bool)

	// GetApplicationMessageType returns the Application Message Type of the message.
	// This value is used by applications only and is passed through the API unmodified.
	// If not set, returns an empty string and ok is false.
	GetApplicationMessageType() (applicationMessageType string, ok bool)

	// GetClassOfService returns the class of service of the message. Class of Service is
	// represented by an integer with:
	//  0 | COS_1
	//  1 | COS_2
	//  2 | COS_3
	GetClassOfService() (cos int)

	// String implements fmt.Stringer. Prints the message as a string. A truncated response
	// may be returned when large payloads or properties are attached.
	String() string
}
