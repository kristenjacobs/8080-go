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

func MOV_MEM_REG(instrName string, ms *machineState, srcReg *uint8) {
	addr := ms.addr(ms.regL, ms.regH)
	ms.writeMem(addr, []uint8{*srcReg}, 1)
	Trace.Printf("0x%04x: %s [0x%04x] 0x%02x\n", ms.pc, instrName, addr, *srcReg)
	ms.pc += 1
}

func LDAX(instrName string, ms *machineState, adrRegHi *uint8, adrRegLo *uint8) {
	var adr uint16 = (uint16(*adrRegHi) << 8) | uint16(*adrRegLo)
	ms.regA = ms.readMem(adr, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x 0x%04x\n", ms.pc, instrName, ms.regA, adr)
	ms.pc += 1
}

func INX(instrName string, ms *machineState, regHi *uint8, regLo *uint8) {
	result := ms.addr(*regLo, *regHi) + 1
	*regHi = uint8(result >> 8)
	*regLo = uint8(result & 0xFF)
	Trace.Printf("0x%04x: %s 0x%04x\n", ms.pc, instrName, result)
	ms.pc += 1
}

func PUSH(instrName string, ms *machineState, regHi *uint8, regLo *uint8) {
	ms.writeMem(ms.sp-2, []uint8{*regHi}, 1)
	ms.writeMem(ms.sp-1, []uint8{*regLo}, 1)
	newSp := ms.sp - 2
	Trace.Printf("0x%02x: %s pc=0x%04x, (0x%04x) <- 0x%02x, (0x%04x) <- 0x%02x, sp <- 0x%04x\n",
		ms.pc, instrName, ms.pc, ms.sp-2, *regHi, ms.sp-1, *regLo, newSp)
	ms.pc += 3
	ms.sp = newSp
}
