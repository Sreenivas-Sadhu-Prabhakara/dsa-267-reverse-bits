package main

import (
	"math/bits"
	"math/rand"
	"testing"
)

func TestReverseBits(t *testing.T) {
	cases := []struct {
		in, want uint32
	}{
		{0b00000010100101000001111010011100, 0b00111001011110000010100101000000},
		{0, 0},
		{0xFFFFFFFF, 0xFFFFFFFF},
		{1, 1 << 31},
		{1 << 31, 1},
	}
	for _, c := range cases {
		if got := ReverseBits(c.in); got != c.want {
			t.Errorf("ReverseBits(%032b) = %032b, want %032b", c.in, got, c.want)
		}
	}
}

func TestProperties(t *testing.T) {
	rng := rand.New(rand.NewSource(3))
	for i := 0; i < 1000; i++ {
		v := rng.Uint32()
		if ReverseBits(ReverseBits(v)) != v {
			t.Fatalf("double reverse not identity at %d", v)
		}
		if ReverseBits(v) != bits.Reverse32(v) {
			t.Fatalf("disagrees with stdlib at %d", v)
		}
	}
}
