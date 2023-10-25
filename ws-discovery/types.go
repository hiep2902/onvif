package wsdiscovery

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/hiep2902/onvif/xsd"
)

type Types []xsd.QName

type Scopes []xsd.AnyURI

type XAddress xsd.AnyURI

type XAddresses []XAddress

type EndpointReference struct {
	Address xsd.String
}

type ProbeMatchItem struct {
	EndpointReference EndpointReference
	Types             Types
	Scopes            Scopes
	XAddrs            XAddresses
	MetadataVersion   xsd.UnsignedInt
}

type ProbeMatches struct {
	ProbeMatch []ProbeMatchItem
}

type ScopeParams struct {
	params map[string][]string
}

func (o *XAddresses) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	vals := strings.Split(v, " ")
	for _, val := range vals {
		val = strings.TrimSpace(val)
		if len(val) > 0 {
			*o = append([]XAddress(*o), XAddress(val))
		}
	}

	return nil
}

func (o *Types) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	vals := strings.Split(v, " ")
	for _, val := range vals {
		val = strings.TrimSpace(val)
		if len(val) > 0 {
			*o = append([]xsd.QName(*o), xsd.QName(val))
		}
	}
	return nil
}

func (o *Scopes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	vals := strings.Split(v, " ")
	for _, val := range vals {
		val = strings.TrimSpace(val)
		if len(val) > 0 {
			*o = append([]xsd.AnyURI(*o), xsd.AnyURI(val))
		}
	}
	return nil
}

func (o Scopes) GetScopeParams() *ScopeParams {
	scopePattern, _ := regexp.Compile(`onvif://www.onvif.org\/(\w+)\/(.+)`)

	params := make(map[string][]string)
	for _, scopeURI := range o {
		submatches := scopePattern.FindStringSubmatch(string(scopeURI))
		if len(submatches) < 3 {
			continue
		}

		param := submatches[1]
		value, _ := url.QueryUnescape(submatches[2])

		var vals []string
		var ok bool
		if vals, ok = params[param]; !ok {
			vals = make([]string, 0)
		}
		vals = append(vals, value)
		params[param] = vals
	}
	return &ScopeParams{
		params: params,
	}
}

func (o ScopeParams) GetProfile() []string {
	return o.GetParam("Profile")
}

func (o ScopeParams) GetLocation() []string {
	return o.GetParam("location")
}

func (o ScopeParams) GetHardware() []string {
	return o.GetParam("hardware")
}

func (o ScopeParams) GetName() []string {
	return o.GetParam("name")
}

func (o ScopeParams) GetParam(param string) []string {
	if vals, ok := o.params[param]; ok {
		return vals
	}
	return nil
}
