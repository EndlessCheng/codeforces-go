package oss

import (
	"fmt"
	"slices"
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

type merchantArrType [merchantNumberInit]point
type stoneArrType [stoneNumberInit + grassNumberInit]point
type stoneFloatArrType [stoneFloatNumberInit]point
type grassArrType [stoneNumberInit + grassNumberInit]point
type goblinArrType [goblinNumberInit]point
type dragonArrType [len(dragonDirInit)]pointWithDir
type mirrorArrType [len(mirrorDirInit) / 2]pointWithDir
type mirrorRefArrType [len(mirrorRefDirInit) / 2]pointWithDir
type mirrorAuxArrType [len(mirrorAuxDirInit) / 2]pointWithDir

type data struct {
	curChar  int8
	warrior  point           // A 推多个对象
	thief    point           // T 拉一个对象
	wizard   point           // W 交换对象
	cleric   point           // C 自己以及上下左右无敌
	bard     point           // B 同时移动切比雪夫距离 <= 2 的对象
	druid    point           // D 把对象变成石头
	merchant merchantArrType // 9 普通角色，推一个对象

	stones stoneArrType // s
	// 用大写 S 表示可以反射的石头？
	stoneFloats stoneFloatArrType // F todo

	grass grassArrType // w

	mirrors     mirrorArrType    // M
	mirrorRefs  mirrorRefArrType // R 主镜子 + 可以被反射
	mirrorAuxes mirrorAuxArrType // m

	beams struct{} // todo

	goblins goblinArrType // g
	dragons dragonArrType // d

	// 门的开闭
	doorOpened        [doorTypes]bool
	monsterDoorOpened bool
}

var n, m int8
var initCharNum uint8

func init() {
	n = int8(len(levelMap))
	m = int8(len(levelMap[0]))

	var stoneNum, grassesNum, goblinNum, dragonNum, mirrorNum, mirrorRefNum, mirrorAuxNum, merchantNum int
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
			case 'R':
				mirrorRefNum++
			case 'm':
				mirrorAuxNum++
			case '9':
				merchantNum++
				initCharNum++
			case 'A', 'T', 'W', 'C', 'B', 'D', '7', '8':
				initCharNum++
			}
		}
	}

	// 检查数组大小是否与 levelMap 匹配
	if stoneNum != stoneNumberInit {
		panic("没有修改 stone number")
	}
	if grassesNum != grassNumberInit {
		panic("没有修改 grass number")
	}
	if goblinNum != len(goblinArrType{}) {
		panic("没有修改 goblin number")
	}
	if dragonNum != len(dragonArrType{}) {
		panic("没有修改 dragon number")
	}
	if mirrorNum != len(mirrorArrType{}) {
		panic("没有修改 mirrorDir")
	}
	if mirrorRefNum != len(mirrorRefArrType{}) {
		panic("没有修改 mirrorRefDir")
	}
	if mirrorAuxNum != len(mirrorAuxArrType{}) {
		panic("没有修改 mirrorAuxDir")
	}
	if !allowCloneMan && merchantNum != len(merchantArrType{}) {
		panic("没有修改 merchant number")
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
	if d.cleric != noPos {
		allChars = append(allChars, d.cleric)
	}
	if d.bard != noPos {
		allChars = append(allChars, d.bard)
	}
	if d.druid != noPos {
		allChars = append(allChars, d.druid)
	}
	for _, p := range d.merchant {
		if p != noPos {
			allChars = append(allChars, p)
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
	for _, p := range d.mirrorRefs {
		if p.point != noPos {
			objs = append(objs, p.point)
		}
	}
	for _, p := range d.mirrorAuxes {
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

// p 不是固体（p 是空地，或者 p 是可移动对象）
func (d *data) isValidPos(p point) bool {
	x, y, z := p.x, p.y, p.z
	if !(0 <= x && x < n && 0 <= y && y < m) {
		return false
	}
	if z == 0 {
		if levelMap[x][y] == '#' { // 墙
			return false
		}
		if slices.Contains(d.grass[:], point{x, y, 0}) { // 草
			return false
		}
		if !d.monsterDoorOpened && slices.Contains(monsterDoors[:], point{x, y, 0}) { // 怪物门
			return false
		}
		for i, opened := range d.doorOpened {
			if !opened && slices.Contains(doors[i][:], point{x, y, 0}) { // 压力门
				return false
			}
		}
		return true
	}
	if z == -1 {
		// todo 水中的门
		if levelMap[x][y] != '~' {
			return false
		}
		return true
	}
	panic("todo z > 0 的情况")
}

func (d *data) isProtected(char point) bool {
	return char == d.cleric || d.cleric != noPos && isNeighbor(char, d.cleric)
}

// 在水面上且下面没有石头（或者门）的对象，落入水中
// todo 摧毁水中的镜子
func (d *data) isFallIntoWater(p point) bool {
	// todo 栏杆
	//if !canFallIntoWater && slices.Contains(d.stones[:], p) { 
	//	return false
	//}
	if p.z == -1 || p == noPos {
		return false
	}
	ch := levelMap[p.x][p.y]
	// todo 门
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
	// todo 忽略向上的门（应该抬高角色）
	for i, opened := range d.doorOpened {
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
	case oldP == d.cleric:
		d.cleric = newP
	case oldP == d.bard:
		d.bard = newP
	case oldP == d.druid:
		d.druid = newP
	default:
		changed := false

		i := slices.Index(d.merchant[:], oldP)
		if i >= 0 {
			changed = true
			d.merchant[i] = newP
		}

		i = pdIndex(d.mirrors[:], oldP)
		if i >= 0 {
			changed = true
			d.mirrors[i].point = newP
		}

		i = pdIndex(d.mirrorRefs[:], oldP)
		if i >= 0 {
			changed = true
			d.mirrorRefs[i].point = newP
		}

		i = pdIndex(d.mirrorAuxes[:], oldP)
		if i >= 0 {
			changed = true
			d.mirrorAuxes[i].point = newP
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

		for j, p := range d.dragons[:] {
			if p.point == oldP {
				changed = true
				d.dragons[j].point = newP
				break
			}
		}

		if !changed {
			panic("没有发生修改，请检查代码")
		}
	}
}

func solveLevel(debug bool) []string {
	merchantInitArr := merchantArrType{}
	for i := range merchantInitArr {
		merchantInitArr[i] = noPos
	}
	mirrorInitArr := mirrorArrType{}
	mirrorRefInitArr := mirrorRefArrType{}
	mirrorAuxInitArr := mirrorAuxArrType{}
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
	__cleric := noPos
	__bard := noPos
	__druid := noPos
	__merchants := merchantInitArr[:0]
	__mirrors := mirrorInitArr[:0]
	__mirrorRefs := mirrorRefInitArr[:0]
	__mirrorAuxes := mirrorAuxInitArr[:0]
	__stones := stoneInitArr[:0]
	__grass := grassInitArr[:0]
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
			case 'C':
				if __curChar < 0 {
					__curChar = charCleric
				}
				if __cleric == noPos {
					__cleric = p
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
			case '9':
				if __curChar < 0 {
					__curChar = charMerchant
				}
				__merchants = append(__merchants, p)
			case 'M':
				idx := len(__mirrors)
				dir0 := getDir(mirrorDirInit[idx*2])
				dir1 := getDir(mirrorDirInit[idx*2+1])
				__mirrors = append(__mirrors, pointWithDir{p, dir1<<4 | dir0})
			case 'R':
				idx := len(__mirrorRefs)
				dir0 := getDir(mirrorRefDirInit[idx*2])
				dir1 := getDir(mirrorRefDirInit[idx*2+1])
				__mirrorRefs = append(__mirrorRefs, pointWithDir{p, dir1<<4 | dir0})
			case 'm':
				idx := len(__mirrorAuxes)
				dir0 := getDir(mirrorAuxDirInit[idx*2])
				dir1 := getDir(mirrorAuxDirInit[idx*2+1])
				__mirrorAuxes = append(__mirrorAuxes, pointWithDir{p, dir1<<4 | dir0})
			case 's':
				__stones = append(__stones, p)
			case 'w':
				__grass = append(__grass, p)
			case 'g':
				__goblins = append(__goblins, p)
			case 'd':
				idx := len(__dragons)
				__dragons = append(__dragons, pointWithDir{p, getDir(dragonDirInit[idx])})
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
			case '.', '#', '~':
				// pass
			default:
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
	if __cleric != noPos {
		validChars = append(validChars, charCleric)
	}
	if __bard != noPos {
		validChars = append(validChars, charBard)
	}
	if __druid != noPos {
		validChars = append(validChars, charDruid)
	}
	if len(merchantInitArr) > 0 {
		validChars = append(validChars, charMerchant)
	}

	levelData := data{
		curChar:  __curChar,
		warrior:  __warrior,
		thief:    __thief,
		wizard:   __wizard,
		cleric:   __cleric,
		bard:     __bard,
		druid:    __druid,
		merchant: merchantInitArr,

		stones:  stoneInitArr,
		grass:   grassInitArr,
		goblins: goblinInitArr,
		dragons: dragonInitArr,

		mirrors:     mirrorInitArr,
		mirrorRefs:  mirrorRefInitArr,
		mirrorAuxes: mirrorAuxInitArr,
	}

	type pair struct {
		data
		info string
	}
	from := map[data]pair{} // 同时充当 vis 的功能
	queue := []data{}

	add := func(last, d data, info string) {
		_, allMovableObjs := d.getAllMovableObjPos()

		if len(weightSwitches[0]) > 0 {
			for i, weightSwitch := range weightSwitches {
				opened := true
				for _, w := range weightSwitch {
					if !slices.Contains(allMovableObjs, w) && !slices.Contains(d.grass[:], w) { // 草也可以按住地板
						opened = false
						break
					}
				}
				//if opened {
				//	fmt.Printf("门 %d 开启\n", i)
				//}
				d.doorOpened[i] = opened

				// 石头被门压碎（石头在门中，但门没有打开）
				if !opened {
					for j, p := range d.stones {
						if slices.Contains(doors[i], p) {
							if !canDestroyObj {
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
				dir := directions[dra.dir]
				cur := point{dra.x, dra.y, dra.z}
				for {
					cur.x += dir.x
					cur.y += dir.y
					cur.z += dir.z
					if !d.isValidPos(cur) {
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
						if !canDestroyObj {
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
						if !canDestroyObj {
							return
						}
						dragons[i].point = noPos
					}
				}
				slices.SortFunc(dragons[:], cmpPointWithDir)
			}

			if canDestroyObj {
				d.goblins = goblins
				d.dragons = dragons
				d.monsterDoorOpened = d.allMonstersDied()
			}
		}

		// todo 石头/镜子落入水中的镜子，水中的镜子会被摧毁

		// 镜子
		if len(d.mirrors) > 0 {
			mir := d.mirrors[:]
			for i, p := range mir {
				if d.isFallIntoWater(p.point) {
					if !canFallIntoWater {
						return
					}
					mir[i].z = -1
				}
			}
			slices.SortFunc(mir, cmpPointWithDir)
		}

		// 可以被反射的镜子
		if len(d.mirrorRefs) > 0 {
			mir := d.mirrorRefs[:]
			for i, p := range mir {
				if d.isFallIntoWater(p.point) {
					if !canFallIntoWater {
						return
					}
					mir[i].z = -1
				}
			}
			slices.SortFunc(mir, cmpPointWithDir)
		}

		// 辅助镜子
		if len(d.mirrorAuxes) > 0 {
			mir := d.mirrorAuxes[:]
			for i, p := range mir {
				if d.isFallIntoWater(p.point) {
					if !canFallIntoWater {
						return
					}
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
					if !canFallIntoWater {
						return
					}
					sto[i].z = -1
				}
			}
			slices.SortFunc(sto, cmpPoint)
		}

		// 草
		if len(d.grass) > 0 {
			slices.SortFunc(d.grass[:], cmpPoint)
		}

		// 人排序
		if len(d.merchant) > 0 {
			slices.SortFunc(d.merchant[:], cmpPoint)
		}

		if _, ok := from[d]; !ok {
			from[d] = pair{last, info}
			queue = append(queue, d)
		}
	}

	info0 := ""
	if debug {
		info0 = "init"
	}
	add(data{}, levelData, info0)

	for len(queue) > 0 {
		// 注意入队的时候修改了物品的位置（重力落下）
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
				//if d.merchant[0] == (point{9, 3, 0}) {
				//	fmt.Println(d)
				//}
				path = append(path, p.info)
				d = p.data
			}
			slices.Reverse(path)
			return path
		}

		// 移动当前角色
		// todo 多控时，如果下一个位置是没有石头的水，则一个角色无法移动（已在商人中实现）
		switch d.curChar {
		case charDefault:
			panic("代码有误，当前角色不能为 charDefault")
		case charWarrior:
			// 普通移动一步
			p0 := d.warrior
			for dIdx, dir := range directions {
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
				if !d.isValidPos(cur) {
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
			for dIdx, dir := range directions {
				x, y, z := p0.x+dir.x, p0.y+dir.y, p0.z+dir.z
				np := point{x, y, z}
				if !d.isValidPos(np) || slices.Contains(allMovableObjs, np) {
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
			for dIdx, dir := range directions {
				// 该方向上是否有人/物
				x, y, z := p0.x, p0.y, p0.z
				for {
					x += dir.x
					y += dir.y
					z += dir.z
					np := point{x, y, z}
					if !d.isValidPos(np) {
						break // 出界或者有障碍物
					}
					if !slices.Contains(allMovableObjs, np) {
						continue // 空地
					}
					// todo 辅助镜子可以吗？
					if i := pdIndex(d.mirrors[:], np); i >= 0 && d.mirrors[i].canReflect(dir) {
						break // 不能刚好面对镜子
					}
					if i := pdIndex(d.mirrorRefs[:], np); i >= 0 && d.mirrorRefs[i].canReflect(dir) {
						break // 不能刚好面对镜子
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
				if !d.isValidPos(np) {
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
		case charCleric:
			// 普通移动一步
			p0 := d.cleric
			for dIdx, dir := range directions {
				np := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				if !d.isValidPos(np) || slices.Contains(allMovableObjs, np) {
					continue // 枚举另一个方向
				}
				newData := d
				newData.cleric = np
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

			// 普通移动一步
			// 切比雪夫距离 <= 2 的物品（包括自己）都移动一步
			for dIdx, dir := range directions {
				x, y, z := p0.x+dir.x, p0.y+dir.y, p0.z+dir.z
				if !d.isValidPos(point{x, y, z}) {
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
				movedItems := []point{}
				for _, oldP := range items {
					// item 往前移动一格
					newP := point{oldP.x + dir.x, oldP.y + dir.y, oldP.z + dir.z}
					if !d.isValidPos(newP) { // 无法移动
						unmovedItems = append(unmovedItems, oldP)
						continue
					}
					// 尝试移动
					if chebyshevDis(newP, p0) > 2 { // item 是力场最前面的点
						if slices.Contains(allMovableObjs, newP) { // 不能与力场外的对象碰撞
							unmovedItems = append(unmovedItems, oldP)
							continue
						}
					} else if slices.Contains(unmovedItems, newP) { // 力场后面的点，不能与前面移动失败的对象碰撞
						unmovedItems = append(unmovedItems, oldP)
						continue
					}
					movedItems = append(movedItems, oldP)
					newData.changePos(oldP, newP)
				}

				if !slices.Contains(unmovedItems, p0) {
					if newData.bard != (point{x, y, z}) {
						panic("移动错误，代码有误")
					}

					// 特性：如果诗人脚下是物品，且该物品移动了，那么诗人可以再走一格
					if slices.Contains(movedItems, point{p0.x, p0.y, p0.z - 1}) {
						nxtP := point{x + dir.x, y + dir.y, z + dir.z}
						if d.isValidPos(nxtP) && !slices.Contains(unmovedItems, nxtP) {
							newData.bard = nxtP
						}
						// todo （待确认）如果 z-2 也移动了，那么再再走一格
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
			for dIdx, dir := range directions {
				newP := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				// todo 目前只实现了草 <-> 石头的逻辑
				// 草变石
				if i := slices.Index(d.grass[:], newP); i >= 0 {
					newData := d
					newData.stones[:][0] = newData.grass[i] // 加个切片避免报错
					newData.grass[i] = noPos
					var info string
					if debug {
						info = fmt.Sprintf("德 %s 草变石", debugDirString[dIdx])
					} else {
						info = dirString[dIdx]
					}
					add(d, newData, info)
					continue
				}

				// 石变草
				if i := slices.Index(d.stones[:], newP); i >= 0 {
					newData := d
					newData.grass[:][0] = newData.stones[i]
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

				// 普通移动一步
				if !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) {
					continue
				}
				newData := d
				newData.druid = newP
				var info string
				if debug {
					info = fmt.Sprintf("德 %s", debugDirString[dIdx])
				} else {
					info = dirString[dIdx]
				}
				add(d, newData, info)
			}
		case charMerchant:
			// 多控
			// 普通移动一步
			for dIdx, dir := range directions {
				newData := d
				oldMerchant := newData.merchant
				man := newData.merchant[:]
				slices.SortFunc(man, func(a, b point) int {
					if dir.x != 0 {
						return int(b.x*dir.x - a.x*dir.x)
					}
					return int(b.y*dir.y - a.y*dir.y)
				})

				unmovedMan := []point{}
				moved := false
				for manIdx, p0 := range man {
					if p0 == noPos {
						continue
					}
					nxt := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
					// 无法移动（注意岸边也是无法移动的）
					if !d.isValidPos(nxt) || d.isFallIntoWater(nxt) || slices.Contains(unmovedMan, nxt) {
						unmovedMan = append(unmovedMan, p0)
						continue
					}
					// 如果前面是物品，则推动（能移动的人已经移动了）
					if !slices.Contains(oldMerchant[:], nxt) && slices.Contains(allMovableObjs, nxt) {
						nxt2 := point{nxt.x + dir.x, nxt.y + dir.y, nxt.z + dir.z}
						// 无法推动前面的物品
						if !d.isValidPos(nxt2) ||
							!slices.Contains(oldMerchant[:], nxt2) && slices.Contains(allMovableObjs, nxt2) ||
							slices.Contains(unmovedMan, nxt2) {
							unmovedMan = append(unmovedMan, p0)
							continue
						}
						newData.changePos(nxt, nxt2)
					}
					moved = true
					man[manIdx] = nxt // 移走！
				}
				if !moved { // 没人动
					continue
				}

				var info string
				if debug {
					info = fmt.Sprintf("人 %s", debugDirString[dIdx])
				} else {
					info = dirString[dIdx]
				}
				add(d, newData, info)
			}
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
				}
				add(d, newData, info)
			}
		}

		// 镜子反射对象
		doMirrors := func() {
			newData := d
			swapped := uint(0)
		nextMirror:
			for _, mirror := range append(d.mirrors[:], d.mirrorRefs[:]...) {
				// 找两个方向最近的可反射的对象
				cur0 := mirror.point
				cur1 := mirror.point
				dir0 := directions[mirror.dir&0xf]
				dir1 := directions[mirror.dir>>4]
				foundMirror := uint8(0)
				for step := 1; ; step++ {
					justFound := uint8(0) // 是否找到了非镜子对象
					// 检查方向 0
					if foundMirror&1 == 0 {
						cur0.x += dir0.x
						cur0.y += dir0.y
						cur0.z += dir0.z
						if !d.isValidPos(cur0) {
							continue nextMirror
						}
						// todo bitset
						if pdContains(d.mirrors[:], cur0) || pdContains(d.mirrorAuxes[:], cur0) {
							foundMirror |= 1
						} else if slices.Contains(allMovableObjs, cur0) {
							justFound |= 1
						}
					}
					// 检查方向 1
					if foundMirror>>1 == 0 {
						cur1.x += dir1.x
						cur1.y += dir1.y
						cur1.z += dir1.z
						if !d.isValidPos(cur1) {
							continue nextMirror
						}
						if pdContains(d.mirrors[:], cur1) || pdContains(d.mirrorAuxes[:], cur1) {
							foundMirror |= 2
						} else if slices.Contains(allMovableObjs, cur1) {
							justFound |= 2
						}
					}
					if foundMirror == 3 {
						return // 不能两方向最近都是镜子
					}
					if justFound == 3 {
						return // 不能反射位置都是对象
					}
					if justFound == 0 {
						continue // 都是空地，继续找
					}

					oldP := cur0
					dir := dir1
					if justFound == 2 {
						oldP = cur1
						dir = dir0 // 往另一个方向反射
					}

					// 无法反射的石头
					if !areStonesReflectable && slices.Contains(d.stones[:], oldP) {
						return
					}

					// 反射：从 mirror.point 出发，往 dir 方向走 step 步
					newP := mirror.point
					for k := range step {
						newP.x += dir.x
						newP.y += dir.y
						newP.z += dir.z
						// 遇到另一面主镜子
						if i := pdIndex(d.mirrors[:], newP); i >= 0 {
							if k == step-1 {
								return // 最终反射到了镜子上
							}
							dir = d.mirrors[i].reflectToAnotherDir(dir)
							if dir == (point{}) {
								return // 镜子背对我们
							}
							continue // 改变光路，继续反射
						}
						// 遇到另一面可以反射的镜子
						if i := pdIndex(d.mirrorRefs[:], newP); i >= 0 {
							if k == step-1 {
								return // 最终反射到了镜子上
							}
							dir = d.mirrorRefs[i].reflectToAnotherDir(dir)
							if dir == (point{}) {
								return // 镜子背对我们
							}
							continue // 改变光路，继续反射
						}
						// 遇到另一面辅助镜子
						if i := pdIndex(d.mirrorAuxes[:], newP); i >= 0 {
							if k == step-1 {
								return // 最终反射到了辅助镜子上
							}
							dir = d.mirrorAuxes[i].reflectToAnotherDir(dir)
							if dir == (point{}) {
								return // 镜子背对我们
							}
							continue // 改变光路，继续反射
						}
						if !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) {
							return // 光路被障碍或非镜子对象挡住
						}
					}

					// 反射成功
					itemIdx := slices.Index(allMovableObjs, oldP)
					if swapped>>itemIdx&1 > 0 {
						// todo 所有对象的分身
						// 不能再分身了
						if slices.Contains(d.merchant[:], oldP) {
							if newData.merchant[:][0] != noPos {
								return
							}
							newData.merchant[:][0] = newP
						} else if areStonesReflectable && slices.Contains(d.stones[:], oldP) {
							//newData.stones[0] = newP // todo
						} else {
							// todo 其他对象的分身
						}
					} else {
						swapped |= 1 << itemIdx
						newData.changePos(oldP, newP)
					}
					break
				}

				// todo 合二为一
				// todo 这里恰有两人
				//if newData.merchant[0] != noPos && newData.merchant[0] == newData.merchant[1] {
				//	//newData.merchant[0] = noPos
				//	return
				//}
			}

			if swapped == 0 {
				return
			}

			var info string
			if !debug {
				info = "x"
			}
			add(d, newData, info)
		}
		doMirrors()
	}

	// 无解
	return nil
}

const digits = "0123456789"

const (
	charDefault = iota // 仅占位，不使用
	charWarrior
	charThief
	charWizard
	charCleric
	charBard
	charDruid
	charExplorer
	charSailor
	charMerchant // Trader
)
