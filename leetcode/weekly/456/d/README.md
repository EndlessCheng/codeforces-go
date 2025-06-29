本文把强度称作**边权**。

## 方法一：二分答案 + 并查集

### 转化

如果最小边权 $\ge 6$ 时可以得到生成树 $T$，那么最小边权 $\ge 5$（或者更小）时，选择同样的边，同样可以得到生成树 $T$。

如果最小边权 $\ge 7$ 时无法得到生成树 $T$，那么最小边权 $\ge 8$（或者更大）时，限制条件更加苛刻，无法得到生成树。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定 $\textit{low}$，能否得到一棵生成树，其边权都 $\ge \textit{low}$？

如果可以得到一棵生成树，说明答案 $\ge \textit{low}$，否则答案 $< \textit{low}$。

### 思路

根据示例 3，我们首先特判必选边成环的情况，此时无法得到生成树，返回 $-1$。

或者，如果原图不是连通的（连通块大于 $1$），也无法得到生成树，返回 $-1$。

否则一定可以得到生成树。

然后二分猜答案，处理上文说的判定性问题：

- 对于必选边，以及边权 $\ge \textit{low}$ 的边，用**并查集**合并。
- 其余边，剩余升级次数大于 $0$，且边权乘 $2$ 后 $\ge \textit{low}$，且边的两个点不在同一个连通块中，那么用**并查集**合并，同时把剩余升级次数减一。

处理结束后，如果连通块个数等于 $1$，说明可以可以得到一棵生成树。

由于本题保证 $s_i > 0$，如果二分结束后答案为 $0$，说明我们无法得到任何生成树，返回 $-1$。

### 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$\min(s_i)$。无需操作，一定满足要求。（注意我们已经提前特判返回 $-1$ 的情况了）
- 开区间右端点初始值：$\min(m_0\cdot 2, m_1)+1$，其中 $m_0$ 是非必选边的最大值（如果不存在则为 $0$），$m_1$ 是必选边的最小值（如果不存在则为 $\infty$）。把 $\min(m_0\cdot 2, m_1)$ 加一后，一定无法满足要求。

对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

完整并查集模板见数据结构题单。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

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
    # 返回是否合并成功
    def merge(self, from_: int, to: int) -> bool:
        x, y = self.find(from_), self.find(to)
        if x == y:  # from 和 to 在同一个集合，不做合并
            return False
        self._fa[x] = y  # 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        self.cc -= 1  # 成功合并，连通块个数减一
        return True


class Solution:
    def maxStability(self, n: int, edges: List[List[int]], k: int) -> int:
        must_uf = UnionFind(n)  # 必选边并查集
        all_uf = UnionFind(n)  # 全图并查集
        min_s, max_s = inf, 0
        for x, y, s, must in edges:
            if must and not must_uf.merge(x, y):  # 必选边成环
                return -1
            all_uf.merge(x, y)
            min_s = min(min_s, s)
            max_s = max(max_s, s)

        if all_uf.cc > 1:  # 图不连通
            return -1

        def check(low: int) -> bool:
            u = UnionFind(n)
            for x, y, s, must in edges:
                if must and s < low:  # 必选边的边权太小
                    return False
                if must or s >= low:
                    u.merge(x, y)

            left_k = k
            for x, y, s, must in edges:
                if left_k == 0 or u.cc == 1:
                    break
                if not must and s < low <= s * 2 and u.merge(x, y):
                    left_k -= 1

            return u.cc == 1

        left, right = min_s, max_s * 2 + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```java [sol-Java]
class UnionFind {
    private final int[] fa; // 代表元
    public int cc; // 连通块个数

    UnionFind(int n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己
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
    // 返回是否合并成功
    public boolean merge(int from, int to) {
        int x = find(from);
        int y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return false;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
        return true;
    }
}

class Solution {
    public int maxStability(int n, int[][] edges, int k) {
        UnionFind mustUf = new UnionFind(n); // 必选边并查集
        UnionFind allUf = new UnionFind(n); // 全图并查集
        int minS = Integer.MAX_VALUE;
        int maxS = 0;
        for (int[] e : edges) {
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (must > 0 && !mustUf.merge(x, y)) { // 必选边成环
                return -1;
            }
            allUf.merge(x, y);
            minS = Math.min(minS, s);
            maxS = Math.max(maxS, s);
        }

        if (allUf.cc > 1) { // 图不连通
            return -1;
        }

        int left = minS;
        int right = maxS * 2 + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(mid, n, edges, k)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(int low, int n, int[][] edges, int k) {
        UnionFind u = new UnionFind(n);
        for (int[] e : edges) {
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (must > 0 && s < low) { // 必选边的边权太小
                return false;
            }
            if (must > 0 || s >= low) {
                u.merge(x, y);
            }
        }

        for (int[] e : edges) {
            if (k == 0 || u.cc == 1) {
                break;
            }
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (must == 0 && s < low && s * 2 >= low && u.merge(x, y)) {
                k--;
            }
        }
        return u.cc == 1;
    }
}
```

```cpp [sol-C++]
class UnionFind {
    vector<int> fa; // 代表元

public:
    int cc; // 连通块个数

    UnionFind(int n) : fa(n), cc(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己
        ranges::iota(fa, 0);
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
    bool merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return false;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
        return true;
    }
};

class Solution {
public:
    int maxStability(int n, vector<vector<int>>& edges, int k) {
        UnionFind must_uf(n); // 必选边并查集
        UnionFind all_uf(n); // 全图并查集
        int min_s = INT_MAX, max_s = 0;
        for (auto& e : edges) {
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (must && !must_uf.merge(x, y)) { // 必选边成环
                return -1;
            }
            all_uf.merge(x, y);
            min_s = min(min_s, s);
            max_s = max(max_s, s);
        }

        if (all_uf.cc > 1) { // 图不连通
            return -1;
        }

        auto check = [&](int low) -> bool {
            UnionFind u(n);
            for (auto& e : edges) {
                int x = e[0], y = e[1], s = e[2], must = e[3];
                if (must && s < low) { // 必选边的边权太小
                    return false;
                }
                if (must || s >= low) {
                    u.merge(x, y);
                }
            }

            int left_k = k;
            for (auto& e : edges) {
                if (left_k == 0 || u.cc == 1) {
                    break;
                }
                int x = e[0], y = e[1], s = e[2], must = e[3];
                if (!must && s < low && s * 2 >= low && u.merge(x, y)) {
                    left_k--;
                }
            }
            return u.cc == 1;
        };

        int left = min_s, right = max_s * 2 + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
type unionFind struct {
	fa []int // 代表元
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, n}
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
// 返回是否合并成功
func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return false
	}
	u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
	u.cc--      // 成功合并，连通块个数减一
	return true
}

func maxStability(n int, edges [][]int, k int) int {
	mustUf := newUnionFind(n) // 必选边并查集
	allUf := newUnionFind(n) // 全图并查集
	minS, maxS := math.MaxInt, 0
	for _, e := range edges {
		x, y, s, must := e[0], e[1], e[2], e[3]
		if must > 0 && !mustUf.merge(x, y) { // 必选边成环
			return -1
		}
		allUf.merge(x, y)
		minS = min(minS, s)
		maxS = max(maxS, s)
	}
	if allUf.cc > 1 { // 图不连通
		return -1
	}

	check := func(low int) bool {
		u := newUnionFind(n)
		for _, e := range edges {
			x, y, s, must := e[0], e[1], e[2], e[3]
			if must > 0 && s < low { // 必选边的边权太小
				return false
			}
			if must > 0 || s >= low {
				u.merge(x, y)
			}
		}
		leftK := k
		for _, e := range edges {
			if leftK == 0 || u.cc == 1 {
				break
			}
			x, y, s, must := e[0], e[1], e[2], e[3]
			if must == 0 && s < low && s*2 >= low && u.merge(x, y) {
				leftK--
			}
		}
		return u.cc == 1
	}

	left, right := minS, maxS*2+1
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log n\log U)$，其中 $m$ 是 $\textit{edges}$ 的长度，$U=\max(s_i) - \min(s_i)$（常数系数可以忽略）。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：如果改成建图连边 + DFS，可以做到 $\mathcal{O}((n+m)\log U)$。

## 方法二：Kruskal 算法

实际上，根据 **Kruskal 算法**，在选出必选边后，剩余的边按照边权从大到小合并即可（类似求最大生成树）。若不按 Kruskal 算法计算，得到的生成树的最小边权只会更小。

设在生成树中的边权 $a$，例如 $a=[9,8,7,6,5]$，贪心地，选最小的边权乘以 $2$，也就是选 $a$ 的后 $k$ 个数乘以 $2$。乘 $2$ 后，答案为如下三者的最小值：

1. 必选边中的最小边权。
2. $a$ 中的最小边权乘以 $2$。
3. $a$ 中的第 $k+1$ 小边权。

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
    # 返回是否合并成功
    def merge(self, from_: int, to: int) -> bool:
        x, y = self.find(from_), self.find(to)
        if x == y:  # from 和 to 在同一个集合，不做合并
            return False
        self._fa[x] = y  # 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        self.cc -= 1  # 成功合并，连通块个数减一
        return True


class Solution:
    def maxStability(self, n: int, edges: List[List[int]], k: int) -> int:
        uf = UnionFind(n)
        all_uf = UnionFind(n)
        min_s1 = inf
        for x, y, s, must in edges:
            if must:
                if not uf.merge(x, y):  # 必选边成环
                    return -1
                min_s1 = min(min_s1, s)
            all_uf.merge(x, y)

        if all_uf.cc > 1:  # 图不连通
            return -1

        # Kruskal 求最大生成树
        edges.sort(key=lambda e: -e[2])
        a = []
        for x, y, s, must in edges:
            if not must and uf.merge(x, y):
                a.append(s)

        if not a:
            return min_s1
        ans = min(min_s1, a[-1] * 2)
        if k < len(a):
            ans = min(ans, a[-1 - k])
        return ans
```

```java [sol-Java]
class UnionFind {
    private final int[] fa; // 代表元
    public int cc; // 连通块个数

    UnionFind(int n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己
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
    // 返回是否合并成功
    public boolean merge(int from, int to) {
        int x = find(from);
        int y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return false;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
        return true;
    }
}

class Solution {
    public int maxStability(int n, int[][] edges, int k) {
        UnionFind uf = new UnionFind(n);
        UnionFind allUf = new UnionFind(n);
        int minS1 = Integer.MAX_VALUE;
        for (int[] e : edges) {
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (must > 0) {
                if (!uf.merge(x, y)) { // 必选边成环
                    return -1;
                }
                minS1 = Math.min(minS1, s);
            }
            allUf.merge(x, y);
        }

        if (allUf.cc > 1) { // 图不连通
            return -1;
        }

        // Kruskal 求最大生成树
        Arrays.sort(edges, (a, b) -> b[2] - a[2]);
        List<Integer> a = new ArrayList<>();
        for (int[] e : edges) {
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (must == 0 && uf.merge(x, y)) {
                a.add(s);
            }
        }

        int m = a.size();
        if (m == 0) {
            return minS1;
        }
        int ans = Math.min(minS1, a.get(m - 1) * 2);
        if (k < m) {
            ans = Math.min(ans, a.get(m - 1 - k));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class UnionFind {
    vector<int> fa; // 代表元

public:
    int cc; // 连通块个数

    UnionFind(int n) : fa(n), cc(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己
        ranges::iota(fa, 0);
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
    bool merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return false;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        cc--; // 成功合并，连通块个数减一
        return true;
    }
};

class Solution {
public:
    int maxStability(int n, vector<vector<int>>& edges, int k) {
        UnionFind uf(n);
        UnionFind all_uf(n);
        int min_s1 = INT_MAX;
        for (auto& e : edges) {
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (must) {
                if (!uf.merge(x, y)) { // 必选边成环
                    return -1;
                }
                min_s1 = min(min_s1, s);
            }
            all_uf.merge(x, y);
        }

        if (all_uf.cc > 1) { // 图不连通
            return -1;
        }

        // Kruskal 求最大生成树
        ranges::sort(edges, {}, [](auto& e) { return -e[2]; });
        vector<int> a;
        for (auto& e : edges) {
            int x = e[0], y = e[1], s = e[2], must = e[3];
            if (!must && uf.merge(x, y)) {
                a.push_back(s);
            }
        }

        int m = a.size();
        if (m == 0) {
            return min_s1;
        }
        int ans = min(min_s1, a[m - 1] * 2);
        if (k < m) {
            ans = min(ans, a[m - 1 - k]);
        }
        return ans;
    }
};
```

```go [sol-Go]
type unionFind struct {
	fa []int // 代表元
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, n}
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
// 返回是否合并成功
func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return false
	}
	u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
	u.cc--      // 成功合并，连通块个数减一
	return true
}

func maxStability(n int, edges [][]int, k int) int {
	uf := newUnionFind(n)
	allUf := newUnionFind(n)
	minS1 := math.MaxInt
	for _, e := range edges {
		x, y, s, must := e[0], e[1], e[2], e[3]
		if must > 0 {
			if !uf.merge(x, y) { // 必选边成环
				return -1
			}
			minS1 = min(minS1, s)
		}
		allUf.merge(x, y)
	}

	if allUf.cc > 1 { // 图不连通
		return -1
	}

	// Kruskal 算法求最大生成树
	slices.SortFunc(edges, func(a, b []int) int { return b[2] - a[2] })
	a := []int{}
	for _, e := range edges {
		x, y, s, must := e[0], e[1], e[2], e[3]
		if must == 0 && uf.merge(x, y) {
			a = append(a, s)
		}
	}

	// 如下三者的最小值：
	// 1. must = 1 中的最小值
	// 2. a 中最小边权 * 2
	// 3. a 中第 k+1 小边权
	m := len(a)
	if m == 0 {
		return minS1
	}
	ans := min(minS1, a[m-1]*2)
	if k < m {
		ans = min(ans, a[m-1-k])
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m + (n+m)\log n)$，其中 $m$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 二分题单的「**§2.5 最大化最小值**」。
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
