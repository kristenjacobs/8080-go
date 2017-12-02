package main

func main() {
	ms := newMachineState()
	for {
		fetchAndDecode(ms)
	}
}
