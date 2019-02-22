package general

import "testing"

func TestProblem20(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "{{",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				s: "{()}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem20(tt.args.s); got != tt.want {
				t.Errorf("Problem20() = %v, want %v", got, tt.want)
			}
		})
	}
}
