package main

import (
	"fmt"
)

func main() {
	storage := createClockMemoryStorage()
	storage.delete("abc")
	fmt.Println("Oh my go")
}
