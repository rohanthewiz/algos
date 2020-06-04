package stringx

import "testing"

func Test_rvsString(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name:    "Fast String reverse",
			args:    args{in: "horse"},
			wantOut: "esroh",
		},
		{
			name:    "Empty String reverse",
			args:    args{in: ""},
			wantOut: "",
		},
		{
			name:    "Single rune String reverse",
			args:    args{in: "b"},
			wantOut: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := rvsString(tt.args.in); gotOut != tt.wantOut {
				t.Errorf("rvsString() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
