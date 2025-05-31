package main

import (
	"fmt"
	"machine/usb"
	"machine/usb/hid/keyboard"
	"time"

	_ "embed"
)

//go:embed password.txt
var password []byte

func init() {
	usb.Manufacturer = "Dobefu"
	usb.Product = "Wi-Fi Typer"
	usb.Serial = fmt.Sprintf("%s/%s", usb.Manufacturer, usb.Product)
}

func main() {
	kb := keyboard.Port()
	time.Sleep(time.Second * 2)

	_, _ = kb.Write(password[:len(password)-1])
}
