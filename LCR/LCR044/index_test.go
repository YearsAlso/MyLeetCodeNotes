package LCR044

import (
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				root: definition.BuildByArray([]any{1, 3, 2, 5, 3, nil, 9}),
			},
			want: []int{1, 3, 9},
		},
		{
			name: "case2",
			args: args{
				root: definition.BuildByArray([]any{1, 3, 2}),
			},
			want: []int{1, 3},
		}, {
			name: "case3",
			args: args{
				root: definition.BuildByArray([]any{1}),
			},
			want: []int{1},
		}, {
			name: "case4",
			args: args{
				root: definition.BuildByArray([]any{1, nil, 2}),
			},
			want: []int{1, 2},
		}, {
			name: "case5",
			args: args{
				root: definition.BuildByArray([]any{}),
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
