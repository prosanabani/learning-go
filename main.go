package main

import "fmt"

type expense interface {
	cost() float64
}

func getExpenseReport(e expense) (text string, cost float64) {
	switch v := e.(type) {
	case  email:
return "Email to " + v.toAddress , v.cost()
	case sms:
		return "SMS to " + v.toPhoneNumber  , v.cost()
	default:
return "Invalid expense type" ,  0.0
	}

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

func main() {

	myEmail := email{
		isSubscribed: true,
		body:         "kjvdnwjfdhv",
		toAddress:    "sdkjdnbkrje",
	}
	mySms := sms{
		isSubscribed:  true,
		body:          "kjvdnwjfdhv",
		toPhoneNumber: "986451",
	}

	fmt.Println(getExpenseReport(myEmail))
	fmt.Println(getExpenseReport(mySms))

}