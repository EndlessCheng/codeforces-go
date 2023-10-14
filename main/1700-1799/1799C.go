package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1799C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		var pre, suf []byte
		i := 0
		for n := len(s); i < n-1; i += 2 {
			if s[i] != s[i+1] {
				if s[i+1] == s[n-1] {
					s[i], s[(i+n)/2] = s[(i+n)/2], s[i]
				} else {
					suf = append(suf, s[i])
					i++
				}
				break
			}
			pre = append(pre, s[i])
			suf = append(suf, s[i])
		}
		for i, n := 0, len(suf); i < n/2; i++ {
			suf[i], suf[n-1-i] = suf[n-1-i], suf[i]
		}
		Fprintf(out, "%s%s%s\n", pre, s[i:], suf)
	}
}

//func main() { CF1799C(os.Stdin, os.Stdout) }
