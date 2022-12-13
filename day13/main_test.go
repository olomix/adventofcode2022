package main

import "testing"

func TestOne(t *testing.T) {
	left := []any{float64(9)}
	right := []any{float64(8), float64(7), float64(6)}
	r := cmp(left, right)
	if r != 1 {
		t.Fatalf("expected 1, got %v", r)
	}
}

func TestTwo(t *testing.T) {
	f := func(a int) float64 { return float64(a) }
	left := []any{[]any{f(2)}}
	right := []any{f(1), f(1), f(3), f(1), f(1)}
	r := cmp(left, right)
	if r != 1 {
		t.Fatalf("expected 1, got %v", r)
	}
}
