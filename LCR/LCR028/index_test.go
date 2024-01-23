package LCR028

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "LCR028 case1",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
