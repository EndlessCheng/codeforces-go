package main

// github.com/EndlessCheng/codeforces-go
type move struct{ dx, dy, t int } // (dx,dy) 表示移动方向，t 表示移动的步数（时间）

func validMovesRook(x, y int) (m []move) {
	m = append(m, move{}) // 为了方便计算皇后
	for i := 1; i <= 8; i++ {
		if i != x {
			m = append(m, move{(i - x) / abs(i-x), 0, abs(i - x)})
		}
	}
	for j := 1; j <= 8; j++ {
		if j != y {
			m = append(m, move{0, (j - y) / abs(j-y), abs(j - y)})
		}
	}
	return
}

func validMovesBishop(x, y int) (m []move) {
	m = append(m, move{}) // 为了方便计算皇后
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if (i != x || j != y) && abs(i-x) == abs(j-y) {
				m = append(m, move{(i - x) / abs(i-x), (j - y) / abs(j-y), abs(i - x)})
			}
		}
	}
	return
}

func validMovesQueen(x, y int) []move { // 皇后可以有上面两种移动方式
	return append(append([]move{{}}, validMovesRook(x, y)[1:]...), validMovesBishop(x, y)[1:]...)
}

// 判断是否合法，即不存在两个棋子占据同一个格子的情况
func isValid(x1, y1, x2, y2 int, m1, m2 move) bool {
	for i := 1; i <= m1.t || i <= m2.t; i++ {
		if i <= m1.t {
			x1 += m1.dx // 每一秒走一步
			y1 += m1.dy
		}
		if i <= m2.t {
			x2 += m2.dx
			y2 += m2.dy
		}
		if x1 == x2 && y1 == y2 { // 两个棋子占据了同一个格子
			return false
		}
	}
	return true
}

func countCombinations(pieces []string, positions [][]int) (ans int) {
	n := len(pieces)
	validMoves := make([][]move, n)
	for i, p := range positions {
		x, y := p[0], p[1]
		if pieces[i] == "rook" {
			validMoves[i] = validMovesRook(x, y) // 预处理所有合法移动
		} else if pieces[i] == "bishop" {
			validMoves[i] = validMovesBishop(x, y)
		} else {
			validMoves[i] = validMovesQueen(x, y)
		}
	}

	moves := make([]move, n)
	var f func(int)
	f = func(i int) {
		if i == n {
			ans++
			return
		}
		x1, y1 := positions[i][0], positions[i][1]
	outer:
		for _, m := range validMoves[i] { // 枚举当前棋子的所有合法移动
			for j, pos := range positions[:i] { // 判断该移动是否与前面的棋子的移动相冲突
				if !isValid(x1, y1, pos[0], pos[1], m, moves[j]) {
					continue outer
				}
			}
			moves[i] = m // 无冲突
			f(i + 1)
		}
	}
	f(0)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
