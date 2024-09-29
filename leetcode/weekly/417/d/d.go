package main

import "math/bits"

// https://space.bilibili.com/206214
func kthCharacter(k int64, operations []int) byte {
	n := min(len(operations), bits.Len64(uint64(k-1)))
	inc := 0
	for i := n - 1; i >= 0; i-- {
		if k > 1<<i { // k 在右半边
			inc += operations[i]
			k -= 1 << i
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
