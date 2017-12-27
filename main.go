package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace *log.Logger
	Debug *log.Logger
)

func InitLogging(
	traceHandle io.Writer,
	debugHandle io.Writer) {
	Trace = log.New(traceHandle, "TRACE: ", log.Ltime)
	Debug = log.New(debugHandle, "DEBUG: ", log.Ltime|log.Lshortfile)
}

func main() {
	trace := flag.Bool("t", false, "enables instruction tracing")
	debug := flag.Bool("d", false, "enables debug")
	flag.Parse()

	traceStream := ioutil.Discard
	if *trace {
		traceStream = os.Stdout
	}
	debugStream := ioutil.Discard
	if *debug {
		debugStream = os.Stdout
	}
	InitLogging(traceStream, debugStream)

	ms := newMachineState()
	for ms.halt == false {
		step(ms)
	}
}
