package main

func tictactoe(moves [][]int) string {
	players := []byte{'A', 'B'}
	g := [3][3]byte{}
	for i, m := range moves {
		g[m[0]][m[1]] = players[i&1]
	}
	same := func(b byte, startI, startJ, dirX, dirY int) bool {
		for i := range g {
			if g[startI+dirX*i][startJ+dirY*i] != b {
				return false
			}
		}
		return true
	}
	for _, b := range players {
		for i := range g {
			if same(b, i, 0, 0, 1) ||
				same(b, 0, i, 1, 0) ||
				same(b, 0, 0, 1, 1) ||
				same(b, 0, len(g)-1, 1, -1) {
				return string(b)
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
