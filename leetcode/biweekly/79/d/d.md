## 题意（换一个场景）

一开始有 $n$ 个**空**水桶，每个水桶的容量都是 $m$ 升。水桶编号从 $0$ 到 $n-1$。

- $\texttt{gather}$：在前 $\textit{maxRow}$ 个水桶中，找第一个还能装至少 $k$ 升水的水桶，往里面倒入 $k$ 升水。如果有这样的水桶，返回水桶编号，以及在倒水前，水桶有多少升水；如果没有这样的水桶，返回空列表。
- $\texttt{scatter}$：往前 $\textit{maxRow}$ 个水桶中倒入**总量**为 $k$ 升的水。从左到右选择没有装满的水桶依次倒入。如果无法倒入总量为 $k$ 升的水，则不执行操作，并返回 $\texttt{false}$；否则执行操作，并返回 $\texttt{true}$。

## 思路

我们需要：

- 求出前 $\textit{maxRow}$ 个水桶中，第一个剩余容量 $\ge k$，也就是接水量 $\le m-k$ 的水桶。
- 维护每个水桶的接水量。
- 维护前 $\textit{maxRow}$ 个水桶的接水量之和，从而判断 $\texttt{scatter}$ 能否倒入总量为 $k$ 升的水。

这些都可以用**线段树**解决。线段树维护每个区间的**接水量的最小值** $\textit{min}$，以及每个区间的**接水量之和** $\textit{sum}$。

对于 $\texttt{gather}$，从线段树的根节点开始递归：

- 如果当前区间 $\textit{min}>m-k$，则无法倒入 $k$ 升水，返回 $0$。
- 如果当前区间长度为 $1$，返回区间端点。
- 如果左半区间 $\textit{min}\le m-k$，则答案在左半区间中，递归左半区间。
- 否则如果 $\textit{maxRow}$ 在右半区间内，递归右半区间。
- 否则返回 $-1$，表示没有这样的水桶。

> 上述过程叫做线段树二分。

对于 $\texttt{scatter}$，如果区间 $[0,\textit{maxRow}]$ 的接水量之和大于 $m\cdot (\textit{maxRow}+1)-k$，则无法执行操作。

否则可以执行操作。从第一个没有装满，也就是接水量 $\le m-1$ 的水桶开始倒水，这也可以用线段树二分求出。

关于线段树需要开多大的数组，推导过程可以看 [OI Wiki](https://oi-wiki.org/ds/seg/#%E5%AE%9E%E7%8E%B0)。

> 本题只有单点修改，没有区间更新，无需懒标记。

```py [sol-Python3]
class BookMyShow:
    def __init__(self, n: int, m: int):
        self.n = n
        self.m = m
        self.min = [0] * (2 << n.bit_length())  # 相比 4n 空间更小
        self.sum = [0] * (2 << n.bit_length())

    # 线段树：把下标 i 上的元素值增加 val
    def update(self, o: int, l: int, r: int, i: int, val: int) -> None:
        if l == r:
            self.min[o] += val
            self.sum[o] += val
            return
        m = (l + r) // 2
        if i <= m:
            self.update(o * 2, l, m, i, val)
        else:
            self.update(o * 2 + 1, m + 1, r, i, val)
        self.min[o] = min(self.min[o * 2], self.min[o * 2 + 1])
        self.sum[o] = self.sum[o * 2] + self.sum[o * 2 + 1]

    # 线段树：返回区间 [L,R] 内的元素和
    def query_sum(self, o: int, l: int, r: int, L: int, R: int) -> int:
        if L <= l and r <= R:
            return self.sum[o]
        res = 0
        m = (l + r) // 2
        if L <= m:
            res = self.query_sum(o * 2, l, m, L, R)
        if R > m:
            res += self.query_sum(o * 2 + 1, m + 1, r, L, R)
        return res

    # 线段树：返回区间 [0,R] 中 <= val 的最靠左的位置，不存在时返回 -1
    def find_first(self, o: int, l: int, r: int, R: int, val: int) -> int:
        if self.min[o] > val:
            return -1  # 整个区间的元素值都大于 val
        if l == r:
            return l
        m = (l + r) // 2
        if self.min[o * 2] <= val:
            return self.find_first(o * 2, l, m, R, val)
        if R > m:
            return self.find_first(o * 2 + 1, m + 1, r, R, val)
        return -1

    def gather(self, k: int, maxRow: int) -> List[int]:
        # 找第一个能倒入 k 升水的水桶
        r = self.find_first(1, 0, self.n - 1, maxRow, self.m - k)
        if r < 0:  # 没有这样的水桶
            return []
        c = self.query_sum(1, 0, self.n - 1, r, r)
        self.update(1, 0, self.n - 1, r, k)  # 倒水
        return [r, c]

    def scatter(self, k: int, maxRow: int) -> bool:
        # [0,maxRow] 的接水量之和
        s = self.query_sum(1, 0, self.n - 1, 0, maxRow)
        if s > self.m * (maxRow + 1) - k:
            return False  # 水桶已经装了太多的水
        # 从第一个没有装满的水桶开始
        i = self.find_first(1, 0, self.n - 1, maxRow, self.m - 1)
        while k:
            left = min(self.m - self.query_sum(1, 0, self.n - 1, i, i), k)
            self.update(1, 0, self.n - 1, i, left)  # 倒水
            k -= left
            i += 1
        return True
```

```java [sol-Java]
class BookMyShow {
    private int n;
    private int m;
    private int[] min;
    private long[] sum;

    public BookMyShow(int n, int m) {
        this.n = n;
        this.m = m;
        int size = 2 << (32 - Integer.numberOfLeadingZeros(n)); // 比 4n 更小
        min = new int[size];
        sum = new long[size];
    }

    public int[] gather(int k, int maxRow) {
        // 找第一个能倒入 k 升水的水桶
        int r = findFirst(1, 0, n - 1, maxRow, m - k);
        if (r < 0) { // 没有这样的水桶
            return new int[]{};
        }
        int c = (int) querySum(1, 0, n - 1, r, r);
        update(1, 0, n - 1, r, k); // 倒水
        return new int[]{r, c};
    }

    public boolean scatter(int k, int maxRow) {
        // [0,maxRow] 的接水量之和
        long s = querySum(1, 0, n - 1, 0, maxRow);
        if (s > (long) m * (maxRow + 1) - k) {
            return false; // 水桶已经装了太多的水
        }
        // 从第一个没有装满的水桶开始
        int i = findFirst(1, 0, n - 1, maxRow, m - 1);
        while (k > 0) {
            int left = Math.min(m - (int) querySum(1, 0, n - 1, i, i), k);
            update(1, 0, n - 1, i, left); // 倒水
            k -= left;
            i++;
        }
        return true;
    }

    // 把下标 i 上的元素值增加 val
    private void update(int o, int l, int r, int i, int val) {
        if (l == r) {
            min[o] += val;
            sum[o] += val;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) {
            update(o * 2, l, m, i, val);
        } else {
            update(o * 2 + 1, m + 1, r, i, val);
        }
        min[o] = Math.min(min[o * 2], min[o * 2 + 1]);
        sum[o] = sum[o * 2] + sum[o * 2 + 1];
    }

    // 返回区间 [L,R] 内的元素和
    private long querySum(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) {
            return sum[o];
        }
        long res = 0;
        int m = (l + r) / 2;
        if (L <= m) {
            res = querySum(o * 2, l, m, L, R);
        }
        if (R > m) {
            res += querySum(o * 2 + 1, m + 1, r, L, R);
        }
        return res;
    }

    // 返回区间 [0,R] 中 <= val 的最靠左的位置，不存在时返回 -1
    private int findFirst(int o, int l, int r, int R, int val) {
        if (min[o] > val) {
            return -1; // 整个区间的元素值都大于 val
        }
        if (l == r) {
            return l;
        }
        int m = (l + r) / 2;
        if (min[o * 2] <= val) {
            return findFirst(o * 2, l, m, R, val);
        }
        if (R > m) {
            return findFirst(o * 2 + 1, m + 1, r, R, val);
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class BookMyShow {
    int n, m;
    vector<int> mn;
    vector<long long> sum;

    // 把下标 i 上的元素值增加 val
    void update(int o, int l, int r, int i, int val) {
        if (l == r) {
            mn[o] += val;
            sum[o] += val;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) {
            update(o * 2, l, m, i, val);
        } else {
            update(o * 2 + 1, m + 1, r, i, val);
        }
        mn[o] = min(mn[o * 2], mn[o * 2 + 1]);
        sum[o] = sum[o * 2] + sum[o * 2 + 1];
    }

    // 返回区间 [L,R] 内的元素和
    long long querySum(int o, int l, int r, int L, int R) {
        if (L <= l && r <= R) {
            return sum[o];
        }
        long long res = 0;
        int m = (l + r) / 2;
        if (L <= m) {
            res = querySum(o * 2, l, m, L, R);
        }
        if (R > m) {
            res += querySum(o * 2 + 1, m + 1, r, L, R);
        }
        return res;
    }

    // 返回区间 [0,R] 中 <= val 的最靠左的位置，不存在时返回 -1
    int findFirst(int o, int l, int r, int R, int val) {
        if (mn[o] > val) {
            return -1; // 整个区间的元素值都大于 val
        }
        if (l == r) {
            return l;
        }
        int m = (l + r) / 2;
        if (mn[o * 2] <= val) {
            return findFirst(o * 2, l, m, R, val);
        }
        if (R > m) {
            return findFirst(o * 2 + 1, m + 1, r, R, val);
        }
        return -1;
    }

public:
    BookMyShow(int n, int m) : n(n), m(m), mn(4 << __lg(n)), sum(4 << __lg(n)) {}

    vector<int> gather(int k, int maxRow) {
        // 找第一个能倒入 k 升水的水桶
        int r = findFirst(1, 0, n - 1, maxRow, m - k);
        if (r < 0) { // 没有这样的水桶
            return {};
        }
        int c = querySum(1, 0, n - 1, r, r);
        update(1, 0, n - 1, r, k); // 倒水
        return {r, c};
    }

    bool scatter(int k, int maxRow) {
        // [0,maxRow] 的接水量之和
        long long s = querySum(1, 0, n - 1, 0, maxRow);
        if (s > (long long) m * (maxRow + 1) - k) {
            return false; // 水桶已经装了太多的水
        }
        // 从第一个没有装满的水桶开始
        int i = findFirst(1, 0, n - 1, maxRow, m - 1);
        while (k) {
            int left = min(m - (int) querySum(1, 0, n - 1, i, i), k);
            update(1, 0, n - 1, i, left); // 倒水
            k -= left;
            i++;
        }
        return true;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

typedef struct {
    int n, m;
    int* min;
    long long* sum;
} BookMyShow;

BookMyShow* bookMyShowCreate(int n, int m) {
    BookMyShow* obj = malloc(sizeof(BookMyShow));
    obj->n = n;
    obj->m = m;
    int size = 2 << (32 - __builtin_clz(n)); // 比 4n 更小
    obj->min = calloc(size, sizeof(int));
    obj->sum = calloc(size, sizeof(long long));
    return obj;
}

// 把下标 i 上的元素值增加 val
void update(BookMyShow* obj, int o, int l, int r, int i, int val) {
    if (l == r) {
        obj->min[o] += val;
        obj->sum[o] += val;
        return;
    }
    int mid = (l + r) / 2;
    if (i <= mid) {
        update(obj, o * 2, l, mid, i, val);
    } else {
        update(obj, o * 2 + 1, mid + 1, r, i, val);
    }
    obj->min[o] = MIN(obj->min[o * 2], obj->min[o * 2 + 1]);
    obj->sum[o] = obj->sum[o * 2] + obj->sum[o * 2 + 1];
}

// 返回区间 [L,R] 内的元素和
long long querySum(BookMyShow* obj, int o, int l, int r, int L, int R) {
    if (L <= l && r <= R) {
        return obj->sum[o];
    }
    long long res = 0;
    int m = (l + r) / 2;
    if (L <= m) {
        res = querySum(obj, o * 2, l, m, L, R);
    }
    if (R > m) {
        res += querySum(obj, o * 2 + 1, m + 1, r, L, R);
    }
    return res;
}

// 返回区间 [0,R] 中 <= val 的最靠左的位置，不存在时返回 -1
int findFirst(BookMyShow* obj, int o, int l, int r, int R, int val) {
    if (obj->min[o] > val) {
        return -1; // 整个区间的元素值都大于 val
    }
    if (l == r) {
        return l;
    }
    int m = (l + r) / 2;
    if (obj->min[o * 2] <= val) {
        return findFirst(obj, o * 2, l, m, R, val);
    }
    if (R > m) {
        return findFirst(obj, o * 2 + 1, m + 1, r, R, val);
    }
    return -1;
}

int* bookMyShowGather(BookMyShow* obj, int k, int maxRow, int* retSize) {
    // 找第一个能倒入 k 升水的水桶
    int r = findFirst(obj, 1, 0, obj->n - 1, maxRow, obj->m - k);
    if (r < 0) {
        *retSize = 0;
        return NULL; // 没有这样的水桶
    }
    int c = querySum(obj, 1, 0, obj->n - 1, r, r);
    update(obj, 1, 0, obj->n - 1, r, k); // 倒水
    *retSize = 2;
    int* ans = malloc(2 * sizeof(int));
    ans[0] = r;
    ans[1] = c;
    return ans;
}

bool bookMyShowScatter(BookMyShow* obj, int k, int maxRow) {
    // [0,maxRow] 的接水量之和
    long long s = querySum(obj, 1, 0, obj->n - 1, 0, maxRow);
    if (s > (long long) obj->m * (maxRow + 1) - k) {
        return false; // 水桶已经装了太多的水
    }
    // 从第一个没有装满的水桶开始
    int i = findFirst(obj, 1, 0, obj->n - 1, maxRow, obj->m - 1);
    while (k > 0) {
        int left = obj->m - (int) querySum(obj, 1, 0, obj->n - 1, i, i);
        left = MIN(left, k);
        update(obj, 1, 0, obj->n - 1, i, left); // 倒水
        k -= left;
        i++;
    }
    return true;
}

void bookMyShowFree(BookMyShow* obj) {
    free(obj->min);
    free(obj->sum);
    free(obj);
}
```

```go [sol-Go]
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

// 把下标 i 上的元素值增加 val
func (t seg) update(o, i, val int) {
    if t[o].l == t[o].r {
        t[o].min += val
        t[o].sum += val
        return
    }
    m := (t[o].l + t[o].r) >> 1
    if i <= m {
        t.update(o<<1, i, val)
    } else {
        t.update(o<<1|1, i, val)
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
        sum = t.querySum(o<<1, l, r)
    }
    if r > m {
        sum += t.querySum(o<<1|1, l, r)
    }
    return
}

// 返回区间 [0,r] 中 <= val 的最靠左的位置，不存在时返回 -1
func (t seg) findFirst(o, r, val int) int {
    if t[o].min > val {
        return -1 // 整个区间的元素值都大于 val
    }
    if t[o].l == t[o].r {
        return t[o].l
    }
    m := (t[o].l + t[o].r) / 2
    if t[o*2].min <= val {
        return t.findFirst(o*2, r, val)
    }
    if r > m {
        return t.findFirst(o*2+1, r, val)
    }
    return -1
}

type BookMyShow struct {
    seg
    n, m int
}

func Constructor(n, m int) BookMyShow {
    t := make(seg, 2<<bits.Len(uint(n-1))) // 比 4n 更小
    t.build(1, 0, n-1)
    return BookMyShow{t, n, m}
}

func (t *BookMyShow) Gather(k, maxRow int) []int {
    // 找第一个能倒入 k 升水的水桶
    r := t.findFirst(1, maxRow, t.m-k)
    if r < 0 { // 没有这样的水桶
        return nil
    }
    c := t.querySum(1, r, r)
    t.update(1, r, k) // 倒水
    return []int{r, c}
}

func (t *BookMyShow) Scatter(k, maxRow int) bool {
    // [0,maxRow] 的接水量之和
    s := t.querySum(1, 0, maxRow)
    if s > t.m*(maxRow+1)-k {
        return false // 水桶已经装了太多的水
    }
    // 从第一个没有装满的水桶开始
    i := t.findFirst(1, maxRow, t.m-1)
    for k > 0 {
        left := min(t.m-t.querySum(1, i, i), k)
        t.update(1, i, left) // 倒水
        k -= left
        i++
    }
    return true
}
```

```js [sol-JavaScript]
class BookMyShow {
    constructor(n, m) {
        this.n = n;
        this.m = m;
        const size = 2 << (32 - Math.clz32(n)); // 比 4n 更小
        this.min = Array(size).fill(0);
        this.sum = Array(size).fill(0);
    }

    // 把下标 i 上的元素值增加 val
    update(o, l, r, i, val) {
        if (l === r) {
            this.min[o] += val;
            this.sum[o] += val;
            return;
        }
        const m = Math.floor((l + r) / 2);
        if (i <= m) {
            this.update(o * 2, l, m, i, val);
        } else {
            this.update(o * 2 + 1, m + 1, r, i, val);
        }
        this.min[o] = Math.min(this.min[o * 2], this.min[o * 2 + 1]);
        this.sum[o] = this.sum[o * 2] + this.sum[o * 2 + 1];
    }

    // 返回区间 [L,R] 内的元素和
    querySum(o, l, r, L, R) {
        if (L <= l && r <= R) {
            return this.sum[o];
        }
        let res = 0;
        const m = Math.floor((l + r) / 2);
        if (L <= m) {
            res = this.querySum(o * 2, l, m, L, R);
        }
        if (R > m) {
            res += this.querySum(o * 2 + 1, m + 1, r, L, R);
        }
        return res;
    }

    // 返回区间 [0,R] 中 <= val 的最靠左的位置，不存在时返回 -1
    findFirst(o, l, r, R, val) {
        if (this.min[o] > val) {
            return -1; // 整个区间的元素值都大于 val
        }
        if (l === r) {
            return l;
        }
        const m = Math.floor((l + r) / 2);
        if (this.min[o * 2] <= val) {
            return this.findFirst(o * 2, l, m, R, val);
        }
        if (R > m) {
            return this.findFirst(o * 2 + 1, m + 1, r, R, val);
        }
        return -1;
    }

    gather(k, maxRow) {
        // 找第一个能倒入 k 升水的水桶
        const r = this.findFirst(1, 0, this.n - 1, maxRow, this.m - k);
        if (r < 0) { // 没有这样的水桶
            return [];
        }
        const c = this.querySum(1, 0, this.n - 1, r, r);
        this.update(1, 0, this.n - 1, r, k); // 倒水
        return [r, c];
    }

    scatter(k, maxRow) {
        // [0,maxRow] 的接水量之和
        const s = this.querySum(1, 0, this.n - 1, 0, maxRow);
        if (s > this.m * (maxRow + 1) - k) {
            return false; // 水桶已经装了太多的水
        }
        // 从第一个没有装满的水桶开始
        let i = this.findFirst(1, 0, this.n - 1, maxRow, this.m - 1);
        while (k) {
            const left = Math.min(this.m - this.querySum(1, 0, this.n - 1, i, i), k);
            this.update(1, 0, this.n - 1, i, left); // 倒水
            k -= left;
            i++;
        }
        return true;
    }
}
```

```rust [sol-Rust]
struct BookMyShow {
    n: usize,
    m: i32,
    min: Vec<i32>,
    sum: Vec<i64>,
}

impl BookMyShow {
    // 把下标 i 上的元素值增加 val
    fn update(&mut self, o: usize, l: usize, r: usize, i: usize, val: i32) {
        if l == r {
            self.min[o] += val;
            self.sum[o] += val as i64;
            return;
        }
        let m = (l + r) / 2;
        if i <= m {
            self.update(o * 2, l, m, i, val);
        } else {
            self.update(o * 2 + 1, m + 1, r, i, val);
        }
        self.min[o] = self.min[o * 2].min(self.min[o * 2 + 1]);
        self.sum[o] = self.sum[o * 2] + self.sum[o * 2 + 1];
    }

    // 返回区间 [L,R] 内的元素和
    fn query_sum(&self, o: usize, l: usize, r: usize, L: usize, R: usize) -> i64 {
        if L <= l && r <= R {
            return self.sum[o];
        }
        let mut res = 0;
        let m = (l + r) / 2;
        if L <= m {
            res = self.query_sum(o * 2, l, m, L, R);
        }
        if R > m {
            res += self.query_sum(o * 2 + 1, m + 1, r, L, R);
        }
        res
    }

    // 返回区间 [0,R] 中 <= val 的最靠左的位置，不存在时返回 -1
    fn find_first(&self, o: usize, l: usize, r: usize, R: usize, val: i32) -> i32 {
        if self.min[o] > val {
            return -1; // 整个区间的元素值都大于 val
        }
        if l == r {
            return l as i32;
        }
        let m = (l + r) / 2;
        if self.min[o * 2] <= val {
            return self.find_first(o * 2, l, m, R, val);
        }
        if R > m {
            return self.find_first(o * 2 + 1, m + 1, r, R, val);
        }
        -1
    }

    fn new(n: i32, m: i32) -> Self {
        let size = 2 << (32 - n.leading_zeros()) as usize;
        BookMyShow {
            n: n as usize,
            m,
            min: vec![0; size],
            sum: vec![0; size],
        }
    }

    fn gather(&mut self, k: i32, max_row: i32) -> Vec<i32> {
        // 找第一个能倒入 k 升水的水桶
        let r = self.find_first(1, 0, self.n - 1, max_row as usize, self.m - k);
        if r < 0 {
            return vec![]; // 没有这样的水桶
        }
        let c = self.query_sum(1, 0, self.n - 1, r as usize, r as usize) as i32;
        self.update(1, 0, self.n - 1, r as usize, k); // 倒水
        vec![r, c]
    }

    fn scatter(&mut self, mut k: i32, max_row: i32) -> bool {
        // [0,maxRow] 的接水量之和
        let s = self.query_sum(1, 0, self.n - 1, 0, max_row as usize);
        if s > (self.m as i64 * (max_row + 1) as i64) - k as i64 {
            return false; // 水桶已经装了太多的水
        }
        // 从第一个没有装满的水桶开始
        let mut i = self.find_first(1, 0, self.n - 1, max_row as usize, self.m - 1) as usize;
        while k > 0 {
            let left = k.min(self.m - self.query_sum(1, 0, self.n - 1, i, i) as i32);
            self.update(1, 0, self.n - 1, i, left); // 倒水
            k -= left;
            i += 1;
        }
        true
    }
}
```

#### 复杂度分析

- 时间复杂度：
  - 初始化为 $\mathcal{O}(n)$。
  - $\texttt{gather}$ 为 $\mathcal{O}(\log n)$。由于每次要么递归左半区间，要么递归右半区间，因此线段树二分的时间复杂度为线段树的树高，即 $\mathcal{O}(\log n)$。
  - $\texttt{scatter}$ 可以从整体上来分析：由于装满的水桶后面不会再遍历了，所有 $\texttt{scatter}$ 的循环次数之和为 $\mathcal{O}(n+q)$（$q$ 为 $\texttt{scatter}$ 的调用次数），所以时间复杂度**之和**为 $\mathcal{O}((n+q)\log n)$。如果近似认为 $n=q$，那么**均摊复杂度**为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。线段树需要 $\mathcal{O}(n)$ 的空间。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
