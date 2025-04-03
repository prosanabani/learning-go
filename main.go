package main

import "fmt"

type expense interface {
	cost(rate int) (expenseCost float64)
}

type formatter interface {
	print() (PrintedString string)
}

type email struct {
	isSubscribed bool
	body         string
}

func (e email) cost(rate int) float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * float64(rate)

	}
	return float64(len(e.body)) * float64(rate)
}

func (e email) print() string {
	return e.body
}

func PrintEmailDetails(e expense, f formatter , rate int) {
	fmt.Println("Email Details:")
	fmt.Println("Cost:", e.cost(rate))
	fmt.Println("Body:", f.print())
}

func main() {

	myEmail := email{
		isSubscribed: true,
		body:         "kjvdnwjfdhv",
	}



	PrintEmailDetails(myEmail, myEmail,  985)
}