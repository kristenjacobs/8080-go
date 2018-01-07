package main

import (
	"io/ioutil"
	"testing"
)

func init() {
	InitLogging(ioutil.Discard, ioutil.Discard)
}

func Test_LXI(t *testing.T) {
	ms := newMachineState()
	check_LXI_pair("Test_0x01_LXI_B_D16", t, instr_0x01_LXI_B_D16, ms, &ms.regB, &ms.regC)
	check_LXI_pair("Test_0x11_LXI_D_D16", t, instr_0x11_LXI_D_D16, ms, &ms.regD, &ms.regE)
	check_LXI_pair("Test_0x21_LXI_H_D16", t, instr_0x21_LXI_H_D16, ms, &ms.regH, &ms.regL)
	check_LXI_single("Test_0x31_LXI_SP_D16", t, instr_0x31_LXI_SP_D16, ms, &ms.sp)
}

func check_LXI_pair(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, regHi *uint8, regLo *uint8) {
	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xBE, 0xBA}, 2)
	instrFunc(ms)
	if *regHi != 0xBA {
		t.Errorf("%s: expected 0xBA, got %02x", testName, *regHi)
	}
	if *regLo != 0xBE {
		t.Errorf("%s: expected 0xBE, got %02x", testName, *regLo)
	}
}

func check_LXI_single(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, res *uint16) {
	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xBE, 0xBA}, 2)
	instrFunc(ms)
	if *res != 0xBABE {
		t.Errorf("%s: expected 0xBABE, got %02x", testName, *res)
	}
}

func Test_MVI(t *testing.T) {
	ms := newMachineState()
	check_MVI("Test_0x06_MVI_B_D8(ms)", t, instr_0x06_MVI_B_D8, ms, &ms.regB)
	check_MVI("Test_0x0e_MVI_C_D8(ms)", t, instr_0x0e_MVI_C_D8, ms, &ms.regC)
	check_MVI("Test_0x16_MVI_D_D8(ms)", t, instr_0x16_MVI_D_D8, ms, &ms.regD)
	check_MVI("Test_0x1e_MVI_E_D8(ms)", t, instr_0x1e_MVI_E_D8, ms, &ms.regE)
	check_MVI("Test_0x26_MVI_H_D8(ms)", t, instr_0x26_MVI_H_D8, ms, &ms.regH)
	check_MVI("Test_0x2e_MVI_L_D8(ms)", t, instr_0x2e_MVI_L_D8, ms, &ms.regL)
	check_MVI("Test_0x3e_MVI_A_D8(ms)", t, instr_0x3e_MVI_A_D8, ms, &ms.regA)

	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xAB}, 1)
	ms.regH = uint8(RAM_BASE >> 8)
	ms.regL = uint8(RAM_BASE & 0xFF)
	instr_0x36_MVI_M_D8(ms)
	result := ms.readMem(RAM_BASE, 1)[0]
	if result != 0xAB {
		t.Errorf("Test_0x3e_MVI_M_D8: expected 0xAB, got %02x", result)
	}
	if ms.pc != RAM_BASE+2 {
		t.Errorf("Test_0x3e_MVI_M_D8: expected pc=0x%04x, got %04x", RAM_BASE+2, ms.pc)
	}
}

func check_MVI(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, res *uint8) {
	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xAB}, 1)
	instrFunc(ms)
	if *res != 0xAB {
		t.Errorf("%s: expected 0xAB, got %02x", testName, *res)
	}
	if ms.pc != RAM_BASE+2 {
		t.Errorf("%s: expected pc=0x%04x, got %04x", testName, RAM_BASE+2, ms.pc)
	}
}

func Test_DCR(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x05_DCR_B", t, instr_0x05_DCR_B, ms, &ms.regB)
	check_DCR("Test_0x0d_DCR_C", t, instr_0x0d_DCR_C, ms, &ms.regC)
	check_DCR("Test_0x15_DCR_D", t, instr_0x15_DCR_D, ms, &ms.regD)
	check_DCR("Test_0x1d_DCR_E", t, instr_0x1d_DCR_E, ms, &ms.regE)
	check_DCR("Test_0x25_DCR_H", t, instr_0x25_DCR_H, ms, &ms.regH)
	check_DCR("Test_0x2d_DCR_L", t, instr_0x2d_DCR_L, ms, &ms.regL)
	check_DCR("Test_0x3d_DCR_A", t, instr_0x3d_DCR_A, ms, &ms.regA)
}

func check_DCR(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, reg *uint8) {
	*reg = 10
	instrFunc(ms)
	if *reg != 9 {
		t.Errorf("%s: expected reg=9, got reg=%d", testName, *reg)
	}
	if ms.flagZ {
		t.Errorf("%s: expected z=false, got z=true", testName)
	}
	if ms.flagS {
		t.Errorf("%s: expected s=false, got s=true", testName)
	}
	if !ms.flagP {
		t.Errorf("%s: expected p=true, got p=false", testName)
	}
}

func Test_MOV_REG_REG(t *testing.T) {
	ms := newMachineState()
	check_MOV_REG_REG("Test_0x40_MOV_B_B", t, instr_0x40_MOV_B_B, ms, &ms.regB, &ms.regB)
	check_MOV_REG_REG("Test_0x41_MOV_B_C", t, instr_0x41_MOV_B_C, ms, &ms.regB, &ms.regC)
	check_MOV_REG_REG("Test_0x42_MOV_B_D", t, instr_0x42_MOV_B_D, ms, &ms.regB, &ms.regD)
	check_MOV_REG_REG("Test_0x43_MOV_B_E", t, instr_0x43_MOV_B_E, ms, &ms.regB, &ms.regE)
	check_MOV_REG_REG("Test_0x44_MOV_B_H", t, instr_0x44_MOV_B_H, ms, &ms.regB, &ms.regH)
	check_MOV_REG_REG("Test_0x45_MOV_B_L", t, instr_0x45_MOV_B_L, ms, &ms.regB, &ms.regL)
	check_MOV_REG_REG("Test_0x47_MOV_B_A", t, instr_0x47_MOV_B_A, ms, &ms.regB, &ms.regA)
	check_MOV_REG_REG("Test_0x48_MOV_C_B", t, instr_0x48_MOV_C_B, ms, &ms.regC, &ms.regB)
	check_MOV_REG_REG("Test_0x49_MOV_C_C", t, instr_0x49_MOV_C_C, ms, &ms.regC, &ms.regC)
	check_MOV_REG_REG("Test_0x4a_MOV_C_D", t, instr_0x4a_MOV_C_D, ms, &ms.regC, &ms.regD)
	check_MOV_REG_REG("Test_0x4b_MOV_C_E", t, instr_0x4b_MOV_C_E, ms, &ms.regC, &ms.regE)
	check_MOV_REG_REG("Test_0x4c_MOV_C_H", t, instr_0x4c_MOV_C_H, ms, &ms.regC, &ms.regH)
	check_MOV_REG_REG("Test_0x4d_MOV_C_L", t, instr_0x4d_MOV_C_L, ms, &ms.regC, &ms.regL)
	check_MOV_REG_REG("Test_0x4f_MOV_C_A", t, instr_0x4f_MOV_C_A, ms, &ms.regC, &ms.regA)
	check_MOV_REG_REG("Test_0x50_MOV_D_B", t, instr_0x50_MOV_D_B, ms, &ms.regD, &ms.regB)
	check_MOV_REG_REG("Test_0x51_MOV_D_C", t, instr_0x51_MOV_D_C, ms, &ms.regD, &ms.regC)
	check_MOV_REG_REG("Test_0x52_MOV_D_D", t, instr_0x52_MOV_D_D, ms, &ms.regD, &ms.regD)
	check_MOV_REG_REG("Test_0x53_MOV_D_E", t, instr_0x53_MOV_D_E, ms, &ms.regD, &ms.regE)
	check_MOV_REG_REG("Test_0x54_MOV_D_H", t, instr_0x54_MOV_D_H, ms, &ms.regD, &ms.regH)
	check_MOV_REG_REG("Test_0x55_MOV_D_L", t, instr_0x55_MOV_D_L, ms, &ms.regD, &ms.regL)
	check_MOV_REG_REG("Test_0x57_MOV_D_A", t, instr_0x57_MOV_D_A, ms, &ms.regD, &ms.regA)
	check_MOV_REG_REG("Test_0x58_MOV_E_B", t, instr_0x58_MOV_E_B, ms, &ms.regE, &ms.regB)
	check_MOV_REG_REG("Test_0x59_MOV_E_C", t, instr_0x59_MOV_E_C, ms, &ms.regE, &ms.regC)
	check_MOV_REG_REG("Test_0x5a_MOV_E_D", t, instr_0x5a_MOV_E_D, ms, &ms.regE, &ms.regD)
	check_MOV_REG_REG("Test_0x5b_MOV_E_E", t, instr_0x5b_MOV_E_E, ms, &ms.regE, &ms.regE)
	check_MOV_REG_REG("Test_0x5c_MOV_E_H", t, instr_0x5c_MOV_E_H, ms, &ms.regE, &ms.regH)
	check_MOV_REG_REG("Test_0x5d_MOV_E_L", t, instr_0x5d_MOV_E_L, ms, &ms.regE, &ms.regL)
	check_MOV_REG_REG("Test_0x5f_MOV_E_A", t, instr_0x5f_MOV_E_A, ms, &ms.regE, &ms.regA)
	check_MOV_REG_REG("Test_0x60_MOV_H_B", t, instr_0x60_MOV_H_B, ms, &ms.regH, &ms.regB)
	check_MOV_REG_REG("Test_0x61_MOV_H_C", t, instr_0x61_MOV_H_C, ms, &ms.regH, &ms.regC)
	check_MOV_REG_REG("Test_0x62_MOV_H_D", t, instr_0x62_MOV_H_D, ms, &ms.regH, &ms.regD)
	check_MOV_REG_REG("Test_0x63_MOV_H_E", t, instr_0x63_MOV_H_E, ms, &ms.regH, &ms.regE)
	check_MOV_REG_REG("Test_0x64_MOV_H_H", t, instr_0x64_MOV_H_H, ms, &ms.regH, &ms.regH)
	check_MOV_REG_REG("Test_0x65_MOV_H_L", t, instr_0x65_MOV_H_L, ms, &ms.regH, &ms.regL)
	check_MOV_REG_REG("Test_0x67_MOV_H_A", t, instr_0x67_MOV_H_A, ms, &ms.regH, &ms.regA)
	check_MOV_REG_REG("Test_0x68_MOV_L_B", t, instr_0x68_MOV_L_B, ms, &ms.regL, &ms.regB)
	check_MOV_REG_REG("Test_0x69_MOV_L_C", t, instr_0x69_MOV_L_C, ms, &ms.regL, &ms.regC)
	check_MOV_REG_REG("Test_0x6a_MOV_L_D", t, instr_0x6a_MOV_L_D, ms, &ms.regL, &ms.regD)
	check_MOV_REG_REG("Test_0x6b_MOV_L_E", t, instr_0x6b_MOV_L_E, ms, &ms.regL, &ms.regE)
	check_MOV_REG_REG("Test_0x6c_MOV_L_H", t, instr_0x6c_MOV_L_H, ms, &ms.regL, &ms.regH)
	check_MOV_REG_REG("Test_0x6d_MOV_L_L", t, instr_0x6d_MOV_L_L, ms, &ms.regL, &ms.regL)
	check_MOV_REG_REG("Test_0x6f_MOV_L_A", t, instr_0x6f_MOV_L_A, ms, &ms.regL, &ms.regA)
	check_MOV_REG_REG("Test_0x78_MOV_A_B", t, instr_0x78_MOV_A_B, ms, &ms.regA, &ms.regB)
	check_MOV_REG_REG("Test_0x79_MOV_A_C", t, instr_0x79_MOV_A_C, ms, &ms.regA, &ms.regC)
	check_MOV_REG_REG("Test_0x7a_MOV_A_D", t, instr_0x7a_MOV_A_D, ms, &ms.regA, &ms.regD)
	check_MOV_REG_REG("Test_0x7b_MOV_A_E", t, instr_0x7b_MOV_A_E, ms, &ms.regA, &ms.regE)
	check_MOV_REG_REG("Test_0x7c_MOV_A_H", t, instr_0x7c_MOV_A_H, ms, &ms.regA, &ms.regH)
	check_MOV_REG_REG("Test_0x7d_MOV_A_L", t, instr_0x7d_MOV_A_L, ms, &ms.regA, &ms.regL)
	check_MOV_REG_REG("Test_0x7f_MOV_A_A", t, instr_0x7f_MOV_A_A, ms, &ms.regA, &ms.regA)
}

func check_MOV_REG_REG(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, dstReg *uint8, srcReg *uint8) {
	*dstReg = 0
	*srcReg = 9
	instrFunc(ms)
	if *dstReg != 9 {
		t.Errorf("%s: expected dstReg=9, got datReg=%d", testName, *dstReg)
	}
}

func Test_MOV_REG_MEM(t *testing.T) {
	ms := newMachineState()
	check_MOV_REG_MEM("Test_0x46_MOV_B_M", t, instr_0x46_MOV_B_M, ms, &ms.regB)
	check_MOV_REG_MEM("Test_0x4e_MOV_C_M", t, instr_0x4e_MOV_C_M, ms, &ms.regC)
	check_MOV_REG_MEM("Test_0x56_MOV_D_M", t, instr_0x56_MOV_D_M, ms, &ms.regD)
	check_MOV_REG_MEM("Test_0x5e_MOV_E_M", t, instr_0x5e_MOV_E_M, ms, &ms.regE)
	check_MOV_REG_MEM("Test_0x66_MOV_H_M", t, instr_0x66_MOV_H_M, ms, &ms.regH)
	check_MOV_REG_MEM("Test_0x6e_MOV_L_M", t, instr_0x6e_MOV_L_M, ms, &ms.regL)
	check_MOV_REG_MEM("Test_0x7e_MOV_A_M", t, instr_0x7e_MOV_A_M, ms, &ms.regA)

}

func check_MOV_REG_MEM(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, dstReg *uint8) {
	*dstReg = 0
	ms.regH = uint8(RAM_BASE >> 8)
	ms.regL = uint8(RAM_BASE & 0xFF)
	ms.writeMem(getPair(ms.regH, ms.regL), []uint8{0xFF}, 1)
	instrFunc(ms)
	if *dstReg != 0xFF {
		t.Errorf("%s: expected dstReg=0xFF, got datReg=%d", testName, *dstReg)
	}
}

func Test_MOV_MEM_REG(t *testing.T) {
	ms := newMachineState()
	check_MOV_MEM_REG("Test_0x70_MOV_M_B", t, instr_0x70_MOV_M_B, ms, &ms.regB)
	check_MOV_MEM_REG("Test_0x71_MOV_M_C", t, instr_0x71_MOV_M_C, ms, &ms.regC)
	check_MOV_MEM_REG("Test_0x72_MOV_M_D", t, instr_0x72_MOV_M_D, ms, &ms.regD)
	check_MOV_MEM_REG("Test_0x73_MOV_M_E", t, instr_0x73_MOV_M_E, ms, &ms.regE)
	check_MOV_MEM_REG("Test_0x74_MOV_M_H", t, instr_0x74_MOV_M_H, ms, &ms.regH)
	check_MOV_MEM_REG("Test_0x75_MOV_M_L", t, instr_0x75_MOV_M_L, ms, &ms.regL)
	check_MOV_MEM_REG("Test_0x77_MOV_M_A", t, instr_0x77_MOV_M_A, ms, &ms.regA)
}

func check_MOV_MEM_REG(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, srcReg *uint8) {
	*srcReg = 0xAB
	ms.regH = uint8(RAM_BASE >> 8)
	ms.regL = uint8(RAM_BASE & 0xFF)
	expected := *srcReg
	ms.writeMem(getPair(ms.regH, ms.regL), []uint8{0x00}, 1)
	instrFunc(ms)
	result := ms.readMem(RAM_BASE, 1)[0]
	if result != expected {
		t.Errorf("%s: expected [0x%04x]=0x%02x, got [0x%04x]=0x%02x", testName, RAM_BASE, expected, RAM_BASE, result)
	}
}

func Test_0xc9_RET(t *testing.T) {
	ms := newMachineState()
	ms.sp = RAM_BASE
	ms.writeMem(ms.sp, []uint8{0x1, 0x2}, 2)
	instr_0xc9_RET(ms)
	if ms.pc != 0x201 {
		t.Errorf("instr_0xc9_RET: expected pc=0x201, got pc=0x%04x", ms.pc)
	}
	if ms.sp != RAM_BASE+2 {
		t.Errorf("instr_0xc9_RET: expected sp=0x%04x, got sp=0x%04x", RAM_BASE+2, ms.sp)
	}
}

func Test_0xcd_CALL_adr(t *testing.T) {
	ms := newMachineState()
	ms.sp = RAM_BASE + 2
	ms.writeMem(ms.sp-2, []uint8{0x00, 0x00}, 2)
	ms.pc = RAM_BASE + 10
	ms.writeMem(ms.pc+1, []uint8{0xAD, 0xDE}, 2)
	instr_0xcd_CALL_adr(ms)
	bytes := ms.readMem(RAM_BASE, 2)
	var sp uint16 = (uint16(bytes[1]) << 8) | uint16(bytes[0])
	if sp != RAM_BASE+13 {
		t.Errorf("instr_0xcd_CALL_adr: expected [sp]=0x%04X, got [sp]=0x%04x", RAM_BASE+13, sp)
	}
	if ms.sp != RAM_BASE {
		t.Errorf("instr_0xcd_CALL_adr: expected sp=0x%04X, got sp=0x%04x", RAM_BASE, ms.sp)
	}
	if ms.pc != 0xDEAD {
		t.Errorf("instr_0xcd_CALL_adr: expected pc=0xDEAD, got pc=0x%04x", ms.pc)
	}
}

func Test_LDAX(t *testing.T) {
	ms := newMachineState()
	check_LDAX("instr_0x0a_LDAX_B", t, instr_0x0a_LDAX_B, ms, &ms.regB, &ms.regC)
	check_LDAX("instr_0x1a_LDAX_D", t, instr_0x1a_LDAX_D, ms, &ms.regD, &ms.regE)
}

func check_LDAX(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, adrRegHi *uint8, adrRegLo *uint8) {
	*adrRegLo = uint8(RAM_BASE >> 8)
	*adrRegHi = uint8(RAM_BASE & 0xFF)
	ms.writeMem(getPair(*adrRegHi, *adrRegLo), []uint8{0xFF}, 1)
	instrFunc(ms)
	if ms.regA != 0xFF {
		t.Errorf("%s: expected regA=0xFF, got regA=0x%02x", testName, ms.regA)
	}
}

func Test_INX(t *testing.T) {
	ms := newMachineState()
	check_INX("Test_0x03_INX_B", t, instr_0x03_INX_B, ms, &ms.regB, &ms.regC)
	check_INX("Test_0x13_INX_D", t, instr_0x13_INX_D, ms, &ms.regD, &ms.regE)
	check_INX("Test_0x23_INX_H", t, instr_0x23_INX_H, ms, &ms.regH, &ms.regL)

	ms.sp = 0x1234
	instr_0x33_INX_SP(ms)
	if ms.sp != 0x1235 {
		t.Errorf("Test_0x33)INX_SP: expected 0x1235, got 0x%02x", ms.sp)
	}
}

func check_INX(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, adrRegHi *uint8, adrRegLo *uint8) {
	*adrRegLo = uint8(0x12)
	*adrRegHi = uint8(0x34)
	instrFunc(ms)
	if *adrRegLo != 0x13 {
		t.Errorf("%s: expected 0x13, got 0x%02x", testName, *adrRegLo)
	}
	if *adrRegHi != 0x34 {
		t.Errorf("%s: expected 0x34, got 0x%02x", testName, *adrRegHi)
	}
}

func Test_conditional_jump(t *testing.T) {
	ms := newMachineState()
	check_conditional_jump("0xc2_JNZ_adr", t, instr_0xc2_JNZ_adr, ms, &ms.flagZ, false)
	check_conditional_jump("0xca_JZ_adr", t, instr_0xca_JZ_adr, ms, &ms.flagZ, true)
	check_conditional_jump("0xd2_JNC_adr", t, instr_0xd2_JNC_adr, ms, &ms.flagCY, false)
	check_conditional_jump("0xda_JC_adr", t, instr_0xda_JC_adr, ms, &ms.flagCY, true)
	check_conditional_jump("0xe2_JPO_adr", t, instr_0xe2_JPO_adr, ms, &ms.flagP, false)
	check_conditional_jump("0xea_JPE_adr", t, instr_0xea_JPE_adr, ms, &ms.flagP, true)
}

func check_conditional_jump(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, flag *bool, val bool) {
	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xBE, 0xBA}, 2)
	*flag = val
	instrFunc(ms)
	if ms.pc != 0xBABE {
		t.Errorf("%s: expected pc=0xBABE, got pc=0x%04x", testName, ms.pc)
	}
	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xBE, 0xBA}, 2)
	*flag = !val
	instrFunc(ms)
	if ms.pc != RAM_BASE+3 {
		t.Errorf("%s: expected pc=0x%04x, got pc=0x%04x", testName, RAM_BASE+3, ms.pc)
	}
}

func Test_0xc3_JMP_adr(t *testing.T) {
	ms := newMachineState()
	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xBE, 0xBA}, 2)
	instr_0xc2_JNZ_adr(ms)
	if ms.pc != 0xBABE {
		t.Errorf("instr_0xc3_JMP_adr: expected pc=0xBABE, got pc=0x%04x", ms.pc)
	}
}

func Test_PUSH(t *testing.T) {
	ms := newMachineState()
	check_PUSH("Test_0xc5_PUSH_B", t, instr_0xc5_PUSH_B, ms, &ms.regC, &ms.regB)
	check_PUSH("Test_0xd5_PUSH_D", t, instr_0xd5_PUSH_D, ms, &ms.regE, &ms.regD)
	check_PUSH("Test_0xe5_PUSH_H", t, instr_0xe5_PUSH_H, ms, &ms.regL, &ms.regH)
}

func check_PUSH(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, regHi *uint8, regLo *uint8) {
	ms.sp = RAM_BASE + 2
	*regLo = uint8(0x13)
	*regHi = uint8(0x34)
	instrFunc(ms)
	resLo := ms.readMem(RAM_BASE+1, 1)[0]
	resHi := ms.readMem(RAM_BASE, 1)[0]
	if resLo != 0x13 {
		t.Errorf("%s: expected 0x13, got 0x%02x", testName, resLo)
	}
	if resHi != 0x34 {
		t.Errorf("%s: expected 0x34, got 0x%02x", testName, resHi)
	}
	if ms.sp != RAM_BASE {
		t.Errorf("%s: expected sp=0x%04x, got 0x%02x", testName, RAM_BASE, ms.sp)
	}
}

func Test_POP(t *testing.T) {
	ms := newMachineState()
	check_POP("Test_0xc1_POP_B", t, instr_0xc1_POP_B, ms, &ms.regC, &ms.regB)
	check_POP("Test_0xd1_POP_D", t, instr_0xd1_POP_D, ms, &ms.regE, &ms.regD)
	check_POP("Test_0xe1_POP_H", t, instr_0xe1_POP_H, ms, &ms.regL, &ms.regH)
}

func check_POP(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, regHi *uint8, regLo *uint8) {
	ms.sp = RAM_BASE
	ms.writeMem(ms.sp, []uint8{0x12, 0x34}, 2)
	instrFunc(ms)
	if *regHi != 0x12 {
		t.Errorf("%s: expected 0x12, got 0x%02x", testName, regHi)
	}
	if *regLo != 0x34 {
		t.Errorf("%s: expected 0x34, got 0x%02x", testName, regLo)
	}
	if ms.sp != RAM_BASE+2 {
		t.Errorf("%s: expected sp=0x%04x, got 0x%02x", testName, RAM_BASE, ms.sp)
	}
}

func Test_DAD(t *testing.T) {
	ms := newMachineState()
	check_DAD("Test_0x09_DAD_B", t, instr_0x09_DAD_B, ms, &ms.regB, &ms.regC)
	check_DAD("Test_0x19_DAD_D", t, instr_0x19_DAD_D, ms, &ms.regD, &ms.regE)
	check_DAD("Test_0x29_DAD_H", t, instr_0x29_DAD_H, ms, &ms.regH, &ms.regL)
}

func check_DAD(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, regHi *uint8, regLo *uint8) {
	// Non-carry test
	ms.flagCY = false
	setPair(&ms.regH, &ms.regL, 14)
	setPair(regHi, regLo, 14)
	instrFunc(ms)
	result := getPair(ms.regH, ms.regL)
	if result != 28 {
		t.Errorf("%s: expected 26, got %d", testName, result)
	}
	if ms.flagCY != false {
		t.Errorf("%s: expected CY=false, CY=%t", testName, ms.flagCY)
	}
	// Carry test
	ms.flagCY = false
	setPair(&ms.regH, &ms.regL, 0x8001)
	setPair(regHi, regLo, 0x8001)
	instrFunc(ms)
	result = getPair(ms.regH, ms.regL)
	if result != 2 {
		t.Errorf("%s: expected 2, got %d", testName, result)
	}
	if ms.flagCY != true {
		t.Errorf("%s: expected CY=true, CY=%t", testName, ms.flagCY)
	}
}

func Test_0xeb_XCHG(t *testing.T) {
	ms := newMachineState()
	ms.regD = 1
	ms.regE = 2
	ms.regH = 3
	ms.regL = 4
	instr_0xeb_XCHG(ms)
	if ms.regD != 3 {
		t.Errorf("instr_0xeb_XCHG: expected regD=3, got regD=%d", ms.regD)
	}
	if ms.regE != 4 {
		t.Errorf("instr_0xeb_XCHG: expected regE=4, got regE=%d", ms.regE)
	}
	if ms.regH != 1 {
		t.Errorf("instr_0xeb_XCHG: expected regH=1, got regH=%d", ms.regH)
	}
	if ms.regL != 2 {
		t.Errorf("instr_0xeb_XCHG: expected regL=2, got regL=%d", ms.regL)
	}
}

func Test_RST(t *testing.T) {
	ms := newMachineState()
	check_RST("Test_0xc7_RST_0", t, instr_0xc7_RST_0, ms, 0x0)
	check_RST("Test_0xcf_RST_1", t, instr_0xcf_RST_1, ms, 0x8)
	check_RST("Test_0xd7_RST_2", t, instr_0xd7_RST_2, ms, 0x10)
	check_RST("Test_0xdf_RST_3", t, instr_0xdf_RST_3, ms, 0x18)
	check_RST("Test_0xe7_RST_4", t, instr_0xe7_RST_4, ms, 0x20)
	check_RST("Test_0xef_RST_5", t, instr_0xef_RST_5, ms, 0x28)
	check_RST("Test_0xf7_RST_6", t, instr_0xf7_RST_6, ms, 0x30)
	check_RST("Test_0xff_RST_7", t, instr_0xff_RST_7, ms, 0x38)
}

func check_RST(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, addr uint16) {
	ms.sp = RAM_BASE + 2
	ms.writeMem(ms.sp-2, []uint8{0x00, 0x00}, 2)
	ms.pc = RAM_BASE + 10
	instrFunc(ms)
	bytes := ms.readMem(RAM_BASE, 2)
	var sp uint16 = (uint16(bytes[1]) << 8) | uint16(bytes[0])
	if sp != RAM_BASE+11 {
		t.Errorf("%s: expected [sp]=0x%04X, got [sp]=0x%04x", testName, RAM_BASE+13, sp)
	}
	if ms.sp != RAM_BASE {
		t.Errorf("%s: expected sp=0x%04X, got sp=0x%04x", testName, RAM_BASE, ms.sp)
	}
	if ms.pc != addr {
		t.Errorf("%s: expected pc=0x%04x, got pc=0x%04x", testName, addr, ms.pc)
	}
}
