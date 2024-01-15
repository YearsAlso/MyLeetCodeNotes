package LCR009

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "LCR009 Case1",
			args: args{
				nums: []int{10, 5, 2, 6},
				k:    100,
			},
			want: 8,
		},
		{
			name: "LCR009 Case2",
			args: args{
				nums: []int{1, 1, 1},
				k:    1,
			},
			want: 8,
		},
		{
			name: "LCR009 Case3",
			args: args{
				nums: []int{10, 2, 2, 5, 4, 4, 4, 3, 7, 7},
				k:    289,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
