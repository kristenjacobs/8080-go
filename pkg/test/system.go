package test

import (
	"github.com/kristenjacobs/8080-go/pkg/core"
)

const (
	RAM_SIZE      uint16 = 0x2000
	RAM_BASE      uint16 = 0x2000
	RAM_MIRROR    uint16 = 0x4000
	TEST_ROM_BASE uint16 = 0x100
	TEST_ROM_SIZE uint16 = 0x1000
)

type System struct {
	ms *core.MachineState
}

func NewSystem() *System {
	return &System{}
}

func (s *System) Run(max int64) {
	s.ms = core.NewMachineState(nil, TEST_ROM_BASE, RAM_BASE)

	// Configures the core ram.
	s.ms.InitialiseRam(RAM_BASE, RAM_SIZE)
	s.ms.InitialiseMirror(RAM_MIRROR)

	// Loads the test rom.
	s.ms.LoadRom(TEST_ROM_BASE, TEST_ROM_SIZE, TestRom)

	// Skips the DAA test
	s.ms.WriteMem(0x59c, []uint8{0xc3}, 1) // JMP
	s.ms.WriteMem(0x59d, []uint8{0xc2}, 1)
	s.ms.WriteMem(0x59e, []uint8{0x05}, 1)

	// Starts the 8080 core running.
	core.Run(s.ms, max)
}

func (s *System) DumpStats() {
	s.ms.DumpStats()
}
