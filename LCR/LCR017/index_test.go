package LCR017

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "LCR017 case1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		}, {
			name: "LCR017 case2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		}, {
			name: "LCR017 case3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myMinWindowPlus(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "LCR017 case1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		}, {
			name: "LCR017 case2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		}, {
			name: "LCR017 case3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myMinWindowPlus(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("myMinWindowPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}
