package core

import "testing"

const (
	RAM_BASE = 0x0
	RAM_SIZE = 0x100
)

func newTestMachineState() *MachineState {
	ms := NewMachineState(nil, RAM_BASE, RAM_BASE)
	ms.InitialiseRam(RAM_BASE, RAM_SIZE)
	return ms
}

func TestSetZ(t *testing.T) {
	ms := newTestMachineState()
	ms.setZ(0x1)
	if ms.flagZ {
		t.Errorf("expected z=false, got z=true")
	}
	ms.setZ(0x0)
	if !ms.flagZ {
		t.Errorf("expected z=true, got z=false")
	}
}

func TestSetS(t *testing.T) {
	ms := newTestMachineState()
	ms.setS(0x1)
	if ms.flagS {
		t.Errorf("expected s=false, got s=true")
	}
	ms.setS(0xFF)
	if !ms.flagS {
		t.Errorf("expected s=true, got s=false")
	}
}

func TestSetP(t *testing.T) {
	ms := newTestMachineState()
	ms.setP(0xF)
	if !ms.flagP {
		t.Errorf("expected p=true, got p=false")
	}
	ms.setP(0x7)
	if ms.flagP {
		t.Errorf("expected p=false, got p=true")
	}
}
