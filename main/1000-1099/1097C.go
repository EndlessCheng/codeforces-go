package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1097C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	var s string
	cnt := map[int]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		c, mn := 0, 0
		for _, b := range s {
			if b == '(' {
				c++
			} else {
				c--
				mn = min(mn, c)
			}
		}
		if mn == 0 || mn == c {
			if cnt[-c] > 0 {
				cnt[-c]--
				ans++
			} else {
				cnt[c]++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1097C(os.Stdin, os.Stdout) }
