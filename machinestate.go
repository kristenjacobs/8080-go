package main

import (
	"fmt"
)

const (
	RAM_SIZE uint16 = 0x800
	RAM_BASE uint16 = 0x2000
)

type memoryRegion struct {
	size  uint16
	base  uint16
	bytes []uint8
}

type machineState struct {
	romE memoryRegion
	romF memoryRegion
	romG memoryRegion
	romH memoryRegion
	ram  memoryRegion

	regA uint8
	regB uint8
	regC uint8
	regD uint8
	regE uint8
	regH uint8
	regL uint8

	flagZ  bool
	flagS  bool
	flagP  bool
	flagCY bool
	flagAC bool

	pc uint16
	sp uint16
}

func newMachineState() *machineState {
	ms := machineState{}

	ms.initialiseRoms()
	ms.initialiseRam()

	ms.pc = ROM_E_BASE
	ms.sp = RAM_BASE

	return &ms
}

func (ms *machineState) initialiseRoms() {
	ms.romE.size = ROM_SIZE
	ms.romE.base = ROM_E_BASE
	ms.romE.bytes = InvadersE

	ms.romF.size = ROM_SIZE
	ms.romF.base = ROM_F_BASE
	ms.romF.bytes = InvadersF

	ms.romG.size = ROM_SIZE
	ms.romG.base = ROM_G_BASE
	ms.romG.bytes = InvadersG

	ms.romH.size = ROM_SIZE
	ms.romH.base = ROM_H_BASE
	ms.romH.bytes = InvadersH
}

func (ms *machineState) initialiseRam() {
	ms.ram.size = RAM_SIZE
	ms.ram.bytes = make([]uint8, RAM_SIZE)
}

func (ms *machineState) readMem(addr uint16, numBytes uint16) ([]uint8, error) {
	if inRegion(addr, numBytes, &ms.romE) {
		fmt.Printf("inRegion E\n")
		return read(addr, numBytes, &ms.romE), nil
	}
	if inRegion(addr, numBytes, &ms.romF) {
		fmt.Printf("inRegion F\n")
		return read(addr, numBytes, &ms.romF), nil
	}
	if inRegion(addr, numBytes, &ms.romG) {
		fmt.Printf("inRegion G\n")
		return read(addr, numBytes, &ms.romG), nil
	}
	if inRegion(addr, numBytes, &ms.romH) {
		fmt.Printf("inRegion H\n")
		return read(addr, numBytes, &ms.romH), nil
	}
	if inRegion(addr, numBytes, &ms.ram) {
		fmt.Printf("inRegion Ram\n")
		return read(addr, numBytes, &ms.ram), nil
	}
	return nil, fmt.Errorf("Cannot read memory, addr: 0x%04x, numBytes: %d", addr, numBytes)
}

func inRegion(addr uint16, numBytes uint16, mr *memoryRegion) bool {
	return addr >= mr.base && (addr+numBytes) < (mr.base+mr.size)
}

func read(addr uint16, numBytes uint16, mr *memoryRegion) []uint8 {
	i := addr - mr.base
	return mr.bytes[i : i+numBytes]
}
