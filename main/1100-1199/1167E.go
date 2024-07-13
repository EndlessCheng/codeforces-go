package main

import (
	. "fmt"
	"io"
)

func cf1167E(in io.Reader, out io.Writer) {
	var n, x, v int
	Fscan(in, &n, &x)
	ps := make([]struct{ l, r, v int }, x+1)
	for i := range ps {
		ps[i].l = 1e9
		ps[i].v = i
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if ps[v].l == 1e9 {
			ps[v].l = i
		}
		ps[v].r = i
	}

	a := ps[:0]
	for _, p := range ps {
		if p.r > 0 {
			a = append(a, p)
		}
	}
	n = len(a)

	i := 0
	for i < n-1 && a[i].r < a[i+1].l {
		i++
	}
	if i == n-1 {
		Fprint(out, x*(x+1)/2)
		return
	}

	ans := a[i+1].v * (x + 1 - a[n-1].v) // 去掉后缀 a[<=i+1] ~ a[n-1]
	for j := n - 1; j == n-1 || a[j].r < a[j+1].l; j-- {
		for i >= 0 && a[i].r >= a[j].l {
			i--
		}
		ans += a[i+1].v * (a[j].v - a[j-1].v) // 去掉 a[<=i+1] ~ a[j-1]
	}
	Fprint(out, ans)
}

//func main() { cf1167E(bufio.NewReader(os.Stdin), os.Stdout) }
