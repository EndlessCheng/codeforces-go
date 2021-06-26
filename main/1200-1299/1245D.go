package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1245D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n int
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
	wireCost := make([]int, n)
	for i := range wireCost {
		Fscan(in, &wireCost[i])
	}

	// 构建距离（费用）矩阵
	dis := make([][]int64, n+1)
	for i := range dis {
		dis[i] = make([]int64, n+1)
	}
	for i, c := range cityCost {
		dis[i][n] = int64(c) // 建立虚点 n，将所有城市连到点 n 上，值为在该城市上建立发电站的费用
		dis[n][i] = int64(c)
		for j := i + 1; j < n; j++ {
			// 搭建电线的费用
			dis[i][j] = int64(abs(city[i].x-city[j].x)+abs(city[i].y-city[j].y)) * int64(wireCost[i]+wireCost[j])
			dis[j][i] = dis[i][j]
		}
	}

	totCost := int64(0)
	ansCity := []interface{}{}
	ansWire := []pair{}

	// 由于是稠密图，用 Prim 算法求最小生成树
	type pair2 struct {
		cost int64
		v    int
	}
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
			ansCity = append(ansCity, minCost[v].v+1)
		} else if minCost[v].v == n { // 发电
			ansCity = append(ansCity, v+1)
		} else if minCost[v].v != v { // 连电线，注意避免自环
			ansWire = append(ansWire, pair{minCost[v].v + 1, v + 1})
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
	Fprintln(out, len(ansWire))
	for _, conn := range ansWire {
		Fprintln(out, conn.x, conn.y)
	}
}

//func main() { CF1245D(os.Stdin, os.Stdout) }
