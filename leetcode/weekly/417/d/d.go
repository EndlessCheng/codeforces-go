package main

import "math/bits"

// https://space.bilibili.com/206214
func kthCharacter(k int64, operations []int) byte {
	k--
	inc := 0
	for i, op := range operations[:bits.Len64(uint64(k))] {
		if k>>i&1 > 0 {
			inc += op
		}
	}
	return 'a' + byte(inc%26)
}

func kthCharacter2(k int64, operations []int) byte {
	n := len(operations)
	if n == 0 {
		return 'a'
	}
	n--
	op := operations[n]
	operations = operations[:n]
	if n >= 63 || k <= 1<<n { // k 在左半边
		return kthCharacter(k, operations)
	}
	// k 在右半边
	ans := kthCharacter(k-1<<n, operations)
	return 'a' + (ans-'a'+byte(op))%26
}
