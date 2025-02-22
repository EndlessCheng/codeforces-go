## 方法一：模拟

- 初始化：创建一个大小为 $n$ 的数组（记作 $a$），初始值全为 $0$。
- $\texttt{allocate}$：找最小的 $i$，满足 $[i,i+\textit{size}-1]$ 全为 $0$，然后把这个区间全部赋值为 $\textit{mID}$（注意题目保证 $\textit{mID}>0$）。如果没有这样的区间，返回 $-1$。怎么找这样的 $i$，可以在遍历 $a$ 的过程中，维护连续为 $0$ 的元素个数 $\textit{free}$。遇到非 $0$ 数字就把 $\textit{free}$ 置为 $0$，否则把 $\textit{free}$ 加一。如果遍历到 $a[i]$ 发现 $\textit{free}=\textit{size}$，则要找的区间为 $[i-\textit{size}+1,i]$。
- $\texttt{freeMemory}$：遍历 $a$，把所有等于 $\textit{mID}$ 的数置为 $0$，同时统计等于 $\textit{mID}$ 的数的个数，作为答案。

```py [sol-Python3]
class Allocator:
    def __init__(self, n: int):
        self.memory = [0] * n

    def allocate(self, size: int, mID: int) -> int:
        free = 0
        for i, id in enumerate(self.memory):
            if id > 0:  # 已分配
                free = 0  # 重新计数
                continue
            free += 1
            if free == size:  # 找到了
                self.memory[i - size + 1: i + 1] = [mID] * size
                return i - size + 1
        return -1  # 无法分配内存

    def freeMemory(self, mID: int) -> int:
        ans = 0
        for i, id in enumerate(self.memory):
            if id == mID:
                ans += 1
                self.memory[i] = 0  # 标记为空闲内存
        return ans
```

```java [sol-Java]
class Allocator {
    private final int[] memory;

    public Allocator(int n) {
        memory = new int[n];
    }

    public int allocate(int size, int mID) {
        int free = 0;
        for (int i = 0; i < memory.length; i++) {
            if (memory[i] > 0) { // 已分配
                free = 0; // 重新计数
                continue;
            }
            free++;
            if (free == size) { // 找到了
                Arrays.fill(memory, i - size + 1, i + 1, mID);
                return i - size + 1;
            }
        }
        return -1; // 无法分配内存
    }

    public int freeMemory(int mID) {
        int ans = 0;
        for (int i = 0; i < memory.length; i++) {
            if (memory[i] == mID) {
                ans++;
                memory[i] = 0; // 标记为空闲内存
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Allocator {
    vector<int> memory;

public:
    Allocator(int n) : memory(n) {}

    int allocate(int size, int mID) {
        int free = 0;
        for (int i = 0; i < memory.size(); i++) {
            if (memory[i] > 0) { // 已分配
                free = 0; // 重新计数
                continue;
            }
            free++;
            if (free == size) { // 找到了
                fill(memory.begin() + (i - size + 1), memory.begin() + (i + 1), mID);
                return i - size + 1;
            }
        }
        return -1; // 无法分配内存
    }

    int freeMemory(int mID) {
        int ans = 0;
        for (int i = 0; i < memory.size(); i++) {
            if (memory[i] == mID) {
                ans++;
                memory[i] = 0; // 标记为空闲内存
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type Allocator []int

func Constructor(n int) Allocator {
    return make([]int, n)
}

func (a Allocator) Allocate(size, mID int) int {
    free := 0
    for i, id := range a {
        if id > 0 { // 已分配
            free = 0 // 重新计数
            continue
        }
        free++
        if free == size { // 找到了
            for j := i - size + 1; j <= i; j++ {
                a[j] = mID
            }
            return i - size + 1
        }
    }
    return -1 // 无法分配内存
}

func (a Allocator) FreeMemory(mID int) (ans int) {
    for i, id := range a {
        if id == mID {
            ans++
            a[i] = 0 // 标记为空闲内存
        }
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：所有操作均为 $\mathcal{O}(n)$。总体时间复杂度 $\mathcal{O}(qn)$，其中 $q$ 为 $\texttt{allocate}$ 和 $\texttt{freeMemory}$ 的调用次数之和。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：Lazy 线段树 + 线段树二分（选读）

**前置知识**：Lazy 线段树，线段树二分。

**前置题目**：[2213. 由单个字符重复的最长子字符串](https://leetcode.cn/problems/longest-substring-of-one-repeating-character/)（思想和本题类似）

把空闲内存单元记作 $0$，分配的内存单元记作 $1$。

- $\texttt{allocate}$ 等价于找最左边的区间（的左端点），满足区间全为 $0$ 且长度 $\ge \textit{size}$。
- $\texttt{freeMemory}$ 等价于把 $\textit{mID}$ 对应的所有区间置为 $0$。

用线段树维护区间内的最长连续 $0$ 的个数 $\textit{max}_0$，这样方便我们计算 $\texttt{allocate}$。

先说怎么计算 $\textit{max}_0$。

初始值为区间长度，因为一开始所有内存单元都是空闲的。

当前节点的 $\textit{max}_0$，等于以下三者的最大值：

- 左儿子的 $\textit{max}_0$。
- 右儿子的 $\textit{max}_0$。
- 左儿子的后缀连续 $0$ 的个数，加上右儿子的前缀连续 $0$ 的个数。
  
所以需要额外维护区间前缀连续 $0$ 的个数 $\textit{pre}_0$，区间后缀连续 $0$ 的个数 $\textit{suf}_0$。

然后来说怎么计算 $\texttt{allocate}$。

从根节点递归这棵线段树：

- 先找左子树，如果左子树中有符合要求的区间，那么返回这个区间左端点。
- 否则，看看左子树的 $\textit{suf}_0$ 加上右子树的 $\textit{pre}_0$ 是否 $\ge \textit{size}$。如果满足，那么返回区间左端点。
- 否则，递归右子树。

递归边界：

- 如果区间 $\textit{max}_0 < \textit{size}$，无需递归，肯定没有符合要求的区间，返回 $-1$。
- 如果区间长度等于 $1$（这意味着 $\textit{size}=1$），返回区间左端点。

此外，还需要一个哈希表，其 key 为 $\textit{mID}$，value 为列表，记录 $\textit{mID}$ 对应的所有内存区间。

- $\texttt{allocate}$：把 $\textit{mID}$ 及其对应区间加入哈希表。
- $\texttt{freeMemory}$：遍历 $\textit{mID}$ 对应的区间列表，用线段树把区间置为 $0$，同时累加区间长度作为答案。最后把 $\textit{mID}$ 从哈希表中删除。

```py [sol-Python3]
class Node:
    __slots__ = 'pre0', 'suf0', 'max0', 'todo'


class SegTree:
    def __init__(self, n: int) -> None:
        self.n = n
        self.t = [Node() for _ in range(2 << (n - 1).bit_length())]
        self.build(1, 0, n - 1)

    def do(self, i: int, l: int, r: int, v: int) -> None:
        size = 0 if v > 0 else r - l + 1
        self.t[i].pre0 = size
        self.t[i].suf0 = size
        self.t[i].max0 = size
        self.t[i].todo = v

    # 下传懒标记
    def spread(self, o: int, l: int, r: int) -> None:
        v = self.t[o].todo
        if v != -1:
            m = (l + r) // 2
            self.do(o * 2, l, m, v)
            self.do(o * 2 + 1, m + 1, r, v)
            self.t[o].todo = -1

    # 初始化线段树
    def build(self, o: int, l: int, r: int) -> None:
        self.do(o, l, r, -1)
        if l == r:
            return
        m = (l + r) // 2
        self.build(o * 2, l, m)
        self.build(o * 2 + 1, m + 1, r)

    # 把区间 [ql, qr] 都置为 v
    def update(self, o: int, l: int, r: int, ql: int, qr: int, v: int) -> None:
        if ql <= l and r <= qr:
            self.do(o, l, r, v)
            return
        self.spread(o, l, r)
        m = (l + r) // 2
        if ql <= m:
            self.update(o * 2, l, m, ql, qr, v)
        if m < qr:
            self.update(o * 2 + 1, m + 1, r, ql, qr, v)

        # 合并左右子树的信息
        lo = self.t[o * 2]
        ro = self.t[o * 2 + 1]
        # 区间前缀连续 0 的个数
        self.t[o].pre0 = lo.pre0
        if lo.pre0 == m - l + 1:
            self.t[o].pre0 += ro.pre0  # 和右子树的 pre0 拼起来
        # 区间后缀连续 0 的个数
        self.t[o].suf0 = ro.suf0
        if ro.suf0 == r - m:
            self.t[o].suf0 += lo.suf0  # 和左子树的 suf0 拼起来
        # 区间最长连续 0 的个数
        self.t[o].max0 = max(lo.max0, ro.max0, lo.suf0 + ro.pre0)

    # 线段树二分，找最左边的区间左端点，满足区间全为 0 且长度 >= size
    # 如果不存在这样的区间，返回 -1
    def find_first(self, o: int, l: int, r: int, size: int) -> int:
        if self.t[o].max0 < size:
            return -1
        if l == r:
            return l
        self.spread(o, l, r)
        m = (l + r) // 2
        idx = self.find_first(o * 2, l, m, size)  # 递归左子树
        if idx < 0:
            # 左子树的后缀 0 个数 + 右子树的前缀 0 个数 >= size
            if self.t[o * 2].suf0 + self.t[o * 2 + 1].pre0 >= size:
                return m - self.t[o * 2].suf0 + 1
            idx = self.find_first(o * 2 + 1, m + 1, r, size)  # 递归右子树
        return idx


class Allocator:
    def __init__(self, n: int):
        self.n = n
        self.tree = SegTree(n)
        self.blocks = defaultdict(list)

    def allocate(self, size: int, mID: int) -> int:
        i = self.tree.find_first(1, 0, self.n - 1, size)
        if i < 0:  # 无法分配内存
            return -1
        # 分配内存 [i, i+size-1]
        self.blocks[mID].append((i, i + size - 1))
        self.tree.update(1, 0, self.n - 1, i, i + size - 1, 1)
        return i

    def freeMemory(self, mID: int) -> int:
        ans = 0
        for l, r in self.blocks[mID]:
            ans += r - l + 1
            self.tree.update(1, 0, self.n - 1, l, r, 0)  # 释放内存
        del self.blocks[mID]
        return ans
```

```java [sol-Java]
class SegTree {
    private final int[] pre0; // 区间前缀连续 0 的个数
    private final int[] suf0; // 区间后缀连续 0 的个数
    private final int[] max0; // 区间最长连续 0 的个数
    private final int[] todo; // 懒标记

    public SegTree(int n) {
        int size = 2 << (32 - Integer.numberOfLeadingZeros(n - 1));
        pre0 = new int[size];
        suf0 = new int[size];
        max0 = new int[size];
        todo = new int[size];
        build(1, 0, n - 1);
    }

    // 把 [ql, qr] 都置为 v
    public void update(int o, int l, int r, int ql, int qr, int v) {
        if (ql <= l && r <= qr) {
            do_(o, l, r, v);
            return;
        }
        spread(o, l, r);
        int m = (l + r) / 2;
        int lo = o * 2;
        int ro = lo + 1;
        if (ql <= m) {
            update(lo, l, m, ql, qr, v);
        }
        if (m < qr) {
            update(ro, m + 1, r, ql, qr, v);
        }

        // 合并左右子树的信息
        pre0[o] = pre0[lo];
        if (pre0[lo] == m - l + 1) {
            pre0[o] += pre0[ro]; // 和右子树的 pre0 拼起来
        }
        suf0[o] = suf0[ro];
        if (suf0[ro] == r - m) {
            suf0[o] += suf0[lo]; // 和左子树的 suf0 拼起来
        }
        max0[o] = Math.max(Math.max(max0[lo], max0[ro]), suf0[lo] + pre0[ro]);
    }

    // 线段树二分，找最左边的区间左端点，满足区间全为 0 且长度 >= size
    // 如果不存在这样的区间，返回 -1
    public int findFirst(int o, int l, int r, int size) {
        if (max0[o] < size) {
            return -1;
        }
        if (l == r) {
            return l;
        }
        spread(o, l, r);
        int m = (l + r) / 2;
        int lo = o * 2;
        int ro = lo + 1;
        int idx = findFirst(lo, l, m, size); // 递归左子树
        if (idx < 0) {
            // 左子树的后缀 0 个数 + 右子树的前缀 0 个数 >= size
            if (suf0[lo] + pre0[ro] >= size) {
                return m - suf0[lo] + 1;
            }
            idx = findFirst(ro, m + 1, r, size); // 递归右子树
        }
        return idx;
    }

    // 初始化线段树
    private void build(int o, int l, int r) {
        do_(o, l, r, -1);
        if (l == r) {
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m);
        build(o * 2 + 1, m + 1, r);
    }

    private void do_(int i, int l, int r, int v) {
        int size = v <= 0 ? r - l + 1 : 0;
        pre0[i] = suf0[i] = max0[i] = size;
        todo[i] = v;
    }

    // 下传懒标记
    private void spread(int o, int l, int r) {
        int v = todo[o];
        if (v != -1) {
            int m = (l + r) / 2;
            do_(o * 2, l, m, v);
            do_(o * 2 + 1, m + 1, r, v);
            todo[o] = -1;
        }
    }
}

class Allocator {
    private final int n;
    private final SegTree tree;
    private final Map<Integer, List<int[]>> blocks = new HashMap<>();

    public Allocator(int n) {
        this.n = n;
        this.tree = new SegTree(n);
    }

    public int allocate(int size, int mID) {
        int i = tree.findFirst(1, 0, n - 1, size);
        if (i < 0) { // 无法分配内存
            return -1;
        }
        // 分配内存 [i, i+size-1]
        blocks.computeIfAbsent(mID, k -> new ArrayList<>()).add(new int[]{i, i + size - 1});
        tree.update(1, 0, n - 1, i, i + size - 1, 1);
        return i;
    }

    public int freeMemory(int mID) {
        int ans = 0;
        List<int[]> list = blocks.get(mID);
        if (list != null) {
            for (int[] range : list) {
                ans += range[1] - range[0] + 1;
                tree.update(1, 0, n - 1, range[0], range[1], 0); // 释放内存
            }
            blocks.remove(mID);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class SegTree {
    struct Node {
        int pre0, suf0, max0, todo;
    };

    vector<Node> t;

    void do_(int i, int l, int r, int v) {
        auto& o = t[i];
        int size = v <= 0 ? r - l + 1 : 0;
        o.pre0 = o.suf0 = o.max0 = size;
        o.todo = v;
    }

    // 下传懒标记
    void spread(int o, int l, int r) {
        int& v = t[o].todo;
        if (v != -1) {
            int m = (l + r) / 2;
            do_(o * 2, l, m, v);
            do_(o * 2 + 1, m + 1, r, v);
            v = -1;
        }
    }

    // 初始化线段树
    void build(int o, int l, int r) {
        do_(o, l, r, -1);
        if (l == r) {
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m);
        build(o * 2 + 1, m + 1, r);
    }

public:
    SegTree(int n) {
        t.resize(2 << bit_width((unsigned) n - 1));
        build(1, 0, n - 1);
    }

    // 把 [ql, qr] 都置为 v
    void update(int o, int l, int r, int ql, int qr, int v) {
        if (ql <= l && r <= qr) {
            do_(o, l, r, v);
            return;
        }
        spread(o, l, r);
        int m = (l + r) / 2;
        if (ql <= m) {
            update(o * 2, l, m, ql, qr, v);
        }
        if (m < qr) {
            update(o * 2 + 1, m + 1, r, ql, qr, v);
        }

        // 合并左右子树的信息
        Node& lo = t[o * 2];
        Node& ro = t[o * 2 + 1];
        // 区间前缀连续 0 的个数
        t[o].pre0 = lo.pre0;
        if (lo.pre0 == m - l + 1) {
            t[o].pre0 += ro.pre0; // 和右子树的 pre0 拼起来
        }
        // 区间后缀连续 0 的个数
        t[o].suf0 = ro.suf0;
        if (ro.suf0 == r - m) {
            t[o].suf0 += lo.suf0; // 和左子树的 suf0 拼起来
        }
        // 区间最长连续 0 的个数
        t[o].max0 = max({lo.max0, ro.max0, lo.suf0 + ro.pre0});
    }

    // 线段树二分，找最左边的区间左端点，满足区间全为 0 且长度 >= size
    // 如果不存在这样的区间，返回 -1
    int find_first(int o, int l, int r, int size) {
        if (t[o].max0 < size) {
            return -1;
        }
        if (l == r) {
            return l;
        }
        spread(o, l, r);
        int m = (l + r) / 2;
        int idx = find_first(o * 2, l, m, size); // 递归左子树
        if (idx < 0) {
            // 左子树的后缀 0 个数 + 右子树的前缀 0 个数 >= size
            if (t[o * 2].suf0 + t[o * 2 + 1].pre0 >= size) {
                return m - t[o * 2].suf0 + 1;
            }
            idx = find_first(o * 2 + 1, m + 1, r, size); // 递归右子树
        }
        return idx;
    }
};

class Allocator {
    int n;
    SegTree tree;
    unordered_map<int, vector<pair<int, int>>> blocks;

public:
    Allocator(int n) : n(n), tree(n) {}

    int allocate(int size, int mID) {
        int i = tree.find_first(1, 0, n - 1, size);
        if (i < 0) { // 无法分配内存
            return -1;
        }
        blocks[mID].emplace_back(i, i + size - 1);
        tree.update(1, 0, n - 1, i, i + size - 1, 1); // 分配内存 [i, i+size-1]
        return i;
    }

    int freeMemory(int mID) {
        int ans = 0;
        for (auto& [l, r] : blocks[mID]) {
            ans += r - l + 1;
            tree.update(1, 0, n - 1, l, r, 0); // 释放内存
        }
        blocks.erase(mID);
        return ans;
    }
};
```

```go [sol-Go]
type segTree []struct {
    l, r, pre0, suf0, max0, todo int
}

func newSegTree(n int) segTree {
    t := make(segTree, 2<<bits.Len(uint(n-1)))
    t.build(1, 0, n-1)
    return t
}

func (t segTree) do(i, v int) {
    o := &t[i]
    size := 0
    if v <= 0 {
        size = o.r - o.l + 1
    }
    o.pre0 = size
    o.suf0 = size
    o.max0 = size
    o.todo = v
}

// 下传懒标记
func (t segTree) spread(o int) {
    v := t[o].todo
    if v != -1 {
        t.do(o<<1, v)
        t.do(o<<1|1, v)
        t[o].todo = -1
    }
}

// 初始化线段树
func (t segTree) build(o, l, r int) {
    t[o].l, t[o].r = l, r
    t.do(o, -1)
    if l == r {
        return
    }
    m := (l + r) >> 1
    t.build(o<<1, l, m)
    t.build(o<<1|1, m+1, r)
}

// 把 [l, r] 都置为 v
func (t segTree) update(o, l, r, v int) {
    if l <= t[o].l && t[o].r <= r {
        t.do(o, v)
        return
    }
    t.spread(o)
    m := (t[o].l + t[o].r) >> 1
    if l <= m {
        t.update(o<<1, l, r, v)
    }
    if m < r {
        t.update(o<<1|1, l, r, v)
    }

    // 合并左右子树的信息
    lo, ro := t[o<<1], t[o<<1|1]
    // 区间前缀连续 0 的个数
    t[o].pre0 = lo.pre0
    if lo.pre0 == m-t[o].l+1 {
        t[o].pre0 += ro.pre0 // 和右子树的 pre0 拼起来
    }
    // 区间后缀连续 0 的个数
    t[o].suf0 = ro.suf0
    if ro.suf0 == t[o].r-m {
        t[o].suf0 += lo.suf0 // 和左子树的 suf0 拼起来
    }
    // 区间最长连续 0 的个数
    t[o].max0 = max(lo.max0, ro.max0, lo.suf0+ro.pre0)
}

// 线段树二分，找最左边的区间左端点，满足区间全为 0 且长度 >= size
// 如果不存在这样的区间，返回 -1
func (t segTree) findFirst(o, size int) int {
    if t[o].max0 < size {
        return -1
    }
    if t[o].l == t[o].r {
        return t[o].l
    }
    t.spread(o)
    idx := t.findFirst(o<<1, size) // 递归左子树
    if idx < 0 {
        // 左子树的后缀 0 个数 + 右子树的前缀 0 个数 >= size
        if t[o<<1].suf0+t[o<<1|1].pre0 >= size {
            m := (t[o].l + t[o].r) >> 1
            return m - t[o<<1].suf0 + 1
        }
        idx = t.findFirst(o<<1|1, size) // 递归右子树
    }
    return idx
}

// 上面为线段树代码

type interval struct {
    l, r int
}

type Allocator struct {
    tree   segTree
    blocks map[int][]interval
}

func Constructor(n int) Allocator {
    return Allocator{
        tree:   newSegTree(n),
        blocks: map[int][]interval{},
    }
}

func (a Allocator) Allocate(size, mID int) int {
    i := a.tree.findFirst(1, size)
    if i < 0 { // 无法分配内存
        return -1
    }
    a.blocks[mID] = append(a.blocks[mID], interval{i, i + size - 1})
    a.tree.update(1, i, i+size-1, 1) // 分配内存 [i, i+size-1]
    return i
}

func (a Allocator) FreeMemory(mID int) (ans int) {
    for _, p := range a.blocks[mID] {
        ans += p.r - p.l + 1
        a.tree.update(1, p.l, p.r, 0) // 释放内存
    }
    delete(a.blocks, mID)
    return
}
```

#### 复杂度分析

- 时间复杂度：初始化 $\mathcal{O}(n)$。$\texttt{allocate}$ $\mathcal{O}(\log n)$。$\texttt{freeMemory}$ 均摊 $\mathcal{O}(\log n)$，因为释放内存单元的次数不会超过 $\texttt{allocate}$ 的调用次数。总体时间复杂度 $\mathcal{O}(n+q\log n)$，其中 $q$ 为 $\texttt{allocate}$ 和 $\texttt{freeMemory}$ 的调用次数之和。
- 空间复杂度：$\mathcal{O}(n)$。注意哈希表的大小不会超过 $n$。

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
