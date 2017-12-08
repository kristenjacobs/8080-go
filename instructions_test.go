package main

import "testing"

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
	check_MOV("Test_0x41_MOV_B_C", t, instr_0x41_MOV_B_C, ms, &ms.regC, &ms.regC)
	check_MOV("Test_0x42_MOV_B_D", t, instr_0x42_MOV_B_D, ms, &ms.regD, &ms.regD)
	check_MOV("Test_0x43_MOV_B_E", t, instr_0x43_MOV_B_E, ms, &ms.regE, &ms.regE)
	check_MOV("Test_0x44_MOV_B_H", t, instr_0x44_MOV_B_H, ms, &ms.regH, &ms.regH)
	check_MOV("Test_0x45_MOV_B_L", t, instr_0x45_MOV_B_L, ms, &ms.regL, &ms.regL)
}

func check_MOV(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, fromReg *uint8, toReg *uint8) {
	*toReg = 0
	*fromReg = 9
	instrFunc(ms)
	if *toReg != 9 {
		t.Errorf("%s: expected toReg=9, got toReg=%d", testName, *toReg)
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
