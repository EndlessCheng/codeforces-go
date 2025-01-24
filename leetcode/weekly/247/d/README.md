按照构筑房间的顺序，给每个节点标记上 $1,2,3,\ldots,n$。第一个构筑的房间标记 $1$，第二个构筑的房间标记 $2$，依此类推。

然后 DFS 遍历这棵树，收集标记的数字，我们会得到一个 $1$ 到 $n$ 的排列。所有排列总共有 $n!$ 种，但其中肯定有不合法的（无法遍历得到的）。

比如，这个排列肯定要以 $1$ 开头，所有不以 $1$ 开头的排列都是不合法的。以 $1$ 开头的排列个数是 $(n-1)!$，相当于把 $n!$ 除以 $n$。

同理，对于每棵子树 $i$（对应着 $1$ 到 $n$ 排列中的一个**子序列**），不以 $i$ 的标记数字开头的排列都是不合法的，所以同样地，要把方案数除以 $\textit{size}[i]$，即子树 $i$ 的大小。

所以最终答案为

$$
\dfrac{n!}{\prod\limits_{i=0}^{n-1} \textit{size}[i]}
$$

分别计算分子和分母，然后用费马小定理计算分母的倒数（逆元）。原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def waysToBuildRooms(self, prevRoom: List[int]) -> int:
        MOD = 1_000_000_007
        n = len(prevRoom)
        g = [[] for _ in range(n)]
        fac = 1  # 分子
        for i in range(1, n):
            fac = fac * (i + 1) % MOD
            g[prevRoom[i]].append(i)

        mul = 1  # 分母
        def dfs(x: int) -> int:
            size = 1
            for y in g[x]:
                size += dfs(y)
            nonlocal mul
            mul = mul * size % MOD
            return size
        dfs(0)

        return fac * pow(mul, -1, MOD) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    private long mul = 1; // 分母

    public int waysToBuildRooms(int[] prevRoom) {
        int n = prevRoom.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        long fac = 1; // 分子
        for (int i = 1; i < n; i++) {
            fac = fac * (i + 1) % MOD;
            g[prevRoom[i]].add(i);
        }
        dfs(0, g);
        return (int) (fac * pow(mul, MOD - 2) % MOD);
    }

    private int dfs(int x, List<Integer>[] g) {
        int size = 1;
        for (int y : g[x]) {
            size += dfs(y, g);
        }
        mul = mul * size % MOD;
        return size;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long qpow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    int waysToBuildRooms(vector<int>& prevRoom) {
        int n = prevRoom.size();
        vector<vector<int>> g(n);
        long long fac = 1; // 分子
        for (int i = 1; i < n; i++) {
            fac = fac * (i + 1) % MOD;
            g[prevRoom[i]].push_back(i);
        }

        long long mul = 1; // 分母
        auto dfs = [&](this auto&& dfs, int x) -> int {
            int size = 1;
            for (int y : g[x]) {
                size += dfs(y);
            }
            mul = mul * size % MOD;
            return size;
        };
        dfs(0);

        return fac * qpow(mul, MOD - 2) % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	g := make([][]int, n)
	fac := 1 // 分子
	for i := 1; i < n; i++ {
		p := prevRoom[i]
		g[p] = append(g[p], i)
		fac = fac * (i + 1) % mod
	}

	mul := 1 // 分母
	var dfs func(int) int
	dfs = func(x int) int {
		size := 1
		for _, y := range g[x] {
			size += dfs(y)
		}
		mul = mul * size % mod
		return size
	}
	dfs(0)

	return fac * pow(mul, mod-2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\log M)$，其中 $n$ 是 $\textit{prevRoom}$ 的长度，$M=10^9+7$。
- 空间复杂度：$\mathcal{O}(n)$。

## 变形题

把树改成无向树。

返回一个长为 $n$ 的数组，分别表示以节点 $0,1,2,\ldots,n-1$ 为根时的答案。

这题是 [ABC160F. Distributing Integers](https://atcoder.jp/contests/abc160/tasks/abc160_f)。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
