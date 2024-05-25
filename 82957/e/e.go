package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, maxW, ans int
	Fscan(in, &n, &maxW)
	a := make([]struct{ w, v int }, n)
	for i := range a {
		Fscan(in, &a[i].w, &a[i].v)
	}
	for i := 10; i >= 0; i-- {
		t := ans | 1<<i
		w := 2047
		for _, p := range a {
			if p.v&t == t {
				w &= p.w
			}
		}
		if w <= maxW {
			ans = t
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
