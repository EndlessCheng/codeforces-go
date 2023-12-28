package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf990C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	n, s, ans := 0, "", 0
	l := map[int]int{}
	r := map[int]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		cnt, mn := 0, 0
		for _, b := range s {
			if b == '(' {
				cnt++
			} else {
				cnt--
				mn = min(mn, cnt)
			}
		}
		if cnt == mn {
			if cnt == 0 {
				ans++
			}
			ans += l[-cnt]
		}
		if mn >= 0 {
			ans += r[-cnt]
			l[cnt]++
		}
		if cnt == mn {
			r[cnt]++
		}
	}
	Fprint(out, ans)
}

//func main() { cf990C(os.Stdin, os.Stdout) }
