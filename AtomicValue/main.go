package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type tConfigData struct {
	number int
	level  int
	name   string
	code   string
}

func loadConfig() tConfigData {
	return tConfigData{4, 12, "kelly", "3.14159"}
}

func requests() []tConfigData {
	return []tConfigData{
		{4, 42, "jim", "3.14159"},
		{1, 11, "allie", "3.14159"},
		{7, 72, "beanie", "3.14159"},
	}
}

func main() {
	var config atomic.Value // holds current server configuration
	// Create initial config value and store into config.
	config.Store(loadConfig())
	go func() {
		// Reload config every 10 seconds
		// and update config value with the new version.
		for {
			time.Sleep(10 * time.Millisecond)
			config.Store(loadConfig())
		}
	}()
	// Create worker goroutines that handle incoming requests
	// using the latest config value.
	for i := 0; i < 10; i++ {
		go func(ix int) {
			for _, r := range requests() {
				c := config.Load()
				fmt.Println(ix, " r: ", r)
				fmt.Println(ix, " c: ", c)
				// Handle request r using config c.
			}
		}(i)
	}
	fmt.Scanln()
}
