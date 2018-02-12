package main

import "testing"

func TestSetZ(t *testing.T) {
	ms := newMachineState(nil)
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
	ms := newMachineState(nil)
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
	ms := newMachineState(nil)
	ms.setP(0xF)
	if !ms.flagP {
		t.Errorf("expected p=true, got p=false")
	}
	ms.setP(0x7)
	if ms.flagP {
		t.Errorf("expected p=false, got p=true")
	}
}
