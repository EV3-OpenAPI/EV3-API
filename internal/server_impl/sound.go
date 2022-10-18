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

const SoundPath = "/dev/input/by-path/platform-sound-event"

func (s *ApiService) SoundTonePost(ctx context.Context, tone openapi.Tone) (openapi.ImplResponse, error) {

	var speaker = ev3dev.NewSpeaker(SoundPath)

	speaker.Init()
	defer speaker.Close()

	d, _ := time.ParseDuration(fmt.Sprintf("%vms", tone.LengthMs))
	speaker.Tone(uint32(tone.Frequency))
	time.Sleep(d)

	speaker.Tone(0)

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *ApiService) SoundTonesPost(ctx context.Context, tones []openapi.Tone) (openapi.ImplResponse, error) {

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
