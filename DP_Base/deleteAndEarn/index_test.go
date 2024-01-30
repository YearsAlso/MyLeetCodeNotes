package deleteAndEarn

import "testing"

func Test_deleteAndEarn(t *testing.T) {
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
			name: "case1",
			args: args{
				nums: []int{3, 4, 2},
			},
			want: 6,
		},
		{
			name: "case1",
			args: args{
				nums: []int{2, 2, 3, 3, 3, 4},
			},
			want: 9,
		},
		{
			name: "case3",
			args: args{
				nums: []int{1, 3},
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				nums: []int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10},
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteAndEarn(tt.args.nums); got != tt.want {
				t.Errorf("deleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
