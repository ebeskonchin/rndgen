package main

import "testing"

func Test_getLetterByInt(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   rune
	}{
		{
			name:   "last letter must be 9",
			number: 25 + 25 + 9,
			want:   '9',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLetterByInt(tt.number); got != tt.want {
				t.Errorf("getLetterByInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
