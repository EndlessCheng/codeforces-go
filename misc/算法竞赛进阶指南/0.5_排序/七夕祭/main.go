package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func f(a []int) (ans int) {
	n := len(a)
	avg := 0
	for _, v := range a {
		avg += v
	}
	if avg%n != 0 {
		return -1
	}
	avg /= n
	sum := make([]int, n)
	sum[0] = a[0] - avg
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + a[i] - avg
	}
	sort.Ints(sum)
	mid := sum[n/2]
	for _, v := range sum {
		ans += abs(v - mid)
	}
	return
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, t, x, y int
	Fscan(in, &n, &m, &t)
	row := make([]int, n)
	col := make([]int, m)
	for ; t > 0; t-- {
		Fscan(in, &x, &y)
		row[x-1]++
		col[y-1]++
	}
	if ansR, ansC := f(row), f(col); ansR >= 0 && ansC >= 0 {
		Fprintln(out, "both", ansR+ansC)
	} else if ansR >= 0 {
		Fprintln(out, "row", ansR)
	} else if ansC >= 0 {
		Fprintln(out, "column", ansC)
	} else {
		Fprintln(out, "impossible")
	}
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
