package main

import (
	"bytes"
	. "fmt"
)

// https://github.com/EndlessCheng
func q37(l, r int) int {
	Println("?", l, r)
	Scan(&r)
	return r
}

func cf2037E() {
	var T, n int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		pre := q37(1, n)
		if pre == 0 {
			Println("! IMPOSSIBLE")
			continue
		}
		ans := bytes.Repeat([]byte{'1'}, n)
		for r := n - 1; r > 1; r-- {
			c := q37(1, r)
			if c == pre {
				ans[r] = '0'
			} else if c == 0 {
				r--
				for range pre {
					ans[r] = '0'
					r--
				}
				goto o
			}
			pre = c
		}
		ans[0] = '0'
	o:
		Printf("! %s\n", ans)
	}
}

//func main() { cf2037E() }
