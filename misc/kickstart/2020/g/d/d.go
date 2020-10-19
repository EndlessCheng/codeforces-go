package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// todo 也可以用调和级数做，那样是线性的
	dp := make([][]float64, 5000)
	dp[1] = []float64{1, 1}
	for i := 2; i < 5000; i++ {
		dp[i] = make([]float64, i+1)
		dp[i][0] = dp[i-1][0] + 1/float64(i)
		dp[i][i] = dp[i][0]
		for j := 1; j <= i-j; j++ {
			dp[i][j] = (dp[i-1][j-1]*float64(j) + float64(i-j)*dp[i-1][j] + 2) / float64(i)
			dp[i][i-j] = dp[i][j]
		}
	}

	solve := func(Case int) {
		var n, v int
		Fscan(in, &n)
		ans := 0.0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			ans += float64(v) * dp[n-1][i]
		}
		Fprintf(out, "%.8f\n", ans)
	}

	var t int
	Fscan(in, &t)
	for Case := 1; Case <= t; Case++ {
		Fprintf(out, "Case #%d: ", Case)
		solve(Case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
