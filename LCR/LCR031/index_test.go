package LCR031

import (
	"reflect"
	"testing"
)

func TestLCROptimize(t *testing.T) {
	type args struct {
		funcNames []string
		param     [][]int
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
				funcNames: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				param:     [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			want: []any{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4},
		},
		{
			name: "case2",
			args: args{
				funcNames: []string{"LRUCache", "put", "get"},
				param:     [][]int{{2}, {2, 1}, {2}},
			},
			want: []any{nil, nil, 1},
		}, {
			name: "case2",
			args: args{
				funcNames: []string{"LRUCache", "put", "put", "get", "put", "put", "get"},
				param:     [][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}},
			},
			want: []any{nil, nil, nil, 2, nil, nil, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCROptimize(tt.args.funcNames, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LCROptimize() = %v, want %v", got, tt.want)
			}
		})
	}
}
