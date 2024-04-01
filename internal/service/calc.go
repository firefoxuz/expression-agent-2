package service

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Calc struct {
	Task string
}

func NewCalc(task string) *Calc {
	return &Calc{
		Task: task,
	}
}

func (c *Calc) Calculate() (int, error) {
	operands := strings.Split(c.Task, " ")
	firstValue, _ := strconv.Atoi(operands[0])
	secondValue, _ := strconv.Atoi(operands[1])
	operator := operands[2]

	if operator == "/" && secondValue == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	result, err := evaluateOperator(operator, firstValue, secondValue)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func evaluateOperator(operator string, a, b int) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "^":
		return int(math.Pow(float64(a), float64(b))), nil
	default:
		return 0, errors.New("Unknown operator: " + operator)
	}
}
