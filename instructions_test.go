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
		t.Errorf("%s: expected 0xBA, got %02x", *regHi)
	}
	if *regLo != 0xBE {
		t.Errorf("%s: expected 0xBE, got %02x", *regLo)
	}
}

func check_LXI_single(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, res *uint16) {
	ms.pc = RAM_BASE
	ms.writeMem(ms.pc+1, []uint8{0xBE, 0xBA}, 2)
	instrFunc(ms)
	if *res != 0xBABE {
		t.Errorf("%s: expected 0xBABE, got %02x", *res)
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

func Test_MOV(t *testing.T) {
	ms := newMachineState()
	check_MOV("Test_0x40_MOV_B_B", t, instr_0x40_MOV_B_B, ms, &ms.regB, &ms.regB)
	check_MOV("Test_0x41_MOV_B_C", t, instr_0x41_MOV_B_C, ms, &ms.regB, &ms.regC)
	check_MOV("Test_0x42_MOV_B_D", t, instr_0x42_MOV_B_D, ms, &ms.regB, &ms.regD)
	check_MOV("Test_0x43_MOV_B_E", t, instr_0x43_MOV_B_E, ms, &ms.regB, &ms.regE)
	check_MOV("Test_0x44_MOV_B_H", t, instr_0x44_MOV_B_H, ms, &ms.regB, &ms.regH)
	check_MOV("Test_0x45_MOV_B_L", t, instr_0x45_MOV_B_L, ms, &ms.regB, &ms.regL)
	check_MOV("Test_0x47_MOV_B_A", t, instr_0x47_MOV_B_A, ms, &ms.regB, &ms.regA)
	check_MOV("Test_0x48_MOV_C_B", t, instr_0x48_MOV_C_B, ms, &ms.regC, &ms.regB)
	check_MOV("Test_0x49_MOV_C_C", t, instr_0x49_MOV_C_C, ms, &ms.regC, &ms.regC)
	check_MOV("Test_0x4a_MOV_C_D", t, instr_0x4a_MOV_C_D, ms, &ms.regC, &ms.regD)
	check_MOV("Test_0x4b_MOV_C_E", t, instr_0x4b_MOV_C_E, ms, &ms.regC, &ms.regE)
	check_MOV("Test_0x4c_MOV_C_H", t, instr_0x4c_MOV_C_H, ms, &ms.regC, &ms.regH)
	check_MOV("Test_0x4d_MOV_C_L", t, instr_0x4d_MOV_C_L, ms, &ms.regC, &ms.regL)
	check_MOV("Test_0x4f_MOV_C_A", t, instr_0x4f_MOV_C_A, ms, &ms.regC, &ms.regA)
	check_MOV("Test_0x50_MOV_D_B", t, instr_0x50_MOV_D_B, ms, &ms.regD, &ms.regB)
	check_MOV("Test_0x51_MOV_D_C", t, instr_0x51_MOV_D_C, ms, &ms.regD, &ms.regC)
	check_MOV("Test_0x52_MOV_D_D", t, instr_0x52_MOV_D_D, ms, &ms.regD, &ms.regD)
	check_MOV("Test_0x53_MOV_D_E", t, instr_0x53_MOV_D_E, ms, &ms.regD, &ms.regE)
	check_MOV("Test_0x54_MOV_D_H", t, instr_0x54_MOV_D_H, ms, &ms.regD, &ms.regH)
	check_MOV("Test_0x55_MOV_D_L", t, instr_0x55_MOV_D_L, ms, &ms.regD, &ms.regL)
	check_MOV("Test_0x57_MOV_D_A", t, instr_0x57_MOV_D_A, ms, &ms.regD, &ms.regA)
	check_MOV("Test_0x58_MOV_E_B", t, instr_0x58_MOV_E_B, ms, &ms.regE, &ms.regB)
	check_MOV("Test_0x59_MOV_E_C", t, instr_0x59_MOV_E_C, ms, &ms.regE, &ms.regC)
	check_MOV("Test_0x5a_MOV_E_D", t, instr_0x5a_MOV_E_D, ms, &ms.regE, &ms.regD)
	check_MOV("Test_0x5b_MOV_E_E", t, instr_0x5b_MOV_E_E, ms, &ms.regE, &ms.regE)
	check_MOV("Test_0x5c_MOV_E_H", t, instr_0x5c_MOV_E_H, ms, &ms.regE, &ms.regH)
	check_MOV("Test_0x5d_MOV_E_L", t, instr_0x5d_MOV_E_L, ms, &ms.regE, &ms.regL)
	check_MOV("Test_0x5f_MOV_E_A", t, instr_0x5f_MOV_E_A, ms, &ms.regE, &ms.regA)
	check_MOV("Test_0x60_MOV_H_B", t, instr_0x60_MOV_H_B, ms, &ms.regH, &ms.regB)
	check_MOV("Test_0x61_MOV_H_C", t, instr_0x61_MOV_H_C, ms, &ms.regH, &ms.regC)
	check_MOV("Test_0x62_MOV_H_D", t, instr_0x62_MOV_H_D, ms, &ms.regH, &ms.regD)
	check_MOV("Test_0x63_MOV_H_E", t, instr_0x63_MOV_H_E, ms, &ms.regH, &ms.regE)
	check_MOV("Test_0x64_MOV_H_H", t, instr_0x64_MOV_H_H, ms, &ms.regH, &ms.regH)
	check_MOV("Test_0x65_MOV_H_L", t, instr_0x65_MOV_H_L, ms, &ms.regH, &ms.regL)
	check_MOV("Test_0x67_MOV_H_A", t, instr_0x67_MOV_H_A, ms, &ms.regH, &ms.regA)
	check_MOV("Test_0x68_MOV_L_B", t, instr_0x68_MOV_L_B, ms, &ms.regL, &ms.regB)
	check_MOV("Test_0x69_MOV_L_C", t, instr_0x69_MOV_L_C, ms, &ms.regL, &ms.regC)
	check_MOV("Test_0x6a_MOV_L_D", t, instr_0x6a_MOV_L_D, ms, &ms.regL, &ms.regD)
	check_MOV("Test_0x6b_MOV_L_E", t, instr_0x6b_MOV_L_E, ms, &ms.regL, &ms.regE)
	check_MOV("Test_0x6c_MOV_L_H", t, instr_0x6c_MOV_L_H, ms, &ms.regL, &ms.regH)
	check_MOV("Test_0x6d_MOV_L_L", t, instr_0x6d_MOV_L_L, ms, &ms.regL, &ms.regL)
	check_MOV("Test_0x6f_MOV_L_A", t, instr_0x6f_MOV_L_A, ms, &ms.regL, &ms.regA)
	check_MOV("Test_0x78_MOV_A_B", t, instr_0x78_MOV_A_B, ms, &ms.regA, &ms.regB)
	check_MOV("Test_0x79_MOV_A_C", t, instr_0x79_MOV_A_C, ms, &ms.regA, &ms.regC)
	check_MOV("Test_0x7a_MOV_A_D", t, instr_0x7a_MOV_A_D, ms, &ms.regA, &ms.regD)
	check_MOV("Test_0x7b_MOV_A_E", t, instr_0x7b_MOV_A_E, ms, &ms.regA, &ms.regE)
	check_MOV("Test_0x7c_MOV_A_H", t, instr_0x7c_MOV_A_H, ms, &ms.regA, &ms.regH)
	check_MOV("Test_0x7d_MOV_A_L", t, instr_0x7d_MOV_A_L, ms, &ms.regA, &ms.regL)
	check_MOV("Test_0x7f_MOV_A_A", t, instr_0x7f_MOV_A_A, ms, &ms.regA, &ms.regA)
}

func check_MOV(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, dstReg *uint8, srcReg *uint8) {
	*dstReg = 0
	*srcReg = 9
	instrFunc(ms)
	if *dstReg != 9 {
		t.Errorf("%s: expected dstReg=9, got datReg=%d", testName, *dstReg)
	}
}

func Test_0xc9_RET(t *testing.T) {
	ms := newMachineState()
	ms.sp = RAM_BASE
	ms.writeMem(ms.sp, []uint8{0x1, 0x2, 0x3}, 3)
	instr_0xc9_RET(ms)
	if ms.pc != 0x201 {
		t.Errorf("instr_0xc9_RET: expected pc=0x201, got pc=0x%04x", ms.pc)
	}
	if ms.sp != 0x3 {
		t.Errorf("instr_0xc9_RET: expected pc=0x3, got pc=0x%04x", ms.sp)
	}
}
