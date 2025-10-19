## 前置题目/知识

1. 本题的简单版本 [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)，[我的题解](https://leetcode.cn/problems/contiguous-array/solutions/3805089/shi-zi-bian-xing-mei-ju-you-wei-hu-zuo-p-x9q2/)。
2. [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。
3. [分块思想](https://oi-wiki.org/ds/decompose/) 或者 [Lazy 线段树](https://oi-wiki.org/ds/seg/#%E7%BA%BF%E6%AE%B5%E6%A0%91%E7%9A%84%E5%8C%BA%E9%97%B4%E4%BF%AE%E6%94%B9%E4%B8%8E%E6%87%92%E6%83%B0%E6%A0%87%E8%AE%B0)。

## 转化

从左到右遍历 $\textit{nums}$，如果固定子数组右端点为 $i$，要想让子数组包含某个元素 $x$，左端点必须 $\le x\ 最后一次出现的位置$。只要子数组包含最近遇到的 $x$，那么无论子数组有多长，都包含了 $x$。所以对于多个相同元素 $x$，我们只关注最近遇到的 $x$。

由于我们只关注最近遇到的 $x$，所以在遍历 $\textit{nums}$ 的过程中，$\textit{nums}$ 是**动态变化**的。

以 $\textit{nums}=[1,2,1,2,3,3]$ 为例：

- 遍历到 $i=0$，把 $\textit{nums}$ 视作 $[1,*,*,*,*,*]$。
- 遍历到 $i=1$，把 $\textit{nums}$ 视作 $[1,2,*,*,*,*]$。
- 遍历到 $i=2$，把 $\textit{nums}$ 视作 $[*,2,1,*,*,*]$。
- 遍历到 $i=3$，把 $\textit{nums}$ 视作 $[*,*,1,2,*,*]$。
- 遍历到 $i=4$，把 $\textit{nums}$ 视作 $[*,*,1,2,3,*]$。
- 遍历到 $i=5$，把 $\textit{nums}$ 视作 $[*,*,1,2,*,3]$。

根据 525 题我的题解，把偶数视作 $-1$，奇数视作 $1$，**遍历过的星号视作** $0$，设这个新数组为 $a$，问题相当于：

- 计算 $a$ 中和为 $0$ 的最长子数组的长度。

设 $a$ 的长为 $n+1$ 的前缀和数组为 $\textit{sum}$。根据 525 题我的题解，问题相当于：

- 枚举 $i$，在 $[0,i-1]$ 中找到一个下标最小的 $\textit{sum}[j]$，满足 $\textit{sum}[j] = \textit{sum}[i]$。
- 用子数组长度 $i-j$ 更新答案的最大值。

根据上面动态变化的过程：

- 设 $x=\textit{nums}[i]$ 对应的 $a[i]$ 值为 $v$。
- 当我们首次遇到 $x$ 时，对于前缀和 $\textit{sum}$ 来说，$[i,n]$ 要全部增加 $v$。
- 当我们再次遇到 $x$ 时，原来的 $\textit{nums}[j]$ 变成星号（$a[j]=0$），$x$ 搬到了新的位置 $i$，所以对于前缀和 $\textit{sum}$ 来说，我们要撤销 $[j,i-1]$ 的加 $v$，也就是把 $[j,i-1]$ 减 $v$。

整理一下，我们需要维护一个动态变化的前缀和数组，需要一个数据结构，支持：

1. 把 $\textit{sum}$ 的某个子数组增加 $1$ 或者 $-1$。
2. 查询 $\textit{sum}[i]$ 在 $\textit{sum}$ 中首次出现的位置。

## 方法一：分块

见前置知识。

每块维护块内 $\textit{sum}[i]$ 首次出现的位置，以及区间加的 Lazy 标记。

```go
func longestBalanced(nums []int) (ans int) {
	n := len(nums)
	B := int(math.Sqrt(float64(n+1)))/2 + 1
	sum := make([]int, n+1)

	// === 分块模板开始 ===
	// 用分块维护 sum
	type block struct {
		l, r int // [l,r) 左闭右开
		todo int
		pos  map[int]int
	}
	blocks := make([]block, n/B+1)
	calcPos := func(l, r int) map[int]int {
		pos := map[int]int{}
		for j := r - 1; j >= l; j-- {
			pos[sum[j]] = j
		}
		return pos
	}
	for i := 0; i <= n; i += B {
		r := min(i+B, n+1)
		pos := calcPos(i, r)
		blocks[i/B] = block{i, r, 0, pos}
	}

	// sum[l:r] 增加 v
	rangeAdd := func(l, r, v int) {
		for i := range blocks {
			b := &blocks[i]
			if b.r <= l {
				continue
			}
			if b.l >= r {
				break
			}
			if l <= b.l && b.r <= r { // 完整块
				b.todo += v
			} else { // 部分块，直接重算
				for j := b.l; j < b.r; j++ {
					sum[j] += b.todo
					if l <= j && j < r {
						sum[j] += v
					}
				}
				b.pos = calcPos(b.l, b.r)
				b.todo = 0
			}
		}
	}

	// 返回 sum[:r] 中第一个 v 的下标
	// 如果没有 v，返回 n
	findFirst := func(r, v int) int {
		for i := range blocks {
			b := &blocks[i]
			if b.r <= r { // 完整块，直接查哈希表
				if j, ok := b.pos[v-b.todo]; ok {
					return j
				}
			} else { // 部分块，暴力查找
				for j := b.l; j < r; j++ {
					if sum[j] == v-b.todo {
						return j
					}
				}
				break
			}
		}
		return n
	}
	// === 分块模板结束 ===

	last := map[int]int{} // nums 的元素上一次出现的位置
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		v := x%2*2 - 1
		if j := last[x]; j == 0 { // 首次遇到 x
			rangeAdd(i, n+1, v) // sum[i:] 增加 v
		} else { // 再次遇到 x
			rangeAdd(j, i, -v) // 撤销之前对 sum[j:i] 的增加
		}
		last[x] = i

		s := sum[i] + blocks[i/B].todo // sum[i] 的实际值
		ans = max(ans, i-findFirst(i-ans, s)) // 优化右边界
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：Lazy 线段树

由于 $a$ 中元素只有 $-1,0,1$，所以 $\textit{sum}$ 数组相邻元素之差 $\le 1$。这意味着，设 $\textit{min}$ 和 $\textit{max}$ 分别为区间的最小值和最大值，只要 $\textit{sum}[i]$ 在 $[\textit{min},\textit{max}]$ 范围中，区间就一定存在等于 $\textit{sum}[i]$ 的数。

用 Lazy 线段树维护区间最小值、区间最大值、区间加的 Lazy tag。

完整线段树模板见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 手写 min max 更快
min = lambda a, b: b if b < a else a
max = lambda a, b: b if b > a else a

class Node:
    __slots__ = 'val', 'todo'

class LazySegmentTree:
    # 懒标记初始值
    _TODO_INIT = 0

    def __init__(self, arr, default=0):
        # 线段树维护一个长为 n 的数组（下标从 0 到 n-1）
        # arr 可以是 list 或者 int
        # 如果 arr 是 int，视作数组大小，默认值为 default
        if isinstance(arr, int):
            arr = [default] * arr
        n = len(arr)
        self._n = n
        self._tree = [Node() for _ in range(2 << (n - 1).bit_length())]
        self._build(arr, 1, 0, n - 1)

    # 合并两个 val
    def _merge_val(self, a, b):
        return [min(a[0], b[0]), max(a[1], b[1])]

    # 合并两个懒标记
    def _merge_todo(self, a: int, b: int) -> int:
        return a + b

    # 把懒标记作用到 node 子树
    def _apply(self, node: int, l: int, r: int, todo: int) -> None:
        cur = self._tree[node]
        # 计算 tree[node] 区间的整体变化
        cur.val[0] += todo
        cur.val[1] += todo
        cur.todo = self._merge_todo(todo, cur.todo)

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
            self._tree[node].val = [a[l], a[l]]  # 初始化叶节点的值
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

    def _find_first(self, node: int, l: int, r: int, ql: int, qr: int, target: int) -> int:
        if l > qr or r < ql or not self._tree[node].val[0] <= target <= self._tree[node].val[1]:
            return -1
        if l == r:
            return l
        self._spread(node, l, r)
        m = (l + r) // 2
        idx = self._find_first(node * 2, l, m, ql, qr, target)
        if idx < 0:
            # 去右子树找
            idx = self._find_first(node * 2 + 1, m + 1, r, ql, qr, target)
        return idx

    # 用 f 更新 [ql, qr] 中的每个 a[i]
    # 0 <= ql <= qr <= n-1
    # 时间复杂度 O(log n)
    def update(self, ql: int, qr: int, f: int) -> None:
        self._update(1, 0, self._n - 1, ql, qr, f)

    # 查询 [ql,qr] 内第一个等于 target 的元素下标
    # 找不到返回 -1
    # 时间复杂度 O(log n)
    def find_first(self, ql: int, qr: int, target: int) -> int:
        return self._find_first(1, 0, self._n - 1, ql, qr, target)

class Solution:
    def longestBalanced(self, nums: List[int]) -> int:
        n = len(nums)
        t = LazySegmentTree(n + 1)

        last = {}  # nums 的元素上一次出现的位置
        ans = cur_sum = 0
        for i, x in enumerate(nums, 1):
            v = 1 if x % 2 else -1
            j = last.get(x, 0)
            if j == 0:  # 首次遇到 x
                cur_sum += v
                t.update(i, n, v)  # sum[i:] 增加 v
            else:  # 再次遇到 x
                t.update(j, i - 1, -v)  # 撤销之前对 sum[j:i] 的增加
            last[x] = i

            j = t.find_first(0, i - 1, cur_sum)
            if j >= 0:
                ans = max(ans, i - j)
        return ans
```

```java [sol-Java]
class LazySegmentTree {
    // 懒标记初始值
    private static final int TODO_INIT = 0;

    private static final class Node {
        int[] val;
        int todo;
    }

    // 合并两个 val
    private int[] mergeVal(int[] a, int[] b) {
        return new int[]{Math.min(a[0], b[0]), Math.max(a[1], b[1])};
    }

    // 合并两个懒标记
    private int mergeTodo(int a, int b) {
        return a + b;
    }

    // 把懒标记作用到 node 子树（本例为区间加）
    private void apply(int node, int l, int r, int todo) {
        Node cur = tree[node];
        // 计算 tree[node] 区间的整体变化
        cur.val[0] += todo;
        cur.val[1] += todo;
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

    // 查询 [ql,qr] 内第一个等于 target 的元素下标
    // 找不到返回 -1
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    public int findFirst(int ql, int qr, int target) {
        return findFirst(1, 0, n - 1, ql, qr, target);
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
            tree[node].val = new int[]{a[l], a[l]}; // 初始化叶节点的值
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

    private int findFirst(int node, int l, int r, int ql, int qr, int target) {
        if (l > qr || r < ql || target < tree[node].val[0] || target > tree[node].val[1]) {
            return -1;
        }
        if (l == r) {
            return l;
        }
        spread(node, l, r);
        int m = (l + r) / 2;
        int idx = findFirst(node * 2, l, m, ql, qr, target);
        if (idx < 0) {
            idx = findFirst(node * 2 + 1, m + 1, r, ql, qr, target);
        }
        return idx;
    }
}

class Solution {
    public int longestBalanced(int[] nums) {
        int n = nums.length;
        LazySegmentTree t = new LazySegmentTree(n + 1, 0);

        Map<Integer, Integer> last = new HashMap<>();
        int ans = 0;
        int curSum = 0;

        for (int i = 1; i <= n; i++) {
            int x = nums[i - 1];
            int v = x % 2 > 0 ? 1 : -1;
            if (!last.containsKey(x)) {
                curSum += v;
                t.update(i, n, v);
            } else {
                int j = last.get(x);
                t.update(j, i - 1, -v);
            }
            last.put(x, i);

            int j = t.findFirst(0, i - 1, curSum);
            if (j >= 0) {
                ans = Math.max(ans, i - j);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class LazySegmentTree {
    using T = pair<int, int>;
    using F = int;

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
        return {min(a.first, b.first), max(a.second, b.second)};
    }

    // 合并两个懒标记
    F merge_todo(F a, F b) const {
        return a + b;
    }

    // 把懒标记作用到 node 子树（本例为区间加）
    void apply(int node, int l, int r, F todo) {
        Node& cur = tree[node];
        // 计算 tree[node] 区间的整体变化
        cur.val.first += todo;
        cur.val.second += todo;
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

    // 查询 [ql,qr] 内第一个等于 target 的元素下标
    // 找不到返回 -1
    int find_first(int node, int l, int r, int ql, int qr, int target) {
        // 不相交 或 target 不在当前区间的 [min,max] 范围内
        if (l > qr || r < ql || target < tree[node].val.first || target > tree[node].val.second) {
            return -1;
        }
        if (l == r) {
            // 此处必然等于 target
            return l;
        }
        spread(node, l, r);
        int m = (l + r) / 2;
        int idx = find_first(node * 2, l, m, ql, qr, target);
        if (idx < 0) {
            // 去右子树找
            idx = find_first(node * 2 + 1, m + 1, r, ql, qr, target);
        }
        return idx;
    }

public:
    // 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
    LazySegmentTree(int n, T init_val = {0, 0}) : LazySegmentTree(vector<T>(n, init_val)) {}

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

    // 查询 [ql,qr] 内第一个等于 target 的元素下标
    // 找不到返回 -1
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    int find_first(int ql, int qr, int target) {
        return find_first(1, 0, n - 1, ql, qr, target);
    }
};

class Solution {
public:
    int longestBalanced(vector<int>& nums) {
        int n = nums.size();
        LazySegmentTree t(n + 1);

        unordered_map<int, int> last; // nums 的元素上一次出现的位置
        int ans = 0, cur_sum = 0;
        for (int i = 1; i <= n; i++) {
            int x = nums[i - 1];
            int v = x % 2 ? 1 : -1;
            auto it = last.find(x);
            if (it == last.end()) { // 首次遇到 x
                cur_sum += v;
                t.update(i, n, v); // sum 的 [i,n] 增加 v
            } else { // 再次遇到 x
                int j = it->second;
                t.update(j, i - 1, -v);  // 撤销之前对 sum 的 [j,i-1] 的增加
            }
            last[x] = i;

            // 在 [0,i-1] 中找第一个等于 cur_sum 的下标
            int j = t.find_first(0, i - 1, cur_sum);
            if (j >= 0) {
                ans = max(ans, i - j);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type pair struct{ min, max int }
type lazySeg []struct {
	l, r int
	pair
	todo int
}

func merge(l, r pair) pair {
	return pair{min(l.min, r.min), max(l.max, r.max)}
}

func (t lazySeg) apply(o int, f int) {
	cur := &t[o]
	cur.min += f
	cur.max += f
	cur.todo += f
}

func (t lazySeg) maintain(o int) {
	t[o].pair = merge(t[o<<1].pair, t[o<<1|1].pair)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t lazySeg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
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

// 查询 [l,r] 内第一个等于 target 的元素下标
func (t lazySeg) findFirst(o, l, r, target int) int {
	if t[o].l > r || t[o].r < l || target < t[o].min || target > t[o].max {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	idx := t.findFirst(o<<1, l, r, target)
	if idx < 0 {
		// 去右子树找
		idx = t.findFirst(o<<1|1, l, r, target)
	}
	return idx
}

func longestBalanced(nums []int) (ans int) {
	n := len(nums)
	t := make(lazySeg, 2<<bits.Len(uint(n)))
	t.build(1, 0, n)

	last := map[int]int{} // nums 的元素上一次出现的位置
	curSum := 0
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		v := x%2*2 - 1
		if j := last[x]; j == 0 { // 首次遇到 x
			curSum += v
			t.update(1, i, n, v) // sum[i:] 增加 v
		} else { // 再次遇到 x
			t.update(1, j, i-1, -v) // 撤销之前对 sum[j:i] 的增加
		}
		last[x] = i

		j := t.findFirst(1, 0, i-1, curSum)
		if j >= 0 {
			ans = max(ans, i-j)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

[HH 的项链](https://www.luogu.com.cn/problem/P1972)

## 专题训练

见下面数据结构题单的「**§8.4 Lazy 线段树**」和「**十、根号算法**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
