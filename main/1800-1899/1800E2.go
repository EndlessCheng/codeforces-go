package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1800E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s, &t)
		k = min(k, n)
		if n-k < k {
			if s[n-k:k] != t[n-k:k] {
				Fprintln(out, "NO")
				continue
			}
			s = s[:n-k] + s[k:]
			t = t[:n-k] + t[k:]
		}
		cnt := [26]int{}
		for i, b := range s {
			cnt[b-'a']++
			cnt[t[i]-'a']--
		}
		if cnt == [26]int{} {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1800E2(os.Stdin, os.Stdout) }
