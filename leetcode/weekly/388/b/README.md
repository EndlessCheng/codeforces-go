请看 [视频讲解](https://www.bilibili.com/video/BV1Xr421J77b/) 第二题。

首先，应当选 $\textit{happiness}$ 中最大的 $k$ 的数。

这些数要按照什么顺序选呢？

由于小的数减成 $0$ 就不再减少了，优先选大的更好。

比如 $2,1,1$，如果按照 $1,1,2$ 的顺序选，答案为 $1+0+0=1$。但按照 $2,1,1$ 的顺序选，答案为 $2+0+0=2$ 更优。

```py [sol-Python3]
class Solution:
    def maximumHappinessSum(self, happiness: List[int], k: int) -> int:
        happiness.sort(reverse=True)
        ans = 0
        for i, x in enumerate(happiness[:k]):
            if x <= i:
                break
            ans += x - i
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumHappinessSum(int[] happiness, int k) {
        Arrays.sort(happiness);
        int n = happiness.length;
        long ans = 0;
        for (int i = n - 1; i >= n - k && happiness[i] > n - 1 - i; i--) {
            ans += happiness[i] - (n - 1 - i);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumHappinessSum(vector<int> &happiness, int k) {
        ranges::sort(happiness, greater<>());
        long long ans = 0;
        for (int i = 0; i < k && happiness[i] > i; i++) {
            ans += happiness[i] - i;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumHappinessSum(happiness []int, k int) (ans int64) {
	slices.SortFunc(happiness, func(a, b int) int { return b - a })
	for i, x := range happiness[:k] {
		if x <= i {
			break
		}
		ans += int64(x - i)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{happiness}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序和切片的开销。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
