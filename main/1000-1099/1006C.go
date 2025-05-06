package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1006C(in io.Reader, out io.Writer) {
	var n, ans, pre, suf int
	Fscan(in, &n)
	cnt := map[int]int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		pre += a[i]
		cnt[pre]++
	}
	for i := n - 1; i >= 0; i-- {
		cnt[pre]--
		pre -= a[i]
		suf += a[i]
		if cnt[suf] > 0 {
			ans = max(ans, suf)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1006C(bufio.NewReader(os.Stdin), os.Stdout) }
