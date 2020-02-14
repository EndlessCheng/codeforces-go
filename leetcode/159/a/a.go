package main

func checkStraightLine(ps [][]int) (ans bool) {
	dx0, dy0 := ps[1][0]-ps[0][0], ps[1][1]-ps[0][1]
	for i := 1; i < len(ps)-1; i++ {
		dx, dy := ps[i+1][0]-ps[i][0], ps[i+1][1]-ps[i][1]
		if dx*dy0 != dx0*dy {
			return
		}
	}
	return true
}
