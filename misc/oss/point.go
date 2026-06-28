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
	return point{(p.x + q.x + mapSizeN) % mapSizeN, (p.y + q.y + mapSizeM) % mapSizeM, p.z + q.z}
}

func isNeighbor4(p, q point) bool {
	for _, dir := range directions4 {
		if (point{p.x + dir.x, p.y + dir.y, p.z + dir.z}) == q {
			return true
		}
	}
	return false
}

func isNeighbor6(p, q point) bool {
	for _, dir := range directions6 {
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
	// 和遍历 levelMap 的顺序保持一致
	return int(cmp.Or(a.z-b.z, a.x-b.x, a.y-b.y))
}

type pointWithDir struct {
	point
	dir uint8
}

func (mirror *pointWithDir) reflectToAnotherDir(dir point) point {
	revDir := point{-dir.x, -dir.y, -dir.z}
	d0, d1 := directions4[mirror.dir&0xf], directions4[mirror.dir>>4]
	if d0 == revDir {
		return d1
	}
	if d1 == revDir {
		return d0
	}
	return point{}
}

func (mirror *pointWithDir) canReflect(dir point) bool {
	revDir := point{-dir.x, -dir.y, -dir.z}
	return directions4[mirror.dir&0xf] == revDir || directions4[mirror.dir>>4] == revDir
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
	// 和遍历 levelMap 的顺序保持一致
	return int(cmp.Or(a.z-b.z, a.x-b.x, a.y-b.y))
}

// 直接改 rawDir 中的顺序
var directions4 = []point{rawDir[0].p, rawDir[1].p, rawDir[2].p, rawDir[3].p}
var directions6 = append(directions4, point{0, 0, 1}, point{0, 0, -1})
var debugDir4String = []string{rawDir[0].dirZH, rawDir[1].dirZH, rawDir[2].dirZH, rawDir[3].dirZH}
var dir4String = []string{rawDir[0].dirEN, rawDir[1].dirEN, rawDir[2].dirEN, rawDir[3].dirEN}
var noPos = point{-60, -60, -60}
var noPosDir = pointWithDir{noPos, uint8(math.MaxUint8)}
