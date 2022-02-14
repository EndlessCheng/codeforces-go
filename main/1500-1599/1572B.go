package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1572B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		x, p := 0, 0
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			x ^= v
			if x == 0 && i&1 == 0 {
				p = i + 1
			}
		}
		if x > 0 || p == 0 {
			Fprintln(out, "NO")
			continue
		}
		Fprintln(out, "YES")
		ans := []int{}
		for i := 1; i < p; i += 2 {
			ans = append(ans, i)
		}
		for i := p + 1; i < n-1; i += 2 {
			ans = append(ans, i)
		}
		Fprintln(out, len(ans)*2)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		for i := len(ans) - 1; i >= 0; i-- {
			Fprint(out, ans[i], " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1572B(os.Stdin, os.Stdout) }
