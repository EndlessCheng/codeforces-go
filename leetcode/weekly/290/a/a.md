由于 $\textit{nums}[i]$ 所有元素均不相同，利用这一性质，用哈希表统计每个元素的出现次数，最后遍历哈希表，返回所有次数等于 $n$ 的元素，这里 $n$ 为 $\textit{nums}$ 的长度。

```go
func intersection(nums [][]int) (ans []int) {
	cnt := map[int]int{}
	for _, a := range nums {
		for _, v := range a {
			cnt[v]++
		}
	}
	for v, c := range cnt {
		if c == len(nums) {
			ans = append(ans, v)
		}
	}
	sort.Ints(ans)
	return
}
```
