package LCR043

import (
	"reflect"
	"testing"
)

func TestConstructorTest(t *testing.T) {
	type args struct {
		funcName []string
		params   [][]int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				funcName: []string{"CBTInserter", "insert", "get_root"},
				params:   [][]int{{1}, {2}, {}},
			},
			want: []interface{}{nil, 1, []int{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructorTest(tt.args.funcName, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstructorTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
