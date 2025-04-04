package main

import "fmt"

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("you can't divide %v by zero" , de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend:  dividend}
	}
	return dividend / divisor, nil
}

func main() {

	end := 40.0
	sor := 0.0

	answer , err :=divide(end,sor)
	if err != nil{
		fmt.Println(err)
		return 
	}
	fmt.Println(answer)
}