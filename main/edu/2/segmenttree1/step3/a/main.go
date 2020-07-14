package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	tree := [1e5 + 1]int{}
	add := func(i int) {
		for ; i <= n; i += i & -i {
			tree[i]++
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		Fprint(out, i-sum(v), " ")
		add(v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
