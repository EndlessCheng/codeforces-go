定义 $e_p(n)$ 为 $n$ 的质因子分解中质数 $p$ 的指数。

对于 $x\cdot y = k^2$，我们有

$$
e_p(x) + e_p(y) = e_p(k^2) = 2\cdot e_p(k)
$$

所以 $e_p(x) + e_p(y)$ 是偶数，由于只有奇+奇=偶或者偶+偶=偶，所以 $e_p(x)$ 和 $e_p(y)$ 的奇偶性相同，即

$$
e_p(x)\bmod 2 = e_p(y)\bmod 2
$$

根据这一等式，定义**无平方因子核** $\text{core}(n)$ 为 $n$ 除去完全平方因子后的剩余结果。

例如 $\text{core}(8)=8/4=2,\ \text{core}(12)=12/4=3, \text{core}(25)=25/25=1, \text{core}(5)=5/1=5$。

根据定义，我们有

$$
e_p(\text{core}(n)) = e_p(n)\bmod 2
$$

设 $x' = \text{core}(x)$，$y' = \text{core}(y)$，等式简化为

$$
e_p(x') = e_p(y')
$$

这意味着 $x'$ 和 $y'$ 完全一样，即

$$
\text{core}(x) = \text{core}(y)
$$

例如 $8\cdot 2 = 16$，除去完全平方因子后是 $2\cdot 2 = 4$。

例如 $4\cdot 9 = 36$，除去完全平方因子后是 $1\cdot 1 = 1$。

现在问题变成：

- 统计点权 core 值相同的点对个数，其中点对必须互为祖先后代。

做法类似 [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/)。本题是在树上做这个问题，递归返回时要恢复现场，即撤销计数器 $\textit{cnt}[\text{core}(\textit{nums}[x])]$ 的加一。

预处理 $\text{core}(n)$ 的过程类似埃氏筛，请看 [我的题解](https://leetcode.cn/problems/maximum-element-sum-of-a-complete-subset-of-indices/solutions/2446037/an-zhao-corei-fen-zu-pythonjavacgo-by-en-i6nu/)。

[本题视频讲解](https://www.bilibili.com/video/BV1FJ4uz1EkN/?t=26m51s)，欢迎点赞关注~

```py [sol-Python3]
# 预处理 core
MX = 100_001
core = [0] * MX
for i in range(1, MX):
    if core[i] == 0:
        for j in range(1, isqrt(MX // i) + 1):
            core[i * j * j] = i

class Solution:
    def sumOfAncestors(self, n: int, edges: List[List[int]], nums: List[int]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        cnt = defaultdict(int)
        ans = 0

        def dfs(x: int, fa: int) -> None:
            nonlocal ans
            c = core[nums[x]]
            # 本题 x 的祖先不包括 x 自己
            ans += cnt[c]
            cnt[c] += 1
            for y in g[x]:
                if y != fa:
                    dfs(y, x)
            cnt[c] -= 1  # 恢复现场

        dfs(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MX = 100_001;
    private static final int[] core = new int[MX];

    static {
        // 预处理 core
        for (int i = 1; i < MX; i++) {
            if (core[i] == 0) {
                for (int j = 1; i * j * j < MX; j++) {
                    core[i * j * j] = i;
                }
            }
        }
    }

    public long sumOfAncestors(int n, int[][] edges, int[] nums) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        Map<Integer, Integer> cnt = new HashMap<>(); // 更快的写法见【Java 数组】
        return dfs(0, -1, g, nums, cnt);
    }

    private long dfs(int x, int fa, List<Integer>[] g, int[] nums, Map<Integer, Integer> cnt) {
        int cr = core[nums[x]];
        // 本题 x 的祖先不包括 x 自己
        int c = cnt.getOrDefault(cr, 0);
        long res = c;
        cnt.put(cr, c + 1);
        for (int y : g[x]) {
            if (y != fa) {
                res += dfs(y, x, g, nums, cnt);
            }
        }
        cnt.put(cr, c); // 恢复现场
        return res;
    }
}
```

```java [sol-Java 数组]
class Solution {
    private static final int MX = 100_001;
    private static final int[] core = new int[MX];

    static {
        // 预处理 core
        for (int i = 1; i < MX; i++) {
            if (core[i] == 0) {
                for (int j = 1; i * j * j < MX; j++) {
                    core[i * j * j] = i;
                }
            }
        }
    }

    public long sumOfAncestors(int n, int[][] edges, int[] nums) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, core[x]);
        }

        int[] cnt = new int[mx + 1];
        return dfs(0, -1, g, nums, cnt);
    }

    private long dfs(int x, int fa, List<Integer>[] g, int[] nums, int[] cnt) {
        int cr = core[nums[x]];
        // 本题 x 的祖先不包括 x 自己
        long res = cnt[cr]++;
        for (int y : g[x]) {
            if (y != fa) {
                res += dfs(y, x, g, nums, cnt);
            }
        }
        cnt[cr]--; // 恢复现场
        return res;
    }
}
```

```cpp [sol-C++]
// 预处理 core
const int MX = 100'001;
int core[MX];

int init = [] {
    for (int i = 1; i < MX; ++i) {
        if (core[i] == 0) {
            for (int j = 1; i * j * j < MX; ++j) {
                core[i * j * j] = i;
            }
        }
    }
    return 0;
}();

class Solution {
public:
    long long sumOfAncestors(int n, vector<vector<int>>& edges, vector<int>& nums) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        unordered_map<int, int> cnt;
        long long ans = 0;

        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            int c = core[nums[x]];
            // 本题 x 的祖先不包括 x 自己
            ans += cnt[c];
            cnt[c]++;
            for (int y : g[x]) {
                if (y != fa) {
                    dfs(y, x);
                }
            }
            cnt[c]--; // 恢复现场
        };

        dfs(0, -1);
        return ans;
    }
};
```

```go [sol-Go]
// 预处理 core
const mx = 100_001
var core = [mx]int{}

func init() {
	for i := 1; i < mx; i++ {
		if core[i] == 0 {
			for j := 1; i*j*j < mx; j++ {
				core[i*j*j] = i
			}
		}
	}
}

func sumOfAncestors(n int, edges [][]int, nums []int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := map[int]int{}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		c := core[nums[x]]
		// 本题 x 的祖先不包括 x 自己
		ans += int64(cnt[c])
		cnt[c]++
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}
		cnt[c]-- // 恢复现场
	}
	dfs(0, -1)
	return
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

[2862. 完全子集的最大元素和](https://leetcode.cn/problems/maximum-element-sum-of-a-complete-subset-of-indices/)

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
