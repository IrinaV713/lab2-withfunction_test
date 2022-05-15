package lab2

import (
	"errors"
	"fmt"
	"strings"
)

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(value string) {
	*s = append(*s, value)
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		elementIndex := len(*s) - 1
		element := (*s)[elementIndex]
		*s = (*s)[:elementIndex]
		return element, true
	}
}
func isOperator(character byte) bool {
	switch character {
	case '+', '-', '*', '/', '^', '%':
		return true
	default:
		return false
	}
}

func isOperand(character byte) bool {
	if (character >= '0' && character <= '9') ||
		(character >= 'a' && character <= 'z') ||
		(character >= 'A' && character <= 'Z') {
		return true
	}
	return false
}

func convert(input string) (string, error) {
	var stack Stack
	//input := "A B C / - A K / L - * "
	byteInput := []byte(input)
	var space byte = 32

	for _, character := range byteInput {
		if isOperator(character) {
			operand1 := stack[len(stack)-1]
			stack.pop()
			operand2 := stack[len(stack)-1]
			stack.pop()

			temp := string(character) + string(operand2) + string(operand1)
			stack.push(temp)

		} else if isOperand(byte(character)) {
			stack.push(string(character))
		} else if character == space {
			continue
		} else {
			return "", errors.New("Could not convert.\n")
		}

	}
	return stackToString(stack), nil

}

func stackToString(arg Stack) string {
	unmodified := []string(arg)
	modified := strings.Join(unmodified, "")
	return modified
}

func postfixToPrefix(input string) (string, error) {
	return convert(input)
}

