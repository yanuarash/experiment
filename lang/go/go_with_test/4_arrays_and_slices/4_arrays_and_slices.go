package main

func Sum(numbers [3]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum
}
