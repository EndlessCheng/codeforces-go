package main

func robot(command string, obstacles [][]int, endX, endY int) (ans bool) {
	path := map[[2]int]bool{{}: true}
	cur := [2]int{}
	for _, c := range command {
		if c == 'R' {
			cur[0]++
		} else {
			cur[1]++
		}
		path[cur] = true
	}
	cx, cy := cur[0], cur[1]

	// 可以推广至能往三个方向移动的情况
	// 四个方向的话就只能枚举障碍物从进入到离开矩形区域的所有可能位置了
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	reduce := func(x, y int) [2]int {
		k := 0
		if abs(cx) + abs(cy) > 0 {
			k = (abs(x) + abs(y)) / (abs(cx) + abs(cy))
		}
		return [2]int{x - k*cx, y - k*cy}
	}
	if !path[reduce(endX, endY)] {
		return
	}
	for _, p := range obstacles {
		if x, y := p[0], p[1]; x >= 0 && x <= endX && y >= 0 && y <= endY && path[reduce(x, y)] {
			return
		}
	}
	return true
}
