package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1931D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, y, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		ans := 0
		type pair struct{ x, y int }
		cnt := map[pair]int{}
		for ; n > 0; n-- {
			Fscan(in, &v)
			ans += cnt[pair{(x - v%x) % x, v % y}]
			cnt[pair{v % x, v % y}]++
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1931D(os.Stdin, os.Stdout) }
