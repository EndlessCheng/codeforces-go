package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1907F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n, n+1)
		for i := range a {
			Fscan(in, &a[i])
		}
		a = append(a, a[0])

		ans := int(1e9)
		c := 0
		for i, v := range a[:n] {
			if v > a[i+1] {
				c++
				p = i + 1
			}
		}
		if c == 0 {
			ans = 0
		} else if c == 1 {
			ans = min(ans, p+2, n-p)
		}

		c = 0
		for i, v := range a[:n] {
			if v < a[i+1] {
				c++
				p = i + 1
			}
		}
		if c == 0 {
			ans = 0
		} else if c == 1 {
			ans = min(ans, min(p, n-p)+1)
		}

		if ans == 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1907F(os.Stdin, os.Stdout) }
