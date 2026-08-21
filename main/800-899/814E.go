package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf814E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n int
	Fscan(in, &n)
	memo := make([][][]int, n+1)
	for i := range memo {
		memo[i] = make([][]int, n+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n+1)
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i < 0 || j < 0 {
			return
		}
		p := &memo[i][j][k]
		if *p != 0 {
			return *p - 1
		}
		if k > 0 {
			res = (i*dfs(i-1, j, k-1) + j*dfs(i+1, j-1, k-1)) % mod
		} else if i > 0 {
			res = ((i-1)*dfs(i-2, j, 0) + j*dfs(i, j-1, 0)) % mod
		} else if j > 2 {
			a := (j - 1) * (j - 2) / 2
			for c := 3; c <= j; c++ {
				res = (res + a*dfs(0, j-c, 0)) % mod
				a = a * (j - c) % mod
			}
		} else if j == 0 {
			res = 1
		}
		*p = res + 1
		return
	}

	d := make([]int, n+1)
	for i := range n {
		Fscan(in, &d[i+1])
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, max(n+1, 5))
	}
	for i := n - 1; i >= 0; i-- {
		a := 0
		b := 0
		if d[i+1] == 2 {
			a++
		}
		if d[i+1] == 3 {
			b++
		}
		for j := i + 1; j <= n; j++ {
			if j == n {
				f[i][j] = dfs(a, b, 0)
			} else {
				for k := j + 1; k <= n; k++ {
					f[i][j] = (f[i][j] + dfs(a, b, k-j)*f[j][k]) % mod
				}
			}
			if j < n {
				if d[j+1] == 2 {
					a++
				}
				if d[j+1] == 3 {
					b++
				}
			}
		}
	}
	Fprint(out, f[1][1+d[1]])
}

//func main() { cf814E(bufio.NewReader(os.Stdin), os.Stdout) }
