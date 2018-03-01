package test

import (
	"github.com/kristenjacobs/8080-go/pkg/core"
)

const (
	ramBase     uint16 = 0x2000
	ramSize     uint16 = 0x2000
	ramMirror   uint16 = 0x4000
	testRomBase uint16 = 0x100
	testRomSize uint16 = 0x1000
)

type System struct {
	ms *core.MachineState
}

func NewSystem() *System {
	return &System{}
}

func (s *System) Run(max int64) {
	s.ms = core.NewMachineState(nil, testRomBase, ramBase)

	// Configures the core ram.
	s.ms.InitialiseRam(ramBase, ramSize)
	s.ms.InitialiseMirror(ramMirror)

	// Loads the test rom.
	s.ms.LoadRom(testRomBase, testRomSize, testRom)

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
