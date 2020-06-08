package main

import (
	"fmt"
	"github.com/gabielfemi/calculator/calculator"
	"github.com/sirupsen/logrus"
	
)

func main() {
	logrus.Println("Calculator App")
	arithmetic := calculator.Arithmetic{}
	result := arithmetic.Add(1, 3)
	fmt.Println("The sum is", result)
}
