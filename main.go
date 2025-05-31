package main

import (
	"fmt"
	"machine"
	"machine/usb"
	"machine/usb/hid/keyboard"

	_ "embed"
)

//go:embed password.txt
var password []byte
var isBtnPressed = false

func init() {
	usb.Manufacturer = "Dobefu"
	usb.Product = "Wi-Fi Typer"
	usb.Serial = fmt.Sprintf("%s/%s", usb.Manufacturer, usb.Product)
}

func main() {
	kb := keyboard.Port()

	button := machine.GPIO23
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !button.Get() {
			if !isBtnPressed {
				_, _ = kb.Write(password[:len(password)-1])
				isBtnPressed = true
			}

			continue
		}

		isBtnPressed = false
	}
}
