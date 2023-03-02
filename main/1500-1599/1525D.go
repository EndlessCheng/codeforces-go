package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1525D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, v int
	Fscan(in, &n)
	a := [2][]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		a[v] = append(a[v], i)
	}
	p := a[1]
	m := len(p)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = 1e9
	}
	for _, v := range a[0] {
		for j := m - 1; j >= 0; j-- {
			f[j+1] = min(f[j+1], f[j]+abs(v-p[j]))
		}
	}
	Fprint(out, f[m])
}

//func main() { CF1525D(os.Stdin, os.Stdout) }
