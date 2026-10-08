package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf261D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, b, t int
	Fscan(in, &T, &n, &b, &t)
	t = min(t, n, b)
	a := make([]int, n+1)
	for range T {
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		f := make([]int, b+1)
		for range t {
			for j := 1; j <= n; j++ {
				f[a[j]] = f[a[j]-1] + 1
				for k := a[j] + 1; k <= b && f[k] < f[a[j]]; k++ {
					f[k] = f[a[j]]
				}
			}
		}
		Fprintln(out, f[b])
	}
}

//func main() { cf261D(bufio.NewReader(os.Stdin), os.Stdout) }
