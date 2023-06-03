## 思考

如果 $k=1$，应该如何选择呢？（思考问题可以先从一些简单的情况开始）

不妨先把奶酪全部给第二只老鼠，然后「撤销」其中的一块奶酪，给第一只老鼠。如何选择可以使得分最大？

你可以把这个结论推广到 $k>1$ 的情况吗？

## 解惑

为方便描述，将 $\textit{reward}$ 简记为 $r$。

先把奶酪全部给第二只老鼠，然后撤销其中的第 $i$ 块奶酪，给第一只老鼠，那么得分增加了 

$$
r_1[i] - r_2[i]
$$ 

在 $k=1$ 时，选上式最大的奶酪，给第一只老鼠，这样可以使得分最大。（注意第一只老鼠一定要吃**恰好** $k$ 块奶酪）

对于 $k>1$ 的情况，可以按照 $r_1[i] - r_2[i]$ 从大到小排序，把得分加上 $r_1[i] - r_2[i]$ 的前 $k$ 大之和。这可以用快速选择优化到 $\mathcal{O}(n)$，具体见 C++ 代码。

```py [sol1-Python3]
class Solution:
    def miceAndCheese(self, reward1: List[int], reward2: List[int], k: int) -> int:
        a = sorted(zip(reward1, reward2), key=lambda p: p[1] - p[0])
        return sum(x for x, _ in a[:k]) + sum(y for _, y in a[k:])
```

```py [sol1-Python3 原地修改]
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
            ans += r2[i]; // 先全部给第二只老鼠
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
        return accumulate(r2.begin(), r2.end(), 0) + // 先全部给第二只老鼠
               accumulate(r1.end() - k, r1.end(), 0); // 再加上增量
    }
};
```

```go [sol1-Go]
func miceAndCheese(reward1, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x // 先全部给第二只老鼠
		reward1[i] -= x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, x := range reward1[:k] {
		ans += x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{reward}_1$ 的长度。快速选择可以做到 $\mathcal{O}(n)$，具体见 C++ 代码。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

#### 相似题目

- [1029. 两地调度](https://leetcode.cn/problems/two-city-scheduling/)

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

---

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
