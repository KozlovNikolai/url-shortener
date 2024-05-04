package random

import "testing"

func TestNewRandomString(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-001",
			args: args{size: 4},
			want: "hjkd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomString(tt.args.size); got != tt.want {
				t.Errorf("NewRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
