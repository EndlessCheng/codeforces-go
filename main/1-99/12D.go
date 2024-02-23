package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf12D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]struct{ x, y, z int }, n)
	for i := range a {
		Fscan(in, &a[i].x)
	}
	sy := make([]int, n)
	for i := range a {
		Fscan(in, &a[i].y)
		sy[i] = a[i].y
	}
	for i := range a {
		Fscan(in, &a[i].z)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x > a[j].x })
	sort.Ints(sy)

	t := make([]int, n+1)
	for i := 0; i < n; {
		st := i
		v := a[st].x
		for ; i < n && a[i].x == v; i++ {
			a[i].y = n - sort.SearchInts(sy, a[i].y) // 要维护的是后缀最大值
			for j := a[i].y - 1; j > 0; j &= j - 1 {
				if t[j] > a[i].z {
					ans++
					break
				}
			}
		}
		for ; st < i; st++ {
			for j := a[st].y; j <= n; j += j & -j {
				t[j] = max(t[j], a[st].z)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf12D(os.Stdin, os.Stdout) }
