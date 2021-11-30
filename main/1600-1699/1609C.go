package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1609C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 1e6
	np := [mx + 1]bool{1: true}
	for i := 2; i <= mx; i++ {
		if !np[i] {
			for j := 2 * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}

	var T, n, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &d)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := int64(0)
		for i, v := range a {
			if np[v] {
				continue
			}
			l, r := i, i
			for l >= d && a[l-d] == 1 {
				l -= d
			}
			for r+d < n && a[r+d] == 1 {
				r += d
			}
			ans += int64((i-l)/d+1)*int64((r-i)/d+1) - 1
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1609C(os.Stdin, os.Stdout) }
