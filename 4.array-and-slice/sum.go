package main

func Sum(numbers []int) int {
	result := 0

	for _, v := range numbers {
		result += v
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			tails := numbers[1:]
			result = append(result, Sum(tails))
		}
	}

	return result
}
