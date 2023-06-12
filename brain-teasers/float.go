package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1.1
	fmt.Println(n * n)

	fmt.Println(math.NaN() == math.NaN())
}
