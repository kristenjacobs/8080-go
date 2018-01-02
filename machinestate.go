package main

import (
	"fmt"
)

const (
	ROM_SIZE      uint16 = 0x800
	ROM_E_BASE    uint16 = 0x1800
	ROM_F_BASE    uint16 = 0x1000
	ROM_G_BASE    uint16 = 0x0800
	ROM_H_BASE    uint16 = 0x0000
	RAM_SIZE      uint16 = 0x2000
	RAM_BASE      uint16 = 0x2000
	TEST_ROM_BASE uint16 = 0x100
	TEST_ROM_SIZE uint16 = 0x1000
)

type memoryRegion struct {
	size  uint16
	base  uint16
	bytes []uint8
}

type machineState struct {
	roms []memoryRegion
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
	ms.initialiseRam()
	ms.initialiseSpaceInvadersRoms()
	ms.pc = ROM_H_BASE
	ms.sp = RAM_BASE
	ms.halt = false
	return &ms
}

func newTestMachineState() *machineState {
	ms := machineState{}
	ms.initialiseRam()
	ms.initialiseTestRom()
	ms.pc = TEST_ROM_BASE
	ms.sp = RAM_BASE
	ms.halt = false
	return &ms
}

func (ms *machineState) initialiseSpaceInvadersRoms() {
	ms.roms = []memoryRegion{
		memoryRegion{ROM_SIZE, ROM_E_BASE, InvadersE},
		memoryRegion{ROM_SIZE, ROM_F_BASE, InvadersF},
		memoryRegion{ROM_SIZE, ROM_G_BASE, InvadersG},
		memoryRegion{ROM_SIZE, ROM_H_BASE, InvadersH},
	}
}

func (ms *machineState) initialiseTestRom() {
	ms.roms = []memoryRegion{
		memoryRegion{TEST_ROM_SIZE, TEST_ROM_BASE, TestRom},
	}
}

func (ms *machineState) initialiseRam() {
	ms.ram.size = RAM_SIZE
	ms.ram.base = RAM_BASE
	ms.ram.bytes = make([]uint8, RAM_SIZE)
}

func (ms *machineState) readMem(addr uint16, numBytes uint16) []uint8 {
	for _, rom := range ms.roms {
		if inRegion(addr, numBytes, &rom) {
			return read(addr, numBytes, &rom)
		}
	}
	if inRegion(addr, numBytes, &ms.ram) {
		return read(addr, numBytes, &ms.ram)
	}
	panic(fmt.Sprintf("Cannot read memory, addr: 0x%04x, numBytes: %d", addr, numBytes))
}

func (ms *machineState) writeMem(addr uint16, bytes []uint8, numBytes uint16) {
	for _, rom := range ms.roms {
		if inRegion(addr, numBytes, &rom) {
			write(addr, bytes, numBytes, &rom)
			return
		}
	}
	if inRegion(addr, numBytes, &ms.ram) {
		write(addr, bytes, numBytes, &ms.ram)
		return
	}
	panic(fmt.Sprintf("Cannot write memory, addr: 0x%04x, numBytes: %d", addr, numBytes))
}

func inRegion(addr uint16, numBytes uint16, mr *memoryRegion) bool {
	result := (addr >= mr.base) && (addr+numBytes) <= (mr.base+mr.size)
	Debug.Printf("In region: 0x%04x, %d, region: 0x%04x->0x%04x, result: %t\n", addr, numBytes, mr.base, mr.base+mr.size, result)
	return result
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

func (ms *machineState) setCY(val bool) {
	ms.flagCY = val
}

func (ms *machineState) setAC(result uint8) {
	// Not yet implemented.
}

func getPair(regHi uint8, regLo uint8) uint16 {
	return (uint16(regHi) << 8) | uint16(regLo&0xFF)
}

func setPair(regHi *uint8, regLo *uint8, val uint16) {
	*regLo = uint8(val & 0xFF)
	*regHi = uint8((val >> 8) & 0xFF)
}
