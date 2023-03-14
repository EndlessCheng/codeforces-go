package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF687C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v int
	Fscan(in, &n, &k)
	f := make([][]bool, k+1)
	for i := range f {
		f[i] = make([]bool, k+1)
	}
	f[0][0] = true
	for ; n > 0; n-- {
		Fscan(in, &v)
		for j1 := k; j1 >= 0; j1-- {
			for j2 := k; j2 >= 0; j2-- {
				f[j1][j2] = f[j1][j2] || j1 >= v && f[j1-v][j2] || j2 >= v && f[j1][j2-v]
			}
		}
	}
	ans := []int{}
	for i, fi := range f {
		if fi[k-i] {
			ans = append(ans, i)
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF687C(os.Stdin, os.Stdout) }
