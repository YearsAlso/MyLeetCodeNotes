package sumIndicesWithKSetBits

import "testing"

func Test_sumIndicesWithKSetBits(t *testing.T) {
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
			name: "case1",
			args: args{
				nums: []int{5, 10, 1, 5, 2},
				k:    1,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumIndicesWithKSetBits(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("sumIndicesWithKSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
