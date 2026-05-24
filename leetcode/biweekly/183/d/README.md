本题有两个要求：

1. 子集点权和是 $k$ 的倍数（模 $k$ 为 $0$）。
2. 子集是**独立集**。见 [337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/)，[视频讲解【基础算法精讲 24】](https://www.bilibili.com/video/BV1vu4y1f7dn/)。

考虑到 $k$ 比较小，我们可以在 337 题的基础上，定义：

- $\textit{dfs}(x)_0[i]$ 表示不选节点 $x$ 时，子树 $x$ 的子集点权和模 $k$ 为 $i$ 的方案数。
- $\textit{dfs}(x)_1[i]$ 表示选节点 $x$ 时，子树 $x$ 的子集点权和模 $k$ 为 $i$ 的方案数。

设 $x$ 的儿子为 $y_1,y_2,\ldots,y_m$。定义：

- $f_{x,0}[c][i]$ 表示不选节点 $x$ 时，从 $x$ 的前 $c$ 个儿子中，选出的子集点权和模 $k$ 为 $i$ 的方案数。
- $f_{x,1}[c][i]$ 表示选节点 $x$ 时，从 $x$ 以及 $x$ 的前 $c$ 个儿子中，选出的子集点权和模 $k$ 为 $i$ 的方案数。

根据定义，$\textit{dfs}(x)_0[i] = f_{x,0}[m][i]$，$\textit{dfs}(x)_1[i] = f_{x,1}[m][i]$。

对于儿子 $y_c$ 来说：

- 如果选 $x$，那么 $y_c$ 一定不能选。在不选 $y_c$ 的情况下，枚举从子树 $y_c$ 中选出的子集点权和模 $k$ 为 $i=0,1,2,\ldots,k-1$，对应的方案数为 $\textit{dfs}(y)_0[i]$；然后枚举从 $x$ 以及子树 $y_1,y_2,\ldots,y_{c-1}$ 中，选出的子集点权和模 $k$ 为 $j$，对应的方案数为 $f_{x,1}[c-1][j]$，根据**乘法原理**，用**刷表法**把 $f_{x,1}[c][(i+j)\bmod k]$ 增加 $\textit{dfs}(y)_0[i]\cdot f_{x,1}[c-1][j]$。
- 如果不选 $x$，那么 $y_c$ 可选可不选，同理，把 $f_{x,0}[c][(i+j)\bmod k]$ 增加 $(\textit{dfs}(y)_0[i] + \textit{dfs}(y)_1[i]) \cdot f_{x,0}[c-1][j]$。

初始值 $f_{x,0}[0][0] = 1$，$f_{x,1}[0][\textit{nums}[x]\bmod k] = 1$。

代码实现时，计算 $f_{x,i}$ 可以用滚动数组优化。

为什么子集点权和可以在中途取模？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def countValidSubsets(self, parent: list[int], nums: list[int], k: int) -> int:
        MOD = 1_000_000_007
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parent[i]].append(i)

        def dfs(x: int) -> tuple[list[int], list[int]]:
            f0 = [0] * k  # f0[i] 表示不选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
            f1 = [0] * k  # f1[i] 表示选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
            f0[0] = 1
            f1[nums[x] % k] = 1

            for y in g[x]:
                fy0, fy1 = dfs(y)

                # 不选 x，那么 y 可选可不选
                nf0 = [0] * k
                for i in range(k):  # 枚举从子树 y 中选出的点权和模 k 为 i
                    v = fy0[i] + fy1[i]
                    if v == 0:  # 优化
                        continue
                    for j, w in enumerate(f0):  # 枚举从之前的子树中选出的点权和模 k 为 j
                        s = (i + j) % k
                        nf0[s] = (nf0[s] + v * w) % MOD

                # 选 x，那么 y 不能选
                nf1 = [0] * k
                for i, v in enumerate(fy0):  # 枚举从子树 y 中选出的点权和模 k 为 i
                    if v == 0:  # 优化
                        continue
                    for j, w in enumerate(f1):  # 枚举从 x 以及之前的子树中选出的点权和模 k 为 j
                        s = (i + j) % k
                        nf1[s] = (nf1[s] + v * w) % MOD

                f0, f1 = nf0, nf1

            return f0, f1

        f0, f1 = dfs(0)
        # 恰好被 k 整除即模 k 为 0，注意减去空集的方案数 1
        return (f0[0] + f1[0] - 1) % MOD
```

```py [sol-Python3 字典]
class Solution:
    def countValidSubsets(self, parent: list[int], nums: list[int], k: int) -> int:
        MOD = 1_000_000_007
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parent[i]].append(i)

        def dfs(x: int) -> tuple[dict[int, int], dict[int, int]]:
            f0 = {0: 1}  # 不选 x
            f1 = {nums[x] % k: 1}  # 选 x

            for y in g[x]:
                fy0, fy1 = dfs(y)

                nf0 = {}
                # 不选 x，那么 y 可选可不选
                for i in fy0.keys() | fy1.keys():
                    v = fy0.get(i, 0) + fy1.get(i, 0)
                    for j, w in f0.items():
                        s = (i + j) % k
                        nf0[s] = (nf0.get(s, 0) + v * w) % MOD

                nf1 = {}
                # 选 x，那么 y 不能选
                for i, v in fy0.items():
                    for j, w in f1.items():
                        s = (i + j) % k
                        nf1[s] = (nf1.get(s, 0) + v * w) % MOD

                f0, f1 = nf0, nf1

            return f0, f1

        f0, f1 = dfs(0)
        return (f0.get(0, 0) + f1.get(0, 0) - 1) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countValidSubsets(int[] parent, int[] nums, int k) {
        int n = parent.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[parent[i]].add(i);
        }

        long[][] res = dfs(0, g, nums, k);

        // 恰好被 k 整除即模 k 为 0，注意减去空集的方案数 1
        return (int) ((res[0][0] + res[1][0] - 1 + MOD) % MOD);
    }

    private long[][] dfs(int x, List<Integer>[] g, int[] nums, int k) {
        long[] f0 = new long[k]; // f0[i] 表示不选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
        long[] f1 = new long[k]; // f1[i] 表示选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
        f0[0] = 1;
        f1[nums[x] % k] = 1;

        for (int y : g[x]) {
            long[][] resY = dfs(y, g, nums, k);
            long[] fy0 = resY[0];
            long[] fy1 = resY[1];

            // 不选 x，那么 y 可选可不选
            long[] nf0 = new long[k];
            for (int i = 0; i < k; i++) { // 枚举从子树 y 中选出的点权和模 k 为 i
                long v = fy0[i] + fy1[i];
                if (v == 0) { // 优化
                    continue;
                }
                for (int j = 0; j < k; j++) { // 枚举从之前的子树中选出的点权和模 k 为 j
                    int s = (i + j) % k;
                    nf0[s] = (nf0[s] + v * f0[j]) % MOD;
                }
            }

            // 选 x，那么 y 不能选
            long[] nf1 = new long[k];
            for (int i = 0; i < k; i++) { // 枚举从子树 y 中选出的点权和模 k 为 i
                long v = fy0[i];
                if (v == 0) { // 优化
                    continue;
                }
                for (int j = 0; j < k; j++) { // 枚举从 x 以及之前的子树中选出的点权和模 k 为 j
                    int s = (i + j) % k;
                    nf1[s] = (nf1[s] + v * f1[j]) % MOD;
                }
            }

            f0 = nf0;
            f1 = nf1;
        }

        return new long[][]{f0, f1};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countValidSubsets(vector<int>& parent, vector<int>& nums, int k) {
        constexpr int MOD = 1'000'000'007;
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[parent[i]].push_back(i);
        }

        auto dfs = [&](this auto&& dfs, int x) -> pair<vector<long long>, vector<long long>> {
            vector<long long> f0(k); // f0[i] 表示不选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
            vector<long long> f1(k); // f1[i] 表示选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
            f0[0] = 1;
            f1[nums[x] % k] = 1;

            for (int y : g[x]) {
                auto [fy0, fy1] = dfs(y);

                // 不选 x，那么 y 可选可不选
                vector<long long> nf0(k);
                for (int i = 0; i < k; i++) { // 枚举从子树 y 中选出的点权和模 k 为 i
                    long long v = fy0[i] + fy1[i];
                    if (v == 0) { // 优化
                        continue;
                    }
                    for (int j = 0; j < k; j++) { // 枚举从之前的子树中选出的点权和模 k 为 j
                        int s = (i + j) % k;
                        nf0[s] = (nf0[s] + v * f0[j]) % MOD;
                    }
                }

                // 选 x，那么 y 不能选
                vector<long long> nf1(k);
                for (int i = 0; i < k; i++) { // 枚举从子树 y 中选出的点权和模 k 为 i
                    long long v = fy0[i];
                    if (v == 0) { // 优化
                        continue;
                    }
                    for (int j = 0; j < k; j++) { // 枚举从 x 以及之前的子树中选出的点权和模 k 为 j
                        int s = (i + j) % k;
                        nf1[s] = (nf1[s] + v * f1[j]) % MOD;
                    }
                }

                f0 = move(nf0);
                f1 = move(nf1);
            }

            return {f0, f1};
        };

        auto [f0, f1] = dfs(0);
        // 恰好被 k 整除即模 k 为 0，注意减去空集的方案数 1
        return (f0[0] + f1[0] - 1 + MOD) % MOD;
    }
};
```

```go [sol-Go]
func countValidSubsets(parent []int, nums []int, k int) int {
	const mod = 1_000_000_007
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	var dfs func(int) ([]int, []int)
	dfs = func(x int) ([]int, []int) {
		f0 := make([]int, k) // f0[i] 表示不选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
		f1 := make([]int, k) // f1[i] 表示选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
		f0[0] = 1
		f1[nums[x]%k] = 1

		for _, y := range g[x] {
			fy0, fy1 := dfs(y)

			// 不选 x，那么 y 可选可不选
			nf0 := make([]int, k)
			for i := range k { // 枚举从子树 y 中选出的点权和模 k 为 i
				v := fy0[i] + fy1[i]
				if v == 0 { // 优化
					continue
				}
				for j, w := range f0 { // 枚举从之前的子树中选出的点权和模 k 为 j
					s := (i + j) % k
					nf0[s] = (nf0[s] + v*w) % mod
				}
			}

			// 选 x，那么 y 不能选
			nf1 := make([]int, k)
			for i, v := range fy0 { // 枚举从子树 y 中选出的点权和模 k 为 i
				if v == 0 { // 优化
					continue
				}
				for j, w := range f1 { // 枚举从 x 以及之前的子树中选出的点权和模 k 为 j
					s := (i + j) % k
					nf1[s] = (nf1[s] + v*w) % mod
				}
			}
			f0, f1 = nf0, nf1
		}

		return f0, f1
	}

	f0, f1 := dfs(0)
	// 恰好被 k 整除即模 k 为 0，注意减去空集的方案数 1
	return (f0[0] + f1[0] - 1 + mod) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk^2)$，其中 $n$ 是 $\textit{parent}$ 的长度。有 $n-1$ 条边，每条边计算一次 $\mathcal{O}(k^2)$ 的 DP 转移。
- 空间复杂度：$\mathcal{O}(nk)$。最坏情况下，递归栈中保存了 $\mathcal{O}(n)$ 个大小为 $\mathcal{O}(k)$ 的 DP 数组。

## 专题训练

见下面动态规划题单的「**§12.2 树上最大独立集**」和「**§3.5 树上背包**」。

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
