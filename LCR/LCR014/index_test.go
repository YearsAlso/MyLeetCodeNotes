package LCR014

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "LCR014 case1",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		}, {
			name: "LCR014 case2",
			args: args{
				s1: "abc",
				s2: "cccccbabbbaaaa",
			},
			want: true,
		}, {
			name: "LCR014 case3",
			args: args{
				s1: "adc",
				s2: "dcda",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
