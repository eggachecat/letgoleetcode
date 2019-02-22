package search

import (
	"reflect"
	"testing"
)

func TestSolutionsFor46At0_Permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		s    *SolutionsFor46At0
		args args
		want [][]int
	}{
		{
			name: "1",
			s:    &SolutionsFor46At0{},
			args: args{
				nums: []int{1, 2},
			},
			want: [][]int{[]int{1, 2}, []int{2, 1}},
		},
		{
			name: "2",
			s:    &SolutionsFor46At0{},
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3},
				[]int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolutionsFor46At0.Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
