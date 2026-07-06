package main

func SumAll(arraysToSum ...[]int) []int {
	arraysLen := len(arraysToSum)
	results := make([]int, arraysLen, arraysLen)

	for index, numbers := range arraysToSum {
		results[index] = Sum(numbers)
	}

	return results
}
