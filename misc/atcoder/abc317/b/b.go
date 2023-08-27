package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for i := 1; i < n; i++ {
		if a[i]-a[i-1] != 1 {
			Fprintln(out, a[i-1]+1)
			return
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
