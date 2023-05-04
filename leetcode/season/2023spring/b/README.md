### 本题视频讲解

见[【力扣杯2023春·个人赛】](https://www.bilibili.com/video/BV1dg4y1j78A/)第二题。

```py [sol1-Python3]
class Solution:
    def adventureCamp(self, a: List[str]) -> int:
        vis = set(a[0].split('->'))  # 这样可能会把空串插入，但是没有关系
        max_cnt, ans = 0, -1
        for i in range(1, len(a)):
            if a[i] == "": continue
            cnt = 0
            for t in a[i].split('->'):
                if t not in vis:
                    vis.add(t)
                    cnt += 1
            if cnt > max_cnt:
                max_cnt, ans = cnt, i
        return ans
```

```go [sol1-Go]
func adventureCamp(a []string) int {
	vis := map[string]bool{}
	for _, s := range strings.Split(a[0], "->") { // 这样可能会把空串插入，但是没有关系
		vis[s] = true
	}
	maxCnt, ans := 0, -1
	for i := 1; i < len(a); i++ {
		if a[i] == "" {
			continue
		}
		cnt := 0
		for _, s := range strings.Split(a[i], "->") {
			if !vis[s] {
				vis[s] = true
				cnt++
			}
		}
		if cnt > maxCnt {
			maxCnt, ans = cnt, i
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(L)$。
