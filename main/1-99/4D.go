package __99

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol4D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	type pair struct{ w, h int }
	env := make([]pair, n+1)
	for i := range env {
		Fscan(in, &env[i].w, &env[i].h)
	}

	next := make([]int, n+1)
	for i := range next {
		next[i] = -1
	}
	dp := make([]int, n+1)
	var dfs func(v int) int
	dfs = func(v int) int {
		if dp[v] == 0 {
			for w, e := range env {
				if env[v].w < e.w && env[v].h < e.h {
					if dpw := dfs(w) + 1; dpw > dp[v] {
						dp[v] = dpw
						next[v] = w
					}
				}
			}
		}
		return dp[v]
	}
	Fprintln(out, dfs(0))
	for v := next[0]; v != -1; v = next[v] {
		Fprint(out, v, " ")
	}
}

//func main() {
//	Sol4D(os.Stdin, os.Stdout)
//}
