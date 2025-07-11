package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2106D(in io.Reader, out io.Writer) {
	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}

		suf := make([]int, n+1)
		for i := n - 1; i >= 0; i-- {
			suf[i] = suf[i+1]
			if a[i] >= b[m-1-suf[i]] {
				suf[i]++
				if suf[i] == m {
					Fprintln(out, 0)
					continue o
				}
			}
		}

		ans := int(2e9)
		if suf[0] == m-1 {
			ans = b[0]
		}
		pre := 0
		for i, v := range a {
			if v >= b[pre] {
				pre++
			}
			if pre+suf[i+1] >= m {
				Fprintln(out, 0)
				continue o
			}
			if pre+suf[i+1] == m-1 {
				ans = min(ans, b[pre])
			}
		}
		if ans == 2e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2106D(bufio.NewReader(os.Stdin), os.Stdout) }
