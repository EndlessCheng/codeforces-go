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
	var n, m, ans int
	Fscan(in, &n, &m)
	op := make([]struct{ a, c int }, m)
	for i := range op {
		Fscan(in, &op[i].a, &op[i].c)
	}
	sort.Slice(op, func(i, j int) bool { return op[i].c < op[j].c })
	for _, o := range op {
		g := gcd(n, o.a)
		ans += (n - g) * o.c
		n = g
	}
	if n > 1 {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
