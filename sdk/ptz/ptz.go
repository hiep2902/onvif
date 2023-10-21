package ptz

//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetServiceCapabilities
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetNodes
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetNode
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetConfiguration
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetConfigurations
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz SetConfiguration
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetConfigurationOptions
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz SendAuxiliaryCommand
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetPresets
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz SetPreset
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz RemovePreset
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GotoPreset
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GotoHomePosition
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz SetHomePosition
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz ContinuousMove
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz RelativeMove
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetStatus
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz AbsoluteMove
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GeoMove
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz Stop
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetPresetTours
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetPresetTour
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetPresetTourOptions
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz CreatePresetTour
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz ModifyPresetTour
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz OperatePresetTour
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz RemovePresetTour
//go:generate go run github.com/hiep2902/onvif/sdk/codegen ptz ptz GetCompatibleConfigurations
