package maps

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]USER, name string) (deleted bool, err error) {

	targetUser, ok := users[name]

	if !ok {
		return false, errors.New("user not found")
	}
	if targetUser.scheduledForDeletion {
		delete(users, targetUser.name)
		return true, nil
	}

	return false, nil
}

func ExampleTwo() {

	fmt.Println(deleteIfNecessary(map[string]USER{
		"ali": {
			name:                 "ali ibrahim",
			phoneNumber:          738,
			scheduledForDeletion: false,
		},
		"moh": {
			name:                 "mohammed yousef",
			phoneNumber:          7354,
			scheduledForDeletion: true,
		},
	}, "ali"))

}
