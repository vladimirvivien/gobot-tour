package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()
	delay := 50 * time.Millisecond

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem411")
	led := gpio.NewLedDriver(firmataAdaptor, "led", "13")

	work := func() {
		gobot.Every(delay, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
