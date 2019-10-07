package main

import (
	"bufio"
	. "fmt"
	"io"
)

func dfs1214D(cell [][]byte, i, j int) bool {
	switch cell[i][j] {
	case '#':
		return false
	case '$':
		return true
	}
	cell[i][j] = '#'
	return j+1 < len(cell[i]) && dfs1214D(cell, i, j+1) ||
		i+1 < len(cell) && dfs1214D(cell, i+1, j)
}

// github.com/EndlessCheng/codeforces-go
func Sol1214D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	cell := make([][]byte, n)
	for i := range cell {
		Fscan(in, &cell[i])
	}

	ans := 0
	cell[n-1][m-1] = '$'
	if dfs1214D(cell, 0, 0) {
		ans++
		cell[0][0] = '.'
		if dfs1214D(cell, 0, 0) {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1214D(os.Stdin, os.Stdout)
//}
