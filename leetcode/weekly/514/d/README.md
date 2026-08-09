**技巧**：如果一个问题可以用**分治**解决，那么这个问题的带修改版本可以用**线段树**解决。

如何用分治计算一个数组的峰值子数组的个数？

把数组分成左右两部分 $A$ 和 $B$：

1. 计算 $A$ 中的峰值子数组的个数。这可以递归解决。
2. 计算 $B$ 中的峰值子数组的个数。这可以递归解决。
3. 计算左端点在 $A$ 中，右端点在 $B$ 中的峰值子数组的个数。我们需要知道 $A$ 中的最后一个峰顶的位置，以及 $B$ 中第一个峰顶的位置。或者计算从峰顶到数组端点的长度。用所有子数组的个数，减去不含峰值的子数组的个数，即为峰值子数组的个数。

这样思考后，发现本题可以用分治解决，那么带修改版本就可以用线段树解决了。

对于线段树的区间 $[L,R]$，维护如下信息：

- $\textit{cnt}$：该区间的峰值子数组的个数。
- $\textit{has}$：该区间是否有峰顶。
- $\textit{pre}$：从 $L$ 到该区间第一个峰顶的长度。如果该区间没有峰顶，则 $\textit{pre}=R-L+1$。
- $\textit{suf}$：从该区间最后一个峰顶到 $R$ 的长度。如果该区间没有峰顶，则 $\textit{suf}=R-L+1$。
- $\textit{len}$：为方便计算，顺带维护区间长度。

合并区间 $A$ 和 $B$ 时，如何计算左端点在 $A$ 中，右端点在 $B$ 中的峰值子数组的个数？

一共有 $A.\textit{len} \cdot B.\textit{len}$ 个横跨 $A$ 和 $B$ 的子数组，其中不含峰顶的子数组，左端点有 $A.\textit{suf}$ 个，右端点有 $B.\textit{pre}$ 个（注意峰顶必须严格在子数组内部，不能在子数组端点上），所以一共有 $A.\textit{suf}\cdot B.\textit{pre}$ 个不含峰顶的子数组。所以左端点在 $A$ 中，右端点在 $B$ 中的峰值子数组的个数为

$$
A.\textit{len} \cdot B.\textit{len} - A.\textit{suf}\cdot B.\textit{pre}
$$

合并区间后，大区间的峰值子数组的个数为

$$
A.\textit{cnt} + B.\textit{cnt} + A.\textit{len} \cdot B.\textit{len} - A.\textit{suf}\cdot B.\textit{pre}
$$

对于更新操作，由于更新 $\textit{nums}[i]$ 会影响 $i-1,i,i+1$ 三个位置是否为峰顶，我们执行三次线段树的单点更新操作。

[本题视频讲解](https://www.bilibili.com/video/BV1ryuy6WEDs/?t=22m22s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
# 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree:
    def __init__(self, a: list[int]):
        self._n = n = len(a)
        self._tree = [None] * (2 << (n - 1).bit_length())
        self._build(a, 1, 0, n - 1)

    def _merge_data(self, a: list, b: list) -> list:
        a_cnt, a_pre, a_suf, a_len, a_has = a
        b_cnt, b_pre, b_suf, b_len, b_has = b
        cnt = a_cnt + b_cnt + a_len * b_len - a_suf * b_pre
        pre = a_pre if a_has else a_len + b_pre
        suf = b_suf if b_has else b_len + a_suf
        return [cnt, pre, suf, a_len + b_len, a_has or b_has]

    def _maintain(self, node: int) -> None:
        self._tree[node] = self._merge_data(self._tree[node * 2], self._tree[node * 2 + 1])

    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        if l == r:  # 叶子
            has_peak = 0 < l < self._n - 1 and a[l - 1] < a[l] > a[l + 1]
            self._tree[node] = [0, 1, 1, 1, has_peak]  # 初始化叶节点的值
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)  # 初始化左子树
        self._build(a, node * 2 + 1, m + 1, r)  # 初始化右子树
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, i: int, has_peak: bool) -> None:
        if l == r:  # 叶子（到达目标）
            self._tree[node][-1] = has_peak
            return
        m = (l + r) // 2
        if i <= m:  # i 在左子树
            self._update(node * 2, l, m, i, has_peak)
        else:  # i 在右子树
            self._update(node * 2 + 1, m + 1, r, i, has_peak)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> list:
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            return self._tree[node]
        m = (l + r) // 2
        if qr <= m:  # [ql, qr] 与右子树无交集，仅需递归左子树
            return self._query(node * 2, l, m, ql, qr)
        if ql > m:  # [ql, qr] 与左子树无交集，仅需递归右子树
            return self._query(node * 2 + 1, m + 1, r, ql, qr)
        # [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
        return self._merge_data(self._query(node * 2, l, m, ql, qr), self._query(node * 2 + 1, m + 1, r, ql, qr))

    def update(self, i: int, val: bool) -> None:
        self._update(1, 0, self._n - 1, i, val)

    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)[0]


class Solution:
    def countOfPeaks(self, nums: list[int], queries: list[list[int]]) -> list[int]:
        n = len(nums)
        t = SegmentTree(nums)
        ans = []
        for op, i, v in queries:
            if op == 1:
                ans.append(t.query(i, v))
                continue
            nums[i] = v
            for j in range(max(i - 1, 1), min(i + 2, n - 1)):
                # 注：这里可以优化一下，如果更新前后 has_peak 不变，则不调用 t.update
                has_peak = nums[j - 1] < nums[j] > nums[j + 1]
                t.update(j, has_peak)
        return ans
```

```java [sol-Java]
// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    private record Data(long cnt, int pre, int suf, int len, boolean hasPeak) {
    }

    private final int n;
    private final Data[] tree;

    private Data mergeData(Data a, Data b) {
        long cnt = a.cnt + b.cnt + (long) a.len * b.len - (long) a.suf * b.pre;
        int pre = a.hasPeak ? a.pre : a.len + b.pre;
        int suf = b.hasPeak ? b.suf : b.len + a.suf;
        return new Data(cnt, pre, suf, a.len + b.len, a.hasPeak || b.hasPeak);
    }

    public SegmentTree(int[] a) {
        n = a.length;
        tree = new Data[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(a, 1, 0, n - 1);
    }

    public void update(int i, boolean hasPeak) {
        update(1, 0, n - 1, i, hasPeak);
    }

    public long query(int ql, int qr) {
        return query(1, 0, n - 1, ql, qr).cnt;
    }

    private void maintain(int node) {
        tree[node] = mergeData(tree[node * 2], tree[node * 2 + 1]);
    }

    private void build(int[] a, int node, int l, int r) {
        if (l == r) { // 叶子
            boolean hasPeak = 0 < l && l < n - 1 && a[l - 1] < a[l] && a[l] > a[l + 1];
            tree[node] = new Data(0, 1, 1, 1, hasPeak); // 初始化叶节点的值
            return;
        }
        int m = (l + r) >>> 1;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    private void update(int node, int l, int r, int i, boolean hasPeak) {
        if (l == r) { // 叶子（到达目标）
            Data d = tree[node];
            tree[node] = new Data(d.cnt, d.pre, d.suf, d.len, hasPeak);
            return;
        }
        int m = (l + r) >>> 1;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i, hasPeak);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i, hasPeak);
        }
        maintain(node);
    }

    private Data query(int node, int l, int r, int ql, int qr) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node];
        }
        int m = (l + r) >>> 1;
        if (qr <= m) { // [ql, qr] 与右子树无交集，仅需递归左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 与左子树无交集，仅需递归右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        // [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
        Data lRes = query(node * 2, l, m, ql, qr);
        Data rRes = query(node * 2 + 1, m + 1, r, ql, qr);
        return mergeData(lRes, rRes);
    }
}

class Solution {
    public long[] countOfPeaks(int[] nums, int[][] queries) {
        int cnt1 = 0;
        for (int[] q : queries) {
            cnt1 += 2 - q[0];
        }

        int n = nums.length;
        SegmentTree t = new SegmentTree(nums);
        long[] ans = new long[cnt1];
        int k = 0;
        for (int[] q : queries) {
            if (q[0] == 1) {
                ans[k++] = t.query(q[1], q[2]);
                continue;
            }
            int i = q[1];
            nums[i] = q[2];
            for (int j = Math.max(i - 1, 1); j <= Math.min(i + 1, n - 2); j++) {
                // 注：这里可以优化一下，如果更新前后 hasPeak 不变，则不调用 t.update
                boolean hasPeak = nums[j - 1] < nums[j] && nums[j] > nums[j + 1];
                t.update(j, hasPeak);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
struct Data {
    long long cnt;
    int pre, suf, len;
    bool has_peak;
};

// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    int n;
    vector<Data> tree;

    Data merge_data(Data a, Data b) const {
        long long cnt = a.cnt + b.cnt + 1LL * a.len * b.len - 1LL * a.suf * b.pre;
        int pre = a.has_peak ? a.pre : a.len + b.pre;
        int suf = b.has_peak ? b.suf : b.len + a.suf;
        return {cnt, pre, suf, a.len + b.len, a.has_peak || b.has_peak};
    }

    void maintain(int node) {
        tree[node] = merge_data(tree[node * 2], tree[node * 2 + 1]);
    }

    void build(const vector<int>& a, int node, int l, int r) {
        if (l == r) { // 叶子
            bool has_peak = 0 < l && l < n - 1 && a[l - 1] < a[l] && a[l] > a[l + 1];
            tree[node] = {0, 1, 1, 1, has_peak}; // 初始化叶节点的值
            return;
        }
        int m = (l + r) >> 1;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    void update(int node, int l, int r, int i, bool has_peak) {
        if (l == r) { // 叶子（到达目标）
            tree[node].has_peak = has_peak;
            return;
        }
        int m = (l + r) >> 1;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i, has_peak);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i, has_peak);
        }
        maintain(node);
    }

    Data query(int node, int l, int r, int ql, int qr) const {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node];
        }
        int m = (l + r) >> 1;
        if (qr <= m) { // [ql, qr] 与右子树无交集，仅需递归左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 与左子树无交集，仅需递归右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        // [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
        return merge_data(query(node * 2, l, m, ql, qr), query(node * 2 + 1, m + 1, r, ql, qr));
    }

public:
    SegmentTree(const vector<int>& a) : n(a.size()), tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, n - 1);
    }

    void update(int i, bool has_peak) {
        update(1, 0, n - 1, i, has_peak);
    }

    long long query(int ql, int qr) const {
        return query(1, 0, n - 1, ql, qr).cnt;
    }
};

class Solution {
public:
    vector<long long> countOfPeaks(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        SegmentTree t(nums);
        vector<long long> ans;
        for (auto& q : queries) {
            if (q[0] == 1) {
                ans.push_back(t.query(q[1], q[2]));
                continue;
            }
            int i = q[1];
            nums[i] = q[2];
            for (int j = max(i - 1, 1); j <= min(i + 1, n - 2); j++) {
                // 注：这里可以优化一下，如果更新前后 has_peak 不变，则不调用 t.update
                bool has_peak = nums[j - 1] < nums[j] && nums[j] > nums[j + 1];
                t.update(j, has_peak);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type data struct {
	cnt, pre, suf, len int
	hasPeak            bool
}

// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
type seg []data

func newSegmentTree(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func (seg) mergeData(a, b data) data {
	cnt := a.cnt + b.cnt + a.len*b.len - a.suf*b.pre
	pre := a.pre
	if !a.hasPeak {
		pre += b.pre
	}
	suf := b.suf
	if !b.hasPeak {
		suf += a.suf
	}
	return data{cnt, pre, suf, a.len + b.len, a.hasPeak || b.hasPeak}
}

func (t seg) maintain(node int) {
	t[node] = t.mergeData(t[node*2], t[node*2+1])
}

func (t seg) build(a []int, node, l, r int) {
	if l == r { // 叶子
		hasPeak := 0 < l && l < len(a)-1 && a[l-1] < a[l] && a[l] > a[l+1]
		t[node] = data{0, 1, 1, 1, hasPeak} // 初始化叶节点的值
		return
	}
	m := (l + r) >> 1
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树
	t.maintain(node)
}

func (t seg) update(node, l, r, i int, hasPeak bool) {
	if l == r { // 叶子（到达目标）
		t[node].hasPeak = hasPeak
		return
	}
	m := (l + r) >> 1
	if i <= m { // i 在左子树
		t.update(node*2, l, m, i, hasPeak)
	} else { // i 在右子树
		t.update(node*2+1, m+1, r, i, hasPeak)
	}
	t.maintain(node)
}

func (t seg) query(node, l, r, ql, qr int) data {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return t[node]
	}
	m := (l + r) >> 1
	if qr <= m { // [ql, qr] 与右子树无交集，仅需递归左子树
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 与左子树无交集，仅需递归右子树
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	// [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
	return t.mergeData(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func countOfPeaks(nums []int, queries [][]int) (ans []int64) {
	n := len(nums)
	t := newSegmentTree(nums)
	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, int64(t.query(1, 0, n-1, q[1], q[2]).cnt))
			continue
		}
		i := q[1]
		nums[i] = q[2]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			// 注：这里可以优化一下，如果更新前后 hasPeak 不变，则不调用 t.update
			hasPeak := nums[j-1] < nums[j] && nums[j] > nums[j+1]
			t.update(1, 0, n-1, j, hasPeak)
		}
	}
	return
}
```

## 优化

如果更新前后 $\textit{hasPeak}$ 不变，则不调用 `t.update`。

```py [sol-Python3]
# 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree:
    def __init__(self, a: list[int]):
        self._n = n = len(a)
        self._tree = [None] * (2 << (n - 1).bit_length())
        self._build(a, 1, 0, n - 1)

    def _merge_data(self, a: list, b: list) -> list:
        a_cnt, a_pre, a_suf, a_len, a_has = a
        b_cnt, b_pre, b_suf, b_len, b_has = b
        cnt = a_cnt + b_cnt + a_len * b_len - a_suf * b_pre
        pre = a_pre if a_has else a_len + b_pre
        suf = b_suf if b_has else b_len + a_suf
        return [cnt, pre, suf, a_len + b_len, a_has or b_has]

    def _maintain(self, node: int) -> None:
        self._tree[node] = self._merge_data(self._tree[node * 2], self._tree[node * 2 + 1])

    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        if l == r:  # 叶子
            has_peak = 0 < l < self._n - 1 and a[l - 1] < a[l] > a[l + 1]
            self._tree[node] = [0, 1, 1, 1, has_peak]  # 初始化叶节点的值
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)  # 初始化左子树
        self._build(a, node * 2 + 1, m + 1, r)  # 初始化右子树
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, i: int) -> None:
        if l == r:  # 叶子（到达目标）
            d = self._tree[node]
            d[-1] = not d[-1]
            return
        m = (l + r) // 2
        if i <= m:  # i 在左子树
            self._update(node * 2, l, m, i)
        else:  # i 在右子树
            self._update(node * 2 + 1, m + 1, r, i)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> list:
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            return self._tree[node]
        m = (l + r) // 2
        if qr <= m:  # [ql, qr] 与右子树无交集，仅需递归左子树
            return self._query(node * 2, l, m, ql, qr)
        if ql > m:  # [ql, qr] 与左子树无交集，仅需递归右子树
            return self._query(node * 2 + 1, m + 1, r, ql, qr)
        # [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
        return self._merge_data(self._query(node * 2, l, m, ql, qr), self._query(node * 2 + 1, m + 1, r, ql, qr))

    def update(self, i: int) -> None:
        self._update(1, 0, self._n - 1, i)

    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)[0]


class Solution:
    def countOfPeaks(self, nums: list[int], queries: list[list[int]]) -> list[int]:
        n = len(nums)
        t = SegmentTree(nums)
        ans = []

        for op, i, v in queries:
            if op == 1:
                ans.append(t.query(i, v))
                continue

            if i > 1:
                old_has = nums[i - 2] < nums[i - 1] > nums[i]
                new_has = nums[i - 2] < nums[i - 1] > v
                if new_has != old_has:
                    t.update(i - 1)

            if 0 < i < n - 1:
                old_has = nums[i - 1] < nums[i] > nums[i + 1]
                new_has = nums[i - 1] < v > nums[i + 1]
                if new_has != old_has:
                    t.update(i)

            if i < n - 2:
                old_has = nums[i] < nums[i + 1] > nums[i + 2]
                new_has = v < nums[i + 1] > nums[i + 2]
                if new_has != old_has:
                    t.update(i + 1)

            nums[i] = v

        return ans
```

```java [sol-Java]
// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    private record Data(long cnt, int pre, int suf, int len, boolean hasPeak) {
    }

    private final int n;
    private final Data[] tree;

    private Data mergeData(Data a, Data b) {
        long cnt = a.cnt + b.cnt + (long) a.len * b.len - (long) a.suf * b.pre;
        int pre = a.hasPeak ? a.pre : a.len + b.pre;
        int suf = b.hasPeak ? b.suf : b.len + a.suf;
        return new Data(cnt, pre, suf, a.len + b.len, a.hasPeak || b.hasPeak);
    }

    public SegmentTree(int[] a) {
        n = a.length;
        tree = new Data[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(a, 1, 0, n - 1);
    }

    public void update(int i) {
        update(1, 0, n - 1, i);
    }

    public long query(int ql, int qr) {
        return query(1, 0, n - 1, ql, qr).cnt;
    }

    private void maintain(int node) {
        tree[node] = mergeData(tree[node * 2], tree[node * 2 + 1]);
    }

    private void build(int[] a, int node, int l, int r) {
        if (l == r) { // 叶子
            boolean hasPeak = 0 < l && l < n - 1 && a[l - 1] < a[l] && a[l] > a[l + 1];
            tree[node] = new Data(0, 1, 1, 1, hasPeak); // 初始化叶节点的值
            return;
        }
        int m = (l + r) >>> 1;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    private void update(int node, int l, int r, int i) {
        if (l == r) { // 叶子（到达目标）
            Data d = tree[node];
            tree[node] = new Data(d.cnt, d.pre, d.suf, d.len, !d.hasPeak);
            return;
        }
        int m = (l + r) >>> 1;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i);
        }
        maintain(node);
    }

    private Data query(int node, int l, int r, int ql, int qr) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node];
        }
        int m = (l + r) >>> 1;
        if (qr <= m) { // [ql, qr] 与右子树无交集，仅需递归左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 与左子树无交集，仅需递归右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        // [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
        Data lRes = query(node * 2, l, m, ql, qr);
        Data rRes = query(node * 2 + 1, m + 1, r, ql, qr);
        return mergeData(lRes, rRes);
    }
}

class Solution {
    public long[] countOfPeaks(int[] nums, int[][] queries) {
        int cnt1 = 0;
        for (int[] q : queries) {
            cnt1 += 2 - q[0];
        }

        int n = nums.length;
        SegmentTree t = new SegmentTree(nums);
        long[] ans = new long[cnt1];
        int k = 0;

        for (int[] q : queries) {
            if (q[0] == 1) {
                ans[k++] = t.query(q[1], q[2]);
                continue;
            }

            int i = q[1];
            int v = q[2];

            if (i > 1) {
                boolean oldHas = nums[i - 2] < nums[i - 1] && nums[i - 1] > nums[i];
                boolean newHas = nums[i - 2] < nums[i - 1] && nums[i - 1] > v;
                if (newHas != oldHas) {
                    t.update(i - 1);
                }
            }

            if (0 < i && i < n - 1) {
                boolean oldHas = nums[i - 1] < nums[i] && nums[i] > nums[i + 1];
                boolean newHas = nums[i - 1] < v && v > nums[i + 1];
                if (newHas != oldHas) {
                    t.update(i);
                }
            }

            if (i < n - 2) {
                boolean oldHas = nums[i] < nums[i + 1] && nums[i + 1] > nums[i + 2];
                boolean newHas = v < nums[i + 1] && nums[i + 1] > nums[i + 2];
                if (newHas != oldHas) {
                    t.update(i + 1);
                }
            }

            nums[i] = v;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
struct Data {
    long long cnt;
    int pre, suf, len;
    bool has_peak;
};

// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    int n;
    vector<Data> tree;

    Data merge_data(Data a, Data b) const {
        long long cnt = a.cnt + b.cnt + 1LL * a.len * b.len - 1LL * a.suf * b.pre;
        int pre = a.has_peak ? a.pre : a.len + b.pre;
        int suf = b.has_peak ? b.suf : b.len + a.suf;
        return {cnt, pre, suf, a.len + b.len, a.has_peak || b.has_peak};
    }

    void maintain(int node) {
        tree[node] = merge_data(tree[node * 2], tree[node * 2 + 1]);
    }

    void build(const vector<int>& a, int node, int l, int r) {
        if (l == r) { // 叶子
            bool has_peak = 0 < l && l < n - 1 && a[l - 1] < a[l] && a[l] > a[l + 1];
            tree[node] = {0, 1, 1, 1, has_peak}; // 初始化叶节点的值
            return;
        }
        int m = (l + r) >> 1;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    void update(int node, int l, int r, int i) {
        if (l == r) { // 叶子（到达目标）
            tree[node].has_peak = !tree[node].has_peak;
            return;
        }
        int m = (l + r) >> 1;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i);
        }
        maintain(node);
    }

    Data query(int node, int l, int r, int ql, int qr) const {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node];
        }
        int m = (l + r) >> 1;
        if (qr <= m) { // [ql, qr] 与右子树无交集，仅需递归左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 与左子树无交集，仅需递归右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        // [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
        return merge_data(query(node * 2, l, m, ql, qr), query(node * 2 + 1, m + 1, r, ql, qr));
    }

public:
    SegmentTree(const vector<int>& a) : n(a.size()), tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, n - 1);
    }

    void update(int i) {
        update(1, 0, n - 1, i);
    }

    long long query(int ql, int qr) const {
        return query(1, 0, n - 1, ql, qr).cnt;
    }
};

class Solution {
public:
    vector<long long> countOfPeaks(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        SegmentTree t(nums);
        vector<long long> ans;

        for (auto& q : queries) {
            if (q[0] == 1) {
                ans.push_back(t.query(q[1], q[2]));
                continue;
            }

            int i = q[1];
            int v = q[2];

            if (i > 1) {
                bool old_has = nums[i - 2] < nums[i - 1] && nums[i - 1] > nums[i];
                bool new_has = nums[i - 2] < nums[i - 1] && nums[i - 1] > v;
                if (new_has != old_has) {
                    t.update(i - 1);
                }
            }

            if (0 < i && i < n - 1) {
                bool old_has = nums[i - 1] < nums[i] && nums[i] > nums[i + 1];
                bool new_has = nums[i - 1] < v && v > nums[i + 1];
                if (new_has != old_has) {
                    t.update(i);
                }
            }

            if (i < n - 2) {
                bool old_has = nums[i] < nums[i + 1] && nums[i + 1] > nums[i + 2];
                bool new_has = v < nums[i + 1] && nums[i + 1] > nums[i + 2];
                if (new_has != old_has) {
                    t.update(i + 1);
                }
            }

            nums[i] = v;
        }

        return ans;
    }
};
```

```go [sol-Go]
type data struct {
	cnt, pre, suf, len int
	hasPeak            bool
}

// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
type seg []data

func newSegmentTree(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func (seg) mergeData(a, b data) data {
	cnt := a.cnt + b.cnt + a.len*b.len - a.suf*b.pre
	pre := a.pre
	if !a.hasPeak {
		pre += b.pre
	}
	suf := b.suf
	if !b.hasPeak {
		suf += a.suf
	}
	return data{cnt, pre, suf, a.len + b.len, a.hasPeak || b.hasPeak}
}

func (t seg) maintain(node int) {
	t[node] = t.mergeData(t[node*2], t[node*2+1])
}

func (t seg) build(a []int, node, l, r int) {
	if l == r { // 叶子
		hasPeak := 0 < l && l < len(a)-1 && a[l-1] < a[l] && a[l] > a[l+1]
		t[node] = data{0, 1, 1, 1, hasPeak} // 初始化叶节点的值
		return
	}
	m := (l + r) >> 1
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树
	t.maintain(node)
}

func (t seg) update(node, l, r, i int) {
	if l == r { // 叶子（到达目标）
		t[node].hasPeak = !t[node].hasPeak
		return
	}
	m := (l + r) >> 1
	if i <= m { // i 在左子树
		t.update(node*2, l, m, i)
	} else { // i 在右子树
		t.update(node*2+1, m+1, r, i)
	}
	t.maintain(node)
}

func (t seg) query(node, l, r, ql, qr int) data {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return t[node]
	}
	m := (l + r) >> 1
	if qr <= m { // [ql, qr] 与右子树无交集，仅需递归左子树
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 与左子树无交集，仅需递归右子树
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	// [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
	return t.mergeData(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func countOfPeaks(nums []int, queries [][]int) (ans []int64) {
	n := len(nums)
	t := newSegmentTree(nums)
	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, int64(t.query(1, 0, n-1, q[1], q[2]).cnt))
			continue
		}

		i, v := q[1], q[2]

		if i > 1 {
			oldHas := nums[i-2] < nums[i-1] && nums[i-1] > nums[i]
			newHas := nums[i-2] < nums[i-1] && nums[i-1] > v
			if newHas != oldHas {
				t.update(1, 0, n-1, i-1)
			}
		}

		if 0 < i && i < n-1 {
			oldHas := nums[i-1] < nums[i] && nums[i] > nums[i+1]
			newHas := nums[i-1] < v && v > nums[i+1]
			if newHas != oldHas {
				t.update(1, 0, n-1, i)
			}
		}

		if i < n-2 {
			oldHas := nums[i] < nums[i+1] && nums[i+1] > nums[i+2]
			newHas := v < nums[i+1] && nums[i+1] > nums[i+2]
			if newHas != oldHas {
				t.update(1, 0, n-1, i+1)
			}
		}

		nums[i] = v
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 专题训练

见下面数据结构题单的「**§8.3 线段树**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
