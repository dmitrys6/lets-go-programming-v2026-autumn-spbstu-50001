package main

import "fmt"

func main() {

	var operand1 int = 0
	var operand2 int = 0
	var operation string

	_, err := fmt.Scan(&operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	if operation == "+" {
		fmt.Println(operand1 + operand2)
	} else if operation == "-" {
		fmt.Println(operand1 - operand2)
	} else if operation == "*" {
		fmt.Println(operand1 * operand2)
	} else if operation == "/" {

		if operand2 == 0 {
			fmt.Println("Division by zero")
			return
		}

		fmt.Println(operand1 / operand2)

	} else {
		fmt.Println("Invalid operation")
	}
}
