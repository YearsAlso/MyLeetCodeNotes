package LCR025

import (
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"reflect"
	"testing"
)

func Test_AddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "LCR025 case1",
			args: args{
				l1: definition.BuildList([]int{7, 2, 4, 3}),
				l2: definition.BuildList([]int{5, 6, 4}),
			},
			want: []int{7, 8, 0, 7},
		},
		{
			name: "LCR025 case2",
			args: args{
				l1: definition.BuildList([]int{2, 4, 3}),
				l2: definition.BuildList([]int{5, 6, 4}),
			},
			want: []int{8, 0, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
