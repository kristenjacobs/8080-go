package core

import (
	"fmt"
	"math"
	"strings"
)

func MVI(instrName string, regName string, ms *MachineState, reg *uint8) {
	*reg = ms.ReadMem(ms.pc+1, 1)[0]
	Trace.Printf("0x%04x: %s reg%s[0x%02x]=0x%02x\n", ms.pc, instrName, regName, *reg, *reg)
	ms.pc += 2
}

func MOV_REG_REG(instrName string, ms *MachineState, dst *uint8, src *uint8) {
	*dst = *src
	Trace.Printf("0x%04x: %s 0x%02x 0x%02x\n", ms.pc, instrName, *dst, *src)
	ms.pc += 1
}

func MOV_REG_MEM(instrName string, ms *MachineState, dst *uint8) {
	addr := getPair(ms.regH, ms.regL)
	*dst = ms.ReadMem(addr, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x [0x%04x]\n", ms.pc, instrName, *dst, addr)
	ms.pc += 1
}

func MOV_MEM_REG(instrName string, ms *MachineState, srcReg *uint8) {
	addr := getPair(ms.regH, ms.regL)
	ms.WriteMem(addr, []uint8{*srcReg}, 1)
	Trace.Printf("0x%04x: %s [0x%04x] 0x%02x\n", ms.pc, instrName, addr, *srcReg)
	ms.pc += 1
}

func LDAX(instrName string, ms *MachineState, adrRegHi *uint8, adrRegLo *uint8) {
	var adr uint16 = (uint16(*adrRegHi) << 8) | uint16(*adrRegLo)
	ms.regA = ms.ReadMem(adr, 1)[0]
	Trace.Printf("0x%04x: %s 0x%02x 0x%04x\n", ms.pc, instrName, ms.regA, adr)
	ms.pc += 1
}

func STAX(instrName string, ms *MachineState, adrRegHi *uint8, adrRegLo *uint8) {
	var adr uint16 = (uint16(*adrRegHi) << 8) | uint16(*adrRegLo)
	ms.WriteMem(adr, []uint8{ms.regA}, 1)
	Trace.Printf("0x%04x: %s (0x%04x)=regA[0x%02x]\n", ms.pc, instrName, adr, ms.regA)
	ms.pc += 1
}

func INX(instrName string, ms *MachineState, regHi *uint8, regLo *uint8) {
	result := getPair(*regHi, *regLo) + 1
	*regHi = uint8(result >> 8)
	*regLo = uint8(result & 0xFF)
	Trace.Printf("0x%04x: %s 0x%04x\n", ms.pc, instrName, result)
	ms.pc += 1
}

func DCX(instrName string, ms *MachineState, regHi *uint8, regLo *uint8) {
	result := getPair(*regHi, *regLo) - 1
	*regHi = uint8(result >> 8)
	*regLo = uint8(result & 0xFF)
	Trace.Printf("0x%04x: %s 0x%04x\n", ms.pc, instrName, result)
	ms.pc += 1
}

func PUSH(instrName string, ms *MachineState, regHi *uint8, regLo *uint8) {
	ms.WriteMem(ms.sp-2, []uint8{*regHi}, 1)
	ms.WriteMem(ms.sp-1, []uint8{*regLo}, 1)
	newSp := ms.sp - 2
	Trace.Printf("0x%04x: %s (0x%04x)<-0x%02x, (0x%04x)<-0x%02x, sp<-0x%04x\n",
		ms.pc, instrName, ms.sp-2, *regHi, ms.sp-1, *regLo, newSp)
	ms.pc += 1
	ms.sp = newSp
}

func POP(instrName string, ms *MachineState, regHi *uint8, regLo *uint8) {
	*regHi = ms.ReadMem(ms.sp, 1)[0]
	*regLo = ms.ReadMem(ms.sp+1, 1)[0]
	newSp := ms.sp + 2
	Trace.Printf("0x%04x: %s 0x%02x<-(0x%04x), 0x%02x<-(0x%04x), sp<-0x%04x\n",
		ms.pc, instrName, *regHi, ms.sp, *regLo, ms.sp+1, newSp)
	ms.pc += 1
	ms.sp = newSp
}

func DAD(instrName string, ms *MachineState, regHi *uint8, regLo *uint8) {
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

func RST(instrName string, ms *MachineState, addr uint16) {
	nextPC := ms.pc + 1
	pcHi := uint8(nextPC >> 8)
	pcLo := uint8(nextPC & 0xFF)
	ms.WriteMem(ms.sp-2, []uint8{pcLo, pcHi}, 2)
	Trace.Printf("0x%04x: %s 0x%04x\n", ms.pc, instrName, addr)
	ms.sp = ms.sp - 2
	ms.pc = addr
}

func CALL(instrName string, ms *MachineState, condFlagName string, condFlagVal bool) {
	byte1 := ms.ReadMem(ms.pc+1, 1)[0]
	byte2 := ms.ReadMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	PC := ms.pc
	nextPC := ms.pc + 3
	var syscall string
	var message string
	if !condFlagVal {
		ms.pc = nextPC
		syscall = ""
		message = ""

	} else if isSyscallAddress(adr) {
		offset := getPair(ms.regD, ms.regE)
		for i := 0; ; i++ {
			char := ms.ReadMem(offset+3+uint16(i), 1)[0]
			if char == '$' {
				message += "\n"
				break
			}
			message += fmt.Sprintf("%c", char)
		}
		ms.pc = nextPC
		syscall = "[SYSCALL]"

	} else {
		pcHi := uint8(nextPC >> 8)
		pcLo := uint8(nextPC & 0xFF)
		ms.WriteMem(ms.sp-2, []uint8{pcLo, pcHi}, 2)
		ms.sp = ms.sp - 2
		ms.pc = adr
		syscall = ""
		message = ""
	}
	if condFlagName != "" {
		Trace.Printf("0x%04x: %s 0x%04x, Taken=%t %s\n", PC, instrName, adr, condFlagVal, syscall)
	} else {
		Trace.Printf("0x%04x: %s 0x%04x %s\n", PC, instrName, adr, syscall)
	}
	if message != "" {
		fmt.Printf("%s", message)
		if strings.Contains(message, "CPU IS OPERATIONAL") {
			ms.Halt = true
		}
	}
}

func RET(instrName string, ms *MachineState, condFlagName string, condFlagVal bool) {
	currentPc := ms.pc
	bytes := ms.ReadMem(ms.sp, 2)
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

func INR(instrName string, ms *MachineState, reg *uint8) {
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

func DCR(instrName string, ms *MachineState, reg *uint8) {
	r := *reg
	*reg = *reg - 1
	ms.setZ(*reg)
	ms.setS(*reg)
	ms.setP(*reg)
	ms.setCY(uint(*reg)-uint(*reg) > math.MaxUint8)
	ms.setAC(*reg)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]-1, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, *reg, r, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ADD(instrName string, ms *MachineState, srcReg string, reg uint8) {
	regA := ms.regA
	ms.regA = regA + reg
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(reg) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]+reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, reg, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func SUB(instrName string, ms *MachineState, srcReg string, reg uint8) {
	regA := ms.regA
	ms.regA = regA - reg
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(reg) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]+reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, reg, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ADC(instrName string, ms *MachineState, srcReg string, reg uint8) {
	regA := ms.regA
	ms.regA = regA + reg
	var carry uint8 = 0
	if ms.flagCY {
		carry = 1
	}
	ms.regA += carry
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(reg)+uint(carry) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]+reg%s[0x%02x]+%d, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, reg, carry, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func SBB(instrName string, ms *MachineState, srcReg string, reg uint8) {
	regA := ms.regA
	ms.regA = regA - reg
	var carry uint8 = 0
	if ms.flagCY {
		carry = 1
	}
	ms.regA -= carry
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)-uint(reg)-uint(carry) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]-reg%s[0x%02x]-%d, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, reg, carry, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ANA(instrName string, ms *MachineState, srcReg string, reg uint8) {
	regA := ms.regA
	ms.regA = regA & reg
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)&uint(reg) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]&reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, reg, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func ORA(instrName string, ms *MachineState, srcReg string, reg uint8) {
	regA := ms.regA
	ms.regA = regA | reg
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)|uint(reg) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]|reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, reg, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func XRA(instrName string, ms *MachineState, srcReg string, reg uint8) {
	regA := ms.regA
	ms.regA = regA ^ reg
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)^uint(reg) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: %s regA[0x%02x]=regA[0x%02x]^reg%s[0x%02x], Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, regA, srcReg, reg, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func CMP(instrName string, ms *MachineState, srcReg string, reg uint8) {
	result := ms.regA - reg
	ms.setZ(result)
	ms.setS(result)
	ms.setP(result)
	ms.setCY(ms.regA < reg)
	ms.setAC(result)
	Trace.Printf("0x%04x: %s regA[0x%02x]-0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, instrName, ms.regA, reg, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func isSyscallAddress(adr uint16) bool {
	return adr == 0x5
}
