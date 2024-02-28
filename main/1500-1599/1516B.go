package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1516B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		all := 0
		for i := range a {
			Fscan(in, &a[i])
			all ^= a[i]
		}
		if all == 0 {
			Fprintln(out, "YES")
			continue
		}
		cnt := 0
		for i := 0; i < n; i++ {
			for s := 0; i < n && s^a[i] != all; i++ {
				s ^= a[i]
			}
			cnt++
		}
		if cnt > 2 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1516B(os.Stdin, os.Stdout) }
