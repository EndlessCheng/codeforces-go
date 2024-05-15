package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf760B(in io.Reader, out io.Writer) {
	var n, tot, k int
	Fscan(in, &n, &tot, &k)
	tot -= n
	ans := 1 + sort.Search(tot, func(m int) bool {
		m++
		st := max(m-k+1, 0)
		cnt := (st + m) * (m - st + 1) / 2
		st = max(m-(n-k), 0)
		cnt += (st + m - 1) * (m - st) / 2
		return cnt > tot
	})
	Fprint(out, ans)
}

//func main() { cf760B(os.Stdin, os.Stdout) }
