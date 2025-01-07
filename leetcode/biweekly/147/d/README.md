## 分析

如果不删数字，那么做法同 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)。

如果删数字，那么只考虑子数组包含被删数字的情况（否则就和不删数字一样了），且被删数字必须是负数（删 $0$ 或者正数不如不删）

## 方法一：线段树

设 $\textit{nums}$ 的最大值为 $m$。

如果 $m\le 0$，那么答案就是 $m$，因为负数只能选一个最大的。

否则答案一定是正数，此时可以把「删除 $x$」视作「把 $x$ 改成 $0$」。

这可以用单点修改的线段树维护，原理见 53 题官方题解的方法二（分治）。题目是 [P4513 小白逛公园](https://www.luogu.com.cn/problem/P4513)。

为了快速知道哪些数要改成 $0$，用哈希表记录每个元素的所有出现位置。

```py [sol-Python3]
max = lambda a, b: b if b > a else a  # 手动比大小，效率更高

class SegmentTree:
    def __init__(self, nums: List[int]):
        n = len(nums)
        self.tree = [(0, 0, 0, 0)] * (2 << n.bit_length())
        self.build(nums, 1, 0, n - 1)

    def set(self, o: int, val: int) -> None:
        self.tree[o] = (val, val, val, val)

    def merge_info(self, a: Tuple[int, int, int, int], b: Tuple[int, int, int, int]) -> Tuple[int, int, int, int]:
        ans = max(max(a[0], b[0]), a[3] + b[2])
        s = a[1] + b[1]
        pre = max(a[2], a[1] + b[2])
        suf = max(b[3], b[1] + a[3])
        return ans, s, pre, suf

    def maintain(self, o: int) -> None:
        self.tree[o] = self.merge_info(self.tree[o * 2], self.tree[o * 2 + 1])

    # 初始化线段树
    def build(self, nums, o: int, l: int, r: int) -> None:
        if l == r:
            self.set(o, nums[l])
            return
        m = (l + r) // 2
        self.build(nums, o * 2, l, m)
        self.build(nums, o * 2 + 1, m + 1, r)
        self.maintain(o)

    # 单点更新
    def update(self, o: int, l: int, r: int, i: int, val: int) -> None:
        if l == r:
            self.set(o, val)
            return
        m = (l + r) // 2
        if i <= m:
            self.update(o * 2, l, m, i, val)
        else:
            self.update(o * 2 + 1, m + 1, r, i, val)
        self.maintain(o)

    # 区间询问（没用到）
    def query(self, o: int, l: int, r: int, L: int, R: int) -> Tuple[int, int, int, int]:
        if L <= l and r <= R:
            return self.tree[o]
        m = (l + r) // 2
        if R <= m:
            return self.query(o * 2, l, m, L, R)
        if m < L:
            return self.query(o * 2 + 1, m + 1, r, L, R)
        return self.merge_info(
            self.query(o * 2, l, m, L, R),
            self.query(o * 2 + 1, m + 1, r, L, R)
        )

class Solution:
    def maxSubarraySum(self, nums: List[int]) -> int:
        n = len(nums)
        t = SegmentTree(nums)
        ans = t.tree[1][0]  # 不删任何数
        if ans <= 0:
            return ans

        pos = defaultdict(list)
        for i, x in enumerate(nums):
            if x < 0:
                pos[x].append(i)
        for idx in pos.values():
            for i in idx:
                t.update(1, 0, n - 1, i, 0)  # 删除
            ans = max(ans, t.tree[1][0])
            for i in idx:
                t.update(1, 0, n - 1, i, nums[i])  # 复原
        return ans
```

```java [sol-Java]
class SegmentTree {
    private record Info(long ans, long sum, long pre, long suf) {}

    private final Info[] tree;

    public SegmentTree(int[] nums) {
        int n = nums.length;
        tree = new Info[2 << (32 - Integer.numberOfLeadingZeros(n))];
        build(nums, 1, 0, n - 1);
    }

    private Info mergeInfo(Info a, Info b) {
        return new Info(
            Math.max(Math.max(a.ans, b.ans), a.suf + b.pre()),
            a.sum + b.sum,
            Math.max(a.pre, a.sum + b.pre),
            Math.max(b.suf, b.sum + a.suf)
        );
    }

    private void set(int o, int val) {
        tree[o] = new Info(val, val, val, val);
    }

    private void maintain(int o) {
        tree[o] = mergeInfo(tree[o << 1], tree[o << 1 | 1]);
    }

    // 初始化线段树
    private void build(int[] nums, int o, int l, int r) {
        if (l == r) {
            set(o, nums[l]);
            return;
        }
        int m = (l + r) >> 1;
        build(nums, o << 1, l, m);
        build(nums, o << 1 | 1, m + 1, r);
        maintain(o);
    }

    // 单点更新
    public void update(int o, int l, int r, int i, int val) {
        if (l == r) {
            set(o, val);
            return;
        }
        int m = (l + r) >> 1;
        if (i <= m) {
            update(o << 1, l, m, i, val);
        } else {
            update(o << 1 | 1, m + 1, r, i, val);
        }
        maintain(o);
    }

    public long queryAll() {
        return tree[1].ans;
    }

    // 区间询问（没用到）
    public Info query(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) {
            return tree[o];
        }
        int m = (l + r) >> 1;
        if (R <= m) {
            return query(o << 1, l, m, L, R);
        }
        if (m < L) {
            return query(o << 1 | 1, m + 1, r, L, R);
        }
        return mergeInfo(
                query(o << 1, l, m, L, R),
                query(o << 1 | 1, m + 1, r, L, R)
        );
    }
}

class Solution {
    public long maxSubarraySum(int[] nums) {
        int n = nums.length;
        SegmentTree t = new SegmentTree(nums);
        long ans = t.queryAll(); // 不删任何数
        if (ans <= 0) {
            return ans;
        }

        Map<Integer, List<Integer>> pos = new HashMap<>();
        for (int i = 0; i < n; i++) {
            if (nums[i] < 0) {
                pos.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
            }
        }
        for (List<Integer> idx : pos.values()) {
            for (int i : idx) {
                t.update(1, 0, n - 1, i, 0); // 删除
            }
            ans = Math.max(ans, t.queryAll());
            for (int i : idx) {
                t.update(1, 0, n - 1, i, nums[i]); // 复原
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
struct SegmentTree {
    struct Info {
        long long ans, sum, pre, suf;
    };

    vector<Info> tree;

    SegmentTree(vector<int>& nums) {
        int n = nums.size();
        tree.resize(2 << (bit_width((unsigned) n) + 1));
        build(nums, 1, 0, n - 1);
    }

    Info mergeInfo(Info& a, Info& b) {
        return {
            max({a.ans, b.ans, a.suf + b.pre}),
            a.sum + b.sum,
            max(a.pre, a.sum + b.pre),
            max(b.suf, b.sum + a.suf),
        };
    }

    void set(int o, int val) {
        tree[o] = {val, val, val, val};
    }

    void maintain(int o) {
        tree[o] = mergeInfo(tree[o * 2], tree[o * 2 + 1]);
    }

    // 初始化线段树
    void build(vector<int>& nums, int o, int l, int r) {
        if (l == r) {
            set(o, nums[l]);
            return;
        }
        int m = (l + r) / 2;
        build(nums, o * 2, l, m);
        build(nums, o * 2 + 1, m + 1, r);
        maintain(o);
    }

    // 单点更新
    void update(int o, int l, int r, int i, int val) {
        if (l == r) {
            set(o, val);
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) {
            update(o * 2, l, m, i, val);
        } else {
            update(o * 2 + 1, m + 1, r, i, val);
        }
        maintain(o);
    }
};

class Solution {
public:
    long long maxSubarraySum(vector<int>& nums) {
        int n = nums.size();
        SegmentTree t(nums);
        long long ans = t.tree[1].ans; // 不删任何数
        if (ans <= 0) {
            return ans;
        }

        unordered_map<int, vector<int>> pos;
        for (int i = 0; i < n; i++) {
            if (nums[i] < 0) {
                pos[nums[i]].push_back(i);
            }
        }
        for (auto& [_, idx] : pos) {
            for (int i : idx) {
                t.update(1, 0, n - 1, i, 0); // 删除
            }
            ans = max(ans, t.tree[1].ans);
            for (int i : idx) {
                t.update(1, 0, n - 1, i, nums[i]); // 复原
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type info struct {
	ans, sum, pre, suf int
}

type seg []info

func (t seg) set(o, val int) {
	t[o] = info{val, val, val, val}
}

func (t seg) mergeInfo(a, b info) info {
	return info{
		max(max(a.ans, b.ans), a.suf+b.pre),
		a.sum + b.sum,
		max(a.pre, a.sum+b.pre),
		max(b.suf, b.sum+a.suf),
	}
}

func (t seg) maintain(o int) {
	t[o] = t.mergeInfo(t[o<<1], t[o<<1|1])
}

// 初始化线段树
func (t seg) build(nums []int, o, l, r int) {
	if l == r {
		t.set(o, nums[l])
		return
	}
	m := (l + r) >> 1
	t.build(nums, o<<1, l, m)
	t.build(nums, o<<1|1, m+1, r)
	t.maintain(o)
}

// 单点更新
func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t.set(o, val)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t.maintain(o)
}

// 区间询问（没用到）
func (t seg) query(o, l, r, L, R int) info {
	if L <= l && r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, L, R)
	}
	if m < L {
		return t.query(o<<1|1, m+1, r, L, R)
	}
	return t.mergeInfo(t.query(o<<1, l, m, L, R), t.query(o<<1|1, m+1, r, L, R))
}

func maxSubarraySum(nums []int) int64 {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	ans := t[1].ans // 不删任何数
	if ans <= 0 {
		return int64(ans)
	}

	pos := map[int][]int{}
	for i, x := range nums {
		if x < 0 {
			pos[x] = append(pos[x], i)
		}
	}
	for _, idx := range pos {
		for _, i := range idx {
			t.update(1, 0, n-1, i, 0) // 删除
		}
		ans = max(ans, t[1].ans)
		for _, i := range idx {
			t.update(1, 0, n-1, i, nums[i]) // 复原
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。每个下标 $i$ 都会恰好调用两次线段树的更新，每次更新是 $\mathcal{O}(\log n)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：前后缀分解

假设删除的是 $x=\textit{nums}[i]$，考虑前后缀分解，问题变成：

- 删除所有 $x$ 后，以 $i$ 结尾的最大子数组和 $\textit{pre}[i]$。
- 删除所有 $x$ 后，以 $i$ 开头的最大子数组和 $\textit{suf}[i]$。

讨论 $\textit{pre}[i]$ 怎么算：

- 设 $i$ 左边最近 $x$ 的下标为 $j$。
- 情况一：如果子数组不包含 $\textit{nums}[j]$，那么问题变成在不删除元素的情况下，以 $i-1$ 结尾的最大子数组和 $f[i-1]$（见下文的「注意」）。
- 情况二：如果子数组包含 $\textit{nums}[j]$，那么问题变成在删除 $x$ 的情况下，以 $j$ 结尾的最大子数组和 $\textit{pre}[j]$，加上子数组 $[j+1,i-1]$ 的元素和，即 $\textit{pre}[i] = \textit{pre}[j] + \sum\limits_{k=j+1}^{i-1} \textit{nums}[k]$。用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化得 $\textit{pre}[i] = \textit{pre}[j] + s[i] - s[j+1]$。

二者取最大值，得

$$
\textit{pre}[i] = \max(f[i-1], \textit{pre}[j] + s[i] - s[j+1])
$$

⚠**注意**：由于被删除的数是负数，那么如果 $f[i-1]$ 包含 $\textit{nums}[j]$，必然不如情况二优。所以当 $f[i-1]$ 比情况二大时，$f[i-1]$ **必然不包含** $\textit{nums}[j]$。实现时，无需判断被删除的数是否为负数，因为删除正数的情况必然不如删除负数优。

代码实现时，与其记录上一个 $x$ 的下标，不如直接把 $\textit{pre}[j] - s[j+1]$ 记在哈希表 $\textit{last}[x]$ 中，这样转移方程为

$$
\textit{pre}[i] = \max(f[i-1], \textit{last}[x] + s[i])
$$

对于 $\textit{suf}[i]$ 的计算同理。

最终答案为如下情况的最大值：

- 不删除元素，即 $f[i]$ 的最大值。
- 删除 $\textit{nums}[i]$，即 $\textit{pre}[i] + \textit{suf}[i]$。
- 如果 $\textit{suf}[i] < 0$，那么只取 $\textit{pre}[i]$。
- 如果 $\textit{pre}[i] < 0$，那么只取 $\textit{suf}[i]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1SzrAYMESJ/?t=23m43s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxSubarraySum(self, nums: List[int]) -> int:
        n = len(nums)
        f = -inf
        s = 0
        last = {}

        def update(x: int) -> int:
            nonlocal f, s
            res = f  # f[i-1]
            f = max(f, 0) + x  # f[i] = max(f[i-1], 0) + x
            if x in last:
                res = max(res, last[x] + s)  # s[i]
            s += x  # s[i+1] = s[i] + x
            last[x] = res - s
            return res

        suf = [0] * n
        for i in range(n - 1, -1, -1):
            suf[i] = update(nums[i])

        ans = f = -inf
        s = 0
        last = {}
        for x, sf in zip(nums, suf):
            pre = update(x)
            ans = max(ans, f, pre + sf, pre, sf)
        return ans
```

```java [sol-Java]
class Solution {
    private long f = Long.MIN_VALUE / 2; // 当前的最大子数组和
    private long s = 0; // 前缀和
    private final Map<Integer, Long> last = new HashMap<>(); // 记录每个元素最后一次更新的状态

    public long maxSubarraySum(int[] nums) {
        int n = nums.length;
        long[] pre = new long[n];
        for (int i = 0; i < n; i++) {
            pre[i] = update(nums[i]);
        }

        long ans = Long.MIN_VALUE;
        f = Long.MIN_VALUE / 2;
        s = 0;
        last.clear();
        for (int i = n - 1; i >= 0; i--) {
            long suf = update(nums[i]);
            ans = Math.max(ans, Math.max(f, Math.max(pre[i] + suf, Math.max(pre[i], suf))));
        }
        return ans;
    }

    private long update(int x) {
        long res = f; // f[i-1]
        f = Math.max(f, 0) + x; // f[i] = max(f[i-1], 0) + x
        if (last.containsKey(x)) {
            res = Math.max(res, last.get(x) + s); // s[i]
        }
        s += x; // s[i+1] = s[i] + x
        last.put(x, res - s);
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarraySum(vector<int>& nums) {
        int n = nums.size();
        long long f = LLONG_MIN / 2;
        long long s = 0;
        unordered_map<int, long long> last;

        auto update = [&](int x) -> long long {
            long long res = f; // f[i-1]
            f = max(f, 0LL) + x; // f[i] = max(f[i-1], 0) + x
            auto it = last.find(x);
            if (it != last.end()) {
                res = max(res, it->second + s); // s[i]
            }
            s += x; // s[i+1] = s[i] + x
            last[x] = res - s;
            return res;
        };

        vector<long long> pre(n);
        for (int i = 0; i < n; i++) {
            pre[i] = update(nums[i]);
        }

        long long ans = LLONG_MIN;
        f = LLONG_MIN / 2;
        s = 0;
        last.clear();
        for (int i = n - 1; i >= 0; i--) {
            long long suf = update(nums[i]);
            ans = max({ans, f, pre[i] + suf, pre[i], suf});
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSubarraySum(nums []int) int64 {
	n := len(nums)
	f := math.MinInt / 2
	s := 0
	last := map[int]int{}

	update := func(x int) int {
		res := f // f[i-1]
		f = max(f, 0) + x // f[i] = max(f[i-1], 0) + x
		if v, ok := last[x]; ok {
			res = max(res, v+s) // s[i]
		}
		s += x // s[i+1] = s[i] + x
		last[x] = res - s
		return res
	}

	pre := make([]int, n)
	for i, x := range nums {
		pre[i] = update(x)
	}

	ans := math.MinInt
	f = math.MinInt / 2
	s = 0
	clear(last)
	for i, x := range slices.Backward(nums) {
		suf := update(x)
		ans = max(ans, f, pre[i]+suf, pre[i], suf)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法三：前缀和+枚举右维护左

前置知识：[前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

计算 $\textit{nums}$ 的前缀和数组 $s$。

设被删的数是 $x$，答案对应的子数组的下标区间为 $[i,j)$，其中包含 $k$ 个 $x$，那么答案为

$$
(s[j] - s[i]) - k\cdot x = s[j] - (s[i] + k\cdot x)
$$

枚举 $j$，那么 $s[i] + k\cdot x$ 越小，答案越大。

所以我们需要在枚举 $j$ 的过程中，维护 $s[i] + k\cdot x$ 的最小值，记作 $\textit{delMinS}[x]$。所有 $\textit{delMinS}[x]$ 的最小值记作 $\textit{allMin}$。用 $s[j]-\textit{allMin}$ 更新答案的最大值。

如何维护 $s[i] + k\cdot x$ 的最小值？

设 $x$ 是当前遍历到的元素，我们要删除它。如果 $x\ge 0$，那么删除没有意义，所以下面只讨论 $x<0$ 的情况：

- 如果 $x$ 是首次遇到（左边没有 $x$），那么 $k=1$，我们需要知道在 $x$ 左边的所有 $s[i]$ 的最小值 $\textit{nonDelMinS}$，这样才能使 $s[i] + k\cdot x = \textit{nonDelMinS} + x$ 最小。把 $\textit{nonDelMinS} + x$ 记录到哈希表 $\textit{delMinS}[x]$ 中。
- 如果 $x$ 是第二次遇到（左边还有一个 $x$）：
  - 如果打算包含两个 $x$，那么需要知道在上一个 $x$ 的位置，$\textit{nonDelMinS} + x$ 是多少，也就是 $\textit{delMinS}[x]$ 中存储的值。于是 $s[i] + 2\cdot x = (s[i] + x) + x$ = $\textit{delMinS}[x] + x$。
  - 如果只打算包含一个 $x$，那么同上，$s[i] + x = \textit{nonDelMinS} + x$。
  - 二者取最小值，得 $\textit{delMinS}[x] = \min(\textit{delMinS}[x], \textit{nonDelMinS}) + x$。
  - 该递推式也适用于有更多个 $x$ 的情况。

代码实现时，前缀和可以只用一个变量 $s$ 表示。其初始值为 $0$，对应着 $s[0]=0$。

### 答疑

**问**：如果从 $\textit{nonDelMinS}$ 到当前 $x$ 之间，有多个 $x$ 呢？这样不就包含多个 $x$ 了吗？

**答**：根据 $\textit{delMinS}[x]$ 的计算过程，如果从 $\textit{nonDelMinS}$ 到当前 $x$ 之间有多个 $x$，那么在此之前，我们已经把 $\textit{nonDelMinS} + x$ 更新到 $\textit{delMinS}[x]$ 中了，由于 $x<0$，所以更新之后的 $\textit{delMinS}[x] < \textit{nonDelMinS}$ 一定成立，所以递推式中的 $\min$ 一定会取 $\textit{delMinS}[x]$。换句话说，如果 $\min$ 取的是 $\textit{nonDelMinS}$，那么对应的子数组一定只包含一个 $x$。

**问**：如何保证子数组一定是非空的？

**答**：先计算 $s-\textit{allMin}$，再更新 $\textit{allMin}$，以保证子数组是非空的。

**问**：对于 $\textit{nums}=[-1,-1,-1]$，如何理解代码中的 $s-\textit{allMin}$？

**答**：此时 $s-\textit{allMin} = -1$，对应着不删除任何元素（即 $k=0$）的情况。

```py [sol-Python3]
class Solution:
    def maxSubarraySum(self, nums: List[int]) -> int:
        ans = -inf
        s = non_del_min_s = all_min = 0
        del_min_s = defaultdict(int)
        for x in nums:
            s += x
            ans = max(ans, s - all_min)
            if x < 0:
                del_min_s[x] = min(del_min_s[x], non_del_min_s) + x
                all_min = min(all_min, del_min_s[x])
                non_del_min_s = min(non_del_min_s, s)
        return ans
```

```py [sol-Python3 手动比大小]
class Solution:
    def maxSubarraySum(self, nums: List[int]) -> int:
        ans = -inf
        s = non_del_min_s = all_min = 0
        del_min_s = defaultdict(int)
        for x in nums:
            s += x
            if s - all_min > ans: ans = s - all_min
            if x < 0:
                if non_del_min_s < del_min_s[x]: del_min_s[x] = non_del_min_s
                del_min_s[x] += x
                if del_min_s[x] < all_min: all_min = del_min_s[x]
                if s < non_del_min_s: non_del_min_s = s
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSubarraySum(int[] nums) {
        long ans = Long.MIN_VALUE;
        long s = 0; // 前缀和
        long nonDelMinS = 0; // 最小前缀和
        Map<Integer, Long> delMinS = new HashMap<>(); // x -> 最小前缀和 + 若干 x
        long allMin = 0; // 所有 delMinS[x] 的最小值
        for (int x : nums) {
            s += x;
            ans = Math.max(ans, s - allMin);
            if (x < 0) {
                long res = Math.min(delMinS.getOrDefault(x, Long.MAX_VALUE), nonDelMinS) + x;
                delMinS.put(x, res);
                allMin = Math.min(allMin, res);
                nonDelMinS = Math.min(nonDelMinS, s);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarraySum(vector<int>& nums) {
        long long ans = INT_MIN;
        long long s = 0, non_del_min_s = 0, all_min = 0;
        unordered_map<int, long long> del_min_s;
        for (int x : nums) {
            s += x;
            ans = max(ans, s - all_min);
            if (x < 0) {
                del_min_s[x] = min(del_min_s[x], non_del_min_s) + x;
                all_min = min(all_min, del_min_s[x]);
                non_del_min_s = min(non_del_min_s, s);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSubarraySum(nums []int) int64 {
	ans := math.MinInt
	var s, nonDelMinS, allMin int
	delMinS := map[int]int{}
	for _, x := range nums {
		s += x
		ans = max(ans, s-allMin)
		if x < 0 {
			delMinS[x] = min(delMinS[x], nonDelMinS) + x
			allMin = min(allMin, delMinS[x])
			nonDelMinS = min(nonDelMinS, s)
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单中的「**§11.4 树状数组/线段树优化 DP**」，数据结构题单中的「**§1.2 前缀和与哈希表**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
