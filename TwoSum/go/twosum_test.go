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
