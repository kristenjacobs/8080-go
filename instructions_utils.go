package main

func MVI(instrName string, ms *machineState, reg *uint8) {
	*reg = ms.readMem(ms.pc+1, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x\n", ms.pc, instrName, *reg)
	ms.pc += 2
}

func MOV(instrName string, ms *machineState, dst *uint8, src *uint8) {
	*dst = *src
	Trace.Printf("0x%04x: %s 0x%02x 0x%02x\n", ms.pc, instrName, *dst, *src)
	ms.pc += 1
}
