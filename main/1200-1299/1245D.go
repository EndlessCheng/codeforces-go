package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1245D(reader io.Reader, writer io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	type pair struct{ x, y int }
	ps := make([]pair, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
	}
	costCity := make([]int, n)
	for i := range costCity {
		Fscan(in, &costCity[i])
	}
	costLine := make([]int, n)
	for i := range costLine {
		Fscan(in, &costLine[i])
	}

	dist := [2001][2001]int64{}
	for i := 0; i < n; i++ {
		dist[i][n] = int64(costCity[i])
		dist[n][i] = dist[i][n]
		for j := i + 1; j < n; j++ {
			dist[i][j] = int64(abs(ps[i].x-ps[j].x)+abs(ps[i].y-ps[j].y)) * int64(costLine[i]+costLine[j])
			dist[j][i] = dist[i][j]
		}
	}

	ans := int64(0)
	ansCity := []interface{}{}
	ansConn := []pair{}
	type distCity struct {
		d int64
		v int
	}
	minCost := make([]distCity, n+1)
	for i := range minCost {
		minCost[i].d = int64(1e18)
	}
	minCost[0].d = 0
	used := make([]bool, n+1)
	for {
		v := -1
		for i := 0; i <= n; i++ {
			if !used[i] && (v == -1 || minCost[i].d < minCost[v].d) {
				v = i
			}
		}
		if v == -1 {
			break
		}
		used[v] = true
		if minCost[v].d > 0 {
			ans += minCost[v].d
			if v == n {
				ansCity = append(ansCity, minCost[v].v+1)
			} else if minCost[v].v == n {
				ansCity = append(ansCity, v+1)
			} else {
				ansConn = append(ansConn, pair{minCost[v].v + 1, v + 1})
			}
		}
		for w := 0; w <= n; w++ {
			if dist[v][w] < minCost[w].d {
				minCost[w] = distCity{dist[v][w], v}
			}
		}
	}
	Fprintln(out, ans)
	Fprintln(out, len(ansCity))
	Fprintln(out, ansCity...)
	Fprintln(out, len(ansConn))
	for _, p := range ansConn {
		Fprintln(out, p.x, p.y)
	}
}

//func main() {
//	Sol1245D(os.Stdin, os.Stdout)
//}
