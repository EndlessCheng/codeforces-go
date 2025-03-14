## 提示 1

核心想法是用 BFS 模拟这个过程。

当我们位于下标 $i$ 时，下一步可以去哪？换句话说，通过**一次**操作，可以到达哪些下标？

## 提示 2

如果翻转子数组 $[L,R]$，那么：

- 下标 $L$ 翻转到下标 $R$。
- 下标 $L+1$ 翻转到下标 $R-1$。
- 下标 $L+2$ 翻转到下标 $R-2$。
- 依此类推，可以发现翻转前后的下标之和恒等于 $L+R$，所以下标 $i$ 在翻转后的下标是 $L+R-i$。

推论：

- 当子数组向右滑动时，每滑动一个单位，$L$ 和 $R$ 都增加 $1$，所以翻转后的下标会增加 $2$。
- 当子数组向左滑动时，每滑动一个单位，$L$ 和 $R$ 都减少 $1$，所以翻转后的下标会减少 $2$。
- 因此，**$i$ 翻转后的下标组成了一个公差为 $2$ 的等差数列**。

$i$ 翻转后的下标，最小是多少，最大是多少？求出最小值和最大值，就能求出 $i$ 翻转后的所有下标。

## 提示 3

如果不考虑下标越界，那么：

- 当 $i$ 在子数组右端点时，可以翻转到子数组左端点 $i-k+1$，这是最小值。
- 当 $i$ 在子数组左端点时，可以翻转到子数组右端点 $i+k-1$，这是最大值。

但如果（举例）$k=4$，长为 $k$ 的子数组，右端点最小是 $k-1=3$，当 $i=0,1,2$ 的时候，$i$ 不可能在子数组右端点。同样地，左端点最大是 $n-k$，当 $i > n-k$ 的时候，$i$ 不可能在子数组左端点。

这意味着 $i$ 比较小（比较大）的情况，应当特殊处理：

- $i < k-1$ 的情况，当子数组在最左边时，$L=0,\ R=k-1$，$i$ 翻转后是 $L+R-i=0+(k-1)-i=k-i-1$，小于 $k-i-1$ 的下标无法翻转得到。
- $i > n-k$ 的情况，当子数组在最右边时，$L=n-k,\ R=n-1$，$i$ 翻转后是 $L+R-i=(n-k)+(n-1) - i=2n-k-i-1$，大于 $2n-k-i-1$ 的下标无法翻转到。

综上所述：

- $i$ 翻转后的**最小值**为 $\max(i-k+1,k-i-1)$。
- $i$ 翻转后的**最大值**为 $\min(i+k-1,2n-k-i-1)$。

## 方法一：BFS + 有序集合

由于等差数列的公差为 $2$，翻转后的下标要么都是偶数，要么都是奇数。我们用两个有序集合 $\textit{indices}_0$ 和 $\textit{indices}_1$ 分别维护**没有访问过的**偶数下标和奇数下标。注意这些下标不能在 $\textit{banned}$ 中。由于 $p$ 是起点（已访问），所以也不在有序集合中。

然后用 BFS 模拟。

在有序集合上，一边遍历翻转后的下标 $j$，一边把 $j$ 从有序集合中删除。这样可以避免重复访问已经访问过的下标，也方便我们查找下一个没有访问过的下标。

关于 BFS 的基本原理，见[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)。

```py [sol-Python3]
class Solution:
    def minReverseOperations(self, n: int, p: int, banned: List[int], k: int) -> List[int]:
        ban = set(banned) | {p}
        indices = [SortedList(), SortedList()]  # import sortedcontainers
        for i in range(n):
            if i not in ban:
                indices[i % 2].add(i)
        indices[0].add(n)  # 哨兵，下面无需判断越界
        indices[1].add(n)

        ans = [-1] * n
        ans[p] = 0  # 起点
        q = deque([p])
        while q:
            i = q.popleft()
            # indices[mn % 2] 中的从 mn 到 mx 的所有下标都可以从 i 翻转到
            mn = max(i - k + 1, k - i - 1)
            mx = min(i + k - 1, n * 2 - k - i - 1)
            sl = indices[mn % 2]
            idx = sl.bisect_left(mn)
            while sl[idx] <= mx:
                j = sl[idx]
                ans[j] = ans[i] + 1  # 移动一步
                q.append(j)
                sl.pop(idx)
                # 注意 pop(idx) 会使后续元素向左移，不需要写 idx += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int[] minReverseOperations(int n, int p, int[] banned, int k) {
        Set<Integer> ban = new HashSet<>();
        for (int b : banned) {
            ban.add(b);
        }

        TreeSet<Integer>[] indices = new TreeSet[]{new TreeSet<>(), new TreeSet<>()};
        for (int i = 0; i < n; i++) {
            if (i != p && !ban.contains(i)) {
                indices[i % 2].add(i);
            }
        }

        int[] ans = new int[n];
        Arrays.fill(ans, -1);
        ans[p] = 0; // 起点
        Queue<Integer> q = new ArrayDeque<>();
        q.offer(p);
        while (!q.isEmpty()) {
            int i = q.poll();
            // indices[mn % 2] 中的从 mn 到 mx 的所有下标都可以从 i 翻转到
            int mn = Math.max(i - k + 1, k - i - 1);
            int mx = Math.min(i + k - 1, n * 2 - k - i - 1);
            TreeSet<Integer> set = indices[mn % 2];
            for (Iterator<Integer> it = set.tailSet(mn).iterator(); it.hasNext(); it.remove()) {
                int j = it.next();
                if (j > mx) {
                    break;
                }
                ans[j] = ans[i] + 1; // 移动一步
                q.offer(j);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minReverseOperations(int n, int p, vector<int>& banned, int k) {
        unordered_set<int> ban{banned.begin(), banned.end()};
        set<int> indices[2];
        for (int i = 0; i < n; i++) {
            if (i != p && !ban.contains(i)) {
                indices[i % 2].insert(i);
            }
        }
        indices[0].insert(n); // 哨兵，下面无需判断 it != st.end()
        indices[1].insert(n);

        vector<int> ans(n, -1);
        ans[p] = 0; // 起点
        queue<int> q;
        q.push(p);
        while (!q.empty()) {
            int i = q.front(); q.pop();
            // indices[mn % 2] 中的从 mn 到 mx 的所有下标都可以从 i 翻转到
            int mn = max(i - k + 1, k - i - 1);
            int mx = min(i + k - 1, n * 2 - k - i - 1);
            auto& st = indices[mn % 2];
            for (auto it = st.lower_bound(mn); *it <= mx; it = st.erase(it)) {
                ans[*it] = ans[i] + 1; // 移动一步
                q.push(*it);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minReverseOperations(n int, p int, banned []int, k int) []int {
    ban := map[int]struct{}{p: {}}
    for _, b := range banned {
        ban[b] = struct{}{}
    }

    indices := [2]*redblacktree.Tree[int, struct{}]{
        redblacktree.New[int, struct{}](),
        redblacktree.New[int, struct{}](),
    }
    for i := range n {
        if _, ok := ban[i]; !ok {
            indices[i%2].Put(i, struct{}{})
        }
    }
    indices[0].Put(n, struct{}{}) // 哨兵，下面无需判断 node != nil
    indices[1].Put(n, struct{}{})

    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    ans[p] = 0 // 起点
    q := []int{p}
    for len(q) > 0 {
        i := q[0]
        q = q[1:]
        // indices[mn%2] 中的从 mn 到 mx 的所有下标都可以从 i 翻转到
        mn := max(i-k+1, k-i-1)
        mx := min(i+k-1, n*2-k-i-1)
        t := indices[mn%2]
        for node, _ := t.Ceiling(mn); node.Key <= mx; node, _ = t.Ceiling(mn) {
            j := node.Key
            ans[j] = ans[i] + 1 // 移动一步
            q = append(q, j)
            t.Remove(j)
        }
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。每个下标入队出队各至多一次，每次（均摊）$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：BFS + 并查集

**前置知识**：并查集。

返璞归真，把方法一的有序集合，合并一个列表 $\textit{indices}=[0,1,2,\ldots,n-1]$。

改成列表后，直接删除列表元素的话，时间复杂度是 $\mathcal{O}(n)$ 的。如何高效地删除元素？

举例说明。假设我们有下标 $0,1,2,3,4$，要把其中的 $2$ 删除。用并查集思考，改成把 $j=2$ 与 $j+2=4$ 合并，也就是调用 $\texttt{merge}(j,j+2)$。这样下次遍历的时候，就可以利用并查集的 $\texttt{find}$ 函数，**跳过**已删除的下标。比如想找 $\ge 2$ 的没有访问过（没被删除）的最小偶数下标，只需要调用并查集的 $\texttt{find}(2)$，结果是 $4$。

代码实现时，$\texttt{merge}(j,j+2)$ 可以优化成 $\texttt{merge}(j,\textit{mx}+2)$。既然范围内的数都要删除，直接一步到位，全部指向下一个没被删除的数，即 $\textit{mx}+2$。为了保证 $\textit{mx}+2$ 一定存在，可以添加哨兵 $n$ 和 $n+1$。也就是说，并查集的大小是 $n+2$ 而不是 $n$。

```py [sol-Python3]
class UnionFind:
    def __init__(self, n: int):
        self.fa = list(range(n))

    def find(self, x: int) -> int:
        if self.fa[x] != x:
            self.fa[x] = self.find(self.fa[x])
        return self.fa[x]

    def merge(self, from_: int, to: int) -> None:
        self.fa[self.find(from_)] = self.find(to)

class Solution:
    def minReverseOperations(self, n: int, p: int, banned: List[int], k: int) -> List[int]:
        indices = UnionFind(n + 2)
        indices.merge(p, p + 2)  # 删除 p
        for i in banned:
            indices.merge(i, i + 2)  # 删除 i

        ans = [-1] * n
        ans[p] = 0
        q = deque([p])
        while q:
            i = q.popleft()
            mn = max(i - k + 1, k - i - 1)
            mx = min(i + k - 1, n * 2 - k - i - 1)
            j = indices.find(mn)
            while j <= mx:
                ans[j] = ans[i] + 1
                q.append(j)
                indices.merge(j, mx + 2)  # 删除 j
                j = indices.find(j + 2)  # 快速跳到 >= j+2 的下一个下标
        return ans
```

```py [sol-Python3 优化]
class UnionFind:
    def __init__(self, n: int):
        self.fa = list(range(n))

    def find(self, x: int) -> int:
        if self.fa[x] != x:
            self.fa[x] = self.find(self.fa[x])
        return self.fa[x]

class Solution:
    def minReverseOperations(self, n: int, p: int, banned: List[int], k: int) -> List[int]:
        indices = UnionFind(n + 2)
        indices.fa[p] += 2  # 删除 p
        for i in banned:
            indices.fa[i] += 2  # 删除 i

        ans = [-1] * n
        ans[p] = 0
        q = deque([p])
        while q:
            i = q.popleft()
            mn = max(i - k + 1, k - i - 1)
            mx = min(i + k - 1, n * 2 - k - i - 1)
            end = indices.find(mx + 2)
            j = indices.find(mn)
            while j <= mx:
                ans[j] = ans[i] + 1
                q.append(j)
                indices.fa[j] = end  # 删除 j
                j = indices.find(j + 2)  # 快速跳到 >= j+2 的下一个下标
        return ans
```

```java [sol-Java]
class UnionFind {
    private final int[] fa;

    public UnionFind(int n) {
        fa = new int[n];
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
        fa[find(from)] = find(to);
    }
}

class Solution {
    public int[] minReverseOperations(int n, int p, int[] banned, int k) {
        UnionFind indices = new UnionFind(n + 2);
        indices.merge(p, p + 2); // 删除 p
        for (int i : banned) {
            indices.merge(i, i + 2); // 删除 i
        }

        int[] ans = new int[n];
        Arrays.fill(ans, -1);
        ans[p] = 0;
        Queue<Integer> q = new ArrayDeque<>(); // 注：如果改用数组模拟队列，可以再快一些
        q.offer(p);
        while (!q.isEmpty()) {
            int i = q.poll();
            int mn = Math.max(i - k + 1, k - i - 1);
            int mx = Math.min(i + k - 1, n * 2 - k - i - 1);
            for (int j = indices.find(mn); j <= mx; j = indices.find(j + 2)) { // 快速跳到 >= j+2 的下一个下标
                ans[j] = ans[i] + 1;
                q.offer(j);
                indices.merge(j, mx + 2); // 删除 j
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class UnionFind {
    vector<int> fa;

public:
    UnionFind(int n) : fa(n) {
        iota(fa.begin(), fa.end(), 0);
    }

    int find(int x) {
        if (fa[x] != x) {
            fa[x] = find(fa[x]);
        }
        return fa[x];
    }

    void merge(int from, int to) {
        fa[find(from)] = find(to);
    }
};

class Solution {
public:
    vector<int> minReverseOperations(int n, int p, vector<int>& banned, int k) {
        UnionFind indices(n + 2);
        indices.merge(p, p + 2); // 删除 p
        for (int i : banned) {
            indices.merge(i, i + 2); // 删除 i
        }

        vector<int> ans(n, -1);
        ans[p] = 0;
        queue<int> q;
        q.push(p);
        while (!q.empty()) {
            int i = q.front(); q.pop();
            int mn = max(i - k + 1, k - i - 1);
            int mx = min(i + k - 1, n * 2 - k - i - 1);
            for (int j = indices.find(mn); j <= mx; j = indices.find(j + 2)) { // 快速跳到 >= j+2 的下一个下标
                ans[j] = ans[i] + 1;
                q.push(j);
                indices.merge(j, mx + 2); // 删除 j
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type unionFind struct {
    fa []int
}

func newUnionFind(n int) unionFind {
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }
    return unionFind{fa}
}

func (uf unionFind) find(x int) int {
    if uf.fa[x] != x {
        uf.fa[x] = uf.find(uf.fa[x])
    }
    return uf.fa[x]
}

func (uf unionFind) merge(from, to int) {
    uf.fa[uf.find(from)] = uf.find(to)
}

func minReverseOperations(n, p int, banned []int, k int) []int {
    indices := newUnionFind(n + 2)
    indices.merge(p, p+2) // 删除 p
    for _, i := range banned {
        indices.merge(i, i+2) // 删除 i
    }

    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    ans[p] = 0
    q := []int{p}
    for len(q) > 0 {
        i := q[0]
        q = q[1:]
        mn := max(i-k+1, k-i-1)
        mx := min(i+k-1, n*2-k-i-1)
        for j := indices.find(mn); j <= mx; j = indices.find(j + 2) { // 快速跳到 >= j+2 的下一个下标
            ans[j] = ans[i] + 1
            q = append(q, j)
            indices.merge(j, mx+2) // 删除 j
        }
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$。
- 空间复杂度：$O(n)$。

能否做到 $\mathcal{O}(n)$？见 [RMQ 标准算法和线性树上并查集](https://ljt12138.blog.uoj.ac/blog/4874)。

更多相似题目，见数据结构题单中的「**§7.4 数组上的并查集**」。

并查集的模板代码也整理在数据结构题单中。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
