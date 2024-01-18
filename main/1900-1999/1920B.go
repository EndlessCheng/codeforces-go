package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1920B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, p, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &p, &q)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
		}
		sort.Ints(s[1:])
		for i := 1; i < n; i++ {
			s[i+1] += s[i]
		}
		ans := int(-1e18)
		for i := 0; i <= p; i++ {
			ans = max(ans, s[max(n-i-q, 0)]*2-s[n-i])
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1920B(os.Stdin, os.Stdout) }
