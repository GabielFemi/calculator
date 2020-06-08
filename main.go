package main

import (
	"bufio"
	"fmt"
	"github.com/gabielfemi/calculator/calculator"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Println("Calculator App created using Golang")


	selectOperation()
	firstNumber := readint("Enter the first number: ")
	secondNumber := readint("Enter the second number: ")
	arithmetic := calculator.Arithmetic{}
	sum := arithmetic.Add(firstNumber, secondNumber)
	fmt.Println("The sum of", firstNumber, "and", secondNumber, "is", sum)
}


func readint(prompt string) int {
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	var input int
	for {
		_, err := fmt.Fscan(stdin, &input)
		if err == nil && input >= 0 {
			break
		}
		_, _ = stdin.ReadString('\n')
		logrus.Fatalln("Please enter a number!")
	}
	return input
}

func selectOperation() {
	fmt.Print("Select your operation (+, -, /, X) : ")
	stdin := bufio.NewReader(os.Stdin)
	var operationSelected string

	_, err := fmt.Fscan(stdin, &operationSelected)
	if err != nil {
		logrus.Fatalln("An Error occurred when receiving input")
	}

	operator, valid := isValidOperator(operationSelected)
	if valid {
		fmt.Println("You chose", operator)
	}

}

func isValidOperator(operator string)  (string, bool){
	validOperators := []string {"+", "-", "/", "x"}
	if Find(validOperators, operator) {
		return operator, true
	 } else {
	 	return operator, false
	}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}