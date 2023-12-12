package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf702C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, ans, i int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	pre := int(-1e18)
	for ; m > 0; m-- {
		Fscan(in, &v)
		for ; i < n && a[i] <= v; i++ {
			ans = max(ans, min(a[i]-pre, v-a[i]))
		}
		pre = v
	}
	if i < n {
		ans = max(ans, a[n-1]-pre)
	}
	Fprint(out, ans)
}

//func main() { cf702C(os.Stdin, os.Stdout) }
