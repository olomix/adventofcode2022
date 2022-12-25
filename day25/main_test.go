package main

import "testing"

func Test_toSnafu(t *testing.T) {
	tests := []struct {
		in   uint
		want string
	}{
		{in: 1, want: "1"},
		{in: 2, want: "2"},
		{in: 3, want: "1="},
		{in: 4, want: "1-"},
		{in: 5, want: "10"},
		{in: 6, want: "11"},
		{in: 7, want: "12"},
		{in: 8, want: "2="},
		{in: 9, want: "2-"},
		{in: 10, want: "20"},
		{in: 15, want: "1=0"},
		{in: 20, want: "1-0"},
		{in: 2022, want: "1=11-2"},
		{in: 12345, want: "1-0---0"},
		{in: 314159265, want: "1121-1110-1=0"},
	}
	for _, tt := range tests {
		if got := toSnafu(tt.in); got != tt.want {
			t.Errorf("toSnafu() = %v, want %v", got, tt.want)
		}
	}
}

func Test_snafuToInt(t *testing.T) {
	tests := []struct {
		in   string
		want uint
	}{
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}
	for _, tt := range tests {
		if got := snafuToInt(tt.in); got != tt.want {
			t.Errorf("snafuToInt() = %v, want %v", got, tt.want)
		}
	}
}
