通知：今晚 8 点在 B 站直播讲双周赛和周赛的题目（包含线段树的讲解）。

感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

本题有点「超纲」（算法竞赛难度），直接讲做法。

本题的关键在于如何求出前 $\textit{maxRow}$ 排座位中，观众数不超过 $m-k$ 且下标最小的那一排座位。同时，我们还需要动态维护每一排的观众数，这可以用**线段树**来解决。

对于 $\texttt{gather}$ 操作，为了求出目标的那一排座位，我们可以**在线段树上二分**。

具体来说，我们可以用线段树维护每个区间上的观众数的最小值 $\textit{min}$：

- 如果当前区间 $\textit{min}>m-k$，则无法坐 $k$ 个人，返回 $0$；
- 如果当前区间只包含一个元素，则返回该元素的下标；
- 如果左半部分 $\textit{min}\le m-k$，则递归左半部分；
- 否则如果 $\textit{maxRow}+1$ 在右半部分内，则递归右半部分；
- 否则返回 $0$。

由于每次要么递归左半，要么递归右半，因此线段树二分的复杂度为线段树的树高，即 $O(\log n)$。

对于 $\texttt{scatter}$ 操作，为了判断能否坐 $k$ 个人，我们还需要用线段树维护区间的元素和 $\textit{sum}$，从而计算出前 $\textit{maxRow}$ 排座位中的人数。

如果剩余座位不低于 $k$，我们可以**从第一个未坐满的排开始**占座，这同样可以在线段树上二分：找第一个小于 $m$ 的元素下标。

#### 复杂度分析

- 时间复杂度：
   - 初始化为 $O(n)$。初始化线段树需要 $O(n)$。
   - $\texttt{gather}$ 为 $O(\log n)$。
   - $\texttt{scatter}$ 可以从整体上来分析：我们要么在填充空排，要么在填充受之前操作影响的未填满的排，所以**所有 $\texttt{scatter}$ 操作**的复杂度之和为 $O((n+q)\log n)$，这里 $q$ 为 $\texttt{gather}$ 和 $\texttt{scatter}$ 的调用次数之和。注意上述题解中的「从第一个未坐满的排开始占座」保证了总体复杂度不会退化至 $O(nq)$。
- 空间复杂度：$O(n)$。线段树需要 $O(n)$ 的空间。

```Python [sol1-Python3]
class BookMyShow:
    def __init__(self, n: int, m: int):
        self.n = n
        self.m = m
        self.i = 1
        self.min = [0] * (n * 4)
        self.sum = [0] * (n * 4)

    # 将 idx 上的元素值增加 val
    def add(self, o: int, l: int, r: int, idx: int, val: int):
        if l == r:
            self.min[o] += val
            self.sum[o] += val
            return
        m = (l + r) // 2
        if idx <= m: self.add(o * 2, l, m, idx, val)
        else: self.add(o * 2 + 1, m + 1, r, idx, val)
        self.min[o] = min(self.min[o * 2], self.min[o * 2 + 1])
        self.sum[o] = self.sum[o * 2] + self.sum[o * 2 + 1]

    # 返回区间 [L,R] 内的元素和
    def query_sum(self, o: int, l: int, r: int, L: int, R: int):
        if L <= l and r <= R: return self.sum[o]
        sum = 0
        m = (l + r) // 2
        if L <= m: sum += self.query_sum(o * 2, l, m, L, R)
        if R > m: sum += self.query_sum(o * 2 + 1, m + 1, r, L, R)
        return sum

    # 返回区间 [1,R] 中 <= val 的最靠左的位置，不存在时返回 0
    def index(self, o: int, l: int, r: int, R: int, val: int):
        if self.min[o] > val: return 0  # 说明整个区间的元素值都大于 val
        if l == r: return l
        m = (l + r) // 2
        if self.min[o * 2] <= val: return self.index(o * 2, l, m, R, val)  # 看看左半部分
        if m < R: return self.index(o * 2 + 1, m + 1, r, R, val)  # 看看右半部分
        return 0

    def gather(self, k: int, maxRow: int) -> List[int]:
        i = self.index(1, 1, self.n, maxRow + 1, self.m - k)
        if i == 0: return []
        seats = self.query_sum(1, 1, self.n, i, i)
        self.add(1, 1, self.n, i, k)  # 占据 k 个座位
        return [i - 1, seats]

    def scatter(self, k: int, maxRow: int) -> bool:
        if (maxRow + 1) * self.m - self.query_sum(1, 1, self.n, 1, maxRow + 1) < k:
            return False  # 剩余座位不足 k 个
        i = self.index(1, 1, self.n, maxRow + 1, self.m - 1)  # 从第一个没有坐满的排开始占座
        while True:
            left_seats = self.m - self.query_sum(1, 1, self.n, i, i)
            if k <= left_seats:  # 剩余人数不够坐后面的排
                self.add(1, 1, self.n, i, k)
                return True
            k -= left_seats
            self.add(1, 1, self.n, i, left_seats)
            i += 1
```

```java [sol1-Java]
class BookMyShow {
    int n, m;
    int[] min;
    long[] sum;

    public BookMyShow(int n, int m) {
        this.n = n;
        this.m = m;
        min = new int[n * 4];
        sum = new long[n * 4];
    }

    public int[] gather(int k, int maxRow) {
        int i = index(1, 1, n, maxRow + 1, m - k);
        if (i == 0) return new int[]{}; // 不存在
        var seats = (int) query_sum(1, 1, n, i, i);
        add(1, 1, n, i, k); // 占据 k 个座位
        return new int[]{i - 1, seats};
    }

    public boolean scatter(int k, int maxRow) {
        if ((long) (maxRow + 1) * m - query_sum(1, 1, n, 1, maxRow + 1) < k) return false; // 剩余座位不足 k 个
        // 从第一个没有坐满的排开始占座
        for (var i = index(1, 1, n, maxRow + 1, m - 1); ; ++i) {
            var left_seats = m - (int) query_sum(1, 1, n, i, i);
            if (k <= left_seats) { // 剩余人数不够坐后面的排
                add(1, 1, n, i, k);
                return true;
            }
            k -= left_seats;
            add(1, 1, n, i, left_seats);
        }
    }

    // 将 idx 上的元素值增加 val
    void add(int o, int l, int r, int idx, int val) {
        if (l == r) {
            min[o] += val;
            sum[o] += val;
            return;
        }
        var m = (l + r) / 2;
        if (idx <= m) add(o * 2, l, m, idx, val);
        else add(o * 2 + 1, m + 1, r, idx, val);
        min[o] = Math.min(min[o * 2], min[o * 2 + 1]);
        sum[o] = sum[o * 2] + sum[o * 2 + 1];
    }

    // 返回区间 [L,R] 内的元素和
    long query_sum(int o, int l, int r, int L, int R) { // L 和 R 在整个递归过程中均不变，将其大写，视作常量
        if (L <= l && r <= R) return sum[o];
        var sum = 0L;
        var m = (l + r) / 2;
        if (L <= m) sum += query_sum(o * 2, l, m, L, R);
        if (R > m) sum += query_sum(o * 2 + 1, m + 1, r, L, R);
        return sum;
    }

    // 返回区间 [1,R] 中 <= val 的最靠左的位置，不存在时返回 0
    int index(int o, int l, int r, int R, int val) { // R 在整个递归过程中均不变，将其大写，视作常量
        if (min[o] > val) return 0; // 说明整个区间的元素值都大于 val
        if (l == r) return l;
        var m = (l + r) / 2;
        if (min[o * 2] <= val) return index(o * 2, l, m, R, val); // 看看左半部分
        if (m < R) return index(o * 2 + 1, m + 1, r, R, val); // 看看右半部分
        return 0;
    }
}
```

```C++ [sol1-C++]
class BookMyShow {
    int n, m;
    vector<int> min;
    vector<long> sum;

    // 将 idx 上的元素值增加 val
    void add(int o, int l, int r, int idx, int val) {
        if (l == r) {
            min[o] += val;
            sum[o] += val;
            return;
        }
        int m = (l + r) / 2;
        if (idx <= m) add(o * 2, l, m, idx, val);
        else add(o * 2 + 1, m + 1, r, idx, val);
        min[o] = std::min(min[o * 2], min[o * 2 + 1]);
        sum[o] = sum[o * 2] + sum[o * 2 + 1];
    }

    // 返回区间 [L,R] 内的元素和
    long query_sum(int o, int l, int r, int L, int R) { // L 和 R 在整个递归过程中均不变，将其大写，视作常量
        if (L <= l && r <= R) return sum[o];
        long sum = 0L;
        int m = (l + r) / 2;
        if (L <= m) sum += query_sum(o * 2, l, m, L, R);
        if (R > m) sum += query_sum(o * 2 + 1, m + 1, r, L, R);
        return sum;
    }

    // 返回区间 [1,R] 中 <= val 的最靠左的位置，不存在时返回 0
    int index(int o, int l, int r, int R, int val) { // R 在整个递归过程中均不变，将其大写，视作常量
        if (min[o] > val) return 0; // 说明整个区间的元素值都大于 val
        if (l == r) return l;
        int m = (l + r) / 2;
        if (min[o * 2] <= val) return index(o * 2, l, m, R, val); // 看看左半部分
        if (m < R) return index(o * 2 + 1, m + 1, r, R, val); // 看看右半部分
        return 0;
    }

public:
    BookMyShow(int n, int m) : n(n), m(m), min(n * 4), sum(n * 4) {}

    vector<int> gather(int k, int maxRow) {
        int i = index(1, 1, n, maxRow + 1, m - k);
        if (i == 0) return {}; // 不存在
        int seats = query_sum(1, 1, n, i, i);
        add(1, 1, n, i, k); // 占据 k 个座位
        return {i - 1, seats};
    }

    bool scatter(int k, int maxRow) {
        if ((long) (maxRow + 1) * m - query_sum(1, 1, n, 1, maxRow + 1) < k) return false; // 剩余座位不足 k 个
        // 从第一个没有坐满的排开始占座
        for (int i = index(1, 1, n, maxRow + 1, m - 1);; ++i) {
            int left_seats = m - query_sum(1, 1, n, i, i);
            if (k <= left_seats) { // 剩余人数不够坐后面的排
                add(1, 1, n, i, k);
                return true;
            }
            k -= left_seats;
            add(1, 1, n, i, left_seats);
        }
    }
};
```

```go [sol1-Go]
type seg []struct{ l, r, min, sum int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

// 将 idx 上的元素值增加 val
func (t seg) add(o, idx, val int) {
	if t[o].l == t[o].r {
		t[o].min += val
		t[o].sum += val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if idx <= m {
		t.add(o<<1, idx, val)
	} else {
		t.add(o<<1|1, idx, val)
	}
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].min = min(lo.min, ro.min)
	t[o].sum = lo.sum + ro.sum
}

// 返回区间 [l,r] 内的元素和
func (t seg) querySum(o, l, r int) (sum int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		sum += t.querySum(o<<1, l, r)
	}
	if r > m {
		sum += t.querySum(o<<1|1, l, r)
	}
	return
}

// 返回区间 [1,R] 中 <= val 的最靠左的位置，不存在时返回 0
func (t seg) index(o, r, val int) int {
	if t[o].min > val { // 说明整个区间的元素值都大于 val
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	m := (t[o].l + t[o].r) >> 1
	if t[o<<1].min <= val { // 看看左半部分
		return t.index(o<<1, r, val)
	}
	if m < r { // 看看右半部分
		return t.index(o<<1|1, r, val)
	}
	return 0
}

type BookMyShow struct {
	seg
	m int
}

func Constructor(n, m int) BookMyShow {
	t := make(seg, n*4)
	t.build(1, 1, n)
	return BookMyShow{t, m}
}

func (t BookMyShow) Gather(k, maxRow int) []int {
	i := t.index(1, maxRow+1, t.m-k)
	if i == 0 { // 不存在
		return nil
	}
	seats := t.querySum(1, i, i)
	t.add(1, i, k) // 占据 k 个座位
	return []int{i - 1, seats}
}

func (t BookMyShow) Scatter(k, maxRow int) bool {
	if (maxRow+1)*t.m-t.querySum(1, 1, maxRow+1) < k { // 剩余座位不足 k 个
		return false
	}
	// 从第一个没有坐满的排开始占座
	for i := t.index(1, maxRow+1, t.m-1); ; i++ {
		leftSeats := t.m - t.querySum(1, i, i)
		if k <= leftSeats { // 剩余人数不够坐后面的排
			t.add(1, i, k)
			return true
		}
		k -= leftSeats
		t.add(1, i, leftSeats)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
```
