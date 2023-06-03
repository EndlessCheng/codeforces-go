package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF351E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	tree := map[int]int{}
	add := func(i int) {
		for i++; i <= 1e5+1; i += i & -i {
			tree[i]++
		}
	}
	preSum := func(i int) (res int) {
		// 本来要 +1 的，但这里求的是 <i 
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	pre := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] = abs(a[i])
		pre[i] = preSum(a[i])
		add(a[i])
	}

	tree = map[int]int{}
	for i := n - 1; i >= 0; i-- {
		ans += min(pre[i], preSum(a[i]))
		add(a[i])
	}
	Fprint(out, ans)
}

//func main() { CF351E(os.Stdin, os.Stdout) }
