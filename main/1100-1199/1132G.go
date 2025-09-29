package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1132G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	left := make([]int, n+1)
	st := []int{0}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		for len(st) > 1 && a[st[len(st)-1]] < a[i] {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	fa := make([]int, n+1)
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

	d := make([]int, n+2)
	sumD := 0
	j := 1
	for i := 1; i <= n; i++ {
		fa[i] = i
		l := find(left[i] + 1)
		if l > 0 {
			fa[l] = l - 1
		}
		if l <= j {
			sumD++
		} else {
			d[l]++
		}
		d[i+1]--
		if i >= k {
			Fprint(out, sumD, " ")
			j++
			sumD += d[j]
		}
	}
}

//func main() { cf1132G(bufio.NewReader(os.Stdin), os.Stdout) }
