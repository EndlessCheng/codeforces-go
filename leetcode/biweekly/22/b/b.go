package main

// github.com/EndlessCheng/codeforces-go
func maxNumberOfFamilies(n int, seats [][]int) (ans int) {
	pos := map[int][]bool{}
	for _, p := range seats {
		x, y := p[0], p[1]
		if pos[x] == nil {
			pos[x] = make([]bool, 11) // 也可以用位运算
		}
		pos[x][y] = true
	}
	ans = 2*n - 2*len(pos)
	for _, ps := range pos {
	o:
		for i := 2; i < 7; i += 2 {
			for j := 0; j < 4; j++ {
				if ps[i+j] {
					continue o
				}
			}
			ans++
			i += 2
		}
	}
	return
}
