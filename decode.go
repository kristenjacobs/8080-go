package main

import (
	"fmt"
)

func unimplemented(opcode uint8) {
	panic(fmt.Sprintf("Unimplemented opcode: 0x%02x\n", opcode))
}

func fetchAndDecode(ms *machineState) {

	// Gets the byte at the current PC
	opcode, err := ms.readMem(ms.pc, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf(">>> 0x%02x\n", opcode[0])

	// Opcode	Instruction	size	flags	function
	switch opcode[0] {
	case 0x00:
		// 0x00	NOP	1
		unimplemented(0x00)
	case 0x01:
		// 0x01	LXI B,D16	3		B <- byte 3, C <- byte 2
		unimplemented(0x01)
	case 0x02:
		// 0x02	STAX B	1		(BC) <- A
		unimplemented(0x02)
	case 0x03:
		// 0x03	INX B	1		BC <- BC+1
		unimplemented(0x03)
	case 0x04:
		// 0x04	INR B	1	Z, S, P, AC	B <- B+1
		unimplemented(0x04)
	case 0x05:
		// 0x05	DCR B	1	Z, S, P, AC	B <- B-1
		unimplemented(0x05)
	case 0x06:
		// 0x06	MVI B, D8	2		B <- byte 2
		unimplemented(0x06)
	case 0x07:
		// 0x07	RLC	1	CY	A = A << 1; bit 0 = prev bit 7; CY = prev bit 7
		unimplemented(0x07)
	case 0x08:
		// 0x08	-
		unimplemented(0x08)
	case 0x09:
		// 0x09	DAD B	1	CY	HL = HL + BC
		unimplemented(0x09)
	case 0x0a:
		// 0x0a	LDAX B	1		A <- (BC)
		unimplemented(0x0a)
	case 0x0b:
		// 0x0b	DCX B	1		BC = BC-1
		unimplemented(0x0b)
	case 0x0c:
		// 0x0c	INR C	1	Z, S, P, AC	C <- C+1
		unimplemented(0x0c)
	case 0x0d:
		// 0x0d	DCR C	1	Z, S, P, AC	C <-C-1
		unimplemented(0x0d)
	case 0x0e:
		// 0x0e	MVI C,D8	2		C <- byte 2
		unimplemented(0x0e)
	case 0x0f:
		// 0x0f	RRC	1	CY	A = A >> 1; bit 7 = prev bit 0; CY = prev bit 0
		unimplemented(0x0f)
	case 0x10:
		// 0x10	-
		unimplemented(0x10)
	case 0x11:
		// 0x11	LXI D,D16	3		D <- byte 3, E <- byte 2
		unimplemented(0x11)
	case 0x12:
		// 0x12	STAX D	1		(DE) <- A
		unimplemented(0x12)
	case 0x13:
		// 0x13	INX D	1		DE <- DE + 1
		unimplemented(0x13)
	case 0x14:
		// 0x14	INR D	1	Z, S, P, AC	D <- D+1
		unimplemented(0x14)
	case 0x15:
		// 0x15	DCR D	1	Z, S, P, AC	D <- D-1
		unimplemented(0x15)
	case 0x16:
		// 0x16	MVI D, D8	2		D <- byte 2
		unimplemented(0x16)
	case 0x17:
		// 0x17	RAL	1	CY	A = A << 1; bit 0 = prev CY; CY = prev bit 7
		unimplemented(0x17)
	case 0x18:
		// 0x18	-
		unimplemented(0x18)
	case 0x19:
		// 0x19	DAD D	1	CY	HL = HL + DE
		unimplemented(0x19)
	case 0x1a:
		// 0x1a	LDAX D	1		A <- (DE)
		unimplemented(0x1a)
	case 0x1b:
		// 0x1b	DCX D	1		DE = DE-1
		unimplemented(0x1b)
	case 0x1c:
		// 0x1c	INR E	1	Z, S, P, AC	E <-E+1
		unimplemented(0x1c)
	case 0x1d:
		// 0x1d	DCR E	1	Z, S, P, AC	E <- E-1
		unimplemented(0x1d)
	case 0x1e:
		// 0x1e	MVI E,D8	2		E <- byte 2
		unimplemented(0x1e)
	case 0x1f:
		// 0x1f	RAR	1	CY	A = A >> 1; bit 7 = prev bit 7; CY = prev bit 0
		unimplemented(0x1f)
	case 0x20:
		// 0x20	RIM	1		special
		unimplemented(0x20)
	case 0x21:
		// 0x21	LXI H,D16	3		H <- byte 3, L <- byte 2
		unimplemented(0x21)
	case 0x22:
		// 0x22	SHLD adr	3		(adr) <-L; (adr+1)<-H
		unimplemented(0x22)
	case 0x23:
		// 0x23	INX H	1		HL <- HL + 1
		unimplemented(0x23)
	case 0x24:
		// 0x24	INR H	1	Z, S, P, AC	H <- H+1
		unimplemented(0x24)
	case 0x25:
		// 0x25	DCR H	1	Z, S, P, AC	H <- H-1
		unimplemented(0x25)
	case 0x26:
		// 0x26	MVI H,D8	2		H <- byte 2
		unimplemented(0x26)
	case 0x27:
		// 0x27	DAA	1		special
		unimplemented(0x27)
	case 0x28:
		// 0x28	-
		unimplemented(0x28)
	case 0x29:
		// 0x29	DAD H	1	CY	HL = HL + HI
		unimplemented(0x29)
	case 0x2a:
		// 0x2a	LHLD adr	3		L <- (adr); H<-(adr+1)
		unimplemented(0x2a)
	case 0x2b:
		// 0x2b	DCX H	1		HL = HL-1
		unimplemented(0x2b)
	case 0x2c:
		// 0x2c	INR L	1	Z, S, P, AC	L <- L+1
		unimplemented(0x2c)
	case 0x2d:
		// 0x2d	DCR L	1	Z, S, P, AC	L <- L-1
		unimplemented(0x2d)
	case 0x2e:
		// 0x2e	MVI L, D8	2		L <- byte 2
		unimplemented(0x2e)
	case 0x2f:
		// 0x2f	CMA	1		A <- !A
		unimplemented(0x2f)
	case 0x30:
		// 0x30	SIM	1		special
		unimplemented(0x30)
	case 0x31:
		// 0x31	LXI SP, D16	3		SP.hi <- byte 3, SP.lo <- byte 2
		unimplemented(0x31)
	case 0x32:
		// 0x32	STA adr	3		(adr) <- A
		unimplemented(0x32)
	case 0x33:
		// 0x33	INX SP	1		SP = SP + 1
		unimplemented(0x33)
	case 0x34:
		// 0x34	INR M	1	Z, S, P, AC	(HL) <- (HL)+1
		unimplemented(0x34)
	case 0x35:
		// 0x35	DCR M	1	Z, S, P, AC	(HL) <- (HL)-1
		unimplemented(0x35)
	case 0x36:
		// 0x36	MVI M,D8	2		(HL) <- byte 2
		unimplemented(0x36)
	case 0x37:
		// 0x37	STC	1	CY	CY = 1
		unimplemented(0x37)
	case 0x38:
		// 0x38	-
		unimplemented(0x38)
	case 0x39:
		// 0x39	DAD SP	1	CY	HL = HL + SP
		unimplemented(0x39)
	case 0x3a:
		// 0x3a	LDA adr	3		A <- (adr)
		unimplemented(0x3a)
	case 0x3b:
		// 0x3b	DCX SP	1		SP = SP-1
		unimplemented(0x3b)
	case 0x3c:
		// 0x3c	INR A	1	Z, S, P, AC	A <- A+1
		unimplemented(0x3c)
	case 0x3d:
		// 0x3d	DCR A	1	Z, S, P, AC	A <- A-1
		unimplemented(0x3d)
	case 0x3e:
		// 0x3e	MVI A,D8	2		A <- byte 2
		unimplemented(0x3e)
	case 0x3f:
		// 0x3f	CMC	1	CY	CY=!CY
		unimplemented(0x3f)
	case 0x40:
		// 0x40	MOV B,B	1		B <- B
		unimplemented(0x40)
	case 0x41:
		// 0x41	MOV B,C	1		B <- C
		unimplemented(0x41)
	case 0x42:
		// 0x42	MOV B,D	1		B <- D
		unimplemented(0x42)
	case 0x43:
		// 0x43	MOV B,E	1		B <- E
		unimplemented(0x43)
	case 0x44:
		// 0x44	MOV B,H	1		B <- H
		unimplemented(0x44)
	case 0x45:
		// 0x45	MOV B,L	1		B <- L
		unimplemented(0x45)
	case 0x46:
		// 0x46	MOV B,M	1		B <- (HL)
		unimplemented(0x46)
	case 0x47:
		// 0x47	MOV B,A	1		B <- A
		unimplemented(0x47)
	case 0x48:
		// 0x48	MOV C,B	1		C <- B
		unimplemented(0x48)
	case 0x49:
		// 0x49	MOV C,C	1		C <- C
		unimplemented(0x49)
	case 0x4a:
		// 0x4a	MOV C,D	1		C <- D
		unimplemented(0x4a)
	case 0x4b:
		// 0x4b	MOV C,E	1		C <- E
		unimplemented(0x4b)
	case 0x4c:
		// 0x4c	MOV C,H	1		C <- H
		unimplemented(0x4c)
	case 0x4d:
		// 0x4d	MOV C,L	1		C <- L
		unimplemented(0x4d)
	case 0x4e:
		// 0x4e	MOV C,M	1		C <- (HL)
		unimplemented(0x4e)
	case 0x4f:
		// 0x4f	MOV C,A	1		C <- A
		unimplemented(0x4f)
	case 0x50:
		// 0x50	MOV D,B	1		D <- B
		unimplemented(0x50)
	case 0x51:
		// 0x51	MOV D,C	1		D <- C
		unimplemented(0x51)
	case 0x52:
		// 0x52	MOV D,D	1		D <- D
		unimplemented(0x52)
	case 0x53:
		// 0x53	MOV D,E	1		D <- E
		unimplemented(0x53)
	case 0x54:
		// 0x54	MOV D,H	1		D <- H
		unimplemented(0x54)
	case 0x55:
		// 0x55	MOV D,L	1		D <- L
		unimplemented(0x55)
	case 0x56:
		// 0x56	MOV D,M	1		D <- (HL)
		unimplemented(0x56)
	case 0x57:
		// 0x57	MOV D,A	1		D <- A
		unimplemented(0x57)
	case 0x58:
		// 0x58	MOV E,B	1		E <- B
		unimplemented(0x58)
	case 0x59:
		// 0x59	MOV E,C	1		E <- C
		unimplemented(0x59)
	case 0x5a:
		// 0x5a	MOV E,D	1		E <- D
		unimplemented(0x5a)
	case 0x5b:
		// 0x5b	MOV E,E	1		E <- E
		unimplemented(0x5b)
	case 0x5c:
		// 0x5c	MOV E,H	1		E <- H
		unimplemented(0x5c)
	case 0x5d:
		// 0x5d	MOV E,L	1		E <- L
		unimplemented(0x5d)
	case 0x5e:
		// 0x5e	MOV E,M	1		E <- (HL)
		unimplemented(0x5e)
	case 0x5f:
		// 0x5f	MOV E,A	1		E <- A
		unimplemented(0x5f)
	case 0x60:
		// 0x60	MOV H,B	1		H <- B
		unimplemented(0x60)
	case 0x61:
		// 0x61	MOV H,C	1		H <- C
		unimplemented(0x61)
	case 0x62:
		// 0x62	MOV H,D	1		H <- D
		unimplemented(0x62)
	case 0x63:
		// 0x63	MOV H,E	1		H <- E
		unimplemented(0x63)
	case 0x64:
		// 0x64	MOV H,H	1		H <- H
		unimplemented(0x64)
	case 0x65:
		// 0x65	MOV H,L	1		H <- L
		unimplemented(0x65)
	case 0x66:
		// 0x66	MOV H,M	1		H <- (HL)
		unimplemented(0x66)
	case 0x67:
		// 0x67	MOV H,A	1		H <- A
		unimplemented(0x67)
	case 0x68:
		// 0x68	MOV L,B	1		L <- B
		unimplemented(0x68)
	case 0x69:
		// 0x69	MOV L,C	1		L <- C
		unimplemented(0x69)
	case 0x6a:
		// 0x6a	MOV L,D	1		L <- D
		unimplemented(0x6a)
	case 0x6b:
		// 0x6b	MOV L,E	1		L <- E
		unimplemented(0x6b)
	case 0x6c:
		// 0x6c	MOV L,H	1		L <- H
		unimplemented(0x6c)
	case 0x6d:
		// 0x6d	MOV L,L	1		L <- L
		unimplemented(0x6d)
	case 0x6e:
		// 0x6e	MOV L,M	1		L <- (HL)
		unimplemented(0x6e)
	case 0x6f:
		// 0x6f	MOV L,A	1		L <- A
		unimplemented(0x6f)
	case 0x70:
		// 0x70	MOV M,B	1		(HL) <- B
		unimplemented(0x70)
	case 0x71:
		// 0x71	MOV M,C	1		(HL) <- C
		unimplemented(0x71)
	case 0x72:
		// 0x72	MOV M,D	1		(HL) <- D
		unimplemented(0x72)
	case 0x73:
		// 0x73	MOV M,E	1		(HL) <- E
		unimplemented(0x73)
	case 0x74:
		// 0x74	MOV M,H	1		(HL) <- H
		unimplemented(0x74)
	case 0x75:
		// 0x75	MOV M,L	1		(HL) <- L
		unimplemented(0x75)
	case 0x76:
		// 0x76	HLT	1		special
		unimplemented(0x76)
	case 0x77:
		// 0x77	MOV M,A	1		(HL) <- A
		unimplemented(0x77)
	case 0x78:
		// 0x78	MOV A,B	1		A <- B
		unimplemented(0x78)
	case 0x79:
		// 0x79	MOV A,C	1		A <- C
		unimplemented(0x79)
	case 0x7a:
		// 0x7a	MOV A,D	1		A <- D
		unimplemented(0x7a)
	case 0x7b:
		// 0x7b	MOV A,E	1		A <- E
		unimplemented(0x7b)
	case 0x7c:
		// 0x7c	MOV A,H	1		A <- H
		unimplemented(0x7c)
	case 0x7d:
		// 0x7d	MOV A,L	1		A <- L
		unimplemented(0x7d)
	case 0x7e:
		// 0x7e	MOV A,M	1		A <- (HL)
		unimplemented(0x7e)
	case 0x7f:
		// 0x7f	MOV A,A	1		A <- A
		unimplemented(0x7f)
	case 0x80:
		// 0x80	ADD B	1	Z, S, P, CY, AC	A <- A + B
		unimplemented(0x80)
	case 0x81:
		// 0x81	ADD C	1	Z, S, P, CY, AC	A <- A + C
		unimplemented(0x81)
	case 0x82:
		// 0x82	ADD D	1	Z, S, P, CY, AC	A <- A + D
		unimplemented(0x82)
	case 0x83:
		// 0x83	ADD E	1	Z, S, P, CY, AC	A <- A + E
		unimplemented(0x83)
	case 0x84:
		// 0x84	ADD H	1	Z, S, P, CY, AC	A <- A + H
		unimplemented(0x84)
	case 0x85:
		// 0x85	ADD L	1	Z, S, P, CY, AC	A <- A + L
		unimplemented(0x85)
	case 0x86:
		// 0x86	ADD M	1	Z, S, P, CY, AC	A <- A + (HL)
		unimplemented(0x86)
	case 0x87:
		// 0x87	ADD A	1	Z, S, P, CY, AC	A <- A + A
		unimplemented(0x87)
	case 0x88:
		// 0x88	ADC B	1	Z, S, P, CY, AC	A <- A + B + CY
		unimplemented(0x88)
	case 0x89:
		// 0x89	ADC C	1	Z, S, P, CY, AC	A <- A + C + CY
		unimplemented(0x89)
	case 0x8a:
		// 0x8a	ADC D	1	Z, S, P, CY, AC	A <- A + D + CY
		unimplemented(0x8a)
	case 0x8b:
		// 0x8b	ADC E	1	Z, S, P, CY, AC	A <- A + E + CY
		unimplemented(0x8b)
	case 0x8c:
		// 0x8c	ADC H	1	Z, S, P, CY, AC	A <- A + H + CY
		unimplemented(0x8c)
	case 0x8d:
		// 0x8d	ADC L	1	Z, S, P, CY, AC	A <- A + L + CY
		unimplemented(0x8d)
	case 0x8e:
		// 0x8e	ADC M	1	Z, S, P, CY, AC	A <- A + (HL) + CY
		unimplemented(0x8e)
	case 0x8f:
		// 0x8f	ADC A	1	Z, S, P, CY, AC	A <- A + A + CY
		unimplemented(0x8f)
	case 0x90:
		// 0x90	SUB B	1	Z, S, P, CY, AC	A <- A - B
		unimplemented(0x90)
	case 0x91:
		// 0x91	SUB C	1	Z, S, P, CY, AC	A <- A - C
		unimplemented(0x91)
	case 0x92:
		// 0x92	SUB D	1	Z, S, P, CY, AC	A <- A + D
		unimplemented(0x92)
	case 0x93:
		// 0x93	SUB E	1	Z, S, P, CY, AC	A <- A - E
		unimplemented(0x93)
	case 0x94:
		// 0x94	SUB H	1	Z, S, P, CY, AC	A <- A + H
		unimplemented(0x94)
	case 0x95:
		// 0x95	SUB L	1	Z, S, P, CY, AC	A <- A - L
		unimplemented(0x95)
	case 0x96:
		// 0x96	SUB M	1	Z, S, P, CY, AC	A <- A + (HL)
		unimplemented(0x96)
	case 0x97:
		// 0x97	SUB A	1	Z, S, P, CY, AC	A <- A - A
		unimplemented(0x97)
	case 0x98:
		// 0x98	SBB B	1	Z, S, P, CY, AC	A <- A - B - CY
		unimplemented(0x98)
	case 0x99:
		// 0x99	SBB C	1	Z, S, P, CY, AC	A <- A - C - CY
		unimplemented(0x99)
	case 0x9a:
		// 0x9a	SBB D	1	Z, S, P, CY, AC	A <- A - D - CY
		unimplemented(0x9a)
	case 0x9b:
		// 0x9b	SBB E	1	Z, S, P, CY, AC	A <- A - E - CY
		unimplemented(0x9b)
	case 0x9c:
		// 0x9c	SBB H	1	Z, S, P, CY, AC	A <- A - H - CY
		unimplemented(0x9c)
	case 0x9d:
		// 0x9d	SBB L	1	Z, S, P, CY, AC	A <- A - L - CY
		unimplemented(0x9d)
	case 0x9e:
		// 0x9e	SBB M	1	Z, S, P, CY, AC	A <- A - (HL) - CY
		unimplemented(0x9e)
	case 0x9f:
		// 0x9f	SBB A	1	Z, S, P, CY, AC	A <- A - A - CY
		unimplemented(0x9f)
	case 0xa0:
		// 0xa0	ANA B	1	Z, S, P, CY, AC	A <- A & B
		unimplemented(0xa0)
	case 0xa1:
		// 0xa1	ANA C	1	Z, S, P, CY, AC	A <- A & C
		unimplemented(0xa1)
	case 0xa2:
		// 0xa2	ANA D	1	Z, S, P, CY, AC	A <- A & D
		unimplemented(0xa2)
	case 0xa3:
		// 0xa3	ANA E	1	Z, S, P, CY, AC	A <- A & E
		unimplemented(0xa3)
	case 0xa4:
		// 0xa4	ANA H	1	Z, S, P, CY, AC	A <- A & H
		unimplemented(0xa4)
	case 0xa5:
		// 0xa5	ANA L	1	Z, S, P, CY, AC	A <- A & L
		unimplemented(0xa5)
	case 0xa6:
		// 0xa6	ANA M	1	Z, S, P, CY, AC	A <- A & (HL)
		unimplemented(0xa6)
	case 0xa7:
		// 0xa7	ANA A	1	Z, S, P, CY, AC	A <- A & A
		unimplemented(0xa7)
	case 0xa8:
		// 0xa8	XRA B	1	Z, S, P, CY, AC	A <- A ^ B
		unimplemented(0xa8)
	case 0xa9:
		// 0xa9	XRA C	1	Z, S, P, CY, AC	A <- A ^ C
		unimplemented(0xa9)
	case 0xaa:
		// 0xaa	XRA D	1	Z, S, P, CY, AC	A <- A ^ D
		unimplemented(0xaa)
	case 0xab:
		// 0xab	XRA E	1	Z, S, P, CY, AC	A <- A ^ E
		unimplemented(0xab)
	case 0xac:
		// 0xac	XRA H	1	Z, S, P, CY, AC	A <- A ^ H
		unimplemented(0xac)
	case 0xad:
		// 0xad	XRA L	1	Z, S, P, CY, AC	A <- A ^ L
		unimplemented(0xad)
	case 0xae:
		// 0xae	XRA M	1	Z, S, P, CY, AC	A <- A ^ (HL)
		unimplemented(0xae)
	case 0xaf:
		// 0xaf	XRA A	1	Z, S, P, CY, AC	A <- A ^ A
		unimplemented(0xaf)
	case 0xb0:
		// 0xb0	ORA B	1	Z, S, P, CY, AC	A <- A | B
		unimplemented(0xb0)
	case 0xb1:
		// 0xb1	ORA C	1	Z, S, P, CY, AC	A <- A | C
		unimplemented(0xb1)
	case 0xb2:
		// 0xb2	ORA D	1	Z, S, P, CY, AC	A <- A | D
		unimplemented(0xb2)
	case 0xb3:
		// 0xb3	ORA E	1	Z, S, P, CY, AC	A <- A | E
		unimplemented(0xb3)
	case 0xb4:
		// 0xb4	ORA H	1	Z, S, P, CY, AC	A <- A | H
		unimplemented(0xb4)
	case 0xb5:
		// 0xb5	ORA L	1	Z, S, P, CY, AC	A <- A | L
		unimplemented(0xb5)
	case 0xb6:
		// 0xb6	ORA M	1	Z, S, P, CY, AC	A <- A | (HL)
		unimplemented(0xb6)
	case 0xb7:
		// 0xb7	ORA A	1	Z, S, P, CY, AC	A <- A | A
		unimplemented(0xb7)
	case 0xb8:
		// 0xb8	CMP B	1	Z, S, P, CY, AC	A - B
		unimplemented(0xb8)
	case 0xb9:
		// 0xb9	CMP C	1	Z, S, P, CY, AC	A - C
		unimplemented(0xb9)
	case 0xba:
		// 0xba	CMP D	1	Z, S, P, CY, AC	A - D
		unimplemented(0xba)
	case 0xbb:
		// 0xbb	CMP E	1	Z, S, P, CY, AC	A - E
		unimplemented(0xbb)
	case 0xbc:
		// 0xbc	CMP H	1	Z, S, P, CY, AC	A - H
		unimplemented(0xbc)
	case 0xbd:
		// 0xbd	CMP L	1	Z, S, P, CY, AC	A - L
		unimplemented(0xbd)
	case 0xbe:
		// 0xbe	CMP M	1	Z, S, P, CY, AC	A - (HL)
		unimplemented(0xbe)
	case 0xbf:
		// 0xbf	CMP A	1	Z, S, P, CY, AC	A - A
		unimplemented(0xbf)
	case 0xc0:
		// 0xc0	RNZ	1		if NZ, RET
		unimplemented(0xc0)
	case 0xc1:
		// 0xc1	POP B	1		C <- (sp); B <- (sp+1); sp <- sp+2
		unimplemented(0xc1)
	case 0xc2:
		// 0xc2	JNZ adr	3		if NZ, PC <- adr
		unimplemented(0xc2)
	case 0xc3:
		// 0xc3	JMP adr	3		PC <= adr
		unimplemented(0xc3)
	case 0xc4:
		// 0xc4	CNZ adr	3		if NZ, CALL adr
		unimplemented(0xc4)
	case 0xc5:
		// 0xc5	PUSH B	1		(sp-2)<-C; (sp-1)<-B; sp <- sp - 2
		unimplemented(0xc5)
	case 0xc6:
		// 0xc6	ADI D8	2	Z, S, P, CY, AC	A <- A + byte
		unimplemented(0xc6)
	case 0xc7:
		// 0xc7	RST 0	1		CALL $0
		unimplemented(0xc7)
	case 0xc8:
		// 0xc8	RZ	1		if Z, RET
		unimplemented(0xc8)
	case 0xc9:
		// 0xc9	RET	1		PC.lo <- (sp); PC.hi<-(sp+1); SP <- SP+2
		unimplemented(0xc9)
	case 0xca:
		// 0xca	JZ adr	3		if Z, PC <- adr
		unimplemented(0xca)
	case 0xcb:
		// 0xcb	-
		unimplemented(0xcb)
	case 0xcc:
		// 0xcc	CZ adr	3		if Z, CALL adr
		unimplemented(0xcc)
	case 0xcd:
		// 0xcd	CALL adr	3		(SP-1)<-PC.hi;(SP-2)<-PC.lo;SP<-SP+2;PC=adr
		unimplemented(0xcd)
	case 0xce:
		// 0xce	ACI D8	2	Z, S, P, CY, AC	A <- A + data + CY
		unimplemented(0xce)
	case 0xcf:
		// 0xcf	RST 1	1		CALL $8
		unimplemented(0xcf)
	case 0xd0:
		// 0xd0	RNC	1		if NCY, RET
		unimplemented(0xd0)
	case 0xd1:
		// 0xd1	POP D	1		E <- (sp); D <- (sp+1); sp <- sp+2
		unimplemented(0xd1)
	case 0xd2:
		// 0xd2	JNC adr	3		if NCY, PC<-adr
		unimplemented(0xd2)
	case 0xd3:
		// 0xd3	OUT D8	2		special
		unimplemented(0xd3)
	case 0xd4:
		// 0xd4	CNC adr	3		if NCY, CALL adr
		unimplemented(0xd4)
	case 0xd5:
		// 0xd5	PUSH D	1		(sp-2)<-E; (sp-1)<-D; sp <- sp - 2
		unimplemented(0xd5)
	case 0xd6:
		// 0xd6	SUI D8	2	Z, S, P, CY, AC	A <- A - data
		unimplemented(0xd6)
	case 0xd7:
		// 0xd7	RST 2	1		CALL $10
		unimplemented(0xd7)
	case 0xd8:
		// 0xd8	RC	1		if CY, RET
		unimplemented(0xd8)
	case 0xd9:
		// 0xd9	-
		unimplemented(0xd9)
	case 0xda:
		// 0xda	JC adr	3		if CY, PC<-adr
		unimplemented(0xda)
	case 0xdb:
		// 0xdb	IN D8	2		special
		unimplemented(0xdb)
	case 0xdc:
		// 0xdc	CC adr	3		if CY, CALL adr
		unimplemented(0xdc)
	case 0xdd:
		// 0xdd	-
		unimplemented(0xdd)
	case 0xde:
		// 0xde	SBI D8	2	Z, S, P, CY, AC	A <- A - data - CY
		unimplemented(0xde)
	case 0xdf:
		// 0xdf	RST 3	1		CALL $18
		unimplemented(0xdf)
	case 0xe0:
		// 0xe0	RPO	1		if PO, RET
		unimplemented(0xe0)
	case 0xe1:
		// 0xe1	POP H	1		L <- (sp); H <- (sp+1); sp <- sp+2
		unimplemented(0xe1)
	case 0xe2:
		// 0xe2	JPO adr	3		if PO, PC <- adr
		unimplemented(0xe2)
	case 0xe3:
		// 0xe3	XTHL	1		L <-> (SP); H <-> (SP+1)
		unimplemented(0xe3)
	case 0xe4:
		// 0xe4	CPO adr	3		if PO, CALL adr
		unimplemented(0xe4)
	case 0xe5:
		// 0xe5	PUSH H	1		(sp-2)<-L; (sp-1)<-H; sp <- sp - 2
		unimplemented(0xe5)
	case 0xe6:
		// 0xe6	ANI D8	2	Z, S, P, CY, AC	A <- A & data
		unimplemented(0xe6)
	case 0xe7:
		// 0xe7	RST 4	1		CALL $20
		unimplemented(0xe7)
	case 0xe8:
		// 0xe8	RPE	1		if PE, RET
		unimplemented(0xe8)
	case 0xe9:
		// 0xe9	PCHL	1		PC.hi <- H; PC.lo <- L
		unimplemented(0xe9)
	case 0xea:
		// 0xea	JPE adr	3		if PE, PC <- adr
		unimplemented(0xea)
	case 0xeb:
		// 0xeb	XCHG	1		H <-> D; L <-> E
		unimplemented(0xeb)
	case 0xec:
		// 0xec	CPE adr	3		if PE, CALL adr
		unimplemented(0xec)
	case 0xed:
		// 0xed	-
		unimplemented(0xed)
	case 0xee:
		// 0xee	XRI D8	2	Z, S, P, CY, AC	A <- A ^ data
		unimplemented(0xee)
	case 0xef:
		// 0xef	RST 5	1		CALL $28
		unimplemented(0xef)
	case 0xf0:
		// 0xf0	RP	1		if P, RET
		unimplemented(0xf0)
	case 0xf1:
		// 0xf1	POP PSW	1		flags <- (sp); A <- (sp+1); sp <- sp+2
		unimplemented(0xf1)
	case 0xf2:
		// 0xf2	JP adr	3		if P=1 PC <- adr
		unimplemented(0xf2)
	case 0xf3:
		// 0xf3	DI	1		special
		unimplemented(0xf3)
	case 0xf4:
		// 0xf4	CP adr	3		if P, PC <- adr
		unimplemented(0xf4)
	case 0xf5:
		// 0xf5	PUSH PSW	1		(sp-2)<-flags; (sp-1)<-A; sp <- sp - 2
		unimplemented(0xf5)
	case 0xf6:
		// 0xf6	ORI D8	2	Z, S, P, CY, AC	A <- A | data
		unimplemented(0xf6)
	case 0xf7:
		// 0xf7	RST 6	1		CALL $30
		unimplemented(0xf7)
	case 0xf8:
		// 0xf8	RM	1		if M, RET
		unimplemented(0xf8)
	case 0xf9:
		// 0xf9	SPHL	1		SP=HL
		unimplemented(0xf9)
	case 0xfa:
		// 0xfa	JM adr	3		if M, PC <- adr
		unimplemented(0xfa)
	case 0xfb:
		// 0xfb	EI	1		special
		unimplemented(0xfb)
	case 0xfc:
		// 0xfc	CM adr	3		if M, CALL adr
		unimplemented(0xfc)
	case 0xfd:
		// 0xfd	-
		unimplemented(0xfd)
	case 0xfe:
		// 0xfe	CPI D8	2	Z, S, P, CY, AC	A - data
		unimplemented(0xfe)
	case 0xff:
		// 0xff	RST 7	1		CALL $38
		unimplemented(0xff)
	}
}
