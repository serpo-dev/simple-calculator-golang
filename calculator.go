package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Hello! It's a simple calculator.")

	fmt.Println("Input the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Input the second value:")
	fmt.Scanln(&num2)

	fmt.Println("Give me an operator ( * / + - ):")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1+num2)
	case "-":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1-num2)
	case "*":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Dividing by zero is restricted.")
		} else {
			fmt.Printf("%f %s %f = %f", num1, operator, num2, num1/num2)
		}
	default:
		fmt.Println("Invalid operator.")
	}
}
