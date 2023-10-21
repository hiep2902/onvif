package imaging

import (
	"github.com/hiep2902/onvif/xsd"
	"github.com/hiep2902/onvif/xsd/onvif"
)

type GetServiceCapabilities struct {
	XMLName string `xml:"GetServiceCapabilities"`
}

type GetImagingSettings struct {
	XMLName          string               `xml:"GetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
}

type SetImagingSettings struct {
	XMLName          string                  `xml:"SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken    `xml:"VideoSourceToken"`
	ImagingSettings  onvif.ImagingSettings20 `xml:"ImagingSettings"`
	ForcePersistence xsd.Boolean             `xml:"ForcePersistence"`
}

type GetOptions struct {
	XMLName          string               `xml:"GetOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
}

type Move struct {
	XMLName          string               `xml:"Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
	Focus            onvif.FocusMove      `xml:"Focus"`
}

type GetMoveOptions struct {
	XMLName          string               `xml:"GetMoveOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
}

type Stop struct {
	XMLName          string               `xml:"Stop"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
}

type GetStatus struct {
	XMLName          string               `xml:"GetStatus"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
}

type GetPresets struct {
	XMLName          string               `xml:"GetPresets"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string               `xml:"GetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string               `xml:"SetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"VideoSourceToken"`
	PresetToken      onvif.ReferenceToken `xml:"PresetToken"`
}
