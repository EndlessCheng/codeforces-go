## 方法一：贪心+暴力

### 提示 1

按照区间右端点从小到大排序。

### 提示 2

排序后，对于区间 $\textit{tasks}[i]$ 来说，它右侧的任务区间要么和它没有交集，要么包含它的一部分**后缀**。

例如排序后的区间为 $[1,5],[3,7],[6,8]$，对于 $[1,5]$ 来说，它右边的区间要么和它没有交集，例如 $[6,8]$；要么交集是 $[1,5]$ 的后缀，例如 $[1,5]$ 和 $[3,7]$ 的交集是 $[3,5]$，这是 $[1,5]$ 的后缀（$3,4,5$ 是 $1,2,3,4,5$ 的后缀）。

### 提示 3

遍历排序后的任务，先统计区间内的已运行的电脑运行时间点，如果个数小于 $\textit{duration}$，则需要新增时间点。根据提示 2，尽量把新增的时间点安排在区间 $[\textit{start},\textit{end}]$ 的后缀上，这样下一个区间就能统计到更多已运行的时间点。

附：[视频讲解](https://www.bilibili.com/video/BV1d54y1M7Qg/) 第四题。

```py [sol-Python3]
class Solution:
    def findMinimumTime(self, tasks: List[List[int]]) -> int:
        tasks.sort(key=lambda t: t[1])
        run = [False] * (tasks[-1][1] + 1)
        for start, end, d in tasks:
            d -= sum(run[start: end + 1])  # 去掉运行中的时间点
            if d <= 0:  # 该任务已完成
                continue
            # 该任务尚未完成，从后往前找没有运行的时间点
            for i in range(end, start - 1, -1):
                if run[i]:  # 已运行
                    continue
                run[i] = True  # 运行
                d -= 1
                if d == 0:
                    break
        return sum(run)
```

```java [sol-Java]
class Solution {
    public int findMinimumTime(int[][] tasks) {
        Arrays.sort(tasks, (a, b) -> a[1] - b[1]);
        int ans = 0;
        int mx = tasks[tasks.length - 1][1];
        boolean[] run = new boolean[mx + 1];
        for (int[] t : tasks) {
            int start = t[0];
            int end = t[1];
            int d = t[2];
            for (int i = start; i <= end; i++) {
                if (run[i]) {
                    d--; // 去掉运行中的时间点
                }
            }
            for (int i = end; d > 0; i--) { // 剩余的 d 填充区间后缀
                if (!run[i]) {
                    run[i] = true; // 运行
                    d--;
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMinimumTime(vector<vector<int>>& tasks) {
        ranges::sort(tasks, [](auto& a, auto& b) { return a[1] < b[1]; });
        int ans = 0;
        vector<int> run(tasks.back()[1] + 1);
        for (auto& t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            d -= reduce(run.begin() + start, run.begin() + end + 1); // 去掉运行中的时间点
            for (int i = end; d > 0; i--) { // 剩余的 d 填充区间后缀
                if (!run[i]) {
                    run[i] = true; // 运行
                    d--;
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMinimumTime(tasks [][]int) (ans int) {
    slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
    run := make([]bool, tasks[len(tasks)-1][1]+1)
    for _, t := range tasks {
        start, end, d := t[0], t[1], t[2]
        for _, b := range run[start : end+1] { // 去掉运行中的时间点
            if b {
                d--
            }
        }
        for i := end; d > 0; i-- { // 剩余的 d 填充区间后缀
            if !run[i] {
                run[i] = true // 运行
                d--
                ans++
            }
        }
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + nU)$，其中 $n$ 为 $\textit{tasks}$ 的长度，$U=\max(\textit{end}_i)$。
- 空间复杂度：$\mathcal{O}(U)$。

## 方法二：线段树优化

**前置知识**：线段树。

在方法一的暴力更新上优化。

由于有区间更新操作，需要用 lazy 线段树，之前在 [双周赛 98](https://www.bilibili.com/video/BV15D4y1G7ms/) 中讲过。

对于本题，在更新的时候需要优先递归右子树，从而保证是从右往左更新。其余细节见代码注释。

> 注：由于线段树常数比较大，在数据范围只有几百几千的情况下，不一定比方法一的暴力快。

```py [sol-Python3]
class Solution:
    def findMinimumTime(self, tasks: List[List[int]]) -> int:
        tasks.sort(key=lambda t: t[1])
        u = tasks[-1][1]
        m = 2 << u.bit_length()
        cnt = [0] * m
        todo = [False] * m

        def do(o: int, l: int, r: int) -> None:
            cnt[o] = r - l + 1
            todo[o] = True

        def spread(o: int, l: int, m: int, r: int) -> None:
            if todo[o]:
                todo[o] = False
                do(o * 2, l, m)
                do(o * 2 + 1, m + 1, r)

        # 查询区间正在运行的时间点 [L,R]   o,l,r=1,1,u
        def query(o: int, l: int, r: int, L: int, R: int) -> int:
            if L <= l and r <= R: return cnt[o]
            m = (l + r) // 2
            spread(o, l, m, r)
            if m >= R: return query(o * 2, l, m, L, R)
            if m < L: return query(o * 2 + 1, m + 1, r, L, R)
            return query(o * 2, l, m, L, R) + query(o * 2 + 1, m + 1, r, L, R)

        # 在区间 [L,R] 的后缀上新增 suffix 个时间点    o,l,r=1,1,u
        # 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log u)
        def update(o: int, l: int, r: int, L: int, R: int) -> None:
            size = r - l + 1
            if cnt[o] == size: return  # 全部为运行中
            nonlocal suffix
            if L <= l and r <= R and size - cnt[o] <= suffix:  # 整个区间全部改为运行中
                suffix -= size - cnt[o]
                do(o, l, r)
                return
            m = (l + r) // 2
            spread(o, l, m, r)
            if m < R: update(o * 2 + 1, m + 1, r, L, R)  # 先更新右子树
            if suffix: update(o * 2, l, m, L, R)  # 再更新左子树（如果还有需要新增的时间点）
            cnt[o] = cnt[o * 2] + cnt[o * 2 + 1]

        for start, end, d in tasks:
            suffix = d - query(1, 1, u, start, end)  # 去掉运行中的时间点
            if suffix > 0: update(1, 1, u, start, end)  # 新增时间点
        return cnt[1]
```

```java [sol-Java]
class Solution {
    public int findMinimumTime(int[][] tasks) {
        Arrays.sort(tasks, (a, b) -> a[1] - b[1]);
        int u = tasks[tasks.length - 1][1];
        int m = 2 << (32 - Integer.numberOfLeadingZeros(u));
        cnt = new int[m];
        todo = new boolean[m];
        for (int[] t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            suffix = d - query(1, 1, u, start, end); // 去掉运行中的时间点
            if (suffix > 0) update(1, 1, u, start, end); // 新增时间点
        }
        return cnt[1];
    }

    private int[] cnt;
    private boolean[] todo;
    private int suffix;

    private void do_(int o, int l, int r) {
        cnt[o] = r - l + 1;
        todo[o] = true;
    }

    private void spread(int o, int l, int m, int r) {
        if (todo[o]) {
            do_(o * 2, l, m);
            do_(o * 2 + 1, m + 1, r);
            todo[o] = false;
        }
    }

    // 查询区间 [L,R]   o,l,r=1,1,u
    private int query(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) return cnt[o];
        int m = (l + r) / 2;
        spread(o, l, m, r);
        if (m >= R) return query(o * 2, l, m, L, R);
        if (m < L) return query(o * 2 + 1, m + 1, r, L, R);
        return query(o * 2, l, m, L, R) + query(o * 2 + 1, m + 1, r, L, R);
    }

    // 新增区间 [L,R] 后缀的 suffix 个时间点    o,l,r=1,1,u
    // 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log u)
    private void update(int o, int l, int r, int L, int R) {
        int size = r - l + 1;
        if (cnt[o] == size) return; // 全部为运行中
        if (L <= l && r <= R && size - cnt[o] <= suffix) { // 整个区间全部改为运行中
            suffix -= size - cnt[o];
            do_(o, l, r);
            return;
        }
        int m = (l + r) / 2;
        spread(o, l, m, r);
        if (m < R) update(o * 2 + 1, m + 1, r, L, R); // 先更新右子树
        if (suffix > 0) update(o * 2, l, m, L, R); // 再更新左子树（如果还有需要新增的时间点）
        cnt[o] = cnt[o * 2] + cnt[o * 2 + 1];
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> cnt, todo;

    void do_(int o, int l, int r) {
        cnt[o] = r - l + 1;
        todo[o] = true;
    }

    void spread(int o, int l, int m, int r) {
        if (todo[o]) {
            do_(o * 2, l, m);
            do_(o * 2 + 1, m + 1, r);
            todo[o] = false;
        }
    }

    // 查询区间 [L,R]   o,l,r=1,1,u
    int query(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) return cnt[o];
        int m = (l + r) / 2;
        spread(o, l, m, r);
        if (m >= R) return query(o * 2, l, m, L, R);
        if (m < L) return query(o * 2 + 1, m + 1, r, L, R);
        return query(o * 2, l, m, L, R) + query(o * 2 + 1, m + 1, r, L, R);
    }

    // 新增区间 [L,R] 后缀的 suffix 个时间点    o,l,r=1,1,u
    // 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log MX)
    void update(int o, int l, int r, int L, int R, int& suffix) {
        int size = r - l + 1;
        if (cnt[o] == size) return; // 全部为运行中
        if (L <= l && r <= R && size - cnt[o] <= suffix) { // 整个区间全部改为运行中
            suffix -= size - cnt[o];
            do_(o, l, r);
            return;
        }
        int m = (l + r) / 2;
        spread(o, l, m, r);
        if (m < R) update(o * 2 + 1, m + 1, r, L, R, suffix); // 先更新右子树
        if (suffix) update(o * 2, l, m, L, R, suffix); // 再更新左子树（如果还有需要新增的时间点）
        cnt[o] = cnt[o * 2] + cnt[o * 2 + 1];
    }

public:
    int findMinimumTime(vector<vector<int>>& tasks) {
        ranges::sort(tasks, [](auto& a, auto& b) { return a[1] < b[1]; });
        int u = tasks.back()[1];
        int m = 2 << (32 - __builtin_clz(u));
        cnt.resize(m);
        todo.resize(m);
        for (auto& t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            d -= query(1, 1, u, start, end);  // 去掉运行中的时间点
            if (d > 0) update(1, 1, u, start, end, d); // 新增时间点
        }
        return cnt[1];
    }
};
```

```go [sol-Go]
type seg []struct {
    l, r, cnt int
    todo      bool
}

func (t seg) build(o, l, r int) {
    t[o].l, t[o].r = l, r
    if l == r {
        return
    }
    m := (l + r) >> 1
    t.build(o<<1, l, m)
    t.build(o<<1|1, m+1, r)
}

func (t seg) do(i int) {
    o := &t[i]
    o.cnt = o.r - o.l + 1
    o.todo = true
}

func (t seg) spread(o int) {
    if t[o].todo {
        t[o].todo = false
        t.do(o << 1)
        t.do(o<<1 | 1)
    }
}

// 查询区间 [l,r]   o=1
func (t seg) query(o, l, r int) int {
    if l <= t[o].l && t[o].r <= r {
        return t[o].cnt
    }
    t.spread(o)
    m := (t[o].l + t[o].r) >> 1
    if r <= m {
        return t.query(o<<1, l, r)
    }
    if m < l {
        return t.query(o<<1|1, l, r)
    }
    return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

// 新增区间 [l,r] 后缀的 suffix 个时间点   o=1
// 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log u)
func (t seg) update(o, l, r int, suffix *int) {
    size := t[o].r - t[o].l + 1
    if t[o].cnt == size { // 全部为运行中
        return
    }
    if l <= t[o].l && t[o].r <= r && size-t[o].cnt <= *suffix { // 整个区间全部改为运行中
        *suffix -= size - t[o].cnt
        t.do(o)
        return
    }
    t.spread(o)
    m := (t[o].l + t[o].r) >> 1
    if r > m { // 先更新右子树
        t.update(o<<1|1, l, r, suffix)
    }
    if *suffix > 0 { // 再更新左子树（如果还有需要新增的时间点）
        t.update(o<<1, l, r, suffix)
    }
    t[o].cnt = t[o<<1].cnt + t[o<<1|1].cnt
}

func findMinimumTime(tasks [][]int) (ans int) {
    slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
    u := tasks[len(tasks)-1][1]
    st := make(seg, 2<<bits.Len(uint(u-1)))
    st.build(1, 1, u)
    for _, t := range tasks {
        start, end, d := t[0], t[1], t[2]
        d -= st.query(1, start, end) // 去掉运行中的时间点
        if d > 0 {
            st.update(1, start, end, &d) // 新增时间点
        }
    }
    return st[1].cnt
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log U)$，其中 $n$ 为 $\textit{tasks}$ 的长度，$U=\max(\textit{end}_i)$。
- 空间复杂度：$\mathcal{O}(U)$。

## 方法三：栈+二分查找

由于每次都是**从右到左**新增时间点，如果把连续的时间点看成闭区间，那么从右到左新增时间点，会把若干**右侧**的区间合并成一个大区间，也就是从 $\textit{end}$ 倒着开始，先合并右边，再合并左边，因此可以用栈来优化。

栈中维护闭区间的左右端点，以及从栈底到栈顶的区间长度之和（类似前缀和）。

由于一旦发现区间相交就立即合并，所以栈中保存的都是**不相交**的区间。

合并前，先尝试在栈中**二分查找**包含左端点 $\textit{start}$ 的区间。由于栈中还保存了区间长度之和，所以可以快速得到 $[\textit{start},\textit{end}]$ 范围内的运行中的时间点个数。

如果还需要新增时间点，那么就从右到左合并，具体细节见代码。

关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

```py [sol-Python3]
class Solution:
    def findMinimumTime(self, tasks: List[List[int]]) -> int:
        tasks.sort(key=lambda t: t[1])
        # 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
        st = [(-2, -2, 0)]  # 哨兵，保证不和任何区间相交
        for start, end, d in tasks:
            _, r, s = st[bisect_left(st, (start,)) - 1]
            d -= st[-1][2] - s  # 去掉运行中的时间点
            if start <= r:  # start 在区间 st[i] 内
                d -= r - start + 1  # 去掉运行中的时间点
            if d <= 0:
                continue
            while end - st[-1][1] <= d:  # 剩余的 d 填充区间后缀
                l, r, _ = st.pop()
                d += r - l + 1  # 合并区间
            st.append((end - d + 1, end, st[-1][2] + d))
        return st[-1][2]
```

```java [sol-Java]
class Solution {
    public int findMinimumTime(int[][] tasks) {
        Arrays.sort(tasks, (a, b) -> a[1] - b[1]);
        // 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
        List<int[]> st = new ArrayList<>();
        st.add(new int[]{-2, -2, 0}); // 哨兵，保证不和任何区间相交
        for (int[] t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            int[] e = st.get(lowerBound(st, start) - 1);
            d -= st.get(st.size() - 1)[2] - e[2]; // 去掉运行中的时间点
            if (start <= e[1]) { // start 在区间 st[i] 内
                d -= e[1] - start + 1; // 去掉运行中的时间点
            }
            if (d <= 0) {
                continue;
            }
            while (end - st.get(st.size() - 1)[1] <= d) { // 剩余的 d 填充区间后缀
                e = st.remove(st.size() - 1);
                d += e[1] - e[0] + 1; // 合并区间
            }
            st.add(new int[]{end - d + 1, end, st.get(st.size() - 1)[2] + d});
        }
        return st.get(st.size() - 1)[2];
    }

    // 开区间二分
    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(List<int[]> st, int target) {
        int left = -1, right = st.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // st[left] < target
            // st[right] >= target
            int mid = (left + right) >>> 1;
            if (st.get(mid)[0] < target) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return right; // 或者 left+1
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMinimumTime(vector<vector<int>>& tasks) {
        ranges::sort(tasks, [](auto& a, auto& b) { return a[1] < b[1]; });
        // 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
        vector<array<int, 3>> st{{-2, -2, 0}}; // 哨兵，保证不和任何区间相交
        for (auto& t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            auto [_, r, s] = *--ranges::lower_bound(st, start, {}, [](auto& x) { return x[0]; });
            d -= st.back()[2] - s; // 去掉运行中的时间点
            if (start <= r) { // start 在区间 st[i] 内
                d -= r - start + 1; // 去掉运行中的时间点
            }
            if (d <= 0) {
                continue;
            }
            while (end - st.back()[1] <= d) { // 剩余的 d 填充区间后缀
                auto [l, r, _] = st.back();
                st.pop_back();
                d += r - l + 1; // 合并区间
            }
            st.push_back({end - d + 1, end, st.back()[2] + d});
        }
        return st.back()[2];
    }
};
```

```go [sol-Go]
func findMinimumTime(tasks [][]int) int {
    slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
    // 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
    type tuple struct{ l, r, s int }
    st := []tuple{{-2, -2, 0}} // 哨兵，保证不和任何区间相交
    for _, p := range tasks {
        start, end, d := p[0], p[1], p[2]
        i := sort.Search(len(st), func(i int) bool { return st[i].l >= start }) - 1
        d -= st[len(st)-1].s - st[i].s // 去掉运行中的时间点
        if start <= st[i].r { // start 在区间 st[i] 内
            d -= st[i].r - start + 1 // 去掉运行中的时间点
        }
        if d <= 0 {
            continue
        }
        for end-st[len(st)-1].r <= d { // 剩余的 d 填充区间后缀
            top := st[len(st)-1]
            st = st[:len(st)-1]
            d += top.r - top.l + 1 // 合并区间
        }
        st = append(st, tuple{end - d + 1, end, st[len(st)-1].s + d})
    }
    return st[len(st)-1].s
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{tasks}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
