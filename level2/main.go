package main

import (
	"fmt"
	"level2/structExample"
)

func main() {
	name1, _ := getNames()
	fmt.Println("Hello", name1)
	structExample.StructExample()
}

func getNames() (string, string) {
	return "Alice", "Bob"
}
