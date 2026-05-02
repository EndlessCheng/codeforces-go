package main

import "bytes"

// github.com/EndlessCheng/codeforces-go
func rotateTheBox1(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
	}

	for i, row := range boxGrid {
		stone := 0
		for j, ch := range row {
			if ch == '#' { // 石头
				stone++
				ch = '.' // 先把石头清空，下面循环会填入石头
			}
			ans[j][m-1-i] = ch
			if j == n-1 || row[j+1] == '*' { // 下一个格子是障碍物
				// 石头垂直掉落后，从 j 往前 stone 个格子都是石头
				for k := j; k > j-stone; k-- {
					ans[k][m-1-i] = '#'
				}
				stone = 0 // 重置计数器
			}
		}
	}

	return ans
}

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = bytes.Repeat([]byte{'.'}, m)
	}

	for i, row := range boxGrid {
		stone := n - 1
		for j := n - 1; j >= 0; j-- {
			if row[j] == '*' { // 障碍物
				ans[j][m-1-i] = '*'
				stone = j - 1 // 障碍物左边最近的石头，在旋转后掉落到 j-1
			} else if row[j] == '#' { // 石头
				ans[stone][m-1-i] = '#' // 旋转后，石头掉落到 stone
				stone--
			}
		}
	}

	return ans
}
