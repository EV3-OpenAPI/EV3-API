package button

import (
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"time"
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
