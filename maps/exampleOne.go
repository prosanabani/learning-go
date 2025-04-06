package maps

import (
	"errors"
	"fmt"
)

type USER struct {
	name                 string
	phoneNumber          int
	scheduledForDeletion bool
}

func getUserMap(names []string, phoneNumbers []int) (map[string]USER, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("length of names and phone numbers must be equal")
	}
	userMap := make(map[string]USER)
	for i := range names {

		userMap[names[i]] = USER{
			name:                 names[i],
			phoneNumber:          phoneNumbers[i],
			scheduledForDeletion: false,
		}

	}
	return userMap, nil
}

func ExampleOne() {

	names := []string{"John", "Jane", "Doe"}
	phoneNumbers := []int{1234567890, 9876543210, 1122334455}
	userMap, err := getUserMap(names, phoneNumbers)
	if err != nil {
		panic(err)
	}
	for name, user := range userMap {
		fmt.Println("Name:", name, "Phone Number:", user.phoneNumber)
	}

}
