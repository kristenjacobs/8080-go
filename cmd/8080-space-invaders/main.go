package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/kristenjacobs/8080-go/pkg/core"
	"github.com/kristenjacobs/8080-go/pkg/system"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	trace := flag.Bool("t", false, "enables instruction tracing")
	stats := flag.Bool("s", false, "enables statistcs output")
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

	sys := system.NewSystem()
	pixelgl.Run(func() {
		sys.Run(*max)
	})
	if *stats {
		sys.DumpStats()
	}

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
