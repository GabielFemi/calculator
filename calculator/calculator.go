package calculator

type Arithmetic struct {}

func (a *Arithmetic) Add(num1 int, num2 int) int {
	return num1 + num2
}

func (a *Arithmetic) Subtract(num1 int, num2 int) int {
	return num1 - num2
}

func (a *Arithmetic) Divide(num1 int, num2 int) int {
	return num1 / num2
}

func (a *Arithmetic) Multiply(num1 int, num2 int) int {
	return num1 * num2
}