package main

import (
	"github.com/gabielfemi/calculator/calculator"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("Calculator App")
	arithmetic := calculator.Arithmetic{}
	result := arithmetic.Add(1, 3)
	logrus.Println(result)
}