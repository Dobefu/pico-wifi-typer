package main

import (
	"fmt"
	"machine/usb"
	"machine/usb/hid/keyboard"
	"time"
)

func init() {
	usb.Manufacturer = "Dobefu"
	usb.Product = "Wi-Fi Typer"
	usb.Serial = fmt.Sprintf("%s/%s", usb.Manufacturer, usb.Product)
}

func main() {
	kb := keyboard.Port()
	time.Sleep(time.Second * 2)

	_, _ = kb.Write([]byte("Test"))
}
