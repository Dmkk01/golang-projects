package main

import (
	"fmt"
	"time"
)

type Log struct {
	Message string
	time.Time
}

func main() {
	ts := time.Date(2017, time.November, 10, 0, 0, 0, 0, time.UTC)
	log := Log{"log message", ts}
	fmt.Println(log)
}
