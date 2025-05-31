.PHONY: build

default: build

build:
	@echo "Flashing the binary..."
	@tinygo flash -target pico
