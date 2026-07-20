package main

import ("time"
		"process-engine/internal/runtime"
		)

func main() {
	run()
}

func run() {
	runtime := runtime.Start()

	for {
		runtime.Tick()
		time.Sleep(time.Second)
	}
}
