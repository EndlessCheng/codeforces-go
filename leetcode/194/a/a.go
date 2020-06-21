package main

func xorOperation(n int, start int) (ans int) {
	for i := 0; i < n; i++ {
		ans ^= i*2 + start
	}
	return
}
