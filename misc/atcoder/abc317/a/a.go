package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n,h,x int
	Fscan(in, &n,&h,&x)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i, v := range a {
		if h + v >= x {
			Fprintln(out, i+1)
			return
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
