### 本题视频讲解

见[【周赛 339】](https://www.bilibili.com/video/BV1va4y1M7Fr/)。

### 思路

为方便描述，将 $\textit{reward}$ 简记为 $r$。

比较两个物品 $i$ 和 $j$ 被吃后的得分：

- 如果 $i$ 给第一只老鼠，$j$ 给第二只老鼠，那么得分为 $r_1[i] + r_2[j]$；
- 如果 $i$ 给第二只老鼠，$j$ 给第一只老鼠，那么得分为 $r_2[i] + r_1[j]$；

如果第一种策略更优，则有 

$$
r_1[i] + r_2[j] > r_2[i] + r_1[j]
$$

变形得

$$
r_1[i] - r_2[i] > r_1[j] - r_2[j]
$$

这说明 $r_1[i] - r_2[i]$ 更大的奶酪，应该给第一只老鼠。

那么按照 $r_1[i] - r_2[i]$ 从大到小排序，前 $k$ 个给第一只老鼠，剩余的给第二只老鼠。

代码实现时，也可以先全部给第二只老鼠，然后再加上 $r_1[i] - r_2[i]$ 的前 $k$ 大之和。这可以用快速选择优化到 $O(n)$，具体见 C++ 代码。

```py [sol1-Python3]
class Solution:
    def miceAndCheese(self, reward1: List[int], reward2: List[int], k: int) -> int:
        a = sorted(zip(reward1, reward2), key=lambda p: p[1] - p[0])
        return sum(x for x, _ in a[:k]) + sum(y for _, y in a[k:])
```

```py [sol1-Python3 O(1) 空间]
class Solution:
    def miceAndCheese(self, r1: List[int], r2: List[int], k: int) -> int:
        for i, x in enumerate(r2):
            r1[i] -= x
        r1.sort(reverse=True)
        return sum(r2) + sum(r1[:k])  # 忽略切片空间
```

```java [sol1-Java]
class Solution {
    public int miceAndCheese(int[] r1, int[] r2, int k) {
        int ans = 0, n = r1.length;
        for (int i = 0; i < n; ++i) {
            ans += r2[i]; // 全部给第二只老鼠
            r1[i] -= r2[i];
        }
        Arrays.sort(r1);
        for (int i = n - k; i < n; ++i)
            ans += r1[i];
        return ans;
    }
}
```

```cpp [sol1-C++ 快速选择]
class Solution {
public:
    int miceAndCheese(vector<int> &r1, vector<int> &r2, int k) {
        for (int i = 0; i < r1.size(); ++i)
            r1[i] -= r2[i];
        nth_element(r1.begin(), r1.end() - k, r1.end());
        return accumulate(r2.begin(), r2.end(), 0) +
               accumulate(r1.end() - k, r1.end(), 0);
    }
};
```

```go [sol1-Go]
func miceAndCheese(reward1, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x // 全部给第二只老鼠
		reward1[i] -= x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, x := range reward1[:k] {
		ans += x
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$ 或 $O(n)$，其中 $n$ 为 $\textit{reward}_1$ 的长度。快速选择可以做到 $O(n)$，具体见 C++ 代码。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

### 相似题目

- [1029. 两地调度](https://leetcode.cn/problems/two-city-scheduling/)
