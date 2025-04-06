package dividepackage

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}

func Main() {

	end := 40.0
	sor := 0.0

	fmt.Println()

	answer, err := divide(end, sor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}
