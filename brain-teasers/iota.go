package main

import "fmt"

const (
	Read = 1 << iota
	Write
	Execute
)

func main() {
	fmt.Printf("%d %d %d", Read, Write, Execute)
}
