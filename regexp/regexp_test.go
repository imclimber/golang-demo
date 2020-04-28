package regexp

import "testing"

func TestRegularExpressionMustMatch(t *testing.T) {
	type args struct {
		source  string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "a",
			args: args{
				source:  "hello, (1000)",
				pattern: "\\([0-9]+\\)$",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegularExpressionMustMatch(tt.args.source, tt.args.pattern); got != tt.want {
				t.Errorf("RegularExpressionMustMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
