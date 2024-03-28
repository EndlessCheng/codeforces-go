package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1446D1(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, mx, ans int
	Fscan(in, &n)
	a := make([]int, n)
	cnt := [101]int{}
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
		if cnt[a[i]] > cnt[mx] {
			mx = a[i]
		}
	}
	pos := make([]int, n*2+1)
	for t, c := range cnt {
		if c == 0 || t == mx {
			continue
		}
		for i := range pos {
			pos[i] = -2
		}
		s := n
		pos[s] = -1
		for i, v := range a {
			if v == mx {
				s--
			} else if v == t {
				s++
			}
			if pos[s] >= -1 {
				ans = max(ans, i-pos[s])
			} else {
				pos[s] = i
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1446D1(os.Stdin, os.Stdout) }
