package maps

import "fmt"

func getUserCount(usersIds []string) map[string]int {
	currentUserCount := map[string]int{}
	for _, userId := range usersIds {
		count := currentUserCount[userId]
		count++

		currentUserCount[userId] = count
	}
	return currentUserCount
}

func ExampleThree() {
	// tempMap := map[string]map[string]int{
	// 	"ali": {
	// 		"age": 7,
	// 	},
	// 	"mohammed": {
	// 		"age": 4,
	// 	},
	// }

	// fmt.Println(tempMap["ali"]["agekjn"])
	usersIds := []string{"123", "123", "1234", "12", "12", "123"}

	fmt.Println(getUserCount(usersIds))
}
