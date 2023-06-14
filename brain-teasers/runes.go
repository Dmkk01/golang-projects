package main

import "fmt"

func main() {
	msg := "π = 3.14159265358..."
	fmt.Println(msg)

	for _, c := range msg {
		fmt.Printf("%c", c)
	}
}
