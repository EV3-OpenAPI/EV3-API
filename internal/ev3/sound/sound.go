package sound

import (
	"EV3-API/internal/ev3"
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"os/exec"
	"time"
)

var speaker *ev3dev.Speaker

const (
	soundPath  = "/dev/input/by-path/platform-sound-event"
	aplayPath  = "/usr/bin/aplay"
	espeakPath = "/usr/bin/espeak"
)

func Init() (err error) {
	log.Printf("INFO - Initializing speaker")
	speaker = ev3dev.NewSpeaker(soundPath)
	err = speaker.Init()

	return
}

func Close() (err error) {
	log.Printf("INFO - Closing speaker")
	err = speaker.Close()

	return
}

type Tone struct {
	Freq     uint32
	Duration time.Duration
}

func Tones(tones []Tone) (err error) {
	for _, tone := range tones {
		if err = speaker.Tone(tone.Freq); err != nil {
			return
		}
		time.Sleep(tone.Duration)
	}

	err = speaker.Tone(0)
	return
}

func Beep() error {
	return Tones([]Tone{{440, ev3.DurationMs(200)}})
}

func Speak(text string) error {
	cmd := fmt.Sprintf("%s --stdout -a 200 -s 130 '%s' | %s -q", espeakPath, text, aplayPath)
	if out, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		return fmt.Errorf("%s - %v", out, err)
	}

	return nil
}
