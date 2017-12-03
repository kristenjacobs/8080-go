package main

import "testing"

func Test_0x05_DCR_B(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x05_DCR_B", t, instr_0x05_DCR_B, ms, &ms.regB)
}

func Test_0x0d_DCR_C(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x0d_DCR_C", t, instr_0x0d_DCR_C, ms, &ms.regC)
}

func Test_0x15_DCR_D(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x15_DCR_D", t, instr_0x15_DCR_D, ms, &ms.regD)
}

func Test_0x1d_DCR_E(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x1d_DCR_E", t, instr_0x1d_DCR_E, ms, &ms.regE)
}

func Test_0x25_DCR_H(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x25_DCR_H", t, instr_0x25_DCR_H, ms, &ms.regH)
}

func Test_0x2d_DCR_L(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x2d_DCR_L", t, instr_0x2d_DCR_L, ms, &ms.regL)
}

func Test_0x3d_DCR_A(t *testing.T) {
	ms := newMachineState()
	check_DCR("Test_0x3d_DCR_A", t, instr_0x3d_DCR_A, ms, &ms.regA)
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

func check_DCR(testName string, t *testing.T, instrFunc func(*machineState), ms *machineState, reg *uint8) {
	*reg = 10
	instrFunc(ms)
	if *reg != 9 {
		t.Errorf("%s: expected reg=9, got reg=%d", *reg)
	}
	if ms.flagZ {
		t.Errorf("%s: expected z=false, got z=true")
	}
	if ms.flagS {
		t.Errorf("%s: expected s=false, got s=true")
	}
	if !ms.flagP {
		t.Errorf("%s: expected p=true, got p=false")
	}
}
