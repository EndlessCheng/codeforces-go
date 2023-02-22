package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF584C(in io.Reader, out io.Writer) {
	var n, t, d int
	s := [2]string{}
	Fscan(bufio.NewReader(in), &n, &t, &s[0], &s[1])
	ans := []byte(s[0])
	add := func(i int) { ans[i] = (ans[i]-'a'+1)%26 + 'a' }
	for i, v := range ans {
		if v != s[1][i] {
			d++
			add(i)
			if ans[i] == s[1][i] {
				add(i)
			}
		}
	}
	d = (d - t) * 2 // d<0 说明缺了，d>0 说明多了
	for i := range ans {
		if d < 0 && s[0][i] == s[1][i] {
			add(i) // 不一样
			d += 2
		} else if d > 0 && s[0][i] != s[1][i] {
			d--
			ans[i] = s[d&1][i] // 和其中一个一样
		}
	}
	if d == 0 {
		Fprint(out, string(ans))
	} else {
		Fprint(out, -1)
	}
}

//func main() { CF584C(os.Stdin, os.Stdout) }
