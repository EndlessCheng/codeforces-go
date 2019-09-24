package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func Sol1133E(reader io.Reader, writer io.Writer) {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	const maxTeamDelta = 5

	var n, k int
	Fscan(in, &n, &k)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &arr[i])
	}
	sort.Ints(arr)

	deltaPos := make([][maxTeamDelta + 1]int, n+1)
	for j := 0; j <= maxTeamDelta; j++ {
		deltaPos[1][j] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j <= maxTeamDelta; j++ {
			deltaPos[i][j] = i
		}
		if delta := arr[i] - arr[i-1]; delta <= maxTeamDelta {
			copy(deltaPos[i][delta:], deltaPos[i-1][:maxTeamDelta+1-delta])
		}
	}

	var dp [2][5001]int
	for j := 1; j <= n; j++ {
		dp[0][j] = max(dp[0][j-1], j-deltaPos[j][maxTeamDelta]+1)
	}
	row := 0
	for i := 2; i <= k; i++ {
		row ^= 1
		for j := i; j <= n; j++ {
			posL := deltaPos[j][maxTeamDelta]
			// max(not count student j, count student j)
			dp[row][j] = max(dp[row][j-1], dp[row^1][posL-1]+j-posL+1)
		}
	}
	Fprintln(out, max(dp[row][n], dp[row^1][n]))
}

func main() {
	Sol1133E(os.Stdin, os.Stdout)
}
