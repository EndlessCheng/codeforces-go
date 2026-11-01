package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf567F(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	size := n*2 + 1
	f := make([][]int, size)
	for i := range f {
		f[i] = make([]int, size)
	}
	for i := 1; i <= n*2-1; i++ {
		f[i][i+1] = 1
	}
	opt := make([]string, k+1)
	x := make([]int, k+1)
	y := make([]int, k+1)

	for i := 1; i <= k; i++ {
		Fscan(in, &x[i], &opt[i], &y[i])
		if x[i] > y[i] {
			x[i], y[i] = y[i], x[i]
			if opt[i][0] != '=' {
				if opt[i][0] == '>' {
					opt[i] = "<" + opt[i][1:]
				} else {
					opt[i] = ">" + opt[i][1:]
				}
			}
		}
		if x[i]+1 == y[i] && opt[i][len(opt[i])-1] != '=' {
			f[x[i]][y[i]] = 0
		}
	}

	for sz := 4; sz <= n*2; sz++ {
		if sz%2 == 0 {
			for i := 1; i <= n*2-sz+1; i++ {
				j := i + sz - 1
				l := 1
				for ; l <= k; l++ {
					if x[l] == i && y[l] == i+1 && opt[l][len(opt[l])-1] != '=' ||
						(x[l] == i || x[l] == i+1) && y[l] > i+1 && y[l] <= j && opt[l][0] != '<' {
						break
					}
				}
				if l > k {
					f[i][j] += f[i+2][j]
				}

				l = 1
				for ; l <= k; l++ {
					if x[l] == j-1 && y[l] == j && opt[l][len(opt[l])-1] != '=' ||
						(y[l] == j || y[l] == j-1) && x[l] < j-1 && x[l] >= i && opt[l][0] != '>' {
						break
					}
				}
				if l > k {
					f[i][j] += f[i][j-2]
				}

				l = 1
				for ; l <= k; l++ {
					if x[l] == i && y[l] == j && opt[l][len(opt[l])-1] != '=' ||
						x[l] == i && y[l] > i && y[l] < j && opt[l][0] != '<' ||
						x[l] > i && x[l] < j && y[l] == j && opt[l][0] != '>' {
						break
					}
				}
				if l > k {
					f[i][j] += f[i+1][j-1]
				}
			}
		}
	}
	Fprint(out, f[1][n*2])
}

//func main() { cf567F(bufio.NewReader(os.Stdin), os.Stdout) }
