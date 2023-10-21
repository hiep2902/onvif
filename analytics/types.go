package analytics

import (
	"github.com/hiep2902/onvif/xsd"
	"github.com/hiep2902/onvif/xsd/onvif"
)

type GetSupportedRules struct {
	XMLName            string               `xml:"GetSupportedRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
}

type CreateRules struct {
	XMLName            string               `xml:"CreateRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
	Rule               onvif.Config         `xml:"Rule"`
}

type DeleteRules struct {
	XMLName            string               `xml:"DeleteRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
	RuleName           xsd.String           `xml:"RuleName"`
}

type GetRules struct {
	XMLName            string               `xml:"GetRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
}

type GetRuleOptions struct {
	XMLName            string               `xml:"GetRuleOptions"`
	RuleType           xsd.QName            `xml:"RuleType"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
}

type ModifyRules struct {
	XMLName            string               `xml:"ModifyRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
	Rule               onvif.Config         `xml:"Rule"`
}

type GetServiceCapabilities struct {
	XMLName string `xml:"GetServiceCapabilities"`
}

type GetSupportedAnalyticsModules struct {
	XMLName            string               `xml:"GetSupportedAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
}

type GetAnalyticsModuleOptions struct {
	XMLName            string               `xml:"GetAnalyticsModuleOptions"`
	Type               xsd.QName            `xml:"Type"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
}

type CreateAnalyticsModules struct {
	XMLName            string               `xml:"CreateAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
	AnalyticsModule    onvif.Config         `xml:"AnalyticsModule"`
}

type DeleteAnalyticsModules struct {
	XMLName             string               `xml:"DeleteAnalyticsModules"`
	ConfigurationToken  onvif.ReferenceToken `xml:"ConfigurationToken"`
	AnalyticsModuleName xsd.String           `xml:"AnalyticsModuleName"`
}

type GetAnalyticsModules struct {
	XMLName            string               `xml:"GetAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
}

type ModifyAnalyticsModules struct {
	XMLName            string               `xml:"ModifyAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken `xml:"ConfigurationToken"`
	AnalyticsModule    onvif.Config         `xml:"AnalyticsModule"`
}
