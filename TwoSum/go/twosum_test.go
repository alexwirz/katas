package twosum_test

import (
	"fmt"
	"testing"
	"twosum"
)

func TestEmptyArray(t *testing.T) {
	t.Parallel()
	first, second, err := twosum.TwoSum([]int{}, 2)
	if err == fmt.Errorf("Nothing found") {
		t.Errorf("Expected an error but got tuple: %d %d", first, second)
	}
}

func TestArrayWithTwoElements(t *testing.T) {
	t.Parallel()
	first, second, err := twosum.TwoSum([]int{1, 2}, 3)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if first != 0 || second != 1 {
		t.Errorf("Expected (0, 1) but got: (%d, %d)", first, second)
	}
}

func TestArrayWithThreeElements(t *testing.T) {
	t.Parallel()
	first, second, err := twosum.TwoSum([]int{1, 2, 3}, 4)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if first != 0 || second != 2 {
		t.Errorf("Expected (0, 2) but got: (%d, %d)", first, second)
	}
}
