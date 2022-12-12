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

	var n int
	Fscan(in, &n)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &p[i])
	}
	ans := [][2]int{}
	f := func(op, i int) {
		ans = append(ans, [2]int{op, i})
		p[i], p[i+1+op] = p[i+1+op], p[i]
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= n-2; j++ {
			if p[j]%2 == j%2 && p[j+2]%2 != j%2 {
				f(1, j)
			}
		}
	}
	for i := 1; i <= n-1; i++ {
		if p[i]%2 != i%2 && p[i+1]%2 != (i+1)%2 {
			f(0, i)
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= n-2; j++ {
			if p[j] > p[j+2] {
				f(1, j)
			}
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintf(out, "%c %d\n", 'A'+byte(p[0]), p[1])
	}
}

func main() { run(os.Stdin, os.Stdout) }
