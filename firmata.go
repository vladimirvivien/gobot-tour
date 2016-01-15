package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hybridgroup/gobot/platforms/firmata"
)

func main() {

	board := firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem411")
	errs := board.Connect()
	if len(errs) != 0 {
		fmt.Println(errs)
		os.Exit(1)
	}
	pin := "9"
	level := byte(0)
	step := byte(5)
	delay := 30 * time.Millisecond
	for {
		board.DigitalWrite(pin, level)
		level += step
		if level == 0 || level == 255 {
			step = -step
		}
		time.Sleep(delay)
	}
}
