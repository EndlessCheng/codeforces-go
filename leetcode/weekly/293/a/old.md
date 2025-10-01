对于根据相邻元素的性质来删除元素的题目，用栈来思考是最简单的。

用栈来模拟题目的操作，遍历的时候每次和栈顶比较，如果不是字母异位词就入栈。

遍历结束后，栈中所有相邻元素均不为字母异位词，栈即为答案。

相似题目：

- [1047. 删除字符串中的所有相邻重复项](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/)
- [2216. 美化数组的最少删除数](https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful/)
- [2197. 替换数组中的非互质数](https://leetcode.cn/problems/replace-non-coprime-numbers-in-array/)

```go
func removeAnagrams(words []string) []string {
	ans := []string{words[0]}
	for _, word := range words[1:] {
		cnt := [26]int{}
		for _, b := range word {
			cnt[b-'a']++
		}
		for _, b := range ans[len(ans)-1] {
			cnt[b-'a']--
		}
		if cnt != [26]int{} { // 不是字母异位词
			ans = append(ans, word)
		}
	}
	return ans
}
```