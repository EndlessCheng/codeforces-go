package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1691D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		type pair struct{ v int; s int64 }
		st, s := []pair{}, int64(0)
		for _, v := range a {
			for len(st) > 0 && st[len(st)-1].v <= v {
				if st[len(st)-1].s < s {
					Fprintln(out, "NO")
					continue o
				}
				st = st[:len(st)-1]
			}
			st = append(st, pair{v, s})
			s += int64(v)
		}

		st, s = []pair{}, 0
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for len(st) > 0 && st[len(st)-1].v <= v {
				if st[len(st)-1].s < s {
					Fprintln(out, "NO")
					continue o
				}
				st = st[:len(st)-1]
			}
			st = append(st, pair{v, s})
			s += int64(v)
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1691D(os.Stdin, os.Stdout) }
