package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Println("Calculator App")

	firstNumber := readint("Enter the first number: ")
	secondNumber := readint("Enter the second number: ")
	fmt.Println(firstNumber, secondNumber)
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
		stdin.ReadString('\n')
		logrus.Fatalln("Please enter a number!")
	}
	return input
}