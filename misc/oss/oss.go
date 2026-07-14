package oss

import (
	"fmt"
	"math"
	"math/bits"
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

type merchantArrType [merchantNumberInit]point
type stoneArrType [stoneNumberInit + grassNumberInit]point
type grassArrType [stoneNumberInit + grassNumberInit]point
type stoneFloatArrType [stoneFloatNumberInit]point
type goblinArrType [goblinNumberInit]pointWithDir
type dragonArrType [len(dragonDirInit)]pointWithDir
type beamArrType [len(beamDirInit)]pointWithDir
type mirrorArrType [len(mirrorDirInit) / 2]pointWithDir
type mirrorRefArrType [len(mirrorRefDirInit) / 2]pointWithDir
type mirrorAuxArrType [len(mirrorAuxDirInit) / 2]pointWithDir

type data struct {
	warrior  point           // A 推多个对象
	thief    point           // T 拉一个对象
	wizard   point           // W 交换对象
	cleric   point           // C 自己以及上下左右无敌
	bard     point           // B 同时移动切比雪夫距离 <= 2 的对象
	druid    point           // D 把对象变成石头
	explorer point           // 7 普通角色，无法推对象
	sailor   point           // 8 普通角色，推一个对象
	merchant merchantArrType // 9 普通角色，推一个对象

	// 石头/水晶
	stones stoneArrType // s
	// 用大写 S 表示可被反射的石头？
	stoneFloats stoneFloatArrType // F todo 跳石

	// 草
	grass grassArrType // w

	// 怪物
	goblins goblinArrType // g
	dragons dragonArrType // d

	// 镜子
	mirrors     mirrorArrType    // M
	mirrorRefs  mirrorRefArrType // R 主镜子 + 可以被反射
	mirrorAuxes mirrorAuxArrType // m 关卡名中称其为 mundane

	// 光束
	beams beamArrType // b 高 4 位是类型，低 4 位是方向

	// 门的开闭
	doorOpened        [doorKinds]bool
	monsterDoorOpened bool

	// 当前角色类型
	curCharTypeNum int8
}

var mapSizeH, mapSizeN, mapSizeM int8
var initCharNum uint8

func init() {
	mapSizeH = int8(len(levelMap))
	mapSizeN = int8(len(levelMap[0]))
	mapSizeM = int8(len(levelMap[0][0]))

	var stoneNum, grassesNum, goblinNum, dragonNum, beamNum, mirrorNum, mirrorRefNum, mirrorAuxNum, merchantNum, doorMask int
	for i, ds := range doors {
		if len(ds) > 0 {
			doorMask |= 1 << i
		}
	}

	for _, grid := range levelMap {
		for _, row := range grid {
			if len(row) != int(mapSizeM) {
				panic("行不等长")
			}
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
				case 'b':
					beamNum++
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
				case 'X', 'Y', 'Z', '[':
					doorMask |= 1 << (ch - 'X')
				}
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
		panic("没有修改 dragon dir")
	}
	if beamNum != len(beamArrType{}) {
		panic("没有修改 beam number")
	}
	if len(beamDirInit) != len(beamTypeInit) {
		panic("没有修改 beam type")
	}
	if mirrorNum != len(mirrorArrType{}) {
		panic("没有修改 mirror dir")
	}
	if mirrorRefNum != len(mirrorRefArrType{}) {
		panic("没有修改 mirror ref dir")
	}
	if mirrorAuxNum != len(mirrorAuxArrType{}) {
		panic("没有修改 mirror aux dir")
	}
	if !allowCloneMan && merchantNum != len(merchantArrType{}) {
		panic("没有修改 merchant number")
	}
	if bits.OnesCount(uint(doorMask)) != doorKinds {
		panic("没有修改 door kinds")
	}
}

func (d *data) areAllMonstersDied() bool {
	for _, p := range d.goblins {
		if p.point != noPos {
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
func (d *data) getAllCharPos(isBigMap bool) []point {
	if isBigMap {
		return nil
	}

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
	if d.explorer != noPos {
		allChars = append(allChars, d.explorer)
	}
	if d.sailor != noPos {
		allChars = append(allChars, d.sailor)
	}
	for _, p := range d.merchant {
		if p != noPos {
			allChars = append(allChars, p)
		}
	}
	return allChars
}

func (d *data) getAllMovableObjPos(isBigMap bool) ([]point, []point) {
	chars := d.getAllCharPos(isBigMap)
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
		if p.point != noPos {
			objs = append(objs, p.point)
		}
	}
	for _, p := range d.dragons {
		if p.point != noPos {
			objs = append(objs, p.point)
		}
	}
	for _, p := range d.beams {
		if p.point != noPos {
			objs = append(objs, p.point)
		}
	}
	return chars, objs
}

func inBound(p point) bool {
	return 0 <= p.x && p.x < mapSizeN &&
		0 <= p.y && p.y < mapSizeM &&
		p.z < mapSizeH
}

// p 不是固体（p 是空地，或者 p 是可移动对象）
func (d *data) isValidPos(p point) bool {
	x, y, z := p.x, p.y, p.z
	if !inBound(p) {
		return false
	}

	if z < 0 {
		if z != -1 {
			panic("invalid z")
		}
		// todo 水中的门
		if levelMap[0][x][y] != '~' {
			return false
		}
		return true
	}

	if levelMap[z][x][y] == '#' { // 墙
		return false
	}
	if slices.Contains(d.grass[:], point{x, y, z}) { // 草
		return false
	}
	if !d.monsterDoorOpened && slices.Contains(monsterDoors[:], point{x, y, z}) { // 怪物门
		return false
	}
	for i, opened := range d.doorOpened {
		if !opened && slices.Contains(doors[i][:], point{x, y, z}) { // 压力门
			return false
		}
	}
	return true
}

// 返回 mask 表示在哪些 beam 中
func (d *data) withinBeams(p point, allMovableObjs []point) (mask uint8) {
	for _, beam := range d.beams {
		dir := directions4[beam.dir&0xf]

		// 剪枝：先粗略判断是否在光束方向上（不考虑障碍）
		if dir.x != 0 { // 上下，必须同 y
			if beam.y != p.y { // todo z
				continue
			}
			if dir.x > 0 != (beam.x < p.x) {
				continue
			}
		} else if dir.y != 0 { // 左右，必须同 x
			if beam.x != p.x { // todo z
				continue
			}
			if dir.y > 0 != (beam.y < p.y) {
				continue
			}
		} else {
			panic("todo dir.z != 0")
		}

		cur := beam.point
		for {
			cur.x += dir.x
			cur.y += dir.y
			cur.z += dir.z
			if cur == p {
				mask |= 1 << (beam.dir >> 4)
				break
			}
			// 出界，或者遇到对象障碍（墙反而不是障碍）
			// todo 门
			if !inBound(cur) || slices.Contains(allMovableObjs, cur) {
				break
			}
		}
	}
	return
}

func (d *data) isProtected(char point) bool {
	return char == d.cleric || d.cleric != noPos && isNeighbor6(char, d.cleric)
}

func (d *data) isFallingIntoGround(p point) bool {
	return p.z > 0 && levelMap[p.z-1][p.x][p.y] == '.'
}

// 在水面上且下面没有石头（或者门）的对象，落入水中
// todo 摧毁水中的镜子
func (d *data) isFallIntoWater(p point) bool {
	// todo z > 0 中途遇到障碍
	// todo 栏杆
	//if !canFallIntoWater && slices.Contains(d.stones[:], p) { 
	//	return false
	//}
	if p.z == -1 || p == noPos || levelMap[0][p.x][p.y] != '~' {
		return false
	}
	downP := point{p.x, p.y, -1}
	// 水中的门
	for i, opened := range d.doorOpened {
		if !opened && slices.Contains(doors[i][:], downP) {
			return false
		}
	}
	// 水中的石头
	if slices.Contains(d.stones[:], downP) {
		return false
	}
	return true
}

func (d *data) isAttacked(p point, burnPos []point) bool {
	// 喷火龙
	if slices.Contains(burnPos, p) {
		return true
	}

	// 哥布林
	for _, g := range d.goblins {
		if isNeighbor4(g.point, p) {
			return true
		}
	}

	return false
}

const (
	dieTypeNo = iota
	dieTypeCrushed
	dieTypeAttacked
	dieTypeDrown
)

func (d *data) getDieType(p point, burnPos []point, isChar bool) int {
	// 被门压死
	// todo 忽略向上的门（应该抬高角色）
	for i, opened := range d.doorOpened {
		if !opened && slices.Contains(doors[i][:], p) {
			return dieTypeCrushed
		}
	}

	// 淹死
	if d.isFallIntoWater(p) {
		return dieTypeDrown
	}

	if isChar && d.isProtected(p) {
		return dieTypeNo
	}

	if d.isAttacked(p, burnPos) {
		return dieTypeAttacked
	}

	return dieTypeNo
}

// 反射：从 mirror.point 出发，往 dir 方向走 step 步
func (d *data) reflectTo(mirror pointWithDir, dir point, step int, allMovableObjs []point) point {
	cur := mirror.point
	for k := range step {
		cur.x += dir.x
		cur.y += dir.y
		cur.z += dir.z
		// 遇到另一面主镜子
		if i := pdIndex(d.mirrors[:], cur); i >= 0 {
			if k == step-1 { // 按 X 反射
				return noPos // 最终反射到了镜子上，这不行
			}
			dir = d.mirrors[i].reflectToAnotherDir(dir)
			if dir == (point{}) {
				// 镜子背对我们
				if step == math.MaxInt { // 法师
					return d.mirrors[i].point
				}
				return noPos
			}
			continue // 改变光路，继续反射
		}
		// 遇到另一面可以反射的镜子
		if i := pdIndex(d.mirrorRefs[:], cur); i >= 0 {
			if k == step-1 {
				return noPos // 最终反射到了镜子上
			}
			dir = d.mirrorRefs[i].reflectToAnotherDir(dir)
			if dir == (point{}) {
				// 镜子背对我们
				if step == math.MaxInt { // 法师
					return d.mirrorRefs[i].point
				}
				return noPos
			}
			continue // 改变光路，继续反射
		}
		// 遇到另一面辅助镜子
		if i := pdIndex(d.mirrorAuxes[:], cur); i >= 0 {
			if k == step-1 {
				return noPos // 最终反射到了辅助镜子上
			}
			dir = d.mirrorAuxes[i].reflectToAnotherDir(dir)
			if dir == (point{}) {
				// 镜子背对我们
				if step == math.MaxInt { // 法师
					return d.mirrorAuxes[i].point
				}
				return noPos
			}
			continue // 改变光路，继续反射
		}
		// 光路被（不可移动对象）挡住
		if !d.isValidPos(cur) {
			return noPos
		}
		// 光路被非镜子对象挡住
		if i := slices.Index(allMovableObjs, cur); i >= 0 {
			if step == math.MaxInt { // 法师
				return allMovableObjs[i]
			}
			return noPos
		}
	}
	// 按 X 反射
	return cur
}

func (d *data) changePos(oldP, newP point, newDir uint8) {
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
	case oldP == d.explorer:
		d.explorer = newP
	case oldP == d.sailor:
		d.sailor = newP
	default:
		changed := false

		//if i := slices.Index(d.sailor[:], oldP); i >= 0 {
		//	changed = true
		//	d.sailor[i] = newP
		//}

		if i := slices.Index(d.merchant[:], oldP); i >= 0 {
			changed = true
			d.merchant[i] = newP
		}

		if i := pdIndex(d.mirrors[:], oldP); i >= 0 {
			changed = true
			d.mirrors[i].point = newP
			if newDir != math.MaxUint8 {
				d.mirrors[i].dir = newDir
			}
		}

		if i := pdIndex(d.mirrorRefs[:], oldP); i >= 0 {
			changed = true
			d.mirrorRefs[i].point = newP
			if newDir != math.MaxUint8 {
				d.mirrorRefs[i].dir = newDir
			}
		}

		if i := pdIndex(d.mirrorAuxes[:], oldP); i >= 0 {
			changed = true
			d.mirrorAuxes[i].point = newP
			if newDir != math.MaxUint8 {
				d.mirrorAuxes[i].dir = newDir
			}
		}

		if i := slices.Index(d.stones[:], oldP); i >= 0 {
			changed = true
			d.stones[i] = newP
		}

		if canPushGoblin {
			if i := pdIndex(d.goblins[:], oldP); i >= 0 {
				changed = true
				d.goblins[i].point = newP
			}
		}

		if canPushDragon {
			if i := pdIndex(d.dragons[:], oldP); i >= 0 {
				changed = true
				d.dragons[i].point = newP
				if newDir != math.MaxUint8 {
					d.dragons[i].dir &^= 7
					d.dragons[i].dir |= newDir
				}
			}
		}

		if canPushBeam {
			if i := pdIndex(d.beams[:], oldP); i >= 0 {
				changed = true
				d.beams[i].point = newP
				if newDir != math.MaxUint8 {
					d.beams[i].dir = newDir
				}
			}
		}

		if !changed {
			panic("没有发生修改，请检查代码")
		}
	}
}

func (d *data) getCurCharTypePos() (pos point) {
	switch d.curCharTypeNum {
	case charDefault:
		panic("代码有误，当前角色不能为 charDefault")
	case charWarrior:
		pos = d.warrior
	case charThief:
		pos = d.thief
	case charWizard:
		pos = d.wizard
	case charCleric:
		pos = d.cleric
	case charBard:
		pos = d.bard
	case charDruid:
		pos = d.druid
	case charExplorer:
		pos = d.explorer
	case charSailor:
		pos = d.sailor
	case charMerchant:
		pos = d.merchant[:][0]
	default:
		panic("未找到当前角色")
	}
	return
}

func solveLevel() []string {
	sailorInitArr := merchantArrType{}
	for i := range sailorInitArr {
		sailorInitArr[i] = noPos
	}
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
	beamInitArr := beamArrType{}

	__curCharTypeNum := initCharTypeNum
	__warrior := warriorPosInit
	__thief := thiefPosInit
	__wizard := wizardPosInit
	__cleric := noPos
	__bard := noPos
	__druid := noPos
	__explorer := noPos
	__sailor := noPos
	__sailors := sailorInitArr[:0]
	__merchants := merchantInitArr[:0]

	__mirrors := mirrorInitArr[:0]
	__mirrorRefs := mirrorRefInitArr[:0]
	__mirrorAuxes := mirrorAuxInitArr[:0]
	__stones := stoneInitArr[:0]
	__grass := grassInitArr[:0]
	__goblins := goblinInitArr[:0]
	__dragons := dragonInitArr[:0]
	__beams := beamInitArr[:0]
	for z, grid := range levelMap {
		for x, row := range grid {
			for y, ch := range row {
				p := point{int8(x), int8(y), int8(z)}
				switch ch {
				case 'A':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charWarrior
					}
					if __warrior == noPos {
						__warrior = p
					}
				case 'T':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charThief
					}
					if __thief == noPos {
						__thief = p
					}
				case 'W':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charWizard
					}
					if __wizard == noPos {
						__wizard = p
					}
				case 'C':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charCleric
					}
					if __cleric == noPos {
						__cleric = p
					}
				case 'B':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charBard
					}
					if __bard == noPos {
						__bard = p
					}
				case 'D':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charDruid
					}
					if __druid == noPos {
						__druid = p
					}
				case '7':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charExplorer
					}
					if __explorer == noPos {
						__explorer = p
					}
				case '8':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charSailor
					}
					if __sailor == noPos {
						__sailor = p
					}
					_ = __sailors
					//__sailors = append(__sailors, p)
				case '9':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charMerchant
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
					__goblins = append(__goblins, pointWithDir{p, math.MaxUint8})
				case 'd':
					idx := len(__dragons)
					__dragons = append(__dragons, pointWithDir{p, getDir(dragonDirInit[idx])})
				case 'b':
					idx := len(__beams)
					dir := getDir(beamDirInit[idx])
					tp := beamTypeInit[idx] - '0'
					__beams = append(__beams, pointWithDir{p, tp<<4 | dir})
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
					panic(fmt.Sprintf("不支持的符号 %c", ch))
				}
			}
		}
	}

	// 有时候会手动添加 finals 的初始值，总体不一定是有序的
	slices.SortFunc(finals, cmpPoint)

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
	if __explorer != noPos {
		validChars = append(validChars, charExplorer)
	}
	if __sailor != noPos {
		validChars = append(validChars, charSailor)
	}
	//if len(sailorInitArr) > 0 {
	//	validChars = append(validChars, charSailor)
	//}
	if len(merchantInitArr) > 0 {
		validChars = append(validChars, charMerchant)
	}

	if !slices.Contains(validChars, __curCharTypeNum) {
		panic(fmt.Sprint("请修改 initCharTypeNum"))
	}

	levelData := data{
		warrior:  __warrior,
		thief:    __thief,
		wizard:   __wizard,
		cleric:   __cleric,
		bard:     __bard,
		druid:    __druid,
		explorer: __explorer,
		sailor:   __sailor,
		merchant: merchantInitArr,

		stones:  stoneInitArr,
		grass:   grassInitArr,
		goblins: goblinInitArr,
		dragons: dragonInitArr,

		mirrors:     mirrorInitArr,
		mirrorRefs:  mirrorRefInitArr,
		mirrorAuxes: mirrorAuxInitArr,

		beams: beamInitArr,

		//doorOpened:        doorOpenedInit,
		monsterDoorOpened: monsterDoorOpenedInit,

		curCharTypeNum: __curCharTypeNum,
	}

	type pair struct {
		data
		info string
	}
	from := map[data]pair{} // 同时充当 vis 的功能
	queue := []data{}

	//hasStone := map[point]bool{}

	add := func(last, d data, info string) {
		//if !hasStone[d.stones[0]] {
		//	hasStone[d.stones[0]] = true
		//	fmt.Println(d.stones[0])
		//}

		_, allMovableObjs := d.getAllMovableObjPos(isBigMap)

		// 先确定门的开闭
		for i, weightSwitch := range weightSwitches {
			opened := !doorOpenedInit[i]
			for _, w := range weightSwitch {
				if !slices.Contains(allMovableObjs, w) && !slices.Contains(d.grass[:], w) { // 草也可以按住地板
					opened = !opened // 反转开闭状态
					break
				}
			}
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

		// 被喷火龙攻击到的位置
		// todo 镜子反射火焰
		var burnedPos []point
		if !d.monsterDoorOpened { // 喷火龙没有 die（如果没有怪物门，monsterDoorOpened = false）
			for _, dra := range d.dragons {
				if dra.dir&dirStoneDelta > 0 { // 是石头
					continue
				}
				dir := directions4[dra.dir]
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

		// 对象下落
		// todo 整合后面的落水逻辑
		if mapSizeH > 1 {
			for _, p := range allMovableObjs {
				// 如果 p 是牧师或其邻居，且正被攻击，那么 p 不会下落
				if d.isProtected(p) && d.isAttacked(p, burnedPos) {
					continue
				}
				oldP := p
				p.z--
				for p.z >= 0 && d.isValidPos(p) && !slices.Contains(allMovableObjs, p) {
					p.z--
				}
				p.z++
				if oldP.z != p.z {
					if !allowFallIntoGround {
						return
					}
					d.changePos(oldP, p, math.MaxUint8)
				}
			}
		}

		// 先判断是否有角色死亡
		for _, char := range d.getAllCharPos(isBigMap) {
			if d.getDieType(char, burnedPos, true) != dieTypeNo {
				return
			}
		}

		dieType := dieTypeNo
		// 一开始，以及切换角色，都不结算攻击
		isSwitching := info[0] == 'c' || '1' <= info[0] && info[0] <= '9'
		if !isSwitching && !d.monsterDoorOpened {
			// 哥布林
			goblins := d.goblins
			if len(d.goblins) > 0 {
				for i, p := range d.goblins {
					// todo 变成石头的哥布林
					if tp := d.getDieType(p.point, burnedPos, false); tp != dieTypeNo {
						if !canDestroyObj {
							return
						}
						dieType = tp
						goblins[i].point = noPos
					}
				}
				slices.SortFunc(goblins[:], cmpPointWithDir) // 一定要排序，不然状态数爆炸了
			}

			// 喷火龙
			dragons := d.dragons
			if len(d.dragons) > 0 {
				for i, p := range d.dragons {
					if p.dir&dirStoneDelta > 0 { // 是石头
						// todo 落水
						continue
					}
					if tp := d.getDieType(p.point, burnedPos, false); tp != dieTypeNo {
						if !canDestroyObj {
							return
						}
						dieType = tp
						dragons[i].point = noPos
					}
				}
				slices.SortFunc(dragons[:], cmpPointWithDir) // 一定要排序，不然状态数爆炸了
			}

			if canDestroyObj {
				d.goblins = goblins
				d.dragons = dragons
				d.monsterDoorOpened = d.areAllMonstersDied()
			}
		}

		// todo 石头/镜子落入水中的镜子，水中的镜子会被摧毁

		// 镜子
		if len(d.mirrors) > 0 {
			mir := d.mirrors[:]
			for i, p := range mir {
				if d.isFallIntoWater(p.point) {
					if !allowFallIntoWater {
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
					if !allowFallIntoWater {
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
					if !allowFallIntoWater {
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
					if !allowFallIntoWater {
						return
					}
					info += "W"
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
		//if len(d.sailor) > 0 {
		//	slices.SortFunc(d.sailor[:], cmpPoint)
		//}
		if len(d.merchant) > 0 {
			slices.SortFunc(d.merchant[:], cmpPoint)
		}

		// 光束排序
		if len(d.beams) > 0 {
			slices.SortFunc(d.beams[:], cmpPointWithDir)
		}

		if _, ok := from[d]; !ok {
			if dieType == dieTypeAttacked {
				info += "K" // 怪物攻击动画
			} else if dieType == dieTypeDrown {
				info += "W"
			}
			from[d] = pair{last, info}
			queue = append(queue, d)
		}
	}

	add(data{}, levelData, "c")

	for len(queue) > 0 {
		// 注意入队的时候修改了物品的位置（重力落下）
		d := queue[0]
		queue = queue[1:]

		allChars, allMovableObjs := d.getAllMovableObjPos(isBigMap)

		var pass bool
		if !targetIsClearAllMonsters {
			// 标准版：所有人都到达终点
			if isBigMap {
				p := d.getCurCharTypePos()
				pass = slices.Equal([]point{p}, finals)
			} else {
				slices.SortFunc(allChars, cmpPoint)
				pass = slices.Equal(allChars, finals)
			}
		} else {
			// 简化版：怪物门开启（怪物都被杀）
			pass = d.monsterDoorOpened
		}
		if pass {
			// 生成操作序列
			path := []string{}
			for {
				var ok bool
				pre, ok := from[d]
				if !ok {
					panic("代码修改了 d，与存入的 d 不符")
				}
				d = pre.data
				if d == (data{}) { // 初始状态
					break
				}
				if pre.info != "IGNORE" {
					path = append(path, pre.info)
				}
			}
			slices.Reverse(path)
			return path
		}

		// todo 如果角色的头上有物品，物品会跟着移动（注意镜子的方向会变）    堆叠上限是多少？？
		// todo 即使人没有移动，切换方向也会改变头上物品（镜子、激光等）的方向
		// todo 多控时，如果下一个位置是没有石头的水，则一个角色无法移动（已在商人中实现）

		// todo 修改 changePos 的代码，添加一个参数 alsoMoveTop bool，
		//      使得当物品移动时，物品上方的物品（如果有）也跟着移动

		// 先考虑按 x 镜子反射对象，这样后面移动更流畅
		doMirrors := func() {
			newData := d
			swapped := uint(0)
		nextMirror:
			for _, mirror := range append(d.mirrors[:], d.mirrorRefs[:]...) {
				// 找两个方向最近的可反射的对象
				cur0 := mirror.point
				cur1 := mirror.point
				dir0 := directions4[mirror.dir&0xf]
				dir1 := directions4[mirror.dir>>4]
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

					// 反射
					newP := d.reflectTo(mirror, dir, step, allMovableObjs)
					if newP == noPos {
						return // 反射失败
					}
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
						// todo 如果是 oldP 是喷火龙，则朝向会变，需要修改朝向
						newData.changePos(oldP, newP, math.MaxUint8)
					}
					break
				}

				// 合二为一
				if allowMerge {
					// todo 这里恰有两人
					man := newData.merchant[:]
					if man[0] != noPos && man[0] == man[1] {
						man[0] = noPos
					}
				}
			}

			if swapped == 0 {
				return
			}

			add(d, newData, "x")
		}
		doMirrors()

		// 移动当前角色
		switch d.curCharTypeNum {
		case charDefault:
			panic("代码有误，当前角色不能为 charDefault")
		case charWarrior:
			// 普通移动一步
			p0 := d.warrior
			for dIdx, dir := range directions4 {
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
					newData.changePos(nxt, cur, math.MaxUint8)
					cur = nxt
				}
				np := point{x, y, z}
				newData.warrior = np
				add(d, newData, dir4String[dIdx])
			}
		case charThief:
			// 普通移动一步
			p0 := d.thief
			for dIdx, dir := range directions4 {
				x, y, z := p0.x+dir.x, p0.y+dir.y, p0.z+dir.z
				np := point{x, y, z}
				if !d.isValidPos(np) || slices.Contains(allMovableObjs, np) {
					continue // 枚举另一个方向
				}
				newData := d
				back := point{p0.x - dir.x, p0.y - dir.y, p0.z - dir.z}
				if slices.Contains(allMovableObjs, back) {
					// 拉人/物 -> 当前位置
					newData.changePos(back, p0, math.MaxUint8)
				}
				newData.thief = np
				add(d, newData, dir4String[dIdx])
			}
		case charWizard:
			p0 := d.wizard
		nextDir:
			for dIdx, dir := range directions4 {
				// dir 方向是否有可交换对象
				x, y, z := p0.x, p0.y, p0.z
				for {
					x += dir.x
					y += dir.y
					z += dir.z
					newP := point{x, y, z}
					if !d.isValidPos(newP) {
						break // 出界或者有障碍物
					}
					if !slices.Contains(allMovableObjs, newP) {
						continue // 空地
					}

					mir := noPosDir
					if i := pdIndex(d.mirrors[:], newP); i >= 0 && d.mirrors[i].canReflect(dir) {
						mir = d.mirrors[i]
					} else if i := pdIndex(d.mirrorRefs[:], newP); i >= 0 && d.mirrorRefs[i].canReflect(dir) {
						mir = d.mirrorRefs[i]
					} else if i := pdIndex(d.mirrorAuxes[:], newP); i >= 0 && d.mirrorAuxes[i].canReflect(dir) {
						mir = d.mirrorAuxes[i]
					}

					// 面对的是镜子的正面
					if mir.point != noPos {
						dir2 := mir.reflectToAnotherDir(dir)
						// 沿着光路搜索，找第一个可交换对象
						newP = d.reflectTo(mir, dir2, math.MaxInt, allMovableObjs)
						if newP == noPos {
							break // 镜子反射路径没有任何对象，只能普通移动一步
						}
					}

					// 和对象交换位置
					newData := d
					newData.changePos(newP, p0, math.MaxUint8) // newP 换到 p0
					newData.wizard = newP                      // 法师换到 newP
					add(d, newData, dir4String[dIdx]+"P")      // swap
					continue nextDir
				}

				// 没有可交换对象，那就普通移动一步
				newP := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				if !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) {
					continue // 枚举另一个方向
				}
				newData := d
				newData.wizard = newP
				add(d, newData, dir4String[dIdx]) // move
			}
		case charCleric:
			// 普通移动一步
			p0 := d.cleric
			withinBeams := d.withinBeams(p0, allMovableObjs)
			for dIdx, dir := range directions4 {
				newP := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				if !d.isValidPos(newP) {
					continue // 枚举另一个方向
				}
				newData := d
				if i := slices.Index(allMovableObjs, newP); i >= 0 {
					if withinBeams>>beamPush&1 == 0 {
						continue // 枚举另一个方向
					}
					if allowPushItem {
						// 可以推物品
						nxt2 := point{newP.x + dir.x, newP.y + dir.y, newP.z + dir.z}
						if !d.isValidPos(nxt2) || slices.Contains(allMovableObjs, nxt2) {
							continue // 枚举另一个方向
						}
						newData.changePos(newP, nxt2, math.MaxUint8)
					}
				}
				newData.cleric = newP
				add(d, newData, dir4String[dIdx])
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
			for dIdx, dir := range directions4 {
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
					newData.changePos(oldP, newP, math.MaxUint8)
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

					add(d, newData, dir4String[dIdx])
				}
			}
		case charDruid:
			p0 := d.druid
			for dIdx, dir := range directions4 {
				newP := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				// 草变石
				if i := slices.Index(d.grass[:], newP); i >= 0 {
					newData := d
					newData.stones[:][0] = newData.grass[i] // 加个切片避免报错
					newData.grass[i] = noPos
					add(d, newData, dir4String[dIdx]+"C") // trans
					continue
				}

				// 石变草
				if !druidOnlyGrassToStone {
					if i := slices.Index(d.stones[:], newP); i >= 0 {
						newData := d
						newData.grass[:][0] = newData.stones[i]
						newData.stones[i] = noPos
						add(d, newData, dir4String[dIdx]+"C") // trans
						continue
					}
				}

				// todo 在有牧师的情况下，哥布林切换
				//if i := pdIndex(d.goblins[:], newP); i >= 0 {
				//	newData := d
				//	newData.goblins[i].dir ^= dirStoneDelta
				//	add(d, newData, dir4String[dIdx]+"C") // trans
				//	continue
				//}

				// 喷火龙切换
				if i := pdIndex(d.dragons[:], newP); i >= 0 {
					newData := d
					newData.dragons[i].dir ^= dirStoneDelta
					add(d, newData, dir4String[dIdx]+"C") // trans
					continue
				}

				// 普通移动一步
				if !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) {
					continue
				}
				newData := d
				newData.druid = newP
				add(d, newData, dir4String[dIdx]) // move
			}
		case charExplorer:
			// 普通移动一步
			p0 := d.explorer
			for dIdx, dir := range directions4 {
				newP := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				if !d.isValidPos(newP) {
					if mapSizeH > 1 {
						// 如果头上有喷火龙或者镜子，修改其朝向
						if i := pdIndex(d.dragons[:], point{p0.x, p0.y, p0.z + 1}); i >= 0 {
							newData := d
							newData.dragons[i].dir &^= 7
							newData.dragons[i].dir |= uint8(dIdx)
							add(d, newData, dir4String[dIdx])
						}
						// todo 镜子
					}
					continue // 枚举另一个方向
				}

				newData := d
				if allowExplorerPushItem {
					if i := slices.Index(allMovableObjs, newP); i >= 0 {
						// 推物品
						nxt2 := point{newP.x + dir.x, newP.y + dir.y, newP.z + dir.z}
						if !d.isValidPos(nxt2) || slices.Contains(allMovableObjs, nxt2) {
							continue // 枚举另一个方向
						}
						newData.changePos(newP, nxt2, math.MaxUint8)
					}
				}

				newData.explorer = newP
				if mapSizeH > 1 {
					oldTop := point{p0.x, p0.y, p0.z + 1}
					// 如果原位置头上有喷火龙或者镜子，修改其位置和朝向
					if i := pdIndex(newData.dragons[:], oldTop); i >= 0 {
						newTop := newP
						newTop.z++
						if !d.isValidPos(newTop) || slices.Contains(allMovableObjs, newTop) {
							continue // todo 暂时禁止喷火龙落地 
						}
						newData.dragons[i] = pointWithDir{newTop, uint8(dIdx)}
					} else if slices.Contains(allMovableObjs, oldTop) {
						newTop := newP
						newTop.z++
						newData.changePos(oldTop, newTop, uint8(dIdx))
					}
					// todo 镜子
				}
				add(d, newData, dir4String[dIdx])
			}
		case charSailor:
			// 普通移动一步
			p0 := d.sailor
			for dIdx, dir := range directions4 {
				newP := point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
				if !d.isValidPos(newP) {
					if mapSizeH > 1 {
						// 如果头上有喷火龙或者镜子，修改其朝向
						if i := pdIndex(d.dragons[:], point{p0.x, p0.y, p0.z + 1}); i >= 0 {
							newData := d
							newData.dragons[i].dir &^= 7
							newData.dragons[i].dir |= uint8(dIdx)
							add(d, newData, dir4String[dIdx])
						}
						// todo 镜子
					}
					continue // 枚举另一个方向
				}

				newData := d
				if allowPushItem {
					if i := slices.Index(allMovableObjs, newP); i >= 0 {
						// 推物品
						nxt2 := point{newP.x + dir.x, newP.y + dir.y, newP.z + dir.z}
						if !d.isValidPos(nxt2) || slices.Contains(allMovableObjs, nxt2) {
							continue // 枚举另一个方向
						}
						newData.changePos(newP, nxt2, math.MaxUint8)
					}
				}

				newData.sailor = newP
				if mapSizeH > 1 {
					oldTop := point{p0.x, p0.y, p0.z + 1}
					// 如果原位置头上有喷火龙或者镜子，修改其位置和朝向
					if i := pdIndex(newData.dragons[:], oldTop); i >= 0 {
						newTop := newP
						newTop.z++
						if !d.isValidPos(newTop) || slices.Contains(allMovableObjs, newTop) {
							continue // todo 暂时禁止喷火龙落地 
						}
						newData.dragons[i] = pointWithDir{newTop, uint8(dIdx)}
					} else if slices.Contains(allMovableObjs, oldTop) {
						newTop := newP
						newTop.z++
						newData.changePos(oldTop, newTop, uint8(dIdx))
					}
					// todo 镜子
				}
				add(d, newData, dir4String[dIdx])
			}
		case charMerchant:
			// 多控
			// 普通移动一步
			for dIdx, dir := range directions4 {
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
						newData.changePos(nxt, nxt2, math.MaxUint8)
					}
					moved = true
					man[manIdx] = nxt // 移走！
				}
				if !moved { // 没人动
					continue
				}
				add(d, newData, dir4String[dIdx])
			}
		default:
			// 跳石
			oriChar := d.curCharTypeNum - skippingStoneDelta
			_ = oriChar

		}

		// 换成其他人
		if isBigMap {
			p := d.getCurCharTypePos()
			ch := levelMap[0][p.x][p.y]
			if ch != charNumToName[d.curCharTypeNum] {
				if i := strings.IndexByte("ATWCDB789", ch); i >= 0 {
					newData := d
					// 复原所有角色位置
					newData.warrior = __warrior
					newData.thief = __thief
					newData.wizard = __wizard
					newData.cleric = __cleric
					newData.druid = __druid
					newData.bard = __bard
					newData.explorer = __explorer
					newData.sailor = __sailor
					newData.curCharTypeNum = int8(i + 1)
					add(d, newData, "IGNORE")
				}
			}
		} else {
			for _, char := range validChars {
				if char != d.curCharTypeNum {
					newData := d
					newData.curCharTypeNum = char
					var info string
					if len(allChars) > 2 {
						info = digits[char : char+1]
					} else {
						info = "c"
					}
					if d.curCharTypeNum == charBard {
						info = "B" + info // 等一下再换人
					}
					add(d, newData, info)
				}
			}
		}
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

var charNumToName = [...]byte{
	charWarrior:  'A',
	charThief:    'T',
	charWizard:   'W',
	charCleric:   'C',
	charBard:     'B',
	charDruid:    'D',
	charExplorer: '7',
	charSailor:   '8',
	charMerchant: '9',
}

// 跳石，无法操纵，只能原地等待
// 当跳石被推动后，额外进入该角色
// 当跳石停止移动后，换回原来的角色（用 skippingStoneDelta + 原来的角色编号表示跳石的情况）
const skippingStoneDelta = 1 << 6

const (
	beamDefault   = iota
	beamOpen      // 红
	beamDouble    // 橙
	beamDestroy   // 黄
	beamPenetrate // 绿
	beamPush      // 青
	beamTeleport  // 紫
)

const dirStoneDelta = 1 << 6
