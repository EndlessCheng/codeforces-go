例如：

- 排序前 $[2,3,1]$。
- 排序后 $[1,2,3]$。

从左到右遍历排序前的数组，依次把元素交换到它排序后的位置：

- $2$ 和 $3$ 交换，得到 $[3,2,1]$。
- $3$ 和 $1$ 交换，得到 $[1,2,3]$。
- $1$ 已经在正确的位置上了。为什么？因为其他 $2$ 个数都在正确位置上了，不可能存在其他目标位置是 $1$，所以 $1$ 一定无需交换。
- 所以这 $3$ 个数需要 $2$ 次交换操作。

在「排序前的元素下标」与「排序后的元素下标」之间连边，得到一个图。

> 为什么是下标连边而不是元素值连边？因为元素值和下标是一一对应的，下标的范围更小，可以用数组记录。

通过上面的例子可知，图中每个大小为 $k$ 连通块（有 $k$ 个点和 $k$ 条边，是个环），需要 $k-1$ 次交换操作，加到答案中。

为什么是 $k-1$ 次？因为交换的过程中，除了最后一个数，其余每个数都不会等于其目标位置。这可以用反证法证明，如果在中途就出现某个数在其目标位置（这个数无需交换）的情况，那么意味着我们找到了一个更小的环，这与实际情况矛盾。视频讲解中给出了基于**置换**的解释，可以看看。

也可以用 $n$ 减去连通块的个数（每个连通块可以少操作一次），即为答案。

计算连通块个数，可以用 DFS、BFS、并查集等，下面用的并查集。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Z3JGzwEU9/?t=3m37s)，欢迎点赞关注~

```py [sol-Python3]
# 完整的并查集模板，见我的数据结构题单
class UnionFind:
    def __init__(self, n: int):
        # 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        # 集合 i 的代表元是自己，大小为 1
        self._fa = list(range(n))  # 代表元
        self.cc = n  # 连通块个数

    # 返回 x 所在集合的代表元
    # 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    def find(self, x: int) -> int:
        # 如果 fa[x] == x，则表示 x 是代表元
        if self._fa[x] != x:
            self._fa[x] = self.find(self._fa[x])  # fa 改成代表元
        return self._fa[x]

    # 把 from 所在集合合并到 to 所在集合中
    def merge(self, from_: int, to: int) -> None:
        x, y = self.find(from_), self.find(to)
        if x == y:  # from 和 to 在同一个集合，不做合并
            return
        self._fa[x] = y  # 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        self.cc -= 1  # 成功合并，连通块个数减一

class Solution:
    def minSwaps(self, nums: List[int]) -> int:
        a = sorted((sum(map(int, str(x))), x, i) for i, x in enumerate(nums))
        n = len(a)
        u = UnionFind(n)
        for i, t in enumerate(a):
            u.merge(i, t[2])
        return n - u.cc
```

```java [sol-Java]
// 完整的并查集模板，见我的数据结构题单
class UnionFind {
    private final int[] fa; // 代表元
    public int cc; // 连通块个数

    UnionFind(int n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己，大小为 1
        fa = new int[n];
        for (int i = 0; i < n; i++) {
            fa[i] = i;
        }
        cc = n;
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
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
    }
}

class Solution {
    public int minSwaps(int[] nums) {
        int n = nums.length;
        int[][] a = new int[n][3];
        for (int i = 0; i < n; i++) {
            int s = 0;
            for (int x = nums[i]; x > 0; x /= 10) {
                s += x % 10;
            }
            a[i][0] = s;
            a[i][1] = nums[i];
            a[i][2] = i;
        }

        Arrays.sort(a, (p, q) -> p[0] != q[0] ? p[0] - q[0] : p[1] - q[1]);

        UnionFind u = new UnionFind(n);
        for (int i = 0; i < n; i++) {
            u.merge(i, a[i][2]);
        }
        return n - u.cc;
    }
}
```

```cpp [sol-C++]
// 完整的并查集模板，见我的数据结构题单
class UnionFind {
    vector<int> fa; // 代表元

public:
    int cc; // 连通块个数

    UnionFind(int n) : fa(n), cc(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己，大小为 1
        ranges::iota(fa, 0); // iota(fa.begin(), fa.end(), 0);
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
    void merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
    }
};

class Solution {
public:
    int minSwaps(vector<int>& nums) {
        int n = nums.size();
        vector<tuple<int, int, int>> a(n);
        for (int i = 0; i < n; i++) {
            int s = 0;
            for (int x = nums[i]; x > 0; x /= 10) {
                s += x % 10;
            }
            a[i] = {s, nums[i], i};
        }

        ranges::sort(a);

        UnionFind u(n);
        for (int i = 0; i < n; i++) {
            u.merge(i, get<2>(a[i]));
        }
        return n - u.cc;
    }
};
```

```go [sol-Go]
// 完整的并查集模板，见我的数据结构题单
type unionFind struct {
	fa []int
	cc int // 连通块的个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, n}
}

func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return
	}
	u.fa[x] = y
	u.cc-- // 合并后，连通块个数减一
}

func minSwaps(nums []int) int {
	n := len(nums)
	type tuple struct{ s, x, i int }
	a := make([]tuple, n)
	for i, num := range nums {
		s := 0
		for x := num; x > 0; x /= 10 {
			s += x % 10
		}
		a[i] = tuple{s, num, i}
	}

	slices.SortFunc(a, func(a, b tuple) int { return cmp.Or(a.s-b.s, a.x-b.x) })

	u := newUnionFind(n)
	for i, p := range a {
		u.merge(i, p.i)
	}
	return n - u.cc
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U + n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

[2471. 逐层排序二叉树所需的最少操作数目](https://leetcode.cn/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/) 1635

更多相似题目，见下面图论题单的 **DFS** 或者数据结构题单的**并查集**。

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
