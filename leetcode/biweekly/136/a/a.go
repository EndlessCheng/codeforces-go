package main

// https://space.bilibili.com/206214
func winningPlayerCount(n int, pick [][]int) (ans int) {
	cnts := make([][11]int, n)
	won := make([]bool, n)
	for _, p := range pick {
		x, y := p[0], p[1]
		cnts[x][y]++
		if !won[x] && cnts[x][y] > x {
			won[x] = true
			ans++
		}
	}
	return
}

func winningPlayerCount2(n int, pick [][]int) (ans int) {
	cnts := make([][11]int, n)
	for _, p := range pick {
		cnts[p[0]][p[1]]++
	}
	for i, cnt := range cnts {
		for _, c := range cnt {
			if c > i {
				ans++
				break
			}
		}
	}
	return
}
