package LCR026

import (
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "LCR026 case1",
			args: args{
				head: definition.BuildList([]int{1, 2, 3, 4}),
			},
			want: []int{1, 4, 2, 3},
		},
		{
			name: "LCR026 case1",
			args: args{
				head: definition.BuildList([]int{1, 2, 3, 4, 5}),
			},
			want: []int{1, 5, 2, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReorderList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
