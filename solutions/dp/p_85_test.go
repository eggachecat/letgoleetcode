package search

import (
	"testing"
)

func TestSolutionsFor85At0_MaximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		s    *SolutionsFor85At0
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MaximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("SolutionsFor85At0.MaximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionsFor85At1_MaximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		s    *SolutionsFor85At1
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor85At1{},
			args: args{
				matrix: [][]byte{
					[]byte("10100"),
					[]byte("10111"),
					[]byte("11111"),
					[]byte("10010"),
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MaximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("SolutionsFor85At1.MaximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionsFor85At1_GetCurrentLeft(t *testing.T) {
	type args struct {
		l   int
		row []int
	}
	tests := []struct {
		name string
		s    *SolutionsFor85At1
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor85At1{},
			args: args{
				l:   3,
				row: []int{0, 1, 1, 1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetCurrentLeft(tt.args.l, tt.args.row); got != tt.want {
				t.Errorf("SolutionsFor85At1.GetCurrentLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionsFor85At1_GetCurrentRight(t *testing.T) {
	type args struct {
		r   int
		row []int
	}
	tests := []struct {
		name string
		s    *SolutionsFor85At1
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor85At1{},
			args: args{
				r:   3,
				row: []int{0, 1, 1, 1, 1},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetCurrentRight(tt.args.r, tt.args.row); got != tt.want {
				t.Errorf("SolutionsFor85At1.GetCurrentRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
