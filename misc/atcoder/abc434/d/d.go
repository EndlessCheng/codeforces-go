package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 2000
	diff := [mx + 2][mx + 2]int{}
	var n, r1, r2, c1, c2 int
	Fscan(in, &n)
	for v := n; v < n*2; v++ {
		Fscan(in, &r1, &r2, &c1, &c2)
		diff[r1][c1] += v
		diff[r1][c2+1] -= v
		diff[r2+1][c1] -= v
		diff[r2+1][c2+1] += v
	}

	ans := make([]int, n)
	zero := 0
	for i := range mx {
		for j := range mx {
			diff[i+1][j+1] += diff[i+1][j] + diff[i][j+1] - diff[i][j]
			v := diff[i+1][j+1]
			if v == 0 {
				zero++
			} else if v < n*2 {
				ans[v-n]++
			}
		}
	}
	for _, v := range ans {
		Fprintln(out, v+zero)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
