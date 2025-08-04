package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2001D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		left := make([]int, n+1)
		for i := range a {
			Fscan(in, &a[i])
			left[a[i]]++
		}

		st := a[:0]
		vis := make([]bool, n+1)
		for _, v := range a {
			left[v]--
			if vis[v] {
				continue
			}
			for len(st) > 0 && left[st[len(st)-1]] > 0 {
				top := st[len(st)-1]
				if len(st)%2 > 0 {
					if v < top && (len(st) == 1 || v > st[len(st)-2] || left[st[len(st)-2]] == 0) {
						break
					}
				} else {
					if v > top && (v < st[len(st)-2] || left[st[len(st)-2]] == 0) {
						break
					}
				}
				st = st[:len(st)-1]
				vis[top] = false
			}
			st = append(st, v)
			vis[v] = true
		}

		Fprintln(out, len(st))
		for _, v := range st {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2001D(bufio.NewReader(os.Stdin), os.Stdout) }
