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

// NewPowerApiService creates a default api service
func NewSoundApiService() openapi.SoundApiServicer {
	return &SoundApiService{}
}

const SoundPath = "/dev/input/by-path/platform-sound-event"

func (s *SoundApiService) SoundTonePost(ctx context.Context, tone openapi.Tone) (openapi.ImplResponse, error) {

	var speaker = ev3dev.NewSpeaker(SoundPath)

	speaker.Init()
	defer speaker.Close()

	d, _ := time.ParseDuration(fmt.Sprintf("%vms", tone.LengthMs))
	speaker.Tone(uint32(tone.Frequency))
	time.Sleep(d)

	speaker.Tone(0)

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *SoundApiService) SoundTonesPost(ctx context.Context, tones []openapi.Tone) (openapi.ImplResponse, error) {

	var speaker = ev3dev.NewSpeaker(SoundPath)

	speaker.Init()
	defer speaker.Close()

	for i, tone := range tones {
		log.Printf("INFO - Playing tone %d/%d", i+1, len(tones))

		d, _ := time.ParseDuration(fmt.Sprintf("%vms", tone.LengthMs))
		speaker.Tone(uint32(tone.Frequency))
		time.Sleep(d)
	}

	speaker.Tone(0)

	return openapi.Response(http.StatusOK, nil), nil
}
