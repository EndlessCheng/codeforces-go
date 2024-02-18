package main

// https://space.bilibili.com/206214
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func mostFrequentPrime(mat [][]int) int {
	dirs := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	cnt := map[int]int{}
	m, n := len(mat), len(mat[0])
	for i, row := range mat {
		for j, v := range row {
			for _, d := range dirs {
				x, y, v := i+d.x, j+d.y, v
				for 0 <= x && x < m && 0 <= y && y < n {
					v = v*10 + mat[x][y]
					// 优化：如果 v 在 cnt 中那么 v 一定是质数
					if cnt[v] > 0 || isPrime(v) {
						cnt[v]++
					}
					x += d.x
					y += d.y
				}
			}
		}
	}

	ans, maxCnt := -1, 0
	for v, c := range cnt {
		if c > maxCnt {
			ans, maxCnt = v, c
		} else if c == maxCnt {
			ans = max(ans, v)
		}
	}
	return ans
}
