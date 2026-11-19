package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf729F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	s := make([]int32, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}

	const maxK = 88
	f := [2][4000][maxK*2 + 1]int32{}
	for k := maxK; k > 0; k-- {
		for i := n - 1; i >= 0; i-- {
			low := i*2 + k - n
			for d := max(-k, low); d <= 0; d++ {
				res := s[i+k] - s[i] - f[k&1][i+k][d+k+maxK]
				if k < maxK && i+k+1 <= n-i+d {
					res = max(res, s[i+k+1]-s[i]-f[k&1^1][i+k+1][d+k+1+maxK])
				}
				f[k&1][i][d+maxK] = res
			}

			for d := max(1, low); d <= min(k, i); d++ {
				r := n - i + d
				res := s[r] - s[r-k] - f[k&1][i][d-k+maxK]
				if k < maxK && r-k-1 >= i {
					res = max(res, s[r]-s[r-k-1]-f[k&1^1][i][d-k-1+maxK])
				}
				f[k&1][i][d+maxK] = res
			}
		}
	}
	Fprint(out, f[1][0][maxK])
}

//func main() { cf729F(bufio.NewReader(os.Stdin), os.Stdout) }
