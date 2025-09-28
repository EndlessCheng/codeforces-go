用并查集把 $p_i$ 和 $q_i$ 合并。

合并后，同一集合中的元素可以随意交换。

贪心地：

- 交替和中取负号的元素位于奇数下标，我们把较小元素交换到奇数下标上。
- 交替和中取正号的元素位于偶数下标，我们把较大元素交换到偶数下标上。

为此，并查集需要额外维护集合中的奇数下标个数。

对于同一组，设奇数下标个数为 $\textit{odd}$。把元素从小到大排序后，前 $\textit{odd}$ 个元素取负号，其余元素取正号。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
class UnionFind:
    def __init__(self, n: int):
        # 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        # 集合 i 的代表元是自己
        self._fa = list(range(n))  # 代表元
        self.odd = [i % 2 for i in range(n)]  # 集合中的奇数个数
        self.cc = n  # 连通块个数

    # 返回 x 所在集合的代表元
    # 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    def find(self, x: int) -> int:
        # 如果 fa[x] == x，则表示 x 是代表元
        fa = self._fa
        if fa[x] != x:
            fa[x] = self.find(fa[x])  # fa 改成代表元
        return fa[x]

    # 把 from 所在集合合并到 to 所在集合中
    def merge(self, from_: int, to: int) -> None:
        x, y = self.find(from_), self.find(to)
        if x == y:  # from 和 to 在同一个集合，不做合并
            return
        self._fa[x] = y  # 合并集合
        self.odd[y] += self.odd[x]  # 更新集合中的奇数个数

class Solution:
    def maxAlternatingSum(self, nums: List[int], swaps: List[List[int]]) -> int:
        uf = UnionFind(len(nums))
        for p, q in swaps:
            uf.merge(p, q)

        g = defaultdict(list)
        for i, x in enumerate(nums):
            g[uf.find(i)].append(x)  # 相同集合的元素分到同一组

        ans = 0
        for i, a in g.items():
            a.sort()
            odd = uf.odd[i]
            # 小的取负号，大的取正号
            for j, x in enumerate(a):
                ans += -x if j < odd else x
        return ans
```

```java [sol-Java]
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
class UnionFind {
    private final int[] fa; // 代表元
    public final int[] odd; // 集合中的奇数个数

    UnionFind(int n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己
        fa = new int[n];
        odd = new int[n];
        for (int i = 0; i < n; i++) {
            fa[i] = i;
            odd[i] = i % 2;
        }
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    public int find(int x) {
        // 如果 fa[x] == x，则表示 x 是代表元
        if (fa[x] != x) {
            fa[x] = find(fa[x]); // fa 改成代表元
        }
        return fa[x];
    }

    // 把 from 所在集合合并到 to 所在集合中
    public void merge(int from, int to) {
        int x = find(from);
        int y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return;
        }
        fa[x] = y; // 合并集合
        odd[y] += odd[x]; // 更新集合中的奇数个数
    }
}

class Solution {
    public long maxAlternatingSum(int[] nums, int[][] swaps) {
        int n = nums.length;
        UnionFind uf = new UnionFind(n);
        for (int[] p : swaps) {
            uf.merge(p[0], p[1]);
        }

        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int i = 0; i < n; i++) {
            g[uf.find(i)].add(nums[i]); // 相同集合的元素分到同一组
        }

        long ans = 0;
        for (int i = 0; i < n; i++) {
            List<Integer> a = g[i];
            if (a.isEmpty()) {
                continue;
            }
            Collections.sort(a);
            int odd = uf.odd[i];
            // 小的取负号，大的取正号
            for (int j = 0; j < a.size(); j++) {
                int x = a.get(j);
                ans += j < odd ? -x : x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
class UnionFind {
    vector<int> fa; // 代表元
public:
    vector<int> odd; // 集合中的奇数个数

    UnionFind(int n) : fa(n), odd(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己
        for (int i = 0; i < n; i++) {
            fa[i] = i;
            odd[i] = i % 2;
        }
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    int find(int x) {
        // 如果 fa[x] == x，则表示 x 是代表元
        if (fa[x] != x) {
            fa[x] = find(fa[x]); // fa 改成代表元
        }
        return fa[x];
    }

    // 把 from 所在集合合并到 to 所在集合中
    // 返回是否合并成功
    void merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        odd[y] += odd[x]; // 更新集合中的奇数个数
    }
};

class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums, vector<vector<int>>& swaps) {
        int n = nums.size();
        UnionFind uf(n); // 假设并查集已实现
        for (auto& p : swaps) {
            uf.merge(p[0], p[1]);
        }

        vector<vector<int>> g(n);
        for (int i = 0; i < n; i++) {
            g[uf.find(i)].push_back(nums[i]); // 相同集合的元素分到同一组
        }

        long long ans = 0;
        for (int i = 0; i < n; i++) {
            auto& a = g[i];
            ranges::sort(a);
            int odd = uf.odd[i];
            // 小的取负号，大的取正号
            for (int j = 0; j < a.size(); j++) {
                int x = a[j];
                ans += j < odd ? -x : x;
            }
        }
        return ans;
    }
};
```

```cpp [sol-C++ 快速选择]
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
class UnionFind {
    vector<int> fa; // 代表元
public:
    vector<int> odd; // 集合中的奇数个数

    UnionFind(int n) : fa(n), odd(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己
        for (int i = 0; i < n; i++) {
            fa[i] = i;
            odd[i] = i % 2;
        }
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    int find(int x) {
        // 如果 fa[x] == x，则表示 x 是代表元
        if (fa[x] != x) {
            fa[x] = find(fa[x]); // fa 改成代表元
        }
        return fa[x];
    }

    // 把 from 所在集合合并到 to 所在集合中
    // 返回是否合并成功
    void merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        odd[y] += odd[x]; // 更新集合中的奇数个数
    }
};

class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums, vector<vector<int>>& swaps) {
        int n = nums.size();
        UnionFind uf(n); // 假设并查集已实现
        for (auto& p : swaps) {
            uf.merge(p[0], p[1]);
        }

        vector<vector<int>> g(n);
        for (int i = 0; i < n; i++) {
            g[uf.find(i)].push_back(nums[i]); // 相同集合的元素分到同一组
        }

        long long ans = 0;
        for (int i = 0; i < n; i++) {
            auto& a = g[i];
            if (a.empty()) {
                continue;
            }
            int odd = uf.odd[i];
            ranges::nth_element(a, a.begin() + odd);
            // 小的取负号，大的取正号
            for (int j = 0; j < a.size(); j++) {
                int x = a[j];
                ans += j < odd ? -x : x;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
type unionFind struct {
	fa  []int // 代表元
	odd []int // 集合中的奇数个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	odd := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
		odd[i] = i % 2
	}
	return unionFind{fa, odd}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	// 如果 fa[x] == x，则表示 x 是代表元
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
	}
	return u.fa[x]
}

// 把 from 所在集合合并到 to 所在集合中
func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return
	}
	u.fa[x] = y          // 合并集合
	u.odd[y] += u.odd[x] // 更新集合中的奇数个数
}

func maxAlternatingSum(nums []int, swaps [][]int) (ans int64) {
	n := len(nums)
	uf := newUnionFind(n)
	for _, p := range swaps {
		uf.merge(p[0], p[1])
	}

	g := make([][]int, n)
	for i, x := range nums {
		f := uf.find(i)
		g[f] = append(g[f], x) // 相同集合的元素分到同一组
	}

	for i, a := range g {
		if a == nil {
			continue
		}
		slices.Sort(a)
		odd := uf.odd[i]
		// 小的取负号，大的取正号
		for j, x := range a {
			if j < odd {
				ans -= int64(x)
			} else {
				ans += int64(x)
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。使用建图 + DFS + 快速选择算法，可以把时间复杂度优化到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 图论题单的「**§1.1 深度优先搜索（DFS）**」。
2. 数据结构题单的「**七、并查集**」。

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
