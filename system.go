package main

import (
	"fmt"
	//"image/color"
	//"os/exec"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

const (
	spriteSizePixels          = 4
	numX                      = 256
	numY                      = 224
	windowWidthPixels         = numX * spriteSizePixels
	windowHeightPixels        = numY * spriteSizePixels
	videoRamAddr       uint16 = 0x2400
	midpointX                 = numX / 2
)

type System struct {
	window              *pixelgl.Window
	shiftRegister       uint16
	shiftRegisterOffset uint16
	numScreenRefreshes  int64
	screenRefreshNS     int64
	pixelsRendered      int64
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
		fmt.Printf("Unimplemented read from port: %d\n", port)

	case 1:
		//0   coin (0 when active)
		//1   P2 start button   (key '2')
		//2   P1 start button   (key '1')
		//3   ?
		//4   P1 shoot button   (key 'e')
		//5   P1 joystick left  (key 'q')
		//6   P1 joystick right (key 'w')
		//7   ?
		fmt.Printf("Read port: %d, value: 0x%02x\n", port, value)

	case 2:
		// 0,1 dipswitch number of lives (0:3,1:4,2:5,3:6)
		value = (value & ^uint8(0x3<<0)) | ((0 & 0x3) << 0)
		// 2   tilt 'button'     (key 'space')
		if system.window.Pressed(pixelgl.KeySpace) {
			value = (value & ^uint8(0x1<<2)) | ((1 & 0x1) << 2)
		}
		// 3   dipswitch bonus life at 1:1000,0:1500
		value = (value & ^uint8(0x1<<3)) | ((1 & 0x1) << 3)
		// 4   P2 shoot button   ('i')
		if system.window.Pressed(pixelgl.KeyI) {
			value = (value & ^uint8(0x1<<4)) | ((1 & 0x1) << 4)
		}
		// 5   P2 joystick left  ('o')
		if system.window.Pressed(pixelgl.KeyO) {
			value = (value & ^uint8(0x1<<5)) | ((1 & 0x1) << 5)
		}
		// 6   P2 joystick right ('p')
		if system.window.Pressed(pixelgl.KeyP) {
			value = (value & ^uint8(0x1<<6)) | ((1 & 0x1) << 6)
		}
		// 7   dipswitch coin info 1:off,0:on
		value = (value & ^uint8(0x1<<7)) | ((0 & 0x1) << 7)

	case 3:
		// shift register result
		value = uint8((system.shiftRegister >> (8 - system.shiftRegisterOffset)) & 0xFF)
		fmt.Printf("Read port: %d, value: 0x%02x\n", port, value)

	case 4:
		fmt.Printf("Unimplemented read from port: %d\n", port)

	case 5:
		fmt.Printf("Unimplemented read from port: %d\n", port)

	case 6:
		fmt.Printf("Unimplemented read from port: %d\n", port)

	case 7:
		fmt.Printf("Unimplemented read from port: %d\n", port)
	}
	return value
}

func (system *System) Write(port uint8, value uint8) {
	switch port {
	case 0:
		fmt.Printf("Unimplemented write to port: %d\n", port)

	case 1:
		fmt.Printf("Unimplemented write to port: %d\n", port)

	case 2:
		// shift register result offset (bits 0,1,2)
		fmt.Printf("Write: port: %d, value: 0x%02x\n", port, value)
		system.shiftRegisterOffset = uint16(value) & 0x7

	case 3:
		// sound related
		fmt.Printf("Unimplemented write to port: %d\n", port)

	case 4:
		// fill shift register
		fmt.Printf("Write: port: %d, value: 0x%02x\n", port, value)
		system.shiftRegister = (system.shiftRegister) >> 8
		system.shiftRegister = system.shiftRegister | ((uint16(value) << 8) & 0xFF00)

	case 5:
		// sound related
		fmt.Printf("Unimplemented write to port: %d\n", port)

	case 6:
		// 'debug' port? eg. it writes to this port when it
		// writes text to the screen (0=a,1=b,2=c, etc)
		fmt.Printf("Unimplemented write to port: %d\n", port)

	case 7:
		fmt.Printf("Unimplemented write to port: %d\n", port)
	}
}

//func (system *System) handleKeys(win *pixelgl.Window) {
//	if win.Pressed(pixelgl.KeyLeft) {
//		fmt.Println("LEFT")
//	}
//	if win.Pressed(pixelgl.KeyRight) {
//		fmt.Println("RIGHT")
//	}
//	if win.Pressed(pixelgl.KeyUp) {
//		fmt.Println("UP")
//	}
//	if win.Pressed(pixelgl.KeyDown) {
//		fmt.Println("DOWN")
//	}
//}

//func playSound() {
//	go func() {
//		cmd := exec.Command("paplay", "./resources/sample.wav")
//		_, err := cmd.CombinedOutput()
//		if err != nil {
//			panic(err)
//		}
//	}()
//}

func (system *System) draw(imd *imdraw.IMDraw, x int, y int) {
	x1 := float64(x * spriteSizePixels)
	y1 := float64(y * spriteSizePixels)
	x2 := x1 + spriteSizePixels
	y2 := y1 + spriteSizePixels
	imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
	imd.Rectangle(0)
	system.pixelsRendered++
}

func (system *System) renderScreen(imd *imdraw.IMDraw, ms *machineState, fromX int, toX int, byteIndex uint16) uint16 {
	var bitIndex uint = 0
	var byteValue uint8
	for x := fromX; x < toX; x++ {
		for y := 0; y < numY; y++ {
			if bitIndex == 0 {
				byteValue = ms.readMem(byteIndex, 1)[0]
				byteIndex++
			}
			if ((byteValue << bitIndex) & 0x1) == 0x1 {
				system.draw(imd, x, y)
			}
			bitIndex++
			if bitIndex == 8 {
				bitIndex = 0
			}
		}
	}
	return byteIndex
}

func (system *System) run(ms *machineState) {
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

	for !system.window.Closed() && ms.halt == false {
		start := time.Now()

		system.window.Clear(colornames.Black)
		imd.Clear()

		if ms.interruptsEnabled {
			var byteIndex uint16 = videoRamAddr

			// Draw the left half of the screen, starting at bottom left.
			// (Note: First row ends at top left).
			byteIndex = system.renderScreen(imd, ms, 0, midpointX, byteIndex)

			// Middle of frame interrupt (RST_1).
			ms.setInterrupt(0x08)

			// Draw the right half of the screen, starting at bottom middle.
			// (Note: Last row ends at top right).
			byteIndex = system.renderScreen(imd, ms, midpointX, numX, byteIndex)

			// End of frame interrupt (RST_2).
			ms.setInterrupt(0x10)

			imd.Draw(system.window)
			system.window.Update()

			system.numScreenRefreshes++
			system.screenRefreshNS += int64(time.Now().Sub(start))
		}
	}
}
