package main

import (
	"time"
)

func generateInterrupts(ms *machineState) {
	if ms.interruptsEnabled {
		// Middle of frame interrupt (RST_1).
		ms.setInterrupt(0x08)
	}
	time.Sleep(16 * time.Millisecond)
	if ms.interruptsEnabled {
		// End of frame interrupt (RST_2).
		ms.setInterrupt(0x10)
	}
	time.Sleep(16 * time.Millisecond)
}
