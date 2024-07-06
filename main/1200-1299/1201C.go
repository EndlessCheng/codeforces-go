package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func CF1201C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	slices.Sort(a)
	a = a[n/2:] // 题目保证 n 是奇数
	i, n := 1, len(a)
	for ; i < n; i++ {
		if a[i] != a[i-1] {
			if a[i]-a[i-1] > k/i {
				break
			}
			k -= i * (a[i] - a[i-1])
		}
	}
	Fprint(out, a[i-1]+k/i)
}

//func main() { CF1201C(os.Stdin, os.Stdout) }
