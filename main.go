package main

import (
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
	InitLogging(os.Stdout, ioutil.Discard)
	//InitLogging(os.Stdout, os.Stdout)

	ms := newMachineState()
	for ms.halt == false {
		step(ms)
	}
}
