package oss

import (
	"fmt"
	"math"
	"math/bits"
	"runtime"
	"slices"
	"strings"
	"time"
	"unsafe"
)

/*
2026.6.13

todo
 荷叶旋转时，其上方的物品也要旋转
    单独整理一个旋转上方物品的函数
 光束 + 镜子反射
 多次反射后，喷火龙的朝向
 诗人 叠加物品机制 https://www.bilibili.com/video/BV1yyT16sEX1/
 7 放置宝石
 植物的根往下长
 https://www.youtube.com/watch?v=ntlY5gDi17E&t=124s
 Order of the Sinking Star interview with Jonathan Blow
 https://www.youtube.com/watch?v=he78sfGu_ww
 https://youtu.be/he78sfGu_ww?si=j3hhlrC-qOwFe0DQ&t=38

todo 重构
 !d.isValidPos(cur) || slices.Contains(_allMovableObjs, cur)
 修改 changePos 的代码，添加一个参数 alsoMoveTop bool，使得当物品移动时，物品上方的物品（如果有）也跟着移动

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
type gemArrType [gemNumberInit]point
type beamArrType [len(beamDirInit)]pointWithDir
type mirrorArrType [len(mirrorDirInit) / 2]pointWithDir
type mirrorRefArrType [len(mirrorRefDirInit) / 2]pointWithDir
type mirrorAuxArrType [len(mirrorAuxDirInit) / 2]pointWithDir

type data struct {
	warrior  warriorArrType  // A 推多个对象
	thief    thiefArrType    // T 拉一个对象
	wizard   wizardArrType   // W 交换对象
	cleric   priestArrType   // C 自己以及上下左右无敌
	bard     bardArrType     // B 同时移动切比雪夫距离 <= 2 的对象
	druid    druidArrType    // D 把对象变成石头
	explorer explorerArrType // 7 普通角色，无法推对象   todo 用 dir（高 4 位）记录宝石个数
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
	mirrorRefs  mirrorRefArrType // R 可以被反射的镜子（可被诗人魅惑）
	mirrorAuxes mirrorAuxArrType // m 关卡名中称其为 mundane

	// 宝石
	gems gemArrType // o

	// 光束
	// 高 4 位是类型，低 4 位是方向
	beams beamArrType // b

	//redDoors [0] // r

	// 哪些角色变成了水晶（从 0 开始）
	crystalHeroMask [min(druidNumberInit, 1)]uint8 // uint8 不支持 merchant

	// todo 用于快速判断牧师是否悬浮
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

func initMap() { // 初始化变量、检查地图是否与 const 匹配
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

	for i, p := range goblinPosInit {
		goblinPosInit[i].point = changeNegPoint(p.point)
	}
	for i, p := range dragonPosInit {
		dragonPosInit[i].point = changeNegPoint(p.point)
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

	var doorMask, switchMask, finalNum, socketNum int
	var warriorNum, thiefNum, wizardNum, priestNum, druidNum, bardNum, explorerNum, sailorNum, merchantNum int
	var goblinNum, dragonNum int
	var stoneNum, crystalNum, skippingStoneNum, skippingCrystalNum, grassesNum, lilyNum,
		gemNum, beamNum, mirrorNum, mirrorRefNum, mirrorAuxNum int

	finalNum += len(finals)
	socketNum += len(sockets)
	stoneNum += len(stonePosInit)
	lilyNum += len(lilyPosInit)
	goblinNum += len(goblinPosInit)
	dragonNum += len(dragonPosInit)

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
	for _, sw := range sameSwitches {
		if sw > 0 {
			switchMask |= 1 << (sw - 'x')
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
				case 'o':
					gemNum++
				case 'O':
					socketNum++
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
				case '~', '^', 'v', '<', '>', '\'', '-', '\\': // 水
					hasWater = true
				case '#', '.', '|', '_', 'L', 'e', 'N':
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

	if !isBigMap {
		if merchantNumberInit <= 1 && !finalsContains && finalNum != allCharNum {
			panic("地图 'f' 设置错误")
		}
	} else {
		if finalNum != 1 {
			panic("地图 'f' 设置错误")
		}
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
	// !allowCloneMan &&
	if merchantNum != merchantNumberInit {
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
	if gemNum != gemNumberInit {
		panic("没有修改 gemNumberInit")
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
		switch grid[p.x][p.y] {
		case '|', 'L', '\'', '\\':
			return true
		default:
			return false
		}
	case dir.y == 1: // 右
		switch grid[p.x][p.y+1] {
		case '|', 'L', '\'', '\\':
			return true
		default:
			return false
		}
	case dir.x == -1: // 上
		switch grid[p.x-1][p.y] {
		case '_', 'L', '-', '\\':
			return true
		default:
			return false
		}
	case dir.x == 1: // 下
		switch grid[p.x][p.y] {
		case '_', 'L', '-', '\\':
			return true
		default:
			return false
		}
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
func (d *data) getAllCharPos() []point {
	//if isBigMap {
	//	return nil
	//}

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
func (d *data) getAllMovableObjPos(onlyLife bool) (all, chars, nonChars []point) {
	chars = d.getAllCharPos()
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
	if gemNumberInit > 0 {
		for _, p := range d.gems {
			if p != noPos {
				all = append(all, p)
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

func (d *data) getAllLife() (life, nonLife []point) {
	life = d.getAllCharPos()
	if mirrorDirInit != "" {
		for _, p := range d.mirrors {
			if p.point != noPos {
				nonLife = append(nonLife, p.point)
			}
		}
	}
	if mirrorRefDirInit != "" { // 水晶镜子（可被反射的镜子）
		for _, p := range d.mirrorRefs {
			if p.point != noPos {
				life = append(life, p.point)
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
	if gemNumberInit > 0 {
		for _, p := range d.gems {
			if p != noPos {
				nonLife = append(nonLife, p)
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
func (d *data) withinBeams(p point, allNonCharObjs []point) (beamIndex []uint8, typeMask uint16, beamNf beamInfo) {
	// 前提是宝石在插座上
	if len(sockets) > 0 {
		gems := d.gems[:]
		for _, socket := range sockets {
			if !slices.Contains(gems, socket) { // 宝石不在插座上
				return
			}
		}
	}

	for bid, beam := range d.beams {
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
			// 出界，或者遇到不可穿透对象（石头、宝石、门）
			// 可以在地图边界加一圈 '#' 
			if !inBound(cur) || // 出界
				slices.Contains(allNonCharObjs, cur) || // 水晶在上面判断了
				d.inAnyClosedDoors(cur) { // 怪物门、活塞门
				cur = old
				break
			}
		}

		if meetP {
			beamIndex = append(beamIndex, uint8(bid))
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
	case '~', '^', 'v', '<', '>', 'l', '\'', '-', '\\': // 水
		// 继续
	default:
		return false
	}

	downP := point{p.x, p.y, -1}
	// 水中的物品（石头、水晶、水漂石、睡莲叶）
	// todo 光束、宝石等
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
			if isNeighbor6(g.point, p) {
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
	// 活
	dieTypeNo = iota
	dieTypeProtected
	dieTypeNoUpper

	// 死
	dieTypeCrushed
	dieTypeAttacked
	dieTypeDrown
)

func (d *data) getDieType(p point, burnPos []point, isChar bool) int {
	// 被门压死
	// todo 忽略向上的门（应该抬高角色）
	if d.inAnyClosedDoors(p) {
		return dieTypeCrushed
	}

	// 被攻击的优先级更高
	if d.isAttacked(p, burnPos) {
		if !isChar {
			return dieTypeAttacked
		}
		if d.isProtected(p) {
			return dieTypeProtected // 注：如果下面是空或者水，不会落下去
		}
		if len(d.crystalHeroMask) > 0 && d.crystalHeroMask[:][0] > 0 && d.crystalHeroMask[:][0]>>(d.getCharType(p)-1)&1 > 0 {
			return dieTypeNo
		}
		return dieTypeAttacked
	}

	// 淹死
	if d.isFallIntoWater(p) {
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

	if gemNumberInit > 0 {
		if i := slices.Index(d.gems[:], oldP); i >= 0 {
			d.gems[i] = newP
			return true
		}
	}

	if beamDirInit != "" && allowPushBeam {
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
		if merchantNumberInit > 0 {
			newData.merchant[:][0] = noPos
		}

		newData.sailor[:][0] = newP
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
	timeAtStart := time.Now()
	memAtStart := runtime.MemStats{}
	runtime.ReadMemStats(&memAtStart)

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
	gemInitArr := gemArrType{}
	beamInitArr := beamArrType{}

	__initCharTypeNum := initCharTypeNum
	if warriorPosInit != noPos {
		__initCharTypeNum = charWarrior
	}
	if thiefPosInit != noPos {
		__initCharTypeNum = charThief
	}
	if wizardPosInit != noPos {
		__initCharTypeNum = charWizard
	}
	if priestPosInit != noPos {
		__initCharTypeNum = charCleric
	}
	if bardPosInit != noPos {
		__initCharTypeNum = charBard
	}
	if sailorPosInit != noPos {
		__initCharTypeNum = charSailor
	}
	if isBigMap {
		if sailorNumberInit != 1 {
			panic("isBigMap 为 true 时，sailorNumberInit 必须为 1")
		}
		__initCharTypeNum = charSailor
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
	for _, pd := range goblinPosInit {
		__goblins = append(__goblins, pd)
	}
	__dragons := dragonInitArr[:0]
	for _, pd := range dragonPosInit {
		__dragons = append(__dragons, pd)
	}
	__gems := gemInitArr[:0]
	__beams := beamInitArr[:0]

	__handledDoorCnt := 0

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
						if __initCharTypeNum < 0 {
							__initCharTypeNum = charWarrior
						}
						__warriors = append(__warriors, p)
					}
				case 'T':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __initCharTypeNum < 0 {
							__initCharTypeNum = charThief
						}
						__thiefs = append(__thiefs, p)
					}
				case 'W':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __initCharTypeNum < 0 {
							__initCharTypeNum = charWizard
						}
						__wizards = append(__wizards, p)
					}
				case 'C':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __initCharTypeNum < 0 {
							__initCharTypeNum = charCleric
						}
						__priests = append(__priests, p)
					}
				case 'D':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __initCharTypeNum < 0 {
							__initCharTypeNum = charDruid
						}
						__druids = append(__druids, p)
					}
				case 'B':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __initCharTypeNum < 0 {
							__initCharTypeNum = charBard
						}
						__bards = append(__bards, p)
					}
				case '7':
					if isBigMap {
						if __sailor == noPos {
							__sailor = p
						}
					} else {
						if __initCharTypeNum < 0 {
							__initCharTypeNum = charExplorer
						}
						__explorers = append(__explorers, p)
					}
				case '8':
					if __initCharTypeNum < 0 {
						__initCharTypeNum = charSailor
					}
					if __sailor == noPos {
						__sailor = p
					}
					__sailors = append(__sailors, p)
				case '9':
					if __initCharTypeNum < 0 {
						__initCharTypeNum = charMerchant
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
					p.z = -1
					__lilies = append(__lilies, pointWithDir{p, dirStop})
				case 'g':
					__goblins = append(__goblins, pointWithDir{p, 0}) // todo 默认方向为 dirs[0]
				case 'd':
					idx := len(__dragons)
					__dragons = append(__dragons, pointWithDir{p, getDir(dragonDirInit[idx])})
				case 'O':
					sockets = append(sockets, p)
				case 'o':
					__gems = append(__gems, p)
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
					if doorDirString != "" {
						dir = getDir(doorDirString[__handledDoorCnt])
						__handledDoorCnt++
					}
					doors[ch-'X'] = append(doors[ch-'X'], pointWithDir{p, dir})
				case 'N':
					monsterDoors = append(monsterDoors, p)
				case 'f':
					finals = append(finals, p)
				case '.', '#', '|', '_', 'L', 'e':
					// ignore
				case '~', '^', 'v', '<', '>', '\'', '-', '\\': // 水
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

	if __handledDoorCnt != len(doorDirString) {
		panic(fmt.Sprintf("有 %d 个门的方向没有设置", len(doorDirString)-__handledDoorCnt))
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
	slices.Sort(validChars)

	if !slices.Contains(validChars, __initCharTypeNum) {
		panic(fmt.Sprintf("请修改 initCharTypeNum, as %v", validChars))
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

		gems:  gemInitArr,
		beams: beamInitArr,

		curCharTypeNum: __initCharTypeNum,
	}

	if initCrystalHeroMask > 0 {
		levelData.crystalHeroMask[:][0] = initCrystalHeroMask
	}

	type pair struct {
		data
		info string
	}
	from := map[data]pair{} // 同时充当 vis 的功能
	queue := []data{}
	defer func() {
		delta := time.Since(timeAtStart)
		memCur := runtime.MemStats{}
		runtime.ReadMemStats(&memCur)
		fmt.Printf("\n")
		fmt.Printf("// 搜索了 %d 个状态 (%.2fs, %.0f MB)\n",
			len(from),
			delta.Seconds(),
			float64(memCur.TotalAlloc-memAtStart.TotalAlloc)/float64(1<<20),
		)
	}()

	//vis := map[string]bool{}

	add := func(last, d data, info string) {
		if !allowSkippingStoneGone && (pdContains(d.skippingStones[:], noPos) || pdContains(d.skippingCrystals[:], noPos)) {
			return
		}

		// 剪枝
		//if d.dragons[0].x != mapSizeN-3 && d.dragons[0].y != 5 {
		//	return
		//}

		/* 常见错误
		没有预处理负数下标 -> 见 initMap
		方向写错
		*/
		//debugInfo := fmt.Sprintf("%v", d.stones)
		//if !vis[debugInfo] {
		//	vis[debugInfo] = true
		//	fmt.Println(debugInfo)
		//}

		allMovableObjs, _, _ := d.getAllMovableObjPos(false)
		animeTypeMask := uint8(0)

		// 水漂石
		changed := false
		doSkippingStones := func(skippingStones []pointWithDir) {
			if isSkippingAndLilySlow {
				lilies := d.lilies[:]
				for i, p := range skippingStones {
					// 不动
					if p.dir == dirStop {
						if d.isFallIntoWater(p.point) {
							if !allowFallIntoWater {
								return
							}
							animeTypeMask |= 1 << animeFallIntoWater
							skippingStones[i].z = -1
						}
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
				// 已经跑完了水漂动画
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
			allMovableObjs, _, _ = d.getAllMovableObjPos(false)
		}

		// 先确定门的开闭，方便后面判断下落
		var pushedDoors []pointWithDir
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

			beforeOpened := d.doorOpened[i]
			d.doorOpened[i] = opened

			// 如果门是侧向门，且是从开到关，则从门的位置开始推物品
			if beforeOpened && !opened && doors[i][0].dir < 4 {
				for _, door := range doors[i] {
					if slices.Contains(allMovableObjs, door.point) {
						pushedDoors = append(pushedDoors, door)
					}
				}
			}

			// 水晶被门压碎（水晶在门中，但门没有打开）
			//if !opened {
			//	for j, p := range d.crystals {
			//		if d.inAnyClosedDoors(p) {
			//			if !canDestroyObj {
			//				return
			//			}
			//			d.crystals[j] = noPos
			//		}
			//	}
			//	for j, p := range d.skippingCrystals {
			//		if d.inAnyClosedDoors(p.point) {
			//			if !canDestroyObj {
			//				return
			//			}
			//			d.skippingCrystals[j].point = noPos
			//		}
			//	}
			//}
		}

		// todo 多个门同时推物品时，优先级是什么样的？
		//  推走的物品又影响了门的开闭，岂不是要写递归？可能无限递归？
		for _, door := range pushedDoors {
			// 从门的位置开始推物品（逻辑同战士）
			p0 := door.point
			dir := directions4[door.dir]

			// 该方向有多少个连续的对象
			cnt := 1
			cur := p0.add(dir)
			for slices.Contains(allMovableObjs, cur) && !hasFence(cur, dir.rev()) {
				cnt++
				cur = cur.add(dir)
			}

			// 如果末端没有空地，则摧毁第一个物品，其余不变
			if !d.isValidPos(cur) || hasFence(cur, dir.rev()) {
				d.changePos(p0, noPos, math.MaxUint8, allMovableObjs)
				continue
			}

			// 倒着回来
			for range cnt {
				back := cur.sub(dir) // cur 前一个位置的物品移到 cur
				d.changePos(back, cur, math.MaxUint8, allMovableObjs)

				if mapSizeH > 1 {
					oldTop := point{back.x, back.y, back.z + 1}
					// todo 喷火龙 / 镜子
					if slices.Contains(allMovableObjs, oldTop) {
						newTop := cur
						newTop.z++
						d.changePos(oldTop, newTop, door.dir, allMovableObjs)
					}
				}

				cur = back
			}
		}

		if len(pushedDoors) > 0 {
			allMovableObjs, _, _ = d.getAllMovableObjPos(false)
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
		isNonPriestProtected := false
		for _, char := range d.getAllCharPos() {
			if len(d.cleric) > 0 && slices.Contains(d.cleric[:], char) {
				continue
			}
			tp := d.getDieType(char, burnedPos, true)
			if tp > dieTypeNoUpper {
				return
			}
			if tp == dieTypeProtected {
				isNonPriestProtected = true
			}
		}
		// 单独判断牧师
		for _, p := range d.cleric {
			// 只需判断压死或者淹死（后者需要保证牧师为非漂浮状态）
			if d.inAnyClosedDoors(p) || !isNonPriestProtected && !d.isAttacked(p, burnedPos) && d.isFallIntoWater(p) {
				return
			}
		}

		// 一开始，以及切换角色，都不结算怪物之间的攻击
		isSwitching := info[0] == 'c' || '1' <= info[0] && info[0] <= '9'
		if !isSwitching && !d.monsterDoorOpened {
			// 哥布林
			goblins := d.goblins
			if len(d.goblins) > 0 {
				//cnt := 0
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
					if tp := d.getDieType(p.point, burnedPos, false); tp > dieTypeNoUpper {
						if !canDestroyObj {
							return
						}
						if tp == dieTypeAttacked {
							animeTypeMask |= 1 << animeKill
						} else if tp == dieTypeDrown {
							animeTypeMask |= 1 << animeFallIntoWater
						}
						//cnt++
						goblins[i] = noPosDir
					}
				}

				//if cnt != 0 && cnt != 4 {
				//	return
				//}

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
					if tp := d.getDieType(p.point, burnedPos, false); tp > dieTypeNoUpper {
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
				if d.inAnyClosedDoors(p.point) { // 被门压碎
					mir[i] = noPosDir
				} else if d.isFallIntoWater(p.point) {
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
				if d.inAnyClosedDoors(p.point) { // 被门压碎
					mir[i] = noPosDir
				} else if d.isFallIntoWater(p.point) {
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
				if d.inAnyClosedDoors(p.point) { // 被门压碎
					mir[i] = noPosDir
				} else if d.isFallIntoWater(p.point) {
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
				if d.inAnyClosedDoors(p) { // 被门压碎
					sto[i] = noPos
				} else if d.isFallIntoWater(p) {
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
				if d.inAnyClosedDoors(p) { // 被门压碎
					cry[i] = noPos
				} else if d.isFallIntoWater(p) {
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

		// 宝石
		if len(d.gems) > 1 {
			slices.SortFunc(d.gems[:], cmpPoint)
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

		allMovableObjs, allChars, allNonChars := d.getAllMovableObjPos(false)

		var pass bool
		if targetIsOpenDoorY {
			pass = d.doorOpened[:][1]
		} else if targetIsTransAllGrass {
			//pass = d.grass[3] == noPos
		} else if targetIsClearAllMonsters {
			// 简化版：怪物门开启（怪物都被杀）
			pass = d.monsterDoorOpened
		} else {
			// 标准版：所有人都到达终点
			if isBigMap {
				p := d.getCurCharPos()
				pass = slices.Equal([]point{p}, finals)
			} else if len(d.crystalHeroMask) == 0 || d.crystalHeroMask[:][0] == 0 { // 不能有人是水晶
				if finalsContains {
					// 只要所有人都在终点即可
					for _, p := range allChars {
						if !slices.Contains(finals, p) {
							goto finalsContainsNext
						}
					}
					pass = true
				finalsContainsNext:
				} else {
					if len(allChars) > 1 {
						slices.SortFunc(allChars, cmpPoint)
					}
					pass = slices.Equal(allChars, finals)
				}
			}
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

		// 原地不动
		if isSkippingAndLilySlow {
			add(d, d, ".")
		}

		doWand := func() {
			if hasMoveWand {
				// 位移杖
				p0 := d.getCurCharPos()
				// todo 还得知道当前角色面朝的方向   也可以自己手调

			nextWandDir:
				for _, dir := range directions4 {
					newP := p0.add(dir)
					if !slices.Contains(allMovableObjs, newP) {
						continue
					}

					// 该方向有多少个连续的对象
					objs := []point{newP}
					cur := newP.add(dir)
					for {
						if !inBound(cur) {
							continue nextWandDir
						}
						if !d.isValidPos(cur) { // 墙、草、怪物门、活塞门
							// ok
						} else if slices.Contains(allMovableObjs, cur) {
							objs = append(objs, cur)
						} else {
							objs = append(objs, cur)
							break
						}
						cur = cur.add(dir)
					}

					newData := d
					for i := len(objs) - 2; i >= 0; i-- {
						newData.changePos(objs[i], objs[i+1], math.MaxUint8, allMovableObjs)
					}
					add(d, newData, "x")
				}
			} else if hasTorcWand {
				// 转向杖（只能用于光束）
			}
		}
		doWand()

		// todo 如果角色的头上有物品，物品会跟着移动（注意镜子的方向会变）    堆叠上限是多少？？
		// todo 即使人没有移动，切换方向也会改变头上物品（镜子、激光等）的方向
		// todo 多控时，如果下一个位置是没有石头的水，则一个角色无法移动（已在商人中实现）

		// 先考虑按 x 镜子反射对象，这样后面移动更流畅
		// 只要有一个镜子反射失败（红光），所有镜子都无法反射，直接 return
		doMirrors := func() {
			newData := d
			refed := uint(0)
		nextMirror:
			for _, mirror := range append(d.mirrors[:], d.mirrorRefs[:]...) {
				if mirror == noPosDir {
					continue
				}

				// 找两个方向最近的可反射的对象
				// 注意：如果一方向是物品，另一方向是墙草门，那么失败
				cur0 := mirror.point
				cur1 := mirror.point
				dir0 := directions6[mirror.dir&0xf]
				dir1 := directions6[mirror.dir>>4]
				wallMask := uint8(0)
				//foundMirror := uint8(0)

				for step := 1; ; step++ {
					justFound := uint8(0) // 是否找到了非镜子对象
					// 检查方向 0
					if wallMask&1 == 0 {
						cur0 = cur0.add(dir0)
						if !d.isValidPos(cur0) { // 无法反射（墙、草、门）
							wallMask |= 1
						} else if pdContains(d.mirrors[:], cur0) || pdContains(d.mirrorAuxes[:], cur0) {
							wallMask |= 1 // 镜子也视作「墙」
						} else if slices.Contains(allMovableObjs, cur0) {
							justFound |= 1
						}
					}
					// 检查方向 1
					if wallMask>>1 == 0 {
						cur1 = cur1.add(dir1)
						if !d.isValidPos(cur1) { // 无法反射（墙、草、门）
							wallMask |= 2
						} else if pdContains(d.mirrors[:], cur1) || pdContains(d.mirrorAuxes[:], cur1) {
							wallMask |= 2 // 镜子也视作「墙」
						} else if slices.Contains(allMovableObjs, cur1) {
							justFound |= 2
						}
					}
					if wallMask == 3 { // 两边都没有东西，继续看下一个镜子
						continue nextMirror
					}
					//if foundMirror == 3 {
					//	return // 不能两方向最近都是镜子
					//}
					if justFound|wallMask == 3 {
						return // 不能反射位置都是对象或者墙
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

					// 无法反射的石头，视作墙壁，都不能反射，那就没有红光
					if slices.Contains(d.stones[:], oldP) || pdContains(d.skippingStones[:], oldP) {
						continue nextMirror
					}

					// 反射
					newP := d.reflectTo(mirror, dir, step, allMovableObjs)
					if newP == noPos {
						return // 反射失败
					}
					itemIdx := slices.Index(allMovableObjs, oldP)
					if refed>>itemIdx&1 > 0 {
						// 复制对象
						if slices.Contains(d.merchant[:], oldP) {
							//if !allowCloneMan {
							//	panic("禁止复制人！")
							//	return
							//}
							if newData.merchant[:][0] != noPos { // 其实这样就可以禁止复制人了吧
								return
							}
							newData.merchant[:][0] = newP
						} else if ci := slices.Index(d.crystals[:], oldP); ci >= 0 {
							// todo 同时复制多个水晶
							// 复制水晶
							if newData.crystals[:][0] != noPos {
								panic("错误，无法复制水晶")
							}
							newData.crystals[:][0] = newP
							//for i := range newData.crystals {
							//	if newData.crystals[i] == noPos {
							//		newData.crystals[i] = newP
							//		goto copyCrystalNext
							//	}
							//}
							//panic("没有复制水晶，代码有误！例如数组大小开小了")
							//copyCrystalNext:
						} else {
							// todo 其他对象的复制
						}
					} else {
						refed |= 1 << itemIdx
						if i := pdIndex(newData.dragons[:], oldP); i >= 0 { // 喷火龙
							// todo 多次反射
							newDir := mirror.reflectDragon(newData.dragons[i].dir)
							newData.dragons[i] = pointWithDir{newP, newDir}
						} else if i := pdIndex(newData.skippingCrystals[:], oldP); i >= 0 { // 水漂石
							// todo 多次反射
							newDir := newData.skippingCrystals[i].dir
							if newDir != dirStop {
								if d.isFallIntoWater(newP) { // 还在水上
									newDir = mirror.reflectDragon(newDir)
								} else { // 传到地上/物品上，立刻停下
									newDir = dirStop
								}
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
					// todo 目前只支持两人
					if merchantNumberInit == 2 {
						man := newData.merchant[:]
						if man[0] != noPos && man[0] == man[1] {
							man[0] = noPos
						}
					}

					// 合并水晶
					if len(d.crystals) > 1 {
						c := d.crystals[:]
						slices.SortFunc(c, cmpPoint)
						// 相同且不等于 noPos
						for i := range len(c) - 1 {
							if c[i] != noPos && c[i] == c[i+1] {
								c[i] = noPos // 连续相同位置的水晶只保留最后一个
							}
						}
					}
				}
			}

			if refed == 0 {
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

		// todo 添加人物朝向
		doTorc := func(beamIndex []uint8, dir uint8) {
			for _, i := range beamIndex {
				if d.beams[i].dir != dir {
					goto doTorcNext
				}
			}
			return // 没有修改

		doTorcNext:
			newData := d
			for _, i := range beamIndex {
				newData.beams[i].dir = dir
			}
			add(d, newData, "x")
		}
		_ = doTorc

		// 移动当前角色
		switch d.curCharTypeNum {
		case charWarrior:
			// 普通移动一步
			p0 := d.warrior[:][0] // todo 暂时支持一个人
			doElevator(p0)

			_, withinBeams, beamNf := d.withinBeams(p0, allNonChars)

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

				if withinBeams>>beamDouble&1 > 0 {
					newP = newP.add(dir)
				}

				// todo 推 double 光中的物品

				// 该方向有多少个连续的对象
				cnt := 0
				cur := newP
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
					back := cur.sub(dir) // 这是个物品
					newData.changePos(back, cur, math.MaxUint8, allMovableObjs)

					// todo 多层
					if mapSizeH > 1 {
						oldTop := point{back.x, back.y, back.z + 1}
						// todo 喷火龙 / 镜子
						if slices.Contains(allMovableObjs, oldTop) {
							newTop := cur
							newTop.z++
							newData.changePos(oldTop, newTop, uint8(dIdx), allMovableObjs)
						}
					}

					cur = back
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
				info := dir4String[dIdx]
				if withinBeams>>beamDouble&1 > 0 {
					info += "L"
				}
				add(d, newData, info)
			}
		case charThief:
			// 普通移动一步
			p0 := d.thief[:][0] // todo
			doElevator(p0)

			_, withinBeams, beamNf := d.withinBeams(p0, allNonChars)

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

				if withinBeams>>beamDouble&1 > 0 {
					newP = newP.add(dir)
				}

				// 前面是否有空地
				if !(withinBeams>>beamThrough&1 > 0 && inBound(newP) && levelMap[newP.z][newP.x][newP.y] == '#') &&
					(!d.isValidPos(newP) || slices.Contains(allMovableObjs, newP) || hasFence(p0, dir)) {
					continue // 枚举另一个方向
				}

				newData := d

				back := p0.sub(dir) // 被拉的物品
				if slices.Contains(allMovableObjs, back) && !hasFence(back, dir) {
					// todo 拉 double 光中的物品

					// 被拉的物品 -> p0
					newData.changePos(back, p0, math.MaxUint8, allMovableObjs)
				}

				newData.thief[:][0] = newP
				newData.bigMapForceSwapChar(p0, newP)
				info := dir4String[dIdx]
				if withinBeams>>beamDouble&1 > 0 {
					info += "L"
				}
				add(d, newData, info)
			}
		case charWizard:
			p0 := d.wizard[:][0]
			doElevator(p0)

			_, withinBeams, beamNf := d.withinBeams(p0, allNonChars)

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

			_, withinBeams, _ := d.withinBeams(p0, allNonChars)

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

			allLife, nonLife := d.getAllLife()
			items := allLife[:0]
			for _, p := range allLife {
				if chebyshevDis(p, p0) <= 2 {
					items = append(items, p)
				}
			}
			//if isBigMap {
			//	items = append(items, p0)
			//}

			// 普通移动一步
			// 切比雪夫距离 <= 2 的物品（包括自己）都移动一步
			for dIdx, dir := range directions4 {
				p0d := p0.add(dir)
				if !d.isValidPos(p0d) || hasFence(p0, dir) {
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
					// 如果是喷火龙，修改喷火龙的朝向
					// 注意这和推拉不同，魅惑是另一套逻辑
					if dragonDirInit != "" && canPushDragon {
						if i := pdIndex(newData.dragons[:], oldP); i >= 0 {
							dr := &newData.dragons[i]
							// todo 水晶状态下的龙会转向吗？
							if dr.dir&dirIsCrystal == 0 {
								dr.dir &^= 7
								dr.dir |= uint8(dIdx)
							}
						}
					}

					// item 往前移动一格
					newP := oldP.add(dir)
					if !d.isValidPos(newP) || slices.Contains(nonLife, newP) || hasFence(oldP, dir) { // 无法移动
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

					// todo if hasWater
					movedItems = append(movedItems, oldP)

					newData.changePos(oldP, newP, math.MaxUint8, allMovableObjs)
				}

				if !slices.Contains(unmovedItems, p0) {
					if newData.bard[:][0] != p0d {
						panic("诗人移动错误，代码有误")
					}

					// 特性：如果诗人脚下是物品，且该物品移动了，那么诗人可以再走一格
					// todo 对于物品叠物品的情况，也是同样的规则？
					if slices.Contains(movedItems, point{p0.x, p0.y, p0.z - 1}) {
						nxtP := p0d.add(dir)
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
						newData.crystalHeroMask[:][0] ^= 1 << (tp - 1)
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
		case charExplorer, charSailor:
			// 普通移动一步
			var p0 point
			if d.curCharTypeNum == charExplorer {
				p0 = d.explorer[:][0]
			} else {
				p0 = d.sailor[:][0]
			}
			doElevator(p0)

			if d.curCharTypeNum == charExplorer {
				// todo v 放下宝石
			}

			_, withinBeams, _ := d.withinBeams(p0, allNonChars)

			// 在墙里面，但不能穿透
			if withinBeams>>beamThrough&1 == 0 && inBound(p0) && levelMap[p0.z][p0.x][p0.y] == '#' {
				goto afterSwitch
			}

			for dIdx, dir := range directions4 {
				//if hasTorcWand {
				//	// 修改光束方向 todo 需要知道人物朝向
				//	doTorc(withinBeamIndex, uint8(dIdx))
				//}

				newP := p0.add(dir)

				// 前面是否有空地
				if !(withinBeams>>beamThrough&1 > 0 && inBound(newP) && levelMap[newP.z][newP.x][newP.y] == '#') &&
					(!d.isValidPos(newP) || d.curCharTypeNum == charExplorer && slices.Contains(allMovableObjs, newP) || hasFence(p0, dir)) {

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
				if allowAllPushItem && (allowExplorerPushItem || d.curCharTypeNum == charSailor) {
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

				if d.curCharTypeNum == charExplorer {
					newData.explorer[:][0] = newP
				} else {
					newData.sailor[:][0] = newP
				}
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
		// 换成其他人
		if !isBigMap {
			for _, char := range validChars {
				if char == d.curCharTypeNum {
					continue
				}
				// 不能是水晶
				if len(d.crystalHeroMask) > 0 && d.crystalHeroMask[:][0]>>(char-1)&1 > 0 {
					continue
				}
				newData := d
				newData.curCharTypeNum = char
				var info string
				if useNumWhenChangeTwoChar || len(allChars) > 2 {
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
		// 大地图换人见 bigMapForceSwapChar
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
	charSailor   // 同大地图角色
	charMerchant // Trader
)

var charNumToName = [...]byte{
	charWarrior:  'A', // 1
	charThief:    'T', // 2
	charWizard:   'W', // 3
	charCleric:   'C', // 4
	charBard:     'B', // 5
	charDruid:    'D', // 6
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
