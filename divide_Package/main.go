package dividepackage

import "fmt"

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("sdkjvd can't dividoing %v by zero", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
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
