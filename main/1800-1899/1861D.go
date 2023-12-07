package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1861D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		suf := make([]int, n)
		for i := n - 2; i >= 0; i-- {
			suf[i] = suf[i+1]
			if a[i] >= a[i+1] {
				suf[i]++
			}
		}
		ans := suf[0]
		pre := 1
		for i := 1; i < n; i++ {
			ans = min(ans, pre + suf[i])
			if a[i] >= a[i-1] {
				pre++
			}
		}
		ans = min(ans, pre)
		Fprintln(out, ans)
	}
}

//func main() { cf1861D(os.Stdin, os.Stdout) }
