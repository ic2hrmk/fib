package model

import "testing"

var fibSeq = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233}

func TestFibonacciGenerator_Next(t *testing.T) {
	gen := NewFibonacciGenerator()

	for _, v := range fibSeq {
		if gen.Current() != v {
			t.Errorf("broken sequence, have [%4d] but expected [%4d]", gen.Current(), v)
		}

		gen.Next()
	}
}