package LCR022

import (
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"reflect"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	type args struct {
		head *definition.ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
