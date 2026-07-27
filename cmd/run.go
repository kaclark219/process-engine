package main

import ("fmt"
		"os"
		"path/filepath"
		"process-engine/internal/runtime"
		"process-engine/internal/server"
		"time")

func main() {
	exePath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("Failed to get executable path: %v", err))
	}
	projectRoot := filepath.Dir(filepath.Dir(exePath))
	rulesDir := filepath.Join(projectRoot, "rules")

	srv := server.NewServer(rulesDir)
	rt := runtime.Start(rulesDir, srv)
	srv.SetAgentRegistrar(rt)

	go func() {
		fmt.Println("Starting server on http://localhost:8080")
		if err := srv.Start(":8080"); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	time.Sleep(time.Second*10)

	for {
		rt.Tick()
		time.Sleep(time.Second)
	}
}
