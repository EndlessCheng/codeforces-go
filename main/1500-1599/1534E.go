package main

import (
	. "fmt"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1534E() {
	var n, k, res, ans int
	Scan(&n, &k)
	if n&1 > 0 && k&1 == 0 {
		Println(-1)
		return
	}
	type pair struct{ v, i int }
	c := make([]pair, n)
	for i := range c {
		c[i] = pair{1, i + 1}
	}
	tot := n
	// 1. tot 必须是 k 的倍数
	// 2. 至少需要 tot/k 次操作，所以需要满足 max(c[i].v) = c[0].v <= tot/k
	for i := 0; tot%k > 0 || c[0].v > tot/k; i = (i + 1) % n {
		c[i].v += 2 // c[i].v 必须是奇数
		tot += 2
	}
	for ; tot > 0; tot -= k {
		sort.Slice(c, func(i, j int) bool { return c[i].v > c[j].v })
		q := make([]interface{}, k)
		for i := range q {
			q[i] = c[i].i
			c[i].v--
		}
		Print("? ")
		Println(q...)
		Scan(&res)
		ans ^= res
	}
	Println("!", ans)
}

//func main() { CF1534E() }
