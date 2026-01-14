水平栅栏和垂直栅栏分开计算。

- 对于水平栅栏，任意两个栅栏之间的距离（中间的栅栏全部删除）都可能是正方形的边长，存到一个哈希表 $\textit{hSet}$ 中。
- 对于垂直栅栏，任意两个栅栏之间的距离（中间的栅栏全部删除）都可能是正方形的边长，存到一个哈希表 $\textit{vSet}$ 中。

答案是 $\textit{hSet}$ 和 $\textit{vSet}$ 交集中的最大值的平方。记得返回之前取模。

如果交集为空，返回 $-1$。

```py [sol-Python3]
class Solution:
    def f(self, a: List[int], mx: int) -> Set[int]:
        a += [1, mx]
        a.sort()
        # 计算 a 中任意两个数的差，保存到哈希集合中
        return set(y - x for x, y in combinations(a, 2))

    def maximizeSquareArea(self, m: int, n: int, hFences: List[int], vFences: List[int]) -> int:
        MOD = 1_000_000_007
        h_set = self.f(hFences, m)
        v_set = self.f(vFences, n)

        ans = max(h_set & v_set, default=0)
        return ans * ans % MOD if ans else -1
```

```java [sol-Java]
class Solution {
    public int maximizeSquareArea(int m, int n, int[] hFences, int[] vFences) {
        final int MOD = 1_000_000_007;
        Set<Integer> hSet = f(hFences, m);
        Set<Integer> vSet = f(vFences, n);

        int ans = 0;
        for (int x : hSet) {
            if (vSet.contains(x)) {
                ans = Math.max(ans, x);
            }
        }
        return ans > 0 ? (int) ((long) ans * ans % MOD) : -1;
    }

    private Set<Integer> f(int[] a, int mx) {
        int n = a.length;
        a = Arrays.copyOf(a, n + 2);
        a[n++] = 1;
        a[n++] = mx;
        Arrays.sort(a);

        // 计算 a 中任意两个数的差，保存到哈希集合中
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                set.add(a[j] - a[i]);
            }
        }
        return set;
    }
}
```

```cpp [sol-C++]
class Solution {
    unordered_set<int> f(vector<int>& a, int mx) {
        a.push_back(1);
        a.push_back(mx);
        ranges::sort(a);

        // 计算 a 中任意两个数的差，保存到哈希集合中
        unordered_set<int> st;
        for (int i = 0; i < a.size(); i++) {
            for (int j = i + 1; j < a.size(); j++) {
                st.insert(a[j] - a[i]);
            }
        }
        return st;
    }

public:
    int maximizeSquareArea(int m, int n, vector<int>& hFences, vector<int>& vFences) {
        constexpr int MOD = 1'000'000'007;
        unordered_set<int> h_set = f(hFences, m);
        unordered_set<int> v_set = f(vFences, n);

        int ans = 0;
        for (int x : h_set) {
            if (v_set.contains(x)) {
                ans = max(ans, x);
            }
        }
        return ans ? 1LL * ans * ans % MOD : -1;
    }
};
```

```go [sol-Go]
func f(a []int, mx int) map[int]bool {
	a = append(a, 1, mx)
	slices.Sort(a)

	// 计算 a 中任意两个数的差，保存到哈希集合中
	set := map[int]bool{}
	for i, x := range a {
		for _, y := range a[i+1:] {
			set[y-x] = true
		}
	}
	return set
}

func maximizeSquareArea(m, n int, hFences, vFences []int) int {
	const mod = 1_000_000_007
	hSet := f(hFences, m)
	vSet := f(vFences, n)

	ans := 0
	for x := range hSet {
		if vSet[x] {
			ans = max(ans, x)
		}
	}

	if ans == 0 {
		return -1
	}
	return ans * ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(h^2+v^2)$，其中 $h$ 为 $\textit{hFences}$ 的长度，$v$ 为 $\textit{vFences}$ 的长度。
- 空间复杂度：$\mathcal{O}(h^2+v^2)$。

#### 相似题目

- [1465. 切割后面积最大的蛋糕](https://leetcode.cn/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/)
- [2943. 最大化网格图中正方形空洞的面积](https://leetcode.cn/problems/maximize-area-of-square-hole-in-grid/)

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
