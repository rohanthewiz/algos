package sliding_win

import (
	"reflect"
	"testing"
)

func TestMoveZerosToArrEnd(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Move 0s to end",
			args: args{[]int{0, 3, 4, 0, 2, 0, 0, 0, 7}},
			want: []int{3, 4, 2, 7, 0, 0, 0, 0, 0},
		},
		{
			name: "Move all 0s to end",
			args: args{[]int{0, 0, 0, 0}},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveZerosToArrEnd(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZerosToArrEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveOnesToArrStart(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Move 1s to start",
			args: args{[]int{3, 4, 1, 2, 1, 1, 1, 7}},
			want: []int{1, 1, 1, 1, 3, 4, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveOnesToArrStart(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveOnesToArrStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
