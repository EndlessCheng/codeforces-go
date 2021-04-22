package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	x := make([]int, n)
	y := make([]int, n)
	for i := range x {
		Fscan(in, &x[i], &y[i])
	}
	sort.Ints(y)
	for _, v := range y {
		ans += abs(v - y[n/2])
	}
	sort.Ints(x)
	for i := range x {
		x[i] -= i
	}
	sort.Ints(x)
	for _, v := range x {
		ans += abs(v - x[n/2])
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
