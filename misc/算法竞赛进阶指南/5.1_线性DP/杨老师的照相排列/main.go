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

	var k int
	for {
		Fscan(in, &k)
		if k == 0 {
			break
		}
		n := [5]int{}
		for i := 0; i < k; i++ {
			Fscan(in, &n[i])
		}
		dp := make([][][][][]int, n[0]+1)
		for i := range dp {
			dp[i] = make([][][][]int, n[1]+1)
			for j := range dp[i] {
				dp[i][j] = make([][][]int, n[2]+1)
				for k := range dp[i][j] {
					dp[i][j][k] = make([][]int, n[3]+1)
					for l := range dp[i][j][k] {
						dp[i][j][k][l] = make([]int, n[4]+1)
					}
				}
			}
		}
		dp[0][0][0][0][0] = 1
		for i, di := range dp {
			for j, dj := range di {
				for k, dk := range dj {
					for l, dl := range dk {
						for m, v := range dl {
							if i < n[0] {
								dp[i+1][j][k][l][m] += v
							}
							if j < n[1] && i > j {
								dp[i][j+1][k][l][m] += v
							}
							if k < n[2] && j > k {
								dp[i][j][k+1][l][m] += v
							}
							if l < n[3] && k > l {
								dp[i][j][k][l+1][m] += v
							}
							if m < n[4] && l > m {
								dp[i][j][k][l][m+1] += v
							}
						}
					}
				}
			}
		}
		Fprintln(out, dp[n[0]][n[1]][n[2]][n[3]][n[4]])
	}
}

func main() { run(os.Stdin, os.Stdout) }
