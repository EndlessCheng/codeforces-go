为方便讨论，下文把 $\textit{colors}$ 数组简记为 $a$。

## 一、非环形数组 + 没有修改操作

这是最简单的情况。

以 $a=[0,1,1,0,1,0,0,1,0]$ 为例，它可以分为三个交替段：$[0,1],[1,0,1,0],[0,1,0]$。

有多少个长度恰好为 $\textit{size}=3$ 的交替子数组？有 $3$ 个：$[1,0,1,0]$ 中有 $2$ 个，$[0,1,0]$ 中有 $1$ 个。

如果遍历所有交替段，时间复杂度就太高了。

一般地，对于长为 $k$ 的交替段，其中有

$$
k - (\textit{size} - 1)
$$

个长度恰好为 $\textit{size}$ 的交替子数组。

假设有 $3$ 个长度均 $\ge \textit{size}$ 的交替段，长度分别为 $k_1,k_2,k_3$，那么其中有

$$
\begin{aligned}
    & k_1 - (\textit{size} - 1) + k_2 - (\textit{size} - 1) + k_3 - (\textit{size} - 1)      \\
={} & (k_1+k_2+k_3) - 3 \cdot (\textit{size} - 1)        \\
\end{aligned}
$$

个长度恰好为 $\textit{size}$ 的交替子数组。

这启发我们维护长度 $\ge \textit{size}$ 的交替段的**个数**以及**元素和**。考虑到后续要执行修改操作，用**树状数组**维护。

## 二、非环形数组 + 有修改操作

考虑交替段的结束位置，即 $i=n-1$ 或者 $a[i]=a[i+1]$ 的位置。

如果修改 $a[i]$ 的值，**会影响哪些结束位置**？**会对交替段的长度产生什么影响**？

结束位置 $i-1$ 和 $i$ 会受到影响。

为避免复杂的分类讨论，在修改之前，先移除掉结束位置 $i-1$ 和 $i$（如果是结束位置的话），然后再根据 $a[i-1]=a[i]$ 以及 $a[i]=a[i+1]$ 是否成立，添加结束位置 $i-1$ 和 $i$。

- 当我们添加结束位置 $i$ 时，设 $\textit{pre}$ 和 $\textit{nxt}$ 是 $i$ 前后两个相邻的结束位置，那么一个长为 $\textit{nxt} - \textit{pre}$ 的交替段会被拆分为两个长度分别为 $i-\textit{pre}$ 和 $\textit{nxt}-i$ 的交替段。
- 当我们移除结束位置 $i$ 时，设 $\textit{pre}$ 和 $\textit{nxt}$ 是 $i$ 前后两个相邻的结束位置，那么两个长度分别为 $i-\textit{pre}$ 和 $\textit{nxt}-i$ 的交替段会合并成一个长为 $\textit{nxt} - \textit{pre}$ 的交替段。

如何快速找到 $i$ 的前后两个相邻的结束位置？

用有序集合维护所有结束位置。

## 三、环形数组 + 有修改操作

$a$ 变成环形数组后，哪些计算会发生变化？

1. $i=n-1$ 不一定是结束位置了，必须判断 $a[n-1]=a[0]$ 是否成立才行。
2. 如果小于 $i$ 的最大结束位置不存在，那么取所有结束位置的最大值。
3. 如果大于 $i$ 的最小结束位置不存在，那么取所有结束位置的最小值。
4. 对于交替段长度，例如 $\textit{nxt} - \textit{pre}$，可能会算出负数或 $0$。可以通过 $(\textit{nxt} - \textit{pre} + n - 1)\bmod n + 1$ 调整成 $[1,n]$ 中的数。

### 细节

1. 如果没有结束位置，那么无论询问多长的交替子数组，答案都是 $n$。
2. 如果修改操作不改变 $a[i]$ 的值，则直接 `continue`。

具体请看 [视频讲解](https://www.bilibili.com/video/BV124421Z78J/) 第四题，欢迎点赞关注！

```py [sol-Python3]
from sortedcontainers import SortedList

class FenwickTree:
    def __init__(self, n: int):
        self.t = [[0, 0] for _ in range(n + 1)]

    # op=1，添加一个 size
    # op=-1，移除一个 size
    def update(self, size: int, op: int) -> None:
        i = len(self.t) - size
        while i < len(self.t):
            self.t[i][0] += op
            self.t[i][1] += op * size
            i += i & -i

    # 返回 >= size 的元素个数，元素和
    def query(self, size: int) -> (int, int):
        cnt = s = 0
        i = len(self.t) - size
        while i > 0:
            cnt += self.t[i][0]
            s += self.t[i][1]
            i &= i - 1
        return cnt, s

class Solution:
    def numberOfAlternatingGroups(self, a: List[int], queries: List[List[int]]) -> List[int]:
        n = len(a)
        sl = SortedList()
        t = FenwickTree(n)

        # op=1，添加一个结束位置 i
        # op=-1，移除一个结束位置 i
        def update(i: int, op: int) -> None:
            idx = sl.bisect_left(i)
            pre = sl[idx - 1]
            nxt = sl[idx % len(sl)]

            t.update((nxt - pre - 1) % n + 1, -op)  # 移除/添加旧长度
            t.update((i - pre) % n, op)
            t.update((nxt - i) % n, op)  # 添加/移除新长度

        # 添加一个结束位置 i
        def add(i: int) -> None:
            if not sl:
                t.update(n, 1)
            else:
                update(i, 1)
            sl.add(i)

        # 移除一个结束位置 i
        def remove(i: int) -> None:
            sl.remove(i)
            if not sl:
                t.update(n, -1)
            else:
                update(i, -1)

        for i, c in enumerate(a):
            if c == a[(i + 1) % n]:
                add(i)  # i 是一个结束位置

        ans = []
        for q in queries:
            if q[0] == 1:
                if not sl:
                    ans.append(n)  # 每个长为 size 的子数组都符合要求
                else:
                    cnt, s = t.query(q[1])
                    ans.append(s - cnt * (q[1] - 1))
            else:
                i, c = q[1], q[2]
                if a[i] == c:  # 无影响
                    continue
                pre, nxt = (i - 1) % n, (i + 1) % n
                # 修改前，先去掉结束位置
                if a[pre] == a[i]:
                    remove(pre)
                if a[i] == a[nxt]:
                    remove(i)
                a[i] = c
                # 修改后，添加新的结束位置
                if a[pre] == a[i]:
                    add(pre)
                if a[i] == a[nxt]:
                    add(i)
        return ans
```

```py [sol-Python3 写法二]
from sortedcontainers import SortedList

class Solution:
    def numberOfAlternatingGroups(self, a: List[int], queries: List[List[int]]) -> List[int]:
        n = len(a)
        sl = SortedList()
        t = [[0, 0] for _ in range(n + 1)]

        # op=1，添加一个 size
        # op=-1，移除一个 size
        def fenwick_update(size: int, op: int) -> None:
            i = len(t) - size
            while i < len(t):
                t[i][0] += op
                t[i][1] += op * size
                i += i & -i

        # 返回 >= size 的元素个数，元素和
        def fenwick_query(size: int) -> (int, int):
            cnt = s = 0
            i = len(t) - size
            while i > 0:
                cnt += t[i][0]
                s += t[i][1]
                i &= i - 1
            return cnt, s

        # op=1，添加一个结束位置 i
        # op=-1，移除一个结束位置 i
        def update(i: int, op: int) -> None:
            idx = sl.bisect_left(i)
            pre = sl[idx - 1]
            nxt = sl[idx % len(sl)]

            fenwick_update((nxt - pre - 1) % n + 1, -op)  # 移除/添加旧长度
            fenwick_update((i - pre) % n, op)
            fenwick_update((nxt - i) % n, op)  # 添加/移除新长度

        # 添加一个结束位置 i
        def add(i: int) -> None:
            if not sl:
                fenwick_update(n, 1)
            else:
                update(i, 1)
            sl.add(i)

        # 移除一个结束位置 i
        def remove(i: int) -> None:
            sl.remove(i)
            if not sl:
                fenwick_update(n, -1)
            else:
                update(i, -1)

        for i, c in enumerate(a):
            if c == a[(i + 1) % n]:
                add(i)  # i 是一个结束位置

        ans = []
        for q in queries:
            if q[0] == 1:
                if not sl:
                    ans.append(n)  # 每个长为 size 的子数组都符合要求
                else:
                    cnt, s = fenwick_query(q[1])
                    ans.append(s - cnt * (q[1] - 1))
            else:
                i, c = q[1], q[2]
                if a[i] == c:  # 无影响
                    continue
                pre, nxt = (i - 1) % n, (i + 1) % n
                # 修改前，先去掉结束位置
                if a[pre] == a[i]:
                    remove(pre)
                if a[i] == a[nxt]:
                    remove(i)
                a[i] = c
                # 修改后，添加新的结束位置
                if a[pre] == a[i]:
                    add(pre)
                if a[i] == a[nxt]:
                    add(i)
        return ans
```

```java [sol-Java]
class FenwickTree {
    private final int[][] t;

    public FenwickTree(int n) {
        t = new int[n + 1][2];
    }

    // op=1，添加一个 size
    // op=-1，移除一个 size
    public void update(int size, int op) {
        for (int i = t.length - size; i < t.length; i += i & -i) {
            t[i][0] += op;
            t[i][1] += op * size;
        }
    }

    // 返回 >= size 的元素个数，元素和
    public int[] query(int size) {
        int cnt = 0, sum = 0;
        for (int i = t.length - size; i > 0; i &= i - 1) {
            cnt += t[i][0];
            sum += t[i][1];
        }
        return new int[]{cnt, sum};
    }
}

class Solution {
    public List<Integer> numberOfAlternatingGroups(int[] a, int[][] queries) {
        int n = a.length;
        TreeSet<Integer> set = new TreeSet<>();
        FenwickTree t = new FenwickTree(n);

        for (int i = 0; i < n; i++) {
            if (a[i] == a[(i + 1) % n]) {
                add(set, t, n, i); // i 是一个结束位置
            }
        }

        List<Integer> ans = new ArrayList<>();
        for (int[] q : queries) {
            if (q[0] == 1) {
                if (set.isEmpty()) {
                    ans.add(n); // 每个长为 size 的子数组都符合要求
                } else {
                    int[] res = t.query(q[1]);
                    ans.add(res[1] - res[0] * (q[1] - 1));
                }
            } else {
                int i = q[1];
                if (a[i] == q[2]) { // 无影响
                    continue;
                }
                int pre = (i - 1 + n) % n;
                int nxt = (i + 1) % n;
                // 修改前，先去掉结束位置
                if (a[pre] == a[i]) {
                    del(set, t, n, pre);
                }
                if (a[i] == a[nxt]) {
                    del(set, t, n, i);
                }
                a[i] ^= 1;
                // 修改后，添加新的结束位置
                if (a[pre] == a[i]) {
                    add(set, t, n, pre);
                }
                if (a[i] == a[nxt]) {
                    add(set, t, n, i);
                }
            }
        }
        return ans;
    }

    // 添加一个结束位置 i
    private void add(TreeSet<Integer> set, FenwickTree t, int n, int i) {
        if (set.isEmpty()) {
            t.update(n, 1);
        } else {
            update(set, t, n, i, 1);
        }
        set.add(i);
    }

    // 移除一个结束位置 i
    private void del(TreeSet<Integer> set, FenwickTree t, int n, int i) {
        set.remove(i);
        if (set.isEmpty()) {
            t.update(n, -1);
        } else {
            update(set, t, n, i, -1);
        }
    }

    // op=1，添加一个结束位置 i
    // op=-1，移除一个结束位置 i
    private void update(TreeSet<Integer> set, FenwickTree t, int n, int i, int op) {
        Integer pre = set.floor(i);
        if (pre == null) {
            pre = set.last();
        }

        Integer nxt = set.ceiling(i);
        if (nxt == null) {
            nxt = set.first();
        }

        t.update((nxt - pre + n - 1) % n + 1, -op); // 移除/添加旧长度
        t.update((i - pre + n) % n, op);
        t.update((nxt - i + n) % n, op); // 添加/移除新长度
    }
}
```

```cpp [sol-C++]
class FenwickTree {
    vector<array<int, 2>> tree;
public:
    FenwickTree(int n) : tree(n + 1) {}

    // op=1，添加一个 size
    // op=-1，移除一个 size
    void update(int size, int op) {
        for (int i = tree.size() - size; i < tree.size(); i += i & -i) {
            tree[i][0] += op;
            tree[i][1] += op * size;
        }
    }

    // 返回 >= size 的元素个数，元素和
    pair<int, int> query(int size) {
        int cnt = 0, sum = 0;
        for (int i = tree.size() - size; i > 0; i &= i - 1) {
            cnt += tree[i][0];
            sum += tree[i][1];
        }
        return {cnt, sum};
    }
};

class Solution {
public:
    vector<int> numberOfAlternatingGroups(vector<int>& a, vector<vector<int>>& queries) {
        int n = a.size();
        set<int> s;
        FenwickTree t(n);

        // op=1，添加一个结束位置 i
        // op=-1，移除一个结束位置 i
        auto update = [&](int i, int op) {
            auto it = s.lower_bound(i);
            int pre = it == s.begin() ? *s.rbegin() : *prev(it);
            int nxt = it == s.end() ? *s.begin() : *it;

            t.update((nxt - pre + n - 1) % n + 1, -op); // 移除/添加旧长度
            t.update((i - pre + n) % n, op);
            t.update((nxt - i + n) % n, op); // 添加/移除新长度
        };

        // 添加一个结束位置 i
        auto add = [&](int i) {
            if (s.empty()) {
                t.update(n, 1);
            } else {
                update(i, 1);
            }
            s.insert(i);
        };

        // 移除一个结束位置 i
        auto del = [&](int i) {
            s.erase(i);
            if (s.empty()) {
                t.update(n, -1);
            } else {
                update(i, -1);
            }
        };

        for (int i = 0; i < n; i++) {
            if (a[i] == a[(i + 1) % n]) {
                add(i); // i 是一个结束位置
            }
        }

        vector<int> ans;
        for (auto& q : queries) {
            if (q[0] == 1) {
                if (s.empty()) {
                    ans.push_back(n); // 每个长为 size 的子数组都符合要求
                } else {
                    auto [cnt, sum] = t.query(q[1]);
                    ans.push_back(sum - cnt * (q[1] - 1));
                }
            } else {
                int i = q[1];
                if (a[i] == q[2]) { // 无影响
                    continue;
                }
                int pre = (i - 1 + n) % n, nxt = (i + 1) % n;
                // 修改前，先去掉结束位置
                if (a[pre] == a[i]) {
                    del(pre);
                }
                if (a[i] == a[nxt]) {
                    del(i);
                }
                a[i] ^= 1;
                // 修改后，添加新的结束位置
                if (a[pre] == a[i]) {
                    add(pre);
                }
                if (a[i] == a[nxt]) {
                    add(i);
                }
            }
        }
        return ans;
    }
};
```

```cpp [sol-C++ 写法二]
class FenwickTree {
    vector<array<int, 2>> tree;
public:
    FenwickTree(int n) : tree(n + 1) {}

    // op=1，添加一个 size
    // op=-1，移除一个 size
    void update(int size, int op) {
        for (int i = tree.size() - size; i < tree.size(); i += i & -i) {
            tree[i][0] += op;
            tree[i][1] += op * size;
        }
    }

    // 返回 >= size 的元素个数，元素和
    pair<int, int> query(int size) {
        int cnt = 0, sum = 0;
        for (int i = tree.size() - size; i > 0; i &= i - 1) {
            cnt += tree[i][0];
            sum += tree[i][1];
        }
        return {cnt, sum};
    }
};

class Solution {
public:
    vector<int> numberOfAlternatingGroups(vector<int>& a, vector<vector<int>>& queries) {
        int n = a.size();
        set<int> s;
        FenwickTree t(n);

        // op=1，添加一个结束位置 i
        // op=-1，移除一个结束位置 i
        auto update = [&](int pre, int nxt, int i, int op) {
            t.update((nxt - pre + n - 1) % n + 1, -op); // 移除/添加旧长度
            t.update((i - pre + n) % n, op);
            t.update((nxt - i + n) % n, op); // 添加/移除新长度
        };

        // 添加一个结束位置 i
        auto add = [&](int i) {
            if (s.empty()) {
                t.update(n, 1);
                s.insert(i);
            } else {
                auto it = s.insert(i).first;
                int pre = it == s.begin() ? *s.rbegin() : *prev(it);
                int nxt = next(it) == s.end() ? *s.begin() : *next(it);
                update(pre, nxt, i, 1);
            }
        };

        // 移除一个结束位置 i
        auto del = [&](int i) {
            auto it = s.erase(s.find(i));
            if (s.empty()) {
                t.update(n, -1);
            } else {
                int pre = it == s.begin() ? *s.rbegin() : *prev(it);
                int nxt = it == s.end() ? *s.begin() : *it;
                update(pre, nxt, i, -1);
            }
        };

        for (int i = 0; i < n; i++) {
            if (a[i] == a[(i + 1) % n]) {
                add(i); // i 是一个结束位置
            }
        }

        vector<int> ans;
        for (auto& q : queries) {
            if (q[0] == 1) {
                if (s.empty()) {
                    ans.push_back(n); // 每个长为 size 的子数组都符合要求
                } else {
                    auto [cnt, sum] = t.query(q[1]);
                    ans.push_back(sum - cnt * (q[1] - 1));
                }
            } else {
                int i = q[1];
                if (a[i] == q[2]) { // 无影响
                    continue;
                }
                int pre = (i - 1 + n) % n, nxt = (i + 1) % n;
                // 修改前，先去掉结束位置
                if (a[pre] == a[i]) {
                    del(pre);
                }
                if (a[i] == a[nxt]) {
                    del(i);
                }
                a[i] ^= 1;
                // 修改后，添加新的结束位置
                if (a[pre] == a[i]) {
                    add(pre);
                }
                if (a[i] == a[nxt]) {
                    add(i);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type fenwickTree [][2]int

// op=1，添加一个 size
// op=-1，移除一个 size
func (t fenwickTree) update(size, op int) {
	for i := len(t) - size; i < len(t); i += i & -i {
		t[i][0] += op
		t[i][1] += op * size
	}
}

// 返回 >= size 的元素个数，元素和
func (t fenwickTree) query(size int) (cnt, sum int) {
	for i := len(t) - size; i > 0; i &= i - 1 {
		cnt += t[i][0]
		sum += t[i][1]
	}
	return
}

func numberOfAlternatingGroups(a []int, queries [][]int) (ans []int) {
	n := len(a)
	set := redblacktree.New[int, struct{}]()
	t := make(fenwickTree, n+1)

	// op=1，添加一个结束位置 i
	// op=-1，移除一个结束位置 i
	update := func(i, op int) {
		prev, ok := set.Floor(i)
		if !ok {
			prev = set.Right()
		}
		pre := prev.Key

		next, ok := set.Ceiling(i)
		if !ok {
			next = set.Left()
		}
		nxt := next.Key

		t.update((nxt-pre+n-1)%n+1, -op) // 移除/添加旧长度
		t.update((i-pre+n)%n, op)
		t.update((nxt-i+n)%n, op) // 添加/移除新长度
	}

	// 添加一个结束位置 i
	add := func(i int) {
		if set.Empty() {
			t.update(n, 1)
		} else {
			update(i, 1)
		}
		set.Put(i, struct{}{})
	}

	// 移除一个结束位置 i
	del := func(i int) {
		set.Remove(i)
		if set.Empty() {
			t.update(n, -1)
		} else {
			update(i, -1)
		}
	}

	for i, c := range a {
		if c == a[(i+1)%n] {
			add(i) // i 是一个结束位置
		}
	}
	for _, q := range queries {
		if q[0] == 1 {
			if set.Empty() {
				ans = append(ans, n) // 每个长为 size 的子数组都符合要求
			} else {
				cnt, sum := t.query(q[1])
				ans = append(ans, sum-cnt*(q[1]-1))
			}
		} else {
			i := q[1]
			if a[i] == q[2] { // 无影响
				continue
			}
			pre, nxt := (i-1+n)%n, (i+1)%n
			// 修改前，先去掉结束位置
			if a[pre] == a[i] {
				del(pre)
			}
			if a[i] == a[nxt] {
				del(i)
			}
			a[i] ^= 1
			// 修改后，添加新的结束位置
			if a[pre] == a[i] {
				add(pre)
			}
			if a[i] == a[nxt] {
				add(i)
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 是 $\textit{colors}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
