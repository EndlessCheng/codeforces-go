package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2171F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		st := []int{}
		ans := [][2]int{}
		for range n {
			Fscan(in, &v)
			mn := v
			for len(st) > 0 && st[len(st)-1] < v {
				top := st[len(st)-1]
				ans = append(ans, [2]int{top, v})
				mn = min(mn, top)
				st = st[:len(st)-1]
			}
			st = append(st, mn)
		}
		if len(st) > 1 {
			Fprintln(out, "No")
			continue
		}
		Fprintln(out, "Yes")
		for _, p := range ans {
			Fprintln(out, p[0], p[1])
		}
	}
}

//func main() { cf2171F(bufio.NewReader(os.Stdin), os.Stdout) }
