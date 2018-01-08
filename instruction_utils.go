package main

import "fmt"

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
	addr := getPair(ms.regH, ms.regL)
	*dst = ms.readMem(addr, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x [0x%04x]\n", ms.pc, instrName, *dst, addr)
	ms.pc += 1
}

func MOV_MEM_REG(instrName string, ms *machineState, srcReg *uint8) {
	addr := getPair(ms.regH, ms.regL)
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
	result := getPair(*regHi, *regLo) + 1
	*regHi = uint8(result >> 8)
	*regLo = uint8(result & 0xFF)
	Trace.Printf("0x%04x: %s 0x%04x\n", ms.pc, instrName, result)
	ms.pc += 1
}

func PUSH(instrName string, ms *machineState, regHi *uint8, regLo *uint8) {
	ms.writeMem(ms.sp-2, []uint8{*regHi}, 1)
	ms.writeMem(ms.sp-1, []uint8{*regLo}, 1)
	newSp := ms.sp - 2
	Trace.Printf("0x%04x: %s (0x%04x) <- 0x%02x, (0x%04x) <- 0x%02x, sp <- 0x%04x\n",
		ms.pc, instrName, ms.sp-2, *regHi, ms.sp-1, *regLo, newSp)
	ms.pc += 1
	ms.sp = newSp
}

func POP(instrName string, ms *machineState, regHi *uint8, regLo *uint8) {
	*regHi = ms.readMem(ms.sp, 1)[0]
	*regLo = ms.readMem(ms.sp+1, 1)[0]
	newSp := ms.sp + 2
	Trace.Printf("0x%04x: %s 0x%02x <- (0x%04x), 0x%02x <- (0x%04x), sp <- 0x%04x\n",
		ms.pc, instrName, *regHi, ms.sp, *regLo, ms.sp+1, newSp)
	ms.pc += 1
	ms.sp = newSp
}

func DAD(instrName string, ms *machineState, regHi *uint8, regLo *uint8) {
	lhs := getPair(ms.regH, ms.regL)
	rhs := getPair(*regHi, *regLo)
	if (uint32(lhs) + uint32(rhs)) > 0xffff {
		ms.setCY(true)
	} else {
		ms.setCY(false)
	}
	result := lhs + rhs
	setPair(&ms.regH, &ms.regL, result)
	Trace.Printf("0x%04x: %s 0x%04x = 0x%04x + 0x%04x, CY=%t\n", ms.pc, instrName, result, lhs, rhs, ms.flagCY)
	ms.pc += 1
}

func RST(instrName string, ms *machineState, addr uint16) {
	nextPC := ms.pc + 1
	pcHi := uint8(nextPC >> 8)
	pcLo := uint8(nextPC & 0xFF)
	ms.writeMem(ms.sp-2, []uint8{pcLo, pcHi}, 2)
	Trace.Printf("0x%04x: %s 0x%04x\n", ms.pc, instrName, addr)
	ms.sp = ms.sp - 2
	ms.pc = addr
}

func handleSyscall(ms *machineState, adr uint16) bool {
	if adr == 0x5 {
		Trace.Printf("0x%04x: 0xcd_CALL_adr 0x%04x [SYSCALL]\n", ms.pc, adr)
		offset := getPair(ms.regD, ms.regE)
		for i := 0; ; i++ {
			char := ms.readMem(offset+3+uint16(i), 1)[0]
			if char == '$' {
				break
			}
			fmt.Printf("%c", char)
			if char == '$' {
				fmt.Print("\n")
				break
			}
		}
		ms.pc += 3
		return true
	}
	return false
}
