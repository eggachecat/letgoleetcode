package search

import "testing"

func TestSolutionsFor127At0_LadderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		s    *SolutionsFor127At0
		args args
		want int
	}{
		{
			name: "0",
			s:    &SolutionsFor127At0{},
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "1",
			s:    &SolutionsFor127At0{},
			args: args{
				beginWord: "a",
				endWord:   "c",
				wordList:  []string{"a", "b", "c"},
			},
			want: 2,
		},
		{
			name: "2",
			s:    &SolutionsFor127At0{},
			args: args{
				beginWord: "qa",
				endWord:   "sq",
				wordList:  []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.LadderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("SolutionsFor127At0.LadderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
