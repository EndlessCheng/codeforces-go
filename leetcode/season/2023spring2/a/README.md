### 本题视频讲解

见[【力扣杯2023春·战队赛】](https://www.bilibili.com/video/TODO/)。

### 思路

排序后，统计相邻相差不超过 $1$ 的元素个数：

- 如果相邻相差大于 $1$，那么重新统计。
- 否则计数器加一，更新答案的最大值。

```py [sol1-Python3]
class Solution:
    def runeReserve(self, runes: List[int]) -> int:
        runes.sort()
        ans = cnt = 1
        for pre, cur in pairwise(runes):
            if cur - pre > 1:
                cnt = 1  # 重新统计
            else:
                cnt += 1
                ans = max(ans, cnt)
        return ans
```

```java [sol1-Java]
class Solution {
    public int runeReserve(int[] runes) {
        Arrays.sort(runes);
        int ans = 1, cnt = 1;
        for (int i = 1; i < runes.length; i++)
            if (runes[i] - runes[i - 1] > 1)
                cnt = 1;
            else
                ans = Math.max(ans, ++cnt);
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int runeReserve(vector<int> &runes) {
        sort(runes.begin(), runes.end());
        int ans = 1, cnt = 1;
        for (int i = 1; i < runes.size(); i++)
            if (runes[i] - runes[i - 1] > 1)
                cnt = 1;
            else
                ans = max(ans, ++cnt);
        return ans;
    }
};
```

```go [sol1-Go]
func runeReserve(runes []int) int {
	sort.Ints(runes)
	ans, cnt := 1, 1
	for i, n := 1, len(runes); i < n; i++ {
		if runes[i]-runes[i-1] > 1 {
			cnt = 1 // 重新统计
		} else if cnt++; cnt > ans {
			ans = cnt
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{runes}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序使用的栈空间，仅用到若干额外变量。
