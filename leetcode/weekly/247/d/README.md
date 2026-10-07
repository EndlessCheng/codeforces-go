按照构筑房间的顺序，依次给每个节点标记整数 $1,2,3,\ldots,n$。第一个构筑的房间（根节点）标记 $1$，第二个构筑的房间标记 $2$，依此类推。

DFS 遍历这棵树，收集标记的数，我们会得到一个 $1$ 到 $n$ 的排列。所有排列一共有 $n!$ 种，但其中肯定有不合法的（无法遍历得到的），如何去掉这些不合法的排列？

例如，房间 $0$ 是树的根，我们收集到的第一个数一定是 $1$，那么所有不以 $1$ 开头的排列都是不合法的。以 $1$ 开头的排列个数是 $(n-1)!$，相当于把 $n!$ 除以 $n$。

又例如，某棵子树的标记为 $2,3,5$，子树的根为 $2$，那么所有排列中，只有 $2$ 在 $3,5$ 左边的排列才合法。包含 $2,3,5$ 的排列可以分为三类：

- $2$ 在 $3,5$ 左边的排列，合法。
- $3$ 在 $2,5$ 左边的排列，不合法。
- $5$ 在 $2,3$ 左边的排列，不合法。

所以，把方案数除以 $3$，就得到了 $2$ 一定在 $3,5$ 左边的排列个数。

设 $\textit{size}[i]$ 是子树 $i$ 的大小（节点个数）。一般地，对于每棵子树 $i$，$i$ 的标记必须在该子树其余标记的左边。把方案数除以 $\textit{size}[i]$，就得到了 $i$ 的标记在该子树其余标记的左边的排列个数。特别地，把方案数除以 $\textit{size}[0] = n$，就得到了根节点 $0$ 的标记在其余标记的左边的排列个数。

所以最终答案为

$$
\dfrac{n!}{\prod\limits_{i=0}^{n-1} \textit{size}[i]}
$$

分别计算分子和分母，然后用费马小定理计算分母的倒数（逆元）。原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/discuss/post/3584387/fen-xiang-gun-mo-yun-suan-de-shi-jie-dan-7xgu/)。

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

        def dfs(x: int) -> int:
            size = 1
            for y in g[x]:
                size += dfs(y)
            nonlocal prod
            prod = prod * size % MOD
            return size

        prod = 1  # 分母
        dfs(0)
        return fac * pow(prod, -1, MOD) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    private long prod = 1; // 分母

    public int waysToBuildRooms(int[] prevRoom) {
        int n = prevRoom.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        long fac = 1; // 分子
        for (int i = 1; i < n; i++) {
            fac = fac * (i + 1) % MOD;
            g[prevRoom[i]].add(i);
        }
        dfs(0, g);
        return (int) (fac * pow(prod, MOD - 2) % MOD);
    }

    private int dfs(int x, List<Integer>[] g) {
        int size = 1;
        for (int y : g[x]) {
            size += dfs(y, g);
        }
        prod = prod * size % MOD;
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
    static constexpr int MOD = 1'000'000'007;

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

        long long prod = 1; // 分母

        auto dfs = [&](this auto&& dfs, int x) -> int {
            int size = 1;
            for (int y : g[x]) {
                size += dfs(y);
            }
            prod = prod * size % MOD;
            return size;
        };

        dfs(0);
        return fac * qpow(prod, MOD - 2) % MOD;
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

	prod := 1 // 分母

	var dfs func(int) int
	dfs = func(x int) int {
		size := 1
		for _, y := range g[x] {
			size += dfs(y)
		}
		prod = prod * size % mod
		return size
	}

	dfs(0)
	return fac * pow(prod, mod-2) % mod
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

## 进阶问题

把树改成无向树。

返回一个长为 $n$ 的数组，分别表示以节点 $0,1,2,\ldots,n-1$ 为根时的答案。

这题是 [ABC160F. Distributing Integers](https://atcoder.jp/contests/abc160/tasks/abc160_f)。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
