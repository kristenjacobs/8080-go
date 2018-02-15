package main

import (
	//"fmt"
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

type IOHandler struct {
	numScreenRefreshes int64
	screenRefreshNS    int64
	pixelsRendered     int64
}

func newIOHandler() *IOHandler {
	return &IOHandler{
		numScreenRefreshes: 0,
		screenRefreshNS:    0,
	}
}

func (io *IOHandler) Read(port uint8) uint8 {
	var value uint8 = 0
	//fmt.Printf("IOHandler Read: %d, value: 0x%02x\n", port, value)
	return value
}

func (io *IOHandler) Write(port uint8, value uint8) {
	//fmt.Printf("IOHandler Write: port: %d, value: 0x%02x\n", port, value)
}

//func (io *IOHandler) handleKeys(win *pixelgl.Window) {
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

func (io *IOHandler) draw(imd *imdraw.IMDraw, x int, y int) {
	x1 := float64(x * spriteSizePixels)
	y1 := float64(y * spriteSizePixels)
	x2 := x1 + spriteSizePixels
	y2 := y1 + spriteSizePixels
	imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
	imd.Rectangle(0)
	io.pixelsRendered++
}

func (io *IOHandler) renderScreen(imd *imdraw.IMDraw, ms *machineState, fromX int, toX int, byteIndex uint16) uint16 {
	var bitIndex uint = 0
	var byteValue uint8
	for x := fromX; x < toX; x++ {
		for y := 0; y < numY; y++ {
			if bitIndex == 0 {
				byteValue = ms.readMem(byteIndex, 1)[0]
				byteIndex++
			}
			if ((byteValue << bitIndex) & 0x1) == 0x1 {
				io.draw(imd, x, y)
			}
			bitIndex++
			if bitIndex == 8 {
				bitIndex = 0
			}
		}
	}
	return byteIndex
}

func (io *IOHandler) run(ms *machineState) {
	cfg := pixelgl.WindowConfig{
		Title:  "Go Pixel Example",
		Bounds: pixel.R(0, 0, windowWidthPixels, windowHeightPixels),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Graphics handling loop.
	imd := imdraw.New(nil)
	imd.Color = colornames.White

	for !win.Closed() && ms.halt == false {
		start := time.Now()

		win.Clear(colornames.Black)
		imd.Clear()

		if ms.interruptsEnabled {
			var byteIndex uint16 = videoRamAddr

			// Draw the left half of the screen, starting at bottom left.
			// (Note: First row ends at top left).
			byteIndex = io.renderScreen(imd, ms, 0, midpointX, byteIndex)

			// Middle of frame interrupt (RST_1).
			ms.setInterrupt(0x08)

			// Draw the right half of the screen, starting at bottom middle.
			// (Note: Last row ends at top right).
			byteIndex = io.renderScreen(imd, ms, midpointX, numX, byteIndex)

			// End of frame interrupt (RST_2).
			ms.setInterrupt(0x10)

			imd.Draw(win)
			win.Update()

			io.numScreenRefreshes++
			io.screenRefreshNS += int64(time.Now().Sub(start))
		}
	}
}
