package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1898B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		for i := n - 2; i >= 0; i-- {
			v := a[i]
			k := (v - 1) / a[i+1]
			ans += k
			a[i] = v / (k + 1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1898B(os.Stdin, os.Stdout) }
