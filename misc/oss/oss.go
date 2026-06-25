package oss

import (
	"fmt"
	"slices"
	"strings"
)

/*
2026.6.13

《沉星之序》（Order of the Sinking Star）游戏原型：
《Heroes of Sokoban》https://www.puzzlescript.net/play.html?p=6860122
《Heroes of Sokoban II: Monsters》https://www.puzzlescript.net/play.html?p=6910207
《Heroes of Sokoban III: The Bard and The Druid》https://www.puzzlescript.net/play.html?p=7072276
《Mirror Isles》https://alan.draknek.org/games/puzzlescript/mirrors.php
《Skipping Stones To Lonely Homes》https://alan.draknek.org/games/puzzlescript/skipping-stones.php
《PROMESST》https://silverspaceship.com/promesst/
《PROMESST2》https://silverspaceship.com/promesst2/
《ENIGMASH》https://jacklance.github.io/PuzzleScript/play.html?p=cfdcc6e23f1fb3e9de2fd42fafaf4d4c

*/

const loopLevel = false

var levelMap = curMap

// . 空地
// ~ 水
// # 墙
// * / xyz 压力开关
// n / XYZ 压力门
// N 怪物门
// f 终点
type mortalArrType [0]point
type stoneArrType [stoneNumberInit + grassNumberInit]point
type grassArrType [stoneNumberInit + grassNumberInit]point
type goblinArrType [goblinNumberInit]point
type dragonArrType [len(dragonDirInit)]pointWithDir
type mirrorArrType [0]pointWithDir
type data struct {
	curChar int8
	warrior point         // A 推多个对象
	thief   point         // T 拉一个对象
	wizard  point         // W 交换对象
	priest  point         // P 自己以及上下左右无敌
	bard    point         // B 同时移动切比雪夫距离 <= 2 的对象
	druid   point         // D 把草变成石头
	mortal  mortalArrType // @ 普通角色，推一个对象

	// todo 辅助镜子用小写 m
	mirrors mirrorArrType // M
	stones  stoneArrType  // s
	grasses grassArrType  // w
	goblins goblinArrType // g
	dragons dragonArrType // d

	// 门的开闭
	opened            [2]bool
	monsterDoorOpened bool
}

var n, m int8
var initCharNum uint8

func init() {
	n = int8(len(levelMap))
	m = int8(len(levelMap[0]))

	var stoneNum, grassesNum, goblinNum, dragonNum, mirrorNum int
	for _, row := range levelMap {
		for _, ch := range row {
			switch ch {
			case 's':
				stoneNum++
			case 'w':
				grassesNum++
			case 'g':
				goblinNum++
			case 'd':
				dragonNum++
			case 'M':
				mirrorNum++
			default:
				if strings.ContainsRune("ATWPBD@", ch) {
					initCharNum++
				}
			}
		}
	}

	// 检查数组大小是否与 levelMap 匹配
	if stoneNum != stoneNumberInit ||
		grassesNum != grassNumberInit ||
		goblinNum != len(goblinArrType{}) ||
		dragonNum != len(dragonArrType{}) ||
		mirrorNum != len(mirrorArrType{}) {
		panic("没有修改数组大小")
	}
}

func (d *data) allMonstersDied() bool {
	for _, p := range d.goblins {
		if p != noPos {
			return false
		}
	}
	for _, p := range d.dragons {
		if p.point != noPos {
			return false
		}
	}
	return true
}

// 可以用 bitset 优化
func (d *data) getAllCharPos() []point {
	// todo 改成直接计算 data 中各个 point 数组的长度之和
	allChars := make([]point, 0, initCharNum)
	if d.warrior != noPos {
		allChars = append(allChars, d.warrior)
	}
	if d.thief != noPos {
		allChars = append(allChars, d.thief)
	}
	if d.wizard != noPos {
		allChars = append(allChars, d.wizard)
	}
	if d.priest != noPos {
		allChars = append(allChars, d.priest)
	}
	if d.bard != noPos {
		allChars = append(allChars, d.bard)
	}
	if d.druid != noPos {
		allChars = append(allChars, d.druid)
	}
	for _, p := range d.mortal {
		if p != noPos {
			allChars = append(allChars)
		}
	}
	return allChars
}

var doors [2][]point // 不在水下的门（长度至少是 1）
var monsterDoors []point

func (d *data) getAllMovableObjPos() ([]point, []point) {
	chars := d.getAllCharPos()
	objs := chars
	for _, p := range d.mirrors {
		if p.point != noPos {
			objs = append(objs, p.point)
		}
	}
	for _, p := range d.stones {
		if p != noPos {
			objs = append(objs, p)
		}
	}
	for _, p := range d.goblins {
		if p != noPos {
			objs = append(objs, p)
		}
	}
	for _, p := range d.dragons {
		if p.point != noPos {
			objs = append(objs, p.point)
		}
	}
	return chars, objs
}

func (d *data) isValidPos(x, y int8) bool {
	if !(0 <= x && x < n && 0 <= y && y < m) || levelMap[x][y] == '#' { // 墙
		return false
	}
	if slices.Contains(d.grasses[:], point{x, y, 0}) { // 草
		return false
	}
	if !d.monsterDoorOpened && slices.Contains(monsterDoors[:], point{x, y, 0}) { // 怪物门
		return false
	}
	for i, opened := range d.opened {
		if !opened && slices.Contains(doors[i][:], point{x, y, 0}) { // 压力门
			return false
		}
	}
	return true
}

func (d *data) isProtected(char point) bool {
	return char == d.priest || d.priest != noPos && isNeighbor(char, d.priest)
}

// 在水面上且下面没有石头的对象，落入水中
func (d *data) isFallIntoWater(p point) bool {
	// todo 简化处理栏杆的逻辑
	//if !canFallIntoWater && slices.Contains(d.stones[:], p) { 
	//	return false
	//}
	if p.z == -1 || p == noPos {
		return false
	}
	ch := levelMap[p.x][p.y]
	if ch == '~' && !slices.Contains(d.stones[:], point{p.x, p.y, -1}) {
		return true
	}
	// todo 特殊处理机关
	//if 'X' <= ch && ch <= 'Z' {
	//	sw := pos[ch-'A'+'a']
	//	if d.wizard != sw && d.thief != sw && d.stones[0] != sw {
	//		return true
	//	}
	//}
	return false
}

func (d *data) isDie(p point, burnPos []point, isChar bool) bool {
	// 被门压死
	for i, opened := range d.opened {
		if !opened && slices.Contains(doors[i][:], p) {
			return false
		}
	}

	// 淹死
	if d.isFallIntoWater(p) {
		return true
	}

	if isChar && d.isProtected(p) {
		return false
	}

	// 哥布林
	for _, g := range d.goblins {
		if isNeighbor(g, p) {
			return true
		}
	}

	// 喷火龙
	if slices.Contains(burnPos, p) {
		return true
	}

	return false
}

func (d *data) changePos(oldP, newP point) {
	switch {
	case oldP == d.warrior:
		d.warrior = newP
	case oldP == d.thief:
		d.thief = newP
	case oldP == d.wizard:
		d.wizard = newP
	case oldP == d.priest:
		d.priest = newP
	case oldP == d.bard:
		d.bard = newP
	case oldP == d.druid:
		d.druid = newP
	default:
		changed := false

		i := slices.Index(d.mortal[:], oldP)
		if i >= 0 {
			changed = true
			d.stones[i] = newP
		}

		i = slices.Index(d.stones[:], oldP)
		if i >= 0 {
			changed = true
			d.stones[i] = newP
		}

		i = slices.Index(d.goblins[:], oldP)
		if i >= 0 {
			changed = true
			d.goblins[i] = newP
		}

		for i, p := range d.dragons[:] {
			if p.point == oldP {
				changed = true
				d.dragons[i].point = newP
				break
			}
		}

		if !changed {
			panic("没有发生修改，请检查代码")
		}
	}
}

func solveLevel(debug bool) []string {
	mortalInitArr := mortalArrType{}
	for i := range mortalInitArr {
		mortalInitArr[i] = noPos
	}
	mirrorInitArr := mirrorArrType{}
	stoneInitArr := stoneArrType{}
	for i := range stoneInitArr {
		stoneInitArr[i] = noPos
	}
	grassInitArr := grassArrType{}
	for i := range grassInitArr {
		grassInitArr[i] = noPos
	}
	goblinInitArr := goblinArrType{}
	dragonInitArr := dragonArrType{}
	__curChar := initChar
	__warrior := warriorPosInit
	__thief := thiefPosInit
	__wizard := wizardPosInit
	__priest := noPos
	__bard := noPos
	__druid := noPos
	__mortals := mortalInitArr[:0]
	__mirrors := mirrorInitArr[:0]
	__stones := stoneInitArr[:0]
	__grasses := grassInitArr[:0]
	__goblins := goblinInitArr[:0]
	__dragons := dragonInitArr[:0]
	weightSwitches := [len(doors)][]point{}
	finals := []point{}
	for i, row := range levelMap {
		for j, ch := range row {
			p := point{int8(i), int8(j), 0}
			switch ch {
			case 'A':
				if __curChar < 0 {
					__curChar = charWarrior
				}
				if __warrior == noPos {
					__warrior = p
				}
			case 'T':
				if __curChar < 0 {
					__curChar = charThief
				}
				if __thief == noPos {
					__thief = p
				}
			case 'W':
				if __curChar < 0 {
					__curChar = charWizard
				}
				if __wizard == noPos {
					__wizard = p
				}
			case 'P':
				if __curChar < 0 {
					__curChar = charPriest
				}
				if __priest == noPos {
					__priest = p
				}
			case 'B':
				if __curChar < 0 {
					__curChar = charBard
				}
				if __bard == noPos {
					__bard = p
				}
			case 'D':
				if __curChar < 0 {
					__curChar = charDruid
				}
				if __druid == noPos {
					__druid = p
				}
			case '@':
				if __curChar < 0 {
					__curChar = charMortal
				}
				__mortals = append(__mortals, p)
			case 'M':
				__mirrors = append(__mirrors, pointWithDir{p, -1})
			case 's':
				__stones = append(__stones, p)
			case 'w':
				__grasses = append(__grasses, p)
			case 'g':
				__goblins = append(__goblins, p)
			case 'd':
				idx := len(__dragons)
				dir := int8(-1)
				for _i, d := range rawDir {
					if d.dirEN[0] == dragonDirInit[idx] {
						dir = int8(_i)
					}
				}
				__dragons = append(__dragons, pointWithDir{p, dir})
			case '*':
				weightSwitches[0] = append(weightSwitches[0], p)
			case 'n':
				doors[0] = append(doors[0], p)
			case 'x', 'y', 'z', '{':
				weightSwitches[ch-'x'] = append(weightSwitches[ch-'x'], p)
			case 'X', 'Y', 'Z', '[':
				doors[ch-'X'] = append(doors[ch-'X'], p)
			case 'N':
				monsterDoors = append(monsterDoors, p)
			case 'f':
				finals = append(finals, p)
			case '$', '%':
				panic("不支持的符号")
			}
		}
	}

	validChars := []int8{}
	if __warrior != noPos {
		validChars = append(validChars, charWarrior)
	}
	if __thief != noPos {
		validChars = append(validChars, charThief)
	}
	if __wizard != noPos {
		validChars = append(validChars, charWizard)
	}
	if __priest != noPos {
		validChars = append(validChars, charPriest)
	}
	if __bard != noPos {
		validChars = append(validChars, charBard)
	}
	if __druid != noPos {
		validChars = append(validChars, charDruid)
	}
	if len(mortalInitArr) > 0 {
		validChars = append(validChars, charMortal)
	}

	levelData := data{
		curChar: __curChar,
		warrior: __warrior,
		thief:   __thief,
		wizard:  __wizard,
		priest:  __priest,
		bard:    __bard,
		druid:   __druid,
		mortal:  mortalInitArr,

		mirrors: mirrorInitArr,
		stones:  stoneInitArr,
		grasses: grassInitArr,
		goblins: goblinInitArr,
		dragons: dragonInitArr,
	}

	vis := map[data]bool{}
	queue := []data{}
	type pair struct {
		data
		info string
	}
	from := map[data]pair{}
	add := func(last, d data, info string) {
		_, allMovableObjs := d.getAllMovableObjPos()

		if len(weightSwitches[0]) > 0 {
			for i, weightSwitch := range weightSwitches {
				opened := true
				for _, w := range weightSwitch {
					if !slices.Contains(allMovableObjs, w) && !slices.Contains(d.grasses[:], w) { // 草也可以按住地板
						opened = false
						break
					}
				}
				//if opened {
				//	fmt.Printf("门 %d 开启\n", i)
				//}
				d.opened[i] = opened

				// 石头被门压碎（石头在门中，但门没有打开）
				if !opened {
					for j, p := range d.stones {
						if slices.Contains(doors[i], p) {
							if !canBrokenItem {
								return
							}
							d.stones[j] = noPos
						}
					}
				}
			}
		}

		// 被喷火龙攻击到的位置
		var burnedPos []point
		if !d.monsterDoorOpened {
			for _, dra := range d.dragons {
				dir := dir4[dra.dir]
				cur := point{dra.x, dra.y, dra.z}
				for {
					cur.x += dir.x
					cur.y += dir.y
					cur.z += dir.z
					if !d.isValidPos(cur.x, cur.y) {
						break
					}
					if slices.Contains(allMovableObjs, cur) {
						burnedPos = append(burnedPos, cur)
						break
					}
				}
			}
		}

		// 先判断是否有角色死亡
		for _, char := range d.getAllCharPos() {
			if d.isDie(char, burnedPos, true) {
				return // 无效状态
			}
		}

		if !d.monsterDoorOpened {
			// 哥布林
			goblins := d.goblins
			if len(d.goblins) > 0 {
				for i, p := range d.goblins {
					if d.isDie(p, burnedPos, false) {
						if !canBrokenItem {
							return
						}
						goblins[i] = noPos
					}
				}
				slices.SortFunc(goblins[:], cmpPoint)
			}

			// 喷火龙
			dragons := d.dragons
			if len(d.dragons) > 0 {
				for i, p := range d.dragons {
					if d.isDie(p.point, burnedPos, false) {
						if !canBrokenItem {
							return
						}
						dragons[i].point = noPos
					}
				}
				slices.SortFunc(dragons[:], cmpPointWithDir)
			}

			if canBrokenItem {
				d.goblins = goblins
				d.dragons = dragons
				d.monsterDoorOpened = d.allMonstersDied()
			}
		}

		// 镜子
		if len(d.mirrors) > 0 {
			mir := d.mirrors[:]
			for i, p := range mir {
				if d.isFallIntoWater(p.point) {
					mir[i].z = -1
				}
			}
			slices.SortFunc(mir, cmpPointWithDir)
		}

		// 石头
		if len(d.stones) > 0 {
			sto := d.stones[:]
			for i, p := range sto {
				if d.isFallIntoWater(p) {
					sto[i].z = -1
				}
			}
			slices.SortFunc(sto, cmpPoint)
		}

		// 草
		if len(d.grasses) > 0 {
			slices.SortFunc(d.grasses[:], cmpPoint)
		}

		if !vis[d] {
			vis[d] = true
			queue = append(queue, d)
			from[d] = pair{last, info}
		}
	}

	info0 := ""
	if debug {
		info0 = "init"
	}
	add(data{}, levelData, info0)

	for len(queue) > 0 {
		// 注意入队的时候修改了物品的位置（落水）
		d := queue[0]
		queue = queue[1:]

		allChars, allMovableObjs := d.getAllMovableObjPos()

		var pass bool
		if !targetIsClearAllMonsters {
			// 标准版：所有人都到达终点
			slices.SortFunc(allChars, cmpPoint)
			pass = slices.Equal(allChars, finals)
		} else {
			// 简化版：门开启
			pass = d.monsterDoorOpened // d.opened
		}
		if pass {
			path := []string{}
			for d != (data{}) {
				var ok bool
				p, ok := from[d]
				if !ok {
					panic("代码修改了 d，与存入的 d 不符")
				}
				path = append(path, p.info)
				d = p.data
			}
			slices.Reverse(path)
			return path
		}

		// 移动当前角色
		switch d.curChar {
		case charDefault:
			panic("代码有误，当前角色不能为 charDefault")
		case charWarrior:
			// 普通移动一步
			p0 := d.warrior
			for dIdx, dir := range dir4 {
				x, y, z := p0.x+dir.x, p0.y+dir.y, p0.z+dir.z
				// 该方向有多少个连续的对象
				cnt := 0
				cur := point{x, y, z}
				for slices.Contains(allMovableObjs, cur) {
					cnt++
					cur.x += dir.x
					cur.y += dir.y
					cur.z += dir.z
				}
				// 前面是否有空地
				if !d.isValidPos(cur.x, cur.y) {
					continue // 枚举另一个方向
				}

				newData := d
				for range cnt {
					nxt := point{cur.x - dir.x, cur.y - dir.y, cur.z - dir.z}
					newData.changePos(nxt, cur)
					cur = nxt
				}
				np := point{x, y, z}
				newData.warrior = np
				var info string
				if debug {
					info = fmt.Sprintf("士 %s", debugDirString[dIdx])
				} else {
					info = dirString[dIdx]
				}
				add(d, newData, info)
			}
		case charThief:
			// 普通移动一步
			p0 := d.thief
			for dIdx, dir := range dir4 {
				x, y, z := p0.x+dir.x, p0.y+dir.y, p0.z+dir.z
				np := point{x, y, z}
				if !d.isValidPos(np.x, np.y) || slices.Contains(allMovableObjs, np) {
					continue // 枚举另一个方向
				}
				newData := d
				back := point{p0.x - dir.x, p0.y - dir.y, p0.z - dir.z}
				if slices.Contains(allMovableObjs, back) {
					// 拉人/物 -> 当前位置
					newData.changePos(back, p0)
				}
				newData.thief = np
				var info string
				if debug {
					info = fmt.Sprintf("贼 %s", debugDirString[dIdx])
				} else {
					info = dirString[dIdx]
				}
				add(d, newData, info)
			}
		case charWizard:
			p0 := d.wizard
		nextDir:
			for dIdx, dir := range dir4 {
				// 该方向上是否有人/物
				x, y, z := p0.x, p0.y, p0.z
				for {
					x += dir.x
					y += dir.y
					z += dir.z
					// 出界或者有障碍物
					np := point{x, y, z}
					if !d.isValidPos(np.x, np.y) {
						break
					}
					if !slices.Contains(allMovableObjs, np) { // 空地
						continue
					}
					// 和对象交换位置
					newData := d
					newData.changePos(np, p0) // 那个位置的对象换到 p0
					newData.wizard = np
					var info string
					if debug {
						info = fmt.Sprintf("法 %s 交换", debugDirString[dIdx])
						//fmt.Println(info)
					} else {
						info = dirString[dIdx]
					}
					add(d, newData, info)
					continue nextDir
				}

				// 没有，那就普通移动一步
				np := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				if !d.isValidPos(np.x, np.y) {
					continue
				}
				newData := d
				newData.wizard = np
				var info string
				if debug {
					info = fmt.Sprintf("法 %s", debugDirString[dIdx])
					//fmt.Println(info)
				} else {
					info = dirString[dIdx]
				}
				add(d, newData, info)
			}
		case charPriest:
			// 普通移动一步
			p0 := d.priest
			for dIdx, dir := range dir4 {
				np := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				if !d.isValidPos(np.x, np.y) || slices.Contains(allMovableObjs, np) {
					continue // 枚举另一个方向
				}
				newData := d
				newData.priest = np
				var info string
				if debug {
					info = fmt.Sprintf("牧 %s", debugDirString[dIdx])
				} else {
					info = dirString[dIdx]
				}
				add(d, newData, info)
			}
		case charBard:
			p0 := d.bard
			items := []point{}
			for _, p := range allMovableObjs {
				if chebyshevDis(p, p0) <= 2 {
					items = append(items, p)
				}
			}

			// todo 如果踩在物品上，可以多走一格（前提是撞墙或者前面一个格子也可以踩）

			// 普通移动一步
			// 切比雪夫距离 <= 2 的物品（包括自己）都移动一步
			for dIdx, dir := range dir4 {
				x, y, z := p0.x+dir.x, p0.y+dir.y, p0.z+dir.z
				if !d.isValidPos(x, y) {
					continue
				}
				slices.SortFunc(items, func(a, b point) int {
					if dir.x != 0 {
						return int(b.x*dir.x - a.x*dir.x)
					}
					return int(b.y*dir.y - a.y*dir.y)
				})

				newData := d
				unmovedItems := []point{}
				for _, item := range items {
					// item 往前移动一格
					nxtPos := point{item.x + dir.x, item.y + dir.y, item.z + dir.z}
					if !d.isValidPos(nxtPos.x, nxtPos.y) { // 无法移动
						unmovedItems = append(unmovedItems, item)
						continue
					}
					// 可以移动
					if chebyshevDis(nxtPos, p0) > 2 { // item 是力场最前面的点
						if slices.Contains(allMovableObjs, nxtPos) { // 不能与力场外的对象碰撞
							unmovedItems = append(unmovedItems, item)
							continue
						}
					} else if slices.Contains(unmovedItems, nxtPos) { // 力场后面的点，不能与前面移动失败的对象碰撞
						unmovedItems = append(unmovedItems, item)
						continue
					}
					newData.changePos(item, nxtPos)
				}

				if !slices.Contains(unmovedItems, p0) {
					np := point{x, y, z}
					if newData.bard != np {
						panic("移动错误")
					}
					var info string
					if debug {
						info = fmt.Sprintf("诗人 %s", debugDirString[dIdx])
					} else {
						info = dirString[dIdx]
					}
					add(d, newData, info)
				}
			}
		case charDruid:
			p0 := d.druid
			for dIdx, dir := range dir4 {
				np := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				// todo 目前只实现了草 <-> 石头的逻辑
				if i := slices.Index(d.grasses[:], np); i >= 0 {
					newData := d
					newData.stones[0] = newData.grasses[i] // 草变石 todo
					newData.grasses[i] = noPos
					var info string
					if debug {
						info = fmt.Sprintf("德 %s 草变石", debugDirString[dIdx])
					} else {
						info = dirString[dIdx]
					}
					add(d, newData, info)
					continue
				}
				if i := slices.Index(d.stones[:], np); i >= 0 {
					newData := d
					newData.grasses[0] = newData.stones[i] // 石变草 todo
					newData.stones[i] = noPos
					var info string
					if debug {
						info = fmt.Sprintf("德 %s 石变草", debugDirString[dIdx])
					} else {
						info = dirString[dIdx]
					}
					add(d, newData, info)
					continue
				}
				if !d.isValidPos(np.x, np.y) || slices.Contains(allMovableObjs, np) {
					continue // 枚举另一个方向
				}
				// 普通移动一步
				newData := d
				newData.druid = np
				var info string
				if debug {
					info = fmt.Sprintf("德 %s", debugDirString[dIdx])
				} else {
					info = dirString[dIdx]
				}
				add(d, newData, info)
			}
		case charMortal:
			// 多控、普通移动一步
			// todo
		}

		// 换成其他人
		for _, char := range validChars {
			if char != d.curChar {
				newData := d
				newData.curChar = char
				var info string
				if !debug {
					if len(allChars) > 2 {
						info = digits[char : char+1]
					} else {
						info = "c"
					}
				} else {
					//fmt.Println("换人")
				}
				add(d, newData, info)
			}
		}

		// 镜子反射对象
		if len(d.mirrors) > 0 {
			// todo
		}
	}

	// 无解
	return nil
}

/* 镜子、多控

"##.#####...####",
"#...####.M..###",
"#.@..##.......#",
"....###........",
"..@.##.........",
"....##....n....",
"#..####........",
"#######...M...#",
"########......#",
"#########...###",
右上X 上右下右下左 右下左上X
上上右

*/
func heroesOfSokobanMirrors() []string {
	type point struct{ x, y int }
	dir4 := []point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	less := func(a, b point) bool {
		return a.x < b.x || a.x == b.x && a.y < b.y
	}

	type mirror struct {
		point
		tp int
	}
	mirrorDirs := [...]map[point]point{
		{{0, 1}: {1, 0}, {-1, 0}: {0, -1}}, // \ 方向，镜面朝下
		{{1, 0}: {0, 1}, {0, -1}: {-1, 0}}, // \ 方向，镜面朝上
	}

	type data struct {
		man     [2]point
		mirrors [2]mirror
		done    bool
	}
	vis := map[data]bool{}
	Q := []data{}
	type pair struct {
		data
		s string
	}
	from := map[data]pair{}
	add := func(fr, d data, s string) {
		if !less(d.man[0], d.man[1]) {
			d.man[0], d.man[1] = d.man[1], d.man[0]
		}
		if !less(d.mirrors[0].point, d.mirrors[1].point) {
			d.mirrors[0].point, d.mirrors[1].point = d.mirrors[1].point, d.mirrors[0].point
		}

		if !vis[d] {
			vis[d] = true
			Q = append(Q, d)
			from[d] = pair{fr, s}
		}
	}

	levelMap := []string{
		"##.#####...####",
		"#...####.M..###",
		"#.@..##.......#",
		"....###........",
		"..@.##.........",
		"....##....n....",
		"#..####........",
		"#######...M...#",
		"########......#",
		"#########...###",
	}

	final := point{}
	for i, row := range levelMap {
		for j, b := range row {
			p := point{i, j}
			if b == 'n' {
				final = p
			}
		}
	}

	n := len(levelMap)
	m := len(levelMap[0])
	isValidPos := func(x, y int) bool {
		// todo 目前只看不是 '#'
		return 0 <= x && x < n && 0 <= y && y < m && levelMap[x][y] != '#'
	}

	var man, mirrors []point
	for i, row := range levelMap {
		for j, b := range row {
			p := point{i, j}
			if b == '@' {
				man = append(man, p)
			} else if b == 'M' {
				mirrors = append(mirrors, p)
			}
		}
	}
	noPos := point{-60, -60}
	levelData := data{
		man:     [2]point{man[0], man[1]},
		mirrors: [2]mirror{{mirrors[0], 0}, {mirrors[1], 1}},
	}

	add(data{}, levelData, "")

	for {
		d := Q[0]
		Q = Q[1:]
		mn := d.man
		done := d.done
		if !done && mn[0] == mn[1] {
			mn[0] = noPos
			done = true
		}

		// 检查两人是否重合
		if done && mn[1] == final {
			path := []string{}
			for d != (data{}) {
				path = append(path, from[d].s)
				d = from[d].data
			}
			slices.Reverse(path)
			return path
		}

	nextDir1:
		for dIdx, dir := range dir4 {
			newMan := mn
			newMirrors := d.mirrors
			pushed := false
		o:
			for idx, p := range mn {
				if p == noPos {
					continue
				}
				x, y := p.x+dir.x, p.y+dir.y
				np := point{x, y}
				if !isValidPos(x, y) { // 人出界
					continue
				}
				for mi, mr := range d.mirrors {
					if mr.point != np {
						continue
					}
					// 推镜子
					nmr := mr.point
					nmr.x += dir.x
					nmr.y += dir.y
					if !isValidPos(x, y) { // 镜子入水
						// 禁止这种操作
						continue nextDir1
					}
					if nmr == d.mirrors[mi^1].point { // 不能连续推多个镜子
						continue o
					}
					pushed = true
					newMirrors[mi].point = nmr // 推镜子，人也移动一步
					break
				}
				newMan[idx] = np
			}

			// 碰撞检测
			valid := newMan[0] != newMan[1]

			if valid {
				info := "人移动"
				if pushed {
					info = "推镜子"
				}
				s := fmt.Sprintf("%s %s", info, debugDirString[dIdx])
				//fmt.Println(s)
				add(d, data{newMan, newMirrors, done}, s)
			}
		}

		// X 键：以人/物为主体，首先算出其到镜子的距离 d，然后让光路走恰好 d 步，反射期间遇到其他镜子就继续反射
		// 如果光路遇到其他非镜子物品，则禁止反射
		newMan := mn
		validSwap := true
		swapped := [2]bool{} // 标记是否换过

	outer:
		for idx, p := range mn {
			if p == noPos {
				continue
			}
		nextDir:
			for _, dir := range dir4 {
				x, y := p.x, p.y
				for step := 1; ; step++ {
					x += dir.x
					y += dir.y

					// 没有任何物品
					if !(0 <= x && x < n && 0 <= y && y < m) {
						break
					}
					// 遇到另一个非镜子物品，不会触发交换
					np := point{x, y}
					if np == mn[idx^1] {
						// todo 这里是遇到另一个人
						break
					}

					// 遇到镜子，但需要保证镜子朝向是对的
					for mi, mr := range d.mirrors {
						if mr.point != np {
							continue
						}
						mDir, ok := mirrorDirs[mr.tp][dir]
						if !ok {
							continue
						}

						// 反射！从 mr.point 出发，往 mDir 方向走 step 步
						nx, ny := mr.x, mr.y
						for range step {
							nx += mDir.x
							nx += mDir.y
							// 出界、遇到，或者光路有其他人/物，挡住了，无法交换
							if (point{nx, ny}) == mn[idx^1] {
								validSwap = false
								break outer
							}
							// 改变光路
							if (point{nx, ny}) == d.mirrors[mi^1].point {
								mDir, ok = mirrorDirs[d.mirrors[mi^1].tp][dir]
								if !ok { // 镜子朝向不对
									// 挡住了
									validSwap = false
									break outer
								}
							}
						}

						// 落脚点出界，或者在水上
						if !isValidPos(nx, ny) {
							validSwap = false
							break outer
						}

						np = point{nx, ny}
						// 目标点不能有任何物品，包括镜子
						// 由于上面检查了人/物，这里只需检查镜子
						for _, tmp := range d.mirrors {
							if tmp.point == np {
								// 挡住了
								validSwap = false
								break outer
							}
						}

						// 可以反射！
						if swapped[idx] {
							// todo 这里禁止分身
							validSwap = false
							break outer
						}
						swapped[idx] = true
						newMan[idx] = np
						continue nextDir
					}
				}
			}
		}

		if validSwap && (swapped[0] || swapped[1]) {
			s := fmt.Sprint("反射 ", newMan)
			//fmt.Println(s)
			add(d, data{newMan, d.mirrors, done}, s)
		}
	}
}

const digits = "0123456789"

const (
	charDefault = iota // 仅占位，不使用
	charWarrior
	charThief
	charWizard
	charPriest
	charBard
	charDruid
	charMortal
)

//charName := [...]string{
//	charWarrior: "战士",
//	charThief:   "盗贼",
//	charWizard:  "法师",
//	charPriest:  "牧师",
//	charBard:    "诗人",
//	charDruid:   "德鲁伊",
//}
//mpChar := [...]int{
//	'A': charWarrior,
//	'T': charThief,
//	'W': charWizard,
//	'P': charPriest,
//	'B': charBard,
//	'D': charDruid,
//	'@': charMortal,
//}
//_ = mpChar
