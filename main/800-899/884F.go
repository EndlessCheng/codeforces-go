package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf884F(in io.Reader, out io.Writer) {
	var n, ans, totC int
	var s string
	Fscan(in, &n, &s)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
		ans += b[i]
	}

	cnt := [26]int{}
	for i := range n / 2 {
		if s[i] == s[n-1-i] {
			cnt[s[i]-'a']++
			totC++
			ans -= min(b[i], b[n-1-i])
		}
	}

	maxC, maxCh := 0, byte(0)
	for ch, c := range cnt {
		if c > maxC {
			maxC, maxCh = c, 'a'+byte(ch)
		}
	}

	if maxC*2 > totC {
		a := []int{}
		for i := range n / 2 {
			if s[i] != maxCh && s[n-1-i] != maxCh && s[i] != s[n-1-i] {
				a = append(a, min(b[i], b[n-1-i]))
			}
		}
		slices.Sort(a)
		for _, v := range a[:maxC*2-totC] {
			ans -= v
		}
	}

	Fprint(out, ans)
}

//func main() { cf884F(os.Stdin, os.Stdout) }
