package recording

import "github.com/hiep2902/onvif/xsd/onvif"

type GetRecordings struct {
	XMLName string `xml:"trc:GetRecordings"`
}

type GetRecordingsResponse struct {
	RecordingItem []onvif.GetRecordingsResponseItem
}
