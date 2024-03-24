package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf166C(in io.Reader, out io.Writer) {
	var n, x, ans int
	Fscan(in, &n, &x)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)

	// 对于每个询问 x 都可以 O(log n) 解决
	i := sort.SearchInts(a, x)
	j := sort.SearchInts(a, x+1) - 1
	if i == n || a[i] != x {
		ans = 1
		j = i
		n++
	}
	m := (n - 1) / 2
	if i > m {
		ans += i*2 - n + 1
	} else if j < m {
		ans += n - j*2 - 2
	}
	Fprint(out, ans)
}

//func main() { cf166C(os.Stdin, os.Stdout) }
