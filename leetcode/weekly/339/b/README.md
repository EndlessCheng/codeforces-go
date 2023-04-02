### 本题视频讲解

见[【周赛 339】](https://www.bilibili.com/video/BV1va4y1M7Fr/)。

### 思路

用一个哈希表 $\textit{cnt}$ 记录每个元素的剩余次数。构造答案时，如果元素 $x$ 的剩余次数为 $0$，则将 $x$ 从 $\textit{cnt}$ 中删除。

```py [sol1-Python3]
class Solution:
    def findMatrix(self, nums: List[int]) -> List[List[int]]:
        ans = []
        cnt = Counter(nums)
        while cnt:
            ans.append(list(cnt))
            for x in ans[-1]:
                cnt[x] -= 1
                if cnt[x] == 0:
                    del cnt[x]
        return ans
```

```go [sol1-Go]
func findMatrix(nums []int) (ans [][]int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for len(cnt) > 0 {
		row := []int{}
		for x := range cnt {
			row = append(row, x)
			if cnt[x]--; cnt[x] == 0 {
				delete(cnt, x)
			}
		}
		ans = append(ans, row)
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
