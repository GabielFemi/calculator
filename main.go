package main

import (
	"bufio"
	"fmt"
	"github.com/gabielfemi/calculator/calculator"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Println("Calculator App")

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