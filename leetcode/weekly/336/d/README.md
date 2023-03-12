# 方法一：贪心+暴力

### 提示 1

按照右端点排序。

### 提示 2

对于 $\textit{tasks}[i]$ 来说，它右侧的任务要么和它没有交集，要么包含它的区间**后缀**。

### 提示 3

遍历排序后的任务，先统计区间内的已有的电脑运行时间点，如果个数小于 $\textit{duration}$，则需要新增时间点。根据提示 2，尽量把新增的时间点安排在区间 $[\textit{start},\textit{end}]$ 的后缀上，这样下一个区间就能统计到更多已有的时间点。

附：[视频讲解](https://www.bilibili.com/video/BV1d54y1M7Qg/)

```py [sol1-Python3]
class Solution:
    def findMinimumTime(self, tasks: List[List[int]]) -> int:
        tasks.sort(key=lambda t: t[1])
        run = [False] * (tasks[-1][1] + 1)
        for start, end, d in tasks:
            d -= sum(run[start:end + 1])  # 去掉运行中的时间点
            if d > 0:
                for i in range(end, start - 1, -1):  # 剩余的 d 填充区间后缀
                    if run[i]: continue
                    run[i] = True
                    d -= 1
                    if d == 0: break
        return sum(run)
```

```java [sol1-Java]
class Solution {
    public int findMinimumTime(int[][] tasks) {
        Arrays.sort(tasks, (a, b) -> a[1] - b[1]);
        int ans = 0;
        var run = new boolean[tasks[tasks.length - 1][1] + 1];
        for (var t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            for (int i = start; i <= end; ++i)
                if (run[i]) --d;  // 去掉运行中的时间点
            for (int i = end; d > 0; --i) // 剩余的 d 填充区间后缀
                if (!run[i]) {
                    run[i] = true;
                    --d;
                    ++ans;
                }
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
    bool run[2001];
public:
    int findMinimumTime(vector<vector<int>> &tasks) {
        sort(tasks.begin(), tasks.end(), [](auto &a, auto &b) {
            return a[1] < b[1];
        });
        int ans = 0;
        for (auto &t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            for (int i = start; i <= end; ++i)
                d -= run[i]; // 去掉运行中的时间点
            for (int i = end; d > 0; --i) // 剩余的 d 填充区间后缀
                if (!run[i]) {
                    run[i] = true;
                    --d;
                    ++ans;
                }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func findMinimumTime(tasks [][]int) (ans int) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
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
				run[i] = true
				d--
				ans++
			}
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(nU)$，其中 $n$ 为 $\textit{tasks}$ 的长度，$U=\max(\textit{end}_i)$。
- 空间复杂度：$O(U)$。

# 方法二：贪心+线段树优化

在方法一的暴力更新上优化。

由于有区间更新操作，需要用 lazy 线段树，之前在 [双周赛 98](https://www.bilibili.com/video/BV15D4y1G7ms/) 中讲过了。

对于本题，在更新的时候需要优先递归右子树，从而保证是从右往左更新。其余细节见代码注释。

> 注：由于线段树常数比较大，在数据范围只有几百几千的情况下，不一定比方法一的暴力快。

```py [sol2-Python3]
class Solution:
    def findMinimumTime(self, tasks: List[List[int]]) -> int:
        tasks.sort(key=lambda t: t[1])
        u = tasks[-1][1]
        cnt = [0] * (4 * u)
        todo = [False] * (4 * u)

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

```java [sol2-Java]
class Solution {
    public int findMinimumTime(int[][] tasks) {
        Arrays.sort(tasks, (a, b) -> a[1] - b[1]);
        int u = tasks[tasks.length - 1][1];
        cnt = new int[u * 4];
        todo = new boolean[u * 4];
        for (var t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            suffix = d - query(1, 1, u, start, end);  // 去掉运行中的时间点
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
    // 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log u)
    void update(int o, int l, int r, int L, int R) {
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

```cpp [sol2-C++]
class Solution {
    static constexpr int MX = 2000;

    int cnt[MX * 4];
    bool todo[MX * 4];

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
    void update(int o, int l, int r, int L, int R, int &suffix) {
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
    int findMinimumTime(vector<vector<int>> &tasks) {
        sort(tasks.begin(), tasks.end(), [](auto &a, auto &b) {
            return a[1] < b[1];
        });
        int u = tasks.back()[1];
        for (auto &t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            d -= query(1, 1, u, start, end);  // 去掉运行中的时间点
            if (d > 0) update(1, 1, u, start, end, d); // 新增时间点
        }
        return cnt[1];
    }
};
```

```go [sol2-Go]
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
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	u := tasks[len(tasks)-1][1]
	st := make(seg, u*4)
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

### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{tasks}$ 的长度，$U=\max(\textit{end}_i)$。
- 空间复杂度：$O(U)$。

**注**：如果改用动态开点线段树，可以做到 $O(n\log n)$ 时间和 $O(n\log n)$ 空间。

# 方法三：贪心+栈优化+二分查找

### 前置知识：二分查找

见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 思路

由于每次都是从右到左新增时间点，相当于把若干右侧的区间合并成一个大区间，因此可以用栈来优化。

栈中保存闭区间的左右端点，以及从栈底到栈顶的区间长度之和（类似前缀和）。

合并前先在栈中二分查找包含左端点 $\textit{start}$ 的区间，由于我们保存了区间长度之和，所以可以算出 $[\textit{start},\textit{end}]$ 范围内的运行中的时间点个数。

如果还需要新增时间点，那么就从右到左合并，具体细节见代码。

```py [sol3-Python3]
class Solution:
    def findMinimumTime(self, tasks: List[List[int]]) -> int:
        tasks.sort(key=lambda t: t[1])
        st = [(-2, -2, 0)]  # 闭区间左右端点，栈底到栈顶的区间长度的和
        for start, end, d in tasks:
            _, r, s = st[bisect_left(st, (start,)) - 1]
            d -= st[-1][2] - s  # 去掉运行中的时间点
            if start <= r:  # start 在区间 st[i] 内
                d -= r - start + 1  # 去掉运行中的时间点
            if d <= 0: continue
            while end - st[-1][1] <= d:  # 剩余的 d 填充区间后缀
                l, r, _ = st.pop()
                d += r - l + 1  # 合并区间
            st.append((end - d + 1, end, st[-1][2] + d))
        return st[-1][2]
```

```java [sol3-Java]
class Solution {
    public int findMinimumTime(int[][] tasks) {
        Arrays.sort(tasks, (a, b) -> a[1] - b[1]);
        var st = new ArrayList<int[]>();
        st.add(new int[]{-2, -2, 0}); // 闭区间左右端点，栈底到栈顶的区间长度的和
        for (var t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            var e = st.get(lowerBound(st, start) - 1);
            d -= st.get(st.size() - 1)[2] - e[2]; // 去掉运行中的时间点
            if (start <= e[1]) // start 在区间 st[i] 内
                d -= e[1] - start + 1; // 去掉运行中的时间点
            if (d <= 0) continue;
            while (end - st.get(st.size() - 1)[1] <= d) { // 剩余的 d 填充区间后缀
                e = st.remove(st.size() - 1);
                d += e[1] - e[0] + 1; // 合并区间
            }
            st.add(new int[]{end - d + 1, end, st.get(st.size() - 1)[2] + d});
        }
        return st.get(st.size() - 1)[2];
    }

    // 开区间写法
    private int lowerBound(List<int[]> st, int target) {
        int left = -1, right = st.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // st[left] < target
            // st[right] >= target
            int mid = (left + right) >>> 1;
            if (st.get(mid)[0] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right; // 或者 left+1
    }
}
```

```cpp [sol3-C++]
class Solution {
public:
    int findMinimumTime(vector<vector<int>> &tasks) {
        sort(tasks.begin(), tasks.end(), [](auto &a, auto &b) {
            return a[1] < b[1];
        });
        vector<tuple<int, int, int>> st{{-2, -2, 0}}; // 闭区间左右端点，栈底到栈顶的区间长度的和
        for (auto &t : tasks) {
            int start = t[0], end = t[1], d = t[2];
            auto[_, r, s] = *--lower_bound(st.begin(), st.end(), start, [](const auto &a, int b) {
                return get<0>(a) < b;
            });
            d -= get<2>(st.back()) - s; // 去掉运行中的时间点
            if (start <= r) // start 在区间 st[i] 内
                d -= r - start + 1; // 去掉运行中的时间点
            if (d <= 0) continue;
            while (end - get<1>(st.back()) <= d) { // 剩余的 d 填充区间后缀
                auto[l, r, _] = st.back();
                d += r - l + 1; // 合并区间
                st.pop_back();
            }
            st.emplace_back(end - d + 1, end, get<2>(st.back()) + d);
        }
        return get<2>(st.back());
    }
};
```

```go [sol3-Go]
func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	type tuple struct{ l, r, s int }
	st := []tuple{{-2, -2, 0}} // 闭区间左右端点，栈底到栈顶的区间长度的和
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

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{tasks}$ 的长度。
- 空间复杂度：$O(n)$。
