package main

import (
	"fmt"
)

func MVI(instrName string, ms *machineState, reg *uint8) {
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	*reg = byte2
	fmt.Printf("0x%04x: %s 0x%02x\n", ms.pc, instrName, *reg)
	ms.pc += 2
}
