package search

import (
	"reflect"
	"testing"
)

func TestSolutionsFor37At0_SolveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		s    *SolutionsFor37At0
		args args
	}{
		{
			name: "0",
			s:    &SolutionsFor37At0{},
			args: args{
				board: [][]byte{
					[]byte("53..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SolveSudoku(tt.args.board)
			t.Errorf("SolveSudoku")
		})
	}
}

func Test_getChoiceSpace(t *testing.T) {
	type args struct {
		boardInt [][]int
		points   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0",
			args: args{
				boardInt: [][]int{
					[]int{5, 3, 0, 0, 7, 0, 0, 0, 0},
					[]int{6, 0, 0, 1, 9, 5, 0, 0, 0},
					[]int{0, 9, 8, 0, 0, 0, 0, 6, 0},
					[]int{8, 0, 0, 0, 6, 0, 0, 0, 3},
					[]int{4, 0, 0, 8, 0, 3, 0, 0, 1},
					[]int{7, 0, 0, 0, 2, 0, 0, 0, 6},
					[]int{0, 6, 0, 0, 0, 0, 2, 8, 0},
					[]int{0, 0, 0, 4, 1, 9, 0, 0, 5},
					[]int{0, 0, 0, 0, 8, 0, 0, 7, 9},
				},
				points: []int{0, 2},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "1",
			args: args{
				boardInt: [][]int{
					[]int{5, 3, 9, 2, 7, 1, 6, 4, 8},
					[]int{6, 8, 4, 1, 9, 5, 3, 2, 7},
					[]int{3, 9, 8, 7, 3, 7, 5, 6, 2},
					[]int{8, 7, 5, 9, 6, 2, 4, 1, 3},
					[]int{4, 2, 6, 8, 5, 3, 9, 0, 1},
					[]int{7, 0, 0, 0, 2, 0, 0, 0, 6},
					[]int{0, 6, 0, 0, 0, 0, 2, 8, 0},
					[]int{0, 0, 0, 4, 1, 9, 0, 0, 5},
					[]int{0, 0, 0, 0, 8, 0, 0, 7, 9},
				},
				points: []int{5, 1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChoiceSpace(tt.args.boardInt, tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getChoiceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateBoard(t *testing.T) {
	type args struct {
		boardInt [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "2",
			args: args{
				boardInt: [][]int{
					[]int{5, 3, 4, 6, 7, 8, 9, 1, 2},
					[]int{6, 7, 2, 1, 9, 5, 3, 4, 8},
					[]int{1, 9, 8, 3, 4, 2, 5, 6, 7},
					[]int{8, 5, 9, 7, 6, 1, 4, 2, 3},
					[]int{4, 2, 6, 8, 5, 3, 7, 9, 1},
					[]int{7, 1, 3, 9, 2, 4, 8, 5, 6},
					[]int{9, 6, 1, 5, 3, 7, 2, 8, 4},
					[]int{2, 8, 7, 4, 1, 9, 6, 3, 5},
					[]int{3, 4, 5, 2, 8, 6, 1, 7, 9},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateBoard(tt.args.boardInt); got != tt.want {
				t.Errorf("validateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateSubBoard(t *testing.T) {
	type args struct {
		boardInt [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0",
			args: args{
				boardInt: [][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
					[]int{7, 8, 9},
				},
			},
			want: true,
		},
		{
			name: "1",
			args: args{
				boardInt: [][]int{
					[]int{1, 2, 3},
					[]int{3, 5, 6},
					[]int{7, 8, 9},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateSubBoard(tt.args.boardInt); got != tt.want {
				t.Errorf("validateSubBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
