package main

// https://space.bilibili.com/206214
func minMoves(classroom []string, energy int) (ans int) {
	m, n := len(classroom), len(classroom[0])
	idx := make([][]int, m)
	for i := range idx {
		idx[i] = make([]int, n)
	}
	var cntL, sx, sy int
	for i, row := range classroom {
		for j, b := range row {
			if b == 'L' {
				idx[i][j] = cntL // 给垃圾分配编号
				cntL++
			} else if b == 'S' {
				sx, sy = i, j
			}
		}
	}

	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	u := 1 << cntL
	maxEnergy := make([][][]int8, m)
	for i := range maxEnergy {
		maxEnergy[i] = make([][]int8, n)
		for j := range maxEnergy[i] {
			maxEnergy[i][j] = make([]int8, u)
			for k := range maxEnergy[i][j] {
				maxEnergy[i][j][k] = -1
			}
		}
	}

	maxEnergy[sx][sy][u-1] = int8(energy)
	type tuple struct{ x, y, e, mask int }
	q := []tuple{{sx, sy, energy, u - 1}}

	for ; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.mask == 0 { // 所有垃圾清理完毕
				return
			}
			if p.e == 0 { // 走不动了
				continue
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && classroom[x][y] != 'X' {
					newE := p.e - 1
					if classroom[x][y] == 'R' {
						newE = energy // 充满能量
					}
					newMask := p.mask
					if classroom[x][y] == 'L' {
						newMask &^= 1 << idx[x][y] // 清理垃圾
					}
					if int8(newE) > maxEnergy[x][y][newMask] {
						maxEnergy[x][y][newMask] = int8(newE)
						q = append(q, tuple{x, y, newE, newMask})
					}
				}
			}
		}
	}
	return -1
}

func minMoves1(classroom []string, energy int) (ans int) {
	m, n := len(classroom), len(classroom[0])
	idx := make([][]int, m)
	for i := range idx {
		idx[i] = make([]int, n)
	}
	var cntL, sx, sy int
	for i, row := range classroom {
		for j, b := range row {
			if b == 'L' {
				idx[i][j] = cntL // 给垃圾分配编号
				cntL++
			} else if b == 'S' {
				sx, sy = i, j
			}
		}
	}

	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	u := 1 << cntL
	vis := make([][][][]bool, m)
	for i := range vis {
		vis[i] = make([][][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([][]bool, energy+1)
			for k := range vis[i][j] {
				vis[i][j][k] = make([]bool, u)
			}
		}
	}

	vis[sx][sy][energy][u-1] = true
	type tuple struct{ x, y, e, mask int }
	q := []tuple{{sx, sy, energy, u - 1}}
	for ; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.mask == 0 { // 所有垃圾清理完毕
				return
			}
			if p.e == 0 { // 走不动了
				continue
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && classroom[x][y] != 'X' {
					newE := p.e - 1
					if classroom[x][y] == 'R' {
						newE = energy // 充满能量
					}
					newMask := p.mask
					if classroom[x][y] == 'L' {
						newMask &^= 1 << idx[x][y] // 清理垃圾
					}
					if !vis[x][y][newE][newMask] {
						vis[x][y][newE][newMask] = true
						q = append(q, tuple{x, y, newE, newMask})
					}
				}
			}
		}
	}
	return -1
}
