package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2107C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		a := make([]int, n)
		pre := make([]int, n)
		maxF, f, has0 := 0, 0, false
		for i, b := range s {
			Fscan(in, &a[i])
			pre[i] = f
			if b == '0' {
				a[i] = -1e18
				f = 0
				has0 = true
			} else {
				f = max(f+a[i], 0)
				maxF = max(maxF, f)
			}
		}
		if maxF > k || maxF < k && !has0 {
			Fprintln(out, "No")
			continue
		}

		suf := 0
		for i := n - 1; i >= 0; i-- {
			if s[i] == '0' {
				a[i] = k - pre[i] - suf
				break
			}
			suf = max(suf+a[i], 0)
		}
		Fprintln(out, "Yes")
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2107C(bufio.NewReader(os.Stdin), os.Stdout) }
