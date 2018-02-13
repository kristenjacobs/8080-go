package main

import (
	"fmt"
	"time"
)

type IOHandler struct {
}

func newIOHandler() *IOHandler {
	return &IOHandler{}
}

func (io *IOHandler) Read(port uint8) uint8 {
	var value uint8 = 0
	fmt.Printf("IOHandler Read: %d, value: 0x%02x\n", port, value)
	return value
}

func (io *IOHandler) Write(port uint8, value uint8) {
	fmt.Printf("IOHandler Write: port: %d, value: 0x%02x\n", port, value)
}

func (io *IOHandler) run(ms *machineState) {
	for ms.halt == false {
		fmt.Printf("running ... \n")
		time.Sleep(1 * time.Second)
	}
}
