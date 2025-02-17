**前置题目**：[850. 矩形面积 II](https://leetcode.cn/problems/rectangle-area-ii/)，[我的题解](https://leetcode.cn/problems/rectangle-area-ii/solutions/3078272/lazy-xian-duan-shu-sao-miao-xian-pythonj-4tkr/)。

首先用 850 题的扫描线方法，求出所有正方形的面积并 $\textit{totArea}$。

然后再次扫描，设扫描线下方的面积和为 $\textit{area}$，那么扫描线上方的面积和为 $\textit{totArea} - \textit{area}$。

题目要求

$$
\textit{area} = \textit{totArea} - \textit{area}
$$

即

$$
\textit{area}\cdot 2 = \textit{totArea}
$$

设当前扫描线的纵坐标为 $y$，下一个需要经过的正方形上/下边界的纵坐标为 $y'$，被至少一个正方形覆盖的底边长之和为 $\textit{sumLen}$，那么新的面积和为

$$
\textit{area} + \textit{sumLen} \cdot (y'-y)
$$

如果发现

$$
(\textit{area} + \textit{sumLen} \cdot (y'-y))\cdot 2 \ge \textit{totArea}
$$

取等号，解得

$$
y' = y + \dfrac{\textit{totalArea}/2 - \textit{area}}{\textit{sumL}} = y + \dfrac{\textit{totalArea} - \textit{area}\cdot 2}{\textit{sumL}\cdot 2}
$$

即为答案。

**编程技巧**：把第一次扫描过程中的关键数据 $\textit{area}$ 和 $\textit{sumLen}$ 记录到一个数组中，然后遍历数组（或者二分），这样可以避免跑两遍线段树（空间换时间）。

```py [sol-Python3]
class Node:
    __slots__ = 'l', 'r', 'min_cover_len', 'min_cover', 'todo'

    def __init__(self):
        self.l = 0
        self.r = 0
        self.min_cover_len = 0  # 区间内被矩形覆盖次数最少的底边长之和
        self.min_cover = 0  # 区间内被矩形覆盖的最小次数
        self.todo = 0  # 子树内的所有节点的 min_cover 需要增加的量，注意这可以是负数


class SegmentTree:
    def __init__(self, xs: List[int]):
        n = len(xs) - 1  # xs.size() 个横坐标有 xs.size()-1 个差值
        self.seg = [Node() for _ in range(2 << (n - 1).bit_length())]
        self.build(xs, 1, 0, n - 1)

    def get_uncovered_length(self) -> int:
        return 0 if self.seg[1].min_cover else self.seg[1].min_cover_len

    # 根据左右儿子的信息，更新当前节点的信息
    def maintain(self, o: int) -> None:
        lo = self.seg[o * 2]
        ro = self.seg[o * 2 + 1]
        mn = min(lo.min_cover, ro.min_cover)
        self.seg[o].min_cover = mn
        # 只统计等于 min_cover 的底边长之和
        self.seg[o].min_cover_len = (lo.min_cover_len if lo.min_cover == mn else 0) + \
                                    (ro.min_cover_len if ro.min_cover == mn else 0)

    # 仅更新节点信息，不下传懒标记 todo
    def do(self, o: int, v: int) -> None:
        self.seg[o].min_cover += v
        self.seg[o].todo += v

    # 下传懒标记 todo
    def spread(self, o: int) -> None:
        v = self.seg[o].todo
        if v:
            self.do(o * 2, v)
            self.do(o * 2 + 1, v)
            self.seg[o].todo = 0

    # 建树
    def build(self, xs: List[int], o: int, l: int, r: int) -> None:
        self.seg[o].l = l
        self.seg[o].r = r
        if l == r:
            self.seg[o].min_cover_len = xs[l + 1] - xs[l]
            return
        m = (l + r) // 2
        self.build(xs, o * 2, l, m)
        self.build(xs, o * 2 + 1, m + 1, r)
        self.maintain(o)

    # 区间更新
    def update(self, o: int, l: int, r: int, v: int) -> None:
        if l <= self.seg[o].l and self.seg[o].r <= r:
            self.do(o, v)
            return
        self.spread(o)
        m = (self.seg[o].l + self.seg[o].r) // 2
        if l <= m:
            self.update(o * 2, l, r, v)
        if m < r:
            self.update(o * 2 + 1, l, r, v)
        self.maintain(o)


# 代码逻辑同 850 题，增加一个 records 数组记录关键数据
class Solution:
    def separateSquares(self, squares: List[List[int]]) -> float:
        xs = []
        events = []
        for (lx, y, l) in squares:
            rx = lx + l
            xs.append(lx)
            xs.append(rx)
            events.append((y, lx, rx, 1))
            events.append((y + l, lx, rx, -1))

        # 排序，方便离散化
        xs = sorted(set(xs))

        # 初始化线段树
        t = SegmentTree(xs)

        # 模拟扫描线从下往上移动
        events.sort(key=lambda e: e[0])
        records = []
        tot_area = 0
        for (y, lx, rx, delta), e2 in pairwise(events):
            l = bisect_left(xs, lx)  # 离散化
            r = bisect_left(xs, rx) - 1  # r 对应着 xs[r] 与 xs[r+1]=rx 的差值
            t.update(1, l, r, delta)  # 更新被 [lx, rx] 覆盖的次数
            sum_len = xs[-1] - xs[0] - t.get_uncovered_length()  # 减去没被矩形覆盖的长度
            records.append((tot_area, sum_len))  # 记录关键数据
            tot_area += sum_len * (e2[0] - y)  # 新增面积 = 被至少一个矩形覆盖的底边长之和 * 矩形高度

        # 二分找最后一个 < tot_area / 2 的面积
        i = bisect_left(records, tot_area, key=lambda r: r[0] * 2) - 1
        area, sum_len = records[i]
        return events[i][0] + (tot_area - area * 2) / (sum_len * 2)
```

```java [sol-Java]
class SegmentTree {
    private final int n;
    private final int[] minCoverLen; // 区间内被矩形覆盖次数最少的底边长之和
    private final int[] minCover;    // 区间内被矩形覆盖的最小次数
    private final int[] todo;        // 子树内的所有节点的 minCover 需要增加的量，注意这可以是负数

    public SegmentTree(int[] xs) {
        n = xs.length - 1; // xs.length 个横坐标有 xs.length-1 个差值
        int size = 2 << (32 - Integer.numberOfLeadingZeros(n - 1));
        minCoverLen = new int[size];
        minCover = new int[size];
        todo = new int[size];
        build(xs, 1, 0, n - 1);
    }

    public void update(int l, int r, int v) {
        update(1, 0, n - 1, l, r, v);
    }

    public int getUncoveredLength() {
        return minCover[1] == 0 ? minCoverLen[1] : 0;
    }

    // 根据左右儿子的信息，更新当前节点的信息
    private void maintain(int o) {
        int mn = Math.min(minCover[o * 2], minCover[o * 2 + 1]);
        minCover[o] = mn;
        // 只统计等于 mn 的底边长之和
        minCoverLen[o] = (minCover[o * 2] == mn ? minCoverLen[o * 2] : 0) +
                (minCover[o * 2 + 1] == mn ? minCoverLen[o * 2 + 1] : 0);
    }

    // 仅更新节点信息，不下传懒标记 todo
    private void do_(int o, int v) {
        minCover[o] += v;
        todo[o] += v;
    }

    // 下传懒标记 todo
    private void spread(int o) {
        if (todo[o] != 0) {
            do_(o * 2, todo[o]);
            do_(o * 2 + 1, todo[o]);
            todo[o] = 0;
        }
    }

    // 建树
    private void build(int[] xs, int o, int l, int r) {
        if (l == r) {
            minCoverLen[o] = xs[l + 1] - xs[l];
            return;
        }
        int m = (l + r) / 2;
        build(xs, o * 2, l, m);
        build(xs, o * 2 + 1, m + 1, r);
        maintain(o);
    }

    // 区间更新
    private void update(int o, int l, int r, int ql, int qr, int v) {
        if (ql <= l && r <= qr) {
            do_(o, v);
            return;
        }
        spread(o);
        int m = (l + r) / 2;
        if (ql <= m) {
            update(o * 2, l, m, ql, qr, v);
        }
        if (m < qr) {
            update(o * 2 + 1, m + 1, r, ql, qr, v);
        }
        maintain(o);
    }
}

// 代码逻辑同 850 题，增加一个 records 数组记录关键数据
class Solution {
    private record Event(int y, int lx, int rx, int delta) {
    }

    private record Record(long area, int sumLen) {
    }

    public double separateSquares(int[][] squares) {
        int n = squares.length * 2;
        int[] xs = new int[n];
        Event[] events = new Event[n];
        n = 0;
        for (int[] sq : squares) {
            int lx = sq[0];
            int y = sq[1];
            int l = sq[2];
            int rx = lx + l;
            xs[n] = lx;
            xs[n + 1] = rx;
            events[n++] = new Event(y, lx, rx, 1);
            events[n++] = new Event(y + l, lx, rx, -1);
        }

        // 排序，方便离散化
        Arrays.sort(xs);

        // 初始化线段树
        SegmentTree t = new SegmentTree(xs);

        // 模拟扫描线从下往上移动
        Arrays.sort(events, (a, b) -> a.y - b.y);
        Record records[] = new Record[n - 1];
        long totArea = 0;
        for (int i = 0; i < n - 1; i++) {
            Event e = events[i];
            int l = Arrays.binarySearch(xs, e.lx); // 离散化
            int r = Arrays.binarySearch(xs, e.rx) - 1; // r 对应着 xs[r] 与 xs[r+1]=rx 的差值
            t.update(l, r, e.delta); // 更新被 [lx, rx] 覆盖的次数
            int sumLen = xs[n - 1] - xs[0] - t.getUncoveredLength(); // 减去没被矩形覆盖的长度
            records[i] = new Record(totArea, sumLen);
            totArea += (long) sumLen * (events[i + 1].y - e.y); // 新增面积 = 被至少一个矩形覆盖的底边长之和 * 矩形高度
        }

        // 找最后一个 < totArea / 2 的面积
        int i = 0;
        while (i < n - 1 && records[i].area * 2 < totArea) {
            i++;
        }
        i--;
        return events[i].y + (totArea - records[i].area * 2) / (records[i].sumLen * 2.0);
    }
}
```

```cpp [sol-C++]
class SegmentTree {
public:
    SegmentTree(vector<int>& xs) {
        unsigned n = xs.size() - 1; // xs.size() 个横坐标有 xs.size()-1 个差值
        seg.resize(2 << bit_width(n - 1));
        build(xs, 1, 0, n - 1);
    }

    void update(int l, int r, int v) {
        update(1, l, r, v);
    }

    int get_uncovered_length() {
        return seg[1].min_cover ? 0 : seg[1].min_cover_len;
    }

private:
    struct Node {
        int l, r;
        int min_cover_len = 0; // 区间内被矩形覆盖次数最少的底边长之和
        int min_cover = 0;     // 区间内被矩形覆盖的最小次数
        int todo = 0;          // 子树内的所有节点的 min_cover 需要增加的量，注意这可以是负数
    };

    vector<Node> seg;

    // 根据左右儿子的信息，更新当前节点的信息
    void maintain(int o) {
        Node& lo = seg[o * 2];
        Node& ro = seg[o * 2 + 1];
        int mn = min(lo.min_cover, ro.min_cover);
        seg[o].min_cover = mn;
        // 只统计等于 min_cover 的底边长之和
        seg[o].min_cover_len = (lo.min_cover == mn ? lo.min_cover_len : 0) +
                               (ro.min_cover == mn ? ro.min_cover_len : 0);
    }

    // 仅更新节点信息，不下传懒标记 todo
    void do_(int o, int v) {
        seg[o].min_cover += v;
        seg[o].todo += v;
    }

    // 下传懒标记 todo
    void spread(int o) {
        int& v = seg[o].todo;
        if (v != 0) {
            do_(o * 2, v);
            do_(o * 2 + 1, v);
            v = 0;
        }
    }

    // 建树
    void build(vector<int>& xs, int o, int l, int r) {
        seg[o].l = l;
        seg[o].r = r;
        if (l == r) {
            seg[o].min_cover_len = xs[l + 1] - xs[l];
            return;
        }
        int m = (l + r) / 2;
        build(xs, o * 2, l, m);
        build(xs, o * 2 + 1, m + 1, r);
        maintain(o);
    }

    // 区间更新
    void update(int o, int l, int r, int v) {
        if (l <= seg[o].l && seg[o].r <= r) {
            do_(o, v);
            return;
        }
        spread(o);
        int m = (seg[o].l + seg[o].r) / 2;
        if (l <= m) {
            update(o * 2, l, r, v);
        }
        if (m < r) {
            update(o * 2 + 1, l, r, v);
        }
        maintain(o);
    }
};

class Solution {
public:
    double separateSquares(vector<vector<int>>& squares) {
        vector<int> xs;
        struct Event { int y, lx, rx, delta; };
        vector<Event> events;
        for (auto& sq : squares) {
            int lx = sq[0], y = sq[1], l = sq[2];
            int rx = lx + l;
            xs.push_back(lx);
            xs.push_back(rx);
            events.emplace_back(y, lx, rx, 1);
            events.emplace_back(y + l, lx, rx, -1);
        }

        // 排序去重，方便离散化
        ranges::sort(xs);
        xs.erase(unique(xs.begin(), xs.end()), xs.end());

        // 初始化线段树
        SegmentTree t(xs);

        // 模拟扫描线从下往上移动
        ranges::sort(events, {}, &Event::y);
        vector<pair<long long, int>> records(events.size() - 1);
        long long tot_area = 0;
        for (int i = 0; i + 1 < events.size(); i++) {
            auto& [y, lx, rx, delta] = events[i];
            int l = ranges::lower_bound(xs, lx) - xs.begin(); // 离散化
            int r = ranges::lower_bound(xs, rx) - xs.begin() - 1; // r 对应着 xs[r] 与 xs[r+1]=rx 的差值
            t.update(l, r, delta); // 更新被 [lx, rx] 覆盖的次数
            int sum_len = xs.back() - xs[0] - t.get_uncovered_length(); // 减去没被矩形覆盖的长度
            records[i] = {tot_area, sum_len};
            tot_area += 1LL * sum_len * (events[i + 1].y - y); // 新增面积 = 被至少一个矩形覆盖的底边长之和 * 矩形高度
        }

        // 二分找最后一个 < tot_area / 2 的面积
        int i = ranges::lower_bound(records, tot_area, {}, [](auto& p) { return p.first * 2; }) - records.begin() - 1;
        auto [area, sum_len] = records[i];
        return events[i].y + (tot_area - area * 2) / (sum_len * 2.0);
    }
};
```

```go [sol-Go]
// 线段树每个节点维护一段横坐标区间 [lx, rx]
type seg []struct {
    l, r        int
    minCoverLen int // 区间内被矩形覆盖次数最少的底边长之和
    minCover    int // 区间内被矩形覆盖的最小次数
    todo        int // 子树内的所有节点的 minCover 需要增加的量，注意这可以是负数
}

// 根据左右儿子的信息，更新当前节点的信息
func (t seg) maintain(o int) {
    lo, ro := &t[o<<1], &t[o<<1|1]
    mn := min(lo.minCover, ro.minCover)
    t[o].minCover = mn
    t[o].minCoverLen = 0
    if lo.minCover == mn { // 只统计等于 minCover 的底边长之和
        t[o].minCoverLen = lo.minCoverLen
    }
    if ro.minCover == mn {
        t[o].minCoverLen += ro.minCoverLen
    }
}

// 仅更新节点信息，不下传懒标记 todo
func (t seg) do(o, v int) {
    t[o].minCover += v
    t[o].todo += v
}

// 下传懒标记 todo
func (t seg) spread(o int) {
    v := t[o].todo
    if v != 0 {
        t.do(o<<1, v)
        t.do(o<<1|1, v)
        t[o].todo = 0
    }
}

// 建树
func (t seg) build(xs []int, o, l, r int) {
    t[o].l, t[o].r = l, r
    t[o].todo = 0
    if l == r {
        t[o].minCover = 0
        t[o].minCoverLen = xs[l+1] - xs[l]
        return
    }
    m := (l + r) >> 1
    t.build(xs, o<<1, l, m)
    t.build(xs, o<<1|1, m+1, r)
    t.maintain(o)
}

// 区间更新
func (t seg) update(o, l, r, v int) {
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
    t.maintain(o)
}

// 代码逻辑同 850 题，增加一个 records 数组记录关键数据
func separateSquares(squares [][]int) float64 {
    m := len(squares) * 2
    xs := make([]int, 0, m)
    type event struct{ y, lx, rx, delta int }
    events := make([]event, 0, m)
    for _, sq := range squares {
        lx, y, l := sq[0], sq[1], sq[2]
        rx := lx + l
        xs = append(xs, lx, rx)
        events = append(events, event{y, lx, rx, 1}, event{y + l, lx, rx, -1})
    }

    // 排序去重，方便离散化
    slices.Sort(xs)
    xs = slices.Compact(xs)

    // 初始化线段树
    n := len(xs) - 1 // len(xs) 个横坐标有 len(xs)-1 个差值
    t := make(seg, 2<<bits.Len(uint(n-1)))
    t.build(xs, 1, 0, n-1)

    // 模拟扫描线从下往上移动
    slices.SortFunc(events, func(a, b event) int { return a.y - b.y })
    type pair struct{ area, sumLen int }
    records := make([]pair, m-1)
    totArea := 0
    for i, e := range events[:m-1] {
        l := sort.SearchInts(xs, e.lx)
        r := sort.SearchInts(xs, e.rx) - 1 // 注意 r 对应着 xs[r] 与 xs[r+1]=e.rx 的差值
        t.update(1, l, r, e.delta)         // 更新被 [e.lx, e.rx] 覆盖的次数
        sumLen := xs[len(xs)-1] - xs[0]    // 总的底边长度
        if t[1].minCover == 0 {            // 需要去掉没被矩形覆盖的长度
            sumLen -= t[1].minCoverLen
        }
        records[i] = pair{totArea, sumLen} // 记录关键数据
        totArea += sumLen * (events[i+1].y - e.y) // 新增面积 = 被至少一个矩形覆盖的底边长之和 * 矩形高度
    }

    // 二分找最后一个 < totArea / 2 的面积
    i := sort.Search(m-1, func(i int) bool { return records[i].area*2 >= totArea }) - 1
    return float64(events[i].y) + float64(totArea-records[i].area*2)/float64(records[i].sumLen*2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{squares}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面数据结构题单中的「**§8.4 Lazy 线段树**」。

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
