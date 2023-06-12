package main

import (
	"fmt"
	"time"
)

func main() {
	// timeout := 3
	var timeout time.Duration = 3
	fmt.Println("before")
	time.Sleep(timeout * time.Millisecond)
	// time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("after")
}
