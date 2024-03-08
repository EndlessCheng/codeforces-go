package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1744F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			pos[v] = i
		}

		ans := 1
		l, r := pos[0], pos[0]
		for sz := 2; sz <= n; sz++ {
			i := pos[(sz-1)/2]
			l = min(l, i, n-sz)
			r = max(r, i, sz-1)
			ans += max(sz-(r-l), 0)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1744F(os.Stdin, os.Stdout) }
