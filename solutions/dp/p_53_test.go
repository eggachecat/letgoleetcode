package search

import (
	"testing"
)

func TestSolutionsFor53At0_MaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		s    *SolutionsFor53At0
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor53At0{},
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "case-2",
			s:    &SolutionsFor53At0{},
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "case-3",
			s:    &SolutionsFor53At0{},
			args: args{
				nums: []int{-1, -2},
			},
			want: -1,
		},
		{
			name: "case-3",
			s:    &SolutionsFor53At0{},
			args: args{
				nums: []int{-2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MaxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("SolutionsFor53At0.MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionsFor53At1_MaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		s    *SolutionsFor53At1
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor53At1{},
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "case-2",
			s:    &SolutionsFor53At1{},
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "case-3",
			s:    &SolutionsFor53At1{},
			args: args{
				nums: []int{-1, -2},
			},
			want: -1,
		},
		{
			name: "case-3",
			s:    &SolutionsFor53At1{},
			args: args{
				nums: []int{-2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MaxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("SolutionsFor53At1.MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionsFor53At2_MaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		s    *SolutionsFor53At2
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor53At2{},
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "case-2",
			s:    &SolutionsFor53At2{},
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "case-3",
			s:    &SolutionsFor53At2{},
			args: args{
				nums: []int{-1, -2},
			},
			want: -1,
		},
		{
			name: "case-3",
			s:    &SolutionsFor53At2{},
			args: args{
				nums: []int{-2, 1},
			},
			want: 1,
		},
		{
			name: "case-4",
			s:    &SolutionsFor53At2{},
			args: args{
				nums: []int{1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MaxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("SolutionsFor53At2.MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
