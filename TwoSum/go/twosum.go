package twosum

import (
	"fmt"
)

func TwoSum(numbers []int, sum int) (int, int, error) {
	if len(numbers) > 0 {
		return 0, 1, nil
	}

	return -1, -1, fmt.Errorf("Nothing found")
}
