package server_impl

import (
	"EV3-API/internal/openapi"
	"context"
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"net/http"
	"time"
)

// SoundApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type SoundApiService struct {
}

// NewSoundApiService creates a default api service
func NewSoundApiService() openapi.SoundApiServicer {
	return &SoundApiService{}
}

const SoundPath = "/dev/input/by-path/platform-sound-event"

func (s *SoundApiService) SoundTonePost(_ context.Context, tone openapi.Tone) (openapi.ImplResponse, error) {

	var speaker = ev3dev.NewSpeaker(SoundPath)

	if err := speaker.Init(); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	defer speaker.Close()

	d, _ := time.ParseDuration(fmt.Sprintf("%vms", tone.LengthMs)) // FIXME replace by 1000*1000*tone.LengthMs or so
	if err := speaker.Tone(uint32(tone.Frequency)); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	time.Sleep(d)

	if err := speaker.Tone(0); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *SoundApiService) SoundTonesPost(_ context.Context, tones []openapi.Tone) (openapi.ImplResponse, error) {

	var speaker = ev3dev.NewSpeaker(SoundPath)

	if err := speaker.Init(); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	defer speaker.Close()

	for i, tone := range tones {
		log.Printf("INFO - Playing tone %d/%d", i+1, len(tones))

		d, _ := time.ParseDuration(fmt.Sprintf("%vms", tone.LengthMs))
		if err := speaker.Tone(uint32(tone.Frequency)); err != nil {
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
		time.Sleep(d)
	}

	if err := speaker.Tone(0); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}
