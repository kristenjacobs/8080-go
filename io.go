package main

import (
	"fmt"
	"time"
)

type IO interface {
	Read(port uint8)
	Write(port uint8)
}

type IOHandler struct {
}

func newIOHandler() *IOHandler {
	return &IOHandler{}
}

func (io *IOHandler) Read(port uint8) {
	fmt.Printf("IOHandler Read: %d\n", port)
}

func (io *IOHandler) Write(port uint8) {
	fmt.Printf("IOHandler Write: %d\n", port)
}

func (io *IOHandler) run(ms *machineState) {
	for ms.halt == false {
		fmt.Printf("running ... \n")
		time.Sleep(1 * time.Second)
	}
}
