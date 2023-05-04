package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1689E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		ans := 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == 0 {
				ans++
				a[i] = 1
			}
		}

		conn := func() bool {
			left := n
			vis := make([]bool, n)
			var f func(int)
			f = func(v int) {
				vis[v] = true
				left--
				for i, b := range vis {
					if !b && a[v]&a[i] > 0 {
						f(i)
					}
				}
			}
			f(0)
			return left == 0
		}
		if !conn() {
			ans++
			for i := range a {
				a[i]++
				if conn() {
					goto o
				}
				a[i] -= 2
				if conn() {
					goto o
				}
				a[i]++
			}

			ans++
			id, mx := []int{}, 0
			for i, v := range a {
				lb := v & -v
				if lb > mx {
					id, mx = []int{i}, lb
				} else if lb == mx {
					id = append(id, i)
				}
			}
			a[id[0]]++
			a[id[1]]--
		}
	o:
		Fprintln(out, ans)
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1689E(os.Stdin, os.Stdout) }
