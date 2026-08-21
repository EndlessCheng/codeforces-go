package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1553H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)

	S := 1 << k
	f := make([][]int, k+1)
	mn := make([][]int, k+1)
	mx := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		f[i] = make([]int, S)
		mn[i] = make([]int, S)
		mx[i] = make([]int, S)
	}
	for i := 0; i < S; i++ {
		f[0][i] = 1e9
		mn[0][i] = 1e9
		mx[0][i] = -1e9
	}

	for i := 1; i <= n; i++ {
		var x int
		Fscan(in, &x)
		mn[0][x] = 0
		mx[0][x] = 0
	}

	for i := 1; i <= k; i++ {
		for x := 0; x < S; x++ {
			y := x ^ 1<<(i-1)
			mn[i][x] = min(mn[i-1][x], mn[i-1][y]+1<<(i-1))
			mx[i][x] = max(mx[i-1][x], mx[i-1][y]+1<<(i-1))
			f[i][x] = min(f[i-1][x], f[i-1][y], mn[i-1][y]+1<<(i-1)-mx[i-1][x])
		}
	}

	for i := 0; i < S; i++ {
		Fprint(out, f[k][i], " ")
	}
}

//func main() { cf1553H(bufio.NewReader(os.Stdin), os.Stdout) }
