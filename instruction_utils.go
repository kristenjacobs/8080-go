package main

import (
	"fmt"
	"math"
)

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

func CALL(instrName string, ms *machineState, condFlagName string, condFlagVal bool) {
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	PC := ms.pc
	nextPC := ms.pc + 3
	var syscall string
	if !condFlagVal {
		ms.pc = nextPC
		syscall = ""

	} else if isSyscallAddress(adr) {
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
		ms.pc = nextPC
		syscall = "[SYSCALL]"

	} else {
		pcHi := uint8(nextPC >> 8)
		pcLo := uint8(nextPC & 0xFF)
		ms.writeMem(ms.sp-2, []uint8{pcLo, pcHi}, 2)
		ms.sp = ms.sp - 2
		ms.pc = adr
		syscall = ""
	}
	if condFlagName != "" {
		Trace.Printf("0x%04x: %s 0x%04x, Taken=%t %s\n", PC, instrName, adr, condFlagVal, syscall)
	} else {
		Trace.Printf("0x%04x: %s 0x%04x %s\n", PC, instrName, adr, syscall)
	}
}

func RET(instrName string, ms *machineState, condFlagName string, condFlagVal bool) {
	currentPc := ms.pc
	bytes := ms.readMem(ms.sp, 2)
	pcLo := bytes[0]
	pcHi := bytes[1]
	newSp := ms.sp + 2
	pc := (uint16(pcHi) << 8) | uint16(pcLo)
	if condFlagVal {
		ms.pc = (uint16(pcHi) << 8) | uint16(pcLo)
		ms.sp = newSp
	} else {
		ms.pc = currentPc + 1
	}
	if condFlagName != "" {
		Trace.Printf("0x%04x: %s 0x%04x, Taken=%t\n", currentPc, instrName, pc, condFlagVal)
	} else {
		Trace.Printf("0x%04x: %s 0x%04x\n", currentPc, instrName, pc)
	}
}

func INR(instrName string, ms *machineState, reg *uint8) {
	r := *reg
	*reg = *reg + 1
	ms.setZ(*reg)
	ms.setS(*reg)
	ms.setP(*reg)
	ms.setCY(uint(*reg)+uint(*reg) > math.MaxUint8)
	ms.setAC(*reg)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]+1, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, *reg, r, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ADD(instrName string, ms *machineState, srcReg string, reg *uint8) {
	regA := ms.regA
	r := *reg
	ms.regA = regA + r
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(r) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]+reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, r, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func SUB(instrName string, ms *machineState, srcReg string, reg *uint8) {
	regA := ms.regA
	r := *reg
	ms.regA = regA - r
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(r) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]+reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, r, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ADC(instrName string, ms *machineState, srcReg string, reg *uint8) {
	regA := ms.regA
	r := *reg
	ms.regA = regA + r
	var carry uint8 = 0
	if ms.flagCY {
		carry = 1
	}
	ms.regA += carry
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(r)+uint(carry) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]+reg%s[0x%02x]+%d, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, r, carry, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func SBB(instrName string, ms *machineState, srcReg string, reg *uint8) {
	regA := ms.regA
	r := *reg
	ms.regA = regA - r
	var carry uint8 = 0
	if ms.flagCY {
		carry = 1
	}
	ms.regA -= carry
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)-uint(r)-uint(carry) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]-reg%s[0x%02x]-%d, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, r, carry, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ANA(instrName string, ms *machineState, srcReg string, reg *uint8) {
	regA := ms.regA
	r := *reg
	ms.regA = regA & r
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)&uint(r) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]&reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, r, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ORA(instrName string, ms *machineState, srcReg string, reg *uint8) {
	regA := ms.regA
	r := *reg
	ms.regA = regA | r
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)|uint(r) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]|reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, r, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func isSyscallAddress(adr uint16) bool {
	return adr == 0x5
}
