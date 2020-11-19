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

func TestArrayWithAnotherThreeElements(t *testing.T) {
	t.Parallel()
	first, second, err := twosum.TwoSum([]int{1, 2, 5}, 7)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if first != 1 || second != 2 {
		t.Errorf("Expected (1, 2) but got: (%d, %d)", first, second)
	}
}

type testCase struct {
	sum           int
	numbers       []int
	first, second int
}

func TestWithCases(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{sum: 4, numbers: []int{1, 2, 3}, first: 0, second: 2},
	}

	for _, tc := range testCases {
		first, second, err := twosum.TwoSum(tc.numbers, tc.sum)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if first != tc.first || second != tc.second {
			t.Errorf("Expected (%d, %d) but got: (%d, %d)", tc.first, tc.second, first, second)
		}
	}
}
