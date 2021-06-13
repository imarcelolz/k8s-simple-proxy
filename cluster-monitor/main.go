package main

import "fmt"

func main() {
	monitor := NewKubeMonitor("default")
	fmt.Println(monitor.Services())
	// serverCreated := make(chan string)
	// serverRemoved := make(chan string)

	// go monitor.Watch(serverCreated, serverRemoved)

	// for server := range serverCreated {
	// 	fmt.Printf("Server Created %s", server)
	// }

	// for server := range serverRemoved {
	// 	fmt.Printf("Server Removed %s", server)
	// }
}
