package arrayscan

import "testing"

func TestPairsSumNearToTarget(t *testing.T) {
	arr1 := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	arr2 := []int{12, 2, 1, 5, 7, 4, 8, 4, 3, 4, 1}

	idx1, idx2 := PairsSumNearToTarget(arr1, arr2, 17)
	t.Log("idx1", idx1, "idx2", idx2)
}
