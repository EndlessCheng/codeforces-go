package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// 一些常量定义
const (
	eps = 1e-8

	maxAnts              = 6   // 最大蚂蚁数
	antRadius            = 0.5 // 蚂蚁半径
	pheromoneWithoutCake = 2   // 不携带蛋糕时产生的信息素
	pheromoneWithCake    = 5   // 携带蛋糕时产生的信息素
)

// 点
type vec struct {
	x, y int
}

// 线段
type line struct {
	p1, p2 vec
}

// 向量基本运算
func (a vec) add(b vec) vec { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) len() float64  { return math.Hypot(float64(a.x), float64(a.y)) }
func (a vec) len2() int     { return a.x*a.x + a.y*a.y }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }
func (a line) vec() vec     { return a.p2.sub(a.p1) }

// 点到线段的距离
func (a vec) disToSeg(l line) float64 {
	v, p1a, p2a := l.vec(), a.sub(l.p1), a.sub(l.p2)
	if float64(v.dot(p1a)) < -eps {
		return p1a.len()
	}
	if float64(v.dot(p2a)) > eps {
		return p2a.len()
	}
	return math.Abs(float64(v.det(p1a))) / v.len()
}

// 地图格点
type grid struct {
	pheromone int  // 信息素
	empty     bool // 是否为空
}

// 蚂蚁
type ant struct {
	age     int  // 年龄
	level   int  // 等级
	maxHP   int  // 初始血量
	curHP   int  // 当前血量
	prevPos vec  // 上一秒位置
	pos     vec  // 当前位置
	hasCake bool // 是否扛着蛋糕
}

// 炮塔（激光塔）
type tower struct {
	damage   int // 伤害
	atkRange int // 攻击范围
	pos      vec // 位置
}

// 游戏数据
type game struct {
	height, width int      // 长，宽
	board         [][]grid // 地图格点
	antNest, cake vec      // 蚁穴位置，蛋糕位置
	antWithCake   *ant     // 拿着蛋糕的蚂蚁
	antGenCounter int      // 蚂蚁生成计数器
	ants          []*ant   // 蚂蚁
	towers        []*tower // 炮塔
}

// 新游戏
func newGame(height, width int) *game {
	g := &game{height: height, width: width, cake: vec{height, width}}
	g.board = make([][]grid, height+1)
	for i := range g.board {
		g.board[i] = make([]grid, width+1)
		for j := range g.board[i] {
			g.board[i][j].empty = true
		}
	}
	return g
}

// 生成蚂蚁
func (g *game) newAnt() {
	// 如果地图上蚂蚁数不足 6，并且洞口没有蚂蚁，一只蚂蚁就会在洞口出生
	if len(g.ants) >= maxAnts || !g.board[0][0].empty {
		return
	}
	level := g.antGenCounter/6 + 1
	hp := int(4 * math.Pow(1.1, float64(level)))
	a := &ant{level: level, maxHP: hp, curHP: hp}
	g.ants = append(g.ants, a)
	g.antGenCounter++
	g.board[0][0].empty = false
}

// 移动前，蚂蚁们在自己所在点留下一些信息素
func (g *game) beginSecond() {
	for _, a := range g.ants {
		if a.hasCake {
			g.board[a.pos.x][a.pos.y].pheromone += pheromoneWithCake
		} else {
			g.board[a.pos.x][a.pos.y].pheromone += pheromoneWithoutCake
		}
	}
}

// 格点 p 是否可达
func (g *game) canReach(p vec) bool {
	return 0 <= p.x && p.x <= g.height && 0 <= p.y && p.y <= g.width && g.board[p.x][p.y].empty
}

// 东南西北
var dir4 = [4]vec{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// 移动所有蚂蚁
func (g *game) moveAnts() {
	// 蚂蚁按出生的顺序移动，出生得比较早的蚂蚁先移动
	for _, a := range g.ants {
		maxPheromone := -1
		var dirI int
		// 蚂蚁的移动方向为东南西北
		for i, d := range dir4 {
			// 蚂蚁只能移动到空格点上，且不能是上一秒所在的点
			// 如果此时有多个选择，蚂蚁会选择信息素最多的那个点爬过去，
			// 若有多个相同的最多信息素，则按照东南西北的顺序选择有最多信息素的格点 todo:check
			if p := a.pos.add(d); p != a.prevPos && g.canReach(p) && g.board[p.x][p.y].pheromone > maxPheromone {
				maxPheromone = g.board[p.x][p.y].pheromone
				dirI = i
			}
		}
		// 如果蚂蚁的四周都是不可达点，那么蚂蚁在这一秒内会选择停留在当前点。下一秒判断移动方向时，它上一秒所在点为其当前停留的点
		if maxPheromone == -1 {
			a.prevPos = a.pos
			continue
		}
		// 若此时蚂蚁的年龄为 5n+4，它会在选择方向后不断逆时针转 90°，直到面对一个可达的点，这样定下的方向才是蚂蚁最终要爬去的方向
		if a.age%5 == 4 {
			for i := (dirI + 3) % 4; ; i = (i + 3) % 4 {
				if p := a.pos.add(dir4[i]); p != a.prevPos && g.canReach(p) {
					dirI = i
					break
				}
			}
		}
		// 移动蚂蚁
		g.board[a.pos.x][a.pos.y].empty = true
		a.prevPos, a.pos = a.pos, a.pos.add(dir4[dirI])
		g.board[a.pos.x][a.pos.y].empty = false
	}
}

// 更新拿到蛋糕的蚂蚁状态
func (a *ant) getCake() {
	a.hasCake = true
	a.curHP += a.maxHP / 2
	if a.curHP > a.maxHP {
		a.curHP = a.maxHP
	}
}

// 如果有蚂蚁在蛋糕的位置上并且蛋糕没被扛走，它把蛋糕扛上，血量增加
func (g *game) checkCake() {
	if g.antWithCake != nil {
		return
	}
	for _, a := range g.ants {
		if a.pos == g.cake {
			a.getCake()
			g.antWithCake = a
			break
		}
	}
}

// 炮台攻击蚂蚁
func (g *game) towerAttack() {
	// 所有塔同时开始攻击
	for _, t := range g.towers {
		// 只有当代表蚂蚁的圆的圆心与塔 t 的直线距离不超过 t.atkRange 时，塔才算打得到那只蚂蚁
		// 如果一只蚂蚁扛着蛋糕，任何打得到它的塔的炮口都会对准它
		if g.antWithCake != nil && t.pos.sub(g.antWithCake.pos).len2() <= t.atkRange*t.atkRange {
			// 塔到目标蚂蚁圆心的连线上的所有蚂蚁都会被打到并损失 t.damage 血量，这里“被打到”指表示激光的线段与表示蚂蚁的圆有公共点
			towerToAntSeg := line{t.pos, g.antWithCake.pos}
			for _, a := range g.ants {
				if a.pos.disToSeg(towerToAntSeg)-eps < antRadius {
					a.curHP -= t.damage
				}
			}
		} else {
			// 否则塔会挑离它最近的蚂蚁进行攻击，如果有多只蚂蚁，它会选出生较早的一只
			// 因为是最近的，所以炮塔是打不到其他蚂蚁的
			var targetAnt *ant
			minDis2 := int(1e9)
			for _, a := range g.ants {
				if dis2 := t.pos.sub(a.pos).len2(); dis2 < minDis2 && dis2 <= t.atkRange*t.atkRange {
					dis2 = minDis2
					targetAnt = a
				}
			}
			if targetAnt != nil {
				targetAnt.curHP -= t.damage
			}
		}
	}

	// 移除死亡蚂蚁
	newAnts := []*ant{}
	for _, a := range g.ants {
		// 当蚂蚁的血被打成负数时，它才算挂了
		if a.curHP >= 0 {
			newAnts = append(newAnts, a)
		} else if a.hasCake {
			// 如果攻击结束后那只扛着蛋糕的蚂蚁挂了，蛋糕瞬间归位
			g.antWithCake = nil
		}
	}
	g.ants = newAnts
}

// 地图上所有点的信息素损失 1 单位，所有蚂蚁的年龄加 1
func (g *game) endSecond() {
	for i, gi := range g.board {
		for j, gij := range gi {
			if gij.pheromone > 0 {
				g.board[i][j].pheromone--
			}
		}
	}
	for _, a := range g.ants {
		a.age++
	}
}

// 如果发现扛蛋糕的蚂蚁没死并在窝的位置，就认为蚂蚁抢到了蛋糕，游戏结束
func (g *game) isGameOver() bool {
	for _, a := range g.ants {
		if a.hasCake && a.pos == g.antNest {
			return true
		}
	}
	return false
}

// 模拟游戏的前 t 秒钟
func (g *game) runSeconds(t int) (gameOverAt int, isGameOver bool) {
	for s := 1; s <= t; s++ {
		g.newAnt()          // 生成蚂蚁
		g.beginSecond()     // 蚂蚁留下信息素
		g.moveAnts()        // 移动蚂蚁
		g.checkCake()       // 蚂蚁能否拿上蛋糕，更新拿到蛋糕的蚂蚁的状态
		g.towerAttack()     // 炮台攻击蚂蚁
		if g.isGameOver() { // 检查游戏是否结束
			// game over 的时候，这个回合不会对蚂蚁的年龄产生贡献，直接 return
			return s, true
		}
		g.endSecond() // 更新信息素和蚂蚁年龄
	}
	return
}

func p2586(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var height, width, numTower, damage, atkRange int
	Fscan(in, &height, &width, &numTower, &damage, &atkRange)

	// 创建新游戏
	g := newGame(height, width)

	// 放置炮塔
	g.towers = make([]*tower, numTower)
	for i := range g.towers {
		var x, y int
		Fscan(in, &x, &y)
		g.towers[i] = &tower{damage, atkRange, vec{x, y}}
		g.board[x][y].empty = false
	}

	// 模拟游戏的前 t 秒钟
	var time int
	Fscan(in, &time)
	if gameOverAt, isGameOver := g.runSeconds(time); isGameOver {
		Fprintf(out, "Game over after %d seconds\n", gameOverAt)
	} else {
		Fprintln(out, "The game is going on")
	}
	Fprintln(out, len(g.ants))
	for _, a := range g.ants {
		Fprintln(out, a.age, a.level, a.curHP, a.pos.x, a.pos.y)
	}
}

func main() { p2586(os.Stdin, os.Stdout) }
