package array_ints

import "testing"

func TestUniqArrayInts(t *testing.T) {
	want := []int{2, 1, 5, 7, 3, 9, 4, 12}
	arr := UniqArrayInts([]int{2, 1, 5, 7, 3, 2, 1, 1, 12, 2, 5, 4, 2, 2, 9})
	if len(want) != len(arr) {
		t.Error("result array len doesn't match expected array. got:", len(arr), "want:", len(want))
	} else {
		for i := 0; i < len(want); i++ {
			if arr[i] != want[i] {
				t.Error("arrays do not match - expected", want, " - got", arr)
				break
			}
		}
	}
}
