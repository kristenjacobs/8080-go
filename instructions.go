package main

import (
	"math"
)

func instr_0x00_NOP(ms *machineState) {
	// 1
	Trace.Printf("0x%04x: 0x00_NOP\n", ms.pc)
	ms.pc += 1
}

func instr_0x01_LXI_B_D16(ms *machineState) {
	// 3		B <- byte 3, C <- byte 2
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	ms.regB = byte3
	ms.regC = byte2
	Trace.Printf("0x%04x: 0x01_LXI_B_D16 0x%02x 0x%02x\n", ms.pc, ms.regB, ms.regC)
	ms.pc += 3
}

func instr_0x02_STAX_B(ms *machineState) {
	// 1		(BC) <- A
	STAX("0x02_STAX_B", ms, &ms.regB, &ms.regC)
}

func instr_0x03_INX_B(ms *machineState) {
	// 1		BC <- BC+1
	INX("0x03_INX_B", ms, &ms.regB, &ms.regC)
}

func instr_0x04_INR_B(ms *machineState) {
	// 1	Z, S, P, AC	B <- B+1
	INR("0x04_INR_B", ms, &ms.regB)
}

func instr_0x05_DCR_B(ms *machineState) {
	// 1	Z, S, P, AC	B <- B-1
	DCR("0x05_DCR_B", ms, &ms.regB)
}

func instr_0x06_MVI_B_D8(ms *machineState) {
	// 2		B <- byte 2
	MVI("0x06_MVI_B_D8", ms, &ms.regB)
}

func instr_0x07_RLC(ms *machineState) {
	// 1       CY      A = A << 1; bit 0 = prev bit 7; CY = prev bit 7
	panic("Unimplemented")
}

func instr_0x09_DAD_B(ms *machineState) {
	// 1	CY	HL = HL + BC
	DAD("0x09_DAD_B", ms, &ms.regB, &ms.regC)
}

func instr_0x0a_LDAX_B(ms *machineState) {
	// 1		A <- (BC)
	LDAX("0x0a_LDAX_B", ms, &ms.regB, &ms.regC)
}

func instr_0x0b_DCX_B(ms *machineState) {
	// 1		BC = BC-1
	DCX("0x0b_DCX_B", ms, &ms.regB, &ms.regC)
}

func instr_0x0c_INR_C(ms *machineState) {
	// 1	Z, S, P, AC	C <- C+1
	INR("0x0c_INR_C", ms, &ms.regC)
}

func instr_0x0d_DCR_C(ms *machineState) {
	// 1	Z, S, P, AC	C <-C-1
	DCR("0x0d_DCR_C", ms, &ms.regC)
}

func instr_0x0e_MVI_C_D8(ms *machineState) {
	// 2		C <- byte 2
	MVI("0x0e_MVI_C_D8", ms, &ms.regC)
}

func instr_0x0f_RRC(ms *machineState) {
	// 1	CY	A = A >> 1; bit 7 = prev bit 0; CY = prev bit 0
	panic("Unimplemented")
}

func instr_0x11_LXI_D_D16(ms *machineState) {
	// 3		D <- byte 3, E <- byte 2
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	ms.regD = byte3
	ms.regE = byte2
	Trace.Printf("0x%04x: 0x11_LXI_D_D16 0x%02x 0x%02x\n", ms.pc, ms.regD, ms.regE)
	ms.pc += 3
}

func instr_0x12_STAX_D(ms *machineState) {
	// 1		(DE) <- A
	STAX("0x12_STAX_D", ms, &ms.regD, &ms.regE)
}

func instr_0x13_INX_D(ms *machineState) {
	// 1		DE <- DE + 1
	INX("0x13_INX_D", ms, &ms.regD, &ms.regE)
}

func instr_0x14_INR_D(ms *machineState) {
	// 1	Z, S, P, AC	D <- D+1
	INR("0x14_INR_D", ms, &ms.regD)
}

func instr_0x15_DCR_D(ms *machineState) {
	// 1	Z, S, P, AC	D <- D-1
	DCR("0x15_DCR_D", ms, &ms.regD)
}

func instr_0x16_MVI_D_D8(ms *machineState) {
	// 2		D <- byte 2
	MVI("0x16_MVI_D_D8", ms, &ms.regD)
}

func instr_0x17_RAL(ms *machineState) {
	// 1	CY	A = A << 1; bit 0 = prev CY; CY = prev bit 7
	panic("Unimplemented")
}

func instr_0x19_DAD_D(ms *machineState) {
	// 1	CY	HL = HL + DE
	DAD("0x19_DAD_D", ms, &ms.regD, &ms.regE)
}

func instr_0x1a_LDAX_D(ms *machineState) {
	// 1		A <- (DE)
	LDAX("0x1a_LDAX_D", ms, &ms.regD, &ms.regE)
}

func instr_0x1b_DCX_D(ms *machineState) {
	// 1		DE = DE-1
	DCX("0x1b_DCX_D", ms, &ms.regD, &ms.regE)
}

func instr_0x1c_INR_E(ms *machineState) {
	// 1	Z, S, P, AC	E <-E+1
	INR("0x1c_INR_E", ms, &ms.regE)
}

func instr_0x1d_DCR_E(ms *machineState) {
	// 1	Z, S, P, AC	E <- E-1
	DCR("0x1d_DCR_E", ms, &ms.regE)
}

func instr_0x1e_MVI_E_D8(ms *machineState) {
	// 2		E <- byte 2
	MVI("0x1e_MVI_E_D8", ms, &ms.regE)
}

func instr_0x1f_RAR(ms *machineState) {
	// 1	CY	A = A >> 1; bit 7 = prev bit 7; CY = prev bit 0
	panic("Unimplemented")
}

func instr_0x20_RIM(ms *machineState) {
	// 1		special
	Trace.Printf("0x%04x: 0x20_RIM\n", ms.pc)
	ms.pc += 1
}

func instr_0x21_LXI_H_D16(ms *machineState) {
	// 3		H <- byte 3, L <- byte 2
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	ms.regH = byte3
	ms.regL = byte2
	Trace.Printf("0x%04x: 0x21_LXI_H_D16 0x%02x 0x%02x\n", ms.pc, ms.regH, ms.regL)
	ms.pc += 3
}

func instr_0x22_SHLD_adr(ms *machineState) {
	// 3		(adr) <-L; (adr+1)<-H
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte3) << 8) | uint16(byte2)
	ms.writeMem(adr, []uint8{ms.regL}, 1)
	ms.writeMem(adr+1, []uint8{ms.regH}, 1)
	Trace.Printf("0x%04x: 0x22_SHLD_adr (0x%04x) = regL[0x%02x], (0x%04x) = regH[0x%02x]\n",
		ms.pc, adr, ms.regL, adr+1, ms.regH)
	ms.pc += 3
}

func instr_0x23_INX_H(ms *machineState) {
	// 1		HL <- HL + 1
	INX("0x23_INX_H", ms, &ms.regH, &ms.regL)
}

func instr_0x24_INR_H(ms *machineState) {
	// 1	Z, S, P, AC	H <- H+1
	INR("0x24_INR_H", ms, &ms.regH)
}

func instr_0x25_DCR_H(ms *machineState) {
	// 1	Z, S, P, AC	H <- H-1
	DCR("0x25_DCR_H", ms, &ms.regH)
}

func instr_0x26_MVI_H_D8(ms *machineState) {
	// 2		H <- byte 2
	MVI("0x26_MVI_H_D8", ms, &ms.regH)
}

func instr_0x27_DAA(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0x29_DAD_H(ms *machineState) {
	// 1	CY	HL = HL + HL
	DAD("0x29_DAD_H", ms, &ms.regH, &ms.regL)
}

func instr_0x2a_LHLD_adr(ms *machineState) {
	// 3		L <- (adr); H<-(adr+1)
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	ms.regL = ms.readMem(adr, 1)[0]
	ms.regH = ms.readMem(adr+1, 1)[0]
	Trace.Printf("0x%04x: 0x2a_LHLD_adr regL[0x%02x] = (0x%04x), regH[0x%02x] = (0x%04x)\n",
		ms.pc, ms.regL, adr, ms.regH, adr+1)
	ms.pc += 3
}

func instr_0x2b_DCX_H(ms *machineState) {
	// 1		HL = HL-1
	DCX("0x2b_DCX_H", ms, &ms.regH, &ms.regL)
}

func instr_0x2c_INR_L(ms *machineState) {
	// 1	Z, S, P, AC	L <- L+1
	INR("0x2c_INR_L", ms, &ms.regL)
}

func instr_0x2d_DCR_L(ms *machineState) {
	// 1	Z, S, P, AC	L <- L-1
	DCR("0x2d_DCR_L", ms, &ms.regL)
}

func instr_0x2e_MVI_L_D8(ms *machineState) {
	// 2		L <- byte 2
	MVI("0x2e_MVI_L_D8", ms, &ms.regL)
}

func instr_0x2f_CMA(ms *machineState) {
	// 1		A <- !A
	panic("Unimplemented")
}

func instr_0x30_SIM(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0x31_LXI_SP_D16(ms *machineState) {
	// 3		SP.hi <- byte 3, SP.lo <- byte 2
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	var sp uint16 = (uint16(byte3) << 8) | uint16(byte2)
	Trace.Printf("0x%04x: 0x31_LXI_SP_D16 0x%04x\n", ms.pc, sp)
	ms.sp = sp
	ms.pc += 3
}

func instr_0x32_STA_adr(ms *machineState) {
	// 3		(adr) <- A
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	var addr uint16 = (uint16(byte3) << 8) | uint16(byte2)
	ms.writeMem(addr, []uint8{ms.regA}, 1)
	Trace.Printf("0x%04x: 0x32_STA_adr (0x%04x) = regA[0x%02x]\n", ms.pc, addr, ms.regA)
	ms.pc += 3
}

func instr_0x33_INX_SP(ms *machineState) {
	// 1		SP = SP + 1
	ms.sp = ms.sp + 1
	Trace.Printf("0x%04x: 0x33_INX_SP 0x%04x\n", ms.pc, ms.sp)
	ms.pc += 1
}

func instr_0x34_INR_M(ms *machineState) {
	// 1	Z, S, P, AC	(HL) <- (HL)+1
	op := ms.getM()
	res := op + 1
	ms.setM(res)
	ms.setZ(res)
	ms.setS(res)
	ms.setP(res)
	ms.setCY(uint(res)+uint(1) > math.MaxUint8)
	ms.setAC(res)
	Trace.Printf("0x%04x: 0x34_INR_M M[0x%02x]=M[0x%02x]+1, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, res, op, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func instr_0x35_DCR_M(ms *machineState) {
	// 1	Z, S, P, AC	(HL) <- (HL)-1
	op := ms.getM()
	res := op - 1
	ms.setM(res)
	ms.setZ(res)
	ms.setS(res)
	ms.setP(res)
	ms.setCY(uint(res)-uint(1) > math.MaxUint8)
	ms.setAC(res)
	Trace.Printf("0x%04x: 0x34_DCR_M M[0x%02x]=M[0x%02x]-1, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, res, op, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 1
}

func instr_0x36_MVI_M_D8(ms *machineState) {
	// 2		(HL) <- byte 2
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	addr := getPair(ms.regH, ms.regL)
	ms.writeMem(addr, []uint8{byte2}, 1)
	Trace.Printf("0x%04x: 0x36_MVI_M_D8 [0x%04x] 0x%02x\n", ms.pc, addr, byte2)
	ms.pc += 2
}

func instr_0x37_STC(ms *machineState) {
	// 1	CY	CY = 1
	panic("Unimplemented")
}

func instr_0x39_DAD(ms *machineState) {
	// SP	1	CY	HL = HL + SP
	panic("Unimplemented")
}

func instr_0x3a_LDA_adr(ms *machineState) {
	// 3		A <- (adr)
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	ms.regA = ms.readMem(adr, 1)[0]
	Trace.Printf("0x%04x: 0x3a_LDA_adr regA[0x%02x] (0x%04x)\n", ms.pc, ms.regA, adr)
	ms.pc += 3
}

func instr_0x3b_DCX_SP(ms *machineState) {
	// SP	1		SP = SP-1
	ms.sp = ms.sp - 1
	Trace.Printf("0x%04x: 0x3b_DCX_SP 0x%04x\n", ms.pc, ms.sp)
	ms.pc += 1
}

func instr_0x3c_INR_A(ms *machineState) {
	// 1	Z, S, P, AC	A <- A+1
	INR("0x3c_INR_A", ms, &ms.regA)
}

func instr_0x3d_DCR_A(ms *machineState) {
	// 1	Z, S, P, AC	A <- A-1
	DCR("0x3d_DCR_A", ms, &ms.regA)
}

func instr_0x3e_MVI_A_D8(ms *machineState) {
	// 2		A <- byte 2
	MVI("0x3e_MVI_A_D8", ms, &ms.regA)
}

func instr_0x3f_CMC(ms *machineState) {
	// 1	CY	CY=!CY
	panic("Unimplemented")
}

func instr_0x40_MOV_B_B(ms *machineState) {
	// 1		B <- B
	MOV_REG_REG("0x40_MOV_B_B", ms, &ms.regB, &ms.regB)
}

func instr_0x41_MOV_B_C(ms *machineState) {
	// 1		B <- C
	MOV_REG_REG("0x41_MOV_B_C", ms, &ms.regB, &ms.regC)
}

func instr_0x42_MOV_B_D(ms *machineState) {
	// 1		B <- D
	MOV_REG_REG("0x42_MOV_B_D", ms, &ms.regB, &ms.regD)
}

func instr_0x43_MOV_B_E(ms *machineState) {
	// 1		B <- E
	MOV_REG_REG("0x43_MOV_B_E", ms, &ms.regB, &ms.regE)
}

func instr_0x44_MOV_B_H(ms *machineState) {
	// 1		B <- H
	MOV_REG_REG("0x44_MOV_B_H", ms, &ms.regB, &ms.regH)
}

func instr_0x45_MOV_B_L(ms *machineState) {
	// 1		B <- L
	MOV_REG_REG("0x45_MOV_B_L", ms, &ms.regB, &ms.regL)
}

func instr_0x46_MOV_B_M(ms *machineState) {
	// 1		B <- (HL)
	MOV_REG_MEM("0x46_MOV_B_M", ms, &ms.regB)
}

func instr_0x47_MOV_B_A(ms *machineState) {
	// 1		B <- A
	MOV_REG_REG("0x47_MOV_B_A", ms, &ms.regB, &ms.regA)
}

func instr_0x48_MOV_C_B(ms *machineState) {
	// 1		C <- B
	MOV_REG_REG("0x48_MOV_C_B", ms, &ms.regC, &ms.regB)
}

func instr_0x49_MOV_C_C(ms *machineState) {
	// 1		C <- C
	MOV_REG_REG("0x49_MOV_C_C", ms, &ms.regC, &ms.regC)
}

func instr_0x4a_MOV_C_D(ms *machineState) {
	// 1		C <- D
	MOV_REG_REG("0x4a_MOV_C_D", ms, &ms.regC, &ms.regD)
}

func instr_0x4b_MOV_C_E(ms *machineState) {
	// 1		C <- E
	MOV_REG_REG("0x4b_MOV_C_E", ms, &ms.regC, &ms.regE)
}

func instr_0x4c_MOV_C_H(ms *machineState) {
	// 1		C <- H
	MOV_REG_REG("0x4c_MOV_C_H", ms, &ms.regC, &ms.regH)
}

func instr_0x4d_MOV_C_L(ms *machineState) {
	// 1		C <- L
	MOV_REG_REG("0x4d_MOV_C_L", ms, &ms.regC, &ms.regL)
}

func instr_0x4e_MOV_C_M(ms *machineState) {
	// 1		C <- (HL)
	MOV_REG_MEM("0x4e_MOV_C_M", ms, &ms.regC)
}

func instr_0x4f_MOV_C_A(ms *machineState) {
	// 1		C <- A
	MOV_REG_REG("0x4f_MOV_C_A", ms, &ms.regC, &ms.regA)
}

func instr_0x50_MOV_D_B(ms *machineState) {
	// 1		D <- B
	MOV_REG_REG("0x50_MOV_D_B", ms, &ms.regD, &ms.regB)
}

func instr_0x51_MOV_D_C(ms *machineState) {
	// 1		D <- C
	MOV_REG_REG("0x51_MOV_D_C", ms, &ms.regD, &ms.regC)
}

func instr_0x52_MOV_D_D(ms *machineState) {
	// 1		D <- D
	MOV_REG_REG("0x52_MOV_D_D", ms, &ms.regD, &ms.regD)
}

func instr_0x53_MOV_D_E(ms *machineState) {
	// 1		D <- E
	MOV_REG_REG("0x53_MOV_D_E", ms, &ms.regD, &ms.regE)
}

func instr_0x54_MOV_D_H(ms *machineState) {
	// 1		D <- H
	MOV_REG_REG("0x54_MOV_D_H", ms, &ms.regD, &ms.regH)
}

func instr_0x55_MOV_D_L(ms *machineState) {
	// 1		D <- L
	MOV_REG_REG("0x55_MOV_D_L", ms, &ms.regD, &ms.regL)
}

func instr_0x56_MOV_D_M(ms *machineState) {
	// 1		D <- (HL)
	MOV_REG_MEM("0x56_MOV_D_M", ms, &ms.regD)
}

func instr_0x57_MOV_D_A(ms *machineState) {
	// 1		D <- A
	MOV_REG_REG("0x57_MOV_D_A", ms, &ms.regD, &ms.regA)
}

func instr_0x58_MOV_E_B(ms *machineState) {
	// 1		E <- B
	MOV_REG_REG("0x58_MOV_E_B", ms, &ms.regE, &ms.regB)
}

func instr_0x59_MOV_E_C(ms *machineState) {
	// 1		E <- C
	MOV_REG_REG("0x59_MOV_E_C", ms, &ms.regE, &ms.regC)
}

func instr_0x5a_MOV_E_D(ms *machineState) {
	// 1		E <- D
	MOV_REG_REG("0x5a_MOV_E_D", ms, &ms.regE, &ms.regD)
}

func instr_0x5b_MOV_E_E(ms *machineState) {
	// 1		E <- E
	MOV_REG_REG("0x5b_MOV_E_E", ms, &ms.regE, &ms.regE)
}

func instr_0x5c_MOV_E_H(ms *machineState) {
	// 1		E <- H
	MOV_REG_REG("0x5c_MOV_E_H", ms, &ms.regE, &ms.regH)
}

func instr_0x5d_MOV_E_L(ms *machineState) {
	// 1		E <- L
	MOV_REG_REG("0x5d_MOV_E_L", ms, &ms.regE, &ms.regL)
}

func instr_0x5e_MOV_E_M(ms *machineState) {
	// 1		E <- (HL)
	MOV_REG_MEM("0x5e_MOV_E_M", ms, &ms.regE)
}

func instr_0x5f_MOV_E_A(ms *machineState) {
	// 1		E <- A
	MOV_REG_REG("0x5f_MOV_E_A", ms, &ms.regE, &ms.regA)
}

func instr_0x60_MOV_H_B(ms *machineState) {
	// 1		H <- B
	MOV_REG_REG("0x60_MOV_H_B", ms, &ms.regH, &ms.regB)
}

func instr_0x61_MOV_H_C(ms *machineState) {
	// 1		H <- C
	MOV_REG_REG("0x61_MOV_H_C", ms, &ms.regH, &ms.regC)
}

func instr_0x62_MOV_H_D(ms *machineState) {
	// 1		H <- D
	MOV_REG_REG("0x62_MOV_H_D", ms, &ms.regH, &ms.regD)
}

func instr_0x63_MOV_H_E(ms *machineState) {
	// 1		H <- E
	MOV_REG_REG("0x63_MOV_H_E", ms, &ms.regH, &ms.regE)
}

func instr_0x64_MOV_H_H(ms *machineState) {
	// 1		H <- H
	MOV_REG_REG("0x64_MOV_H_H", ms, &ms.regH, &ms.regH)
}

func instr_0x65_MOV_H_L(ms *machineState) {
	// 1		H <- L
	MOV_REG_REG("0x65_MOV_H_L", ms, &ms.regH, &ms.regL)
}

func instr_0x66_MOV_H_M(ms *machineState) {
	// 1		H <- (HL)
	MOV_REG_MEM("0x66_MOV_H_M", ms, &ms.regH)
}

func instr_0x67_MOV_H_A(ms *machineState) {
	// 1		H <- A
	MOV_REG_REG("0x67_MOV_H_A", ms, &ms.regH, &ms.regA)
}

func instr_0x68_MOV_L_B(ms *machineState) {
	// 1		L <- B
	MOV_REG_REG("0x68_MOV_L_B", ms, &ms.regL, &ms.regB)
}

func instr_0x69_MOV_L_C(ms *machineState) {
	// 1		L <- C
	MOV_REG_REG("0x69_MOV_L_C", ms, &ms.regL, &ms.regC)
}

func instr_0x6a_MOV_L_D(ms *machineState) {
	// 1		L <- D
	MOV_REG_REG("0x6a_MOV_L_D", ms, &ms.regL, &ms.regD)
}

func instr_0x6b_MOV_L_E(ms *machineState) {
	// 1		L <- E
	MOV_REG_REG("0x6b_MOV_L_E", ms, &ms.regL, &ms.regE)
}

func instr_0x6c_MOV_L_H(ms *machineState) {
	// 1		L <- H
	MOV_REG_REG("0x6c_MOV_L_H", ms, &ms.regL, &ms.regH)
}

func instr_0x6d_MOV_L_L(ms *machineState) {
	// 1		L <- L
	MOV_REG_REG("0x6d_MOV_L_L", ms, &ms.regL, &ms.regL)
}

func instr_0x6e_MOV_L_M(ms *machineState) {
	// 1		L <- (HL)
	MOV_REG_MEM("0x6e_MOV_L_M", ms, &ms.regL)
}

func instr_0x6f_MOV_L_A(ms *machineState) {
	// 1		L <- A
	MOV_REG_REG("0x6f_MOV_L_A", ms, &ms.regL, &ms.regA)
}

func instr_0x70_MOV_M_B(ms *machineState) {
	// 1		(HL) <- B
	MOV_MEM_REG("0x70_MOV_M_B", ms, &ms.regB)
}

func instr_0x71_MOV_M_C(ms *machineState) {
	// 1		(HL) <- C
	MOV_MEM_REG("0x71_MOV_M_C", ms, &ms.regC)
}

func instr_0x72_MOV_M_D(ms *machineState) {
	// 1		(HL) <- D
	MOV_MEM_REG("0x72_MOV_M_D", ms, &ms.regD)
}

func instr_0x73_MOV_M_E(ms *machineState) {
	// 1		(HL) <- E
	MOV_MEM_REG("0x73_MOV_M_E", ms, &ms.regE)
}

func instr_0x74_MOV_M_H(ms *machineState) {
	// 1		(HL) <- H
	MOV_MEM_REG("0x74_MOV_M_H", ms, &ms.regH)
}

func instr_0x75_MOV_M_L(ms *machineState) {
	// 1		(HL) <- L
	MOV_MEM_REG("0x75_MOV_M_L", ms, &ms.regL)
}

func instr_0x76_HLT(ms *machineState) {
	// 1		special
	Trace.Printf("0x%04x: 0x76_HLT\n", ms.pc)
	ms.halt = true
}

func instr_0x77_MOV_M_A(ms *machineState) {
	// 1		(HL) <- A
	MOV_MEM_REG("0x77_MOV_M_A", ms, &ms.regA)
}

func instr_0x78_MOV_A_B(ms *machineState) {
	// 1		A <- B
	MOV_REG_REG("0x78_MOV_A_B", ms, &ms.regA, &ms.regB)
}

func instr_0x79_MOV_A_C(ms *machineState) {
	// 1		A <- C
	MOV_REG_REG("0x79_MOV_A_C", ms, &ms.regA, &ms.regC)
}

func instr_0x7a_MOV_A_D(ms *machineState) {
	// 1		A <- D
	MOV_REG_REG("0x7a_MOV_A_D", ms, &ms.regA, &ms.regD)
}

func instr_0x7b_MOV_A_E(ms *machineState) {
	// 1		A <- E
	MOV_REG_REG("0x7b_MOV_A_E", ms, &ms.regA, &ms.regE)
}

func instr_0x7c_MOV_A_H(ms *machineState) {
	// 1		A <- H
	MOV_REG_REG("0x7c_MOV_A_H", ms, &ms.regA, &ms.regH)
}

func instr_0x7d_MOV_A_L(ms *machineState) {
	// 1		A <- L
	MOV_REG_REG("0x7d_MOV_A_L", ms, &ms.regA, &ms.regL)
}

func instr_0x7e_MOV_A_M(ms *machineState) {
	// 1		A <- (HL)
	MOV_REG_MEM("0x7e_MOV_A_M", ms, &ms.regA)
}

func instr_0x7f_MOV_A_A(ms *machineState) {
	// 1		A <- A
	MOV_REG_REG("0x7f_MOV_A_A", ms, &ms.regA, &ms.regA)
}

func instr_0x80_ADD_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + B
	ADD("0x80_ADD_B", ms, "B", ms.regB)
}

func instr_0x81_ADD_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + C
	ADD("0x81_ADD_C", ms, "C", ms.regC)
}

func instr_0x82_ADD_D(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + D
	ADD("0x82_ADD_D", ms, "D", ms.regD)
}

func instr_0x83_ADD_E(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + E
	ADD("0x83_ADD_E", ms, "E", ms.regE)
}

func instr_0x84_ADD_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + H
	ADD("0x84_ADD_H", ms, "H", ms.regH)
}

func instr_0x85_ADD_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + L
	ADD("0x85_ADD_L", ms, "L", ms.regL)
}

func instr_0x86_ADD_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + (HL)
	ADD("0x86_ADD_M", ms, "M", ms.getM())
}

func instr_0x87_ADD_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + A
	ADD("0x87_ADD_A", ms, "A", ms.regA)
}

func instr_0x88_ADC_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + B + CY
	ADC("0x88_ADC_B", ms, "B", ms.regB)
}

func instr_0x89_ADC_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + C + CY
	ADC("0x89_ADC_C", ms, "C", ms.regC)
}

func instr_0x8a_ADC_D(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + D + CY
	ADC("0x8a_ADC_D", ms, "D", ms.regD)
}

func instr_0x8b_ADC_E(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + E + CY
	ADC("0x8b_ADC_E", ms, "E", ms.regE)
}

func instr_0x8c_ADC_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + H + CY
	ADC("0x8c_ADC_H", ms, "H", ms.regH)
}

func instr_0x8d_ADC_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + L + CY
	ADC("0x8d_ADC_L", ms, "L", ms.regL)
}

func instr_0x8e_ADC_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + (HL) + CY
	ADC("0x8e_ADC_M", ms, "M", ms.getM())
}

func instr_0x8f_ADC_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + A + CY
	ADC("0x8f_ADC_A", ms, "A", ms.regA)
}

func instr_0x90_SUB_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - B
	SUB("0x90_SUB_B", ms, "B", ms.regB)
}

func instr_0x91_SUB_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - C
	SUB("0x91_SUB_C", ms, "C", ms.regC)
}

func instr_0x92_SUB_D(ms *machineState) {
	// 1       Z, S, P, CY, AC A <- A + D
	SUB("0x92_SUB_D", ms, "D", ms.regD)
}

func instr_0x93_SUB_E(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - E
	SUB("0x93_SUB_E", ms, "E", ms.regE)
}

func instr_0x94_SUB_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + H
	SUB("0x94_SUB_H", ms, "H", ms.regH)
}

func instr_0x95_SUB_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - L
	SUB("0x95_SUB_L", ms, "L", ms.regL)
}

func instr_0x96_SUB_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A + (HL)
	SUB("0x96_SUB_L", ms, "M", ms.getM())
}

func instr_0x97_SUB_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - A
	SUB("0x97_SUB_A", ms, "A", ms.regA)
}

func instr_0x98_SBB_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - B - CY
	SBB("0x98_SBB_B", ms, "B", ms.regB)
}

func instr_0x99_SBB_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - C - CY
	SBB("0x99_SBB_C", ms, "C", ms.regC)
}

func instr_0x9a_SBB_D(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - D - CY
	SBB("0x9a_SBB_D", ms, "D", ms.regD)
}

func instr_0x9b_SBB_E(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - E - CY
	SBB("0x9b_SBB_E", ms, "E", ms.regE)
}

func instr_0x9c_SBB_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - H - CY
	SBB("0x9c_SBB_H", ms, "H", ms.regH)
}

func instr_0x9d_SBB_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - L - CY
	SBB("0x9d_SBB_L", ms, "L", ms.regL)
}

func instr_0x9e_SBB_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - (HL) - CY
	SBB("0x9e_SBB_M", ms, "M", ms.getM())
}

func instr_0x9f_SBB_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A - A - CY
	SBB("0x9f_SBB_A", ms, "A", ms.regA)
}

func instr_0xa0_ANA_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & B
	ANA("0xa0_ANA_B", ms, "B", ms.regB)
}

func instr_0xa1_ANA_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & C
	ANA("0xa1_ANA_C", ms, "C", ms.regC)
}

func instr_0xa2_ANA_D(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & D
	ANA("0xa2_ANA_D", ms, "D", ms.regD)
}

func instr_0xa3_ANA_E(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & E
	ANA("0xa3_ANA_E", ms, "E", ms.regE)
}

func instr_0xa4_ANA_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & H
	ANA("0xa4_ANA_H", ms, "H", ms.regH)
}

func instr_0xa5_ANA_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & L
	ANA("0xa5_ANA_L", ms, "L", ms.regL)
}

func instr_0xa6_ANA_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & (HL)
	ANA("0xa6_ANA_M", ms, "M", ms.getM())
}

func instr_0xa7_ANA_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A & A
	ANA("0xa7_ANA_A", ms, "A", ms.regA)
}

func instr_0xa8_XRA_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ B
	XRA("0xa8_XRA_B", ms, "B", ms.regB)
}

func instr_0xa9_XRA_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ C
	XRA("0xa9_XRA_C", ms, "C", ms.regC)
}

func instr_0xaa_XRA_D(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ D
	XRA("0xaa_XRA_D", ms, "D", ms.regD)
}

func instr_0xab_XRA_E(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ E
	XRA("0xab_XRA_E", ms, "E", ms.regE)
}

func instr_0xac_XRA_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ H
	XRA("0xac_XRA_H", ms, "H", ms.regH)
}

func instr_0xad_XRA_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ L
	XRA("0xad_XRA_L", ms, "L", ms.regL)
}

func instr_0xae_XRA_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ (HL)
	XRA("0xae_XRA_M", ms, "M", ms.getM())
}

func instr_0xaf_XRA_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A ^ A
	XRA("0xaf_XRA_A", ms, "A", ms.regA)
}

func instr_0xb0_ORA_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | B
	ORA("0xb0_ORA_B", ms, "B", ms.regB)
}

func instr_0xb1_ORA_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | C
	ORA("0xb1_ORA_C", ms, "C", ms.regC)
}

func instr_0xb2_ORA_D(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | D
	ORA("0xb2_ORA_D", ms, "D", ms.regD)
}

func instr_0xb3_ORA_E(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | E
	ORA("0xb3_ORA_E", ms, "E", ms.regE)
}

func instr_0xb4_ORA_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | H
	ORA("0xb4_ORA_H", ms, "H", ms.regH)
}

func instr_0xb5_ORA_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | L
	ORA("0xb5_ORA_L", ms, "L", ms.regL)
}

func instr_0xb6_ORA_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | (HL)
	ORA("0xb6_ORA_M", ms, "M", ms.getM())
}

func instr_0xb7_ORA_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A <- A | A
	ORA("0xb7_ORA_A", ms, "A", ms.regA)
}

func instr_0xb8_CMP_B(ms *machineState) {
	// 1	Z, S, P, CY, AC	A - B
	CMP("0xb8_CMP_B", ms, "B", ms.regB)
}

func instr_0xb9_CMP_C(ms *machineState) {
	// 1	Z, S, P, CY, AC	A - C
	CMP("0xb9_CMP_C", ms, "C", ms.regC)
}

func instr_0xba_CMP_D(ms *machineState) {
	// 1	Z, S, P, CY, AC	A - D
	CMP("0xba_CMP_D", ms, "D", ms.regD)
}

func instr_0xbb_CMP_E(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A - E
	CMP("0xbb_CMP_E", ms, "E", ms.regE)
}

func instr_0xbc_CMP_H(ms *machineState) {
	// 1	Z, S, P, CY, AC	A - H
	CMP("0xbc_CMP_H", ms, "H", ms.regH)
}

func instr_0xbd_CMP_L(ms *machineState) {
	// 1	Z, S, P, CY, AC	A - L
	CMP("0xbd_CMP_L", ms, "L", ms.regL)
}

func instr_0xbe_CMP_M(ms *machineState) {
	// 1	Z, S, P, CY, AC	A - (HL)
	CMP("0xbe_CMP_M", ms, "M", ms.getM())
}

func instr_0xbf_CMP_A(ms *machineState) {
	// 1	Z, S, P, CY, AC	A - A
	CMP("0xbf_CMP_A", ms, "A", ms.regA)
}

func instr_0xc0_RNZ(ms *machineState) {
	// 1		if NZ, RET
	RET("0xc0_RNZ", ms, "Z", !ms.flagZ)
}

func instr_0xc1_POP_B(ms *machineState) {
	// 1		C <- (sp); B <- (sp+1); sp <- sp+2
	POP("0xc1_POP_B", ms, &ms.regC, &ms.regB)
}

func instr_0xc2_JNZ_adr(ms *machineState) {
	// 3		if NZ, PC <- adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xc2_JNZ_adr 0x%04x, Z=%t\n", ms.pc, adr, ms.flagZ)
	if !ms.flagZ {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xc3_JMP_adr(ms *machineState) {
	// 3		PC <= adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xc3_JMP_adr 0x%04x\n", ms.pc, adr)
	ms.pc = adr
}

func instr_0xc4_CNZ_adr(ms *machineState) {
	// 3		if NZ, CALL adr
	CALL("0xc4_CNZ_adr", ms, "Z", !ms.flagZ)
}

func instr_0xc5_PUSH_B(ms *machineState) {
	// B	1		(sp-2)<-C; (sp-1)<-B; sp <- sp - 2
	PUSH("0xc5_PUSH_B", ms, &ms.regC, &ms.regB)
}

func instr_0xc6_ADI_D8(ms *machineState) {
	// 2	Z, S, P, CY, AC	A <- A + byte
	data := ms.readMem(ms.pc+1, 1)[0]
	regA := ms.regA
	ms.regA = regA + data
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(data) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: 0xc6_ADI_D8 regA[0x%02x]=regA[0x%02x]+0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, regA, data, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xc7_RST_0(ms *machineState) {
	// 1		CALL $0
	RST("instr_0xc7_RST_0", ms, 0)
}

func instr_0xc8_RZ(ms *machineState) {
	// 1		if Z, RET
	RET("0xc8_RZ", ms, "Z", ms.flagZ)
}

func instr_0xc9_RET(ms *machineState) {
	// 1		PC.lo <- (sp); PC.hi<-(sp+1); SP <- SP+2
	RET("0xc9_RET", ms, "", true)
}

func instr_0xca_JZ_adr(ms *machineState) {
	// adr	3		if Z, PC <- adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xca_JZ_adr Z=%t 0x%04x\n", ms.pc, ms.flagZ, adr)
	if ms.flagZ {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xcc_CZ_adr(ms *machineState) {
	// 3		if Z, CALL adr
	CALL("0xcc_CZ_adr", ms, "Z", ms.flagZ)
}

func instr_0xcd_CALL_adr(ms *machineState) {
	// 3		(SP-1)<-PC.hi;(SP-2)<-PC.lo;SP<-SP-2;PC=adr
	CALL("0xcd_CALL_adr", ms, "", true)
}

func instr_0xce_ACI_D8(ms *machineState) {
	// 2	Z, S, P, CY, AC	A <- A + data + CY
	data := ms.readMem(ms.pc+1, 1)[0]
	regA := ms.regA
	var carry uint8
	if ms.flagCY {
		carry = 1
	} else {
		carry = 0
	}
	ms.regA = regA + data + carry
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(data)+uint(carry) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: 0xce_ACI_D8 regA[0x%02x]=regA[0x%02x]+0x%02x+0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, regA, data, carry, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xcf_RST_1(ms *machineState) {
	// 1		CALL $8
	RST("instr_0xcf_RST_1", ms, 0x8)
}

func instr_0xd0_RNC(ms *machineState) {
	// 1		if NCY, RET
	RET("0xd0_RNC", ms, "CY", !ms.flagCY)
}

func instr_0xd1_POP_D(ms *machineState) {
	// 1		E <- (sp); D <- (sp+1); sp <- sp+2
	POP("0xd1_POP_D", ms, &ms.regE, &ms.regD)
}

func instr_0xd2_JNC_adr(ms *machineState) {
	// 3		if NCY, PC<-adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xd2_JNC_adr 0x%04x, C=%t\n", ms.pc, adr, ms.flagCY)
	if !ms.flagCY {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xd3_OUT_D8(ms *machineState) {
	// 2		special
	port := ms.readMem(ms.pc+1, 1)[0]
	Trace.Printf("0x%04x: 0xd3_OUT_D8 %d\n", ms.pc, port)
	ms.pc += 2
}

func instr_0xd4_CNC_adr(ms *machineState) {
	// 3		if NCY, CALL adr
	CALL("0xd4_CNC_adr", ms, "CY", !ms.flagCY)
}

func instr_0xd5_PUSH_D(ms *machineState) {
	// D	1		(sp-2)<-E; (sp-1)<-D; sp <- sp - 2
	PUSH("0xd5_PUSH_D", ms, &ms.regE, &ms.regD)
}

func instr_0xd6_SUI_D8(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A - data
	data := ms.readMem(ms.pc+1, 1)[0]
	regA := ms.regA
	ms.regA = regA - data
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)-uint(data) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: 0xd6_SUI_D8 regA[0x%02x]=regA[0x%02x]-0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, regA, data, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xd7_RST_2(ms *machineState) {
	// 1		CALL $10
	RST("instr_0xd7_RST_2", ms, 0x10)
}

func instr_0xd8_RC(ms *machineState) {
	// 1		if CY, RET
	RET("0xd8_RC", ms, "CY", ms.flagCY)
}

func instr_0xda_JC_adr(ms *machineState) {
	// 3		if CY, PC<-adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xda_JC_adr 0x%04x, C=%t\n", ms.pc, adr, ms.flagCY)
	if ms.flagCY {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xdb_IN(ms *machineState) {
	// D8	2		special
	panic("Unimplemented")
}

func instr_0xdc_CC_adr(ms *machineState) {
	// 3		if CY, CALL adr
	CALL("0xdc_CC_adr", ms, "CY", ms.flagCY)
}

func instr_0xde_SBI_D8(ms *machineState) {
	// 2	Z, S, P, CY, AC	A <- A - data - CY
	data := ms.readMem(ms.pc+1, 1)[0]
	regA := ms.regA
	var carry uint8
	if ms.flagCY {
		carry = 1
	} else {
		carry = 0
	}
	ms.regA = regA - data - carry
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)+uint(data)+uint(carry) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: 0xde_SBI_D8 regA[0x%02x]=regA[0x%02x]+0x%02x+0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, regA, data, carry, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xdf_RST_3(ms *machineState) {
	// 1		CALL $18
	RST("instr_0xdf_RST_3", ms, 0x18)
}

func instr_0xe0_RPO(ms *machineState) {
	// 1		if PO, RET
	RET("0xe0_RPO", ms, "P", !ms.flagP)
}

func instr_0xe1_POP_H(ms *machineState) {
	// 1		L <- (sp); H <- (sp+1); sp <- sp+2
	POP("0xe1_POP_H", ms, &ms.regL, &ms.regH)
}

func instr_0xe2_JPO_adr(ms *machineState) {
	// 3		if PO, PC <- adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xe2_JPO_adr 0x%04x, P=%t\n", ms.pc, adr, ms.flagP)
	if !ms.flagP {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xe3_XTHL(ms *machineState) {
	// 1		L <-> (SP); H <-> (SP+1)
	panic("Unimplemented")
}

func instr_0xe4_CPO_adr(ms *machineState) {
	// 3		if PO, CALL adr
	CALL("0xe4_CPO_adr", ms, "P", !ms.flagP)

}

func instr_0xe5_PUSH_H(ms *machineState) {
	// 1		(sp-2)<-L; (sp-1)<-H; sp <- sp - 2
	PUSH("0xe5_PUSH_H", ms, &ms.regL, &ms.regH)
}

func instr_0xe6_ANI_D8(ms *machineState) {
	// 2	Z, S, P, CY, AC	A <- A & data
	data := ms.readMem(ms.pc+1, 1)[0]
	regA := ms.regA
	ms.regA = regA & data
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(false)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: 0xe6_ANI_D8 regA[0x%02x]=regA[0x%02x]&0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, regA, data, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xe7_RST_4(ms *machineState) {
	// 1		CALL $20
	RST("instr_0xe7_RST_4", ms, 0x20)
}

func instr_0xe8_RPE(ms *machineState) {
	// 1		if PE, RET
	RET("0xe8_RPE", ms, "P", ms.flagP)
}

func instr_0xe9_PCHL(ms *machineState) {
	// 1		PC.hi <- H; PC.lo <- L
	panic("Unimplemented")
}

func instr_0xea_JPE_adr(ms *machineState) {
	// 3		if PE, PC <- adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xea_JPE_adr 0x%04x, P=%t\n", ms.pc, adr, ms.flagP)
	if ms.flagP {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xeb_XCHG(ms *machineState) {
	// 1		H <-> D; L <-> E
	regD := ms.regD
	regE := ms.regE
	regH := ms.regH
	regL := ms.regL
	ms.regH = regD
	ms.regD = regH
	ms.regL = regE
	ms.regE = regL
	Trace.Printf("0x%04x: 0xeb_XCHG regH[0x%02x]<->regD[0x%02x], regH[0x%02x]<->regD[0x%02x]\n",
		ms.pc, ms.regH, ms.regD, ms.regL, ms.regE)
	ms.pc += 1
}

func instr_0xec_CPE_adr(ms *machineState) {
	// 3		if PE, CALL adr
	CALL("0xec_CPE_adr", ms, "P", ms.flagP)
}

func instr_0xee_XRI_D8(ms *machineState) {
	// 2	Z, S, P, CY, AC	A <- A ^ data
	data := ms.readMem(ms.pc+1, 1)[0]
	regA := ms.regA
	ms.regA = regA ^ data
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)^uint(data) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: 0xee_XRI_D8 regA[0x%02x]=regA[0x%02x]^0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, regA, data, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xef_RST_5(ms *machineState) {
	// 1		CALL $28
	RST("instr_0xef_RST_5", ms, 0x28)
}

func instr_0xf0_RP(ms *machineState) {
	// 1		if P, RET
	RET("0xf0_RP", ms, "S", !ms.flagS)
}

func instr_0xf1_POP(ms *machineState) {
	// PSW	1		flags <- (sp); A <- (sp+1); sp <- sp+2
	panic("Unimplemented")
}

func instr_0xf2_JP_adr(ms *machineState) {
	// 3		if P=1 PC <- adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xf2_JP_adr 0x%04x, S=%t\n", ms.pc, adr, ms.flagS)
	if !ms.flagS {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xf3_DI(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0xf4_CP_adr(ms *machineState) {
	// 3		if P, CALL adr
	CALL("0xf4_CP_adr", ms, "S", !ms.flagS)
}

func instr_0xf5_PUSH(ms *machineState) {
	// PSW        1               (sp-2)<-flags; (sp-1)<-A; sp <- sp - 2
	panic("Unimplemented")
}

func instr_0xf6_ORI_D8(ms *machineState) {
	// 2	Z, S, P, CY, AC	A <- A | data
	data := ms.readMem(ms.pc+1, 1)[0]
	regA := ms.regA
	ms.regA = regA | data
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setCY(uint(regA)|uint(data) > math.MaxUint8)
	ms.setAC(ms.regA)
	Trace.Printf("0x%04x: 0xf6_ORI_D8 regA[0x%02x]=regA[0x%02x]|0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, regA, data, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xf7_RST_6(ms *machineState) {
	// 1		CALL $30
	RST("instr_0xf7_RST_6", ms, 0x30)
}

func instr_0xf8_RM(ms *machineState) {
	// 1		if M, RET
	RET("0xf8_RM", ms, "S", ms.flagS)
}

func instr_0xf9_SPHL(ms *machineState) {
	// 1		SP=HL
	panic("Unimplemented")
}

func instr_0xfa_JM_adr(ms *machineState) {
	// 3		if M, PC <- adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	Trace.Printf("0x%04x: 0xfa_JM_adr 0x%04x, S=%t\n", ms.pc, adr, ms.flagS)
	if ms.flagS {
		ms.pc = adr
	} else {
		ms.pc += 3
	}
}

func instr_0xfb_EI(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0xfc_CM_adr(ms *machineState) {
	// 3		if M, CALL adr
	CALL("0xfc_CM_adr", ms, "S", ms.flagS)
}

func instr_0xfe_CPI_D8(ms *machineState) {
	// 2	Z, S, P, CY, AC	A - data
	data := ms.readMem(ms.pc+1, 1)[0]
	result := ms.regA - data
	ms.setZ(result)
	ms.setS(result)
	ms.setP(result)
	ms.setCY(ms.regA < data)
	ms.setAC(result)
	Trace.Printf("0x%04x: 0xfe_CPI_D8 regA[0x%02x]-0x%02x, Z=%t, S=%t, P=%t, CY=%t, AC=%t\n",
		ms.pc, ms.regA, data, ms.flagZ, ms.flagS, ms.flagP, ms.flagCY, ms.flagAC)
	ms.pc += 2
}

func instr_0xff_RST_7(ms *machineState) {
	// 1		CALL $38
	RST("instr_0xff_RST_7", ms, 0x38)
}
