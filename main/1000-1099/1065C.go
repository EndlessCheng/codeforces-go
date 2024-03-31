package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1065C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, s, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for i := 1; i < n; i++ {
		d := a[i-1] - a[i]
		if s+d*i < k {
			s += d * i
			continue
		}
		d -= (k - s) / i
		ans += 1 + d/(k/i)
		s = d % (k / i) * i
	}
	if s > 0 {
		ans++
	}
	Fprint(out, ans)
}

//func main() { cf1065C(os.Stdin, os.Stdout) }
