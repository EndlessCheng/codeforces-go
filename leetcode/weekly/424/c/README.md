## 方法一：二分答案+差分数组

请先完成上一题 [3355. 零数组变换 I](https://leetcode.cn/problems/zero-array-transformation-i/)。

本题由于 $k$ 越大，越能满足要求；$k$ 越小，越无法满足要求。有**单调性**，可以二分答案求最小的 $k$。

问题变成：

- 能否用前 $k$ 个询问（下标从 $0$ 到 $k-1$）把 $\textit{nums}$ 的所有元素都变成 $\le 0$？

用上一题的差分数组计算。

### 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$-1$。一定无法满足要求。
- 开区间右端点初始值：$q+1$，其中 $q$ 为 $\textit{queries}$ 的长度。假定 $q+1$ 一定可以满足要求，如果二分结果等于 $q+1$，那么返回 $-1$。

对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1yiU6YnEfU/?t=16m57s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        # 3355. 零数组变换 I
        def check(k: int) -> bool:
            diff = [0] * (len(nums) + 1)
            for l, r, val in queries[:k]:  # 前 k 个询问
                diff[l] += val
                diff[r + 1] -= val

            for x, sum_d in zip(nums, accumulate(diff)):
                if x > sum_d:
                    return False
            return True

        q = len(queries)
        left, right = -1, q + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                right = mid
            else:
                left = mid
        return right if right <= q else -1
```

```py [sol-Python3 库函数]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        # 3355. 零数组变换 I
        def check(k: int) -> bool:
            diff = [0] * (len(nums) + 1)
            for l, r, val in queries[:k]:  # 前 k 个询问
                diff[l] += val
                diff[r + 1] -= val

            for x, sum_d in zip(nums, accumulate(diff)):
                if x > sum_d:
                    return False
            return True

        q = len(queries)
        ans = bisect_left(range(q + 1), True, key=check)
        return ans if ans <= q else -1
```

```java [sol-Java]
class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        int q = queries.length;
        int left = -1, right = q + 1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(mid, nums, queries)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right <= q ? right : -1;
    }

    // 3355. 零数组变换 I
    private boolean check(int k, int[] nums, int[][] queries) {
        int n = nums.length;
        int[] diff = new int[n + 1];
        for (int i = 0; i < k; i++) { // 前 k 个询问
            int[] q = queries[i];
            int l = q[0], r = q[1], val = q[2];
            diff[l] += val;
            diff[r + 1] -= val;
        }

        int sumD = 0;
        for (int i = 0; i < n; i++) {
            sumD += diff[i];
            if (nums[i] > sumD) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        // 3355. 零数组变换 I
        int n = nums.size();
        vector<int> diff(n + 1);
        auto check = [&](int k) -> bool {
            ranges::fill(diff, 0);
            for (int i = 0; i < k; i++) { // 前 k 个询问
                auto& q = queries[i];
                int l = q[0], r = q[1], val = q[2];
                diff[l] += val;
                diff[r + 1] -= val;
            }

            int sum_d = 0;
            for (int i = 0; i < n; i++) {
                sum_d += diff[i];
                if (nums[i] > sum_d) {
                    return false;
                }
            }
            return true;
        };

        int q = queries.size();
        int left = -1, right = q + 1;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right <= q ? right : -1;
    }
};
```

```go [sol-Go]
func minZeroArray(nums []int, queries [][]int) int {
	q := len(queries)
	diff := make([]int, len(nums)+1)
	ans := sort.Search(q+1, func(k int) bool {
		// 3355. 零数组变换 I
		clear(diff)
		for _, q := range queries[:k] { // 前 k 个询问
			l, r, val := q[0], q[1], q[2]
			diff[l] += val
			diff[r+1] -= val
		}

		sumD := 0
		for i, x := range nums {
			sumD += diff[i]
			if x > sumD {
				return false
			}
		}
		return true
	})
	if ans > q {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。Python 忽略切片空间。

## 方法二：Lazy 线段树

直接用 Lazy 线段树模拟区间减法。

线段树维护区间最大值。

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        n = len(nums)
        m = 2 << n.bit_length()
        mx = [0] * m
        todo = [0] * m

        def do(o: int, v: int) -> None:
            mx[o] -= v
            todo[o] += v

        def spread(o: int) -> None:
            if todo[o] != 0:
                do(o * 2, todo[o])
                do(o * 2 + 1, todo[o])
                todo[o] = 0

        def maintain(o: int) -> None:
            mx[o] = max(mx[o * 2], mx[o * 2 + 1])

        def build(o: int, l: int, r: int) -> None:
            if l == r:
                mx[o] = nums[l]
                return
            m = (l + r) // 2
            build(o * 2, l, m)
            build(o * 2 + 1, m + 1, r)
            maintain(o)

        def update(o: int, l: int, r: int, ql: int, qr: int, v: int) -> None:
            if ql <= l and r <= qr:
                do(o, v)
                return
            spread(o)
            m = (l + r) // 2
            if ql <= m:
                update(o * 2, l, m, ql, qr, v)
            if m < qr:
                update(o * 2 + 1, m + 1, r, ql, qr, v)
            maintain(o)

        build(1, 0, n - 1)
        if mx[1] <= 0:
            return 0

        for i, (ql, qr, v) in enumerate(queries):
            update(1, 0, n - 1, ql, qr, v)
            if mx[1] <= 0:
                return i + 1
        return -1
```

```java [sol-Java]
class SegmentTree {
    private final int[] mx;
    private final int[] todo;

    public SegmentTree(int[] nums) {
        int n = nums.length;
        int m = 2 << (32 - Integer.numberOfLeadingZeros(n));
        mx = new int[m];
        todo = new int[m];
        build(1, 0, n - 1, nums);
    }

    private void do_(int o, int v) {
        mx[o] -= v;
        todo[o] += v;
    }

    private void spread(int o) {
        if (todo[o] != 0) {
            do_(o * 2, todo[o]);
            do_(o * 2 + 1, todo[o]);
            todo[o] = 0;
        }
    }

    private void maintain(int o) {
        mx[o] = Math.max(mx[o * 2], mx[o * 2 + 1]);
    }

    private void build(int o, int l, int r, int[] nums) {
        if (l == r) {
            mx[o] = nums[l];
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m, nums);
        build(o * 2 + 1, m + 1, r, nums);
        maintain(o);
    }

    public void update(int o, int l, int r, int ql, int qr, int v) {
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

    public int queryAll() {
        return mx[1];
    }
}

class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        SegmentTree tree = new SegmentTree(nums);
        if (tree.queryAll() <= 0) {
            return 0;
        }
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            tree.update(1, 0, nums.length - 1, q[0], q[1], q[2]);
            if (tree.queryAll() <= 0) {
                return i + 1;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class SegmentTree {
    int n;
    vector<int> mx;
    vector<int> todo;

    void do_(int o, int v) {
        mx[o] -= v;
        todo[o] += v;
    }

    void spread(int o) {
        if (todo[o]) {
            do_(o * 2, todo[o]);
            do_(o * 2 + 1, todo[o]);
            todo[o] = 0;
        }
    }

    void maintain(int o) {
        mx[o] = max(mx[o * 2], mx[o * 2 + 1]);
    }

    void build(int o, int l, int r, vector<int>& nums) {
        if (l == r) {
            mx[o] = nums[l];
            return;
        }
        int m = (l + r) / 2;
        build(o * 2, l, m, nums);
        build(o * 2 + 1, m + 1, r, nums);
        maintain(o);
    }

    void update(int o, int l, int r, int ql, int qr, int v) {
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

public:
    SegmentTree(vector<int>& nums) {
        n = nums.size();
        int m = 2 << (32 - __builtin_clz(n));
        mx.resize(m);
        todo.resize(m);
        build(1, 0, n - 1, nums);
    }

    void update(int ql, int qr, int v) {
        update(1, 0, n - 1, ql, qr, v);
    }

    int query_all() {
        return mx[1];
    }
};

class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        SegmentTree tree(nums);
        if (tree.query_all() <= 0) {
            return 0;
        }
        for (int i = 0; i < queries.size(); ++i) {
            auto& q = queries[i];
            tree.update(q[0], q[1], q[2]);
            if (tree.query_all() <= 0) {
                return i + 1;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
type seg []struct {
	l, r, mx, todo int
}

func (t seg) do(o, v int) {
	t[o].mx -= v
	t[o].todo += v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg) maintain(o int) {
	t[o].mx = max(t[o<<1].mx, t[o<<1|1].mx)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mx = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

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

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	if t[1].mx <= 0 {
		return 0
	}
	for i, q := range queries {
		t.update(1, q[0], q[1], q[2])
		if t[1].mx <= 0 {
			return i + 1
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法三：双指针+差分数组

和方法一一样，用一个差分数组处理询问。

这次我们从左到右遍历 $x=\textit{nums}[i]$，如果发现 $x>\textit{sumD}$，那么就必须处理询问，直到 $x\le \textit{sumD}$ 为止。

对于询问 $[l,r,\textit{val}]$，如果发现 $l\le i \le r$，那么直接把 $\textit{sumD}$ 增加 $\textit{val}$。

由于处理过的询问无需再处理，所以上述过程可以用双指针实现。

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        diff = [0] * (len(nums) + 1)
        sum_d = k = 0
        for i, (x, d) in enumerate(zip(nums, diff)):
            sum_d += d
            while k < len(queries) and sum_d < x:  # 需要添加询问，把 x 减小
                l, r, val = queries[k]
                diff[l] += val
                diff[r + 1] -= val
                if l <= i <= r:  # x 在更新范围中
                    sum_d += val
                k += 1
            if sum_d < x:  # 无法更新
                return -1
        return k
```

```java [sol-Java]
class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        int n = nums.length;
        int[] diff = new int[n + 1];
        int sumD = 0;
        int k = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            sumD += diff[i];
            while (k < queries.length && sumD < x) { // 需要添加询问，把 x 减小
                int[] q = queries[k];
                int l = q[0], r = q[1], val = q[2];
                diff[l] += val;
                diff[r + 1] -= val;
                if (l <= i && i <= r) { // x 在更新范围中
                    sumD += val;
                }
                k++;
            }
            if (sumD < x) { // 无法更新
                return -1;
            }
        }
        return k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<int> diff(n + 1);
        int sum_d = 0, k = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            sum_d += diff[i];
            while (k < queries.size() && sum_d < x) { // 需要添加询问，把 x 减小
                auto& q = queries[k];
                int l = q[0], r = q[1], val = q[2];
                diff[l] += val;
                diff[r + 1] -= val;
                if (l <= i && i <= r) { // x 在更新范围中
                    sum_d += val;
                }
                k++;
            }
            if (sum_d < x) { // 无法更新
                return -1;
            }
        }
        return k;
    }
};
```

```go [sol-Go]
func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	diff := make([]int, n+1)
	sumD, k := 0, 0
	for i, x := range nums {
		sumD += diff[i]
		for k < len(queries) && sumD < x { // 需要添加询问，把 x 减小
			q := queries[k]
			l, r, val := q[0], q[1], q[2]
			diff[l] += val
			diff[r+1] -= val
			if l <= i && i <= r { // x 在更新范围中
				sumD += val
			}
			k++
		}
		if sumD < x { // 无法更新
			return -1
		}
	}
	return k
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果询问可以按照任意顺序执行呢？这里限制 $\textit{val}=1$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. 【本题相关】[滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. 【本题相关】[二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
