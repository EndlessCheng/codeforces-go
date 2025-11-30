## 前置知识

1. **同余**：[模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。
2. **中位数贪心**：把区间内的数都变成区间的**中位数**是最优的。[证明](https://zhuanlan.zhihu.com/p/1922938031687595039)。
3. **距离和**：[图解距离和](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)。
4. **可持久化线段树**求区间中位数（第 $k$ 小）：[视频讲解](https://www.bilibili.com/video/BV1D4SiB5Ee3/)。

## 什么情况下无解？

比如 $k=2$，那么偶数无论如何操作，仍然是偶数；奇数无论如何操作，仍然是奇数。在这种情况下，区间内所有元素的奇偶性必须都相同。

一般地，区间内所有元素必须都关于模 $k$ **同余**。

定义 $\textit{left}[i]$ 表示区间 $[\textit{left}[i],i]$ 中的元素都与 $\textit{nums}[i]$ 关于模 $k$ 同余，且 $\textit{left}[i]$ 尽量小。

对于询问的区间 $[l,r]$，我们只需判断 $\textit{left}[r]\le l$ 是否成立，不成立就无解。

根据定义，我们有

$$
\textit{left}[i] =
\begin{cases}
\textit{left}[i-1], & i>0\ 且\ \textit{nums}[i]\bmod k = \textit{nums}[i-1]\bmod k    \\
i, & 其他     \\
\end{cases}
$$

## 中位数贪心

如果有解，那么把区间内的数都变成区间的中位数是最优的。[证明](https://zhuanlan.zhihu.com/p/1922938031687595039)。

**推荐先完成相关题目**：

- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- [2033. 获取单值网格的最小操作数](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/)

## 距离和

设区间的中位数为 $m$。我们还需要计算区间内的元素到 $m$ 的距离之和 $s$，那么 $\dfrac{s}{k}$ 就是最小操作次数。

**推荐先完成相关题目**：

- [2602. 使数组元素全部相等的最少操作次数](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/)

本题询问的是 $\textit{nums}$ 的子数组，子数组内的元素并不是有序的。在这种情况下，如何计算区间的中位数？如何计算区间内的元素到中位数的距离之和？

## 可持久化线段树

类似 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)，考虑对 $\textit{nums}$ 的每个前缀建立一棵**值域线段树**。两棵线段树的差，就是一个子数组对应的的线段树。我们可以在这棵线段树上求出第 $k$ 小（$k$ 从 $1$ 开始）。

但对每个前缀都建立一棵值域线段树，时间空间都是 $\mathcal{O}(n^2)$ 的，这太大了。

类似 Git，考虑我们**在上一个版本的基础上，修改了什么**。在线段树上把一个数的出现次数加一（单点修改），只会更新 $\mathcal{O}(\log n)$ 个节点，其余节点保存的内容是不变的。所以每次只发生了 $\mathcal{O}(\log n)$ 个变动。把这些变动记录下来。

设区间的长度为 $\textit{sz} =r-l+1$，那么区间中位数就是区间第 $\left\lfloor\dfrac{sz}{2}\right\rfloor+1$ 小。如果有两个中位数，取左边的还是右边的都可以，这里算的是右边那个。

本题还需要算距离和，从 [图解距离和](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/) 中我们知道，关键是求出有多少个数 $\le $ 中位数，以及这些数的元素和，这同样可以用可持久化线段树解决。

[本题视频讲解（3D 视角）](https://www.bilibili.com/video/BV1D4SiB5Ee3/)，欢迎点赞关注~

```py [sol-Python3]
class Node:
    __slots__ = 'l', 'r', 'lo', 'ro', 'cnt', 'sum'

    def __init__(self, l: int, r: int, lo=None, ro=None, cnt=0, sum=0):
        self.l = l
        self.r = r
        self.lo = lo
        self.ro = ro
        self.cnt = cnt
        self.sum = sum

    def maintain(self):
        self.cnt = self.lo.cnt + self.ro.cnt
        self.sum = self.lo.sum + self.ro.sum

    @staticmethod
    def build(l: int, r: int) -> 'Node':
        o = Node(l, r)
        if l == r:
            return o
        mid = (l + r) // 2
        o.lo = Node.build(l, mid)
        o.ro = Node.build(mid + 1, r)
        return o

    # 在线段树的位置 i 添加 val
    def add(self, i: int, val: int) -> 'Node':
        # 复制一份当前节点
        o = Node(self.l, self.r, self.lo, self.ro, self.cnt, self.sum)
        if o.l == o.r:
            o.cnt += 1
            o.sum += val
            return o
        mid = (o.l + o.r) // 2
        if i <= mid:
            o.lo = o.lo.add(i, val)
        else:
            o.ro = o.ro.add(i, val)
        o.maintain()
        return o

    # 查询 old 和 self 对应区间的第 k 小，k 从 1 开始
    def kth(self, old: 'Node', k: int) -> int:
        if self.l == self.r:
            return self.l
        cnt_l = self.lo.cnt - old.lo.cnt
        if k <= cnt_l:  # 答案在左子树中
            return self.lo.kth(old.lo, k)
        return self.ro.kth(old.ro, k - cnt_l)  # 答案在右子树中

    # 查询 old 和 self 对应区间，有多少个数 <= i，这些数的元素和是多少
    def query(self, old: 'Node', i: int) -> Tuple[int, int]:
        if self.r <= i:
            return self.cnt - old.cnt, self.sum - old.sum
        cnt, sum_ = self.lo.query(old.lo, i)
        mid = (self.l + self.r) // 2
        if i > mid:
            c, s = self.ro.query(old.ro, i)
            cnt += c
            sum_ += s
        return cnt, sum_


class Solution:
    def minOperations(self, nums: List[int], k: int, queries: List[List[int]]) -> List[int]:
        n = len(nums)
        left = [0] * n
        for i in range(1, n):
            left[i] = left[i - 1] if nums[i] % k == nums[i - 1] % k else i

        # 准备离散化
        sorted_nums = sorted(set(nums))
        mp = {x: i for i, x in enumerate(sorted_nums)}

        # 构建可持久化线段树
        t = [None] * (n + 1)
        t[0] = Node.build(0, len(sorted_nums) - 1)
        for i, x in enumerate(nums):
            j = mp[x]  # 离散化
            t[i + 1] = t[i].add(j, x)

        ans = []
        for l, r in queries:
            if left[r] > l:  # 无解
                ans.append(-1)
                continue

            r += 1  # 改成左闭右开，方便计算

            # 计算区间中位数
            sz = r - l
            i = t[r].kth(t[l], sz // 2 + 1)
            median = sorted_nums[i]  # 离散化后的值 -> 原始值

            # 计算区间所有元素到中位数的距离和
            total = t[r].sum - t[l].sum  # 区间元素和
            cnt_left, sum_left = t[r].query(t[l], i)
            s = median * cnt_left - sum_left  # 蓝色面积
            s += total - sum_left - median * (sz - cnt_left)  # 绿色面积
            ans.append(s // k)  # 操作次数 = 距离和 / k

        return ans
```

```java [sol-Java]
class Node {
    private final int l;
    private final int r;
    private Node lo;
    private Node ro;
    private int cnt;
    public long sum;

    private void maintain() {
        cnt = lo.cnt + ro.cnt;
        sum = lo.sum + ro.sum;
    }

    public Node(int l, int r, Node lo, Node ro, int cnt, long sum) {
        this.l = l;
        this.r = r;
        this.lo = lo;
        this.ro = ro;
        this.cnt = cnt;
        this.sum = sum;
    }

    public Node(int l, int r) {
        this(l, r, null, null, 0, 0);
    }

    public static Node build(int l, int r) {
        Node o = new Node(l, r);
        if (l == r) {
            return o;
        }
        int mid = (l + r) / 2;
        o.lo = build(l, mid);
        o.ro = build(mid + 1, r);
        return o;
    }

    // 在线段树的位置 i 添加 val
    public Node add(int i, int val) {
        Node o = new Node(l, r, lo, ro, cnt, sum); // 复制当前节点
        if (l == r) {
            o.cnt++;
            o.sum += val;
            return o;
        }
        int mid = (l + r) / 2;
        if (i <= mid) {
            o.lo = o.lo.add(i, val);
        } else {
            o.ro = o.ro.add(i, val);
        }
        o.maintain();
        return o;
    }

    // 查询 old 和 this 对应区间的第 k 小，k 从 1 开始
    public int kth(Node old, int k) {
        if (l == r) {
            return l;
        }
        int cntL = lo.cnt - old.lo.cnt;
        if (k <= cntL) { // 答案在左子树中
            return lo.kth(old.lo, k);
        }
        return ro.kth(old.ro, k - cntL); // 答案在右子树中
    }

    // 查询 old 和 this 对应区间，有多少个数 <= i，这些数的元素和是多少
    public long[] query(Node old, int i) {
        if (r <= i) {
            return new long[]{cnt - old.cnt, sum - old.sum};
        }
        long[] left = lo.query(old.lo, i);
        long cnt = left[0];
        long s = left[1];
        int mid = (l + r) / 2;
        if (i > mid) {
            long[] right = ro.query(old.ro, i);
            cnt += right[0];
            s += right[1];
        }
        return new long[]{cnt, s};
    }
}

class Solution {
    public long[] minOperations(int[] nums, int k, int[][] queries) {
        int n = nums.length;
        int[] left = new int[n];
        for (int i = 1; i < n; i++) {
            left[i] = nums[i] % k == nums[i - 1] % k ? left[i - 1] : i;
        }

        // 准备离散化
        int[] sorted = nums.clone();
        Arrays.sort(sorted);

        // 构建可持久化线段树
        Node[] t = new Node[n + 1];
        t[0] = Node.build(0, n - 1);
        for (int i = 0; i < n; i++) {
            int j = Arrays.binarySearch(sorted, nums[i]);
            t[i + 1] = t[i].add(j, nums[i]);
        }

        long[] ans = new long[queries.length];

        for (int qi = 0; qi < queries.length; qi++) {
            int[] q = queries[qi];
            int l = q[0];
            int r = q[1];
            if (left[r] > l) { // 无解
                ans[qi] = -1;
                continue;
            }

            r++; // 改成左闭右开

            // 计算区间中位数
            int sz = r - l;
            int i = t[r].kth(t[l], sz / 2 + 1);
            long median = sorted[i]; // 离散化后的值 -> 原始值

            // 计算区间所有元素到中位数的距离和
            long totalSum = t[r].sum - t[l].sum;
            long[] res = t[r].query(t[l], i);
            long cntLeft = res[0];
            long sumLeft = res[1];
            long sum = median * cntLeft - sumLeft; // 蓝色面积
            sum += totalSum - sumLeft - median * (sz - cntLeft); // 绿色面积
            ans[qi] = sum / k; // 操作次数 = 距离和 / k
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Node {
    int l, r;
    Node* lo;
    Node* ro;
    int cnt;

    void maintain() {
        cnt = lo->cnt + ro->cnt;
        sum = lo->sum + ro->sum;
    }

public:
    long long sum;

    Node(int l, int r, Node* lo = nullptr, Node* ro = nullptr, long long cnt = 0, long long sum = 0)
        : l(l), r(r), lo(lo), ro(ro), cnt(cnt), sum(sum) {}

    static Node* build(int l, int r) {
        Node* o = new Node(l, r);
        if (l == r) {
            return o;
        }
        int mid = (l + r) / 2;
        o->lo = build(l, mid);
        o->ro = build(mid + 1, r);
        return o;
    }

    // 在线段树的位置 i 添加 val
    Node* add(int i, int val) {
        Node* o = new Node(l, r, lo, ro, cnt, sum); // 复制当前节点
        if (l == r) {
            o->cnt++;
            o->sum += val;
            return o;
        }
        int mid = (l + r) / 2;
        if (i <= mid) {
            o->lo = o->lo->add(i, val);
        } else {
            o->ro = o->ro->add(i, val);
        }
        o->maintain();
        return o;
    }

    // 查询 old 和当前节点对应区间的第 k 小，k 从 1 开始
    int kth(Node* old, int k) {
        if (l == r) {
            return l;
        }
        int cnt_l = lo->cnt - old->lo->cnt;
        if (k <= cnt_l) { // 答案在左子树中
            return lo->kth(old->lo, k);
        }
        return ro->kth(old->ro, k - cnt_l); // 答案在右子树中
    }

    // 查询 old 和当前节点对应区间，有多少个数 <= i，这些数的元素和是多少
    pair<int, long long> query(Node* old, int i) {
        if (r <= i) {
            return {cnt - old->cnt, sum - old->sum};
        }
        auto [cnt, sum] = lo->query(old->lo, i);
        int mid = (l + r) / 2;
        if (i > mid) {
            auto [c, s] = ro->query(old->ro, i);
            cnt += c;
            sum += s;
        }
        return {cnt, sum};
    }
};

class Solution {
public:
    vector<long long> minOperations(vector<int>& nums, int k, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<int> left(n);
        for (int i = 1; i < n; i++) {
            left[i] = nums[i] % k == nums[i - 1] % k ? left[i - 1] : i;
        }

        // 准备离散化
        vector<int> sorted_nums = nums;
        ranges::sort(sorted_nums);
        sorted_nums.erase(ranges::unique(sorted_nums).begin(), sorted_nums.end());
        int m = sorted_nums.size();

        // 构建可持久化线段树
        vector<Node*> t(n + 1);
        t[0] = Node::build(0, m - 1);
        for (int i = 0; i < n; i++) {
            int j = ranges::lower_bound(sorted_nums, nums[i]) - sorted_nums.begin();
            t[i + 1] = t[i]->add(j, nums[i]);
        }

        vector<long long> ans;
        ans.reserve(queries.size()); // 预分配空间

        for (auto& q : queries) {
            int l = q[0], r = q[1];
            if (left[r] > l) { // 无解
                ans.push_back(-1);
                continue;
            }

            r++; // 改成左闭右开，方便计算

            // 计算区间中位数
            int sz = r - l;
            int i = t[r]->kth(t[l], sz / 2 + 1);
            long long median = sorted_nums[i]; // 离散化后的值 -> 原始值

            // 计算区间所有元素到中位数的距离和
            long long total = t[r]->sum - t[l]->sum; // 区间元素和
            auto [cnt_left, sum_left] = t[r]->query(t[l], i);
            long long sum = median * cnt_left - sum_left; // 蓝色面积
            sum += total - sum_left - median * (sz - cnt_left); // 绿色面积
            ans.push_back(sum / k); // 操作次数 = 距离和 / k
        }

        // 省略 delete 线段树节点的代码
        return ans;
    }
};
```

```go [sol-Go]
// 有大量指针的题目，关闭 GC 更快
func init() { debug.SetGCPercent(-1) } 

type node struct {
	lo, ro   *node
	l, r     int
	cnt, sum int
}

func (o *node) maintain() {
	o.cnt = o.lo.cnt + o.ro.cnt
	o.sum = o.lo.sum + o.ro.sum
}

func build(l, r int) *node {
	o := &node{l: l, r: r}
	if l == r {
		return o
	}
	mid := (l + r) / 2
	o.lo = build(l, mid)
	o.ro = build(mid+1, r)
	return o
}

// 在线段树的位置 i 添加 val
// 注意这里传的不是指针，会把 node 复制一份，而这正好是我们需要的
func (o node) add(i, val int) *node {
	if o.l == o.r {
		o.cnt++
		o.sum += val
		return &o
	}
	mid := (o.l + o.r) / 2
	if i <= mid {
		o.lo = o.lo.add(i, val)
	} else {
		o.ro = o.ro.add(i, val)
	}
	o.maintain()
	return &o
}

// 查询 old 和 o 对应区间的第 k 小，k 从 1 开始
func (o *node) kth(old *node, k int) int {
	if o.l == o.r {
		return o.l
	}
	cntL := o.lo.cnt - old.lo.cnt
	if k <= cntL { // 答案在左子树中
		return o.lo.kth(old.lo, k)
	}
	return o.ro.kth(old.ro, k-cntL) // 答案在右子树中
}

// 查询 old 和 o 对应区间，有多少个数 <= i，这些数的元素和是多少
func (o *node) query(old *node, i int) (int, int) {
	if o.r <= i {
		return o.cnt - old.cnt, o.sum - old.sum
	}
	cnt, tot := o.lo.query(old.lo, i)
	mid := (o.l + o.r) / 2
	if i > mid {
		c, t := o.ro.query(old.ro, i)
		cnt += c
		tot += t
	}
	return cnt, tot
}

func minOperations(nums []int, k int, queries [][]int) []int64 {
	n := len(nums)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i]%k != nums[i-1]%k {
			left[i] = i
		} else {
			left[i] = left[i-1]
		}
	}

	// 准备离散化
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	t := make([]*node, n+1)
	t[0] = build(0, len(sorted)-1)
	for i, x := range nums {
		j := sort.SearchInts(sorted, x) // 离散化
		t[i+1] = t[i].add(j, x)         // 构建可持久化线段树
	}

	ans := make([]int64, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		if left[r] > l { // 无解
			ans[qi] = -1
			continue
		}

		r++ // 改成左闭右开，方便计算

		// 计算区间中位数
		sz := r - l
		i := t[r].kth(t[l], sz/2+1)
		median := sorted[i] // 离散化后的值 -> 原始值

		// 计算区间所有元素到中位数的距离和
		total := t[r].sum - t[l].sum // 区间元素和
		cntLeft, sumLeft := t[r].query(t[l], i)
		sum := median*cntLeft - sumLeft              // 蓝色面积
		sum += total - sumLeft - median*(sz-cntLeft) // 绿色面积
		ans[qi] = int64(sum / k)                     // 操作次数 = 距离和 / k
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log n)$。返回值不计入。

## 专题训练

1. 贪心题单的「**§4.5 中位数贪心**」。
2. 数据结构题单的「**§1.3 距离和**」。

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
