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
	spriteSizePixels   = 4
	numX               = 256
	numY               = 224
	windowWidthPixels  = numX * spriteSizePixels
	windowHeightPixels = numY * spriteSizePixels
	videoRamAddr       = 0x2400
	midpointX          = numX / 2
)

type IOHandler struct {
	numScreenRefreshes int64
	screenRefreshNS    int64
}

func newIOHandler() *IOHandler {
	return &IOHandler{}
}

func (io *IOHandler) Read(port uint8) uint8 {
	var value uint8 = 0
	//fmt.Printf("IOHandler Read: %d, value: 0x%02x\n", port, value)
	return value
}

func (io *IOHandler) Write(port uint8, value uint8) {
	//fmt.Printf("IOHandler Write: port: %d, value: 0x%02x\n", port, value)
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

func (io *IOHandler) draw(win *pixelgl.Window, imd *imdraw.IMDraw, x int, y int, ms *machineState) {
	// Read pixel value from memory...
	// TODO
	value := 0

	if value == 1 {
		x1 := 0.0 // float64(x * spriteSizePixels)
		y1 := 0.0 //float64(y * spriteSizePixels)
		x2 := x1 + spriteSizePixels
		y2 := y1 + spriteSizePixels
		imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
		imd.Rectangle(0)
	}
}

func (io *IOHandler) handleKeys(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeft) {
		fmt.Println("LEFT")
	}
	if win.Pressed(pixelgl.KeyRight) {
		fmt.Println("RIGHT")
	}
	if win.Pressed(pixelgl.KeyUp) {
		fmt.Println("UP")
	}
	if win.Pressed(pixelgl.KeyDown) {
		fmt.Println("DOWN")
	}
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

	// Starts the keyboard handling goroutine.
	go func() {
		for ms.halt == false {
			io.handleKeys(win)
		}
	}()

	// Starts the sound handling goroutine.
	// TODO

	// Graphics handling loop.
	imd := imdraw.New(nil)
	imd.Color = colornames.White

	for !win.Closed() {
		start := time.Now()

		win.Clear(colornames.Black)
		imd.Clear()

		if ms.halt == true {
			return
		}

		// Draw the left half of the screen, starting at bottom left.
		// (Note: First row ends at top left).
		//fmt.Printf("DRAWNING LEFT HALF\n")
		for x := 0; x < midpointX; x++ {
			for y := 0; y < numY; y++ {
				io.draw(win, imd, x, y, ms)
			}
		}

		if ms.interruptsEnabled {
			// Middle of frame interrupt (RST_1).
			ms.setInterrupt(0x08)
		}

		// Draw the right half of the screen, starting at bottom middle.
		// (Note: Last row ends at top right).
		//fmt.Printf("DRAWNING RIGHT HALF\n")
		for x := midpointX; x < numX; x++ {
			for y := 0; y < numY; y++ {
				io.draw(win, imd, x, y, ms)
			}
		}

		if ms.interruptsEnabled {
			// End of frame interrupt (RST_2).
			ms.setInterrupt(0x10)
		}

		imd.Draw(win)
		win.Update()

		t := time.Now()
		elapsed := t.Sub(start)

		io.numScreenRefreshes++
		io.screenRefreshNS += int64(elapsed)

		// TODO: Remove this!
		//if io.numScreenRefreshes == 2000 {
		//	ms.halt = true
		//}
	}
}
