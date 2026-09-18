package oss

import (
	"fmt"
	"math"
	"math/bits"
	"slices"
	"strings"
	"unsafe"
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

type warriorArrType [warriorNumberInit]point
type thiefArrType [thiefNumberInit]point
type wizardArrType [wizardNumberInit]point
type priestArrType [priestNumberInit]point
type druidArrType [druidNumberInit]point
type bardArrType [bardNumberInit]point
type explorerArrType [explorerNumberInit]point
type sailorArrType [sailorNumberInit]point
type merchantArrType [merchantNumberInit]point

type stoneArrType [stoneNumberInit]point
type crystalArrType [crystalNumberInit + grassNumberInit]point
type grassArrType [(crystalNumberInit + grassNumberInit) * min(druidNumberInit, 1)]point
type skippingStoneArrType [skippingStoneNumberInit]pointWithDir
type skippingCrystalArrType [skippingCrystalNumberInit]pointWithDir
type lilyArrType [lilyNumberInit]pointWithDir
type goblinArrType [goblinNumberInit]pointWithDir
type dragonArrType [len(dragonDirInit)]pointWithDir
type beamArrType [len(beamDirInit)]pointWithDir
type mirrorArrType [len(mirrorDirInit) / 2]pointWithDir
type mirrorRefArrType [len(mirrorRefDirInit) / 2]pointWithDir
type mirrorAuxArrType [len(mirrorAuxDirInit) / 2]pointWithDir

type data struct {
	warrior  warriorArrType  // A 推多个对象
	thief    thiefArrType    // T 拉一个对象
	wizard   wizardArrType   // W 交换对象
	cleric   priestArrType   // C 自己以及上下左右无敌
	druid    druidArrType    // D 把对象变成石头
	bard     bardArrType     // B 同时移动切比雪夫距离 <= 2 的对象
	explorer explorerArrType // 7 普通角色，无法推对象
	sailor   sailorArrType   // 8 普通角色，推一个对象
	merchant merchantArrType // 9 普通角色，推一个对象

	// 石头   挡光，无法被反射
	stones stoneArrType // S

	// 水晶   透明，可以被反射
	crystals crystalArrType // s

	// 草
	grass grassArrType // w

	// 水漂石
	skippingStones   skippingStoneArrType   // K
	skippingCrystals skippingCrystalArrType // k

	// 睡莲叶，z=-1，放在 waterMap 中
	lilies lilyArrType // l

	// 怪物
	goblins goblinArrType // g
	dragons dragonArrType // d

	// 镜子
	mirrors     mirrorArrType    // M
	mirrorRefs  mirrorRefArrType // R 可以被反射的镜子
	mirrorAuxes mirrorAuxArrType // m 关卡名中称其为 mundane

	// 光束
	// 高 4 位是类型，低 4 位是方向
	beams beamArrType // b

	// 哪些角色变成了水晶
	isCrystalMask [min(druidNumberInit, 1)]uint8 // uint8 不支持 merchant

	// todo 用于判断牧师是否漂浮
	//isPriestAttacked bool

	// 门的开闭，避免反复计算
	doorOpened        [doorKinds]bool
	monsterDoorOpened bool

	// 当前角色类型
	curCharTypeNum int8
}

const mapSizeH = int8(len(levelMap))

var mapSizeN, mapSizeM int8
var hasWater = false

// 初始化变量、检查地图是否与 const 匹配
func initMap() {
	fmt.Println("data 结构体大小:", unsafe.Sizeof(data{}), "bytes")
	fmt.Println()

	mapSizeN = int8(len(levelMap[0]))
	mapSizeM = int8(len(levelMap[0][0]))

	for i, p := range finals {
		finals[i] = changeNegPoint(p)
	}
	for i, p := range monsterDoors {
		monsterDoors[i] = changeNegPoint(p)
	}
	for _, ps := range doors {
		for i, p := range ps {
			ps[i].point = changeNegPoint(p.point)
		}
	}
	for _, ps := range switches {
		for i, p := range ps {
			ps[i] = changeNegPoint(p)
		}
	}
	for i, p := range stonePosInit {
		stonePosInit[i] = changeNegPoint(p)
	}
	for i, p := range lilyPosInit {
		if p.z != -1 {
			panic("lilyPosInit[i].z 必须是 -1")
		}
		lilyPosInit[i] = changeNegPoint(p)
	}

	if warriorPosInit != noPos {
		warriorPosInit = changeNegPoint(warriorPosInit)
	}
	if thiefPosInit != noPos {
		thiefPosInit = changeNegPoint(thiefPosInit)
	}
	if wizardPosInit != noPos {
		wizardPosInit = changeNegPoint(wizardPosInit)
	}
	if priestPosInit != noPos {
		priestPosInit = changeNegPoint(priestPosInit)
	}
	if bardPosInit != noPos {
		bardPosInit = changeNegPoint(bardPosInit)
	}
	if sailorPosInit != noPos {
		sailorPosInit = changeNegPoint(sailorPosInit)
	}

	var doorMask, switchMask, finalNum int
	var warriorNum, thiefNum, wizardNum, priestNum, druidNum, bardNum, explorerNum, sailorNum, merchantNum int
	var stoneNum, crystalNum, skippingStoneNum, skippingCrystalNum, grassesNum, lilyNum, beamNum, mirrorNum, mirrorRefNum, mirrorAuxNum int
	var goblinNum, dragonNum int

	finalNum += len(finals)
	for i, ps := range doors {
		if len(ps) > 0 {
			doorMask |= 1 << i
		}
	}
	for i, ps := range switches {
		if len(ps) > 0 {
			switchMask |= 1 << i
		}
	}

	if warriorPosInit != noPos {
		warriorNum++
	}
	if thiefPosInit != noPos {
		thiefNum++
	}
	if wizardPosInit != noPos {
		wizardNum++
	}
	if priestPosInit != noPos {
		priestNum++
	}
	if bardPosInit != noPos {
		bardNum++
	}
	if sailorPosInit != noPos {
		sailorNum++
	}

	stoneNum += len(stonePosInit)
	lilyNum += len(lilyPosInit)

	checkGrid := func(grid []string) {
		for _, row := range grid {
			if len(row) != int(mapSizeM) {
				panic("行不等长")
			}
			for _, ch := range row {
				switch ch {
				case 'A':
					warriorNum++
				case 'T':
					thiefNum++
				case 'W':
					wizardNum++
				case 'C':
					priestNum++
				case 'D':
					druidNum++
				case 'B':
					bardNum++
				case '7':
					explorerNum++
				case '8':
					sailorNum++
				case '9':
					merchantNum++
				case 'S':
					stoneNum++
				case 's':
					crystalNum++
				case 'w':
					grassesNum++
				case 'K':
					skippingStoneNum++
				case 'k':
					skippingCrystalNum++
				case 'l':
					lilyNum++
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
				case 'f':
					finalNum++
				case 'x', 'y', 'z', '{':
					switchMask |= 1 << (ch - 'x')
				case 'X', 'Y', 'Z', '[':
					doorMask |= 1 << (ch - 'X')
				case '~', '^', 'v', '<', '>': // 水
					hasWater = true
				case '#', '.', 'e':
					// ignore
				default:
					fmt.Printf("【警告】没有处理字符 %c\n", ch)
				}
			}
		}
	}
	if len(waterMap) > 0 {
		checkGrid(waterMap)
	}
	for _, grid := range levelMap {
		checkGrid(grid)
	}

	if finalNum != allCharNum {
		panic("地图 'f' 设置错误")
	}
	if bits.OnesCount(uint(doorMask)) != doorKinds {
		panic("没有修改 door kinds")
	}
	if bits.OnesCount(uint(switchMask)) != doorKinds {
		panic("压力开关错误，请检查地图")
	}

	if warriorNum != warriorNumberInit {
		panic("没有修改 warrior number")
	}
	if thiefNum != thiefNumberInit {
		panic("没有修改 thief number")
	}
	if wizardNum != wizardNumberInit {
		panic("没有修改 wizard number")
	}
	if priestNum != priestNumberInit {
		panic("没有修改 priest number")
	}
	if druidNum != druidNumberInit {
		panic("没有修改 druid number")
	}
	if bardNum != bardNumberInit {
		panic("没有修改 bard number")
	}
	if explorerNum != explorerNumberInit {
		panic("没有修改 explorer number")
	}
	if sailorNum != sailorNumberInit {
		panic("没有修改 sailor number")
	}
	if !allowCloneMan && merchantNum != merchantNumberInit {
		panic("没有修改 merchant number")
	}

	// 检查数组大小是否与 levelMap 匹配
	if stoneNum != stoneNumberInit {
		panic("没有修改 stone number")
	}
	if crystalNum != crystalNumberInit {
		panic("没有修改 crystal number")
	}
	if grassesNum != grassNumberInit {
		panic("没有修改 grass number")
	}
	if skippingStoneNum != skippingStoneNumberInit {
		panic("没有修改 skipping stone number")
	}
	if skippingCrystalNum != skippingCrystalNumberInit {
		panic("没有修改 skipping crystal number")
	}
	if lilyNum != lilyNumberInit {
		panic("没有修改 lily number")
	}
	if goblinNum != len(goblinArrType{}) {
		panic("没有修改 goblin number")
	}
	if dragonNum != len(dragonArrType{}) {
		panic("没有修改 dragon dir")
	}
	if beamNum != len(beamArrType{}) {
		panic("没有修改 beamDirInit")
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
}

func hasFence(p, dir point) bool {
	grid := levelMap[p.z]
	switch {
	case dir.y == -1: // 左
		return grid[p.x][p.y] == '|' || grid[p.x][p.y] == 'L'
	case dir.y == 1: // 右
		return grid[p.x][p.y+1] == '|' || grid[p.x][p.y+1] == 'L'
	case dir.x == -1: // 上
		return grid[p.x-1][p.y] == '_' || grid[p.x-1][p.y] == 'L'
	case dir.x == 1: // 下
		return grid[p.x][p.y] == '_' || grid[p.x][p.y] == 'L'
	default:
		panic("不支持的移动")
	}
}

// 所有怪物都杀死或者变成水晶
func (d *data) areAllMonstersDied() bool {
	for _, p := range d.goblins {
		if p.point != noPos && p.dir&dirIsCrystal == 0 { // 没有变成水晶
			return false
		}
	}
	for _, p := range d.dragons {
		if p.point != noPos && p.dir&dirIsCrystal == 0 { // 没有变成水晶
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

	allChars := make([]point, 0, allCharNum)
	if warriorNumberInit > 0 {
		for _, p := range d.warrior {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if thiefNumberInit > 0 {
		for _, p := range d.thief {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if wizardNumberInit > 0 {
		for _, p := range d.wizard {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if priestNumberInit > 0 {
		for _, p := range d.cleric {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if druidNumberInit > 0 {
		for _, p := range d.druid {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if bardNumberInit > 0 {
		for _, p := range d.bard {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if explorerNumberInit > 0 {
		for _, p := range d.explorer {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if sailorNumberInit > 0 {
		for _, p := range d.sailor {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	if merchantNumberInit > 0 {
		for _, p := range d.merchant {
			if p != noPos {
				allChars = append(allChars, p)
			}
		}
	}
	return allChars
}

// onlyLife=true 表示只限于可被诗人魅惑的物品
func (d *data) getAllMovableObjPos(isBigMap bool, onlyLife bool) (all, chars, nonChars []point) {
	chars = d.getAllCharPos(isBigMap)
	all = chars
	if mirrorDirInit != "" && !onlyLife {
		for _, p := range d.mirrors {
			if p.point != noPos {
				all = append(all, p.point)
			}
		}
	}
	if mirrorRefDirInit != "" && !onlyLife {
		for _, p := range d.mirrorRefs {
			if p.point != noPos {
				all = append(all, p.point)
			}
		}
	}
	if mirrorAuxDirInit != "" && !onlyLife {
		for _, p := range d.mirrorAuxes {
			if p.point != noPos {
				all = append(all, p.point)
			}
		}
	}
	if !onlyLife {
		for _, p := range d.stones {
			if p != noPos {
				all = append(all, p)
			}
		}
	}
	for _, p := range d.crystals {
		if p != noPos {
			all = append(all, p)
		}
	}
	if !onlyLife {
		for _, p := range d.skippingStones {
			if p.point != noPos {
				all = append(all, p.point)
			}
		}
	}
	for _, p := range d.skippingCrystals {
		if p.point != noPos {
			all = append(all, p.point)
		}
	}
	if goblinNumberInit > 0 {
		for _, p := range d.goblins {
			if p.point != noPos {
				all = append(all, p.point)
			}
		}
	}
	if dragonDirInit != "" {
		for _, p := range d.dragons {
			if p.point != noPos {
				all = append(all, p.point)
			}
		}
	}
	if beamDirInit != "" && !onlyLife {
		for _, p := range d.beams {
			if p.point != noPos {
				all = append(all, p.point)
			}
		}
	}
	return all, chars, all[len(chars):]
}

func (d *data) getAllLife(isBigMap bool) (life, nonLife []point) {
	life = d.getAllCharPos(isBigMap)
	if mirrorDirInit != "" {
		for _, p := range d.mirrors {
			if p.point != noPos {
				nonLife = append(nonLife, p.point)
			}
		}
	}
	if mirrorRefDirInit != "" {
		for _, p := range d.mirrorRefs {
			if p.point != noPos {
				nonLife = append(nonLife, p.point)
			}
		}
	}
	if mirrorAuxDirInit != "" {
		for _, p := range d.mirrorAuxes {
			if p.point != noPos {
				nonLife = append(nonLife, p.point)
			}
		}
	}
	for _, p := range d.stones {
		if p != noPos {
			nonLife = append(nonLife, p)
		}
	}
	for _, p := range d.crystals {
		if p != noPos {
			life = append(life, p)
		}
	}
	for _, p := range d.skippingStones {
		if p.point != noPos {
			nonLife = append(nonLife, p.point)
		}
	}
	for _, p := range d.skippingCrystals {
		if p.point != noPos {
			life = append(life, p.point)
		}
	}
	if goblinNumberInit > 0 {
		for _, p := range d.goblins {
			if p.point != noPos {
				life = append(life, p.point)
			}
		}
	}
	if dragonDirInit != "" {
		for _, p := range d.dragons {
			if p.point != noPos {
				life = append(life, p.point)
			}
		}
	}
	if beamDirInit != "" {
		for _, p := range d.beams {
			if p.point != noPos {
				nonLife = append(nonLife, p.point)
			}
		}
	}
	return
}

func inBound(p point) bool {
	return 0 <= p.x && p.x < mapSizeN &&
		0 <= p.y && p.y < mapSizeM &&
		p.z < mapSizeH
}

func (d *data) inAnyClosedDoors(p point) bool {
	if !d.monsterDoorOpened && slices.Contains(monsterDoors[:], p) { // 怪物门
		return true
	}
	for i, opened := range d.doorOpened {
		if !opened && pdContains(doors[i], p) { // 活塞门
			return true
		}
	}
	return false
}

// p 是空气、水、电梯、可移动对象 -> true
// p 是出界、墙、草、怪物门、活塞门 -> false   todo 植物的根
func (d *data) isValidPos(p point) bool {
	if !inBound(p) || // 出界
		p.z >= 0 && levelMap[p.z][p.x][p.y] == '#' || // 墙
		len(d.grass) > 0 && slices.Contains(d.grass[:], p) || // 草
		d.inAnyClosedDoors(p) { // 怪物门、活塞门（包括水中的门）
		return false
	}
	return true
}

// 返回 mask 表示在哪些类型的 beam 中
// endPoints（若 mask 有 endpoint）0 为发射器一端，1 为远端
// todo 多个 endPoints
type beamInfo struct {
	endPoints   [2]point
	endPointDir point
}

// todo 镜子
func (d *data) withinBeams(p point, allNonCharObjs []point) (typeMask uint16, beamNf beamInfo) {
	// todo 前提是钻石在正确位置上

	for _, beam := range d.beams {
		// beam.dir 高 4 位是类型，低 4 位是方向
		beamDir := directions6[beam.dir&0xf]

		const hasMirror = mirrorDirInit != "" || mirrorRefDirInit != "" || mirrorAuxDirInit != ""
		if !hasMirror {
			// 剪枝：先粗略判断是否在光束方向上（不考虑障碍）
			if beamDir.x != 0 {
				// 上下，必须同 y 同 z
				if beam.y != p.y || beam.z != p.z {
					continue
				}
				if beamDir.x > 0 != (beam.x < p.x) {
					continue
				}
			} else if beamDir.y != 0 {
				// 左右，必须同 x 同 z
				if beam.x != p.x || beam.z != p.z {
					continue
				}
				if beamDir.y > 0 != (beam.y < p.y) {
					continue
				}
			} else { // beamDir.z != 0
				// 高低，必须同 x 同 y
				if beam.x != p.x || beam.y != p.y {
					continue
				}
			}
		}

		cur := beam.point
		meetP := false
		for {
			old := cur
			cur = cur.add(beamDir)
			if cur == p {
				meetP = true
				if beam.dir>>4 != beamEndpoint {
					break
				}
				// 继续走，走到末端
			}
			// 特判：水晶、怪物可以穿透
			// todo 镜子能从背面穿透吗？
			if slices.Contains(d.crystals[:], cur) || pdContains(d.dragons[:], cur) || pdContains(d.goblins[:], cur) {
				continue
			}
			// 出界，或者遇到不可穿透对象（石头、钻石、门）
			// 可以在地图边界加一圈 '#' 
			if !inBound(cur) || // 出界
				slices.Contains(allNonCharObjs, cur) || // 水晶在上面判断了
				d.inAnyClosedDoors(cur) { // 怪物门、活塞门
				cur = old
				break
			}
		}

		if meetP {
			beamType := beam.dir >> 4
			typeMask |= 1 << beamType
			if beamType == beamEndpoint {
				beamNf.endPoints[0] = beam.point.add(beamDir)
				beamNf.endPoints[1] = cur
				beamNf.endPointDir = beamDir
			}
		}
	}
	return
}

// 是否被牧师保护（或者自己是牧师）
func (d *data) isProtected(char point) bool {
	if priestNumberInit == 0 {
		return false
	}
	priest := d.cleric[:][0] // todo 多个牧师
	if char == priest {
		return true
	}
	if mapSizeH > 1 {
		return isNeighbor6(char, priest)
	}
	return isNeighbor4(char, priest)
}

// 在水面上（z=0）且下面（z=-1）没有物品（或者门）的对象，落入水中
// todo 摧毁水中的镜子
func (d *data) isFallIntoWater(p point) bool {
	if !hasWater ||
		p.z != 0 { // todo z > 0 中途遇到障碍
		return false
	}
	switch levelMap[0][p.x][p.y] {
	case '~', '^', 'v', '<', '>', 'l': // 水
		// 继续
	default:
		return false
	}

	downP := point{p.x, p.y, -1}
	// 水中的物品（石头、水晶、水漂石、睡莲叶）
	// todo 水平栏杆？
	if len(d.stones) > 0 && slices.Contains(d.stones[:], downP) ||
		len(d.crystals) > 0 && slices.Contains(d.crystals[:], downP) ||
		len(d.dragons) > 0 && len(d.druid) > 0 && pdContains(d.dragons[:], downP) ||
		len(d.goblins) > 0 && len(d.druid) > 0 && pdContains(d.goblins[:], downP) ||
		len(d.skippingStones) > 0 && pdContains(d.skippingStones[:], downP) ||
		len(d.skippingCrystals) > 0 && pdContains(d.skippingCrystals[:], downP) ||
		len(d.lilies) > 0 && pdContains(d.lilies[:], downP) ||
		d.inAnyClosedDoors(downP) { // 水中的门
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
		if g.point == noPos || g.dir&dirIsCrystal > 0 { // 是石头
			continue
		}
		if mapSizeH > 1 {
			if isNeighbor6(g.point, p) { // todo 是这样吗？
				return true
			}
		} else {
			if isNeighbor4(g.point, p) {
				return true
			}
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
		if !opened && pdContains(doors[i], p) {
			return dieTypeCrushed
		}
	}

	// 被攻击的优先级更高
	if d.isAttacked(p, burnPos) {
		if isChar && (d.isProtected(p) || len(d.isCrystalMask) > 0 && d.isCrystalMask[:][0] > 0 &&
			d.isCrystalMask[:][0]>>(d.getCharType(p)-1)&1 > 0) {
			return dieTypeNo // 注：如果下面是空或者水，不会落下去
		}
		return dieTypeAttacked
	}

	// 淹死
	if d.isFallIntoWater(p) {
		// todo 如果自己是牧师且周围人被攻击，那么自己是悬浮的，不会淹死
		// todo 牧师不会落水？
		if priestNumberInit > 0 && isChar && p == d.cleric[:][0] {
			return dieTypeNo
		}
		return dieTypeDrown
	}

	return dieTypeNo
}

func (d *data) getCharType(p point) uint8 {
	switch {
	case len(d.warrior) > 0 && d.warrior[:][0] == p:
		return charWarrior
	case len(d.thief) > 0 && d.thief[:][0] == p:
		return charThief
	case len(d.wizard) > 0 && d.wizard[:][0] == p:
		return charWizard
	case len(d.cleric) > 0 && d.cleric[:][0] == p:
		return charCleric
	case len(d.druid) > 0 && d.druid[:][0] == p:
		return charDruid
	case len(d.bard) > 0 && d.bard[:][0] == p:
		return charBard
	case len(d.explorer) > 0 && d.explorer[:][0] == p:
		return charExplorer
	case len(d.sailor) > 0 && d.sailor[:][0] == p:
		return charSailor
	case len(d.merchant) > 0 && d.merchant[:][0] == p:
		return charMerchant
	default:
		return 0
	}
}

// 反射：从 mirror.point 出发，往 dir 方向走 step 步
// todo 原方向在多次反射后的新方向（喷火龙、可被反射的镜子）
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
				if step == math.MaxInt { // 法师 todo 喷火龙
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
				if step == math.MaxInt { // 法师 todo 喷火龙
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
				if step == math.MaxInt { // 法师 todo 喷火龙
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

func (d *data) changeSkippingPos(skipping []pointWithDir, oldP, newP point, allMovableObjs []point) bool {
	i := pdIndex(skipping, oldP)
	if i < 0 {
		return false
	}

	// newP 下面不是水（例如睡莲叶），只修改位置
	if !d.isFallIntoWater(newP) {
		skipping[i] = pointWithDir{newP, dirStop}
		return true
	}

	// 注意，镜子反射的情况已经在 doMirror 中单独处理了
	dir := newP.sub(oldP)

	if !slices.Contains(directions4, dir) {
		skipping[i].point = newP
		if skipping[i].dir != dirStop {
			// todo 法师交换？移位杖？
			panic("未实现")
		}
		return true
	}

	// newP 在水上，开始移动
	if isSkippingAndLilySlow {
		// 一步一步走
		// 只修改 dir，位置仍然是 oldP，我们会在入队的时候修改位置（when isLilySlow = true）
		skipping[i].dir = getDirIndexByDir(dir)

	} else {
		// 走到底
		type pair struct{ point, dir point }
		lilies := []pair{}
		cur := newP
		for {
			// 出界
			if !inBound(cur) {
				newP = noPos
				break
			}

			// 撞上物品
			if !d.isValidPos(cur) || slices.Contains(allMovableObjs, cur) {
				// 回到前一个位置，然后落水
				newP = cur.sub(dir)
				newP.z = -1
				break
			}

			// 下面不是水
			if !d.isFallIntoWater(cur) {
				// 停在这里
				newP = cur // todo

				p := cur
				p.z = -1
				if j := pdIndex(d.lilies[:], p); j >= 0 {
					// 停在了睡莲叶上
					lilies = append(lilies, pair{p, dir})
				}
				break
			}

			// 收集遇到的睡莲叶
			if dir.x != 0 {
				p := point{cur.x, cur.y - 1, -1}
				if j := pdIndex(d.lilies[:], p); j >= 0 {
					lilies = append(lilies, pair{p, point{0, -1, 0}})
				}
				p = point{cur.x, cur.y + 1, -1}
				if j := pdIndex(d.lilies[:], p); j >= 0 {
					lilies = append(lilies, pair{p, point{0, 1, 0}})
				}
			} else {
				p := point{cur.x - 1, cur.y, -1}
				if j := pdIndex(d.lilies[:], p); j >= 0 {
					lilies = append(lilies, pair{p, point{-1, 0, 0}})
				}
				p = point{cur.x + 1, cur.y, -1}
				if j := pdIndex(d.lilies[:], p); j >= 0 {
					lilies = append(lilies, pair{p, point{1, 0, 0}})
				}
			}

			cur = cur.add(dir)
		}

		skipping[i].point = newP

		// 同时处理睡莲叶的移动
		// 睡莲叶遇到纯水（不能有其他睡莲叶或石头）才继续移动，其余情况都会停下
		// todo 先简单点，倒着遍历遇到的睡莲   或者粗略地按照移动距离排序
		for j := len(lilies) - 1; j >= 0; j-- {
			lDir := lilies[j].dir
			p0 := lilies[j].point
			cur := p0.add(lDir)
			for {
				// 出界
				if !inBound(cur) {
					cur = noPos
					break
				}
				// 不是纯水
				if !d.isFallIntoWater(point{cur.x, cur.y, 0}) {
					cur = cur.sub(lDir)
					break
				}
				ch := levelMap[0][cur.x][cur.y]
				switch ch {
				case '^':
					cur.x--
					lDir = point{-1, 0, 0} // 最后一步上岸需要用到
				case 'v':
					cur.x++
					lDir = point{1, 0, 0}
				case '<':
					cur.y--
					lDir = point{0, -1, 0}
				case '>':
					cur.y++
					lDir = point{0, 1, 0}
				default:
					cur = cur.add(lDir)
				}
			}
			if cur == p0 { // 不动
				continue
			}
			d.lilies[j].point = cur
			if cur == noPos {
				continue
			}
			// 睡莲叶承载的物品（如果有）也移动
			d.changePos(point{p0.x, p0.y, p0.z + 1}, point{cur.x, cur.y, cur.z + 1},
				dirIgnore, nil) // 上面的东西不会被阻挡，直接 nil
		}
	}

	return true
}

// todo 添加一个参数 alsoMoveTop bool，
//      使得当物品移动时，物品上方的物品（如果有）也跟着移动
// 如果只是普通推物品，那么 newDir = math.MaxUint8
func (d *data) changePos(oldP, newP point, newDir uint8, allMovableObjs []point) (changed bool) {
	// 人
	if warriorNumberInit > 0 {
		if i := slices.Index(d.warrior[:], oldP); i >= 0 {
			d.warrior[i] = newP
			return true
		}
	}
	if thiefNumberInit > 0 {
		if i := slices.Index(d.thief[:], oldP); i >= 0 {
			d.thief[i] = newP
			return true
		}
	}
	if wizardNumberInit > 0 {
		if i := slices.Index(d.wizard[:], oldP); i >= 0 {
			d.wizard[i] = newP
			return true
		}
	}
	if priestNumberInit > 0 {
		if i := slices.Index(d.cleric[:], oldP); i >= 0 {
			d.cleric[i] = newP
			return true
		}
	}
	if druidNumberInit > 0 {
		if i := slices.Index(d.druid[:], oldP); i >= 0 {
			d.druid[i] = newP
			return true
		}
	}
	if bardNumberInit > 0 {
		if i := slices.Index(d.bard[:], oldP); i >= 0 {
			d.bard[i] = newP
			return true
		}
	}
	if sailorNumberInit > 0 {
		if i := slices.Index(d.sailor[:], oldP); i >= 0 {
			d.sailor[i] = newP
			return
		}
	}
	if explorerNumberInit > 0 {
		if i := slices.Index(d.explorer[:], oldP); i >= 0 {
			d.explorer[i] = newP
			return true
		}
	}
	if merchantNumberInit > 0 {
		if i := slices.Index(d.merchant[:], oldP); i >= 0 {
			d.merchant[i] = newP
			return true
		}
	}

	// 物
	if mirrorDirInit != "" {
		if i := pdIndex(d.mirrors[:], oldP); i >= 0 {
			d.mirrors[i].point = newP
			if newDir&dirIgnore == 0 {
				d.mirrors[i].dir = newDir
			}
			return true
		}
	}
	if mirrorRefDirInit != "" {
		if i := pdIndex(d.mirrorRefs[:], oldP); i >= 0 {
			d.mirrorRefs[i].point = newP
			if newDir&dirIgnore == 0 {
				d.mirrorRefs[i].dir = newDir
			}
			return
		}
	}
	if mirrorAuxDirInit != "" {
		if i := pdIndex(d.mirrorAuxes[:], oldP); i >= 0 {
			d.mirrorAuxes[i].point = newP
			if newDir&dirIgnore == 0 {
				d.mirrorAuxes[i].dir = newDir
			}
			return true
		}
	}

	if len(d.stones) > 0 {
		if i := slices.Index(d.stones[:], oldP); i >= 0 {
			d.stones[i] = newP
			return true
		}
	}

	if len(d.crystals) > 0 {
		if i := slices.Index(d.crystals[:], oldP); i >= 0 {
			d.crystals[i] = newP
			return true
		}
	}

	// 水漂石
	if skippingStoneNumberInit > 0 {
		if d.changeSkippingPos(d.skippingStones[:], oldP, newP, allMovableObjs) {
			return true
		}
	}
	if skippingCrystalNumberInit > 0 {
		if d.changeSkippingPos(d.skippingCrystals[:], oldP, newP, allMovableObjs) {
			return true
		}
	}

	if goblinNumberInit > 0 && canPushGoblin {
		if i := pdIndex(d.goblins[:], oldP); i >= 0 {
			d.goblins[i].point = newP
			return true
		}
	}

	if dragonDirInit != "" && canPushDragon {
		if i := pdIndex(d.dragons[:], oldP); i >= 0 {
			d.dragons[i].point = newP
			if newDir&dirIgnore == 0 {
				d.dragons[i].dir &^= 7
				d.dragons[i].dir |= newDir
			}
			return true
		}
	}

	if beamDirInit != "" && canPushBeam {
		if i := pdIndex(d.beams[:], oldP); i >= 0 {
			d.beams[i].point = newP
			if newDir&dirIgnore == 0 {
				d.beams[i].dir = newDir
			}
			return true
		}
	}

	if newDir != dirIgnore {
		panic("没有发生修改，请检查代码")
	}

	return false
}

func (d *data) getCurCharPos() (pos point) {
	switch d.curCharTypeNum {
	case charDefault:
		panic("代码有误，当前角色不能为 charDefault")
	case charWarrior:
		pos = d.warrior[:][0]
	case charThief:
		pos = d.thief[:][0]
	case charWizard:
		pos = d.wizard[:][0]
	case charCleric:
		pos = d.cleric[:][0]
	case charDruid:
		pos = d.druid[:][0]
	case charBard:
		pos = d.bard[:][0]
	case charExplorer:
		pos = d.explorer[:][0]
	case charSailor:
		pos = d.sailor[:][0]
	case charMerchant:
		pos = d.merchant[:][0]
	default:
		panic("未找到当前角色")
	}
	return
}

// 进入的时候切回 '8'，离开的时候才切换角色
func (newData *data) bigMapForceSwapChar(oldP, newP point) {
	if !isBigMap {
		return
	}

	isOldOutside := strings.ContainsRune("ATWCDB789", rune(levelMap[0][oldP.x][oldP.y]))
	isNewOutside := strings.ContainsRune("ATWCDB789", rune(levelMap[0][newP.x][newP.y]))

	if !isOldOutside && isNewOutside {
		// 从场景内部移到场景外部
		// 重置所有人的位置，除了 8
		if warriorNumberInit > 0 {
			newData.warrior[:][0] = noPos
		}
		if thiefNumberInit > 0 {
			newData.thief[:][0] = noPos
		}
		if wizardNumberInit > 0 {
			newData.wizard[:][0] = noPos
		}
		if priestNumberInit > 0 {
			newData.cleric[:][0] = noPos
		}
		if druidNumberInit > 0 {
			newData.druid[:][0] = noPos
		}
		if bardNumberInit > 0 {
			newData.bard[:][0] = noPos
		}
		if explorerNumberInit > 0 {
			newData.explorer[:][0] = noPos
		}
		//if sailorNumberInit > 0 {
		//	newData.sailor[:][0] = noPos
		//}
		if merchantNumberInit > 0 {
			newData.merchant[:][0] = noPos
		}
		newData.curCharTypeNum = charSailor
	} else if isOldOutside && !isNewOutside {
		// 从场景外部移到场景内部
		switch levelMap[0][oldP.x][oldP.y] {
		case 'A':
			newData.sailor[:][0] = noPos
			newData.warrior[:][0] = newP
			newData.curCharTypeNum = charWarrior
		case 'T':
			newData.sailor[:][0] = noPos
			newData.thief[:][0] = newP
			newData.curCharTypeNum = charThief
		case 'W':
			newData.sailor[:][0] = noPos
			newData.wizard[:][0] = newP
			newData.curCharTypeNum = charWizard
		case 'C':
			newData.sailor[:][0] = noPos
			newData.cleric[:][0] = newP
			newData.curCharTypeNum = charCleric
		case 'D':
			newData.sailor[:][0] = noPos
			newData.druid[:][0] = newP
			newData.curCharTypeNum = charDruid
		case 'B':
			newData.sailor[:][0] = noPos
			newData.bard[:][0] = newP
			newData.curCharTypeNum = charBard
		case '7':
			newData.sailor[:][0] = noPos
			newData.explorer[:][0] = newP
			newData.curCharTypeNum = charExplorer
		case '8':
			newData.sailor[:][0] = newP
			newData.curCharTypeNum = charSailor
		case '9':
			newData.sailor[:][0] = noPos
			newData.merchant[:][0] = newP
			newData.curCharTypeNum = charMerchant
		}
	}
}

func solveLevel() []string {
	initMap()

	warriorInitArr := warriorArrType{}
	for i := range warriorInitArr {
		warriorInitArr[i] = noPos
	}
	thiefInitArr := thiefArrType{}
	for i := range thiefInitArr {
		thiefInitArr[i] = noPos
	}
	wizardInitArr := wizardArrType{}
	for i := range wizardInitArr {
		wizardInitArr[i] = noPos
	}
	priestInitArr := priestArrType{}
	for i := range priestInitArr {
		priestInitArr[i] = noPos
	}
	druidInitArr := druidArrType{}
	for i := range druidInitArr {
		druidInitArr[i] = noPos
	}
	bardInitArr := bardArrType{}
	for i := range bardInitArr {
		bardInitArr[i] = noPos
	}
	explorerInitArr := explorerArrType{}
	for i := range explorerInitArr {
		explorerInitArr[i] = noPos
	}
	sailorInitArr := sailorArrType{}
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
	crystalInitArr := crystalArrType{}
	for i := range crystalInitArr {
		crystalInitArr[i] = noPos
	}
	grassInitArr := grassArrType{}
	for i := range grassInitArr {
		grassInitArr[i] = noPos
	}
	skippingStoneInitArr := skippingStoneArrType{}
	skippingCrystalInitArr := skippingCrystalArrType{}
	lilyInitArr := lilyArrType{}
	goblinInitArr := goblinArrType{}
	dragonInitArr := dragonArrType{}
	beamInitArr := beamArrType{}

	__curCharTypeNum := initCharTypeNum
	if isBigMap {
		__curCharTypeNum = charSailor
	}

	__warriors := warriorInitArr[:0]
	if warriorPosInit != noPos {
		__warriors = append(__warriors, warriorPosInit)
	}
	__thiefs := thiefInitArr[:0]
	if thiefPosInit != noPos {
		__thiefs = append(__thiefs, thiefPosInit)
	}
	__wizards := wizardInitArr[:0]
	if wizardPosInit != noPos {
		__wizards = append(__wizards, wizardPosInit)
	}
	__priests := priestInitArr[:0]
	if priestPosInit != noPos {
		__priests = append(__priests, priestPosInit)
	}
	__druids := druidInitArr[:0]
	__bards := bardInitArr[:0]
	if bardPosInit != noPos {
		__bards = append(__bards, bardPosInit)
	}
	__explorers := explorerInitArr[:0]
	__sailor := sailorPosInit
	__sailors := sailorInitArr[:0]
	if sailorPosInit != noPos {
		__sailors = append(__sailors, sailorPosInit)
	}
	__merchants := merchantInitArr[:0]

	__mirrors := mirrorInitArr[:0]
	__mirrorRefs := mirrorRefInitArr[:0]
	__mirrorAuxes := mirrorAuxInitArr[:0]
	__stones := stoneInitArr[:0]
	for _, p := range stonePosInit {
		__stones = append(__stones, p)
	}
	__crystals := crystalInitArr[:0]
	__grass := grassInitArr[:0]
	__skippingStones := skippingStoneInitArr[:0]
	__skippingCrystals := skippingCrystalInitArr[:0]
	__lilies := lilyInitArr[:0]
	for _, p := range lilyPosInit {
		__lilies = append(__lilies, pointWithDir{p, dirStop})
	}
	__goblins := goblinInitArr[:0]
	__dragons := dragonInitArr[:0]
	__beams := beamInitArr[:0]

	parseGrid := func(z int, grid []string) {
		for x, row := range grid {
			for y, ch := range row {
				p := point{int8(x), int8(y), int8(z)}
				switch ch {
				case 'A':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __curCharTypeNum < 0 {
							__curCharTypeNum = charWarrior
						}
						__warriors = append(__warriors, p)
					}
				case 'T':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __curCharTypeNum < 0 {
							__curCharTypeNum = charThief
						}
						__thiefs = append(__thiefs, p)
					}
				case 'W':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __curCharTypeNum < 0 {
							__curCharTypeNum = charWizard
						}
						__wizards = append(__wizards, p)
					}
				case 'C':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __curCharTypeNum < 0 {
							__curCharTypeNum = charCleric
						}
						__priests = append(__priests, p)
					}
				case 'D':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __curCharTypeNum < 0 {
							__curCharTypeNum = charDruid
						}
						__druids = append(__druids, p)
					}
				case 'B':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __curCharTypeNum < 0 {
							__curCharTypeNum = charBard
						}
						__bards = append(__bards, p)
					}
				case '7':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __curCharTypeNum < 0 {
							__curCharTypeNum = charExplorer
						}
						__explorers = append(__explorers, p)
					}
				case '8':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charSailor
					}
					if __sailor == noPos {
						__sailor = p
					}
					__sailors = append(__sailors, p)
				case '9':
					if __curCharTypeNum < 0 {
						__curCharTypeNum = charMerchant
					}
					__merchants = append(__merchants, p)
				case 'M':
					i := len(__mirrors)
					__mirrors = append(__mirrors, pointWithDir{p, makeMirrorDir(mirrorDirInit[i*2 : i*2+2])})
				case 'R':
					i := len(__mirrorRefs)
					__mirrorRefs = append(__mirrorRefs, pointWithDir{p, makeMirrorDir(mirrorRefDirInit[i*2 : i*2+2])})
				case 'm':
					i := len(__mirrorAuxes)
					__mirrorAuxes = append(__mirrorAuxes, pointWithDir{p, makeMirrorDir(mirrorAuxDirInit[i*2 : i*2+2])})
				case 'S': // 大写，不透光
					__stones = append(__stones, p)
				case 's': // 小写，透光
					__crystals = append(__crystals, p)
				case 'w':
					__grass = append(__grass, p)
				case 'K':
					__skippingStones = append(__skippingStones, pointWithDir{p, dirStop})
				case 'k':
					__skippingCrystals = append(__skippingCrystals, pointWithDir{p, dirStop})
				case 'l':
					p.z = -1 // todo
					__lilies = append(__lilies, pointWithDir{p, dirStop})
				case 'g':
					__goblins = append(__goblins, pointWithDir{p, 0}) // todo 默认方向为 dirs[0]
				case 'd':
					idx := len(__dragons)
					__dragons = append(__dragons, pointWithDir{p, getDir(dragonDirInit[idx])})
				case 'b':
					idx := len(__beams)
					dir := getDir(beamDirInit[idx])
					tp := beamTypeInit[idx] - '0'
					__beams = append(__beams, pointWithDir{p, tp<<4 | dir})
				case 'x', 'y', 'z', '{':
					switches[ch-'x'] = append(switches[ch-'x'], p)
					if also := sameSwitches[ch]; also > 0 {
						switches[also-'x'] = append(switches[also-'x'], p)
					}
				case 'X', 'Y', 'Z', '[':
					dir := getDir('n') // 默认方向向下，除非手动设置 doorDirString
					if len(doors) < len(doorDirString) {
						dir = getDir(doorDirString[len(doors)])
					}
					doors[ch-'X'] = append(doors[ch-'X'], pointWithDir{p, dir})
				case 'N':
					monsterDoors = append(monsterDoors, p)
				case 'f':
					finals = append(finals, p)
				case '.', '#', '|', '_', 'L', 'e':
					// ignore
				case '~', '^', 'v', '<', '>': // 水
					// ignore
				default:
					panic(fmt.Sprintf("不支持的符号 %c", ch))
				}
			}
		}
	}
	if len(waterMap) > 0 {
		parseGrid(-1, waterMap)
	}
	for z, grid := range levelMap {
		parseGrid(z, grid)
	}

	// 有时候会手动添加 finals 的初始值，总体不一定是有序的
	slices.SortFunc(finals, cmpPoint)

	validChars := []int8{}
	if warriorNumberInit > 0 {
		validChars = append(validChars, charWarrior)
	}
	if thiefNumberInit > 0 {
		validChars = append(validChars, charThief)
	}
	if wizardNumberInit > 0 {
		validChars = append(validChars, charWizard)
	}
	if priestNumberInit > 0 {
		validChars = append(validChars, charCleric)
	}
	if druidNumberInit > 0 {
		validChars = append(validChars, charDruid)
	}
	if bardNumberInit > 0 {
		validChars = append(validChars, charBard)
	}
	if explorerNumberInit > 0 {
		validChars = append(validChars, charExplorer)
	}
	if sailorNumberInit > 0 { // todo __sailor != noPos
		validChars = append(validChars, charSailor)
	}
	if merchantNumberInit > 0 {
		validChars = append(validChars, charMerchant)
	}

	if !slices.Contains(validChars, __curCharTypeNum) {
		panic(fmt.Sprint("请修改 initCharTypeNum"))
	}

	levelData := data{
		warrior:  warriorInitArr,
		thief:    thiefInitArr,
		wizard:   wizardInitArr,
		cleric:   priestInitArr,
		bard:     bardInitArr,
		druid:    druidInitArr,
		explorer: explorerInitArr,
		sailor:   sailorInitArr,
		merchant: merchantInitArr,

		stones:   stoneInitArr,
		crystals: crystalInitArr,
		grass:    grassInitArr,
		goblins:  goblinInitArr,
		dragons:  dragonInitArr,

		skippingStones:   skippingStoneInitArr,
		skippingCrystals: skippingCrystalInitArr,
		lilies:           lilyInitArr,

		mirrors:     mirrorInitArr,
		mirrorRefs:  mirrorRefInitArr,
		mirrorAuxes: mirrorAuxInitArr,

		beams: beamInitArr,

		curCharTypeNum: __curCharTypeNum,
	}

	type pair struct {
		data
		info string
	}
	from := map[data]pair{} // 同时充当 vis 的功能
	queue := []data{}
	defer func() { fmt.Printf("// 搜索了 %d 个状态\n", len(from)) }()

	add := func(last, d data, info string) {
		if !allowSkippingStoneGone && (pdContains(d.skippingStones[:], noPos) || pdContains(d.skippingCrystals[:], noPos)) {
			return
		}

		allMovableObjs, _, _ := d.getAllMovableObjPos(isBigMap, false)
		animeTypeMask := uint8(0)

		// 水漂石
		changed := false
		doSkippingStones := func(skippingStones []pointWithDir) {
			if isSkippingAndLilySlow {
				lilies := d.lilies[:]
				for i, p := range skippingStones {
					// 不动
					if p.dir == dirStop {
						continue
					}

					changed = true

					// 移动一步
					dir := directions4[p.dir]
					nxtP := p.point.add(dir)

					// 出界
					if !inBound(nxtP) {
						skippingStones[i] = noPosDir
						continue
					}

					// 撞上物品
					if !d.isValidPos(nxtP) || slices.Contains(allMovableObjs, nxtP) {
						// 落水
						skippingStones[i].z--
						skippingStones[i].dir = dirStop

						animeTypeMask |= 1 << animeFallIntoWater
						continue
					}

					// 下面不是水
					if !d.isFallIntoWater(nxtP) {
						// 停在这里
						skippingStones[i] = pointWithDir{nxtP, dirStop}
						if j := pdIndex(lilies, point{nxtP.x, nxtP.y, -1}); j >= 0 {
							// 停在了睡莲叶上，动量传给睡莲叶
							lilies[j].dir = p.dir
						}
						continue
					}

					// NOTE：为防止 bug，如果下下个位置是睡莲叶，先让睡莲叶停下来
					nxtP2 := nxtP.add(dir)
					if j := pdIndex(lilies, point{nxtP2.x, nxtP2.y, nxtP2.z - 1}); j >= 0 {
						lilies[j].dir = dirStop
					}

					// 两侧的睡莲叶
					if dir.x != 0 {
						// 左右
						p := point{nxtP.x, nxtP.y - 1, -1}
						if j := pdIndex(lilies, p); j >= 0 {
							lilies[j].dir = getDirIndexByDir(point{0, -1, 0})
						}
						p = point{nxtP.x, nxtP.y + 1, -1}
						if j := pdIndex(d.lilies[:], p); j >= 0 {
							lilies[j].dir = getDirIndexByDir(point{0, 1, 0})
						}
					} else {
						// 上下
						p := point{nxtP.x - 1, nxtP.y, -1}
						if j := pdIndex(lilies[:], p); j >= 0 {
							lilies[j].dir = getDirIndexByDir(point{-1, 0, 0})
						}
						p = point{nxtP.x + 1, nxtP.y, -1}
						if j := pdIndex(lilies[:], p); j >= 0 {
							lilies[j].dir = getDirIndexByDir(point{1, 0, 0})
						}
					}

					skippingStones[i].point = nxtP
				}
			} else {
				for i, p := range skippingStones {
					// todo 其实不需要判断？在 changePos 中已经处理好了
					if d.isFallIntoWater(p.point) {
						if !allowFallIntoWater {
							return
						}
						animeTypeMask |= 1 << animeFallIntoWater
						skippingStones[i].z = -1
					}
				}
			}
			if len(skippingStones) > 1 {
				slices.SortFunc(skippingStones, cmpPointWithDir)
			}
		}

		if len(d.skippingStones) > 0 {
			doSkippingStones(d.skippingStones[:])
		}
		if len(d.skippingCrystals) > 0 {
			doSkippingStones(d.skippingCrystals[:])
		}

		// 睡莲叶（相当于地形），最高优先级
		if len(d.lilies) > 0 {
			if isSkippingAndLilySlow {
				lilies := d.lilies[:]
				for i, p := range lilies {
					// 不动
					if p.dir == dirStop {
						continue
					}

					// 移动一步
					dir := directions4[p.dir]
					nxtP := p.point.add(dir)

					// 出界
					if !inBound(nxtP) {
						lilies[i] = noPosDir
						continue
					}

					// 不是纯水，停下来
					if !d.isFallIntoWater(point{nxtP.x, nxtP.y, 0}) {
						lilies[i].dir = dirStop
						continue
					}

					// 可以移动
					lilies[i].point = nxtP

					// 根据水流（如果有）修改移动方向
					ch := levelMap[0][nxtP.x][nxtP.y]
					switch ch {
					case '^', 'v', '<', '>':
						lilies[i].dir = flowDirMapping[ch]
					}

					// 睡莲叶承载的物品（如果有）也移动
					// dirIgnore 表示没找到的时候不报错
					if d.changePos(point{p.x, p.y, p.z + 1}, point{nxtP.x, nxtP.y, nxtP.z + 1},
						dirIgnore, nil) { // 上面的东西不会被阻挡，直接 nil
						changed = true
					}
				}
			}

			if len(d.lilies) > 1 {
				slices.SortFunc(d.lilies[:], cmpPointWithDir)
			}
		}

		if changed {
			allMovableObjs, _, _ = d.getAllMovableObjPos(isBigMap, false)
		}

		// 确定门的开闭，方便后面判断下落
		for i, sw := range switches {
			opened := !doorOpenedInit[i]
			// 如果有一个开关没有被压住，那么 opened 为初始状态
			for _, p := range sw {
				pushed := slices.Contains(allMovableObjs, p) ||
					len(d.grass) > 0 && slices.Contains(d.grass[:], p) // 草也可以压住开关
				if !pushed {
					opened = !opened // 变回 doorOpenedInit[i]
					break
				}
			}
			d.doorOpened[i] = opened

			// 水晶被门压碎（水晶在门中，但门没有打开）
			if !opened {
				for j, p := range d.stones {
					if pdContains(doors[i], p) {
						if !canDestroyObj {
							return
						}
						d.stones[j] = noPos
					}
				}
				for j, p := range d.crystals {
					if pdContains(doors[i], p) {
						if !canDestroyObj {
							return
						}
						d.crystals[j] = noPos
					}
				}
			}
		}

		// 被喷火龙攻击到的位置
		var burnedPos []point
		if len(d.dragons) > 0 {
			for _, dra := range d.dragons {
				if dra.z < 0 || dra.dir&dirIsCrystal > 0 { // 是石头
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

					if len(d.mirrors) > 0 || len(d.mirrorRefs) > 0 || len(d.mirrorAuxes) > 0 {
						// 这里的逻辑和法师是一样的
						mir := noPosDir
						if i := pdIndex(d.mirrors[:], cur); i >= 0 && d.mirrors[i].canReflect(dir) {
							mir = d.mirrors[i]
						} else if i := pdIndex(d.mirrorRefs[:], cur); i >= 0 && d.mirrorRefs[i].canReflect(dir) {
							mir = d.mirrorRefs[i]
						} else if i := pdIndex(d.mirrorAuxes[:], cur); i >= 0 && d.mirrorAuxes[i].canReflect(dir) {
							mir = d.mirrorAuxes[i]
						}

						// 面对的是镜子的正面
						if mir.point != noPos {
							dir2 := mir.reflectToAnotherDir(dir)
							// 沿着光路搜索，找第一个可交换对象
							refP := d.reflectTo(mir, dir2, math.MaxInt, allMovableObjs)
							if refP != noPos {
								burnedPos = append(burnedPos, refP)
							}
							break
						}
					}

					if slices.Contains(allMovableObjs, cur) {
						burnedPos = append(burnedPos, cur)
						break
					}
				}
			}
		}

		// 对象下落到 z >= 0
		if mapSizeH > 1 {
			// todo 多个 Movable Obj 会不会互相影响？
			for _, p := range allMovableObjs {
				if p.z <= 0 {
					continue
				}

				// 如果 p 是电梯
				if levelMap[p.z][p.x][p.y] == 'e' {
					continue
				}

				// 如果 p 是牧师或其邻居，且正被攻击，那么 p 不会下落
				if d.isProtected(p) && d.isAttacked(p, burnedPos) {
					continue
				}

				oldP := p
				for p.z > 0 {
					p.z--

					// 如果下面是镜子，则踩碎镜子，继续下落
					if i := pdIndex(d.mirrors[:], p); i >= 0 {
						d.mirrors[i] = noPosDir
						continue
					}
					if i := pdIndex(d.mirrorRefs[:], p); i >= 0 {
						d.mirrorRefs[i] = noPosDir
						continue
					}
					if i := pdIndex(d.mirrorAuxes[:], p); i >= 0 {
						d.mirrorAuxes[i] = noPosDir
						continue
					}

					if !d.isValidPos(p) || slices.Contains(allMovableObjs, p) {
						p.z++
						break // 下面不是空
					}
				}

				if oldP.z != p.z {
					if !allowFallIntoGround {
						return
					}
					info += strings.Repeat("W", int(oldP.z-p.z)) // todo
					d.changePos(oldP, p, math.MaxUint8, allMovableObjs)
				}
			}
		}

		// 先判断是否有角色死亡
		for _, char := range d.getAllCharPos(false) {
			if d.getDieType(char, burnedPos, true) != dieTypeNo {
				return
			}
		}

		// 一开始，以及切换角色，都不结算怪物之间的攻击
		isSwitching := info[0] == 'c' || '1' <= info[0] && info[0] <= '9'
		if !isSwitching && !d.monsterDoorOpened {
			// 哥布林
			goblins := d.goblins
			if len(d.goblins) > 0 {
				for i, p := range d.goblins {
					if p.z < 0 {
						continue
					}
					if p.dir&dirIsCrystal > 0 { // 是石头
						// 落水
						if d.isFallIntoWater(p.point) {
							if !allowFallIntoWater {
								return
							}
							animeTypeMask |= 1 << animeFallIntoWater
							goblins[i].z = -1
						}
						continue
					}
					// todo 变成水晶的哥布林 + 水晶哥布林落水
					if tp := d.getDieType(p.point, burnedPos, false); tp != dieTypeNo {
						if !canDestroyObj {
							return
						}
						if tp == dieTypeAttacked {
							animeTypeMask |= 1 << animeKill
						} else if tp == dieTypeDrown {
							animeTypeMask |= 1 << animeFallIntoWater
						}
						goblins[i] = noPosDir
					}
				}
				if len(goblins) > 1 {
					slices.SortFunc(goblins[:], cmpPointWithDir)
				}
			}

			// 喷火龙
			dragons := d.dragons
			if len(d.dragons) > 0 {
				for i, p := range d.dragons {
					if p.z < 0 {
						continue
					}
					if p.dir&dirIsCrystal > 0 { // 是石头
						// 落水
						if d.isFallIntoWater(p.point) {
							if !allowFallIntoWater {
								return
							}
							animeTypeMask |= 1 << animeFallIntoWater
							dragons[i].z = -1
						}
						continue
					}
					if tp := d.getDieType(p.point, burnedPos, false); tp != dieTypeNo {
						if !canDestroyObj {
							return
						}
						if tp == dieTypeAttacked {
							animeTypeMask |= 1 << animeKill
						} else if tp == dieTypeDrown {
							animeTypeMask |= 1 << animeFallIntoWater
						}
						dragons[i] = noPosDir
					}
				}
				if len(dragons) > 1 {
					slices.SortFunc(dragons[:], cmpPointWithDir)
				}
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
					animeTypeMask |= 1 << animeFallIntoWater
					mir[i].z = -1
				}
			}
			if len(d.mirrors) > 1 {
				slices.SortFunc(mir, cmpPointWithDir)
			}
		}

		// 可以被反射的镜子
		if len(d.mirrorRefs) > 0 {
			mir := d.mirrorRefs[:]
			for i, p := range mir {
				if d.isFallIntoWater(p.point) {
					if !allowFallIntoWater {
						return
					}
					animeTypeMask |= 1 << animeFallIntoWater
					mir[i].z = -1
				}
			}
			if len(d.mirrorRefs) > 1 {
				slices.SortFunc(mir, cmpPointWithDir)
			}
		}

		// 辅助镜子
		if len(d.mirrorAuxes) > 0 {
			mir := d.mirrorAuxes[:]
			for i, p := range mir {
				if d.isFallIntoWater(p.point) {
					if !allowFallIntoWater {
						return
					}
					animeTypeMask |= 1 << animeFallIntoWater
					mir[i].z = -1
				}
			}
			if len(d.mirrorAuxes) > 1 {
				slices.SortFunc(mir, cmpPointWithDir)
			}
		}

		// 石头
		if len(d.stones) > 0 {
			sto := d.stones[:]
			for i, p := range sto {
				if d.isFallIntoWater(p) {
					if !allowFallIntoWater {
						return
					}
					animeTypeMask |= 1 << animeFallIntoWater
					sto[i].z = -1
				}
			}
			if len(d.stones) > 1 {
				slices.SortFunc(sto, cmpPoint)
			}
		}

		// 水晶
		if len(d.crystals) > 0 {
			cry := d.crystals[:]
			for i, p := range cry {
				if d.isFallIntoWater(p) {
					if !allowFallIntoWater {
						return
					}
					animeTypeMask |= 1 << animeFallIntoWater
					cry[i].z = -1
				}
			}
			if len(d.crystals) > 1 {
				slices.SortFunc(cry, cmpPoint)
			}
		}

		// 草
		if len(d.grass) > 1 {
			slices.SortFunc(d.grass[:], cmpPoint)
		}

		// 光束
		if len(d.beams) > 1 {
			slices.SortFunc(d.beams[:], cmpPointWithDir)
		}

		// 人
		if len(d.warrior) > 1 {
			slices.SortFunc(d.warrior[:], cmpPoint)
		}
		if len(d.thief) > 1 {
			slices.SortFunc(d.thief[:], cmpPoint)
		}
		if len(d.wizard) > 1 {
			slices.SortFunc(d.wizard[:], cmpPoint)
		}
		if len(d.cleric) > 1 {
			slices.SortFunc(d.cleric[:], cmpPoint)
		}
		if len(d.druid) > 1 {
			slices.SortFunc(d.druid[:], cmpPoint)
		}
		if len(d.bard) > 1 {
			slices.SortFunc(d.bard[:], cmpPoint)
		}
		if len(d.explorer) > 1 {
			slices.SortFunc(d.explorer[:], cmpPoint)
		}
		if len(d.sailor) > 1 {
			slices.SortFunc(d.sailor[:], cmpPoint)
		}
		if len(d.merchant) > 1 {
			slices.SortFunc(d.merchant[:], cmpPoint)
		}

		if _, ok := from[d]; !ok {
			if animeTypeMask>>animeKill&1 > 0 {
				info += "K"
			}
			if animeTypeMask>>animeFallIntoWater&1 > 0 {
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

		allMovableObjs, allChars, allNonChars := d.getAllMovableObjPos(isBigMap, false)

		var pass bool
		if !targetIsClearAllMonsters {
			// 标准版：所有人都到达终点
			if isBigMap {
				p := d.getCurCharPos()
				pass = slices.Equal([]point{p}, finals)
			} else if len(d.isCrystalMask) == 0 || d.isCrystalMask[:][0] == 0 { // 不能有人是水晶
				if len(allChars) > 1 {
					slices.SortFunc(allChars, cmpPoint)
				}
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
				if pre.data == (data{}) { // 初始状态
					break
				}

				infoStr := pre.info
				if infoStr != "IGNORE" {
					diffNeg1 := 0
					for _, p := range pre.skippingStones {
						if p.z == -1 {
							diffNeg1++
						}
					}
					for _, p := range d.skippingStones {
						if p.z == -1 {
							diffNeg1--
						}
					}

					diffNeg2 := 0
					for _, p := range pre.skippingCrystals {
						if p.z == -1 {
							diffNeg2++
						}
					}
					for _, p := range d.skippingCrystals {
						if p.z == -1 {
							diffNeg2--
						}
					}
					if (diffNeg1 != 0 || diffNeg2 != 0) && !strings.Contains(infoStr, "W") {
						infoStr += "W" // 水漂石落水
					}

					if !isSkippingAndLilySlow && pre.lilies != d.lilies {
						// 大致估算睡莲叶的移动步数（前后排序了，不准）
						maxStep := 0
						for i, p := range pre.lilies {
							maxStep = max(maxStep, int(abs(p.x-d.lilies[i].x))+int(abs(p.y-d.lilies[i].y)))
						}
						if maxStep == 0 {
							panic("代码有误！maxStep = 0")
						}
						infoStr += strings.Repeat(".", maxStep)
					}

					path = append(path, infoStr)

					// 最上面输出的是最后的
					//fmt.Println(infoStr, pre.sailor[0], pre.skippingStones, pre.lilies) // DEBUG
				}
				d = pre.data
			}
			slices.Reverse(path)
			return path
		}

		// todo 如果角色的头上有物品，物品会跟着移动（注意镜子的方向会变）    堆叠上限是多少？？
		// todo 即使人没有移动，切换方向也会改变头上物品（镜子、激光等）的方向
		// todo 多控时，如果下一个位置是没有石头的水，则一个角色无法移动（已在商人中实现）

		// 先考虑按 x 镜子反射对象，这样后面移动更流畅
		// 只要有一个镜子反射失败（红光），就直接 return
		doMirrors := func() {
			newData := d
			swapped := uint(0)
		nextMirror:
			for _, mirror := range append(d.mirrors[:], d.mirrorRefs[:]...) {
				if mirror == noPosDir {
					continue
				}

				// 找两个方向最近的可反射的对象
				cur0 := mirror.point
				cur1 := mirror.point
				dir0 := directions6[mirror.dir&0xf]
				dir1 := directions6[mirror.dir>>4]
				foundMirror := uint8(0)
				for step := 1; ; step++ {
					justFound := uint8(0) // 是否找到了非镜子对象
					// 检查方向 0
					if foundMirror&1 == 0 {
						cur0 = cur0.add(dir0)
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
						cur1 = cur1.add(dir1)
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
						continue // 都是空，继续找
					}

					oldP := cur0
					dir := dir1
					if justFound == 2 {
						oldP = cur1
						dir = dir0 // 往另一个方向反射
					}

					// 无法反射的石头，视作墙壁
					if slices.Contains(d.stones[:], oldP) || pdContains(d.skippingStones[:], oldP) {
						continue nextMirror
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
						} else if slices.Contains(d.crystals[:], oldP) {
							//newData.crystals[0] = newP // todo
						} else {
							// todo 其他对象的分身
						}
					} else {
						swapped |= 1 << itemIdx
						if i := pdIndex(newData.dragons[:], oldP); i >= 0 { // 喷火龙
							// todo 多次反射
							newDir := mirror.reflectDragon(newData.dragons[i].dir)
							newData.dragons[i] = pointWithDir{newP, newDir}
						} else if i := pdIndex(newData.skippingCrystals[:], oldP); i >= 0 { // 水漂石
							// todo 多次反射
							newDir := newData.skippingCrystals[i].dir
							if newDir != dirStop {
								newDir = mirror.reflectDragon(newDir)
							}
							newData.skippingCrystals[i] = pointWithDir{newP, newDir}
						} else if i := pdIndex(newData.mirrorRefs[:], oldP); i >= 0 { // 可被反射的镜子
							// 如果是 oldP 是可被反射的镜子，则与 mir 垂直的镜子会前后翻转
							newDir := mirror.reflectMirrorRef(newData.mirrorRefs[i].dir)
							newData.mirrorRefs[i] = pointWithDir{newP, newDir}
						} else {
							newData.changePos(oldP, newP, math.MaxUint8, allMovableObjs)
						}
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

		// 只有当前角色会坐电梯？
		// todo 多控？
		doElevator := func(p point) {
			if mapSizeH == 1 {
				return
			}
			if p.z == 0 && levelMap[p.z][p.x][p.y] == 'e' ||
				p.z == mapSizeH-1 && levelMap[p.z][p.x][p.y] == 'e' {
				newData := d
				newData.changePos(p, point{p.x, p.y, p.z ^ (mapSizeH - 1)}, math.MaxUint8, allMovableObjs)
				add(d, newData, "v")
			}
		}

		// 移动当前角色
		switch d.curCharTypeNum {
		case charWarrior:
			// 普通移动一步
			p0 := d.warrior[:][0] // todo 暂时支持一个人
			doElevator(p0)

			withinBeams, beamNf := d.withinBeams(p0, allNonChars)

			// 在墙里面，但不能穿透
			if withinBeams>>beamThrough&1 == 0 && inBound(p0) && levelMap[p0.z][p0.x][p0.y] == '#' {
				goto afterSwitch
			}

			for dIdx, dir := range directions4 {
				newP := p0.add(dir)
				if withinBeams>>beamEndpoint&1 > 0 && (dir == beamNf.endPointDir || dir == beamNf.endPointDir.rev()) {
					// 如果在 endpoint 光中，优先级更高，只能往该方向走到终点
					// 如果面朝发射器移动，移动到发射器前一格子
					// 如果背朝发射器移动，移动到末端格子
					// 如果移动到的格子不合法，则不能移动，否则移动过去
					if dir != beamNf.endPointDir { // 方向相反
						newP = beamNf.endPoints[0]
					} else {
						newP = beamNf.endPoints[1]
					}
					if newP == p0 || !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) {
						continue // 原地不动 or 出界或者有障碍物
					}
					newData := d
					newData.warrior[:][0] = newP
					newData.bigMapForceSwapChar(p0, newP)
					add(d, newData, dir4String[dIdx]+"L")
					continue
				}

				// 该方向有多少个连续的对象
				cnt := 0
				cur := p0.add(dir)
				for slices.Contains(allMovableObjs, cur) && !hasFence(cur, dir.rev()) {
					cnt++
					cur = cur.add(dir)
				}
				// 前面是否有空地
				if !(withinBeams>>beamThrough&1 > 0 && inBound(cur) && levelMap[cur.z][cur.x][cur.y] == '#') && // obj 可以到墙中
					(!d.isValidPos(cur) || hasFence(cur, dir.rev())) {
					continue // 枚举另一个方向
				}

				newData := d
				for range cnt {
					// 倒着回来
					nxt := cur.sub(dir) // 这是个物品
					newData.changePos(nxt, cur, math.MaxUint8, allMovableObjs)

					// todo 多层
					oldTop := point{nxt.x, nxt.y, nxt.z + 1}
					// todo 喷火龙 / 镜子
					if slices.Contains(allMovableObjs, oldTop) {
						newTop := cur
						newTop.z++
						newData.changePos(oldTop, newTop, uint8(dIdx), allMovableObjs)
					}
					cur = nxt
				}

				newP = p0.add(dir)

				if mapSizeH > 1 {
					oldTop := point{p0.x, p0.y, p0.z + 1}
					// 如果原位置头上有喷火龙或者镜子，修改其位置和朝向
					if i := pdIndex(newData.dragons[:], oldTop); i >= 0 {
						newTop := newP
						newTop.z++
						if !d.isValidPos(newTop) || slices.Contains(allMovableObjs, newTop) {
							continue // todo 暂时禁止喷火龙落地 
						}
						// todo 如果喷火龙和人的方向不同呢？
						newData.dragons[i] = pointWithDir{newTop, uint8(dIdx)}
					} else if i := pdIndex(newData.mirrors[:], oldTop); i >= 0 {
						newTop := newP
						newTop.z++
						if !d.isValidPos(newTop) || slices.Contains(allMovableObjs, newTop) {
							continue
						}
						newData.mirrors[i] = pointWithDir{newTop, defaultMirrorDirs[dIdx]}
					} else if slices.Contains(allMovableObjs, oldTop) {
						newTop := newP
						newTop.z++
						newData.changePos(oldTop, newTop, uint8(dIdx), allMovableObjs)
					}
					// todo 镜子
				}

				newData.warrior[:][0] = newP // todo 暂时支持一个人
				newData.bigMapForceSwapChar(p0, newP)
				add(d, newData, dir4String[dIdx])
			}
		case charThief:
			// 普通移动一步
			p0 := d.thief[:][0] // todo
			doElevator(p0)

			withinBeams, beamNf := d.withinBeams(p0, allNonChars)

			// 在墙里面，但不能穿透
			if withinBeams>>beamThrough&1 == 0 && inBound(p0) && levelMap[p0.z][p0.x][p0.y] == '#' {
				goto afterSwitch
			}

			for dIdx, dir := range directions4 {
				newP := p0.add(dir)

				if withinBeams>>beamEndpoint&1 > 0 && (dir == beamNf.endPointDir || dir == beamNf.endPointDir.rev()) {
					// 如果在 endpoint 光中，优先级更高，只能往该方向走到终点
					// 如果面朝发射器移动，移动到发射器前一格子
					// 如果背朝发射器移动，移动到末端格子
					// 如果移动到的格子不合法，则不能移动，否则移动过去
					if dir != beamNf.endPointDir { // 方向相反
						newP = beamNf.endPoints[0]
					} else {
						newP = beamNf.endPoints[1]
					}
					if newP == p0 || !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) {
						continue // 原地不动 or 出界或者有障碍物
					}
					newData := d
					newData.thief[:][0] = newP
					newData.bigMapForceSwapChar(p0, newP)
					add(d, newData, dir4String[dIdx]+"L")
					continue
				}

				// 前面是否有空地
				if !(withinBeams>>beamThrough&1 > 0 && inBound(newP) && levelMap[newP.z][newP.x][newP.y] == '#') &&
					(!d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) || hasFence(p0, dir)) {
					continue // 枚举另一个方向
				}

				newData := d
				back := p0.sub(dir)
				if slices.Contains(allMovableObjs, back) && !hasFence(back, dir) {
					// 拉人/物 -> 当前位置
					newData.changePos(back, p0, math.MaxUint8, allMovableObjs)
				}
				newData.thief[:][0] = newP
				newData.bigMapForceSwapChar(p0, newP)
				add(d, newData, dir4String[dIdx])
			}
		case charWizard:
			p0 := d.wizard[:][0]
			doElevator(p0)

			withinBeams, beamNf := d.withinBeams(p0, allNonChars)

			// 在墙里面，但不能穿透
			if withinBeams>>beamThrough&1 == 0 && inBound(p0) && levelMap[p0.z][p0.x][p0.y] == '#' {
				goto afterSwitch
			}

		nextDir:
			for dIdx, dir := range directions4 {
				var newP point
				if withinBeams>>beamDouble&1 > 0 {
					// todo 先看走一步是不是物品，黄光的规则是这样的吗？
					//newP = point{p0.x + dir.x, p0.y + dir.y, p0.z + dir.z}
					//if slices.Contains(allMovableObjs, newP) {
					//	// 和对象交换位置
					//	newData := d
					//	newData.changePos(newP, p0, math.MaxUint8) // newP 换到 p0
					//	newData.wizard[:][0] = newP                // 法师换到 newP
					//	add(d, newData, dir4String[dIdx]+"P")      // swap
					//	continue
					//}

					// 如果在 double 光中，优先级更高，只能往该方向走两步
					// todo 多个 double 光的情况，要叠加
					// todo 绿光
					const multi = 2
					newP = point{p0.x + dir.x*multi, p0.y + dir.y*multi, p0.z + dir.z*multi}
					if !d.isValidPos(newP) {
						continue // 出界或者有障碍物（墙、草）
					}
					if slices.Contains(allMovableObjs, newP) {
						// 和对象交换位置
						newData := d
						newData.changePos(newP, p0, math.MaxUint8, allMovableObjs) // newP 换到 p0
						newData.wizard[:][0] = newP                                // 法师换到 newP
						add(d, newData, dir4String[dIdx]+"P")                      // swap
						continue
					}
					// 空地，直接走过去
					newData := d
					newData.wizard[:][0] = newP
					newData.bigMapForceSwapChar(p0, newP)
					add(d, newData, dir4String[dIdx]+"L")
				} else if withinBeams>>beamEndpoint&1 > 0 && (dir == beamNf.endPointDir || dir == beamNf.endPointDir.rev()) {
					// 如果在 endpoint 光中，优先级更高，只能往该方向走到终点
					// 如果面朝发射器移动，移动到发射器前一格子
					// 如果背朝发射器移动，移动到末端格子
					// 如果移动到的格子不合法，则不能移动，否则移动过去 / 交换物品
					if dir != beamNf.endPointDir { // 方向相反
						newP = beamNf.endPoints[0]
					} else {
						newP = beamNf.endPoints[1]
					}
					if newP == p0 || !d.isValidPos(newP) {
						continue // 原地不动 or 出界或者有障碍物（墙、草）
					}
					if slices.Contains(allMovableObjs, newP) {
						// 和对象交换位置
						newData := d
						newData.changePos(newP, p0, math.MaxUint8, allMovableObjs) // newP 换到 p0
						newData.wizard[:][0] = newP                                // 法师换到 newP
						add(d, newData, dir4String[dIdx]+"P")                      // swap
						continue
					}
					// 空地，直接走过去
					newData := d
					newData.wizard[:][0] = newP
					newData.bigMapForceSwapChar(p0, newP)
					add(d, newData, dir4String[dIdx]+"L")
					continue
				} else {
					// dir 方向是否有可交换对象
					cur := p0
					for {
						cur = cur.add(dir)
						newP := cur
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
						// 注：这里可能自己和自己交换
						newData := d
						newData.changePos(newP, p0, math.MaxUint8, allMovableObjs) // newP 换到 p0
						newData.wizard[:][0] = newP                                // 法师换到 newP
						add(d, newData, dir4String[dIdx]+"P")                      // swap
						continue nextDir
					}

					// 没有可交换对象，那就普通移动
					newP = p0.add(dir)
					if !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) || hasFence(p0, dir) {
						continue // 枚举另一个方向
					}
				}

				newData := d
				newData.wizard[:][0] = newP
				newData.bigMapForceSwapChar(p0, newP)
				add(d, newData, dir4String[dIdx]) // move
			}
		case charCleric:
			// 普通移动一步
			p0 := d.cleric[:][0]
			doElevator(p0)

			withinBeams, _ := d.withinBeams(p0, allNonChars)

			for dIdx, dir := range directions4 {
				newP := p0.add(dir)
				if !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) || hasFence(p0, dir) {
					continue // 枚举另一个方向
				}
				newData := d
				if allowAllPushItem && withinBeams>>beamStrong&1 > 0 {
					if i := slices.Index(allMovableObjs, newP); i >= 0 {
						// 可以推物品
						nxt2 := newP.add(dir)
						if !d.isValidPos(nxt2) || slices.Contains(allMovableObjs, nxt2) {
							continue // 枚举另一个方向
						}
						newData.changePos(newP, nxt2, math.MaxUint8, allMovableObjs)
					}
				}
				newData.cleric[:][0] = newP
				newData.bigMapForceSwapChar(p0, newP)
				add(d, newData, dir4String[dIdx])
			}
		case charBard:
			p0 := d.bard[:][0]
			doElevator(p0)

			// todo 栅栏

			allLife, nonLife := d.getAllLife(isBigMap)
			items := allLife[:0]
			for _, p := range allLife {
				if chebyshevDis(p, p0) <= 2 {
					items = append(items, p)
				}
			}
			if isBigMap {
				items = append(items, p0)
			}

			// 普通移动一步
			// 切比雪夫距离 <= 2 的物品（包括自己）都移动一步
			for dIdx, dir := range directions4 {
				x, y, z := p0.x+dir.x, p0.y+dir.y, p0.z+dir.z
				if !d.isValidPos(point{x, y, z}) {
					continue
				}
				if len(items) > 1 {
					slices.SortFunc(items, func(a, b point) int {
						if dir.x != 0 {
							return int(b.x*dir.x - a.x*dir.x)
						}
						return int(b.y*dir.y - a.y*dir.y)
					})
				}

				newData := d
				unmovedItems := []point{}
				movedItems := []point{}
				for _, oldP := range items {
					// item 往前移动一格
					newP := oldP.add(dir)
					if !d.isValidPos(newP) || slices.Contains(nonLife, newP) { // 无法移动
						unmovedItems = append(unmovedItems, oldP)
						continue
					}
					// 大地图物品不能出界 
					// todo 目前只实现了诗人的逻辑
					if isBigMap && oldP != p0 && strings.ContainsRune("ATWCDB789", rune(levelMap[0][newP.x][newP.y])) {
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
					newData.changePos(oldP, newP, math.MaxUint8, allMovableObjs)
				}

				if !slices.Contains(unmovedItems, p0) {
					if newData.bard[:][0] != (point{x, y, z}) {
						panic("诗人移动错误，代码有误")
					}

					// 特性：如果诗人脚下是物品，且该物品移动了，那么诗人可以再走一格
					// todo 对于物品叠物品的情况，也是同样的规则？
					if slices.Contains(movedItems, point{p0.x, p0.y, p0.z - 1}) {
						nxtP := point{x + dir.x, y + dir.y, z + dir.z}
						if d.isValidPos(nxtP) && !slices.Contains(unmovedItems, nxtP) {
							newData.bard[:][0] = nxtP
						}
						// todo （待确认）如果 z-2 也移动了，那么再再走一格
					}

					newData.bigMapForceSwapChar(p0, newData.bard[:][0])
					add(d, newData, dir4String[dIdx])
				}
			}
		case charDruid:
			p0 := d.druid[:][0]
			doElevator(p0)

			for dIdx, dir := range directions4 {
				newP := p0.add(dir)
				// 草变水晶
				if len(d.grass) > 0 {
					if i := slices.Index(d.grass[:], newP); i >= 0 {
						newData := d
						newData.crystals[:][0] = newData.grass[i] // 加个切片避免报错
						newData.grass[i] = noPos
						add(d, newData, dir4String[dIdx]+"C") // trans
						continue
					}
				}

				// 水晶变草
				if len(d.crystals) > 0 && druidCrystalToGrass {
					if i := slices.Index(d.crystals[:], newP); i >= 0 {
						newData := d
						newData.grass[:][0] = newData.crystals[i]
						newData.crystals[i] = noPos
						add(d, newData, dir4String[dIdx]+"C") // trans
						continue
					}
				}

				// 哥布林 <-> 水晶
				if len(d.goblins) > 0 && priestNumberInit > 0 && d.cleric[:][0] != noPos {
					if i := pdIndex(d.goblins[:], newP); i >= 0 {
						newData := d
						newData.goblins[i].dir ^= dirIsCrystal
						add(d, newData, dir4String[dIdx]+"C") // trans
						continue
					}
				}

				// 喷火龙 <-> 水晶
				if len(d.dragons) > 0 {
					if i := pdIndex(d.dragons[:], newP); i >= 0 {
						newData := d
						newData.dragons[i].dir ^= dirIsCrystal
						add(d, newData, dir4String[dIdx]+"C") // trans
						continue
					}
				}

				// 人 <-> 水晶
				// 目前只支持单人
				if druidTransMan {
					if tp := d.getCharType(newP); tp > 0 {
						newData := d
						newData.isCrystalMask[:][0] ^= 1 << (tp - 1)
						add(d, newData, dir4String[dIdx]+"C") // trans
						continue
					}
				}

				// 普通移动一步
				if !d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) || hasFence(p0, dir) {
					continue
				}
				newData := d
				newData.druid[:][0] = newP
				newData.bigMapForceSwapChar(p0, newP)
				add(d, newData, dir4String[dIdx]) // move
			}
		case charExplorer:
			// 普通移动一步
			p0 := d.explorer[:][0]
			doElevator(p0)

			for dIdx, dir := range directions4 {
				newP := p0.add(dir)
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
						nxt2 := newP.add(dir)
						if !d.isValidPos(nxt2) || slices.Contains(allMovableObjs, nxt2) {
							continue // 枚举另一个方向
						}
						newData.changePos(newP, nxt2, math.MaxUint8, allMovableObjs)
					}
				} else if slices.Contains(allMovableObjs, newP) || hasFence(p0, dir) {
					continue
				}

				if mapSizeH > 1 {
					oldTop := point{p0.x, p0.y, p0.z + 1}
					// 如果原位置头上有喷火龙或者镜子，修改其位置和朝向
					if i := pdIndex(newData.dragons[:], oldTop); i >= 0 {
						newTop := newP
						newTop.z++
						if !d.isValidPos(newTop) || slices.Contains(allMovableObjs, newTop) {
							continue // todo 暂时禁止喷火龙落地 
						}
						// todo 如果喷火龙和人的方向不同呢？
						newData.dragons[i] = pointWithDir{newTop, uint8(dIdx)}
					} else if slices.Contains(allMovableObjs, oldTop) {
						newTop := newP
						newTop.z++
						newData.changePos(oldTop, newTop, uint8(dIdx), allMovableObjs)
					}
					// todo 镜子
				}

				newData.explorer[:][0] = newP
				newData.bigMapForceSwapChar(p0, newP)
				add(d, newData, dir4String[dIdx])
			}
		case charSailor:
			// 普通移动一步
			p0 := d.sailor[:][0]
			doElevator(p0)

			for dIdx, dir := range directions4 {
				newP := p0.add(dir)
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
				if allowAllPushItem {
					if i := slices.Index(allMovableObjs, newP); i >= 0 {
						// 推物品
						nxt2 := newP.add(dir)
						if !d.isValidPos(nxt2) || slices.Contains(allMovableObjs, nxt2) {
							continue // 枚举另一个方向
						}
						newData.changePos(newP, nxt2, math.MaxUint8, allMovableObjs)
					}
				} else if slices.Contains(allMovableObjs, newP) || hasFence(p0, dir) {
					continue
				}

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
						newData.changePos(oldTop, newTop, uint8(dIdx), allMovableObjs)
					}
					// todo 镜子
				}

				newData.sailor[:][0] = newP
				newData.bigMapForceSwapChar(p0, newP)
				add(d, newData, dir4String[dIdx])
			}
		case charMerchant:
			doElevator(d.merchant[:][0]) // todo

			// todo 栅栏

			// 多控
			// 普通移动一步
			for dIdx, dir := range directions4 {
				newData := d
				oldMerchant := newData.merchant
				man := newData.merchant[:]
				if len(newData.merchant) > 1 {
					slices.SortFunc(man, func(a, b point) int {
						if dir.x != 0 {
							return int(b.x*dir.x - a.x*dir.x)
						}
						return int(b.y*dir.y - a.y*dir.y)
					})
				}

				unmovedMan := []point{}
				moved := false
				for manIdx, p0 := range man {
					if p0 == noPos {
						continue
					}
					nxt := p0.add(dir)
					// 无法移动（注意岸边也是无法移动的）
					if !d.isValidPos(nxt) || d.isFallIntoWater(nxt) || slices.Contains(unmovedMan, nxt) {
						unmovedMan = append(unmovedMan, p0)
						continue
					}
					// 如果前面是物品，则推动（能移动的人已经移动了）
					if !slices.Contains(oldMerchant[:], nxt) && slices.Contains(allMovableObjs, nxt) {
						nxt2 := nxt.add(dir)
						// 无法推动前面的物品
						if !d.isValidPos(nxt2) ||
							!slices.Contains(oldMerchant[:], nxt2) && slices.Contains(allMovableObjs, nxt2) ||
							slices.Contains(unmovedMan, nxt2) {
							unmovedMan = append(unmovedMan, p0)
							continue
						}
						newData.changePos(nxt, nxt2, math.MaxUint8, allMovableObjs)
					}
					moved = true
					man[manIdx] = nxt // 移走！
				}
				if !moved { // 没人动
					continue
				}
				add(d, newData, dir4String[dIdx])
			}
		case charDefault:
			panic("代码有误，当前角色不能为 charDefault")
		default:
		}

	afterSwitch:
		// 原地不动
		if isSkippingAndLilySlow {
			add(d, d, ".")
		}

		// 换成其他人
		if !isBigMap {
			for _, char := range validChars {
				if char == d.curCharTypeNum {
					continue
				}
				// 不能是水晶
				if len(d.isCrystalMask) > 0 && d.isCrystalMask[:][0]>>(char-1)&1 > 0 {
					continue
				}
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
	charDruid
	charBard
	charExplorer
	charSailor   // 同大地图角色
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

const (
	beamDefault  = iota
	beamOpen     // 红 1
	beamDouble   // 橙 2
	beamSmash    // 黄 3
	beamThrough  // 绿 4
	beamStrong   // 青 5
	beamUnknown  // 蓝 6 todo
	beamEndpoint // 紫 7
	beamActive   // 白 8
	beamModify   // 彩 9
)

const dirIgnore = 1 << 7
const dirIsCrystal = 1 << 6
const dirStop = 1 << 5
