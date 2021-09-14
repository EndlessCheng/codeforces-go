package main

/*
由于奇数的最后一个数字是奇数，从后往前找到第一个奇数数字，删掉后面的字符串，剩下的就是最大的奇数
*/

// github.com/EndlessCheng/codeforces-go
func largestOddNumber(s string) string {
	i := len(s) - 1
	for ; i >= 0 && s[i]&1 == 0; i-- {
	}
	return s[:i+1]
}
