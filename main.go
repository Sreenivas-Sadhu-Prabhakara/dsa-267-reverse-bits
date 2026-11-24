package main

import "fmt"

// ReverseBits mirrors a 32-bit word: bit 0 swaps with bit 31, and so on.
// Build the result by shifting it left while feeding in n's lowest bits.
func ReverseBits(n uint32) uint32 {
	var out uint32
	for i := 0; i < 32; i++ {
		out = out<<1 | n&1
		n >>= 1
	}
	return out
}

func main() {
	fmt.Printf("%032b\n", ReverseBits(0b00000010100101000001111010011100))
}
