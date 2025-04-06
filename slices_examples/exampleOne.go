package slicesexamples

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	dayCosts := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(dayCosts) {
			dayCosts = append(dayCosts, 0.0)
		}
		dayCosts[cost.day] += cost.value
	}
	return dayCosts
}

func ExampleOne() {
	fmt.Println(getCostByDay([]cost{
		{day: 3, value: 30.0},
		{day: 1, value: 10.0},
		{day: 2, value: 20.0},
		{day: 1, value: 10.0},
		{day: 4, value: 40.0},
	}))
}
