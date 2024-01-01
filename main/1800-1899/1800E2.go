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
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s, &t)
		cnt := [26]int{}
		for i := range s {
			if n-k <= i && i < k && s[i] != t[i] {
				Fprintln(out, "NO")
				continue o
			}
			cnt[s[i]-'a']++
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
