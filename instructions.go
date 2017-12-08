package main

import (
	"fmt"
)

func instr_0x00_NOP(ms *machineState) {
	// 1
	fmt.Printf("0x%04x: 0x00_NOP\n", ms.pc)
	ms.pc += 1
}

func instr_0x01_LXI_B_D16(ms *machineState) {
	// 3		B <- byte 3, C <- byte 2
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	ms.regB = byte3
	ms.regC = byte2
	fmt.Printf("0x%04x: 0x01_LXI_B_D16 0x%02x 0x%02x\n", ms.pc, ms.regB, ms.regC)
	ms.pc += 3
}

func instr_0x02_STAX(ms *machineState) {
	// B	1		(BC) <- A
	panic("Unimplemented")
}

func instr_0x03_INX(ms *machineState) {
	// B	1		BC <- BC+1
	panic("Unimplemented")
}

func instr_0x04_INR(ms *machineState) {
	// B	1	Z, S, P, AC	B <- B+1
	panic("Unimplemented")
}

func instr_0x05_DCR_B(ms *machineState) {
	// 1	Z, S, P, AC	B <- B-1
	regB := ms.regB
	ms.regB = regB - 1
	ms.setZ(ms.regB)
	ms.setS(ms.regB)
	ms.setP(ms.regB)
	ms.setAC(ms.regB)
	fmt.Printf("0x%04x: 0x05_DCR_B regB[0x%02x]=regB[0x%02x]-1, Z=%t, S=%t, P=%t, AC=%t\n",
		ms.pc, ms.regB, regB, ms.flagZ, ms.flagS, ms.flagP, ms.flagAC)
	ms.pc += 1
}

func instr_0x06_MVI_B_D8(ms *machineState) {
	// 2		B <- byte 2
	MVI("0x06_MVI_B_D8", ms, &ms.regB)
}

func instr_0x07_RLC(ms *machineState) {
	// 1       CY      A = A << 1; bit 0 = prev bit 7; CY = prev bit 7
	panic("Unimplemented")
}

func instr_0x09_DAD(ms *machineState) {
	// B	1	CY	HL = HL + BC
	panic("Unimplemented")
}

func instr_0x0a_LDAX(ms *machineState) {
	// B	1		A <- (BC)
	panic("Unimplemented")
}

func instr_0x0b_DCX(ms *machineState) {
	// B	1		BC = BC-1
	panic("Unimplemented")
}

func instr_0x0c_INR(ms *machineState) {
	// C	1	Z, S, P, AC	C <- C+1
	panic("Unimplemented")
}

func instr_0x0d_DCR_C(ms *machineState) {
	// 1	Z, S, P, AC	C <-C-1
	regC := ms.regC
	ms.regC = regC - 1
	ms.setZ(ms.regC)
	ms.setS(ms.regC)
	ms.setP(ms.regC)
	ms.setAC(ms.regC)
	fmt.Printf("0x%04x: 0x0d_DCR_C regC[0x%02x]=regC[0x%02x]-1, Z=%t, S=%t, P=%t, AC=%t\n",
		ms.pc, ms.regC, regC, ms.flagZ, ms.flagS, ms.flagP, ms.flagAC)
	ms.pc += 1
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
	fmt.Printf("0x%04x: 0x11_LXI_D_D16 0x%02x 0x%02x\n", ms.pc, ms.regD, ms.regE)
	ms.pc += 3
}

func instr_0x12_STAX(ms *machineState) {
	// D	1		(DE) <- A
	panic("Unimplemented")
}

func instr_0x13_INX(ms *machineState) {
	// D	1		DE <- DE + 1
	panic("Unimplemented")
}

func instr_0x14_INR(ms *machineState) {
	// D	1	Z, S, P, AC	D <- D+1
	panic("Unimplemented")
}

func instr_0x15_DCR_D(ms *machineState) {
	// 1	Z, S, P, AC	D <- D-1
	regD := ms.regD
	ms.regD = regD - 1
	ms.setZ(ms.regD)
	ms.setS(ms.regD)
	ms.setP(ms.regD)
	ms.setAC(ms.regD)
	fmt.Printf("0x%04x: 0x15_DCR_D regD[0x%02x]=regD[0x%02x]-1, Z=%t, S=%t, P=%t, AC=%t\n",
		ms.pc, ms.regD, regD, ms.flagZ, ms.flagS, ms.flagP, ms.flagAC)
	ms.pc += 1
}

func instr_0x16_MVI_D_D8(ms *machineState) {
	// 2		D <- byte 2
	MVI("0x16_MVI_D_D8", ms, &ms.regD)
}

func instr_0x17_RAL(ms *machineState) {
	// 1	CY	A = A << 1; bit 0 = prev CY; CY = prev bit 7
	panic("Unimplemented")
}

func instr_0x19_DAD(ms *machineState) {
	// D	1	CY	HL = HL + DE
	panic("Unimplemented")
}

func instr_0x1a_LDAX(ms *machineState) {
	// D	1		A <- (DE)
	panic("Unimplemented")
}

func instr_0x1b_DCX(ms *machineState) {
	// D	1		DE = DE-1
	panic("Unimplemented")
}

func instr_0x1c_INR(ms *machineState) {
	// E	1	Z, S, P, AC	E <-E+1
	panic("Unimplemented")
}

func instr_0x1d_DCR_E(ms *machineState) {
	// 1	Z, S, P, AC	E <- E-1
	regE := ms.regE
	ms.regE = regE - 1
	ms.setZ(ms.regE)
	ms.setS(ms.regE)
	ms.setP(ms.regE)
	ms.setAC(ms.regE)
	fmt.Printf("0x%04x: 0x1d_DCR_E regE[0x%02x]=regE[0x%02x]-1, Z=%t, S=%t, P=%t, AC=%t\n",
		ms.pc, ms.regE, regE, ms.flagZ, ms.flagS, ms.flagP, ms.flagAC)
	ms.pc += 1
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
	fmt.Printf("0x%04x: 0x20_RIM\n", ms.pc)
	ms.pc += 1
}

func instr_0x21_LXI_H_D16(ms *machineState) {
	// 3		H <- byte 3, L <- byte 2
	byte2 := ms.readMem(ms.pc+1, 1)[0]
	byte3 := ms.readMem(ms.pc+2, 1)[0]
	ms.regH = byte3
	ms.regL = byte2
	fmt.Printf("0x%04x: 0x21_LXI_H_D16 0x%02x 0x%02x\n", ms.pc, ms.regH, ms.regL)
	ms.pc += 3
}

func instr_0x22_SHLD(ms *machineState) {
	// adr	3		(adr) <-L; (adr+1)<-H
	panic("Unimplemented")
}

func instr_0x23_INX(ms *machineState) {
	// H	1		HL <- HL + 1
	panic("Unimplemented")
}

func instr_0x24_INR(ms *machineState) {
	// H	1	Z, S, P, AC	H <- H+1
	panic("Unimplemented")
}

func instr_0x25_DCR_H(ms *machineState) {
	// 1	Z, S, P, AC	H <- H-1
	regH := ms.regH
	ms.regH = regH - 1
	ms.setZ(ms.regH)
	ms.setS(ms.regH)
	ms.setP(ms.regH)
	ms.setAC(ms.regH)
	fmt.Printf("0x%04x: 0x25_DCR_H regH[0x%02x]=regH[0x%02x]-1, Z=%t, S=%t, P=%t, AC=%t\n",
		ms.pc, ms.regH, regH, ms.flagZ, ms.flagS, ms.flagP, ms.flagAC)
	ms.pc += 1
}

func instr_0x26_MVI_H_D8(ms *machineState) {
	// 2		H <- byte 2
	MVI("0x26_MVI_H_D8", ms, &ms.regH)
}

func instr_0x27_DAA(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0x29_DAD(ms *machineState) {
	// H	1	CY	HL = HL + HI
	panic("Unimplemented")
}

func instr_0x2a_LHLD(ms *machineState) {
	// adr	3		L <- (adr); H<-(adr+1)
	panic("Unimplemented")
}

func instr_0x2b_DCX(ms *machineState) {
	// H	1		HL = HL-1
	panic("Unimplemented")
}

func instr_0x2c_INR(ms *machineState) {
	// L	1	Z, S, P, AC	L <- L+1
	panic("Unimplemented")
}

func instr_0x2d_DCR_L(ms *machineState) {
	// 1	Z, S, P, AC	L <- L-1
	regL := ms.regL
	ms.regL = regL - 1
	ms.setZ(ms.regL)
	ms.setS(ms.regL)
	ms.setP(ms.regL)
	ms.setAC(ms.regL)
	fmt.Printf("0x%04x: 0x2d_DCR_L regL[0x%02x]=regL[0x%02x]-1, Z=%t, S=%t, P=%t, AC=%t\n",
		ms.pc, ms.regL, regL, ms.flagZ, ms.flagS, ms.flagP, ms.flagAC)
	ms.pc += 1
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
	fmt.Printf("0x%04x: 0x31_LXI_SP_D16 0x%04x\n", ms.pc, sp)
	ms.sp = sp
	ms.pc += 3
}

func instr_0x32_STA(ms *machineState) {
	// adr	3		(adr) <- A
	panic("Unimplemented")
}

func instr_0x33_INX(ms *machineState) {
	// SP	1		SP = SP + 1
	panic("Unimplemented")
}

func instr_0x34_INR(ms *machineState) {
	// M	1	Z, S, P, AC	(HL) <- (HL)+1
	panic("Unimplemented")
}

func instr_0x35_DCR_M(ms *machineState) {
	// 1	Z, S, P, AC	(HL) <- (HL)-1
	panic("Unimplemented")
}

func instr_0x36_MVI(ms *machineState) {
	// M,D8	2		(HL) <- byte 2
	panic("Unimplemented")
}

func instr_0x37_STC(ms *machineState) {
	// 1	CY	CY = 1
	panic("Unimplemented")
}

func instr_0x39_DAD(ms *machineState) {
	// SP	1	CY	HL = HL + SP
	panic("Unimplemented")
}

func instr_0x3a_LDA(ms *machineState) {
	// adr	3		A <- (adr)
	panic("Unimplemented")
}

func instr_0x3b_DCX(ms *machineState) {
	// SP	1		SP = SP-1
	panic("Unimplemented")
}

func instr_0x3c_INR(ms *machineState) {
	// A	1	Z, S, P, AC	A <- A+1
	panic("Unimplemented")
}

func instr_0x3d_DCR_A(ms *machineState) {
	// 1	Z, S, P, AC	A <- A-1
	regA := ms.regA
	ms.regA = regA - 1
	ms.setZ(ms.regA)
	ms.setS(ms.regA)
	ms.setP(ms.regA)
	ms.setAC(ms.regA)
	fmt.Printf("0x%04x: 0x3d_DCR_A regA[0x%02x]=regA[0x%02x]-1, Z=%t, S=%t, P=%t, AC=%t\n",
		ms.pc, ms.regA, regA, ms.flagZ, ms.flagS, ms.flagP, ms.flagAC)
	ms.pc += 1
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
	ms.regB = ms.regB
	fmt.Printf("0x%04x: instr_0x40_MOV_B_B regB[0x%02x]=regB[0x%02x]\n",
		ms.pc, ms.regB, ms.regB)
	ms.pc += 1
}

func instr_0x41_MOV_B_C(ms *machineState) {
	// 1		B <- C
	ms.regB = ms.regC
	fmt.Printf("0x%04x: instr_0x41_MOV_B_C regB[0x%02x]=regC[0x%02x]\n",
		ms.pc, ms.regB, ms.regC)
	ms.pc += 1
}

func instr_0x42_MOV_B_D(ms *machineState) {
	// 1		B <- D
	ms.regB = ms.regD
	fmt.Printf("0x%04x: instr_0x42_MOV_B_D regB[0x%02x]=regD[0x%02x]\n",
		ms.pc, ms.regB, ms.regD)
	ms.pc += 1
}

func instr_0x43_MOV_B_E(ms *machineState) {
	// 1		B <- E
	ms.regB = ms.regE
	fmt.Printf("0x%04x: instr_0x43_MOV_B_E regB[0x%02x]=regE[0x%02x]\n",
		ms.pc, ms.regB, ms.regE)
	ms.pc += 1
}

func instr_0x44_MOV_B_H(ms *machineState) {
	// 1		B <- H
	ms.regB = ms.regH
	fmt.Printf("0x%04x: instr_0x44_MOV_B_H regB[0x%02x]=regH[0x%02x]\n",
		ms.pc, ms.regB, ms.regH)
	ms.pc += 1
}

func instr_0x45_MOV_B_L(ms *machineState) {
	// 1		B <- L
	ms.regB = ms.regL
	fmt.Printf("0x%04x: instr_0x45_MOV_B_L regB[0x%02x]=regL[0x%02x]\n",
		ms.pc, ms.regB, ms.regL)
	ms.pc += 1
}

func instr_0x46_MOV(ms *machineState) {
	// B,M	1		B <- (HL)
	panic("Unimplemented")
}

func instr_0x47_MOV(ms *machineState) {
	// B,A	1		B <- A
	panic("Unimplemented")
}

func instr_0x48_MOV(ms *machineState) {
	// C,B	1		C <- B
	panic("Unimplemented")
}

func instr_0x49_MOV(ms *machineState) {
	// C,C	1		C <- C
	panic("Unimplemented")
}

func instr_0x4a_MOV(ms *machineState) {
	// C,D	1		C <- D
	panic("Unimplemented")
}

func instr_0x4b_MOV(ms *machineState) {
	// C,E	1		C <- E
	panic("Unimplemented")
}

func instr_0x4c_MOV(ms *machineState) {
	// C,H	1		C <- H
	panic("Unimplemented")
}

func instr_0x4d_MOV(ms *machineState) {
	// C,L	1		C <- L
	panic("Unimplemented")
}

func instr_0x4e_MOV(ms *machineState) {
	// C,M	1		C <- (HL)
	panic("Unimplemented")
}

func instr_0x4f_MOV(ms *machineState) {
	// C,A	1		C <- A
	panic("Unimplemented")
}

func instr_0x50_MOV(ms *machineState) {
	// D,B	1		D <- B
	panic("Unimplemented")
}

func instr_0x51_MOV(ms *machineState) {
	// D,C	1		D <- C
	panic("Unimplemented")
}

func instr_0x52_MOV(ms *machineState) {
	// D,D	1		D <- D
	panic("Unimplemented")
}

func instr_0x53_MOV(ms *machineState) {
	// D,E	1		D <- E
	panic("Unimplemented")
}

func instr_0x54_MOV(ms *machineState) {
	// D,H	1		D <- H
	panic("Unimplemented")
}

func instr_0x55_MOV(ms *machineState) {
	// D,L	1		D <- L
	panic("Unimplemented")
}

func instr_0x56_MOV(ms *machineState) {
	// D,M	1		D <- (HL)
	panic("Unimplemented")
}

func instr_0x57_MOV(ms *machineState) {
	// D,A	1		D <- A
	panic("Unimplemented")
}

func instr_0x58_MOV(ms *machineState) {
	// E,B	1		E <- B
	panic("Unimplemented")
}

func instr_0x59_MOV(ms *machineState) {
	// E,C	1		E <- C
	panic("Unimplemented")
}

func instr_0x5a_MOV(ms *machineState) {
	// E,D	1		E <- D
	panic("Unimplemented")
}

func instr_0x5b_MOV(ms *machineState) {
	// E,E	1		E <- E
	panic("Unimplemented")
}

func instr_0x5c_MOV(ms *machineState) {
	// E,H	1		E <- H
	panic("Unimplemented")
}

func instr_0x5d_MOV(ms *machineState) {
	// E,L	1		E <- L
	panic("Unimplemented")
}

func instr_0x5e_MOV(ms *machineState) {
	// E,M	1		E <- (HL)
	panic("Unimplemented")
}

func instr_0x5f_MOV(ms *machineState) {
	// E,A	1		E <- A
	panic("Unimplemented")
}

func instr_0x60_MOV(ms *machineState) {
	// H,B	1		H <- B
	panic("Unimplemented")
}

func instr_0x61_MOV(ms *machineState) {
	// H,C	1		H <- C
	panic("Unimplemented")
}

func instr_0x62_MOV(ms *machineState) {
	// H,D	1		H <- D
	panic("Unimplemented")
}

func instr_0x63_MOV(ms *machineState) {
	// H,E	1		H <- E
	panic("Unimplemented")
}

func instr_0x64_MOV(ms *machineState) {
	// H,H	1		H <- H
	panic("Unimplemented")
}

func instr_0x65_MOV(ms *machineState) {
	// H,L	1		H <- L
	panic("Unimplemented")
}

func instr_0x66_MOV(ms *machineState) {
	// H,M	1		H <- (HL)
	panic("Unimplemented")
}

func instr_0x67_MOV(ms *machineState) {
	// H,A	1		H <- A
	panic("Unimplemented")
}

func instr_0x68_MOV(ms *machineState) {
	// L,B	1		L <- B
	panic("Unimplemented")
}

func instr_0x69_MOV(ms *machineState) {
	// L,C	1		L <- C
	panic("Unimplemented")
}

func instr_0x6a_MOV(ms *machineState) {
	// L,D	1		L <- D
	panic("Unimplemented")
}

func instr_0x6b_MOV(ms *machineState) {
	// L,E	1		L <- E
	panic("Unimplemented")
}

func instr_0x6c_MOV(ms *machineState) {
	// L,H	1		L <- H
	panic("Unimplemented")
}

func instr_0x6d_MOV(ms *machineState) {
	// L,L	1		L <- L
	panic("Unimplemented")
}

func instr_0x6e_MOV(ms *machineState) {
	// L,M	1		L <- (HL)
	panic("Unimplemented")
}

func instr_0x6f_MOV(ms *machineState) {
	// L,A	1		L <- A
	panic("Unimplemented")
}

func instr_0x70_MOV(ms *machineState) {
	// M,B	1		(HL) <- B
	panic("Unimplemented")
}

func instr_0x71_MOV(ms *machineState) {
	// M,C	1		(HL) <- C
	panic("Unimplemented")
}

func instr_0x72_MOV(ms *machineState) {
	// M,D	1		(HL) <- D
	panic("Unimplemented")
}

func instr_0x73_MOV(ms *machineState) {
	// M,E	1		(HL) <- E
	panic("Unimplemented")
}

func instr_0x74_MOV(ms *machineState) {
	// M,H	1		(HL) <- H
	panic("Unimplemented")
}

func instr_0x75_MOV(ms *machineState) {
	// M,L	1		(HL) <- L
	panic("Unimplemented")
}

func instr_0x76_HLT(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0x77_MOV(ms *machineState) {
	// M,A	1		(HL) <- A
	panic("Unimplemented")
}

func instr_0x78_MOV(ms *machineState) {
	// A,B	1		A <- B
	panic("Unimplemented")
}

func instr_0x79_MOV(ms *machineState) {
	// A,C	1		A <- C
	panic("Unimplemented")
}

func instr_0x7a_MOV(ms *machineState) {
	// A,D	1		A <- D
	panic("Unimplemented")
}

func instr_0x7b_MOV(ms *machineState) {
	// A,E	1		A <- E
	panic("Unimplemented")
}

func instr_0x7c_MOV(ms *machineState) {
	// A,H	1		A <- H
	panic("Unimplemented")
}

func instr_0x7d_MOV(ms *machineState) {
	// A,L	1		A <- L
	panic("Unimplemented")
}

func instr_0x7e_MOV(ms *machineState) {
	// A,M	1		A <- (HL)
	panic("Unimplemented")
}

func instr_0x7f_MOV(ms *machineState) {
	// A,A	1		A <- A
	panic("Unimplemented")
}

func instr_0x80_ADD(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A <- A + B
	panic("Unimplemented")
}

func instr_0x81_ADD(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A <- A + C
	panic("Unimplemented")
}

func instr_0x82_ADD(ms *machineState) {
	// D	1	Z, S, P, CY, AC	A <- A + D
	panic("Unimplemented")
}

func instr_0x83_ADD(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A <- A + E
	panic("Unimplemented")
}

func instr_0x84_ADD(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A <- A + H
	panic("Unimplemented")
}

func instr_0x85_ADD(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A <- A + L
	panic("Unimplemented")
}

func instr_0x86_ADD(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A <- A + (HL)
	panic("Unimplemented")
}

func instr_0x87_ADD(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A <- A + A
	panic("Unimplemented")
}

func instr_0x88_ADC(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A <- A + B + CY
	panic("Unimplemented")
}

func instr_0x89_ADC(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A <- A + C + CY
	panic("Unimplemented")
}

func instr_0x8a_ADC(ms *machineState) {
	// D	1	Z, S, P, CY, AC	A <- A + D + CY
	panic("Unimplemented")
}

func instr_0x8b_ADC(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A <- A + E + CY
	panic("Unimplemented")
}

func instr_0x8c_ADC(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A <- A + H + CY
	panic("Unimplemented")
}

func instr_0x8d_ADC(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A <- A + L + CY
	panic("Unimplemented")
}

func instr_0x8e_ADC(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A <- A + (HL) + CY
	panic("Unimplemented")
}

func instr_0x8f_ADC(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A <- A + A + CY
	panic("Unimplemented")
}

func instr_0x90_SUB(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A <- A - B
	panic("Unimplemented")
}

func instr_0x91_SUB(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A <- A - C
	panic("Unimplemented")
}

func instr_0x92_SUB(ms *machineState) {
	// D   1       Z, S, P, CY, AC A <- A + D
	panic("Unimplemented")
}

func instr_0x93_SUB(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A <- A - E
	panic("Unimplemented")
}

func instr_0x94_SUB(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A <- A + H
	panic("Unimplemented")
}

func instr_0x95_SUB(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A <- A - L
	panic("Unimplemented")
}

func instr_0x96_SUB(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A <- A + (HL)
	panic("Unimplemented")
}

func instr_0x97_SUB(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A <- A - A
	panic("Unimplemented")
}

func instr_0x98_SBB(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A <- A - B - CY
	panic("Unimplemented")
}

func instr_0x99_SBB(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A <- A - C - CY
	panic("Unimplemented")
}

func instr_0x9a_SBB(ms *machineState) {
	// D	1	Z, S, P, CY, AC	A <- A - D - CY
	panic("Unimplemented")
}

func instr_0x9b_SBB(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A <- A - E - CY
	panic("Unimplemented")
}

func instr_0x9c_SBB(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A <- A - H - CY
	panic("Unimplemented")
}

func instr_0x9d_SBB(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A <- A - L - CY
	panic("Unimplemented")
}

func instr_0x9e_SBB(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A <- A - (HL) - CY
	panic("Unimplemented")
}

func instr_0x9f_SBB(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A <- A - A - CY
	panic("Unimplemented")
}

func instr_0xa0_ANA(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A <- A & B
	panic("Unimplemented")
}

func instr_0xa1_ANA(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A <- A & C
	panic("Unimplemented")
}

func instr_0xa2_ANA(ms *machineState) {
	// D	1	Z, S, P, CY, AC	A <- A & D
	panic("Unimplemented")
}

func instr_0xa3_ANA(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A <- A & E
	panic("Unimplemented")
}

func instr_0xa4_ANA(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A <- A & H
	panic("Unimplemented")
}

func instr_0xa5_ANA(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A <- A & L
	panic("Unimplemented")
}

func instr_0xa6_ANA(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A <- A & (HL)
	panic("Unimplemented")
}

func instr_0xa7_ANA(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A <- A & A
	panic("Unimplemented")
}

func instr_0xa8_XRA(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A <- A ^ B
	panic("Unimplemented")
}

func instr_0xa9_XRA(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A <- A ^ C
	panic("Unimplemented")
}

func instr_0xaa_XRA(ms *machineState) {
	// D	1	Z, S, P, CY, AC	A <- A ^ D
	panic("Unimplemented")
}

func instr_0xab_XRA(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A <- A ^ E
	panic("Unimplemented")
}

func instr_0xac_XRA(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A <- A ^ H
	panic("Unimplemented")
}

func instr_0xad_XRA(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A <- A ^ L
	panic("Unimplemented")
}

func instr_0xae_XRA(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A <- A ^ (HL)
	panic("Unimplemented")
}

func instr_0xaf_XRA(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A <- A ^ A
	panic("Unimplemented")
}

func instr_0xb0_ORA(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A <- A | B
	panic("Unimplemented")
}

func instr_0xb1_ORA(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A <- A | C
	panic("Unimplemented")
}

func instr_0xb2_ORA(ms *machineState) {
	// D	1	Z, S, P, CY, AC	A <- A | D
	panic("Unimplemented")
}

func instr_0xb3_ORA(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A <- A | E
	panic("Unimplemented")
}

func instr_0xb4_ORA(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A <- A | H
	panic("Unimplemented")
}

func instr_0xb5_ORA(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A <- A | L
	panic("Unimplemented")
}

func instr_0xb6_ORA(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A <- A | (HL)
	panic("Unimplemented")
}

func instr_0xb7_ORA(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A <- A | A
	panic("Unimplemented")
}

func instr_0xb8_CMP(ms *machineState) {
	// B	1	Z, S, P, CY, AC	A - B
	panic("Unimplemented")
}

func instr_0xb9_CMP(ms *machineState) {
	// C	1	Z, S, P, CY, AC	A - C
	panic("Unimplemented")
}

func instr_0xba_CMP(ms *machineState) {
	// D	1	Z, S, P, CY, AC	A - D
	panic("Unimplemented")
}

func instr_0xbb_CMP(ms *machineState) {
	// E	1	Z, S, P, CY, AC	A - E
	panic("Unimplemented")
}

func instr_0xbc_CMP(ms *machineState) {
	// H	1	Z, S, P, CY, AC	A - H
	panic("Unimplemented")
}

func instr_0xbd_CMP(ms *machineState) {
	// L	1	Z, S, P, CY, AC	A - L
	panic("Unimplemented")
}

func instr_0xbe_CMP(ms *machineState) {
	// M	1	Z, S, P, CY, AC	A - (HL)
	panic("Unimplemented")
}

func instr_0xbf_CMP(ms *machineState) {
	// A	1	Z, S, P, CY, AC	A - A
	panic("Unimplemented")
}

func instr_0xc0_RNZ(ms *machineState) {
	// 1		if NZ, RET
	panic("Unimplemented")
}

func instr_0xc1_POP(ms *machineState) {
	// B	1		C <- (sp); B <- (sp+1); sp <- sp+2
	panic("Unimplemented")
}

func instr_0xc2_JNZ(ms *machineState) {
	// adr	3		if NZ, PC <- adr
	panic("Unimplemented")
}

func instr_0xc3_JMP_adr(ms *machineState) {
	// 3		PC <= adr
	byte1 := ms.readMem(ms.pc+1, 1)[0]
	byte2 := ms.readMem(ms.pc+2, 1)[0]
	var adr uint16 = (uint16(byte2) << 8) | uint16(byte1)
	fmt.Printf("0x%04x: 0xc3_JMP_adr 0x%04x\n", ms.pc, adr)
	ms.pc = adr
}

func instr_0xc4_CNZ(ms *machineState) {
	// adr	3		if NZ, CALL adr
	panic("Unimplemented")
}

func instr_0xc5_PUSH(ms *machineState) {
	// B	1		(sp-2)<-C; (sp-1)<-B; sp <- sp - 2
	panic("Unimplemented")
}

func instr_0xc6_ADI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A + byte
	panic("Unimplemented")
}

func instr_0xc7_RST(ms *machineState) {
	// 0	1		CALL $0
	panic("Unimplemented")
}

func instr_0xc8_RZ(ms *machineState) {
	// 1		if Z, RET
	panic("Unimplemented")
}

func instr_0xc9_RET(ms *machineState) {
	// 1		PC.lo <- (sp); PC.hi<-(sp+1); SP <- SP+2
	bytes := ms.readMem(ms.sp, 3)
	pcLo := bytes[0]
	pcHi := bytes[1]
	newSp := uint16(bytes[2])
	pc := (uint16(pcHi) << 8) | uint16(pcLo)
	fmt.Printf("0x%02x: 0xc9_RET pc=0x%04x, sp=0x%04x\n", ms.pc, pc, newSp)
	ms.pc = (uint16(pcHi) << 8) | uint16(pcLo)
	ms.sp = newSp
}

func instr_0xca_JZ(ms *machineState) {
	// adr	3		if Z, PC <- adr
	panic("Unimplemented")
}

func instr_0xcc_CZ(ms *machineState) {
	// adr	3		if Z, CALL adr
	panic("Unimplemented")
}

func instr_0xcd_CALL(ms *machineState) {
	// adr	3		(SP-1)<-PC.hi;(SP-2)<-PC.lo;SP<-SP+2;PC=adr
	panic("Unimplemented")
}

func instr_0xce_ACI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A + data + CY
	panic("Unimplemented")
}

func instr_0xcf_RST(ms *machineState) {
	// 1	1		CALL $8
	panic("Unimplemented")
}

func instr_0xd0_RNC(ms *machineState) {
	// 1		if NCY, RET
	panic("Unimplemented")
}

func instr_0xd1_POP(ms *machineState) {
	// D	1		E <- (sp); D <- (sp+1); sp <- sp+2
	panic("Unimplemented")
}

func instr_0xd2_JNC(ms *machineState) {
	// adr	3		if NCY, PC<-adr
	panic("Unimplemented")
}

func instr_0xd3_OUT(ms *machineState) {
	// D8	2		special
	panic("Unimplemented")
}

func instr_0xd4_CNC(ms *machineState) {
	// adr	3		if NCY, CALL adr
	panic("Unimplemented")
}

func instr_0xd5_PUSH(ms *machineState) {
	// D	1		(sp-2)<-E; (sp-1)<-D; sp <- sp - 2
	panic("Unimplemented")
}

func instr_0xd6_SUI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A - data
	panic("Unimplemented")
}

func instr_0xd7_RST(ms *machineState) {
	// 2	1		CALL $10
	panic("Unimplemented")
}

func instr_0xd8_RC(ms *machineState) {
	// 1		if CY, RET
	panic("Unimplemented")
}

func instr_0xda_JC(ms *machineState) {
	// adr	3		if CY, PC<-adr
	panic("Unimplemented")
}

func instr_0xdb_IN(ms *machineState) {
	// D8	2		special
	panic("Unimplemented")
}

func instr_0xdc_CC(ms *machineState) {
	// adr	3		if CY, CALL adr
	panic("Unimplemented")
}

func instr_0xde_SBI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A - data - CY
	panic("Unimplemented")
}

func instr_0xdf_RST(ms *machineState) {
	// 3	1		CALL $18
	panic("Unimplemented")
}

func instr_0xe0_RPO(ms *machineState) {
	// 1		if PO, RET
	panic("Unimplemented")
}

func instr_0xe1_POP(ms *machineState) {
	// H	1		L <- (sp); H <- (sp+1); sp <- sp+2
	panic("Unimplemented")
}

func instr_0xe2_JPO(ms *machineState) {
	// adr	3		if PO, PC <- adr
	panic("Unimplemented")
}

func instr_0xe3_XTHL(ms *machineState) {
	// 1		L <-> (SP); H <-> (SP+1)
	panic("Unimplemented")
}

func instr_0xe4_CPO(ms *machineState) {
	// adr	3		if PO, CALL adr
	panic("Unimplemented")
}

func instr_0xe5_PUSH(ms *machineState) {
	// H	1		(sp-2)<-L; (sp-1)<-H; sp <- sp - 2
	panic("Unimplemented")
}

func instr_0xe6_ANI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A & data
	panic("Unimplemented")
}

func instr_0xe7_RST(ms *machineState) {
	// 4	1		CALL $20
	panic("Unimplemented")
}

func instr_0xe8_RPE(ms *machineState) {
	// 1		if PE, RET
	panic("Unimplemented")
}

func instr_0xe9_PCHL(ms *machineState) {
	// 1		PC.hi <- H; PC.lo <- L
	panic("Unimplemented")
}

func instr_0xea_JPE(ms *machineState) {
	// adr	3		if PE, PC <- adr
	panic("Unimplemented")
}

func instr_0xeb_XCHG(ms *machineState) {
	// 1		H <-> D; L <-> E
	panic("Unimplemented")
}

func instr_0xec_CPE(ms *machineState) {
	// adr	3		if PE, CALL adr
	panic("Unimplemented")
}

func instr_0xee_XRI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A ^ data
	panic("Unimplemented")
}

func instr_0xef_RST(ms *machineState) {
	// 5	1		CALL $28
	panic("Unimplemented")
}

func instr_0xf0_RP(ms *machineState) {
	// 1		if P, RET
	panic("Unimplemented")
}

func instr_0xf1_POP(ms *machineState) {
	// PSW	1		flags <- (sp); A <- (sp+1); sp <- sp+2
	panic("Unimplemented")
}

func instr_0xf2_JP(ms *machineState) {
	// adr	3		if P=1 PC <- adr
	panic("Unimplemented")
}

func instr_0xf3_DI(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0xf4_CP(ms *machineState) {
	// adr	3		if P, PC <- adr
	panic("Unimplemented")
}

func instr_0xf5_PUSH(ms *machineState) {
	// PSW        1               (sp-2)<-flags; (sp-1)<-A; sp <- sp - 2
	panic("Unimplemented")
}

func instr_0xf6_ORI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A <- A | data
	panic("Unimplemented")
}

func instr_0xf7_RST(ms *machineState) {
	// 6	1		CALL $30
	panic("Unimplemented")
}

func instr_0xf8_RM(ms *machineState) {
	// 1		if M, RET
	panic("Unimplemented")
}

func instr_0xf9_SPHL(ms *machineState) {
	// 1		SP=HL
	panic("Unimplemented")
}

func instr_0xfa_JM(ms *machineState) {
	// adr	3		if M, PC <- adr
	panic("Unimplemented")
}

func instr_0xfb_EI(ms *machineState) {
	// 1		special
	panic("Unimplemented")
}

func instr_0xfc_CM(ms *machineState) {
	// adr	3		if M, CALL adr
	panic("Unimplemented")
}

func instr_0xfe_CPI(ms *machineState) {
	// D8	2	Z, S, P, CY, AC	A - data
	panic("Unimplemented")
}

func instr_0xff_RST(ms *machineState) {
	// 7	1		CALL $38
	panic("Unimplemented")
}
