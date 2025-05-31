package main

import (
	"fmt"
	"machine/usb"
)

func init() {
	usb.Manufacturer = "Dobefu"
	usb.Product = "Wi-Fi Typer"
	usb.Serial = fmt.Sprintf("%s/%s", usb.Manufacturer, usb.Product)
}

func main() {
}
