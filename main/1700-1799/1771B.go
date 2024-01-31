package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1771B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		l := make([]int, n+1)
		for ; m > 0; m-- {
			Fscan(in, &x, &y)
			if x > y {
				x, y = y, x
			}
			l[y] = max(l[y], x)
		}
		ans, maxL := 0, 0
		for i, x := range l {
			maxL = max(maxL, x)
			ans += i - maxL
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1771B(os.Stdin, os.Stdout) }
