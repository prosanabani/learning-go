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

// func modifySth(just []int) []int {
// 	just[0] = 88888
// 	return just
// }

func variaticFunc(text string, numbers ...int) string {
	sum := 0
	for number := 0; number < len(numbers); number++ {
		sum += numbers[number]
	}
	return fmt.Sprintf("%s %d", text, sum)
}

func Main() {
	fmt.Println(variaticFunc("hello", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// arrayToSend := [3]string{"message1", "message2", "message3"}
	// freePlans, err := getMessageWithRetriesForPlan(planFree, arrayToSend)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(freePlans)

	// proPlans, err := getMessageWithRetriesForPlan(planPro, arrayToSend)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(proPlans)

	// errorPlans, err := getMessageWithRetriesForPlan("kdsjnvd", arrayToSend)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return

	// }

	// fmt.Println(errorPlans)

	// just_array := [5]int{1, 2, 3, 4, 5}
	// slice := just_array[0:4]
	// fmt.Println(modifySth(slice)) // pass the slice of the array
	// fmt.Println(just_array)

	// sliceArray := make([]int, 5, 10)
	// sliceArray[0] = 1
	// sliceArray[1] = 2
	// sliceArray[2] = 3
	// sliceArray[3] = 4
	// sliceArray[4] = 5

	// sliceArray = append(sliceArray, 6, 7, 8, 9, 10, 11) // append to the slice
	// fmt.Println(len(sliceArray))
	// fmt.Println(cap(sliceArray))

}
