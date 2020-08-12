package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1045I(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	mp := map[string]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		cnt := [26]int{}
		for i := range s {
			cnt[s[i]-'a']++
		}
		t := []byte{}
		for i, c := range cnt {
			if c&1 == 1 {
				t = append(t, byte(i+'a'))
			}
		}
		mp[string(t)]++
	}
	ans := int64(0)
	for s, _c := range mp {
		c := int64(_c)
		ans += c * (c - 1) / 2
		for i := range s {
			ans += c * int64(mp[s[:i]+s[i+1:]])
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF1045I(os.Stdin, os.Stdout)
//}
