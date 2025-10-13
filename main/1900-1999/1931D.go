package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1931D(in io.Reader, out io.Writer) {
	var T, n, x, y, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		type pair struct{ x, y int }
		cnt := map[pair]int{}
		ans := 0
		for range n {
			Fscan(in, &v)
			ans += cnt[pair{(x - v%x) % x, v % y}]
			cnt[pair{v % x, v % y}]++
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1931D(bufio.NewReader(os.Stdin), os.Stdout) }
