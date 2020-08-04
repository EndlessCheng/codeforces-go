package __99

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF35D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, maxW int
	Fscan(in, &n, &maxW)
	ws := make([]int, n)
	for i := range ws {
		Fscan(in, &ws[i])
		ws[i] *= n - i
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxW+1)
	}
	for i, wi := range ws {
		for j, dpij := range dp[i] {
			if j < wi {
				dp[i+1][j] = dpij
			} else {
				dp[i+1][j] = max(dpij, dp[i][j-wi]+1)
			}
		}
	}
	Fprint(_w, dp[n][maxW])
}

//func main() { r, _ := os.Open("input.txt"); w, _ := os.Create("output.txt"); CF35D(r, w) }
