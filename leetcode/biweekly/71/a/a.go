package main

import (
	"sort"
	"strconv"
)

/* 贪心 + 排序

通过比对可以发现，拆分成两位数+两位数是最优的。因此要使 $\textit{new1}+\textit{new2}$ 最小，应当最小化十位数字。

我们可以将 $\textit{num}$ 的字符串形式从小到大排序，那么前两个数字为十位数，后两个数字为个位数。

*/

// github.com/EndlessCheng/codeforces-go
func minimumSum(num int) int {
	s := []byte(strconv.Itoa(num))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return int(s[0]&15+s[1]&15)*10 + int(s[2]&15+s[3]&15)
}
