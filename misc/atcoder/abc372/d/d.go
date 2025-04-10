package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := make([]any, n)
	st := []int{}
	for i := n - 1; i >= 0; i-- {
		ans[i] = len(st)
		v := a[i]
		for len(st) > 0 && v > a[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	Fprintln(out, ans...)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
