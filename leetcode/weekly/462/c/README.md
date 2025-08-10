如果先选 $\textit{limit}[i]$ 大的价值，再选 $\textit{limit}[i]$ 小的价值，那么 $\textit{limit}[i]$ 小的价值可能没法选。

所以按照 $\textit{limit}[i]$ 从小到大选是最优的。

### 例一

比如 $\textit{limit}=[3,3,3,3,3]$。

当我们选了其中 $3$ 个数时：

- 已选的数永久变为非活跃状态，不算在「活跃元素个数」中。
- 剩余的 $2$ 个未选的 $\textit{limit}[i]$ 等于 $3$ 的数，也永久变为非活跃状态，**不能选**。
- 因此，贪心地，选这 $5$ 个数中，价值最大的 $3$ 个数。

### 例二

比如 $\textit{limit}=[1,2,3,3,3,3,3]$。

- 先选 $1$，选完 $1$ 之后，根据题意，$1$ 永久变为非活跃状态，不算在「活跃元素个数」中。现在活跃元素个数等于 $0$。
- 再选 $2$，现在活跃元素个数等于 $1$。
- 再选一个 $3$，现在活跃元素个数等于 $2$。根据题意，$2$ 永久变为非活跃状态，不算在「活跃元素个数」中。现在活跃元素个数等于 $1$，即我们刚才选的 $3$。
- 还可以再选 $2$ 个 $3$。
- 所以和例一一样，$3,3,3,3,3$ 中可以选 $3$ 个 $3$，和之前的 $1,2$ 怎么选是**无关**的！ 

### 思路

所以本题只需把 $\textit{limit}[i]$ 相同的价值分到同一组。同一组内取最大的 $\textit{limit}[i]$ 个价值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1QNbNzxEtZ/?t=18m26s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxTotal(self, value: List[int], limit: List[int]) -> int:
        groups = defaultdict(list)
        for lim, v in zip(limit, value):
            groups[lim].append(v)

        ans = 0
        for lim, a in groups.items():
            # 取最大的 lim 个数
            a.sort()
            ans += sum(a[-lim:])
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def maxTotal(self, value: List[int], limit: List[int]) -> int:
        groups = defaultdict(list)
        for lim, v in zip(limit, value):
            groups[lim].append(v)

        ans = 0
        for lim, a in groups.items():
            # 取最大的 lim 个数
            ans += sum(nlargest(lim, a))
        return ans
```

```java [sol-Java]
class Solution {
    public long maxTotal(int[] value, int[] limit) {
        int n = value.length;
        List<Integer>[] groups = new ArrayList[n + 1];
        Arrays.setAll(groups, _ -> new ArrayList<>());
        for (int i = 0; i < n; i++) {
            groups[limit[i]].add(value[i]);
        }

        long ans = 0;
        for (int lim = 1; lim <= n; lim++) {
            List<Integer> a = groups[lim];
            // 取最大的 lim 个数
            a.sort(Collections.reverseOrder());
            if (a.size() > lim) {
                a = a.subList(0, lim);
            }
            for (int x : a) {
                ans += x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxTotal(vector<int>& value, vector<int>& limit) {
        int n = value.size();
        vector<vector<int>> groups(n + 1);
        for (int i = 0; i < n; i++) {
            groups[limit[i]].push_back(value[i]);
        }

        long long ans = 0;
        for (int lim = 1; lim <= n; lim++) {
            auto& a = groups[lim];
            if (a.size() > lim) {
                // 取最大的 lim 个数
                ranges::nth_element(a, a.begin() + lim, greater());
                a.resize(lim);
            }
            ans += reduce(a.begin(), a.end(), 0LL);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxTotal(value, limit []int) (ans int64) {
	n := len(value)
	groups := make([][]int, n+1)
	for i, lim := range limit {
		groups[lim] = append(groups[lim], value[i])
	}
	for lim, a := range groups {
		// 取最大的 lim 个数
		slices.SortFunc(a, func(a, b int) int { return b - a })
		if len(a) > lim {
			a = a[:lim]
		}
		for _, x := range a {
			ans += int64(x)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{value}$ 的长度。如果用快速选择算法，可以做到 $\mathcal{O}(n)$，见 C++ 代码。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面贪心与思维题单的「**§5.2 脑筋急转弯**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
