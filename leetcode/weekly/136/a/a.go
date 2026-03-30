package main

// https://space.bilibili.com/206214
func isRobotBounded(instructions string) bool {
	z := 0i
	d := 1i // 初始朝北
	for _, c := range instructions {
		if c == 'G' {
			z += d
		} else if c == 'L' {
			d *= 1i
		} else {
			d *= -1i
		}
	}
	return d != 1i || z == 0
}

type pair struct{ x, y int }

var dirs = [...]pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 上右下左（顺时针）

func isRobotBounded2(instructions string) bool {
	x, y, k := 0, 0, 0
	for _, c := range instructions {
		if c == 'G' {
			x += dirs[k].x
			y += dirs[k].y
		} else {
			k = (k + int(c) + 3) % 4
		}
	}
	return k != 0 || x == 0 && y == 0
}
