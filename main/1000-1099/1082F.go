package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1082F(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	sum := []int{0}
	ch := [][10]int{{}}
	for i := 1; i <= n; i++ {
		var s string
		var x int
		Fscan(in, &s, &x)
		cur := 0
		for _, b := range s {
			d := int(b - '0')
			if ch[cur][d] == 0 {
				ch[cur][d] = len(sum)
				sum = append(sum, 0)
				ch = append(ch, [10]int{})
			}
			cur = ch[cur][d]
		}
		sum[cur] += x
	}

	f := make([][][]int, len(sum))
	for i := range f {
		f[i] = make([][]int, m+1)
	}

	var dfs func(int, int)
	dfs = func(x, d int) {
		for i := range 10 {
			if ch[x][i] != 0 {
				dfs(ch[x][i], d+1)
			}
		}

		f[x] = make([][]int, m+1)
		for i := range m + 1 {
			f[x][i] = make([]int, d+1)
			for j := 0; j <= d; j++ {
				f[x][i][j] = j * sum[x]
			}
		}

		for i := range 10 {
			if ch[x][i] == 0 {
				continue
			}
			y := ch[x][i]
			for j := m; j >= 0; j-- {
				for k := range d + 1 {
					f[x][j][k] += f[y][0][k+1]
					for l := 1; l <= j; l++ {
						f[x][j][k] = min(f[x][j][k], f[x][j-l][k]+min(f[y][l][k+1], f[y][l-1][0]))
					}
				}
			}
		}
	}

	dfs(0, 0)
	Fprint(out, f[0][m][0])
}

//func main() { cf1082F(bufio.NewReader(os.Stdin), os.Stdout) }
