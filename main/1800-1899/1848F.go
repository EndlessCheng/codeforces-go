package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1848F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for m := n / 2; m > 0; m /= 2 {
		for i, w := range a[m:] {
			if a[i] != w {
				for j, w := range a[m:] {
					a[j] ^= w
				}
				ans += m
				break
			}
		}
		a = a[:m]
	}
	if a[0] > 0 {
		ans++
	}
	Fprint(out, ans)
}

//func main() { CF1848F(os.Stdin, os.Stdout) }
