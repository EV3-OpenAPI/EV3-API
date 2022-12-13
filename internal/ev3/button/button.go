package button

import (
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"time"
)

var (
	lastButtonEvent *ev3dev.ButtonEvent
)

func Poll() {
	var b ev3dev.ButtonPoller

	for i := 0; i < 30; i++ {
		b, err := b.Poll()
		if err != nil {
			log.Fatalf("failed to poll keys: %v", err)
		}
		fmt.Printf("%6b\n", b)
		time.Sleep(5 * time.Second)
	}
}

func Wait() {
	go wait()
}

func wait() {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}

	for e := range w.Events {
		lastButtonEvent = &e
		log.Printf("DEBUG - %+v\n", e)
	}
}

// GetLastButtonEvent gets last ev3dev.ButtonEvent
func GetLastButtonEvent(clear bool) (btnEvt *ev3dev.ButtonEvent) {
	if lastButtonEvent == nil {
		return nil
	}

	*btnEvt = *lastButtonEvent
	if clear {
		lastButtonEvent = nil
	}

	return btnEvt
}
