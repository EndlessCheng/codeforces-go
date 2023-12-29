package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1768D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		p := make([]int, n)
		for i := range p {
			Fscan(in, &p[i])
			p[i]--
		}
		ans := n + 1
		id := make([]int, n)
		for v, i := range id {
			if i > 0 {
				continue
			}
			for id[v] == 0 {
				id[v] = ans
				v = p[v]
			}
			ans--
		}
		for i := 1; i < n; i++ {
			if id[i] == id[i-1] {
				ans -= 2
				break
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1768D(os.Stdin, os.Stdout) }
