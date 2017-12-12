package main

import (
	"fmt"
)

const (
	ROM_SIZE   uint16 = 0x800
	ROM_E_BASE uint16 = 0x1800
	ROM_F_BASE uint16 = 0x1000
	ROM_G_BASE uint16 = 0x0800
	ROM_H_BASE uint16 = 0x0000
	RAM_SIZE   uint16 = 0x800
	RAM_BASE   uint16 = 0x2000
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

	halt bool
}

func newMachineState() *machineState {
	ms := machineState{}

	ms.initialiseRoms()
	ms.initialiseRam()

	ms.pc = ROM_H_BASE
	ms.sp = RAM_BASE

	ms.halt = false
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
	ms.ram.base = RAM_BASE
	ms.ram.bytes = make([]uint8, RAM_SIZE)
}

func (ms *machineState) readMem(addr uint16, numBytes uint16) []uint8 {
	if inRegion(addr, numBytes, &ms.romE) {
		return read(addr, numBytes, &ms.romE)
	}
	if inRegion(addr, numBytes, &ms.romF) {
		return read(addr, numBytes, &ms.romF)
	}
	if inRegion(addr, numBytes, &ms.romG) {
		return read(addr, numBytes, &ms.romG)
	}
	if inRegion(addr, numBytes, &ms.romH) {
		return read(addr, numBytes, &ms.romH)
	}
	if inRegion(addr, numBytes, &ms.ram) {
		return read(addr, numBytes, &ms.ram)
	}
	panic(fmt.Sprintf("Cannot read memory, addr: 0x%04x, numBytes: %d", addr, numBytes))
}

func (ms *machineState) writeMem(addr uint16, bytes []uint8, numBytes uint16) {
	if inRegion(addr, numBytes, &ms.romE) {
		write(addr, bytes, numBytes, &ms.romE)
		return
	}
	if inRegion(addr, numBytes, &ms.romF) {
		write(addr, bytes, numBytes, &ms.romF)
		return
	}
	if inRegion(addr, numBytes, &ms.romG) {
		write(addr, bytes, numBytes, &ms.romG)
		return
	}
	if inRegion(addr, numBytes, &ms.romH) {
		write(addr, bytes, numBytes, &ms.romH)
		return
	}
	if inRegion(addr, numBytes, &ms.ram) {
		write(addr, bytes, numBytes, &ms.ram)
		return
	}
	panic(fmt.Sprintf("Cannot write memory, addr: 0x%04x, numBytes: %d", addr, numBytes))
}

func inRegion(addr uint16, numBytes uint16, mr *memoryRegion) bool {
	//fmt.Printf("In region: 0x%04x, %d, region: 0x%04x, %d\n", addr, numBytes, mr.base, mr.size)
	return (addr >= mr.base) && (addr+numBytes) < (mr.base+mr.size)
}

func read(addr uint16, numBytes uint16, mr *memoryRegion) []uint8 {
	i := addr - mr.base
	bytes := mr.bytes[i : i+numBytes]
	Debug.Printf("read %d bytes at addr: 0x%04x: %v\n", numBytes, addr, bytes)
	return bytes
}

func write(addr uint16, bytes []uint8, numBytes uint16, mr *memoryRegion) {
	a := addr - mr.base
	for i, b := range bytes {
		mr.bytes[a+uint16(i)] = b
	}
	Debug.Printf("wrote %d bytes at addr: 0x%04x: %v\n", numBytes, addr, bytes)
}

func (ms *machineState) setZ(result uint8) {
	ms.flagZ = result == 0
}

func (ms *machineState) setS(result uint8) {
	ms.flagS = (result >> 7) == 0x1
}

func (ms *machineState) setP(result uint8) {
	numBitsSet := 0
	for i := uint(0); i < 8; i++ {
		if ((result >> i) & 0x1) == 0x1 {
			numBitsSet++
		}
	}
	ms.flagP = (numBitsSet & 0x1) == 0
}

func (ms *machineState) setAC(result uint8) {
	// Not yet implemented.
}

func (ms *machineState) addr(regLo uint8, regHi uint8) uint16 {
	return (uint16(regHi) << 8) | uint16(regLo&0xFF)
}
