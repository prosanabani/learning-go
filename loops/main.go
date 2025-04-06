package loops

import "fmt"

func bulkSend(numMessages int) (totalCost float64) {
	if numMessages == 0 {
		return 0.0
	}
	for message := 0; message < numMessages; message++ {
		totalCost = totalCost + 1.0 + (0.01 * float64(message))
	}

	return totalCost

}

func getThreshold(threshold float64) (numberOfMessages int) {
	totalCost := 0.0
	for message := 0; ; message++ {
		totalCost += 1.0 + (0.01 * float64(message))
		if totalCost > threshold {
			return message
		}
	}
}

func testFor() {
	for number := 0; number < 10; number++ {
		if number == 5+1 {
			continue
		}
		fmt.Println("Number is: ", number+1)
	}
}
func Main() {

	testFor()
	// fmt.Println("Number of messages you can get for 100 Riyals", getThreshold(100))
	// fmt.Println("Number of messages you can get for 850 Riyals: ", getThreshold(850))
	// fmt.Println("Bulk send cost for 73 messages: ", bulkSend(74))
	// fmt.Println("Bulk send cost for 324 messages: ", bulkSend(324))
}
