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
	var n, mx, ds, carry int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		mx = max(mx, a[i])
		for x := a[i]; x > 0; x /= 10 {
			ds += x % 10
		}
	}

	for pow10 := 10; pow10 <= mx*2; pow10 *= 10 {
		sort.Slice(a, func(i, j int) bool { return a[i]%pow10 < a[j]%pow10 })
		j := n - 1
		for _, v := range a {
			for j >= 0 && v%pow10+a[j]%pow10 >= pow10 {
				j--
			}
			carry += n - 1 - j
		}
	}
	Fprint(out, ds*n*2-carry*9)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
