package main

// github.com/EndlessCheng/codeforces-go
type move struct {
	x0, y0 int // 起点 
	dx, dy int // 移动方向
	step   int // 移动的步数（时间）
}

// 计算位于 (x0,y0) 的棋子在 dirs 这些方向上的所有合法移动
func generateMoves(x0, y0 int, dirs []struct{ x, y int }) []move {
	const size = 8
	moves := []move{{x0, y0, 0, 0, 0}} // 原地不动
	for _, d := range dirs {
		// 往 d 方向走 1,2,3,... 步
		x, y := x0+d.x, y0+d.y
		for step := 1; 0 < x && x <= size && 0 < y && y <= size; step++ {
			moves = append(moves, move{x0, y0, d.x, d.y, step})
			x += d.x
			y += d.y
		}
	}
	return moves
}

// 判断两个移动是否合法，即不存在同一时刻两个棋子重叠的情况
func isValid(m1, m2 move) bool {
	x1, y1 := m1.x0, m1.y0
	x2, y2 := m2.x0, m2.y0
	for i := range max(m1.step, m2.step) {
		// 每一秒走一步
		if i < m1.step {
			x1 += m1.dx
			y1 += m1.dy
		}
		if i < m2.step {
			x2 += m2.dx
			y2 += m2.dy
		}
		if x1 == x2 && y1 == y2 { // 重叠
			return false
		}
	}
	return true
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}} // 上下左右 + 斜向
var pieceDirs = map[byte][]struct{ x, y int }{'r': dirs[:4], 'b': dirs[4:], 'q': dirs}

func countCombinations(pieces []string, positions [][]int) (ans int) {
	n := len(pieces)
	// 预处理所有合法移动
	allMoves := make([][]move, n)
	for i, pos := range positions {
		allMoves[i] = generateMoves(pos[0], pos[1], pieceDirs[pieces[i][0]])
	}

	path := make([]move, n) // 注意 path 的长度是固定的
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}
	outer:
		// 枚举当前棋子的所有合法移动
		for _, move1 := range allMoves[i] {
			// 判断合法移动 move1 是否有效
			for _, move2 := range path[:i] {
				if !isValid(move1, move2) {
					continue outer // 无效，枚举下一个 m1
				}
			}
			path[i] = move1 // 直接覆盖，无需恢复现场
			dfs(i + 1)      // 枚举后续棋子的所有合法移动组合
		}
	}
	dfs(0)
	return
}
