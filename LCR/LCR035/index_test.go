package LCR035

import "testing"

func Test_findMinDifference(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "LCR035 case1",
			args: args{
				timePoints: []string{"23:59", "00:00"},
			},
			want: 1,
		},
		{
			name: "LCR035 case1",
			args: args{
				timePoints: []string{"00:00", "23:59", "00:00"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinDifference(tt.args.timePoints); got != tt.want {
				t.Errorf("findMinDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
