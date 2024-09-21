package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf2008E(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := [2][26]int{}
		for i, b := range s {
			cnt[i%2][b-'a']++
		}
		if n%2 == 0 {
			Fprintln(out, n-slices.Max(cnt[0][:])-slices.Max(cnt[1][:]))
			continue
		}
		mx := 0
		for i, b := range s {
			cnt[i%2][b-'a']--
			mx = max(mx, slices.Max(cnt[0][:])+slices.Max(cnt[1][:]))
			cnt[i%2^1][b-'a']++
		}
		Fprintln(out, n-mx)
	}
}

//func main() { cf2008E(bufio.NewReader(os.Stdin), os.Stdout) }
