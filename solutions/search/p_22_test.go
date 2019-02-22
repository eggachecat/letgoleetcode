package search

import (
	"reflect"
	"testing"
)

func TestSolutionsFor22At0_GenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		s    *SolutionsFor22At0
		args args
		want []string
	}{
		{
			name: "0",
			s:    &SolutionsFor22At0{},
			args: args{
				n: 2,
			},
			want: []string{"(())", "()()"},
		},
		{
			name: "1",
			s:    &SolutionsFor22At0{},
			args: args{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GenerateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolutionsFor22At0.GenerateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
