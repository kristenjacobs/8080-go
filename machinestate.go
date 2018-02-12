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

	halt              bool
	interruptsEnabled bool
	ioHandler         *IOHandler
}

func newMachineState(ioHandler *IOHandler) *machineState {
	ms := machineState{}
	ms.initialiseRam()
	ms.initialiseSpaceInvadersRoms()
	ms.initialiseFlags()
	ms.pc = ROM_H_BASE
	ms.sp = RAM_BASE
	ms.halt = false
	ms.ioHandler = ioHandler
	return &ms
}

func newTestMachineState() *machineState {
	ms := machineState{}
	ms.initialiseRam()
	ms.initialiseTestRom()
	ms.initialiseFlags()
	ms.pc = TEST_ROM_BASE
	ms.sp = RAM_BASE
	ms.halt = false
	ms.ioHandler = nil
	return &ms
}

func (ms *machineState) initialiseSpaceInvadersRoms() {
	ms.roms = []memoryRegion{
		memoryRegion{ROM_SIZE, ROM_E_BASE, newRomBytes(ROM_SIZE, InvadersE)},
		memoryRegion{ROM_SIZE, ROM_F_BASE, newRomBytes(ROM_SIZE, InvadersF)},
		memoryRegion{ROM_SIZE, ROM_G_BASE, newRomBytes(ROM_SIZE, InvadersG)},
		memoryRegion{ROM_SIZE, ROM_H_BASE, newRomBytes(ROM_SIZE, InvadersH)},
	}
}

func (ms *machineState) initialiseTestRom() {
	ms.roms = []memoryRegion{
		memoryRegion{TEST_ROM_SIZE, TEST_ROM_BASE, newRomBytes(TEST_ROM_SIZE, TestRom)},
	}

	// Skips the DAA test
	ms.writeMem(0x59c, []uint8{0xc3}, 1) // JMP
	ms.writeMem(0x59d, []uint8{0xc2}, 1)
	ms.writeMem(0x59e, []uint8{0x05}, 1)
}

func (ms *machineState) initialiseRam() {
	ms.ram.size = RAM_SIZE
	ms.ram.base = RAM_BASE
	ms.ram.bytes = make([]uint8, RAM_SIZE)
}

func (ms *machineState) initialiseFlags() {
	ms.flagZ = false
	ms.flagS = false
	ms.flagP = false
	ms.flagCY = false
	ms.flagAC = false
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

func (ms *machineState) getM() uint8 {
	return ms.readMem(getPair(ms.regH, ms.regL), 1)[0]
}

func (ms *machineState) setM(val uint8) {
	ms.writeMem(getPair(ms.regH, ms.regL), []uint8{val}, 1)
}

func (ms *machineState) getFlags() uint8 {
	var f uint8 = 0
	if ms.flagS {
		f |= 1 << 7
	}
	if ms.flagZ {
		f |= 1 << 6
	}
	if ms.flagAC {
		f |= 1 << 4
	}
	if ms.flagP {
		f |= 1 << 2
	}
	f |= 1 << 1
	if ms.flagCY {
		f |= 1
	}
	return f
}

func (ms *machineState) setFlags(val uint8) {
	ms.flagS = ((val >> 7) & 0x1) == 0x1
	ms.flagZ = ((val >> 6) & 0x1) == 0x1
	ms.flagAC = ((val >> 4) & 0x1) == 0x1
	ms.flagP = ((val >> 2) & 0x1) == 0x1
	ms.flagCY = (val & 0x1) == 0x1
}

func getPair(regHi uint8, regLo uint8) uint16 {
	return (uint16(regHi) << 8) | uint16(regLo&0xFF)
}

func setPair(regHi *uint8, regLo *uint8, val uint16) {
	*regLo = uint8(val & 0xFF)
	*regHi = uint8((val >> 8) & 0xFF)
}

func newRomBytes(romSize uint16, romBytes []uint8) []uint8 {
	bytes := make([]uint8, romSize)
	copy(bytes, romBytes)
	return bytes
}
