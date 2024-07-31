### 方法一：正序回答询问+有序集合+线段树

![b131d.png](https://pic.leetcode.cn/1716685098-OuZZol-b131d.png)

看示例 2：

- 首先，在 $x=7$ 处放置一个障碍物，这会产生一个长度为 $7$ 的没有障碍物的区域。
- 然后，在 $x=2$ 处放置一个障碍物，去掉原来的 $7$，产生两个长度分别为 $2,5$ 的没有障碍物的区域。

这里的区域长度 $2,5,7$ 应该保存在哪个位置，从而方便我们查询？

查询的范围是 $[0,x]$，这可以分为两个部分。以 $x=8$ 为例：

- 设 $\textit{pre}$ 是在 $x$ 左侧最近的障碍物的位置，这里 $\textit{pre}=7$。
- 查询「完整」的没有障碍物的区域，这有 $[0,2],[2,7]$ 两段。
- 查询「不完整」的没有障碍物的区域，也就是 $[7,8]$ 这一段。

如果区域的**右端点**在 $[0,x]$ 中，这个区域一定是「完整」的。

所以，**把区域的长度保存在区域的右端点处。**

设 $d[x]$ 为右端点为 $x$ 的区域的长度。

- 如果 $x$ 没有障碍物则 $d[x]=0$；
- 否则 $d[x]$ 等于 $x$ 到其左侧最近障碍物的距离。

为方便讨论，假设 $x=0$ 处有障碍物。

例如示例 2：

- 一开始所有 $d[x]=0$。
- 首先，在 $x=7$ 处放置一个障碍物，现在 $d[7]=7$。
- 然后，在 $x=2$ 处放置一个障碍物，现在 $d[2]=2,\ d[7]=5$。

问题变成如何维护和查询 $d$ 数组，我们需要支持单点更新，区间查询，这可以用**线段树**解决。

此外，我们还需要知道离 $x$ 左右最近的障碍物的位置，这可以用平衡树维护。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1SU411d7wj/) 第四题，欢迎点赞关注！

```py [sol-Python3]
from sortedcontainers import SortedList

class Solution:
    def getResults(self, queries: List[List[int]]) -> List[bool]:
        m = max(q[1] for q in queries) + 1
        t = [0] * (2 << m.bit_length())

        # 把 i 处的值改成 val
        def update(o: int, l: int, r: int, i: int, val: int) -> None:
            if l == r:
                t[o] = val
                return
            m = (l + r) // 2
            if i <= m:
                update(o * 2, l, m, i, val)
            else:
                update(o * 2 + 1, m + 1, r, i, val)
            t[o] = max(t[o * 2], t[o * 2 + 1])

        # 查询 [0,R] 中的最大值
        def query(o: int, l: int, r: int, R: int) -> int:
            if r <= R:
                return t[o]
            m = (l + r) // 2
            if R <= m:
                return query(o * 2, l, m, R)
            return max(t[o * 2], query(o * 2 + 1, m + 1, r, R))

        sl = SortedList([0, m])  # 哨兵
        ans = []
        for q in queries:
            x = q[1]
            i = sl.bisect_left(x)
            pre = sl[i - 1]  # x 左侧最近障碍物的位置
            if q[0] == 1:
                nxt = sl[i]  # x 右侧最近障碍物的位置
                sl.add(x)
                update(1, 0, m, x, x - pre)    # 更新 d[x] = x - pre
                update(1, 0, m, nxt, nxt - x)  # 更新 d[nxt] = nxt - x
            else:
                # 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                max_gap = max(query(1, 0, m, pre), x - pre)
                ans.append(max_gap >= q[2])
        return ans
```

```java [sol-Java]
class Solution {
    public List<Boolean> getResults(int[][] queries) {
        int m = 0;
        for (int[] q : queries) {
            m = Math.max(m, q[1]);
        }
        m++;

        TreeSet<Integer> set = new TreeSet<>(List.of(0, m)); // 哨兵
        int[] t = new int[2 << (32 - Integer.numberOfLeadingZeros(m))];

        List<Boolean> ans = new ArrayList<>();
        for (int[] q : queries) {
            int x = q[1];
            int pre = set.floor(x - 1); // x 左侧最近障碍物的位置
            if (q[0] == 1) {
                int nxt = set.ceiling(x); // x 右侧最近障碍物的位置
                set.add(x);
                update(t, 1, 0, m, x, x - pre);   // 更新 d[x] = x - pre
                update(t, 1, 0, m, nxt, nxt - x); // 更新 d[nxt] = nxt - x
            } else {
                // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                int maxGap = Math.max(query(t, 1, 0, m, pre), x - pre);
                ans.add(maxGap >= q[2]);
            }
        }
        return ans;
    }

    // 把 i 处的值改成 val
    private void update(int[] t, int o, int l, int r, int i, int val) {
        if (l == r) {
            t[o] = val;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) {
            update(t, o * 2, l, m, i, val);
        } else {
            update(t, o * 2 + 1, m + 1, r, i, val);
        }
        t[o] = Math.max(t[o * 2], t[o * 2 + 1]);
    }

    // 查询 [0,R] 中的最大值
    private int query(int[] t, int o, int l, int r, int R) {
        if (r <= R) {
            return t[o];
        }
        int m = (l + r) / 2;
        if (R <= m) {
            return query(t, o * 2, l, m, R);
        }
        return Math.max(t[o * 2], query(t, o * 2 + 1, m + 1, r, R));
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> t;

    // 把 i 处的值改成 val
    void update(int o, int l, int r, int i, int val) {
        if (l == r) {
            t[o] = val;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) {
            update(o * 2, l, m, i, val);
        } else {
            update(o * 2 + 1, m + 1, r, i, val);
        }
        t[o] = max(t[o * 2], t[o * 2 + 1]);
    }

    // 查询 [0,R] 中的最大值
    int query(int o, int l, int r, int R) {
        if (r <= R) {
            return t[o];
        }
        int m = (l + r) / 2;
        if (R <= m) {
            return query(o * 2, l, m, R);
        }
        return max(t[o * 2], query(o * 2 + 1, m + 1, r, R));
    }

public:
    vector<bool> getResults(vector<vector<int>>& queries) {
        int m = 0;
        for (auto& q : queries) {
            m = max(m, q[1]);
        }
        m++;

        set<int> st{0, m}; // 哨兵
        t.resize(2 << (32 - __builtin_clz(m)));

        vector<bool> ans;
        for (auto& q : queries) {
            int x = q[1];
            auto it = st.lower_bound(x);
            int pre = *prev(it); // x 左侧最近障碍物的位置
            if (q[0] == 1) {
                int nxt = *it; // x 右侧最近障碍物的位置
                st.insert(x);
                update(1, 0, m, x, x - pre);   // 更新 d[x] = x - pre
                update(1, 0, m, nxt, nxt - x); // 更新 d[nxt] = nxt - x
            } else {
                // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                int max_gap = max(query(1, 0, m, pre), x - pre);
                ans.push_back(max_gap >= q[2]);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type seg []int

// 把 i 处的值改成 val
func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o] = val
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 查询 [0,R] 中的最大值
func (t seg) query(o, l, r, R int) int {
	if r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, R)
	}
	return max(t[o<<1], t.query(o<<1|1, m+1, r, R))
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	for _, q := range queries {
		m = max(m, q[1])
	}
	m++

	set := redblacktree.New[int, struct{}]()
	set.Put(0, struct{}{}) // 哨兵
	set.Put(m, struct{}{})
	t := make(seg, 2<<bits.Len(uint(m)))

	for _, q := range queries {
		x := q[1]
        pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			nxt, _ := set.Ceiling(x) // x 右侧最近障碍物的位置
			set.Put(x, struct{}{})
			t.update(1, 0, m, x, x-pre.Key)       // 更新 d[x] = x - pre
			t.update(1, 0, m, nxt.Key, nxt.Key-x) // 更新 d[nxt] = nxt - x
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.query(1, 0, m, pre.Key), x-pre.Key)
			ans = append(ans, maxGap >= q[2])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(U + q\log U)$，其中 $q$ 是 $\textit{queries}$ 的长度，$U$ 是 $\textit{x}$ 的最大值。注意题目保证 $U\le 3q$。
- 空间复杂度：$\mathcal{O}(U)$。

注：如果要做到值域无关，可以用动态开点线段树。

### 方法二：倒序回答询问+有序集合+树状数组

倒序回答询问，在 $x$ 处添加障碍物就变成删除障碍物了。

设 $x$ 左右最近元素分别为 $\textit{pre}$ 和 $\textit{nxt}$。

删除后，需要把 $d[\textit{nxt}]$ 更新为 $\textit{nxt}-\textit{pre}$。注意这会让 $d[\textit{nxt}]$ 增加。

由于询问的是前缀，并且 $d$ 不会变小，所以可以用树状数组维护 $d$。

```py [sol-Python3]
from sortedcontainers import SortedList

class Solution:
    def getResults(self, queries: List[List[int]]) -> List[bool]:
        m = max(q[1] for q in queries) + 1
        t = [0] * m

        def update(i: int, val: int) -> None:
            while i < m:
                t[i] = max(t[i], val)
                i += i & -i

        def pre_max(i: int) -> int:
            res = 0
            while i:
                res = max(res, t[i])
                i &= i - 1
            return res

        pos = [0] + sorted(q[1] for q in queries if q[0] == 1)
        for p, q in pairwise(pos):
            update(q, q - p)
        sl = SortedList(pos)
        sl.add(m)  # 哨兵

        ans = []
        for q in reversed(queries):
            x = q[1]
            i = sl.bisect_left(x)
            pre = sl[i - 1]  # x 左侧最近障碍物的位置
            if q[0] == 1:
                sl.discard(x)
                nxt = sl[i]  # x 右侧最近障碍物的位置
                update(nxt, nxt - pre)  # 更新 d[nxt] = nxt - pre
            else:
                # 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                max_gap = max(pre_max(pre), x - pre)
                ans.append(max_gap >= q[2])
        return ans[::-1]
```

```java [sol-Java]
class Fenwick {
    private final int[] tree;

    public Fenwick(int size) {
        tree = new int[size];
    }

    public void update(int i, int val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] = Math.max(tree[i], val);
        }
    }

    public int preMax(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res = Math.max(res, tree[i]);
        }
        return res;
    }
}

class Solution {
    public static List<Boolean> getResults(int[][] queries) {
        int m = 0;
        List<Integer> pos = new ArrayList<>();
        pos.add(0);
        for (int[] q : queries) {
            m = Math.max(m, q[1]);
            if (q[0] == 1) {
                pos.add(q[1]);
            }
        }
        m++;
        Collections.sort(pos);

        TreeSet<Integer> set = new TreeSet<>(pos);
        set.add(m); // 哨兵
        Fenwick t = new Fenwick(m);
        for (int i = 1; i < pos.size(); i++) {
            t.update(pos.get(i), pos.get(i) - pos.get(i - 1));
        }

        List<Boolean> ans = new ArrayList<>();
        for (int i = queries.length - 1; i >= 0; i--) {
            int[] q = queries[i];
            int x = q[1];
            int pre = set.floor(x - 1); // x 左侧最近障碍物的位置
            if (q[0] == 1) {
                set.remove(x);
                int nxt = set.ceiling(x); // x 右侧最近障碍物的位置
                t.update(nxt, nxt - pre); // 更新 d[nxt] = nxt - pre
            } else {
                // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                int maxGap = Math.max(t.preMax(pre), x - pre);
                ans.add(maxGap >= q[2]);
            }
        }
        Collections.reverse(ans);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> t;

    void update(int i, int val) {
        for (; i < t.size(); i += i & -i) {
            t[i] = max(t[i], val);
        }
    }

    int pre_max(int i) {
        int res = 0;
        for (; i; i &= i - 1) {
            res = max(res, t[i]);
        }
        return res;
    }

public:
    vector<bool> getResults(vector<vector<int>>& queries) {
        int m = 0;
        vector<int> pos = {0};
        for (auto& q : queries) {
            m = max(m, q[1]);
            if (q[0] == 1) {
                pos.push_back(q[1]);
            }
        }
        m++;
        ranges::sort(pos);

        set<int> st(pos.begin(), pos.end());
        st.insert(m); // 哨兵
        t.resize(m);
        for (int i = 1; i < pos.size(); i++) {
            update(pos[i], pos[i] - pos[i - 1]);
        }

        vector<bool> ans;
        for (int i = queries.size() - 1; i >= 0; i--) {
            auto& q = queries[i];
            int x = q[1];
            auto it = st.lower_bound(x);
            int pre = *prev(it); // x 左侧最近障碍物的位置
            if (q[0] == 1) {
                int nxt = *next(it); // x 右侧最近障碍物的位置
                st.erase(it);
                update(nxt, nxt - pre); // 更新 d[nxt] = nxt - pre
            } else {
                // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                int max_gap = max(pre_max(pre), x - pre);
                ans.push_back(max_gap >= q[2]);
            }
        }
        reverse(ans.begin(), ans.end());
        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	set := redblacktree.New[int, struct{}]()
	set.Put(0, struct{}{})
	pos := []int{0}
	for _, q := range queries {
		m = max(m, q[1])
		if q[0] == 1 {
			set.Put(q[1], struct{}{})
			pos = append(pos, q[1])
		}
	}
	m++
	set.Put(m, struct{}{}) // 哨兵

	t := make(fenwick, m)
	slices.Sort(pos)
	for i := 1; i < len(pos); i++ {
		t.update(pos[i], pos[i]-pos[i-1])
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		x := q[1]
		pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			set.Remove(x)
			nxt, _ := set.Ceiling(x) // x 右侧最近障碍物的位置
			t.update(nxt.Key, nxt.Key-pre.Key) // 更新 d[nxt] = nxt - pre
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.preMax(pre.Key), x-pre.Key)
			ans = append(ans, maxGap >= q[2])
		}
	}
	slices.Reverse(ans)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(U + q\log U)$，其中 $q$ 是 $\textit{queries}$ 的长度，$U$ 是 $\textit{x}$ 的最大值。注意题目保证 $U\le 3q$。
- 空间复杂度：$\mathcal{O}(U)$。

### 方法三：倒序回答询问+并查集+树状数组

用并查集实现方法二中有序集合的删除、查找前驱和查找后继。

需要维护两个并查集，分别记作 $\textit{left}$ 和 $\textit{right}$。

- 删除 $x$：把 $\textit{left}$ 中的 $x$ 指向 $x-1$，把 $\textit{right}$ 中的 $x$ 指向 $x+1$。如果要批量删除开区间 $(p,q)$ 中的元素 $x$，可以把 $\textit{left}$ 中的 $x$ 指向 $p$，把 $\textit{right}$ 中的 $x$ 指向 $q$。
- 查找 $x$ 的前驱：$x-1$ 在 $\textit{left}$ 中的代表元。
- 查找 $x$ 的后继：$x+1$ 在 $\textit{right}$ 中的代表元。如果 $x$ 已经删除，也可以查询 $x$ 在 $\textit{right}$ 中的代表元。

```py [sol-Python3]
class Solution:
    def getResults(self, queries: List[List[int]]) -> List[bool]:
        m = max(q[1] for q in queries) + 1
        t = [0] * m

        def update(i: int, val: int) -> None:
            while i < m:
                t[i] = max(t[i], val)
                i += i & -i

        def pre_max(i: int) -> int:
            res = 0
            while i:
                res = max(res, t[i])
                i &= i - 1
            return res

        left = list(range(m + 1))
        right = left.copy()

        def find(fa: List[int], x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa, fa[x])
            return fa[x]

        pos = [0] + sorted(q[1] for q in queries if q[0] == 1)
        for p, q in pairwise(pos):
            update(q, q - p)
            for j in range(p + 1, q):
                left[j] = p  # 删除 j
                right[j] = q
        for j in range(pos[-1] + 1, m):
            left[j] = pos[-1]  # 删除 j
            right[j] = m

        ans = []
        for q in reversed(queries):
            x = q[1]
            pre = find(left, x - 1)  # x 左侧最近障碍物的位置
            if q[0] == 1:
                left[x] = x - 1  # 删除 x
                right[x] = x + 1
                nxt = find(right, x)  # x 右侧最近障碍物的位置
                update(nxt, nxt - pre)  # 更新 d[nxt] = nxt - pre
            else:
                # 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                max_gap = max(pre_max(pre), x - pre)
                ans.append(max_gap >= q[2])
        return ans[::-1]
```

```java [sol-Java]
class Fenwick {
    private final int[] tree;

    public Fenwick(int size) {
        tree = new int[size];
    }

    public void update(int i, int val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] = Math.max(tree[i], val);
        }
    }

    public int preMax(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res = Math.max(res, tree[i]);
        }
        return res;
    }
}

class UnionFind {
    public final int[] fa;

    public UnionFind(int size) {
        fa = new int[size];
        for (int i = 1; i < size; i++) {
            fa[i] = i;
        }
    }

    public int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]);
        }
        return fa[x];
    }
}

class Solution {
    public static List<Boolean> getResults(int[][] queries) {
        int m = 0;
        List<Integer> pos = new ArrayList<>();
        pos.add(0);
        for (int[] q : queries) {
            m = Math.max(m, q[1]);
            if (q[0] == 1) {
                pos.add(q[1]);
            }
        }
        m++;

        UnionFind left = new UnionFind(m + 1);
        UnionFind right = new UnionFind(m + 1);
        Fenwick t = new Fenwick(m);
        Collections.sort(pos);
        for (int i = 1; i < pos.size(); i++) {
            int p = pos.get(i - 1);
            int q = pos.get(i);
            t.update(q, q - p);
            for (int j = p + 1; j < q; j++) {
                left.fa[j] = p; // 删除 j
                right.fa[j] = q;
            }
        }
        for (int j = pos.get(pos.size() - 1) + 1; j < m; j++) {
            left.fa[j] = pos.get(pos.size() - 1); // 删除 j
            right.fa[j] = m;
        }

        List<Boolean> ans = new ArrayList<>();
        for (int i = queries.length - 1; i >= 0; i--) {
            int[] q = queries[i];
            int x = q[1];
            int pre = left.find(x - 1); // x 左侧最近障碍物的位置
            if (q[0] == 1) {
                left.fa[x] = x - 1; // 删除 x
                right.fa[x] = x + 1;
                int nxt = right.find(x); // x 右侧最近障碍物的位置
                t.update(nxt, nxt - pre); // 更新 d[nxt] = nxt - pre
            } else {
                // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                int maxGap = Math.max(t.preMax(pre), x - pre);
                ans.add(maxGap >= q[2]);
            }
        }
        Collections.reverse(ans);
        return ans;
    }
}
```

```cpp [sol-C++]
class UnionFind {
public:
    vector<int> fa;

    UnionFind(int size) : fa(size) {
        for (int i = 1; i < size; i++) {
            fa[i] = i;
        }
    }

    int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]);
        }
        return fa[x];
    }
};

class Solution {
    vector<int> t;

    void update(int i, int val) {
        for (; i < t.size(); i += i & -i) {
            t[i] = max(t[i], val);
        }
    }

    int pre_max(int i) {
        int res = 0;
        for (; i; i &= i - 1) {
            res = max(res, t[i]);
        }
        return res;
    }

public:
    vector<bool> getResults(vector<vector<int>>& queries) {
        int m = 0;
        vector<int> pos = {0};
        for (auto& q : queries) {
            m = max(m, q[1]);
            if (q[0] == 1) {
                pos.push_back(q[1]);
            }
        }
        m++;
        ranges::sort(pos);

        UnionFind left(m + 1);
        UnionFind right(m + 1);
        t.resize(m);
        for (int i = 1; i < pos.size(); i++) {
            int p = pos[i - 1];
            int q = pos[i];
            update(q, q - p);
            for (int j = p + 1; j < q; j++) {
                left.fa[j] = p; // 删除 j
                right.fa[j] = q;
            }
        }
        for (int j = pos.back() + 1; j < m; j++) {
            left.fa[j] = pos.back(); // 删除 j
            right.fa[j] = m;
        }

        vector<bool> ans;
        for (int i = queries.size() - 1; i >= 0; i--) {
            auto& q = queries[i];
            int x = q[1];
            int pre = left.find(x - 1); // x 左侧最近障碍物的位置
            if (q[0] == 1) {
                left.fa[x] = x - 1; // 删除 x
                right.fa[x] = x + 1;
                int nxt = right.find(x); // x 右侧最近障碍物的位置
                update(nxt, nxt - pre); // 更新 d[nxt] = nxt - pre
            } else {
                // 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
                int max_gap = max(pre_max(pre), x - pre);
                ans.push_back(max_gap >= q[2]);
            }
        }
        reverse(ans.begin(), ans.end());
        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

type uf []int

func (f uf) find(x int) int {
	if f[x] != x {
		f[x] = f.find(f[x])
	}
	return f[x]
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	pos := []int{0}
	for _, q := range queries {
		m = max(m, q[1])
		if q[0] == 1 {
			pos = append(pos, q[1])
		}
	}
	m++

	left := make(uf, m+1)
	right := make(uf, m+1)
	for i := range left {
		left[i] = i
		right[i] = i
	}
	t := make(fenwick, m)
	slices.Sort(pos)
	for i := 1; i < len(pos); i++ {
		p, q := pos[i-1], pos[i]
		t.update(q, q-p)
		for j := p + 1; j < q; j++ {
			left[j] = p // 删除 j
			right[j] = q
		}
	}
	for j := pos[len(pos)-1] + 1; j < m; j++ {
		left[j] = pos[len(pos)-1] // 删除 j
		right[j] = m
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		x := q[1]
		pre := left.find(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			left[x] = x - 1 // 删除 x
			right[x] = x + 1
			nxt := right.find(x)   // x 右侧最近障碍物的位置
			t.update(nxt, nxt-pre) // 更新 d[nxt] = nxt - pre
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.preMax(pre), x-pre)
			ans = append(ans, maxGap >= q[2])
		}
	}
	slices.Reverse(ans)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(U + q\log U)$，其中 $q$ 是 $\textit{queries}$ 的长度，$U$ 是 $\textit{x}$ 的最大值。注意题目保证 $U\le 3q$。
- 空间复杂度：$\mathcal{O}(U)$。

## 思考题

改成在 $[\textit{left},x]$ 区间内询问，要怎么做？

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
