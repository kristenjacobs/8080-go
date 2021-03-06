package system

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/kristenjacobs/8080-go/pkg/core"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

const (
	romEBase           uint16 = 0x1800
	romFBase           uint16 = 0x1000
	romGBase           uint16 = 0x0800
	romHBase           uint16 = 0x0000
	romSize            uint16 = 0x800
	ramBase            uint16 = 0x2000
	ramSize            uint16 = 0x2000
	ramMirror          uint16 = 0x4000
	videoRamAddr       uint16 = 0x2400
	spriteSizePixels          = 3
	numX                      = 224
	numY                      = 256
	windowWidthPixels         = numX * spriteSizePixels
	windowHeightPixels        = numY * spriteSizePixels
)

type System struct {
	window               *pixelgl.Window
	ms                   *core.MachineState
	shiftRegister        uint16
	shiftRegisterOffset  uint16
	numScreenRefreshes   int64
	screenRefreshNS      int64
	screenRefreshSleepNS int64
	soundEffects1        uint8
	soundEffects2        uint8
}

func NewSystem() *System {
	return &System{
		numScreenRefreshes: 0,
		screenRefreshNS:    0,
		soundEffects1:      0,
		soundEffects2:      0,
	}
}

func (s *System) Run(max int64) {
	s.ms = core.NewMachineState(s, romHBase, ramBase)

	// Configures the core ram.
	s.ms.InitialiseRam(ramBase, ramSize)
	s.ms.InitialiseMirror(ramMirror)

	// Loads the space invaders roms.
	s.ms.LoadRom(romGBase, romSize, invadersG)
	s.ms.LoadRom(romHBase, romSize, invadersH)
	s.ms.LoadRom(romEBase, romSize, invadersE)
	s.ms.LoadRom(romFBase, romSize, invadersF)

	// Starts the 8080 core running.
	go func() {
		core.Run(s.ms, max)
	}()

	// Handles the screen updates on the main thread.
	s.handleScreen()
}

func (s *System) DumpStats() {
	s.ms.DumpStats()
	fmt.Printf("========== SYSTEM STATS ==========\n")
	fmt.Printf("Total screen refresh time: %.3fms\n", float64(s.screenRefreshNS/1000000.0))
	fmt.Printf("Total screen refresh sleep time: %.3fms\n", float64(s.screenRefreshSleepNS/1000000.0))
	fmt.Printf("Number of screen refreshes: %d\n", s.numScreenRefreshes)
	if s.numScreenRefreshes > 0 {
		fmt.Printf("Average time per refresh: %.3fms\n", float64(s.screenRefreshNS/s.numScreenRefreshes)/1000000.0)
		fmt.Printf("Average time per refresh sleep: %.3fms\n", float64(s.screenRefreshSleepNS/s.numScreenRefreshes)/1000000.0)
		fmt.Printf("Max screen refresh rate: %.3f per sec\n", 1000000000.0/float64(s.screenRefreshNS/s.numScreenRefreshes))
		fmt.Printf("Actual screen refresh rate: %.3f per sec\n", 1000000000.0/float64((s.screenRefreshNS+s.screenRefreshSleepNS)/s.numScreenRefreshes))
	}
	fmt.Printf("\n")
}

func (s *System) Read(port uint8) uint8 {
	var value uint8 = 0
	switch port {
	case 0:
		// bit 0 DIP4 (Seems to be self-test-request read at power up)
		value |= 0x1
		// bit 1 Always 1
		value |= 0x1
		// bit 2 Always 1
		value |= 0x1
		// bit 3 Always 1
		value |= 0x1
		// bit 4 Fire
		if s.window.Pressed(pixelgl.KeySpace) {
			value |= (0x1 << 4)
		}
		// bit 5 Left
		if s.window.Pressed(pixelgl.KeyLeft) {
			value |= (0x1 << 5)
		}
		// bit 6 Right
		if s.window.Pressed(pixelgl.KeyRight) {
			value |= (0x1 << 6)
		}
		// bit 7 ? tied to demux port 7 ?

	case 1:
		// bit 0 = CREDIT (1 if deposit)
		if s.window.Pressed(pixelgl.KeyEnter) {
			value |= 0x1
		}
		// bit 1 = 2P start (1 if pressed)
		if s.window.Pressed(pixelgl.Key2) {
			value |= 0x1 << 1
		}
		// bit 2 = 1P start (1 if pressed)
		if s.window.Pressed(pixelgl.Key1) {
			value |= 0x1 << 2
		}
		// bit 3 = Always 1
		value |= 0x1 << 3
		// bit 4 = 1P shot (1 if pressed)
		if s.window.Pressed(pixelgl.KeyE) {
			value |= 0x1 << 4
		}
		// bit 5 = 1P left (1 if pressed)
		if s.window.Pressed(pixelgl.KeyQ) {
			value |= 0x1 << 5
		}
		// bit 6 = 1P right (1 if pressed)
		if s.window.Pressed(pixelgl.KeyW) {
			value |= 0x1 << 6
		}
		// bit 7 = Not connected

	case 2:
		// bit 0 = DIP3 00 = 3 ships  10 = 5 ships
		// bit 1 = DIP5 01 = 4 ships  11 = 6 ships
		value |= 0x0 << 0
		// bit 2 = Tilt
		if s.window.Pressed(pixelgl.KeyT) {
			value |= 0x1 << 2
		}
		// bit 3 = DIP6 0 = extra ship at 1500, 1 = extra ship at 1000
		value |= 0x0 << 3
		// bit 4 = P2 shot (1 if pressed)
		if s.window.Pressed(pixelgl.KeyI) {
			value |= 0x1 << 4
		}
		// bit 5 = P2 left (1 if pressed)
		if s.window.Pressed(pixelgl.KeyO) {
			value |= 0x1 << 5
		}
		// bit 6 = P2 right (1 if pressed)
		if s.window.Pressed(pixelgl.KeyP) {
			value |= 0x1 << 6
		}
		// bit 7 = DIP7 Coin info displayed in demo screen 0=ON
		value |= 0x0 << 7

	case 3:
		// shift register result
		value = uint8((s.shiftRegister >> (8 - s.shiftRegisterOffset)) & 0xFF)
	}
	return value
}

func (s *System) Write(port uint8, value uint8) {
	switch port {
	case 2:
		// shift register result offset (bits 0,1,2)
		s.shiftRegisterOffset = uint16(value) & 0x7

	case 3:
		// sound related
		handleSoundEffect(s.soundEffects1, value, 1, "./res/shoot.wav")
		handleSoundEffect(s.soundEffects1, value, 2, "./res/explosion.wav")
		handleSoundEffect(s.soundEffects1, value, 3, "./res/invaderkilled.wav")
		s.soundEffects1 = value

	case 4:
		// fill shift register
		s.shiftRegister = (s.shiftRegister) >> 8
		s.shiftRegister = s.shiftRegister | ((uint16(value) << 8) & 0xFF00)

	case 5:
		// sound related
		handleSoundEffect(s.soundEffects2, value, 0, "./res/fastinvader1.wav")
		handleSoundEffect(s.soundEffects2, value, 1, "./res/fastinvader2.wav")
		handleSoundEffect(s.soundEffects2, value, 2, "./res/fastinvader3.wav")
		handleSoundEffect(s.soundEffects2, value, 3, "./res/fastinvader4.wav")
		handleSoundEffect(s.soundEffects2, value, 4, "./res/ufo_lowpitch.wav")
		s.soundEffects2 = value
	}
}

func (s *System) handleScreen() {
	cfg := pixelgl.WindowConfig{
		Title:  "8080 Space Invaders",
		Bounds: pixel.R(0, 0, windowWidthPixels, windowHeightPixels),
	}
	var err error
	s.window, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Graphics handling loop.
	imd := imdraw.New(nil)
	imd.Color = colornames.White

	period := 8 * time.Millisecond

	for !s.window.Closed() && s.ms.Halt == false {
		start := time.Now()

		s.window.Clear(colornames.Black)
		imd.Clear()

		var byteIndex uint16 = videoRamAddr

		// Draw the left half of the screen.
		byteIndex = s.drawScreen(imd, s.ms, 0, numX/2, byteIndex)

		// Middle of frame interrupt (RST_1).
		s.ms.SetInterrupt(0x08)

		elapsed := time.Now().Sub(start)
		if elapsed < period {
			sleep := period - elapsed
			s.screenRefreshSleepNS += int64(sleep)
			time.Sleep(sleep)
		}
		s.screenRefreshNS += int64(elapsed)

		start = time.Now()

		// Draw the right half of the screen.
		byteIndex = s.drawScreen(imd, s.ms, numX/2, numX, byteIndex)

		// End of frame interrupt (RST_2).
		s.ms.SetInterrupt(0x10)

		imd.Draw(s.window)
		s.window.Update()

		elapsed = time.Now().Sub(start)
		if elapsed < period {
			sleep := period - elapsed
			s.screenRefreshSleepNS += int64(sleep)
			time.Sleep(sleep)
		}
		s.screenRefreshNS += int64(elapsed)

		s.numScreenRefreshes++
	}
}

func (s *System) drawScreen(imd *imdraw.IMDraw, ms *core.MachineState, fromX int, toX int, byteIndex uint16) uint16 {
	var bitIndex uint = 0
	var byteValue uint8
	for x := fromX; x < toX; x++ {
		for y := 0; y < numY; y++ {
			if bitIndex == 0 {
				byteValue = ms.ReadMem(byteIndex, 1)[0]
				byteIndex++
			}
			if ((byteValue >> bitIndex) & 0x1) == 0x1 {
				s.drawPixel(imd, x, y)
			}
			bitIndex++
			if bitIndex == 8 {
				bitIndex = 0
			}
		}
	}
	return byteIndex
}

func (s *System) drawPixel(imd *imdraw.IMDraw, x int, y int) {
	x1 := float64(x * spriteSizePixels)
	y1 := float64(y * spriteSizePixels)
	x2 := x1 + spriteSizePixels
	y2 := y1 + spriteSizePixels
	imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
	imd.Rectangle(0)
}

func handleSoundEffect(previous uint8, value uint8, bitPos uint, file string) {
	// If the required bit has transitioned from 0->1,
	// play the corresponding sound effect.
	if (((previous >> bitPos) & 0x1) == 0x0) && (((value >> bitPos) & 0x1) == 0x1) {
		go func() {
			var cmd *exec.Cmd
			if runtime.GOOS == "linux" {
				cmd = exec.Command("paplay", file)
			} else if runtime.GOOS == "darwin" {
				cmd = exec.Command("afplay", file)
			} else {
				panic(fmt.Sprintf("runtime.GOOS: %s not supported", runtime.GOOS))
			}
			_, err := cmd.CombinedOutput()
			if err != nil {
				panic(err)
			}
		}()
	}
}
