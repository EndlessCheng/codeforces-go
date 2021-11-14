package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func decodeCiphertext(encodedText string, row int) string {
	ans := []byte{}
	for i, j, k, col := 0, 0, 0, len(encodedText)/row; k < col; {
		ans = append(ans, encodedText[i*col+j]) // 转换成在 encodedText 上的下标
		i++
		j++
		if i == row || j == col { // 触及边界
			k++
			i, j = 0, k // 移至下一条斜向
		}
	}
	return strings.TrimRight(string(ans), " ") // 移除末尾多余空格
}
