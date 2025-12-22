package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1822E(in io.Reader, out io.Writer) {
	var T, n int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		if n%2 > 0 {
			Fprintln(out, -1)
			continue
		}

		total := [26]int{}
		for _, b := range s {
			total[b-'a']++
			if total[b-'a'] > n/2 {
				Fprintln(out, -1)
				continue o
			}
		}

		cnt := [26]int{}
		k, mx := 0, 0
		for i := range n / 2 {
			if s[i] == s[n-1-i] {
				k++
				cnt[s[i]-'a']++
				mx = max(mx, cnt[s[i]-'a'])
			}
		}
		Fprintln(out, max((k+1)/2, mx))
	}
}

//func main() { cf1822E(bufio.NewReader(os.Stdin), os.Stdout) }
