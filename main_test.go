package main

import "testing"

func Test_getLetterByInt(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   rune
	}{
		{
			name:   "9",
			number: 25 + 25 + 9,
			want:   '9',
		},
		{
			name:   "Z",
			number: 25,
			want:   'Z',
		},
		{
			name:   "z",
			number: 25 + 25,
			want:   'z',
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
