// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/hiep2902/onvif"
	"github.com/hiep2902/onvif/sdk"
	"github.com/hiep2902/onvif/media"
)

// Call_GetAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourceConfigurationResponse.
func Call_GetAudioSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioSourceConfiguration) (media.GetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationResponse media.GetAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourceConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioSourceConfiguration")
		return reply.Body.GetAudioSourceConfigurationResponse, errors.Annotate(err, "reply")
	}
}
