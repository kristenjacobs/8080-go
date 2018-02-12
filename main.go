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
	Trace = log.New(traceHandle, "TRACE: ", 0)
	Debug = log.New(debugHandle, "DEBUG: ", log.Lshortfile)
}

func main() {
	trace := flag.Bool("t", false, "enables instruction tracing")
	debug := flag.Bool("d", false, "enables debug")
	test := flag.Bool("test", false, "executes inbuilt instruction test rom")
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

	ioHandler := newIOHandler()

	var ms *machineState
	if *test {
		ms = newTestMachineState()
	} else {
		ms = newMachineState(ioHandler)
	}

	go func() {
		for ms.halt == false {
			step(ms)
		}
	}()

	ioHandler.run(ms)
}
