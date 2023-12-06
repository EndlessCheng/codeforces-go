package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1907C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, n, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := [26]int{}
		for _, c := range s {
			cnt[c-'a']++
		}
		mx := 0
		for _, c := range cnt {
			mx = max(mx, c)
		}
		Fprintln(out, max(mx*2-n, n%2))
	}
}

//func main() { cf1907C(os.Stdin, os.Stdout) }
