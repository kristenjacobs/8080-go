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

func dumpStats(ms *machineState) {
	simulationTime := ms.endTime.Sub(ms.startTime)
	fmt.Printf("Simulation time: %v\n", simulationTime)
	fmt.Printf("Instructions executed: %d\n", ms.numInstructionsExecuted)
	fmt.Printf("Average time per instruction: %.3fus\n", float64(int64(simulationTime)/ms.numInstructionsExecuted)/1000.0)
}

func run(stats bool, test bool, max int64) {
	var ms *machineState
	if test {
		ms = newTestMachineState()
		start(ms, max)
	} else {
		ioHandler := newIOHandler()
		ms = newMachineState(ioHandler)
		go func() {
			start(ms, max)
		}()
		ioHandler.run(ms)
	}
	if stats {
		dumpStats(ms)
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
		run(*stats, *test, *max)
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
