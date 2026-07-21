package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var m, a, b int
	Fscan(in, &m, &a, &b)
	vis := make([][]int8, m)
	for i := range vis {
		vis[i] = make([]int8, m)
	}

	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if x == 0 || y == 0 {
			vis[x][y] = -1
			return false
		}
		if vis[x][y] != 0 {
			return vis[x][y] > 0
		}
		vis[x][y] = 1
		if !dfs(y, (a*y+b*x)%m) {
			vis[x][y] = -1
			return false
		}
		return true
	}

	ans := 0
	for x := range m {
		for y := range m {
			if dfs(x, y) {
				ans++
			}
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
