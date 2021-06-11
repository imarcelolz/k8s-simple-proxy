package main

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	monitor := NewKubeMonitor("teste")
	fmt.Print(monitor.Services())
	// serverCreated := make(chan string)
	// serverRemoved := make(chan string)

	// for server := range serverCreated {
	// 	fmt.Printf("Server Created %s", server)
	// }

	// for server := range serverRemoved {
	// 	fmt.Printf("Server Removed %s", server)
	// }

	// go monitor.Watch(serverCreated, serverRemoved)
}
