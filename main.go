package main

import (
	"fmt"
	"machine"
	"machine/usb"
	"machine/usb/hid/keyboard"
	"math"

	_ "embed"
)

//go:embed password.txt
var password []byte
var btnPressDuration float64 = 0.0

func init() {
	usb.Manufacturer = "Dobefu"
	usb.Product = "Wi-Fi Typer"
	usb.Serial = fmt.Sprintf("%s/%s", usb.Manufacturer, usb.Product)
}

func main() {
	kb := keyboard.Port()

	button := machine.GPIO7
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if button.Get() {
			if btnPressDuration <= 0 {
				_, _ = kb.Write(password[:len(password)-1])
			}

			btnPressDuration = 10000
			continue
		}

		btnPressDuration = math.Max(btnPressDuration-1, 0)
	}
}
