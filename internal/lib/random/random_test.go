package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "size-1",
			args: args{size: 1},
		},
		{
			name: "size-4",
			args: args{size: 4},
		},
		{
			name: "size-10",
			args: args{size: 10},
		},
		{
			name: "size-30",
			args: args{size: 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.args.size)
			str2 := NewRandomString(tt.args.size)

			assert.Len(t, str1, tt.args.size)
			assert.Len(t, str2, tt.args.size)

			assert.NotEqual(t, str1, str2)
		})
	}
}
