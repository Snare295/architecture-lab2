package lab2

import (
	"fmt"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	const operators = "+-*/^"

	//verify empty input
	if len(input) == 0 {
		return "", fmt.Errorf("empty input")
	}

	//verify first char is operand
	if !strings.Contains(operators, input[0:1]) {
		return input, nil
	}

	var row = strings.Split(input, " ")

	var stack []string
	var res = ""
	op1, op2, exp := "", "", ""
	for i := len(row) - 1; i >= 0; i-- {
		var err = false
		if !strings.Contains(operators, row[i]) {
			Push(&stack, row[i])
		} else {
			op1, err = Pop(&stack)
			if err {
				return "", fmt.Errorf("incorrect expression")
			}
			op2, err = Pop(&stack)
			if err {
				return "", fmt.Errorf("incorrect expression")
			}
			exp = op1 + " " + op2 + " " + row[i]
			Push(&stack, exp)
		}
	}

	var err = false
	res, err = Pop(&stack)
	if err {
		return "", fmt.Errorf("incorrect expression")
	}

	if len(stack) != 0 {
		return "", fmt.Errorf("incorrect expression")
	}
	return res, nil

}

func Push(stack *[]string, v string) {
	*stack = append(*stack, v)
}

func Pop(stack *[]string) (string, bool) {
	if len(*stack) == 0 {
		return "", true
	} else {
		index := len(*stack) - 1
		element := (*stack)[index]
		*stack = (*stack)[:index]
		return element, false
	}
}
