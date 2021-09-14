package main

func maxDistance(grid [][]int) int {
	n := len(grid)
	type p struct{ i, j int }
	lands := []p{}
	vis := [100][100]bool{}
	left := n * n
	for i, g := range grid {
		for j, c := range g {
			if c == 1 {
				lands = append(lands, p{i, j})
				vis[i][j] = true
				left--
			}
		}
	}
	if left == 0 || left == n*n {
		return -1
	}

	dirOffset4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	searchDirOffset4 := func(n, centerI, centerJ, dis int) {
		for i, dir := range dirOffset4 {
			dx := dirOffset4[(i+1)%4][0] - dir[0]
			dy := dirOffset4[(i+1)%4][1] - dir[1]
			x := centerI + dir[0]*dis
			y := centerJ + dir[1]*dis
			for _i := 0; _i < dis; _i++ {
				if x >= 0 && x < n && y >= 0 && y < n {
					if !vis[x][y] {
						vis[x][y] = true
						left--
					}
				}
				x += dx
				y += dy
			}
		}
	}
	ans := 1
	for ; left > 0; ans++ {
		for _, land := range lands {
			searchDirOffset4(n, land.i, land.j, ans)
		}
	}
	return ans - 1
}
