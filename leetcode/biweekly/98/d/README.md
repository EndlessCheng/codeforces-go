下午两点【biIibiIi@灵茶山艾府】直播讲 **Lazy 线段树**，记得关注哦~

---

由于操作 2 和操作 3 更新和统计的是所有 $\textit{nums}_2[i]$ 的值，那么我们其实只需要维护 $\textit{nums}_1$ 中 $1$ 的个数。

用线段树维护区间内 $1$ 的个数 $\textit{cnt}_1$，以及区间反转标记 $\textit{rev}$。

```py [sol1-Python3]
class Solution:
    def handleQuery(self, nums1: List[int], nums2: List[int], queries: List[List[int]]) -> List[int]:
        n = len(nums1)
        cnt1 = [0] * (4 * n)
        rev = [False] * (4 * n)

        def maintain(o: int) -> None:
            cnt1[o] = cnt1[o * 2] + cnt1[o * 2 + 1]

        def do(o: int, l: int, r: int) -> None:
            cnt1[o] = r - l + 1 - cnt1[o]
            rev[o] = not rev[o]

        # 初始化线段树   o,l,r=1,1,n
        def build(o: int, l: int, r: int) -> None:
            if l == r:
                cnt1[o] = nums1[l - 1]
                return
            m = (l + r) // 2
            build(o * 2, l, m)
            build(o * 2 + 1, m + 1, r)
            maintain(o)

        # 反转区间 [L,R]   o,l,r=1,1,n
        def update(o: int, l: int, r: int, L: int, R: int) -> None:
            if L <= l and r <= R:
                do(o, l, r)
                return
            m = (l + r) // 2
            if rev[o]:
                do(o * 2, l, m)
                do(o * 2 + 1, m + 1, r)
                rev[o] = False
            if m >= L: update(o * 2, l, m, L, R)
            if m < R: update(o * 2 + 1, m + 1, r, L, R)
            maintain(o)

        build(1, 1, n)
        ans, s = [], sum(nums2)
        for op, l, r in queries:
            if op == 1: update(1, 1, n, l + 1, r + 1)
            elif op == 2: s += l * cnt1[1]
            else: ans.append(s)
        return ans
```

```java [sol1-Java]
class Solution {
    public long[] handleQuery(int[] nums1, int[] nums2, int[][] queries) {
        int n = nums1.length, m = 0, i = 0;
        cnt1 = new int[n * 4];
        rev = new boolean[n * 4];
        build(nums1, 1, 1, n);

        var sum = 0L;
        for (var x : nums2)
            sum += x;

        for (var q : queries)
            if (q[0] == 3) ++m;
        var ans = new long[m];
        for (var q : queries) {
            if (q[0] == 1) update(1, 1, n, q[1] + 1, q[2] + 1);
            else if (q[0] == 2) sum += (long) q[1] * cnt1[1];
            else ans[i++] = sum;
        }
        return ans;
    }

    private int[] cnt1;
    private boolean[] rev;

    private void maintain(int o) {
        cnt1[o] = cnt1[o * 2] + cnt1[o * 2 + 1];
    }

    private void do_(int o, int l, int r) {
        cnt1[o] = r - l + 1 - cnt1[o];
        rev[o] = !rev[o];
    }

    // 初始化线段树   o,l,r=1,1,n
    private void build(int[] a, int o, int l, int r) {
        if (l == r) {
            cnt1[o] = a[l - 1];
            return;
        }
        int m = (l + r) / 2;
        build(a, o * 2, l, m);
        build(a, o * 2 + 1, m + 1, r);
        maintain(o);
    }

    // 反转区间 [L,R]   o,l,r=1,1,n
    private void update(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) {
            do_(o, l, r);
            return;
        }
        int m = (l + r) / 2;
        if (rev[o]) {
            do_(o * 2, l, m);
            do_(o * 2 + 1, m + 1, r);
            rev[o] = false;
        }
        if (m >= L) update(o * 2, l, m, L, R);
        if (m < R) update(o * 2 + 1, m + 1, r, L, R);
        maintain(o);
    }
}
```

```cpp [sol1-C++]
class Solution {
    static constexpr int MX = 4e5 + 1;

    int cnt1[MX];
    bool rev[MX];

    void maintain(int o) {
        cnt1[o] = cnt1[o * 2] + cnt1[o * 2 + 1];
    }

    void do_(int o, int l, int r) {
        cnt1[o] = r - l + 1 - cnt1[o];
        rev[o] = !rev[o];
    }

    // 初始化线段树   o,l,r=1,1,n
    void build(vector<int> &a, int o, int l, int r) {
        if (l == r) {
            cnt1[o] = a[l - 1];
            return;
        }
        int m = (l + r) / 2;
        build(a, o * 2, l, m);
        build(a, o * 2 + 1, m + 1, r);
        maintain(o);
    }

    // 反转区间 [L,R]   o,l,r=1,1,n
    void update(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) {
            do_(o, l, r);
            return;
        }
        int m = (l + r) / 2;
        if (rev[o]) {
            do_(o * 2, l, m);
            do_(o * 2 + 1, m + 1, r);
            rev[o] = false;
        }
        if (m >= L) update(o * 2, l, m, L, R);
        if (m < R) update(o * 2 + 1, m + 1, r, L, R);
        maintain(o);
    }

public:
    vector<long long> handleQuery(vector<int> &nums1, vector<int> &nums2, vector<vector<int>> &queries) {
        int n = nums1.size();
        build(nums1, 1, 1, n);
        vector<long long> ans;
        long long sum = accumulate(nums2.begin(), nums2.end(), 0LL);
        for (auto &q : queries) {
            if (q[0] == 1) update(1, 1, n, q[1] + 1, q[2] + 1);
            else if (q[0] == 2) sum += 1LL * q[1] * cnt1[1];
            else ans.push_back(sum);
        }
        return ans;
    }
};
```

```go [sol1-Go]
type seg []struct {
	l, r, cnt1 int
	rev        bool
}

func (t seg) maintain(o int) { t[o].cnt1 = t[o<<1].cnt1 + t[o<<1|1].cnt1 }

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].cnt1 = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) do(O int) {
	o := &t[O]
	o.cnt1 = o.r - o.l + 1 - o.cnt1
	o.rev = !o.rev
}

func (t seg) spread(o int) {
	if t[o].rev {
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].rev = false
	}
}

func (t seg) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t.maintain(o)
}

func handleQuery(nums1, nums2 []int, queries [][]int) (ans []int64) {
	sum := 0
	for _, x := range nums2 {
		sum += x
	}
	t := make(seg, len(nums1)*4)
	t.build(nums1, 1, 1, len(nums1))
	for _, q := range queries {
		if q[0] == 1 {
			t.update(1, q[1]+1, q[2]+1)
		} else if q[0] == 2 {
			sum += q[1] * t[1].cnt1
		} else {
			ans = append(ans, int64(sum))
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n+q\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(n)$。
