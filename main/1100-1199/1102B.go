package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1102B(in io.Reader, out io.Writer) {
	var n, k, v, mx int
	Fscan(in, &n, &k)
	cnt := [5001]int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &v)
		cnt[v]++
		a[i] = cnt[v]
		mx = max(mx, cnt[v])
	}
	if mx > k {
		Fprint(out, "NO")
		return
	}

	vis := make([]bool, mx+1)
	color := mx + 1
	for i, v := range a {
		if color > k {
			break
		}
		if vis[v] {
			a[i] = color
			color++
		} else {
			vis[v] = true
		}
	}

	Fprintln(out, "YES")
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { cf1102B(bufio.NewReader(os.Stdin), os.Stdout) }
