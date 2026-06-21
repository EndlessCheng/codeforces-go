先创建一个全为 $\texttt{#}$ 的网格，然后选任意一条从左上角到右下角的移动路径，把路径上的字符改成 $\texttt{.}$ 号，即可满足要求。

最简单的方式是，选一条先往右到右上角，再往下到右下角的路径（也可以先往下再往右）。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def createGrid(self, m: int, n: int) -> list[str]:
        return ['.' * n] + ['#' * (n - 1) + '.' for _ in range(m - 1)]
```

```java [sol-Java]
class Solution {
    public String[] createGrid(int m, int n) {
        String[] ans = new String[m];
        ans[0] = ".".repeat(n);
        for (int i = 1; i < m; i++) {
            ans[i] = "#".repeat(n - 1) + ".";
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> createGrid(int m, int n) {
        vector<string> ans(m);
        ans[0].resize(n, '.');
        for (int i = 1; i < m; i++) {
            ans[i].resize(n, '#');
            ans[i][n - 1] = '.';
        }
        return ans;
    }
};
```

```go [sol-Go]
func createGrid(m, n int) []string {
	ans := make([]string, m)
	ans[0] = strings.Repeat(".", n)
	for i := 1; i < m; i++ {
		ans[i] = strings.Repeat("#", n-1) + "."
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 专题训练

见下面思维题单的「**六、构造题**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
