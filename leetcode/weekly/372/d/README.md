[本题视频讲解](https://www.bilibili.com/video/BV1pC4y1j7Pw/)

## 方法一：离线+最小堆

> 离线：按照自己定义的某种顺序回答询问，而不是按照输入顺序 $\textit{queries}[0],\textit{queries}[1],\cdots$ 回答询问。

下文把 $a_i$ 和 $b_i$ 简称为 $a$ 和 $b$。

不妨设 $a \le b$。

首先遍历 $\textit{queries}$。如果 $a = b$ 或者 $\textit{heights}[a] <\textit{heights}[b]$，那么 Alice 可以直接跳到 Bob 的位置，即 $\textit{ans}[i] = b$。

否则 $\textit{heights}[a] \ge \textit{heights}[b]$，我们可以在位置 $b$ 记录「左边有个 $\textit{heights}[a]$，它属于第 $i$ 个询问」，把数对 $(\textit{heights}[a],i)$ 加到列表 $\textit{qs}[b]$ 中。

然后遍历 $\textit{heights}$，同时用一个最小堆维护上面说的记录：遍历到 $\textit{heights}[i]$ 时，把 $\textit{qs}[i]$ 中的数对全部加入最小堆中。

在加到最小堆之前，我们可以回答堆中所有满足 $\textit{heights}[a] < \textit{heights}[i]$ 的询问，由于 $\textit{heights}[b]\le \textit{heights}[a] < \textit{heights}[i]$，所以该询问的答案是 $i$。

> 为什么要用最小堆？如果堆顶的 $\textit{heights}[a]\ge \textit{heights}[i]$，那么堆中的其余元素也满足 $\textit{heights}[a]\ge \textit{heights}[i]$，这些询问的答案肯定不是 $i$。

### 总结

算法涉及到三个位置，假定 $a \le b$，按照**从左到右**的顺序，它们分别是：

1. $a$：回答询问时，用其高度 $\textit{heights}[a]$ 和当前高度 $\textit{heights}[i]$ 比大小，如果 $\textit{heights}[a] < \textit{heights}[i]$ 则找到答案。
2. $b$：决定了在什么位置把询问加入堆中。注意在遍历到位置 $b$ 之前是不能入堆的。在遍历到位置 $b$ 时入堆，这样后续只需要比较 $\textit{heights}[a] < \textit{heights}[i]$，如果成立，就间接地说明 $\textit{heights}[b] < \textit{heights}[i]$ 也成立。并且，由于我们是从左往右遍历 $\textit{heights}$ 的，当前下标 $i$ 就是 Alice 和 Bob 可以相遇的最左边建筑的下标。
3. 回答询问的位置 $i$。如果堆顶 $\textit{heights}[a]$ 小于当前位置的高度 $\textit{heights}[i]$，则回答堆顶询问，并弹出堆顶。

```py [sol-Python3]
class Solution:
    def leftmostBuildingQueries(self, heights: List[int], queries: List[List[int]]) -> List[int]:
        ans = [-1] * len(queries)
        qs = [[] for _ in heights]
        for i, (a, b) in enumerate(queries):
            if a > b:
                a, b = b, a  # 保证 a <= b
            if a == b or heights[a] < heights[b]:
                ans[i] = b  # a 直接跳到 b
            else:
                qs[b].append((heights[a], i))  # 离线询问

        h = []
        for i, x in enumerate(heights):
            while h and h[0][0] < x:
                # 堆顶的 heights[a] 可以跳到 heights[i]
                ans[heappop(h)[1]] = i
            for q in qs[i]:
                heappush(h, q)  # 后面再回答
        return ans
```

```java [sol-Java]
class Solution {
    public int[] leftmostBuildingQueries(int[] heights, int[][] queries) {
        int[] ans = new int[queries.length];
        Arrays.fill(ans, -1);
        List<int[]>[] qs = new ArrayList[heights.length];
        Arrays.setAll(qs, i -> new ArrayList<>());

        for (int i = 0; i < queries.length; i++) {
            int a = queries[i][0];
            int b = queries[i][1];
            if (a > b) {
                int tmp = a;
                a = b;
                b = tmp; // 保证 a <= b
            }
            if (a == b || heights[a] < heights[b]) {
                ans[i] = b; // a 直接跳到 b
            } else {
                qs[b].add(new int[]{heights[a], i}); // 离线询问
            }
        }

        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        for (int i = 0; i < heights.length; i++) {
            while (!pq.isEmpty() && pq.peek()[0] < heights[i]) {
                // 堆顶的 heights[a] 可以跳到 heights[i]
                ans[pq.poll()[1]] = i;
            }
            for (int[] q : qs[i]) {
                pq.offer(q); // 后面再回答
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> leftmostBuildingQueries(vector<int>& heights, vector<vector<int>>& queries) {
        vector<int> ans(queries.size(), -1);
        vector<vector<pair<int, int>>> qs(heights.size());
        for (int i = 0; i < queries.size(); i++) {
            int a = queries[i][0], b = queries[i][1];
            if (a > b) {
                swap(a, b); // 保证 a <= b
            }
            if (a == b || heights[a] < heights[b]) {
                ans[i] = b; // a 直接跳到 b
            } else {
                qs[b].emplace_back(heights[a], i); // 离线询问
            }
        }

        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        for (int i = 0; i < heights.size(); i++) {
            while (!pq.empty() && pq.top().first < heights[i]) {
                // 堆顶的 heights[a] 可以跳到 heights[i]
                ans[pq.top().second] = i;
                pq.pop();
            }
            for (auto& p : qs[i]) {
                pq.emplace(p); // 后面再回答
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
    ans := make([]int, len(queries))
    for i := range ans {
        ans[i] = -1
    }
    qs := make([][]pair, len(heights))
    for i, q := range queries {
        a, b := q[0], q[1]
        if a > b {
            a, b = b, a // 保证 a <= b
        }
        if a == b || heights[a] < heights[b] {
            ans[i] = b // a 直接跳到 b
        } else {
            qs[b] = append(qs[b], pair{heights[a], i}) // 离线询问
        }
    }

    h := hp{}
    for i, x := range heights {
        for h.Len() > 0 && h[0].h < x {
            // 堆顶的 heights[a] 可以跳到 heights[i]
            ans[heap.Pop(&h).(pair).i] = i
        }
        for _, p := range qs[i] {
            heap.Push(&h, p) // 后面再回答
        }
    }
    return ans
}

type pair struct{ h, i int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log q)$，其中 $n$ 为 $\textit{heights}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n + q)$。

## 方法二：离线+单调栈二分

同方法一，先遍历 $\textit{queries}$，处理出 $\textit{qs}$。

然后**倒序遍历** $\textit{heights}$。试想一下，如果 $\textit{heights}[2]=8,\ \textit{heights}[3]=6$，那么对于在 $\textit{heights}[2]$ 左边的高度来说，$\textit{heights}[3]$ 必然不是第一个相遇的位置，因为我们总是可以选择比 $\textit{heights}[3]$ 更大且更靠左的 $\textit{heights}[2]$。这意味着，当我们遍历到一个更大的高度时，之前遍历过的更小的高度就是无用数据了，要及时清除掉。

这启发我们用一个**底大顶小**的**单调栈**维护高度。原理请看 [单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)。

由于栈中高度严格递减（从栈底到栈顶），可以**二分查找**最后一个大于 $\textit{heights}[a]$ 的高度。原理请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

代码实现时，为方便计算下标，栈中保存的是高度的下标。

```py [sol-Python3]
class Solution:
    def leftmostBuildingQueries(self, heights: List[int], queries: List[List[int]]) -> List[int]:
        ans = [-1] * len(queries)
        qs = [[] for _ in heights]
        for i, (a, b) in enumerate(queries):
            if a > b:
                a, b = b, a  # 保证 a <= b
            if a == b or heights[a] < heights[b]:
                ans[i] = b  # a 直接跳到 b
            else:
                qs[b].append((heights[a], i))  # 离线询问

        st = []
        for i in range(len(heights) - 1, -1, -1):
            for ha, qi in qs[i]:
                # 取反后，相当于找 < -ha 的最大下标，这可以先找 >= -ha 的最小下标，然后减一得到
                j = bisect_left(st, -ha, key=lambda i: -heights[i]) - 1
                if j >= 0:
                    ans[qi] = st[j]
            while st and heights[i] >= heights[st[-1]]:
                st.pop()
            st.append(i)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] leftmostBuildingQueries(int[] heights, int[][] queries) {
        int n = heights.length;
        int[] ans = new int[queries.length];
        List<int[]>[] qs = new ArrayList[n];
        Arrays.setAll(qs, i -> new ArrayList<>());

        for (int i = 0; i < queries.length; i++) {
            int a = queries[i][0];
            int b = queries[i][1];
            if (a > b) {
                int tmp = a;
                a = b;
                b = tmp; // 保证 a <= b
            }
            if (a == b || heights[a] < heights[b]) {
                ans[i] = b; // a 直接跳到 b
            } else {
                qs[b].add(new int[]{heights[a], i}); // 离线询问
            }
        }

        int[] st = new int[n];
        int top = 0;
        for (int i = n - 1; i >= 0; i--) {
            for (int[] q : qs[i]) {
                ans[q[1]] = binarySearch(heights, st, top, q[0]);
            }
            while (top > 0 && heights[i] >= heights[st[top - 1]]) {
                top--;
            }
            st[top++] = i;
        }
        return ans;
    }

    // 返回 st 中最后一个 > x 的高度的下标
    // 如果不存在，返回 -1
    // https://www.bilibili.com/video/BV1AP41137w7/
    private int binarySearch(int[] heights, int[] st, int right, int x) {
        int left = -1; // 开区间 (left, right)
        while (left + 1 < right) { // 开区间不为空
            int mid = (left + right) >>> 1;
            if (heights[st[mid]] > x) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return left < 0 ? -1 : st[left];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> leftmostBuildingQueries(vector<int>& heights, vector<vector<int>>& queries) {
        vector<int> ans(queries.size());
        vector<vector<pair<int, int>>> qs(heights.size());
        for (int i = 0; i < queries.size(); i++) {
            int a = queries[i][0], b = queries[i][1];
            if (a > b) {
                swap(a, b); // 保证 a <= b
            }
            if (a == b || heights[a] < heights[b]) {
                ans[i] = b; // a 直接跳到 b
            } else {
                qs[b].emplace_back(heights[a], i); // 离线询问
            }
        }

        vector<int> st;
        for (int i = heights.size() - 1; i >= 0; i--) {
            for (auto& [ha, qi] : qs[i]) {
                // 取反后，相当于找 < -ha 的最大下标，这可以先找 >= -ha 的最小下标，然后减一得到
                auto it = ranges::lower_bound(st, -ha, {}, [&](int j) { return -heights[j]; });
                ans[qi] = it > st.begin() ? *prev(it) : -1;
            }
            while (!st.empty() && heights[i] >= heights[st.back()]) {
                st.pop_back();
            }
            st.push_back(i);
        }
        return ans;
    }
};
```

```go [sol-Go]
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
    ans := make([]int, len(queries))
    type pair struct{ h, i int }
    qs := make([][]pair, len(heights))
    for i, q := range queries {
        a, b := q[0], q[1]
        if a > b {
            a, b = b, a // 保证 a <= b
        }
        if a == b || heights[a] < heights[b] {
            ans[i] = b // a 直接跳到 b
        } else {
            qs[b] = append(qs[b], pair{heights[a], i}) // 离线询问
        }
    }

    st := []int{}
    for i := len(heights) - 1; i >= 0; i-- {
        for _, q := range qs[i] {
            j := sort.Search(len(st), func(i int) bool { return heights[st[i]] <= q.h }) - 1
            if j >= 0 {
                ans[q.i] = st[j]
            } else {
                ans[q.i] = -1
            }
        }
        for len(st) > 0 && heights[i] >= heights[st[len(st)-1]] {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 为 $\textit{heights}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 方法三：在线+线段树二分

> 在线：按照输入顺序 $\textit{queries}[0],\textit{queries}[1],\cdots$ 一个一个地回答询问。

问题相当于计算区间 $[b+1,n-1]$ 中第一个大于 $v = \textit{heights}[a]$ 的高度的位置。这可以用**线段树二分**解决。

创建一棵维护区间最大值 $\textit{mx}$ 的线段树。

对于每个询问，递归这棵线段树，分类讨论：

- 如果当前区间（线段树的节点对应的区间）最大值 $\textit{mx}\le v$，则当前区间没有大于 $v$ 的数，返回 $-1$。
- 如果当前区间只包含一个元素，则找到答案，返回该元素的下标。
- 如果左子树包含 $b+1$，则递归左子树。
- 如果左子树返回 $-1$，则返回递归右子树的结果。

> 注：方法三是最灵活的，如果题目还有动态修改 $\textit{heights}[i]$ 的操作，方法三也可以做。

```py [sol-Python3]
class Solution:
    def leftmostBuildingQueries(self, heights: List[int], queries: List[List[int]]) -> List[int]:
        n = len(heights)
        mx = [0] * (2 << n.bit_length())

        # 用 heights 初始化线段树，维护区间最大值
        def build(o: int, l: int, r: int) -> None:
            if l == r:
                mx[o] = heights[l]
                return
            m = (l + r) // 2
            build(o * 2, l, m)
            build(o * 2 + 1, m + 1, r)
            mx[o] = max(mx[o * 2], mx[o * 2 + 1])

        # 返回 [L,n-1] 中第一个 > v 的值的下标
        # 如果不存在，返回 -1
        def query(o: int, l: int, r: int, L: int, v: int) -> int:
            if mx[o] <= v:  # 区间最大值 <= v
                return -1  # 没有 > v 的数
            if l == r:  # 找到了
                return l
            m = (l + r) // 2
            if L <= m and (pos := query(o * 2, l, m, L, v)) >= 0:  # 递归左子树
                return pos
            return query(o * 2 + 1, m + 1, r, L, v)  # 递归右子树

        build(1, 0, n - 1)
        ans = []
        for a, b in queries:
            if a > b:
                a, b = b, a  # 保证 a <= b
            if a == b or heights[a] < heights[b]:
                ans.append(b)  # a 直接跳到 b
            else:
                # 线段树二分，找 [b+1,n-1] 中的第一个 > heights[a] 的位置
                ans.append(query(1, 0, n - 1, b + 1, heights[a]))
        return ans
```

```java [sol-Java]
class Solution {
    public int[] leftmostBuildingQueries(int[] heights, int[][] queries) {
        int n = heights.length;
        mx = new int[2 << (Integer.SIZE - Integer.numberOfLeadingZeros(n))];
        build(1, 0, n - 1, heights);

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int a = queries[i][0];
            int b = queries[i][1];
            if (a > b) {
                int tmp = a;
                a = b;
                b = tmp; // 保证 a <= b
            }
            if (a == b || heights[a] < heights[b]) {
                ans[i] = b; // a 直接跳到 b
            } else {
                // 线段树二分，找 [b+1,n-1] 中的第一个 > heights[a] 的位置
                ans[i] = query(1, 0, n - 1, b + 1, heights[a]);
            }
        }
        return ans;
    }

    private int[] mx;

    // 用 heights 初始化线段树，维护区间最大值
    private void build(int o, int l, int r, int[] heights) {
        if (l == r) {
            mx[o] = heights[l];
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m, heights);
        build(o * 2 + 1, m + 1, r, heights);
        mx[o] = Math.max(mx[o * 2], mx[o * 2 + 1]);
    }

    // 返回 [L,n-1] 中第一个 > v 的值的下标
    // 如果不存在，返回 -1
    private int query(int o, int l, int r, int L, int v) {
        if (mx[o] <= v) { // 区间最大值 <= v
            return -1; // 没有 > v 的数
        }
        if (l == r) { // 找到了
            return l;
        }
        int m = (l + r) / 2;
        if (L <= m) {
            int pos = query(o * 2, l, m, L, v); // 递归左子树
            if (pos >= 0) { // 找到了
                return pos;
            }
        }
        return query(o * 2 + 1, m + 1, r, L, v); // 递归右子树
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> mx;

    // 用 heights 初始化线段树，维护区间最大值
    void build(int o, int l, int r, vector<int>& heights) {
        if (l == r) {
            mx[o] = heights[l];
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m, heights);
        build(o * 2 + 1, m + 1, r, heights);
        mx[o] = max(mx[o * 2], mx[o * 2 + 1]);
    }

    // 返回 [L,n-1] 中第一个 > v 的值的下标
    // 如果不存在，返回 -1
    int query(int o, int l, int r, int L, int v) {
        if (mx[o] <= v) { // 区间最大值 <= v
            return -1; // 没有 > v 的数
        }
        if (l == r) { // 找到了
            return l;
        }
        int m = (l + r) / 2;
        if (L <= m) {
            int pos = query(o * 2, l, m, L, v); // 递归左子树
            if (pos >= 0) { // 找到了
                return pos;
            }
        }
        return query(o * 2 + 1, m + 1, r, L, v); // 递归右子树
    }

public:
    vector<int> leftmostBuildingQueries(vector<int>& heights, vector<vector<int>>& queries) {
        int n = heights.size();
        mx.resize(4 << __lg(n));
        build(1, 0, n - 1, heights);

        vector<int> ans;
        for (auto& q : queries) {
            int a = q[0], b = q[1];
            if (a > b) {
                swap(a, b); // 保证 a <= b
            }
            if (a == b || heights[a] < heights[b]) {
                ans.push_back(b); // a 直接跳到 b
            } else {
                // 线段树二分，找 [b+1,n-1] 中的第一个 > heights[a] 的位置
                ans.push_back(query(1, 0, n - 1, b + 1, heights[a]));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type seg []int

// 初始化线段树，维护区间最大值
func (t seg) build(a []int, o, l, r int) {
    if l == r {
        t[o] = a[l]
        return
    }
    m := (l + r) >> 1
    t.build(a, o<<1, l, m)
    t.build(a, o<<1|1, m+1, r)
    t[o] = max(t[o<<1], t[o<<1|1])
}

// 返回 [L,n-1] 中第一个 > v 的值的下标
// 如果不存在，返回 -1
func (t seg) query(o, l, r, L, v int) int {
    if t[o] <= v { // 区间最大值 <= v
        return -1 // 没有 > v 的数
    }
    if l == r { // 找到了
        return l
    }
    m := (l + r) >> 1
    if L <= m {
        pos := t.query(o<<1, l, m, L, v) // 递归左子树
        if pos >= 0 { // 找到了
            return pos
        }
    }
    return t.query(o<<1|1, m+1, r, L, v) // 递归右子树
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
    n := len(heights)
    t := make(seg, 2<<bits.Len(uint(n-1)))
    t.build(heights, 1, 0, n-1)

    ans := make([]int, len(queries))
    for i, q := range queries {
        a, b := q[0], q[1]
        if a > b {
            a, b = b, a // 保证 a <= b
        }
        if a == b || heights[a] < heights[b] {
            ans[i] = b // a 直接跳到 b
        } else {
            // 线段树二分，找 [b+1,n-1] 中的第一个 > heights[a] 的位置
            ans[i] = t.query(1, 0, n-1, b+1, heights[a])
        }
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 为 $\textit{heights}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。在线段树二分中，对于左子树的递归，时间是 $\mathcal{O}(\log n)$ 的（同单点更新）；对于右子树的递归，由于区间满足 $\textit{mx}\le v$ 则不递归，否则只会向下递归，所以这部分的时间也是 $\mathcal{O}(\log n)$ 的，所以线段树二分的时间复杂度为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

更多相似题目，见下面数据结构题单中的「**线段树**」。

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
