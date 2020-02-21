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

	normal := func(x, y int) [2]int {
		k := (x + y) / len(command)
		return [2]int{x - k*cur[0], y - k*cur[1]}
	}
	if !path[normal(endX, endY)] {
		return
	}
	for _, p := range obstacles {
		if x, y := p[0], p[1]; x >= 0 && x <= endX && y >= 0 && y <= endY && path[normal(x, y)] {
			return
		}
	}
	return true
}
