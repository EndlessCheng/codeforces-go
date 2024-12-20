package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2040C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		k--
		if k>>(n-1) > 0 {
			Fprintln(out, -1)
			continue
		}
		a := make([]int, n)
		l, r := 0, n-1
		for i := 1; i <= n; i++ {
			if k<<1>>(n-i)&1 == 0 {
				a[l] = i
				l++
			} else {
				a[r] = i
				r--
			}
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2040C(bufio.NewReader(os.Stdin), os.Stdout) }
