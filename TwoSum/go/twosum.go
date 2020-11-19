package twosum

import (
	"fmt"
)

func TwoSum(numbers []int, sum int) (int, int, error) {
	if len(numbers) < 2 {
		return -1, -1, fmt.Errorf("Need at least two numbers")
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == sum {
				return i, j, nil
			}
		}
	}

	return -1, -1, fmt.Errorf("Nothing found")
}
