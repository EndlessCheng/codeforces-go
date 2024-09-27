package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf755C(in io.Reader, out io.Writer) {
	var n, v int
	Fscan(in, &n)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	ans := n
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		v = find(v - 1)
		w := find(i)
		if v != w {
			ans--
			fa[v] = w
		}
	}
	Fprint(out, ans)
}

//func main() { cf755C(bufio.NewReader(os.Stdin), os.Stdout) }
