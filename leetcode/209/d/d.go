package main

// github.com/EndlessCheng/codeforces-go
func minimumOneBitOperations(n int) (ans int) {
	bits := []int{}
	for ; n > 0; n >>= 1 {
		bits = append(bits, n&1)
	}
	sign := 1
	for i := len(bits) - 1; i >= 0; i-- {
		if bits[i] > 0 {
			ans += sign * (1<<(i+1) - 1)
			sign = -sign
		}
	}
	return
}

func minimumOneBitOperations2(n int) (ans int) {
	for ; n > 0; n >>= 1 {
		ans ^= n
	}
	return
}
