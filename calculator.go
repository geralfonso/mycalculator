package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (c calc) operate(input string, operation string) (int, error) {
	values := strings.Split(input, operation)
	fistNumber, err := parseString(values[0])
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	secondNumber, err := parseString(values[1])
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	switch operation {
	case "+":
		return fistNumber + secondNumber, nil
	case "-":
		return fistNumber - secondNumber, nil
	case "*":
		return fistNumber * secondNumber, nil
	case "/":
		return fistNumber / secondNumber, nil
	default:
		fmt.Println(operation, "is not supported")
		return 0, nil
	}

}

func parseString(value string) (int, error) {
	result, err := strconv.Atoi(value)
	return result, err
}

// ReadInput returns an string from the input
func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
