package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1607D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		Fscan(in, &s)
		cnt := make([]int, n+1)
		cnt2 := make([]int, n+1)
		for i, v := range a {
			if s[i] == 'B' {
				if v < 1 {
					Fprintln(out, "NO")
					continue o
				}
				cnt[min(v, n)]++
			} else {
				if v > n {
					Fprintln(out, "NO")
					continue o
				}
				cnt2[max(v, 1)]++
			}
		}
		left := 0
		for i := 1; i <= n; i++ {
			left += 1 - cnt[i]
			if left < 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		left = 0
		for i := n; i > 0; i-- {
			left += 1 - cnt2[i]
			if left < 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1607D(os.Stdin, os.Stdout) }
