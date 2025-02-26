package main

import (
	"fmt"
	"os"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) float64 {
	if b == 0 {
		fmt.Println("Error: Cannot divide by zero")
		os.Exit(1)
	}
	return float64(a) / float64(b)
}

func Modulo(a, b int) int {
	return a % b
}

func main() {
	var a, b, choice int
	fmt.Println("Welcome to the calculator program!")
	fmt.Println("Please select an operation: ")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Modulo")
	fmt.Scanln(&choice)
	fmt.Println("Please enter two numbers: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	switch choice {
	case 1:
		fmt.Println("Sum: ", Add(a, b))
	case 2:
		fmt.Println("Difference: ", Subtract(a, b))
	case 3:
		fmt.Println("Prouduct: ", Multiply(a, b))
	case 4:
		fmt.Println("Quotient: ", Divide(a, b))
	case 5:
		fmt.Println("Modulo: ", Modulo(a, b))
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

