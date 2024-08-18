package main

import "math"

// https://space.bilibili.com/206214
func maximumValueSum(board [][]int) int64 {
	m := len(board)
	type pair struct{ x, j int }
	suf := make([][3]pair, m)
	p := [3]pair{} // 最大、次大、第三大
	for i := range p {
		p[i].x = math.MinInt
	}
	update := func(row []int) {
		for j, x := range row {
			if x > p[0].x {
				if p[0].j != j { // 如果相等，仅更新最大
					if p[1].j != j { // 如果相等，仅更新最大和次大
						p[2] = p[1]
					}
					p[1] = p[0]
				}
				p[0] = pair{x, j}
			} else if x > p[1].x && j != p[0].j {
				if p[1].j != j { // 如果相等，仅更新次大
					p[2] = p[1]
				}
				p[1] = pair{x, j}
			} else if x > p[2].x && j != p[0].j && j != p[1].j {
				p[2] = pair{x, j}
			}
		}
	}
	for i := m - 1; i > 1; i-- {
		update(board[i])
		suf[i] = p
	}

	ans := math.MinInt
	for i := range p {
		p[i].x = math.MinInt // 重置，计算 pre
	}
	for i, row := range board[:m-2] {
		update(row)
		for j, x := range board[i+1] { // 第二个车
			for _, p := range p { // 第一个车
				if p.j == j {
					continue
				}
				for _, q := range suf[i+2] { // 第三个车
					if q.j != j && q.j != p.j { // 没有同列的车
						ans = max(ans, p.x+x+q.x)
						break
					}
				}
			}
		}
	}
	return int64(ans)
}
