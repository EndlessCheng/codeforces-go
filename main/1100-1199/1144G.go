package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1144G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := range n {
		Fscan(in, &a[i])
	}

	inc, dec := -1, int(1e9)
	for i, v := range a[:n] {
		// 如果 a[i+1] > v，把 v 放到递增使得 a[i+1] 仍有机会放入递增（因为 a[i+1] > v = new last_inc），
		// 若把 v 放到递减，则 a[i+1] 也可能没法放进递减（取决于 last_dec），因此把 v 放进递增通常更有利
		// 如果 a[i+1] <= v，下一个元素无法进入以 v 为 last 的递增序列，
		// 但可能可以进入以 v 为 last 的递减序列；于是把 v 放入递减更合适
		if v > inc && (v >= dec || v < a[i+1]) {
			inc = v
			a[i] = 0
		} else if v < dec {
			dec = v
			a[i] = 1
		} else {
			Fprint(out, "NO")
			return
		}
	}

	Fprintln(out, "YES")
	for _, v := range a[:n] {
		Fprint(out, v, " ")
	}
}

//func main() { cf1144G(bufio.NewReader(os.Stdin), os.Stdout) }
