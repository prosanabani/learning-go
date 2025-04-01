package interfacesexamples

import "fmt"


type expense interface {
	cost(tax int ) int
}
type printer interface{
	print() string
}

type email struct {
	email    string
	password string
	tax int // Required field
}



func (currentEmail email) print() string {
	return currentEmail.email 
}

func (currentEmail email) cost(tax int) int {
	return tax *203
}


// func getCost(currentInterface expense){
// 	fmt.Println(currentInterface.cost())
// }


// func getPrint(currentPrinter printer){
// 	fmt.Println(currentPrinter.print())
// }


func getCostPrint(currentInterface expense, currentPrinter printer , tax ...int){
	if len(tax) > 0 {
		fmt.Println(currentInterface.cost(tax[0]))
		}else
		{
			fmt.Println(currentInterface.cost(0))
		}
fmt.Println(currentPrinter.print())
}	


func Interfacesexamples() {

	myEmail := email{
		email:    "skdjv",
		password: "kjnfkjn",
		tax:  10,
	}

	getCostPrint(myEmail , myEmail  , myEmail.tax  , myEmail.tax ,	myEmail.tax)
	getCostPrint(myEmail , myEmail)
}