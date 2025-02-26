package main

import (
	"fmt"
	"trial2/getter" // Import the "getter" package
)

func main() {
	name1, name2 := getter.GetNames()
	fmt.Println("Hello", name1, "and", name2)
}
