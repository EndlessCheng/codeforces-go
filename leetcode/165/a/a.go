package main

func tictactoe(moves [][]int) string {
	g := [3][3]byte{}
	for i, m := range moves {
		if i&1 == 0 {
			g[m[0]][m[1]] = 'A'
		} else {
			g[m[0]][m[1]] = 'B'
		}
	}
	for i := range g {
		if g[i][0] == 'A' && g[i][1] == 'A' && g[i][2] == 'A' {
			return "A"
		}
		if g[i][0] == 'B' && g[i][1] == 'B' && g[i][2] == 'B' {
			return "B"
		}
		if g[0][i] == 'A' && g[1][i] == 'A' && g[2][i] == 'A' {
			return "A"
		}
		if g[0][i] == 'B' && g[1][i] == 'B' && g[2][i] == 'B' {
			return "B"
		}
	}
	if g[0][0] == 'A' && g[1][1] == 'A' && g[2][2] == 'A' {
		return "A"
	}
	if g[0][0] == 'B' && g[1][1] == 'B' && g[2][2] == 'B' {
		return "B"
	}
	if g[2][0] == 'A' && g[1][1] == 'A' && g[0][2] == 'A' {
		return "A"
	}
	if g[2][0] == 'B' && g[1][1] == 'B' && g[0][2] == 'B' {
		return "B"
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
