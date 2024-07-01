package main

import (
	. "fmt"
	"io"
)

func cf1987D(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int16, n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			cnt[v-1]++
		}
		f := make([][]int16, n+1)
		for i := range f {
			f[i] = make([]int16, n+1)
		}
		for i := n - 1; i >= 0; i-- {
			c := int(cnt[i])
			for j := 0; j < n; j++ {
				if c == 0 {
					f[i][j] = f[i+1][j]
				} else {
					res := f[i+1][j+1] + 1
					if j >= c {
						res = min(res, f[i+1][j-c])
					}
					f[i][j] = res
				}
			}
		}
		Fprintln(out, f[0][0])
	}
}

//func main() { cf1987D(bufio.NewReader(os.Stdin), os.Stdout) }
