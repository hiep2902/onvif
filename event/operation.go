package event

import (
	"github.com/hiep2902/onvif/xsd"
)

// GetServiceCapabilities action
type GetServiceCapabilities struct {
	XMLName string `xml:"GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

// SubscriptionPolicy action
type SubscriptionPolicy struct { //tev http://www.onvif.org/ver10/events/wsdl
	ChangedOnly xsd.Boolean `xml:"ChangedOnly,attr"`
}

// Subscribe action for subscribe event topic
type Subscribe struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	XMLName                struct{}                   `xml:"Subscribe"`
	ConsumerReference      EndpointReferenceType      `xml:"ConsumerReference"`
	Filter                 FilterType                 `xml:"Filter"`
	SubscriptionPolicy     SubscriptionPolicy         `xml:"SubscriptionPolicy"`
	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime"`
}

// SubscribeResponse message for subscribe event topic
type SubscribeResponse struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	SubscriptionReference EndpointReferenceType
	CurrentTime           CurrentTime
	TerminationTime       TerminationTime
}

// Renew action for refresh event topic subscription
type Renew struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	TerminationTime AbsoluteOrRelativeTimeType `xml:"TerminationTime"`
}

// RenewResponse for Renew action
type RenewResponse struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	TerminationTime TerminationTime `xml:"TerminationTime"`
	CurrentTime     CurrentTime     `xml:"CurrentTime"`
}

// Unsubscribe action for Unsubscribe event topic
type Unsubscribe struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	Any string
}

// UnsubscribeResponse message for Unsubscribe event topic
type UnsubscribeResponse struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	Any string
}

// CreatePullPointSubscription action
type CreatePullPointSubscription struct {
	XMLName                string                     `xml:"CreatePullPointSubscription"`
	Filter                 FilterType                 `xml:"Filter"`
	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime"`
	SubscriptionPolicy     SubscriptionPolicy         `xml:"sSubscriptionPolicy"`
}

// CreatePullPointSubscriptionResponse action
type CreatePullPointSubscriptionResponse struct {
	SubscriptionReference EndpointReferenceType
	CurrentTime           CurrentTime
	TerminationTime       TerminationTime
}

// GetEventProperties action
type GetEventProperties struct {
	XMLName string `xml:"GetEventProperties"`
}

// GetEventPropertiesResponse action
type GetEventPropertiesResponse struct {
	TopicNamespaceLocation          xsd.AnyURI
	FixedTopicSet                   FixedTopicSet
	TopicSet                        TopicSet
	TopicExpressionDialect          TopicExpressionDialect
	MessageContentFilterDialect     xsd.AnyURI
	ProducerPropertiesFilterDialect xsd.AnyURI
	MessageContentSchemaLocation    xsd.AnyURI
}

//Port type PullPointSubscription

// PullMessages Action
type PullMessages struct {
	XMLName      string       `xml:"PullMessages"`
	Timeout      xsd.Duration `xml:"Timeout"`
	MessageLimit xsd.Int      `xml:"MessageLimit"`
}

// PullMessagesResponse response type
type PullMessagesResponse struct {
	CurrentTime         CurrentTime
	TerminationTime     TerminationTime
	NotificationMessage NotificationMessage
}

// PullMessagesFaultResponse response type
type PullMessagesFaultResponse struct {
	MaxTimeout      xsd.Duration
	MaxMessageLimit xsd.Int
}

// Seek action
type Seek struct {
	XMLName string       `xml:"Seek"`
	UtcTime xsd.DateTime `xml:"UtcTime"`
	Reverse xsd.Boolean  `xml:"Reverse"`
}

// SeekResponse action
type SeekResponse struct {
}

// SetSynchronizationPoint action
type SetSynchronizationPoint struct {
	XMLName string `xml:"SetSynchronizationPoint"`
}

// SetSynchronizationPointResponse action
type SetSynchronizationPointResponse struct {
}
