package replay

import (
	"github.com/hiep2902/onvif/xsd"
	"github.com/hiep2902/onvif/xsd/onvif"
)

type GetReplayUri struct {
	XMLName        string `xml:"trp:GetReplayUri"`
	StreamSetup    onvif.StreamSetup
	RecordingToken onvif.ReferenceToken
}

type GetReplayUriResponse struct {
	XMLName string `xml:"trp:GetReplayUriResponse"`
	Uri     xsd.AnyURI
}
