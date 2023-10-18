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
	l := make([]int, n)
	r := make([]int, n)
	for i := range l {
		Fscan(in, &l[i], &r[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	sort.Ints(r)
	for i := 0; l[i] > r[i]; i++ {
		ans += (n - 1 - i*2) * (l[i] - r[i])
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
