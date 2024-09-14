package main

import (
	. "fmt"
	"io"
)

func cf1919D(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c0 := 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == 0 {
				c0++
			}
		}
		if c0 != 1 {
			Fprintln(out, "NO")
			continue
		}

		left := make([]int, n)
		st := []int{-1}
		for i, v := range a {
			for len(st) > 1 && st[len(st)-1] >= v {
				st = st[:len(st)-1]
			}
			left[i] = st[len(st)-1]
			st = append(st, v)
		}

		st = []int{n}
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for len(st) > 1 && st[len(st)-1] >= v {
				st = st[:len(st)-1]
			}
			right := st[len(st)-1]
			if v > 0 && right != v-1 && left[i] != v-1 {
				Fprintln(out, "NO")
				continue o
			}
			st = append(st, v)
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1919D(bufio.NewReader(os.Stdin), os.Stdout) }
