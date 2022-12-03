package main

import (
	"strconv"
	"testing"
)

func TestPriority(t *testing.T) {
	testCases := []struct {
		in   rune
		want int
	}{
		{in: 'a', want: 1},
		{in: 'p', want: 16},
		{in: 'L', want: 38},
		{in: 'P', want: 42},
		{in: 'v', want: 22},
		{in: 't', want: 20},
		{in: 's', want: 19},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.want), func(t *testing.T) {
			got := priority(tc.in)
			if got != tc.want {
				t.Errorf("priority(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
func TestCommonItem(t *testing.T) {
	testCases := []struct {
		in   [3]string
		want rune
	}{
		{
			in: [3]string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			want: 'r',
		},
		{
			in: [3]string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: 'Z',
		},
	}

	for _, tc := range testCases {
		t.Run(string(tc.want), func(t *testing.T) {
			got := commonItem(tc.in)
			if got != tc.want {
				t.Errorf("commonItem(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
