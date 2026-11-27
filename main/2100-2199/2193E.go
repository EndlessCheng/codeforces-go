package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2193E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		has := make([]bool, n+1)
		for range n {
			Fscan(in, &v)
			has[v] = true
		}

		f := make([]int, n+1)
		for i := 2; i <= n; i++ {
			f[i] = 1e9
		}
		for i := 1; i <= n; i++ {
			if f[i] == 1e9 {
				continue
			}
			for j := 2; i*j <= n; j++ {
				if has[j] {
					f[i*j] = min(f[i*j], f[i]+1)
				}
			}
		}

		f[1] = -1
		if has[1] {
			f[1] = 1
		}
		for _, v := range f[1:] {
			if v == 1e9 {
				v = -1
			}
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2193E(bufio.NewReader(os.Stdin), os.Stdout) }
