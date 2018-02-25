package main

import (
	"flag"

	"github.com/kristenjacobs/8080-go/pkg/core"
	"github.com/kristenjacobs/8080-go/pkg/test"
)

func main() {
	trace := flag.Bool("t", false, "enables instruction tracing")
	stats := flag.Bool("s", false, "enables statistcs output")
	flag.Parse()

	core.InitTracing(*trace)

	s := test.NewSystem()
	s.Run(0)
	if *stats {
		s.DumpStats()
	}
}
