package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(_case int) {
		var r, s int
		Fscan(in, &r, &s)
		ans := (r*(s-1) + 1) / 2
		Fprintln(out, ans)
		a := make([]int, 0, r*s)
		for i := 0; i < s; i++ {
			for j := 1; j <= r; j++ {
				a = append(a, j)
			}
		}
		n := len(a)
		for ; ans > 0; ans-- {
			for i := 1; i < n; i++ {
				// 相邻元素不等，记这一对为 (w,v)
				if v := a[i]; v != a[i-1] {
					st := i
					for i++; i < n && a[i] == v; i++ {
					}
					found := false
					for j := i; j < n; j++ {
						// 找下一对 (w,v) 的位置，模拟交换
						if a[j] == v {
							found = true
							Fprintln(out, i, j-i)
							a = append(append(append([]int{}, a[i:j]...), a[:i]...), a[j:]...)
							break
						}
					}
					// 若没找到则直接和剩下的交换
					if !found {
						Fprintln(out, st, n-st)
					}
					break
				}
			}
		}
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: ", _case)
		solve(_case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
