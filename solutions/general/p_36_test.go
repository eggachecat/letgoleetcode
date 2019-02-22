package general

import "testing"

func TestSolutionsFor36At0_IsValidSudoku(t *testing.T) {
	type args struct {
		boardInt [][]byte
	}
	tests := []struct {
		name string
		s    *SolutionsFor36At0
		args args
		want bool
	}{
		{
			name: "0",
			s:    &SolutionsFor36At0{},
			args: args{
				boardInt: [][]byte{
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
			want: true,
		},
		{
			name: "1",
			s:    &SolutionsFor36At0{},
			args: args{
				boardInt: [][]byte{
					[]byte("83..7...."),
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
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsValidSudoku(tt.args.boardInt); got != tt.want {
				t.Errorf("SolutionsFor36At0.IsValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
