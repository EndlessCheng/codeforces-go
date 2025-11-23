package main

import (
	"math/bits"
	"strconv"
)

// https://space.bilibili.com/206214
func minimumFlips(num int) int {
	n := uint32(num)
	rev := bits.Reverse32(n) >> bits.LeadingZeros32(n)
	return bits.OnesCount32(n ^ rev)
}

func minimumFlips1(n int) (ans int) {
	s := strconv.FormatUint(uint64(n), 2)
	m := len(s)
	for i := range m / 2 {
		if s[i] != s[m-1-i] {
			ans += 2
		}
	}
	return
}
