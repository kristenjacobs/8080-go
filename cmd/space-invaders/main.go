package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/kristenjacobs/8080-go/pkg/core"

	"github.com/faiface/pixel/pixelgl"
)

func dumpStats(ms *core.MachineState, system *System) {
	fmt.Printf("========== CORE STATS ==========\n")
	simulationTimeNS := int64(ms.endTime.Sub(ms.startTime))
	fmt.Printf("Total simulation time: %.3fms\n", float64(simulationTimeNS/1000000.0))
	fmt.Printf("Total core sleep time: %.3fms\n", float64(ms.coreSleepNS/1000000.0))
	fmt.Printf("Instructions executed: %d\n", ms.numInstructionsExecuted)
	if ms.numInstructionsExecuted > 0 {
		fmt.Printf("Average time per instruction: %.3fus\n", float64(simulationTimeNS/ms.numInstructionsExecuted)/1000.0)
		fmt.Printf("Average sleep time per instruction: %.3fus\n", float64(ms.coreSleepNS/ms.numInstructionsExecuted)/1000.0)
	}
	fmt.Printf("\n")
	if system != nil {
		fmt.Printf("========== SYSTEM STATS ==========\n")
		fmt.Printf("Total screen refresh time: %.3fms\n", float64(system.screenRefreshNS/1000000.0))
		fmt.Printf("Total screen refresh sleep time: %.3fms\n", float64(system.screenRefreshSleepNS/1000000.0))
		fmt.Printf("Number of screen refreshes: %d\n", system.numScreenRefreshes)
		if system.numScreenRefreshes > 0 {
			fmt.Printf("Average time per refresh: %.3fms\n", float64(system.screenRefreshNS/system.numScreenRefreshes)/1000000.0)
			fmt.Printf("Average time per refresh sleep: %.3fms\n", float64(system.screenRefreshSleepNS/system.numScreenRefreshes)/1000000.0)
			fmt.Printf("Max screen refresh rate: %.3f per sec\n", 1000000000.0/float64(system.screenRefreshNS/system.numScreenRefreshes))
			fmt.Printf("Actual screen refresh rate: %.3f per sec\n", 1000000000.0/float64((system.screenRefreshNS+system.screenRefreshSleepNS)/system.numScreenRefreshes))
		}
		fmt.Printf("\n")
	}
}

func main() {
	trace := flag.Bool("t", false, "enables instruction tracing")
	stats := flag.Bool("s", false, "enables statistcs output")
	test := flag.Bool("test", false, "executes inbuilt instruction test rom")
	max := flag.Int64("m", 0, "exit after `n` instructions")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()

	core.InitTracing(*trace)

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

	pixelgl.Run(func() {
		var ms *core.MachineState
		var system *System
		if *test {
			ms = core.NewTestMachineState()
			start(ms, *max)
		} else {
			system = newSystem()
			ms = core.NewMachineState(system)
			go func() {
				start(ms, *max)
			}()
			system.run(ms)
		}
		if *stats {
			dumpStats(ms, system)
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
