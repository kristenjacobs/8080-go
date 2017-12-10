package main

func MVI(instrName string, ms *machineState, reg *uint8) {
	*reg = ms.readMem(ms.pc+1, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x\n", ms.pc, instrName, *reg)
	ms.pc += 2
}

func MOV_REG_REG(instrName string, ms *machineState, dst *uint8, src *uint8) {
	*dst = *src
	Trace.Printf("0x%04x: %s 0x%02x 0x%02x\n", ms.pc, instrName, *dst, *src)
	ms.pc += 1
}

func MOV_REG_MEM(instrName string, ms *machineState, dst *uint8) {
	addr := ms.addr(ms.regL, ms.regH)
	*dst = ms.readMem(addr, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x [0x%04x]\n", ms.pc, instrName, *dst, addr)
	ms.pc += 1
}

func LDAX(instrName string, ms *machineState, adrRegLo *uint8, adrRegHi *uint8) {
	var adr uint16 = (uint16(*adrRegHi) << 8) | uint16(*adrRegLo)
	ms.regA = ms.readMem(adr, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x 0x%04x\n", ms.pc, instrName, ms.regA, adr)
	ms.pc += 1
}
