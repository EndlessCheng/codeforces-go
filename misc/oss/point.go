package oss

import (
	"cmp"
	"math"
)

var rawDir = [...]struct {
	p     point
	dirZH string
	dirEN string
}{
	{point{0, -1, 0}, "左", "a"},
	{point{0, 1, 0}, "右", "d"},
	{point{-1, 0, 0}, "上", "w"},
	{point{1, 0, 0}, "下", "s"},
}

func getDir(en byte) uint8 {
	dir := uint8(math.MaxUint8)
	for i, d := range rawDir {
		if d.dirEN[0] == en {
			dir = uint8(i)
			break
		}
	}
	if dir == math.MaxUint8 {
		panic("xxxDirInit 打错了")
	}
	return dir
}

type point struct {
	x, y, z int8
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

func cmpPoint(a, b point) int {
	return int(cmp.Or(a.x-b.x, a.y-b.y, a.z-b.z))
}

type pointWithDir struct {
	point
	dir uint8
}

func reflectToDir(dir uint8, v point) point {
	rev := point{-v.x, -v.y, -v.z}
	d0, d1 := dir4[dir&0xf], dir4[dir>>4]
	if d0 == rev {
		return d1
	}
	if d1 == rev {
		return d0
	}
	return point{}
}

func pdContains(a []pointWithDir, p point) bool {
	for _, pd := range a {
		if pd.point == p {
			return true
		}
	}
	return false
}

func pdIndex(a []pointWithDir, p point) int {
	for i, pd := range a {
		if pd.point == p {
			return i
		}
	}
	return -1
}

func cmpPointWithDir(a, b pointWithDir) int {
	return int(cmp.Or(a.x-b.x, a.y-b.y, a.z-b.z))
}

//func sortPoints(a ...point) []point {
//	slices.SortFunc(a, cmpPoint)
//	return a
//}

// 直接改 rawDir 中的顺序
var dir4 = []point{rawDir[0].p, rawDir[1].p, rawDir[2].p, rawDir[3].p}
var debugDirString = []string{rawDir[0].dirZH, rawDir[1].dirZH, rawDir[2].dirZH, rawDir[3].dirZH}
var dirString = []string{rawDir[0].dirEN, rawDir[1].dirEN, rawDir[2].dirEN, rawDir[3].dirEN}
var noPos = point{-60, -60, -60}
