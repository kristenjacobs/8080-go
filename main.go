package main

func main() {
	ms := newMachineState()
	for {
		step(ms)
	}
}
