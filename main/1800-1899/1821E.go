package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1821E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, k, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k, &s)
		var a, st []int
		for i, b := range s {
			if b == '(' {
				st = append(st, i)
			} else {
				a = append(a, (i-st[len(st)-1])/2)
				st = st[:len(st)-1]
			}
		}
		sort.Ints(a)
		ans := 0
		for i := 0; i < len(a)-k; i++ {
			ans += a[i]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1821E(os.Stdin, os.Stdout) }
