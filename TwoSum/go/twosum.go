package twosum

import (
	"fmt"
)

func TwoSum(numbers []int, sum int) (int, int, error) {
	if len(numbers) < 2 {
		return -1, -1, fmt.Errorf("Need at least two numbers")
	}

	if numbers[0]+numbers[1] == sum {
		return 0, 1, nil
	}

	if numbers[1]+numbers[2] == sum {
		return 1, 2, nil
	}

	if len(numbers) > 0 {
		return 0, 2, nil
	}

	return -1, -1, fmt.Errorf("Nothing found")
}
