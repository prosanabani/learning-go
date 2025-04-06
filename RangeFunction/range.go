package rangefunction

import (
	"fmt"
	"slices"
)

func badWordDetection(messages []string, badWords []string) string {
	for index, msg := range messages {
		if slices.Contains(badWords, msg) {
			return fmt.Sprintf("Bad word detected: %s at index %d", msg, index)
		}
	}
	return "No bad words detected"
}

func Range() {
	BadWords := []string{"badword1", "badword2", "badword3"}

	Messages := []string{"hello", "badword1", "world", "badword2"}
	// Call the function
	result := badWordDetection(Messages, BadWords)
	fmt.Println(result)

}
