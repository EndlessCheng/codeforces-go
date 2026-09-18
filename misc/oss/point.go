package oss

import (
	"cmp"
	"fmt"
	"math"
)

var rawDir = [...]struct {
	p     point
	dirZH string
	dirEN string
}{
	// 由于屏幕是列数更多，所以更多的是左右移动，要左右移动优先
	{point{0, -1, 0}, "左", "a"},
	{point{0, 1, 0}, "右", "d"},

	{point{-1, 0, 0}, "上", "w"},
	{point{1, 0, 0}, "下", "s"},

	{point{0, 0, -1}, "落", "n"}, // down
	{point{0, 0, 1}, "升", "u"},  // up
}

// todo
var defaultMirrorDirs = [4]uint8{
	makeMirrorDir("aw"),
	makeMirrorDir("sd"),
	makeMirrorDir("wd"),
	makeMirrorDir("as"),
}

func makeMirrorDir(s string) uint8 {
	dir0 := getDir(s[0])
	dir1 := getDir(s[1])
	if dir0 > dir1 {
		dir0, dir1 = dir1, dir0
	}
	// 低小，高大
	return dir1<<4 | dir0
}

// 0 -> '/'
// 1 -> '\'
var mirrorReflectDragonMapping [2][4]uint8

func init() {
	// '/' 左上互换，右下互换
	for i, s := range rawDir[:4] {
		p := s.p
		p.x, p.y = p.y, p.x
		// 找等于 p 的
		for j, t := range rawDir[:4] {
			if t.p == p {
				mirrorReflectDragonMapping[0][i] = uint8(j)
				break
			}
		}
	}

	// '\' 右上互换，左下互换
	for i, s := range rawDir[:4] {
		p := s.p
		p.x, p.y = -p.y, -p.x
		// 找等于 p 的
		for j, t := range rawDir[:4] {
			if t.p == p {
				mirrorReflectDragonMapping[1][i] = uint8(j)
				break
			}
		}
	}
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

func getDirIndexByDir(dir point) uint8 {
	for i, d := range directions6 {
		if d == dir {
			return uint8(i)
		}
	}
	panic(fmt.Sprintf("没找到 dir %v", dir))
}

type point struct {
	x, y, z int8
}

// todo 全部替换成 add
func (p point) add(q point) point {
	if !isLevelLooped {
		return point{p.x + q.x, p.y + q.y, p.z + q.z}
	}
	// todo 取模很慢，换成 if
	return point{(p.x + q.x + mapSizeN) % mapSizeN, (p.y + q.y + mapSizeM) % mapSizeM, p.z + q.z}
}

func (p point) sub(q point) point {
	if !isLevelLooped {
		return point{p.x - q.x, p.y - q.y, p.z - q.z}
	}
	return point{(p.x - q.x + mapSizeN) % mapSizeN, (p.y - q.y + mapSizeM) % mapSizeM, p.z - q.z}
}

func (p point) rev() point {
	return point{-p.x, -p.y, -p.z}
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

func cmpPoint(a, b point) int {
	// 和遍历 levelMap 的顺序保持一致
	return int(cmp.Or(a.z-b.z, a.x-b.x, a.y-b.y))
}

type pointWithDir struct {
	point
	// 镜子：高 4 位和低 4 位分别保存两个方向（0 ~ 5）   低位保存小的 dir-index，高位保存大的 dir-index
	// 光束：高 4 位是类型，低 4 位是方向
	// 其他：dir-index
	dir uint8
}

func (mirror *pointWithDir) reflectToAnotherDir(dir point) point {
	revDir := dir.rev()
	d0, d1 := directions6[mirror.dir&0xf], directions6[mirror.dir>>4]
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
	return directions6[mirror.dir&0xf] == revDir || directions6[mirror.dir>>4] == revDir
}

// 把 refMirror 反射一下
func (mirror *pointWithDir) reflectMirrorRef(mirrorRefDir uint8) uint8 {
	// todo dir6
	// 同向 or 面对面
	if mirrorRefDir == mirror.dir || mirrorRefDir^0x11 == mirror.dir {
		return mirrorRefDir
	}

	// 垂直，翻转镜子即可
	return mirrorRefDir ^ 0x11 // (ref1^1)<<4 | (ref0 ^ 1)
}

func (mirror *pointWithDir) reflectDragon(dragonDir uint8) uint8 {
	// todo dir6
	d0, d1 := directions4[mirror.dir&0xf], directions4[mirror.dir>>4]
	if d0.x == 0 {
		d0, d1 = d1, d0
	}
	if d0.x > 0 == (d1.y > 0) {
		// '/'
		return mirrorReflectDragonMapping[0][dragonDir]
	} else {
		// '\'
		return mirrorReflectDragonMapping[1][dragonDir]
	}
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

// 语法糖，可以用负数位置表示倒数行/列（相对 levelMap）
func changeNegPoint(p point) point {
	if p.x < 0 {
		p.x += mapSizeN
	}
	if p.y < 0 {
		p.y += mapSizeM
	}
	// 请勿调整负数 z
	return p
}

var noPos = point{-60, -60, -60}
var noPosDir = pointWithDir{noPos, dirStop}

// 下面代码勿动，请修改 rawDir 中的顺序
var directions6 = []point{rawDir[0].p, rawDir[1].p, rawDir[2].p, rawDir[3].p, rawDir[4].p, rawDir[5].p}
var directions4 = directions6[:4]
var dir4String = []string{rawDir[0].dirEN, rawDir[1].dirEN, rawDir[2].dirEN, rawDir[3].dirEN}
var debugDir4String = []string{rawDir[0].dirZH, rawDir[1].dirZH, rawDir[2].dirZH, rawDir[3].dirZH}
var flowDirMapping = [...]uint8{
	'^': getDirIndexByDir(point{-1, 0, 0}),
	'v': getDirIndexByDir(point{1, 0, 0}),
	'<': getDirIndexByDir(point{0, -1, 0}),
	'>': getDirIndexByDir(point{0, 1, 0}),
}
