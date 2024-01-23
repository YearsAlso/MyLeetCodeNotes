package alternatingSubarray

import "testing"

func Test_alternatingSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "alternatingSubarray case1",
			args: args{
				nums: []int{2, 3, 4, 3, 4},
			},
			want: 4,
		},
		{
			name: "alternatingSubarray case2",
			args: args{
				nums: []int{4, 5, 6},
			},
			want: 2,
		}, {
			name: "alternatingSubarray case3",
			args: args{
				nums: []int{1, 29, 30, 5},
			},
			want: 2,
		}, {
			name: "alternatingSubarray case3",
			args: args{
				nums: []int{1, 2, 2, 1, 2, 3, 6, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatingSubarray(tt.args.nums); got != tt.want {
				t.Errorf("alternatingSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
