package copypasta

import (
	"cmp"
	"fmt"
	"slices"
)

/*
2026.6.13
《Heroes of Sokoban》https://www.puzzlescript.net/play.html?p=6860122
《Heroes of Sokoban II: Monsters》https://www.puzzlescript.net/play.html?p=6910207
《Heroes of Sokoban III: The Bard and The Druid》https://www.puzzlescript.net/play.html?p=7072276
《Mirror Isles》https://alan.draknek.org/games/puzzlescript/mirrors.php
《Skipping Stones To Lonely Homes》https://alan.draknek.org/games/puzzlescript/skipping-stones.php
《PROMESST》https://silverspaceship.com/promesst/
《PROMESST2》https://silverspaceship.com/promesst2/
《ENIGMASH》https://jacklance.github.io/PuzzleScript/play.html?p=cfdcc6e23f1fb3e9de2fd42fafaf4d4c

牧师 {2 0}
===
法师换牧师 {0 0} {2 0}
法师换怪物 {2 0} {4 0}
===
牧师 {4 1}
===
法师换牧师 {4 0} {4 1}
法师换怪物 {4 1} {4 4}
法师换怪物 {4 4} {0 4}
法师 {1 1}
法师换怪物 {1 1} {4 1}
法师换怪物 {4 1} {4 4}
法师 {2 2}
===
牧师 {2 1}
===
法师换牧师 {2 2} {2 1}
法师换怪物 {2 1} {4 1}

*/
func heroesOfSokoban() []string {
	type point struct{ x, y int }
	dir4 := []point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	getMask := func(a []point) int {
		return a[0].x<<15 | a[0].y<<12 | a[1].x<<9 | a[1].y<<6 | a[2].x<<3 | a[2].y
	}

	const n = 5
	done := map[int]bool{}
	for x := range n {
		for y := range n {
			for i, d1 := range dir4 {
				x2, y2 := x+d1.x, y+d1.y
				if !(0 <= x2 && x2 < n && 0 <= y2 && y2 < n) {
					continue
				}
				for j := i + 1; j < len(dir4); j++ {
					d2 := dir4[j]
					x3, y3 := x+d2.x, y+d2.y
					if !(0 <= x3 && x3 < n && 0 <= y3 && y3 < n) {
						continue
					}
					a := []point{{x, y}, {x2, y2}, {x3, y3}}
					slices.SortFunc(a, func(a, b point) int { return cmp.Or(a.x-b.x, a.y-b.y) })
					done[getMask(a)] = true
				}
			}
		}
	}

	near := func(p, q point) bool {
		for _, d := range dir4 {
			if (point{p.x + d.x, p.y + d.y}) == q {
				return true
			}
		}
		return false
	}

	const whoPriest = 0
	const whoMage = 1
	type data struct {
		monster [3]point
		priest  point
		mage    point
		who     int
	}
	vis := map[data]bool{}
	Q := []data{}
	type pair struct {
		data
		s string
	}
	from := map[data]pair{}
	add := func(fr, d data, s string) {
		if !vis[d] {
			vis[d] = true
			Q = append(Q, d)
			from[d] = pair{fr, s}
		}
	}

	levelMap := data{[3]point{{0, 4}, {4, 0}, {4, 4}}, point{2, 2}, point{0, 0}, whoMage}
	add(data{}, levelMap, "")

	for {
		d := Q[0]
		Q = Q[1:]
		// 检查三个怪是否相邻
		if done[getMask(d.monster[:])] {
			path := []string{}
			for d != (data{}) {
				path = append(path, from[d].s)
				d = from[d].data
			}
			slices.Reverse(path)
			return path
		}

		// 只有两个怪相邻
		if near(d.monster[0], d.monster[1]) || near(d.monster[0], d.monster[2]) || near(d.monster[1], d.monster[2]) {
			continue
		}

		// 如果法师旁边没有牧师，但有怪
		if !near(d.mage, d.priest) && (near(d.mage, d.monster[0]) || near(d.mage, d.monster[1]) || near(d.mage, d.monster[2])) {
			continue
		}

		if d.who == whoPriest {
			// 移动牧师
			for _, dir := range dir4 {
				x, y := d.priest.x+dir.x, d.priest.y+dir.y
				np := point{x, y}
				if 0 <= x && x < n && 0 <= y && y < n && np != d.mage && !slices.Contains(d.monster[:], np) {
					add(d, data{d.monster, np, d.mage, d.who}, fmt.Sprint("牧师 ", np))
				}
			}
		} else {
			// 移动魔法师
		o:
			for _, dir := range dir4 {
				// 该方向上是否有其他人？
				x, y := d.mage.x, d.mage.y
				for {
					x += dir.x
					y += dir.y
					if !(0 <= x && x < n && 0 <= y && y < n) { // 没人
						break
					}

					np := point{x, y}
					if np == d.priest { // 和牧师交换位置
						add(d, data{d.monster, d.mage, np, d.who}, fmt.Sprint("法师换牧师 ", d.mage, np))
						continue o
					}

					for i, q := range d.monster {
						if np != q {
							continue
						}
						// 和怪交换位置
						newMonster := d.monster
						newMonster[i] = d.mage
						slices.SortFunc(newMonster[:], func(a, b point) int { return cmp.Or(a.x-b.x, a.y-b.y) })
						add(d, data{newMonster, d.priest, np, d.who}, fmt.Sprint("法师换怪物 ", d.mage, np))
						continue o
					}
				}

				x, y = d.mage.x+dir.x, d.mage.y+dir.y
				// 普通移动一步
				np := point{x, y}
				if 0 <= x && x < n && 0 <= y && y < n && np != d.priest && !slices.Contains(d.monster[:], np) {
					add(d, data{d.monster, d.priest, np, d.who}, fmt.Sprint("法师 ", np))
				}
			}
		}

		// 单纯换人
		add(d, data{d.monster, d.priest, d.mage, d.who ^ 1}, "===")
	}
}

/*

".....#f",
".Ws..#n",
".ss....",
"...#...",
".*.*.*.",
"...#...",
法师交换 下 {2 1}
法师 下 {3 1}
法师 下 {4 1}
法师交换 上 {1 1}
法师交换 右 {1 2}
法师 右 {1 3}
法师 右 {1 4}
法师 下 {2 4}
法师 右 {2 5}
法师交换 左 {2 2}
法师 下 {3 2}
法师 下 {4 2}
法师 右 {4 3}
法师交换 左 {4 1}
法师交换 上 {1 1}
法师 右 {1 2}
法师 右 {1 3}
法师 右 {1 4}
法师 下 {2 4}
法师 下 {3 4}
法师 右 {3 5}
法师 下 {4 5}
法师交换 上 {2 5}
法师 右 {2 6}
法师 上 {1 6}
法师 上 {0 6}
小结：和推箱子反过来，从近到远，链式把石头传递到远方


原型
".......",
"...#...",
"g..B..g",
"...#...",
".......",
左上上右右下下右上上右下左


"g..B..g",
".......",
"...#..f",
".......",
".......",
下左下下右右上右下右下左 右上上

原型
"sf*###",
"##.###",
"...nW#",
"##.#.B",
"sfs###",
法师 下 {3 4}
法师交换 右 {3 5}
=== 换成诗人 ===
诗人 上 {2 4}
=== 换成法师 ===
法师 左 {3 4}
法师交换 上 {2 4}
=== 换成诗人 ===
诗人 右 {3 5}
=== 换成法师 ===
法师 下 {3 4}
法师交换 右 {3 5}
=== 换成诗人 ===
诗人 上 {2 4}
=== 换成法师 ===
法师 左 {3 4}
法师交换 上 {2 4}
=== 换成诗人 ===
诗人 右 {3 5}
=== 换成法师 ===
法师 下 {3 4}
法师交换 右 {3 5}
=== 换成诗人 ===
诗人 上 {2 4}
=== 换成法师 ===
法师 左 {3 4}
法师交换 上 {2 4}
=== 换成诗人 ===
诗人 右 {3 5}
=== 换成法师 ===
法师 下 {3 4}
法师交换 右 {3 5}
=== 换成诗人 ===
诗人 上 {2 4}
=== 换成法师 ===
法师 左 {3 4}
法师交换 上 {2 4}
法师 左 {2 3}
法师 左 {2 2}
法师 左 {2 1}
=== 换成诗人 ===
诗人 上 {2 4}
=== 换成法师 ===
法师交换 右 {2 4}
=== 换成诗人 ===
诗人 右 {2 2}
=== 换成法师 ===
法师交换 左 {2 2}
法师交换 上 {0 2}
=== 换成诗人 ===
诗人 左 {2 3}
诗人 左 {2 2}
诗人 下 {3 2}
诗人 下 {4 2}
诗人 左 {4 1}


原型
".....##",
".#s..##",
".ss....",
".......",
".*P*B*.",
"...*...",
只需移动诗人
上下下左上上上上右下右下下


大地图 法师
#......nf
.......##
..~*~W.##
.......##
~*~~g####
.g...####
上上左换下换上
左左左换换
上上左到头上右下
右右右上上左上左
下下上


"#.######",
"ww######",
"ww######",
"ww######",
"ww##wwff",
"wwwwww##",
"wwDwW###",
德 法
左上右 左左
 上 右上左下左
 左 上右下左上
 左 下左
 左上 上右上
 上 左下右上左
 上 右上
 上 左下右上
 上 左上
 上 右下
 上 上上上下下
 上左 上左下下下右
 上 上左下下下右
 上 上左下下下右
 上 上左下右下左左上右下右上
 左 左下右右右
 右左 上左下右右
 右上左 上
 右 下右
 右 上右右
 右右

*/
func orderOfTheSinkingStar() []string {
	abs := func(x int8) int8 {
		if x < 0 {
			return -x
		}
		return x
	}

	type point struct{ x, y, z int8 }
	dir4 := []point{{0, 1, 0}, {-1, 0, 0}, {1, 0, 0}, {0, -1, 0}}
	dirString := []rune("右上下左")
	//dirString := "dwsa"
	type pointWithDir struct {
		point
		dir int8
	}
	manhattanDis := func(p, q point) int { return int(abs(p.x-q.x) + abs(p.y-q.y)) }     // todo z 轴
	chebyshevDis := func(p, q point) int { return int(max(abs(p.x-q.x), abs(p.y-q.y))) } // todo z 轴
	_ = manhattanDis
	noPos := point{-60, -60, -60}

	cmpPoint := func(a, b point) int { return int(cmp.Or(a.x-b.x, a.y-b.y, a.z-b.z)) }
	cmpPointWithDir := func(a, b pointWithDir) int { return int(cmp.Or(a.x-b.x, a.y-b.y)) }
	sort := func(a ...point) []point {
		slices.SortFunc(a, cmpPoint)
		return a
	}
	_ = sort
	isNeighbor := func(p, q point) bool {
		for _, dir := range dir4 {
			if (point{p.x + dir.x, p.y + dir.y, p.z + dir.z}) == q {
				return true
			}
		}
		return false
	}

	const (
		charDefault = iota
		charWarrior
		charThief
		charWizard
		charPriest
		charBard
		charDruid
	)
	charName := [...]string{
		charWarrior: "战士",
		charThief:   "盗贼",
		charWizard:  "法师",
		charPriest:  "牧师",
		charBard:    "诗人",
		charDruid:   "德鲁伊",
	}
	_ = charName
	mpChar := [...]int{
		'A': charWarrior,
		'T': charThief,
		'W': charWizard,
		'P': charPriest,
		'B': charBard,
		'D': charDruid,
	}
	_ = mpChar

	// todo 改下面的个数
	levelMap := []string{
		"#.######",
		"ww######",
		"ww######",
		"ww######",
		"ww##wwff",
		"wwwwww##",
		"wwDwW###",
	}
	__done := true // todo
	hasMonster := false

	// . 空地
	// ~ 水
	// # 墙
	// * / xyz 压力开关
	// n / XYZ 活塞门
	// f 终点
	type grassArrType [19]point // todo
	type stoneArrType [19]point // todo
	type goblinArrType [0]point
	type dragonArrType [0]pointWithDir
	type data struct {
		curChar int
		warrior point // A 推多个物品
		thief   point // T 拉一个物品
		wizard  point // W 交换
		priest  point // P 自己以及上下左右无敌
		bard    point // B 同时移动切比雪夫距离 <= 2 的对象
		druid   point // D 把草变成石头

		grasses grassArrType  // w
		stones  stoneArrType  // s
		goblins goblinArrType // g
		dragons dragonArrType // d

		//mirrors // m

		done bool
	}

	getAllChar := func(d *data) []point {
		allChars := []point{}
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
		return allChars
	}
	changePos := func(d *data, oldP, newP point) {
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
			i := slices.Index(d.stones[:], oldP)
			if i >= 0 {
				d.stones[i] = newP
			}

			j := slices.Index(d.goblins[:], oldP)
			if j >= 0 {
				d.goblins[j] = newP
			}

			if i < 0 && j < 0 {
				panic("没有发生修改")
			}
		}
	}

	grassInitArr := grassArrType{}
	for i := range grassInitArr {
		grassInitArr[i] = noPos
	}
	stoneInitArr := stoneArrType{}
	for i := range stoneInitArr {
		stoneInitArr[i] = noPos
	}
	goblinInitArr := goblinArrType{}
	dragonInitArr := dragonArrType{}
	__curChar := -1
	__warrior := noPos
	__thief := noPos
	__wizard := noPos
	__priest := noPos
	__bard := noPos
	__druid := noPos
	__grasses := grassInitArr[:0]
	__stones := stoneInitArr[:0]
	__goblins := goblinInitArr[:0]
	__dragons := dragonInitArr[:0]
	var weightSwitches, doors, finals []point
	pos := [128]point{}
	for i, row := range levelMap {
		for j, ch := range row {
			p := point{int8(i), int8(j), 0}
			pos[ch] = p
			switch ch {
			case 'A':
				if __curChar < 0 {
					__curChar = charWarrior
				}
				__warrior = p
			case 'T':
				if __curChar < 0 {
					__curChar = charThief
				}
				__thief = p
			case 'W':
				if __curChar < 0 {
					__curChar = charWizard
				}
				__wizard = p
			case 'P':
				if __curChar < 0 {
					__curChar = charPriest
				}
				__priest = p
			case 'B':
				if __curChar < 0 {
					__curChar = charBard
				}
				__bard = p
			case 'D':
				if __curChar < 0 {
					__curChar = charDruid
				}
				__druid = p
			case 'w':
				__grasses = append(__grasses, p)
			case 's':
				__stones = append(__stones, p)
			case 'g':
				hasMonster = true
				__goblins = append(__goblins, p)
			case 'd':
				hasMonster = true
				__dragons = append(__dragons, pointWithDir{p, -1}) // todo 用 <>^v 表示？
			case '*':
				weightSwitches = append(weightSwitches, p)
			case 'n':
				doors = append(doors, p)
			case 'f':
				finals = append(finals, p)
			}
		}
	}

	//if len(__stones) != len(stoneInitArr) {
	//	panic("石头个数错误")
	//}
	if len(__goblins) != len(goblinInitArr) {
		panic("哥布林个数错误")
	}
	if len(__dragons) != len(dragonInitArr) {
		panic("喷火龙个数错误")
	}

	fallIntoWater := func(d *data, p point) bool {
		if p.z == -1 || p == noPos {
			return false
		}
		ch := levelMap[p.x][p.y]
		if ch == '~' && !slices.Contains(d.stones[:], point{p.x, p.y, -1}) {
			return true
		}
		// todo
		//if 'X' <= ch && ch <= 'Z' {
		//	sw := pos[ch-'A'+'a']
		//	if d.wizard != sw && d.thief != sw && d.stones[0] != sw {
		//		return true
		//	}
		//}
		return false
	}

	validChars := []int{}
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

	levelData := data{
		curChar: __curChar,

		warrior: __warrior,
		thief:   __thief,
		wizard:  __wizard,
		priest:  __priest,
		bard:    __bard,
		druid:   __druid,

		grasses: grassInitArr,
		stones:  stoneInitArr,
		goblins: goblinInitArr,
		dragons: dragonInitArr,

		done: __done,
	}

	for _, row := range levelMap {
		fmt.Println(row)
	}
	fmt.Println("levelData", levelData)
	fmt.Println("weightSwitches", weightSwitches)
	fmt.Println("doors", doors)
	fmt.Println("finals", finals)

	n, m := int8(len(levelMap)), int8(len(levelMap[0]))
	isValidPos := func(d *data, x, y int8) bool {
		return 0 <= x && x < n && 0 <= y && y < m && levelMap[x][y] != '#' &&
			!slices.Contains(d.grasses[:], point{x, y, 0}) // todo
	}

	vis := map[data]bool{}
	queue := []data{}
	type pair struct {
		data
		info string
	}
	from := map[data]pair{}
	add := func(last, d data, info string) {
		// 先判断是否有人被怪物攻击
		if hasMonster { // todo 改成实时的？
			allChars := getAllChar(&d)
			for _, char := range allChars {
				if char == d.priest || d.priest != noPos && isNeighbor(char, d.priest) {
					continue
				}
				for _, g := range d.goblins {
					if isNeighbor(g, char) {
						return
					}
				}
			}

			// todo 喷火龙
		}

		// 草排序
		slices.SortFunc(d.grasses[:], cmpPoint)

		// 在水中且下面没有物品的石头，落入水中
		sto := d.stones[:]
		if len(sto) > 0 {
			for i, p := range sto {
				if fallIntoWater(&d, p) {
					sto[i].z = -1
				}
			}
			slices.SortFunc(sto, cmpPoint)
		}

		// 在水中且下面没有物品的哥布林，落入水中淹死，坐标改成 noPos
		gob := d.goblins[:]
		for i, p := range gob {
			if fallIntoWater(&d, p) {
				gob[i] = noPos
			}
		}
		// todo 怪物之间的攻击（如果是落水的情况呢？攻击还是死）
		if !d.done && isNeighbor(gob[0], gob[1]) {
			return
			gob[0] = noPos
			gob[1] = noPos
			d.done = false // todo
		}
		slices.SortFunc(gob, cmpPoint)

		slices.SortFunc(d.dragons[:], cmpPointWithDir)

		// todo 其他物品（镜子）的排序

		if !vis[d] {
			vis[d] = true
			queue = append(queue, d)
			from[d] = pair{last, info}
		}
	}

	add(data{}, levelData, "")

nextQ:
	for len(queue) > 0 {
		// 注意入队的时候修改了物品的位置（落水）
		d := queue[0]
		queue = queue[1:]

		allChars := getAllChar(&d)
		slices.SortFunc(allChars, cmpPoint)

		allObj := allChars
		for _, p := range d.stones {
			if p != noPos {
				allObj = append(allObj, p)
			}
		}
		//for _, p := range d.grasses {
		//	if p != noPos {
		//		allObj = append(allObj, p)
		//	}
		//}
		for _, p := range d.goblins {
			if p != noPos {
				allObj = append(allObj, p)
			}
		}
		for _, p := range d.dragons {
			if p.point != noPos {
				allObj = append(allObj, p.point)
			}
		}

		open := true
		for _, w := range weightSwitches {
			if !slices.Contains(allObj, w) {
				open = false
			}
		}

		canMoveTo := func(d *data, p point) bool {
			return isValidPos(d, p.x, p.y) &&
				!slices.Contains(allObj, p) &&
				(open || !slices.Contains(doors, p))
		}

		// 是否有人死了（被怪物攻击的情况在 add 中判断了）
		for _, char := range allChars {
			// 被压死
			if !open && slices.Contains(doors, char) {
				continue nextQ
			}

			// 淹死
			if fallIntoWater(&d, char) {
				continue nextQ
			}
		}

		// todo 是否达成目标
		// 标准版：所有人都到达终点     d.done && 
		pass := slices.Equal(allChars, finals)
		// 简化版：石头都在开关上
		//pass := slices.Equal(sort(d.stones[0], d.stones[1], d.stones[2], d.bard), weightSwitches)
		if pass {
			path := []string{}
			for d != (data{}) {
				var ok bool
				pr, ok := from[d]
				if !ok {
					panic("代码修改了 d，与存入的 d 不符")
				}
				path = append(path, pr.info)
				d = pr.data
			}
			slices.Reverse(path)
			return path
		}

		// 移动当前角色
		switch d.curChar {
		case charWarrior:
			// 普通移动一步
			cur := d.warrior
			for dIdx, dir := range dir4 {
				x, y, z := cur.x+dir.x, cur.y+dir.y, cur.z+dir.z
				// 找到该方向前面有连续多少个物品
				cnt := 0
				projP := point{x, y, z}
				for slices.Contains(allObj, projP) {
					cnt++
					projP.x += dir.x
					projP.y += dir.y
					projP.z += dir.z
				}
				// 前面是否有空地
				if !canMoveTo(&d, projP) {
					continue
				}

				newData := d
				for range cnt {
					nxt := point{projP.x - dir.x, projP.y - dir.y, projP.z - dir.z}
					changePos(&newData, nxt, projP)
					projP = nxt
				}
				np := point{x, y, z}
				newData.warrior = np
				info := fmt.Sprintf("士 %c", dirString[dIdx])
				add(d, newData, info)
			}
		case charThief:
			// 普通移动一步
			cur := d.thief
			for dIdx, dir := range dir4 {
				x, y, z := cur.x+dir.x, cur.y+dir.y, cur.z+dir.z
				np := point{x, y, z}
				if !canMoveTo(&d, np) {
					continue
				}
				newData := d
				back := point{cur.x - dir.x, cur.y - dir.y, cur.z - dir.z}
				if slices.Contains(allObj, back) {
					// 拉人/物 -> 当前位置
					changePos(&newData, back, cur)
				}
				newData.thief = np
				info := fmt.Sprintf("贼 %c", dirString[dIdx])

				//fmt.Println(info)

				add(d, newData, info)
			}
		case charWizard:
			cur := d.wizard
		nextDir:
			for dIdx, dir := range dir4 {
				// 该方向上是否有人/物
				x, y, z := cur.x, cur.y, cur.z
				for {
					x += dir.x
					y += dir.y
					z += dir.z
					// 出界或者有障碍物
					np := point{x, y, z}
					if !isValidPos(&d, np.x, np.y) || !open && slices.Contains(doors, np) {
						break
					}
					if !slices.Contains(allObj, np) { // 空地
						continue
					}
					// 和人/物交换位置
					newData := d
					changePos(&newData, np, cur)
					newData.wizard = np
					info := fmt.Sprintf("法 %c", dirString[dIdx]) // 法交换

					//fmt.Println(info)

					add(d, newData, info)
					continue nextDir
				}

				// 没有，那就普通移动一步
				np := point{cur.x + dir.x, cur.y + dir.y, cur.z + dir.z}
				if !canMoveTo(&d, np) {
					continue
				}
				newData := d
				newData.wizard = np
				info := fmt.Sprintf("法 %c", dirString[dIdx])

				//fmt.Println(info)

				add(d, newData, info)
			}
		case charPriest:
			// 普通移动一步
			cur := d.priest
			for dIdx, dir := range dir4 {
				np := point{cur.x + dir.x, cur.y + dir.y, cur.z + dir.z}
				if !canMoveTo(&d, np) {
					continue
				}
				newData := d
				newData.priest = np
				info := fmt.Sprintf("牧 %c %v", dirString[dIdx], np)
				add(d, newData, info)
			}
		case charBard:
			cur := d.bard
			items := []point{}
			for _, p := range allObj {
				if chebyshevDis(p, cur) <= 2 {
					items = append(items, p)
				}
			}

			// todo 如果踩在物品上，可以多走一格（前提是撞墙或者前面一个格子也可以踩）

			// 普通移动一步
			// 切比雪夫距离 <= 2 的物品（包括自己）都移动一步
			for dIdx, dir := range dir4 {
				x, y, z := cur.x+dir.x, cur.y+dir.y, cur.z+dir.z
				if !isValidPos(&d, x, y) {
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
					np := point{item.x + dir.x, item.y + dir.y, item.z + dir.z}
					if !isValidPos(&d, np.x, np.y) || !open && slices.Contains(doors, np) { // 挡住了
						unmovedItems = append(unmovedItems, item)
						continue
					}
					if chebyshevDis(np, cur) > 2 { // 力场最前面的点
						if !canMoveTo(&d, np) { // 不能与力场外的物品碰撞
							unmovedItems = append(unmovedItems, item)
							continue
						}
					} else if slices.Contains(unmovedItems, np) {
						// 力场后面的点，不能与前面移动失败的物品碰撞
						unmovedItems = append(unmovedItems, item)
						continue
					}
					changePos(&newData, item, np)
				}

				if !slices.Contains(unmovedItems, cur) {
					np := point{x, y, z}
					if newData.bard != np {
						panic("移动错误")
					}
					info := fmt.Sprintf("诗人 %c %v", dirString[dIdx], np)
					add(d, newData, info)
				}
			}
		case charDruid:
			cur := d.druid
			for dIdx, dir := range dir4 {
				np := point{cur.x + dir.x, cur.y + dir.y, cur.z + dir.z}
				// todo 目前只实现了草 <-> 石头的逻辑
				if i := slices.Index(d.grasses[:], np); i >= 0 {
					newData := d
					newData.stones[0] = newData.grasses[i] // 草变石
					newData.grasses[i] = noPos
					info := fmt.Sprintf("德 %c 草变石", dirString[dIdx])
					add(d, newData, info)
					continue
				}
				if i := slices.Index(d.stones[:], np); i >= 0 {
					newData := d
					newData.grasses[0] = newData.stones[i] // 石变草
					newData.stones[i] = noPos
					info := fmt.Sprintf("德 %c 石变草", dirString[dIdx])
					add(d, newData, info)
					continue
				}
				if !canMoveTo(&d, np) {
					continue
				}
				// 普通移动一步
				newData := d
				newData.druid = np
				info := fmt.Sprintf("德 %c", dirString[dIdx])
				add(d, newData, info)
			}
		}

		// 换成其他人
		for _, char := range validChars {
			if char != d.curChar {
				newData := d
				newData.curChar = char
				//info := "换" // fmt.Sprint() charName[char]
				info := "c"
				add(d, newData, info)
			}
		}
	}

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
	dirString := []rune("左右上下")
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

		// 枚举键盘输入（左右上下）
	nextMan:
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
						continue nextMan
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
				s := fmt.Sprintf("%s %c %v", info, dirString[dIdx], newMan)
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
