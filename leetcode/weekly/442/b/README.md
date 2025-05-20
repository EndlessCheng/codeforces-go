暴力枚举所有 $\textit{properties}[i]$ 和 $\textit{properties}[j]$，如果交集大小 $\ge k$，那么用并查集合并 $i$ 和 $j$。

初始连通块个数 $\textit{cc}=n$，每成功合并一次，就把 $\textit{cc}$ 减一。

完整的并查集模板，见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/)。

[本题视频讲解](https://www.bilibili.com/video/BV12eXYYVE5H/?t=2m21s)，欢迎点赞关注~

```py [sol-Python3]
class UnionFind:
    def __init__(self, n: int):
        # 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        # 集合 i 的代表元是自己
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
        self.cc -= 1  # 合并后，连通块个数减少了 1

class Solution:
    def numberOfComponents(self, properties: List[List[int]], k: int) -> int:
        sets = list(map(set, properties))
        uf = UnionFind(len(properties))
        for i, a in enumerate(sets):
            for j, b in enumerate(sets[:i]):
                if len(a & b) >= k:
                    uf.merge(j, i)
        return uf.cc
```

```py [sol-Python3 位运算]
class UnionFind:
    def __init__(self, n: int):
        # 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        # 集合 i 的代表元是自己
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
        self.cc -= 1  # 合并后，连通块个数减少了 1

class Solution:
    def numberOfComponents(self, properties: List[List[int]], k: int) -> int:
        sets = [reduce(or_, (1 << x for x in a)) for a in properties]
        uf = UnionFind(len(properties))
        for i, a in enumerate(sets):
            for j, b in enumerate(sets[:i]):
                if (a & b).bit_count() >= k:
                    uf.merge(j, i)
        return uf.cc
```

```java [sol-Java]
class UnionFind {
    private final int[] fa;
    public int cc; // 连通块个数

    UnionFind(int n) {
        fa = new int[n];
        cc = n;
        for (int i = 0; i < n; i++) {
            fa[i] = i;
        }
    }

    public int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]);
        }
        return fa[x];
    }

    public void merge(int from, int to) {
        int x = find(from);
        int y = find(to);
        if (x == y) {
            return;
        }
        fa[x] = y;
        cc--;
    }
}

class Solution {
    public int numberOfComponents(int[][] properties, int k) {
        int n = properties.length;
        int m = properties[0].length;
        Set<Integer>[] sets = new HashSet[n];
        Arrays.setAll(sets, i -> new HashSet<>(m));
        for (int i = 0; i < n; i++) {
            for (int x : properties[i]) {
                sets[i].add(x);
            }
        }

        UnionFind u = new UnionFind(n);
        for (int i = 0; i < n; i++) {
            Set<Integer> a = sets[i];
            for (int j = 0; j < i; j++) {
                int cnt = 0;
                for (int x : sets[j]) {
                    if (a.contains(x)) {
                        cnt++;
                    }
                }
                if (cnt >= k) {
                    u.merge(i, j);
                }
            }
        }
        return u.cc;
    }
}
```

```cpp [sol-C++]
class UnionFind {
    vector<int> fa;

public:
    int cc; // 连通块个数

    UnionFind(int n) : fa(n), cc(n) {
        ranges::iota(fa, 0);
    }

    int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]);
        }
        return fa[x];
    }

    void merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) {
            return;
        }
        fa[x] = y;
        cc--;
    }
};

class Solution {
public:
    int numberOfComponents(vector<vector<int>>& properties, int k) {
        int n = properties.size();
        vector<unordered_set<int>> sets(n);
        for (int i = 0; i < n; i++) {
            sets[i] = unordered_set(properties[i].begin(), properties[i].end());
        }

        UnionFind uf(n);
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < i; j++) {
                int cnt = 0;
                for (int x : sets[j]) {
                    if (sets[i].contains(x)) {
                        cnt++;
                    }
                }
                if (cnt >= k) {
                    uf.merge(i, j);
                }
            }
        }
        return uf.cc;
    }
};
```

```go [sol-Go]
type uf struct {
	fa []int
	cc int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa, n}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return
	}
	u.fa[x] = y
	u.cc--
}

func numberOfComponents(properties [][]int, k int) int {
	sets := make([]map[int]bool, len(properties))
	for i, a := range properties {
		sets[i] = map[int]bool{}
		for _, x := range a {
			sets[i][x] = true
		}
	}

	u := newUnionFind(len(properties))
	for i, a := range sets {
		for j, b := range sets[:i] {
			cnt := 0
			for x := range b {
				if a[x] {
					cnt++
				}
			}
			if cnt >= k {
				u.merge(i, j)
			}
		}
	}
	return u.cc
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2m)$ 或者 $\mathcal{O}(n^2m/w)$，其中 $n$ 是 $\textit{properties}$ 的长度，$m$ 是 $\textit{properties}[i]$ 的长度，$w$ 等于 $32$ 或 $64$。
- 空间复杂度：$\mathcal{O}(nm)$ 或者 $\mathcal{O}(nm/w)$。

更多相似题目，见下面数据结构题单中的「**并查集**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
