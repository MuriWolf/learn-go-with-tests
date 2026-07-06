package main

func SumAllTails(arraysToSum ...[]int) []int {
	var results []int

	for _, numbers := range arraysToSum {
		if len(numbers) == 0 {
			results = append(results, 0)
		} else {
			results = append(results, Sum(numbers[1:]))
		}
	}

	return results
}
