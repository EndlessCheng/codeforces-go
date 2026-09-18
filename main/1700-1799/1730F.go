package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1730F(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	k++
	p := make([]int, n+1)
	pos := make([]int, n+k+2)
	sum := make([]int, n+k+3)
	f := make([][]int, n+k+1)
	for i := range f {
		f[i] = make([]int, 1<<k)
		for j := range f[i] {
			f[i][j] = 1e18
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		Fscan(in, &p[i])
		pos[p[i]] = i
	}

	for i := range n + 1 {
		for j := 1; j <= pos[i]; j++ {
			sum[j]++
		}
		for j := 0; j < 1<<k; j += 2 {
			for x := range k {
				if i+x >= n {
					break
				}
				if j>>x&1 != 0 {
					continue
				}
				p := i
				q := j | 1<<x
				cnt := sum[pos[i+x+1]+1] + f[i][j]
				for q&1 != 0 {
					p++
					q >>= 1
				}
				for y := range k {
					if j>>y&1 != 0 && pos[i+y+1] > pos[i+x+1] {
						cnt++
					}
				}
				f[p][q] = min(f[p][q], cnt)
			}
		}
	}
	Fprint(out, f[n][0])
}

//func main() { cf1730F(bufio.NewReader(os.Stdin), os.Stdout) }
