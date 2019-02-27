package search

import (
	"testing"
)

func TestSolutionsFor72At0_MinDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		s    *SolutionsFor72At0
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor72At0{},
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 3,
		},
		// {
		// 	name: "case-2",
		// 	s:    &SolutionsFor72At0{},
		// 	args: args{
		// 		word1: "hjros",
		// 		word2: "hoasdrse",
		// 	},
		// 	want: 2,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MinDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("SolutionsFor72At0.MinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionsFor72At1_MinDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		s    *SolutionsFor72At1
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor72At1{},
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MinDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("SolutionsFor72At1.MinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
