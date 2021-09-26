package main

// github.com/EndlessCheng/codeforces-go
func placeWordInCrossword(board [][]byte, word string) bool {
	m, n, k := len(board), len(board[0]), len(word)
	// 遍历行
	for _, row := range board {
		for j := 0; j < n; j++ {
			if row[j] == '#' {
				continue
			}
			// 遍历并匹配两个 # 之间的字符
			j0, ok1, ok2 := j, true, true
			for ; j < n && row[j] != '#'; j++ { // 注意这里的 j 就是外层循环的 j，因此整体复杂度是线性的
				if j-j0 >= k || row[j] != ' ' && row[j] != word[j-j0] { // 正序匹配 word
					ok1 = false
				}
				if j-j0 >= k || row[j] != ' ' && row[j] != word[k-1-j+j0] { // 倒序匹配 word
					ok2 = false
				}
			}
			if (ok1 || ok2) && j-j0 == k { // 只要正序和倒序中有一个匹配成功，且两个 # 之间的字符长度恰好为 word 的长度，则返回 true
				return true
			}
		}
	}

	// 遍历列（同上）
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if board[i][j] == '#' {
				continue
			}
			i0, ok1, ok2 := i, true, true
			for ; i < m && board[i][j] != '#'; i++ {
				if i-i0 >= k || board[i][j] != ' ' && board[i][j] != word[i-i0] {
					ok1 = false
				}
				if i-i0 >= k || board[i][j] != ' ' && board[i][j] != word[k-1-i+i0] {
					ok2 = false
				}
			}
			if (ok1 || ok2) && i-i0 == k {
				return true
			}
		}
	}
	return false
}
