package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()
	pin := "9"
	brightness := byte(0)
	fade := byte(5)
	delay := 30 * time.Millisecond

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem411")
	led := gpio.NewLedDriver(firmataAdaptor, "led", pin)

	work := func() {
		for {
			led.Brightness(brightness)
			brightness += fade
			if brightness == 0 || brightness == 255 {
				fade = -fade
			}
			time.Sleep(delay)
		}
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
