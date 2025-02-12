package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n, n+1)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(a, -1)
	st := []int{-1}
	for r, x := range a {
		for len(st) > 1 && a[st[len(st)-1]] >= x {
			i := st[len(st)-1]
			st = st[:len(st)-1]
			ans += a[i] * (i - st[len(st)-1]) * (r - i)
		}
		st = append(st, r)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
