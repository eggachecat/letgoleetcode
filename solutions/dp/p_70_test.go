package search

import "testing"

func TestSolutionsFor70At0_ClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		s    *SolutionsFor70At0
		args args
		want int
	}{
		{
			name: "case-0",
			s:    &SolutionsFor70At0{},
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("SolutionsFor70At0.ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
