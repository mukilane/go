package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	ir := gpio.NewDirectPinDriver(firmataAdaptor, "A0")

	work := func() {
		gobot.Every(1*time.Second, func() {
			fmt.Println(ir.DigitalRead())
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{ir},
		work,
	)

	robot.Start()
}
