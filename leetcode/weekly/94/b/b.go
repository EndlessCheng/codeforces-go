package main

// https://space.bilibili.com/206214
type pair struct{ x, y int }
var dirs = [...]pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 上右下左（顺时针）

func robotSim(commands []int, obstacles [][]int) (ans int) {
	isObstacle := make(map[pair]bool, len(obstacles)) // 预分配空间
	for _, p := range obstacles {
		isObstacle[pair{p[0], p[1]}] = true
	}

	x, y, k := 0, 0, 0
	for _, c := range commands {
		if c < 0 {
			k = (k + c*2 + 7) % 4 // -2 把 k 减一，-1 把 k 加一
			continue
		}
		for ; c > 0 && !isObstacle[pair{x + dirs[k].x, y + dirs[k].y}]; c-- {
			x += dirs[k].x
			y += dirs[k].y
		}
		ans = max(ans, x*x+y*y)
	}
	return
}
