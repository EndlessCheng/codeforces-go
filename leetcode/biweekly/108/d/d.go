package main

// https://space.bilibili.com/206214
func countBlackBlocks(m, n int, coordinates [][]int) []int64 {
	type pair struct{ x, y int }
	set := make(map[pair]int, len(coordinates))
	for _, p := range coordinates {
		set[pair{p[0], p[1]}] = 1
	}

	arr := make([]int64, 5)
	vis := make(map[pair]bool, len(set)*4)
	for _, p := range coordinates {
		x, y := p[0], p[1]
		for i := max(x-1, 0); i <= x && i < m-1; i++ {
			for j := max(y-1, 0); j <= y && j < n-1; j++ {
				if !vis[pair{i, j}] {
					vis[pair{i, j}] = true
					cnt := set[pair{i, j}] + set[pair{i, j + 1}] +
						   set[pair{i + 1, j}] + set[pair{i + 1, j + 1}]
					arr[cnt]++
				}
			}
		}
	}
	arr[0] = int64(m-1)*int64(n-1) - int64(len(vis))
	return arr
}

func max(a, b int) int { if b > a { return b }; return a }
