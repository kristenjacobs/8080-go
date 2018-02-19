package system

import (
	//"fmt"
	//"image/color"
	//"os/exec"
	"time"

	"github.com/kristenjacobs/8080-go/pkg/core"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

const (
	ROM_SIZE      uint16 = 0x800
	ROM_E_BASE    uint16 = 0x1800
	ROM_F_BASE    uint16 = 0x1000
	ROM_G_BASE    uint16 = 0x0800
	ROM_H_BASE    uint16 = 0x0000
	RAM_SIZE      uint16 = 0x2000
	RAM_BASE      uint16 = 0x2000
	RAM_MIRROR    uint16 = 0x4000
	TEST_ROM_BASE uint16 = 0x100
	TEST_ROM_SIZE uint16 = 0x1000
)

const (
	spriteSizePixels          = 3
	numX                      = 224
	numY                      = 256
	windowWidthPixels         = numX * spriteSizePixels
	windowHeightPixels        = numY * spriteSizePixels
	videoRamAddr       uint16 = 0x2400
)

type System struct {
	window               *pixelgl.Window
	shiftRegister        uint16
	shiftRegisterOffset  uint16
	numScreenRefreshes   int64
	screenRefreshNS      int64
	screenRefreshSleepNS int64
}

func newSystem() *System {
	return &System{
		numScreenRefreshes: 0,
		screenRefreshNS:    0,
	}
}

func (system *System) Read(port uint8) uint8 {
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
		if system.window.Pressed(pixelgl.KeySpace) {
			value |= (0x1 << 4)
		}
		// bit 5 Left
		if system.window.Pressed(pixelgl.KeyLeft) {
			value |= (0x1 << 5)
		}
		// bit 6 Right
		if system.window.Pressed(pixelgl.KeyRight) {
			value |= (0x1 << 6)
		}
		// bit 7 ? tied to demux port 7 ?
		//fmt.Printf("Unimplemented read from port: %d\n", port)

	case 1:
		// bit 0 = CREDIT (1 if deposit)
		if system.window.Pressed(pixelgl.KeyEnter) {
			value |= 0x1
		}
		// bit 1 = 2P start (1 if pressed)
		if system.window.Pressed(pixelgl.Key2) {
			value |= 0x1 << 1
		}
		// bit 2 = 1P start (1 if pressed)
		if system.window.Pressed(pixelgl.Key1) {
			value |= 0x1 << 2
		}
		// bit 3 = Always 1
		value |= 0x1 << 3
		// bit 4 = 1P shot (1 if pressed)
		if system.window.Pressed(pixelgl.KeyE) {
			value |= 0x1 << 4
		}
		// bit 5 = 1P left (1 if pressed)
		if system.window.Pressed(pixelgl.KeyQ) {
			value |= 0x1 << 5
		}
		// bit 6 = 1P right (1 if pressed)
		if system.window.Pressed(pixelgl.KeyW) {
			value |= 0x1 << 6
		}
		// bit 7 = Not connected
		//fmt.Printf("Read port: %d, value: 0x%02x\n", port, value)

	case 2:

		// bit 0 = DIP3 00 = 3 ships  10 = 5 ships
		// bit 1 = DIP5 01 = 4 ships  11 = 6 ships
		value |= 0x0 << 0
		// bit 2 = Tilt
		if system.window.Pressed(pixelgl.KeyT) {
			value |= 0x1 << 2
		}
		// bit 3 = DIP6 0 = extra ship at 1500, 1 = extra ship at 1000
		value |= 0x0 << 3
		// bit 4 = P2 shot (1 if pressed)
		if system.window.Pressed(pixelgl.KeyI) {
			value |= 0x1 << 4
		}
		// bit 5 = P2 left (1 if pressed)
		if system.window.Pressed(pixelgl.KeyO) {
			value |= 0x1 << 5
		}
		// bit 6 = P2 right (1 if pressed)
		if system.window.Pressed(pixelgl.KeyP) {
			value |= 0x1 << 6
		}
		// bit 7 = DIP7 Coin info displayed in demo screen 0=ON
		value |= 0x0 << 7
		//fmt.Printf("Read port: %d, value: 0x%02x\n", port, value)

	case 3:
		// shift register result
		value = uint8((system.shiftRegister >> (8 - system.shiftRegisterOffset)) & 0xFF)
		//fmt.Printf("Read port: %d, value: 0x%02x, shiftRegister: 0x%04x, offset: %d\n",
		//	port, value, system.shiftRegister, system.shiftRegisterOffset)

	case 4:
		//fmt.Printf("Unimplemented read from port: %d\n", port)

	case 5:
		//fmt.Printf("Unimplemented read from port: %d\n", port)

	case 6:
		//fmt.Printf("Unimplemented read from port: %d\n", port)

	case 7:
		//fmt.Printf("Unimplemented read from port: %d\n", port)
	}
	return value
}

func (system *System) Write(port uint8, value uint8) {
	switch port {
	case 0:
		//fmt.Printf("Unimplemented write to port: %d\n", port)

	case 1:
		//fmt.Printf("Unimplemented write to port: %d\n", port)

	case 2:
		// shift register result offset (bits 0,1,2)
		//fmt.Printf("Write: port: %d, value: 0x%02x\n", port, value)
		system.shiftRegisterOffset = uint16(value) & 0x7

	case 3:
		// sound related
		//fmt.Printf("Unimplemented write to port: %d\n", port)

	case 4:
		// fill shift register
		//before := system.shiftRegister
		system.shiftRegister = (system.shiftRegister) >> 8
		system.shiftRegister = system.shiftRegister | ((uint16(value) << 8) & 0xFF00)
		//fmt.Printf("Write: port: %d, value: 0x%02x, befofe: 0x%04x, after: 0x%04x, offset: %d\n",
		//	port, value, before, system.shiftRegister, system.shiftRegisterOffset)

	case 5:
		// sound related
		//fmt.Printf("Unimplemented write to port: %d\n", port)

	case 6:
		// 'debug' port? eg. it writes to this port when it
		// writes text to the screen (0=a,1=b,2=c, etc)
		//fmt.Printf("Unimplemented write to port: %d\n", port)

	case 7:
		//fmt.Printf("Unimplemented write to port: %d\n", port)
	}
}

//func playSound() {
//	go func() {
//		cmd := exec.Command("paplay", "./resources/sample.wav")
//		_, err := cmd.CombinedOutput()
//		if err != nil {
//			panic(err)
//		}
//	}()
//}

func (system *System) drawPixel(imd *imdraw.IMDraw, x int, y int) {
	x1 := float64(x * spriteSizePixels)
	y1 := float64(y * spriteSizePixels)
	x2 := x1 + spriteSizePixels
	y2 := y1 + spriteSizePixels
	imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
	imd.Rectangle(0)
}

func (system *System) drawScreen(imd *imdraw.IMDraw, ms *core.MachineState, fromX int, toX int, byteIndex uint16) uint16 {
	var bitIndex uint = 0
	var byteValue uint8
	for x := fromX; x < toX; x++ {
		for y := 0; y < numY; y++ {
			if bitIndex == 0 {
				byteValue = ms.readMem(byteIndex, 1)[0]
				byteIndex++
			}
			if ((byteValue >> bitIndex) & 0x1) == 0x1 {
				system.drawPixel(imd, x, y)
			}
			bitIndex++
			if bitIndex == 8 {
				bitIndex = 0
			}
		}
	}
	return byteIndex
}

func (system *System) run(ms *core.MachineState) {
	// Loads the space invaders roms.
	ms.LoadRom(ROM_G_BASE, ROM_SIZE, InvadersG)
	ms.LoadRom(ROM_H_BASE, ROM_SIZE, InvadersH)
	ms.LoadRom(ROM_E_BASE, ROM_SIZE, InvadersE)
	ms.LoadRom(ROM_F_BASE, ROM_SIZE, InvadersF)

	cfg := pixelgl.WindowConfig{
		Title:  "Go Pixel Example",
		Bounds: pixel.R(0, 0, windowWidthPixels, windowHeightPixels),
	}
	var err error
	system.window, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Graphics handling loop.
	imd := imdraw.New(nil)
	imd.Color = colornames.White

	period := 8 * time.Millisecond

	for !system.window.Closed() && ms.halt == false {
		start := time.Now()

		system.window.Clear(colornames.Black)
		imd.Clear()

		var byteIndex uint16 = videoRamAddr

		// Draw the left half of the screen.
		byteIndex = system.drawScreen(imd, ms, 0, numX/2, byteIndex)

		// Middle of frame interrupt (RST_1).
		ms.setInterrupt(0x08)

		elapsed := time.Now().Sub(start)
		if elapsed < period {
			sleep := period - elapsed
			system.screenRefreshSleepNS += int64(sleep)
			time.Sleep(sleep)
		}
		system.screenRefreshNS += int64(elapsed)

		start = time.Now()

		// Draw the right half of the screen.
		byteIndex = system.drawScreen(imd, ms, numX/2, numX, byteIndex)

		// End of frame interrupt (RST_2).
		ms.setInterrupt(0x10)

		imd.Draw(system.window)
		system.window.Update()

		elapsed = time.Now().Sub(start)
		if elapsed < period {
			sleep := period - elapsed
			system.screenRefreshSleepNS += int64(sleep)
			time.Sleep(sleep)
		}
		system.screenRefreshNS += int64(elapsed)

		system.numScreenRefreshes++
	}
	ms.halt = true
}
