package main

import (
	. "fmt"
	"math/rand"
	"time"
)

// https://space.bilibili.com/206214
func CF1840G2() {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	rand.Seed(time.Now().UnixNano())

	const d = 333
	const maxN int = 1e9
	mx, sum, n := 0, 0, maxN
	Scan(&mx)
	pos := map[int]int{mx: 0}
	q := func(x int) {
		sum += x
		Println("+", x)
		Scan(&x)
		if pos[x] > 0 {
			n = min(n, sum-pos[x])
		}
		pos[x] = sum
		mx = max(mx, x)
	}
	for i := 1; i < d; i++ {
		q(rand.Intn(maxN) + 1)
	}
	for i := 0; i < d; i++ {
		q(1)
	}
	q(mx) // 倒着走 n-mx 步，且在下一圈，这样 sum-pos[x] 就是 n
	for i := 1; i < d; i++ {
		q(d)
	}
	Println("!", n)
}

//func main() { CF1840G2() }
