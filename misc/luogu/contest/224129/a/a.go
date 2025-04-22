package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n,t int
	Fscan(in, &n, &t)
	type pair struct{ x, y int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}

	ans := 0
	for i := 1; i < n; i++ {
		d := a[i].x - a[i-1].y
		ans += max(d - t, 0)
	}
	Fprintln(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
