package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF106C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var maxW, n, a, b int
	Fscan(in, &maxW, &n)
	n++
	weights := make([]int, n)
	values := make([]int, n)
	stocks := make([]int, n)
	Fscan(in, &weights[0], &values[0])
	stocks[0] = maxW / weights[0]
	for i := 1; i < n; i++ {
		Fscan(in, &a, &b, &weights[i], &values[i])
		stocks[i] = a / b
	}

	dp := make([]int, maxW+1)
	for i, num := range stocks {
		v, w := values[i], weights[i]
		for a := 0; a < w; a++ {
			type pair struct{ i, v int }
			q := []pair{}
			for j := 0; j*w+a <= maxW; j++ {
				val := dp[j*w+a] - j*v
				for len(q) > 0 && q[len(q)-1].v <= val {
					q = q[:len(q)-1]
				}
				q = append(q, pair{j, val})
				dp[j*w+a] = q[0].v + j*v
				if q[0].i == j-num {
					q = q[1:]
				}
			}
		}
	}
	Fprint(out, dp[maxW])
}

//func main() { CF106C(os.Stdin, os.Stdout) }
