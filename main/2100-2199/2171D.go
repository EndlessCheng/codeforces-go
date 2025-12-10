package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2171D(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		st := []int{}
		for range n {
			Fscan(in, &v)
			mn := v
			for len(st) > 0 && st[len(st)-1] < v {
				mn = min(mn, st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = append(st, mn)
		}
		if len(st) > 1 {
			Fprintln(out, "No")
		} else {
			Fprintln(out, "Yes")
		}
	}
}

//func main() { cf2171D(bufio.NewReader(os.Stdin), os.Stdout) }
