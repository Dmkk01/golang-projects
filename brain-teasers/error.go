package main

import "fmt"

type OSError int

func (e *OSError) Error() string {
	return "OS Error"
}

func FileExists(path string) (bool, error) {
	var err *OSError
	return false, err
}

func main() {
	if _, err := FileExists("path"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File exists")
	}
}
