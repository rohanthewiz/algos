package sliding_win

import "testing"

func Test_longestSetOfCharsNotAlreadySeen(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{name: "Longest set chars not seen - middle",
			args:    args{in: "micoemozfklyddoreos"},
			wantOut: "zfklyd",
		},
		{name: "Longest chars not seen - empty",
			args:    args{in: ""},
			wantOut: "",
		},
		{name: "Longest chars not seen - end",
			args:    args{in: "mousemicoemozfklydr"},
			wantOut: "zfklydr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := longestSetOfCharsNotAlreadySeen(tt.args.in); gotOut != tt.wantOut {
				t.Errorf("longestSetOfCharsNotAlreadySeen() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
