package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1822E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		if n%2 > 0 {
			Fprintln(out, -1)
			continue
		}
		cnt := [26]int{}
		for i := 0; i < n/2; i++ {
			if s[i] == s[n-1-i] {
				cnt[s[i]-'a']++
			}
		}
		tot, mxI := 0, 0
		for i, c := range cnt {
			tot += c
			if c > cnt[mxI] {
				mxI = i
			}
		}
		mx := cnt[mxI]
		if mx*2 <= tot {
			Fprintln(out, (tot+1)/2)
			continue
		}
		left := mx*2 - tot
		t := 'a' + byte(mxI)
		for i := 0; i < n/2; i++ {
			if s[i] != s[n-1-i] && s[i] != t && s[n-1-i] != t {
				left--
			}
		}
		if left > 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, mx)
		}
	}
}

//func main() { cf1822E(os.Stdin, os.Stdout) }
