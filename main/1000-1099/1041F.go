package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1041F(in io.Reader, out io.Writer) {
	var n, m, y, ans int
	Fscan(in, &n, &y)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &m, &y)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}

	f := func(d, m int) {
		cnt := map[int]int{}
		for _, v := range a {
			cnt[v%m]++
		}
		for _, v := range b {
			cnt[(v+d)%m]++
		}
		for _, c := range cnt {
			ans = max(ans, c)
		}
	}
	f(0, 1<<32)
	for d := 1; d <= 1e9; d *= 2 {
		f(d, d*2)
	}
	Fprint(out, ans)
}

//func main() { cf1041F(bufio.NewReader(os.Stdin), os.Stdout) }
