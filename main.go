package main

import "fmt"

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func getExpenseReport(e expense) (string, float64) {
	email, isEmail := e.(email)
	sms, isSms := e.(sms)

	if isEmail {
		return email.toAddress, e.cost()
	}
	if isSms {
		return sms.toPhoneNumber, e.cost()
	}

	return "", 0.0
}

func printType(e expense) (string  , float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress ,  v.cost()

	case sms:
		return v.toPhoneNumber   ,  v.cost()

	default:
		return "" , 0.0
	}
}

func main() {

	myEmail := email{
		isSubscribed: true,
		body:         "dskjnvskjn",
		toAddress:    "tusin",
	}

	mySms := sms{
		isSubscribed:  false,
		body:          "sdlkjnvd",
		toPhoneNumber: "42458645312",
	}


	myInvalid := invalid{}



	// x , y := getExpenseReport(myEmail)
	// fmt.Println(x)
	// fmt.Println(y)

	// smsX, smsY := getExpenseReport(mySms)
	// fmt.Println(smsX)
	// fmt.Println(smsY)


	smsX, smsY := printType(mySms)
	fmt.Println(smsX)
	fmt.Println(smsY)

	emailX, emailY := printType(myEmail)
	fmt.Println(emailX)
	fmt.Println(emailY)

	myInvalidX, myInvalidY := printType(myInvalid)
	fmt.Println(myInvalidX)
	fmt.Println(myInvalidY)


}