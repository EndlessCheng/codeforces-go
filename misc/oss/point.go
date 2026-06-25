package oss

import "cmp"

type point struct {
	x, y, z int8
}

type pointWithDir struct {
	point
	dir int8
}

func (p point) add(q point) point {
	if !loopLevel {
		return point{p.x + q.x, p.y + q.y, p.z + q.z}
	}
	return point{(p.x + q.x + n) % n, (p.y + q.y + m) % m, p.z + q.z}
}

func isNeighbor(p, q point) bool {
	for _, dir := range dir4 {
		if (point{p.x + dir.x, p.y + dir.y, p.z + dir.z}) == q {
			return true
		}
	}
	return false
}

// 吟游诗人
func chebyshevDis(p, q point) int {
	return int(max(abs(p.x-q.x), abs(p.y-q.y), abs(p.z-q.z)))
}

//func manhattanDis(p, q point) int {
//	return int(abs(p.x-q.x) + abs(p.y-q.y) + abs(p.z-q.z))
//}

var dir4 = []point{{0, 1, 0}, {-1, 0, 0}, {1, 0, 0}, {0, -1, 0}}
var debugDirString = []string{"右", "上", "下", "左"}
var dirString = []string{"d", "w", "s", "a"}

var noPos = point{-60, -60, -60}

func cmpPoint(a, b point) int {
	return int(cmp.Or(a.x-b.x, a.y-b.y, a.z-b.z))
}

func cmpPointWithDir(a, b pointWithDir) int {
	return int(cmp.Or(a.x-b.x, a.y-b.y, a.z-b.z))
}

//func sortPoints(a ...point) []point {
//	slices.SortFunc(a, cmpPoint)
//	return a
//}
