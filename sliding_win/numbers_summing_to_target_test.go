package sliding_win

import (
	"reflect"
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Set 1",
			args: args{array: []int{2, 3, 4, 5, 6}, target: 10},
			want: []int{2, 4},
		},
		{
			name: "Set 2",
			args: args{array: []int{2, 13, 4, 8, 6}, target: 15},
			want: []int{2, 4},
		},
		{
			name: "Set 2",
			args: args{array: []int{2, 13, 4, 8, 6}, target: 3},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoNumberSum(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoNumberSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
