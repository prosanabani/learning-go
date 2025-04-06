package arrays

import "fmt"

const (
	planFree = "free"
	planPro  = "pro"
)

// func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
// 	if plan == "free" {
// 		return messages[0:2], nil
// 	}
// 	if plan == "pro" {
// 		return messages[:], nil
// 	}
// 	return nil, errors.New("unsupported plan")

// }

func modifySth(just []int) []int {
	just[0] = 88888
	return just
}

func Main() {

	// freePlans, err := getMessageWithRetriesForPlan(planFree, [3]string{"message1", "message2", "message3"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(freePlans)

	// proPlans, err := getMessageWithRetriesForPlan(planPro, [3]string{"message1", "message2", "message3"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(proPlans)

	// errorPlans, err := getMessageWithRetriesForPlan("kdsjnvd", [3]string{"message1", "message2", "message3"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return

	// }

	// fmt.Println(errorPlans)

	just_array := [5]int{1, 2, 3, 4, 5}
	slice := just_array[0:4]
	fmt.Println(modifySth(slice)) // pass the slice of the array
	fmt.Println(just_array)
}
