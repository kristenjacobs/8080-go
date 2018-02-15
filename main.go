package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/faiface/pixel/pixelgl"
)

var (
	Trace *log.Logger
	Debug *log.Logger
)

func initLogging(
	traceHandle io.Writer,
	debugHandle io.Writer) {
	Trace = log.New(traceHandle, "TRACE: ", 0)
	Debug = log.New(debugHandle, "DEBUG: ", log.Lshortfile)
}

func dumpStats(ms *machineState, ioHandler *IOHandler) {
	fmt.Printf("========== CORE STATS ==========\n")
	simulationTimeNS := int64(ms.endTime.Sub(ms.startTime))
	fmt.Printf("Simulation time: %.3fms\n", float64(simulationTimeNS/1000000.0))
	fmt.Printf("Instructions executed: %d\n", ms.numInstructionsExecuted)
	if ms.numInstructionsExecuted > 0 {
		fmt.Printf("Average time per instruction: %.3fms\n", float64(simulationTimeNS/ms.numInstructionsExecuted)/1000000.0)
	}
	fmt.Printf("\n")
	if ioHandler != nil {
		fmt.Printf("========== SYSTEM STATS ==========\n")
		fmt.Printf("Total screen refresh time: %.3fms\n", float64(ioHandler.screenRefreshNS/1000000.0))
		fmt.Printf("Number of screen refreshes: %d\n", ioHandler.numScreenRefreshes)
		if ioHandler.numScreenRefreshes > 0 {
			fmt.Printf("Average time per refresh: %.3fms\n", float64(ioHandler.screenRefreshNS/ioHandler.numScreenRefreshes)/1000000.0)
			fmt.Printf("Screen refresh rate: %.3f per sec\n", 1000000000.0/float64(ioHandler.screenRefreshNS/ioHandler.numScreenRefreshes))
			fmt.Printf("Total num pixels rendered: %d\n", ioHandler.pixelsRendered)
			fmt.Printf("Average num pixels rendered per frame: %.3f\n", float64(ioHandler.pixelsRendered/ioHandler.numScreenRefreshes))
		}
		fmt.Printf("\n")
	}
}

func main() {
	trace := flag.Bool("t", false, "enables instruction tracing")
	debug := flag.Bool("d", false, "enables debug")
	stats := flag.Bool("s", false, "enables statistcs output")
	test := flag.Bool("test", false, "executes inbuilt instruction test rom")
	max := flag.Int64("m", 0, "exit after `n` instructions")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()

	traceStream := ioutil.Discard
	if *trace {
		traceStream = os.Stdout
	}
	debugStream := ioutil.Discard
	if *debug {
		debugStream = os.Stdout
	}
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	initLogging(traceStream, debugStream)

	pixelgl.Run(func() {
		var ms *machineState
		var ioHandler *IOHandler
		if *test {
			ms = newTestMachineState()
			start(ms, *max)
		} else {
			ioHandler = newIOHandler()
			ms = newMachineState(ioHandler)
			go func() {
				start(ms, *max)
			}()
			ioHandler.run(ms)
		}
		if *stats {
			dumpStats(ms, ioHandler)
		}
	})

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
