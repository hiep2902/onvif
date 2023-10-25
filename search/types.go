package search

import (
	"github.com/hiep2902/onvif/xsd"
	"github.com/hiep2902/onvif/xsd/onvif"
)

type FindRecordings struct {
	XMLName       string `xml:"tse:FindRecordings"`
	Scope         onvif.SearchScope
	MaxMatches    xsd.Int
	KeepAliveTime xsd.Duration
}

type FindRecordingsResponse struct {
	SearchToken onvif.JobToken
}

type GetRecordingSearchResults struct {
	XMLName     string `xml:"tse:GetRecordingSearchResults"`
	SearchToken onvif.JobToken
	MinResults  xsd.Int
	MaxResults  xsd.Int
	WaitTime    xsd.Duration
}

type GetRecordingSearchResultsResponse struct {
	ResultList onvif.FindRecordingResultList
}

type FindEvents struct {
	XMLName           string `xml:"tse:FindEvents"`
	StartPoint        xsd.DateTime
	EndPoint          xsd.DateTime
	Scope             onvif.SearchScope
	SearchFilter      onvif.EventFilter
	IncludeStartState xsd.Boolean
	MaxMatches        xsd.Int
	KeepAliveTime     xsd.Duration
}

type FindEventsResponse struct {
	SearchToken onvif.JobToken
}

type GetEventSearchResults struct {
	XMLName     string `xml:"tse:GetEventSearchResults"`
	SearchToken onvif.JobToken
	MinResults  xsd.Int
	MaxResults  xsd.Int
	WaitTime    xsd.Duration
}

type GetEventSearchResultsResponse struct {
	ResultList onvif.FindEventResultList
}

type EndSearch struct {
	XMLName     string `xml:"tse:EndSearch"`
	SearchToken onvif.JobToken
}

type EndSearchResponse struct {
	Endpoint xsd.DateTime
}
