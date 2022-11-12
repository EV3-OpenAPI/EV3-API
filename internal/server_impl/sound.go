package server_impl

import (
	"EV3-API/internal/ev3"
	"EV3-API/internal/ev3/sound"
	"EV3-API/internal/gen/openapi"
	"context"
	"net/http"
)

// SoundApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type SoundApiService struct {
}

func (s *SoundApiService) SoundSpeakPost(_ context.Context, text openapi.Text) (openapi.ImplResponse, error) {
	if err := sound.Speak(text.Text); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

// NewSoundApiService creates a default api service
func NewSoundApiService() openapi.SoundApiServicer {
	return &SoundApiService{}
}

func (s *SoundApiService) SoundBeepPost(_ context.Context) (openapi.ImplResponse, error) {
	if err := sound.Beep(); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *SoundApiService) SoundTonePost(_ context.Context, tone openapi.Tone) (openapi.ImplResponse, error) {
	tones := []sound.Tone{{uint32(tone.Frequency), ev3.DurationMs(tone.LengthMs)}}

	if err := sound.Tones(tones); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *SoundApiService) SoundTonesPost(_ context.Context, tonesReq []openapi.Tone) (openapi.ImplResponse, error) {
	tones := make([]sound.Tone, len(tonesReq))

	// Parse requested tones
	for i, tone := range tonesReq {
		tones[i] = sound.Tone{Freq: uint32(tone.Frequency), Duration: ev3.DurationMs(tone.LengthMs)}
	}

	if err := sound.Tones(tones); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}
