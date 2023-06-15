package main

var (
	id = nextID()
)

func nextID() int {
	id++
	return id
}

func main() {
	println(id)
}
