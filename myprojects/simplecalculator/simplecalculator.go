package simplecalculator

import "fmt"

func SimpleCalculator() {
	greetings()	
	for {
		firstNumber, secondNumber, calculationSign := takeInputs()
		calculate(firstNumber, secondNumber, calculationSign)
	}
	
}

func greetings() {
	appName := "Go calculator"
	fmt.Printf("Welcome to our %v\n", appName)
}
func takeInputs() (float64, float64, string) {
	var firstNumber float64
	var secondNumber float64
	var calculationSign string

	fmt.Println("Enter the first number")
	fmt.Scan(&firstNumber)
	fmt.Println("Enter + for addition\n\nEnter - for subtraction\n\nEnter * for multiplication\n\nEnter / for division")
	fmt.Scan(&calculationSign)
	fmt.Println("Enter the second number")
	fmt.Scan(&secondNumber)	
	return firstNumber, secondNumber, calculationSign
}

func calculate(firstNumber float64, secondNumber float64, calculationSign string) {
	var result float64
	if calculationSign == "+" {
		result = firstNumber + secondNumber
	}else if calculationSign == "-" {
		result = firstNumber - secondNumber
	}else if calculationSign == "/" {
		result = firstNumber / secondNumber
	}else if calculationSign == "*" {
		result = firstNumber * secondNumber
	}
	fmt.Printf("The result is %v\n", result)
}