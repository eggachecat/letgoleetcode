package search

import (
	"reflect"
	"testing"
)

func Test_getChoices(t *testing.T) {
	type args struct {
		board      [][]byte
		targetChar byte
		duplicate  [][]bool
		point      []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "0",
			args: args{
				board: [][]byte{
					[]byte("ABC"),
					[]byte("HGH"),
					[]byte("LHN"),
				},
				targetChar: byte('H'),
				duplicate: [][]bool{
					[]bool{false, false, false},
					[]bool{false, false, false},
					[]bool{false, false, false},
				},
				point: []int{1, 1},
			},
			want: [][]int{{1, 2}, {1, 0}, {2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChoices(tt.args.board, tt.args.targetChar, tt.args.duplicate, tt.args.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getChoices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionsFor79At0_Exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		s    *SolutionsFor79At0
		args args
		want bool
	}{
		{
			name: "0",
			s:    &SolutionsFor79At0{},
			args: args{
				board: [][]byte{
					[]byte("ABC"),
					[]byte("HGH"),
					[]byte("LHN"),
				},
				word: "GHN",
			},
			want: true,
		},
		{
			name: "1",
			s:    &SolutionsFor79At0{},
			args: args{
				board: [][]byte{
					[]byte("aa"),
				},
				word: "aaa",
			},
			want: false,
		},
		{
			name: "0",
			s:    &SolutionsFor79At0{},
			args: args{
				board: [][]byte{
					[]byte("CAA"),
					[]byte("AAA"),
					[]byte("BCD"),
				},
				word: "AAB",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("SolutionsFor79At0.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
