package LCR045

import (
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"testing"
)

func Test_findBottomLeftValue(t *testing.T) {
	type args struct {
		root *TreeNode
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
				root: definition.BuildByArray([]any{2, 1, 3}),
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				root: definition.BuildByArray([]any{1, 2, 3, 4, nil, 5, 6, nil, nil, 7}),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
