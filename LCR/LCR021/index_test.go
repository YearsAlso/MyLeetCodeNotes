package LCR021

import (
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *definition.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "LCR021 case1",
			args: args{
				head: (*ListNode)(definition.BuildList([]int{1, 2, 3, 4, 5})),
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
