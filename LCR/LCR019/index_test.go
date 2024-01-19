package LCR019

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "LCR019 case1",
			args: args{
				s: "aba",
			},
			want: true,
		},
		{
			name: "LCR019 case2",
			args: args{
				s: "abca",
			},
			want: true,
		},
		{
			name: "LCR019 case3",
			args: args{
				s: "abc",
			},
			want: false,
		}, {
			name: "LCR019 case4",
			args: args{
				s: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
