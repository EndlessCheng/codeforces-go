package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2110C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		pos := []int{}
		h := 0
		for i, d := range a {
			if d < 0 {
				pos = append(pos, i)
				a[i] = 0
			} else {
				h += d
			}
			Fscan(in, &l, &r)
			if a[0] < 0 || h > r || len(pos) < l-h {
				a[0] = -1
				continue
			}
			if l > h {
				for range l - h {
					a[pos[len(pos)-1]] = 1
					pos = pos[:len(pos)-1]
				}
				h = l
			}
			pos = pos[:min(r-h, len(pos))]
		}
		if a[0] < 0 {
			Fprintln(out, -1)
		} else {
			for _, v := range a {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf2110C(bufio.NewReader(os.Stdin), os.Stdout) }
