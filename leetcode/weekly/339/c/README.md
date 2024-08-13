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

```py [sol-Python3]
class Solution:
    def miceAndCheese(self, r1: List[int], r2: List[int], k: int) -> int:
        a = sorted(zip(r1, r2), key=lambda p: p[1] - p[0])
        return sum(x for x, _ in a[:k]) + sum(y for _, y in a[k:])
```

```py [sol-Python3 原地修改]
class Solution:
    def miceAndCheese(self, r1: List[int], r2: List[int], k: int) -> int:
        for i, x in enumerate(r2):
            r1[i] -= x
        r1.sort(reverse=True)
        return sum(r2) + sum(r1[:k])  # 忽略切片空间
```

```java [sol-Java]
class Solution {
    public int miceAndCheese(int[] r1, int[] r2, int k) {
        int n = r1.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans += r2[i]; // 先全部给第二只老鼠
            r1[i] -= r2[i];
        }
        Arrays.sort(r1);
        for (int i = n - k; i < n; i++) {
            ans += r1[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    int miceAndCheese(vector<int>& r1, vector<int>& r2, int k) {
        for (int i = 0; i < r1.size(); i++) {
            r1[i] -= r2[i];
        }
        ranges::nth_element(r1, r1.end() - k);
        return reduce(r2.begin(), r2.end()) + // 先全部给第二只老鼠
               reduce(r1.end() - k, r1.end()); // 再加上增量
    }
};
```

```go [sol-Go]
func miceAndCheese(r1, r2 []int, k int) (ans int) {
	for i, x := range r2 {
		ans += x // 先全部给第二只老鼠
		r1[i] -= x
	}
	slices.SortFunc(r1, func(a, b int) int { return b - a })
	for _, x := range r1[:k] {
		ans += x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{reward}_1$ 的长度。**快速选择**可以做到 $\mathcal{O}(n)$，具体见 C++ 代码。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 相似题目

- [1029. 两地调度](https://leetcode.cn/problems/two-city-scheduling/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
