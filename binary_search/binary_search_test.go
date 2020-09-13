package binary_search

import "testing"

func TestIterativeBinarySearchInt(t *testing.T) {
	type args struct {
		arrInt []int
		target int
	}

	var sortedArr = []int{1, 2, 3, 5, 7}

	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{
			name:      "Sorted array search - mid",
			args:      args{arrInt: sortedArr, target: 3},
			wantIndex: 2,
		},
		{
			name:      "Sorted array search - upper extreme",
			args:      args{arrInt: sortedArr, target: 7},
			wantIndex: 4,
		},
		{
			name:      "Sorted array search - lower extreme",
			args:      args{arrInt: sortedArr, target: 1},
			wantIndex: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := IterativeBinarySearchInt(tt.args.arrInt, tt.args.target); gotIndex != tt.wantIndex {
				t.Errorf("IterativeBinarySearchInt() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestIterativeBinarySearchString(t *testing.T) {
	type args struct {
		arrInt []string
		target string
	}

	var sortedArr = []string{"cat", "dog", "elephant", "horse", "mouse"}

	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{
			name:      "Sorted array search - mid",
			args:      args{arrInt: sortedArr, target: "horse"},
			wantIndex: 3,
		},
		{
			name:      "Sorted array search - upper extreme",
			args:      args{arrInt: sortedArr, target: "mouse"},
			wantIndex: 4,
		},
		{
			name:      "Sorted array search - lower extreme",
			args:      args{arrInt: sortedArr, target: "cat"},
			wantIndex: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := IterativeBinarySearchString(tt.args.arrInt, tt.args.target); gotIndex != tt.wantIndex {
				t.Errorf("IterativeBinarySearchString() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
