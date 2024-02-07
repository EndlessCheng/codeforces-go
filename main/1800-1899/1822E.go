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
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		if n%2 > 0 {
			Fprintln(out, -1)
			continue
		}

		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
			if cnt[b-'a'] > n/2 {
				Fprintln(out, -1)
				continue o
			}
		}

		tot, mx := 0, 0
		same := [26]int{}
		for i := 0; i < n/2; i++ {
			if s[i] == s[n-1-i] {
				tot++
				same[s[i]-'a']++
				mx = max(mx, same[s[i]-'a'])
			}
		}
		Fprintln(out, max((tot+1)/2, mx))
	}
}

//func main() { cf1822E(os.Stdin, os.Stdout) }
