package core

import (
	"fmt"
	"time"
)

type memoryRegion struct {
	base  uint16
	size  uint16
	bytes []uint8
}

type IO interface {
	Read(port uint8) uint8
	Write(port uint8, value uint8)
}

type MachineState struct {
	roms      []memoryRegion
	ram       memoryRegion
	rammirror memoryRegion

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
	interrupt         bool
	interruptAddr     uint16
	io                IO

	startTime               time.Time
	endTime                 time.Time
	numInstructionsExecuted int64
	coreSleepNS             int64
}

func NewMachineState(io IO) *MachineState {
	ms := MachineState{}
	ms.initialiseRam()
	ms.initialiseRoms()
	ms.initialiseFlags()
	ms.pc = ROM_H_BASE
	ms.sp = RAM_BASE
	ms.halt = false
	ms.io = io
	return &ms
}

func newTestMachineState() *MachineState {
	ms := MachineState{}
	ms.initialiseRam()
	ms.initialiseTestRom()
	ms.initialiseFlags()
	ms.pc = TEST_ROM_BASE
	ms.sp = RAM_BASE
	ms.halt = false
	ms.io = nil
	return &ms
}

func (ms *MachineState) initialiseRoms() {
	ms.roms = []memoryRegion{}
}

func (ms *MachineState) initialiseTestRom() {
	ms.roms = []memoryRegion{
	//	memoryRegion{TEST_ROM_SIZE, TEST_ROM_BASE, newRomBytes(TEST_ROM_SIZE, TestRom)},
	}

	// Skips the DAA test
	//	ms.writeMem(0x59c, []uint8{0xc3}, 1) // JMP
	//	ms.writeMem(0x59d, []uint8{0xc2}, 1)
	//	ms.writeMem(0x59e, []uint8{0x05}, 1)
}

func (ms *MachineState) initialiseRam() {
	ms.ram.size = RAM_SIZE
	ms.ram.base = RAM_BASE
	ms.ram.bytes = make([]uint8, RAM_SIZE)

	ms.rammirror.size = RAM_SIZE
	ms.rammirror.base = RAM_MIRROR
	ms.rammirror.bytes = ms.ram.bytes
}

func (ms *MachineState) initialiseFlags() {
	ms.flagZ = false
	ms.flagS = false
	ms.flagP = false
	ms.flagCY = false
	ms.flagAC = false
}

func (ms *MachineState) LoadRom(base uint16, size uint16, bytes []uint8) {
	ms.roms = append(ms.roms, memoryRegion{size, base, newRomBytes(size, bytes)})
}

func (ms *MachineState) readMem(addr uint16, numBytes uint16) []uint8 {
	for _, rom := range ms.roms {
		if inRegion(addr, numBytes, &rom) {
			return read(addr, numBytes, &rom)
		}
	}
	if inRegion(addr, numBytes, &ms.ram) {
		return read(addr, numBytes, &ms.ram)
	}
	if inRegion(addr, numBytes, &ms.rammirror) {
		return read(addr, numBytes, &ms.rammirror)
	}
	panic(fmt.Sprintf("Cannot read memory, addr: 0x%04x, numBytes: %d", addr, numBytes))
}

func (ms *MachineState) writeMem(addr uint16, bytes []uint8, numBytes uint16) {
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
	if inRegion(addr, numBytes, &ms.rammirror) {
		write(addr, bytes, numBytes, &ms.rammirror)
		return
	}
	panic(fmt.Sprintf("Cannot write memory, addr: 0x%04x, numBytes: %d", addr, numBytes))
}

func inRegion(addr uint16, numBytes uint16, mr *memoryRegion) bool {
	result := (addr >= mr.base) && (addr+numBytes) <= (mr.base+mr.size)
	//Debug.Printf("In region: 0x%04x, %d, region: 0x%04x->0x%04x, result: %t\n", addr, numBytes, mr.base, mr.base+mr.size, result)
	return result
}

func read(addr uint16, numBytes uint16, mr *memoryRegion) []uint8 {
	i := addr - mr.base
	bytes := mr.bytes[i : i+numBytes]
	//Debug.Printf("read %d bytes at addr: 0x%04x: %v\n", numBytes, addr, bytes)
	return bytes
}

func write(addr uint16, bytes []uint8, numBytes uint16, mr *memoryRegion) {
	a := addr - mr.base
	for i, b := range bytes {
		mr.bytes[a+uint16(i)] = b
	}
	//Debug.Printf("wrote %d bytes at addr: 0x%04x: %v\n", numBytes, addr, bytes)
}

func (ms *MachineState) setZ(result uint8) {
	ms.flagZ = result == 0
}

func (ms *MachineState) setS(result uint8) {
	ms.flagS = (result >> 7) == 0x1
}

func (ms *MachineState) setP(result uint8) {
	numBitsSet := 0
	for i := uint(0); i < 8; i++ {
		if ((result >> i) & 0x1) == 0x1 {
			numBitsSet++
		}
	}
	ms.flagP = (numBitsSet & 0x1) == 0
}

func (ms *MachineState) setCY(val bool) {
	ms.flagCY = val
}

func (ms *MachineState) setAC(result uint8) {
	// Not yet implemented.
}

func (ms *MachineState) getM() uint8 {
	return ms.readMem(getPair(ms.regH, ms.regL), 1)[0]
}

func (ms *MachineState) setM(val uint8) {
	ms.writeMem(getPair(ms.regH, ms.regL), []uint8{val}, 1)
}

func (ms *MachineState) getFlags() uint8 {
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

func (ms *MachineState) setFlags(val uint8) {
	ms.flagS = ((val >> 7) & 0x1) == 0x1
	ms.flagZ = ((val >> 6) & 0x1) == 0x1
	ms.flagAC = ((val >> 4) & 0x1) == 0x1
	ms.flagP = ((val >> 2) & 0x1) == 0x1
	ms.flagCY = (val & 0x1) == 0x1
}

func (ms *MachineState) setInterrupt(addr uint16) {
	ms.interrupt = true
	ms.interruptAddr = addr
}

func (ms *MachineState) handleInterrupt() bool {
	if ms.interruptsEnabled && ms.interrupt {
		nextPC := ms.pc
		pcHi := uint8(nextPC >> 8)
		pcLo := uint8(nextPC & 0xFF)
		ms.writeMem(ms.sp-2, []uint8{pcLo, pcHi}, 2)
		ms.sp = ms.sp - 2
		ms.pc = ms.interruptAddr
		Trace.Printf("********** INTERRUPT: addr: 0x%04x **********\n", ms.interruptAddr)
		ms.interrupt = false
		ms.interruptAddr = 0
		ms.interruptsEnabled = false
		return true
	}
	//ms.interrupt = false
	//ms.interruptAddr = 0
	return false
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
