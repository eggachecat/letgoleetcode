package search

import "testing"

func TestSolutionsFor62At0_UniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		s    *SolutionsFor62At0
		args args
		want int
	}{
		{
			name: "case-1",
			s:    &SolutionsFor62At0{},
			args: args{
				m: 5,
				n: 4,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("SolutionsFor62At0.UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
