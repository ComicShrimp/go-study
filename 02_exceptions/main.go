package main

import (
	"errors"
	"fmt"
)

func main() {
	numerator := 11
	denominator := 0

	fmt.Printf("Numerator: %v | Denominator: %v\n", numerator, denominator)

	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("Result: %v | Remainder: %v\n", result, remainder)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		var err error
		err = errors.New("Zero division is not allowed\n")

		return 0, 0, err
	}

	result := numerator / denominator
	remainder := numerator % denominator

	return result, remainder, nil
}
