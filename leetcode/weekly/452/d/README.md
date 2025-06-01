题目让我们把 $\textit{nums}$ 切一刀，分别计算左右两部分的不同质数个数，再相加，计算相加结果的最大值。

相加结果等价于如下两部分之和：

1. 切一刀之前，整个数组的不同质数之和。
2. 切一刀之后，带来的**额外收益**。考虑一个至少出现两次的质数，其最左边的出现位置 $l$ 和最右边的出现位置 $r$，如果 $l+1\le k\le r$，那么分割后，左右两边都有同一个质数，答案可以额外加一。这等价于创建一个数组 $\textit{cnt}$，把 $\textit{cnt}$ 数组中的区间 $[l+1,r]$ 加一（也可以把区间 $[l,r]$ 加一）。区间加一后，我们需要知道整个 $\textit{cnt}$ 数组的最大值，即为切一刀（相比不切）带来的额外收益。

据此，本题思路为：

1. 用 **Lazy 线段树**维护区间加一、区间最大值。
2. 用**哈希表套有序集合**维护元素及其下标，哈希表的 key 是质数，value 是这个质数在 $\textit{nums}$ 中的下标列表。这个下标列表可以用有序集合维护，方便求最小最大，即区间左右端点。

如何预处理质数，见 [204. 计数质数](https://leetcode.cn/problems/count-primes/)。

常用数据结构代码模板可以看我的 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/)。

[本题视频讲解](https://www.bilibili.com/video/BV1Dz76zfEdi/?t=39m41s)，欢迎点赞关注~

```py [sol-Python3]
MX = 100_001
is_p = [True] * MX
is_p[0] = is_p[1] = False
for i in range(2, MX):
    if is_p[i]:
        for j in range(i * i, MX, i):
            is_p[j] = False


class Node:
    __slots__ = 'val', 'todo'

class LazySegmentTree:
    # 懒标记初始值
    _TODO_INIT = 0

    def __init__(self, n: int, init_val=0):
        # 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
        # init_val 可以是 list 或者 int
        # 如果是 int，那么创建一个 list
        if isinstance(init_val, int):
            init_val = [init_val] * n
        self._n = n
        self._tree = [Node() for _ in range(2 << (n - 1).bit_length())]
        self._build(init_val, 1, 0, n - 1)

    # 合并两个 val
    def _merge_val(self, a: int, b: int) -> int:
        return a if a > b else b

    # 把懒标记作用到 node 子树（本例为区间加）
    def _apply(self, node: int, l: int, r: int, todo: int) -> None:
        cur = self._tree[node]
        # 计算 tree[node] 区间的整体变化
        cur.val += todo
        cur.todo += todo

    # 把当前节点的懒标记下传给左右儿子
    def _spread(self, node: int, l: int, r: int) -> None:
        todo = self._tree[node].todo
        if todo == self._TODO_INIT:  # 没有需要下传的信息
            return
        m = (l + r) // 2
        self._apply(node * 2, l, m, todo)
        self._apply(node * 2 + 1, m + 1, r, todo)
        self._tree[node].todo = self._TODO_INIT  # 下传完毕

    # 合并左右儿子的 val 到当前节点的 val
    def _maintain(self, node: int) -> None:
        self._tree[node].val = self._merge_val(self._tree[node * 2].val, self._tree[node * 2 + 1].val)

    # 用 a 初始化线段树
    # 时间复杂度 O(n)
    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        self._tree[node].todo = self._TODO_INIT
        if l == r:  # 叶子
            self._tree[node].val = a[l]  # 初始化叶节点的值
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)  # 初始化左子树
        self._build(a, node * 2 + 1, m + 1, r)  # 初始化右子树
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, ql: int, qr: int, f: int) -> None:
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            self._apply(node, l, r, f)
            return
        self._spread(node, l, r)
        m = (l + r) // 2
        if ql <= m:  # 更新左子树
            self._update(node * 2, l, m, ql, qr, f)
        if qr > m:  # 更新右子树
            self._update(node * 2 + 1, m + 1, r, ql, qr, f)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            return self._tree[node].val
        self._spread(node, l, r)
        m = (l + r) // 2
        if qr <= m:  # [ql, qr] 在左子树
            return self._query(node * 2, l, m, ql, qr)
        if ql > m:  # [ql, qr] 在右子树
            return self._query(node * 2 + 1, m + 1, r, ql, qr)
        l_res = self._query(node * 2, l, m, ql, qr)
        r_res = self._query(node * 2 + 1, m + 1, r, ql, qr)
        return self._merge_val(l_res, r_res)

    # 用 f 更新 [ql, qr] 中的每个 a[i]
    # 0 <= ql <= qr <= n-1
    # 时间复杂度 O(log n)
    def update(self, ql: int, qr: int, f: int) -> None:
        self._update(1, 0, self._n - 1, ql, qr, f)

    # 返回用 _merge_val 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    # 0 <= ql <= qr <= n-1
    # 时间复杂度 O(log n)
    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)


class Solution:
    def maximumCount(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n = len(nums)
        pos = defaultdict(SortedList)
        for i, x in enumerate(nums):
            if is_p[x]:
                pos[x].add(i)

        t = LazySegmentTree(n, 0)
        for sl in pos.values():
            if len(sl) > 1:
                t.update(sl[0], sl[-1], 1)

        ans = []
        for i, x in queries:
            old = nums[i]
            nums[i] = x

            # 处理旧值
            if is_p[old]:
                sl = pos[old]
                if len(sl) > 1:
                    t.update(sl[0], sl[-1], -1)
                sl.remove(i)
                if len(sl) > 1:
                    t.update(sl[0], sl[-1], 1)
                elif not sl:
                    del pos[old]

            # 处理新值
            if is_p[x]:
                sl = pos[x]
                if len(sl) > 1:
                    t.update(sl[0], sl[-1], -1)
                sl.add(i)
                if len(sl) > 1:
                    t.update(sl[0], sl[-1], 1)

            ans.append(len(pos) + t.query(0, n - 1))

        return ans
```

```java [sol-Java]
class LazySegmentTree {
    // 懒标记初始值
    private static final int TODO_INIT = 0; // **根据题目修改**

    private static final class Node {
        int val;
        int todo;
    }

    // 合并两个 val
    private int mergeVal(int a, int b) {
        return Math.max(a, b);
    }

    // 合并两个懒标记
    private int mergeTodo(int a, int b) {
        return a + b;
    }

    // 把懒标记作用到 node 子树（本例为区间加）
    private void apply(int node, int l, int r, int todo) {
        Node cur = tree[node];
        // 计算 tree[node] 区间的整体变化
        cur.val += todo;
        cur.todo = mergeTodo(todo, cur.todo);
    }

    private final int n;
    private final Node[] tree;

    // 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 initVal
    public LazySegmentTree(int n, int initVal) {
        this.n = n;
        int[] a = new int[n];
        Arrays.fill(a, initVal);
        tree = new Node[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(a, 1, 0, n - 1);
    }

    // 线段树维护数组 a
    public LazySegmentTree(int[] a) {
        n = a.length;
        tree = new Node[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(a, 1, 0, n - 1);
    }

    // 用 f 更新 [ql, qr] 中的每个 a[i]
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    public void update(int ql, int qr, int f) {
        update(1, 0, n - 1, ql, qr, f);
    }

    // 返回用 mergeVal 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    public int query(int ql, int qr) {
        return query(1, 0, n - 1, ql, qr);
    }

    // 把当前节点的懒标记下传给左右儿子
    private void spread(int node, int l, int r) {
        int todo = tree[node].todo;
        if (todo == TODO_INIT) { // 没有需要下传的信息
            return;
        }
        int m = (l + r) / 2;
        apply(node * 2, l, m, todo);
        apply(node * 2 + 1, m + 1, r, todo);
        tree[node].todo = TODO_INIT; // 下传完毕
    }

    // 合并左右儿子的 val 到当前节点的 val
    private void maintain(int node) {
        tree[node].val = mergeVal(tree[node * 2].val, tree[node * 2 + 1].val);
    }

    // 用 a 初始化线段树
    // 时间复杂度 O(n)
    private void build(int[] a, int node, int l, int r) {
        tree[node] = new Node();
        tree[node].todo = TODO_INIT;
        if (l == r) { // 叶子
            tree[node].val = a[l]; // 初始化叶节点的值
            return;
        }
        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    private void update(int node, int l, int r, int ql, int qr, int f) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            apply(node, l, r, f);
            return;
        }
        spread(node, l, r);
        int m = (l + r) / 2;
        if (ql <= m) { // 更新左子树
            update(node * 2, l, m, ql, qr, f);
        }
        if (qr > m) { // 更新右子树
            update(node * 2 + 1, m + 1, r, ql, qr, f);
        }
        maintain(node);
    }

    private int query(int node, int l, int r, int ql, int qr) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node].val;
        }
        spread(node, l, r);
        int m = (l + r) / 2;
        if (qr <= m) { // [ql, qr] 在左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 在右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        int lRes = query(node * 2, l, m, ql, qr);
        int rRes = query(node * 2 + 1, m + 1, r, ql, qr);
        return mergeVal(lRes, rRes);
    }
}

class Solution {
    private static final int MX = 100_000;
    private static final boolean[] np = new boolean[MX + 1];
    private static boolean initialized = false;

    // 这样写比 static block 快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        np[0] = np[1] = true;
        for (int i = 2; i <= MX; i++) {
            if (!np[i]) {
                for (int j = i; j <= MX / i; j++) { // 避免溢出的写法
                    np[i * j] = true;
                }
            }
        }
    }

    public int[] maximumCount(int[] nums, int[][] queries) {
        init();

        int n = nums.length;
        Map<Integer, TreeSet<Integer>> pos = new HashMap<>();

        LazySegmentTree t = new LazySegmentTree(n, 0);
        for (int i = 0; i < n; i++) {
            int v = nums[i];
            if (!np[v]) {
                pos.computeIfAbsent(v, k -> new TreeSet<>()).add(i);
            }
        }

        // 对出现次数大于 1 的质数，其最小最大索引之间的区间 +1
        for (TreeSet<Integer> s : pos.values()) {
            if (s.size() > 1) {
                t.update(s.first(), s.last(), 1);
            }
        }

        int[] ans = new int[queries.length];
        for (int qi = 0; qi < queries.length; qi++) {
            int[] q = queries[qi];
            int i = q[0];
            int v = q[1];
            int old = nums[i];
            nums[i] = v;

            // 删除旧值 old 的影响
            if (!np[old]) {
                TreeSet<Integer> s = pos.get(old);
                if (s.size() > 1) {
                    t.update(s.first(), s.last(), -1);
                }
                s.remove(i);
                if (s.size() > 1) {
                    t.update(s.first(), s.last(), 1);
                } else if (s.isEmpty()) {
                    pos.remove(old);
                }
            }

            // 插入新值 v 的影响
            if (!np[v]) {
                TreeSet<Integer> s = pos.computeIfAbsent(v, k -> new TreeSet<>());
                if (s.size() > 1) {
                    t.update(s.first(), s.last(), -1);
                }
                s.add(i);
                if (s.size() > 1) {
                    t.update(s.first(), s.last(), 1);
                }
            }

            ans[qi] = pos.size() + t.query(0, n - 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
const int MX = 100'000;
bool np[MX + 1];

int init = []() {
    np[0] = np[1] = true;
    for (int i = 2; i <= MX; i++) {
        if (!np[i]) {
            for (int j = i; j <= MX / i; j++) { // 避免溢出的写法
                np[i * j] = true;
            }
        }
    }
    return 0;
}();

template<typename T, typename F>
class LazySegmentTree {
    // 注：也可以去掉 template<typename T, typename F>，改在这里定义 T 和 F
    // using T = pair<int, int>;
    // using F = pair<int, int>;

    // 懒标记初始值
    const F TODO_INIT = 0;

    struct Node {
        T val;
        F todo;
    };

    int n;
    vector<Node> tree;

    // 合并两个 val
    T merge_val(T a, T b) const {
        return max(a, b);
    }

    // 合并两个懒标记
    F merge_todo(F a, F b) const {
        return a + b;
    }

    // 把懒标记作用到 node 子树（本例为区间加）
    void apply(int node, int l, int r, F todo) {
        Node& cur = tree[node];
        // 计算 tree[node] 区间的整体变化
        cur.val += todo;
        cur.todo = merge_todo(todo, cur.todo);
    }

    // 把当前节点的懒标记下传给左右儿子
    void spread(int node, int l, int r) {
        Node& cur = tree[node];
        F todo = cur.todo;
        if (todo == TODO_INIT) { // 没有需要下传的信息
            return;
        }
        int m = (l + r) / 2;
        apply(node * 2, l, m, todo);
        apply(node * 2 + 1, m + 1, r, todo);
        cur.todo = TODO_INIT; // 下传完毕
    }

    // 合并左右儿子的 val 到当前节点的 val
    void maintain(int node) {
        tree[node].val = merge_val(tree[node * 2].val, tree[node * 2 + 1].val);
    }

    // 用 a 初始化线段树
    // 时间复杂度 O(n)
    void build(const vector<T>& a, int node, int l, int r) {
        Node& cur = tree[node];
        cur.todo = TODO_INIT;
        if (l == r) { // 叶子
            cur.val = a[l]; // 初始化叶节点的值
            return;
        }
        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    void update(int node, int l, int r, int ql, int qr, F f) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            apply(node, l, r, f);
            return;
        }
        spread(node, l, r);
        int m = (l + r) / 2;
        if (ql <= m) { // 更新左子树
            update(node * 2, l, m, ql, qr, f);
        }
        if (qr > m) { // 更新右子树
            update(node * 2 + 1, m + 1, r, ql, qr, f);
        }
        maintain(node);
    }

    T query(int node, int l, int r, int ql, int qr) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node].val;
        }
        spread(node, l, r);
        int m = (l + r) / 2;
        if (qr <= m) { // [ql, qr] 在左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 在右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        T l_res = query(node * 2, l, m, ql, qr);
        T r_res = query(node * 2 + 1, m + 1, r, ql, qr);
        return merge_val(l_res, r_res);
    }

public:
    // 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
    LazySegmentTree(int n, T init_val = 0) : LazySegmentTree(vector<T>(n, init_val)) {}

    // 线段树维护数组 a
    LazySegmentTree(const vector<T>& a) : n(a.size()), tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, n - 1);
    }

    // 用 f 更新 [ql, qr] 中的每个 a[i]
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    void update(int ql, int qr, F f) {
        update(1, 0, n - 1, ql, qr, f);
    }

    // 返回用 merge_val 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    T query(int ql, int qr) {
        return query(1, 0, n - 1, ql, qr);
    }
};

class Solution {
public:
    vector<int> maximumCount(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();

        unordered_map<int, set<int>> pos;

        // 初始化 pos，记录每个质数的位置（有序）
        for (int i = 0; i < n; i++) {
            int v = nums[i];
            if (!np[v]) {
                pos[v].insert(i);
            }
        }

        LazySegmentTree<int, int> t(n);
        for (auto& [_, s] : pos) {
            if (s.size() > 1) {
                t.update(*s.begin(), *s.rbegin(), 1);
            }
        }

        vector<int> ans;
        for (auto& q : queries) {
            int i = q[0], v = q[1];
            int old = nums[i];
            nums[i] = v;

            // 删除旧值 old 的影响
            if (!np[old]) {
                auto& s = pos[old];
                if (s.size() > 1) {
                    t.update(*s.begin(), *s.rbegin(), -1);
                }
                s.erase(i);
                if (s.size() > 1) {
                    t.update(*s.begin(), *s.rbegin(), 1);
                } else if (s.empty()) {
                    pos.erase(old);
                }
            }

            // 插入新值 v 的影响
            if (!np[v]) {
                auto& s = pos[v];
                if (s.size() > 1) {
                    t.update(*s.begin(), *s.rbegin(), -1);
                }
                s.insert(i);
                if (s.size() > 1) {
                    t.update(*s.begin(), *s.rbegin(), 1);
                }
            }

            ans.push_back(pos.size() + t.query(0, n - 1));
        }

        return ans;
    }
};
```

```go [sol-Go]
// import "github.com/emirpasic/gods/v2/trees/redblacktree"

const mx int = 1e5

var np = [mx + 1]bool{true, true}

func init() {
	for i := 2; i <= mx; i++ {
		if !np[i] {
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}
}

type lazySeg []struct {
	l, r int
	mx   int
	todo int
}

func mergeInfo(a, b int) int {
	return max(a, b)
}

const todoInit = 0

func mergeTodo(f, old int) int {
	return f + old
}

func (t lazySeg) apply(o int, f int) {
	cur := &t[o]
	cur.mx += f
	cur.todo = mergeTodo(f, cur.todo)
}

func (t lazySeg) maintain(o int) {
	t[o].mx = mergeInfo(t[o<<1].mx, t[o<<1|1].mx)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t lazySeg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].mx = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazySeg) update(o, l, r int, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t lazySeg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mx
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func newLazySegmentTreeWithArray(a []int) lazySeg {
	n := len(a)
	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func newLazySegmentTree(n int, initVal int) lazySeg {
	a := make([]int, n)
	for i := range a {
		a[i] = initVal
	}
	return newLazySegmentTreeWithArray(a)
}

func maximumCount(nums []int, queries [][]int) (ans []int) {
	n := len(nums)
	pos := map[int]*redblacktree.Tree[int, struct{}]{}
	for i, v := range nums {
		if np[v] {
			continue
		}
		if _, ok := pos[v]; !ok {
			pos[v] = redblacktree.New[int, struct{}]()
		}
		pos[v].Put(i, struct{}{})
	}

	t := newLazySegmentTree(n, 0)
	for _, ps := range pos {
		if ps.Size() > 1 {
			t.update(1, ps.Left().Key, ps.Right().Key, 1)
		}
	}

	for _, q := range queries {
		i, v := q[0], q[1]
		old := nums[i]
		nums[i] = v

		// 删除旧值 old 的影响
		if !np[old] {
			ps := pos[old]
			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, -1)
			}
			ps.Remove(i)

			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, 1)
			} else if ps.Empty() {
				delete(pos, old)
			}
		}

		// 插入新值 v 的影响
		if !np[v] {
			if _, ok := pos[v]; !ok {
				pos[v] = redblacktree.New[int, struct{}]()
			}
			ps := pos[v]
			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, -1)
			}
			ps.Put(i, struct{}{})
			if ps.Size() > 1 {
				t.update(1, ps.Left().Key, ps.Right().Key, 1)
			}
		}

		ans = append(ans, len(pos)+t.query(1, 0, n-1))
	}

	return
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

更多相似题目，见下面数据结构题单的「**§8.4 Lazy 线段树**」。

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
