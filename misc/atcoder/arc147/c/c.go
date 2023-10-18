package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	b := append(a[:0:0], a...)
	sort.Slice(b, func(i, j int) bool { return b[i].l > b[j].l })

	for i := 0; a[i].r < b[i].l; i++ {
		ans += (n - 1 - i*2) * (b[i].l - a[i].r)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
