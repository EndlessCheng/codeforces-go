package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, totCost int
	Fscan(in, &n)
	type pair struct{ x, y int }
	city := make([]pair, n)
	for i := range city {
		Fscan(in, &city[i].x, &city[i].y)
	}
	cityCost := make([]int, n)
	for i := range cityCost {
		Fscan(in, &cityCost[i])
	}
	lineCost := make([]int, n)
	for i := range lineCost {
		Fscan(in, &lineCost[i])
	}

	// 构建距离（费用）矩阵
	dis := make([][]int, n+1)
	for i := range dis {
		dis[i] = make([]int, n+1)
	}
	for i, c := range cityCost {
		dis[i][n] = c // 建立虚点 n，将所有城市连到点 n 上，值为在该城市上建立发电站的费用
		dis[n][i] = c
		for j := i + 1; j < n; j++ {
			// 搭建电线的费用
			dis[i][j] = (abs(city[i].x-city[j].x) + abs(city[i].y-city[j].y)) * (lineCost[i] + lineCost[j])
			dis[j][i] = dis[i][j]
		}
	}

	ansCity := []interface{}{}
	ansConn := []pair{}

	// 由于是稠密图，用 Prim 算法求最小生成树
	type pair2 struct{ cost, city int }
	minCost := make([]pair2, n+1)
	for i := range minCost {
		minCost[i].cost = math.MaxInt64
	}
	minCost[0].cost = 0
	used := make([]bool, n+1)
	for {
		v := -1
		for w, b := range used {
			if !b && (v < 0 || minCost[w].cost < minCost[v].cost) {
				v = w
			}
		}
		if v < 0 {
			break
		}
		used[v] = true
		totCost += minCost[v].cost
		if v == n { // 发电
			ansCity = append(ansCity, minCost[v].city+1)
		} else if minCost[v].city == n { // 发电
			ansCity = append(ansCity, v+1)
		} else if minCost[v].city != v { // 连电线
			ansConn = append(ansConn, pair{minCost[v].city + 1, v + 1})
		}
		for w, d := range dis[v] {
			if d < minCost[w].cost {
				minCost[w] = pair2{d, v}
			}
		}
	}

	Fprintln(out, totCost)
	Fprintln(out, len(ansCity))
	Fprintln(out, ansCity...)
	Fprintln(out, len(ansConn))
	for _, conn := range ansConn {
		Fprintln(out, conn.x, conn.y)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
