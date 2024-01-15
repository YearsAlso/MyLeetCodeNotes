package LCR008

import (
	"testing"
)

func TestPerform(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perform(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("Perform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LCR008 Case1",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			name: "LCR008 Case2",
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			name: "LCR008 Case3",
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "LCR008 Case4",
			args: args{
				target: 11,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "LCR008 Case5",
			args: args{
				target: 80,
				nums:   []int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.args.target, tt.args.nums)
			if got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
