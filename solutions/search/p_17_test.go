package search

import (
	"reflect"
	"testing"
)

func TestSolutionsFor17At0_LetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		s    *SolutionsFor17At0
		args args
		want []string
	}{
		{
			name: "1",
			s:    &SolutionsFor17At0{},
			args: args{
				digits: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "2",
			s:    &SolutionsFor17At0{},
			args: args{
				digits: "",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.LetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolutionsFor17At0.LetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
