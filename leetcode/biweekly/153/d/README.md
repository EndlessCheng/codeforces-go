**前置题目**：[3499. 操作后最大活跃区段数 I](https://leetcode.cn/problems/maximize-active-section-with-trade-i/)。

设 $s$ 中的 $\texttt{1}$ 的个数为 $\textit{total}_1$。

设 $s$ 中的所有连续 $\texttt{0}$ 对应的（左闭右开）区间列表为 $a$。例如 $s=\texttt{0100}$ 中有两段连续 $\texttt{0}$，区间分别为 $[0,1)$ 和 $[2,4)$。

在前置题目中，我们知道，答案为 $\textit{total}_1$ 加上 $\texttt{0}$ 最多的 $\texttt{010}$ 子串中的 $\texttt{0}$ 的个数 $\textit{mx}$。

对于询问 $[\textit{ql},\textit{qr})$（改成右开），分类讨论：

- 如果 $[\textit{ql},\textit{qr})$ 中没有完整的区间，但包含一段完整的 $\texttt{1}$，那么 $\textit{mx}$ 为两个残缺的区间长度之和。
- 如果 $[\textit{ql},\textit{qr})$ 中有完整的区间，那么 $\textit{mx}$ 为以下三种情况的最大值：
   - $[\textit{ql},\textit{qr})$ 中的相邻完整区间的长度之和的最大值。这可以用线段树或者 ST 表统计。线段树的模板见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/)。
   - $\textit{ql}$ 所处的残缺区间与 $[\textit{ql},\textit{qr})$ 的第一个完整区间的长度之和。
   - $\textit{qr}$ 所处的残缺区间与 $[\textit{ql},\textit{qr})$ 的最后一个完整区间的长度之和。

计算 $[\textit{ql},\textit{qr})$ 中的第一个完整区间和最后一个完整区间，可以用二分查找。

对于最后一个完整区间，可以先二分找到右端点 $> \textit{qr}$ 的第一个区间，这个区间的左边相邻区间，就是最后一个完整区间。

代码实现时，可以用哨兵简化代码，无需判断下标是否在边界上。可以把计算两个区间长度之和的逻辑，封装成一个函数。

[本题视频讲解](https://www.bilibili.com/video/BV1JrZzYhEHt/?t=6m9s)，欢迎点赞关注~

## 写法一：二分查找

```py [sol-Python3]
class SparseTable:
    def __init__(self, nums: List[int]):
        n = len(nums) - 1
        w = n.bit_length()
        st = [[0] * n for _ in range(w)]
        st[0] = [r1 - l1 + r2 - l2 for (l1, r1), (l2, r2) in pairwise(nums)]
        for i in range(1, w):
            for j in range(n - (1 << i) + 1):
                st[i][j] = max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))])
        self.st = st

    # 查询区间最大值，[l, r) 左闭右开
    def query(self, l: int, r: int) -> int:
        if l >= r:
            return 0
        k = (r - l).bit_length() - 1
        return max(self.st[k][l], self.st[k][r - (1 << k)])


class Solution:
    def maxActiveSectionsAfterTrade(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s)
        total1 = 0
        # 统计连续 0 段对应的区间（左闭右开）
        a = [(-1, -1)]  # 哨兵
        start = 0
        for i, b in enumerate(s):
            if i == n - 1 or b != s[i + 1]:
                if b == '1':
                    total1 += i - start + 1
                else:
                    a.append((start, i + 1))  # 左闭右开
                start = i + 1
        a.append((n + 1, n + 1))  # 哨兵

        def merge(x: int, y: int) -> int:
            return x + y if x > 0 and y > 0 else 0

        st = SparseTable(a)
        ans = []
        for ql, qr in queries:
            qr += 1  # 左闭右开
            i = bisect_left(a, ql, key=lambda p: p[0])
            j = bisect_right(a, qr, key=lambda p: p[1]) - 1
            mx = 0
            if i <= j:  # [ql,qr) 中有完整的区间
                mx = max(
                    st.query(i, j),  # 相邻完整区间的长度之和的最大值
                    merge(a[i - 1][1] - ql, a[i][1] - a[i][0]),  # 残缺区间 i-1 + 完整区间 i
                    merge(qr - a[j + 1][0], a[j][1] - a[j][0]),  # 残缺区间 j+1 + 完整区间 j
                )
            elif i == j + 1:  # [ql,qr) 中有两个相邻的残缺区间
                mx = merge(a[i - 1][1] - ql, qr - a[j + 1][0])  # 残缺区间 i-1 + 残缺区间 j+1
            ans.append(total1 + mx)
        return ans
```

```java [sol-Java]
class Solution {
    private record Pair(int l, int r) { // 左闭右开
    }

    private static class SparseTable {
        private final int[][] st;

        SparseTable(List<Pair> a) {
            int n = a.size() - 1;
            int w = 32 - Integer.numberOfLeadingZeros(n);
            st = new int[w][n];
            for (int j = 0; j < n; j++) {
                st[0][j] = a.get(j).r - a.get(j).l + a.get(j + 1).r - a.get(j + 1).l;
            }
            for (int i = 1; i < w; i++) {
                for (int j = 0; j + (1 << i) <= n; j++) {
                    st[i][j] = Math.max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))]);
                }
            }
        }

        // 查询区间最大值，[l,r) 左闭右开
        int query(int l, int r) {
            if (l >= r) {
                return 0;
            }
            int k = 32 - Integer.numberOfLeadingZeros(r - l) - 1;
            return Math.max(st[k][l], st[k][r - (1 << k)]);
        }
    }

    public List<Integer> maxActiveSectionsAfterTrade(String S, int[][] queries) {
        char[] s = S.toCharArray();
        int n = s.length;
        int total1 = 0;
        List<Pair> a = new ArrayList<>();
        a.add(new Pair(-1, -1)); // 哨兵
        int start = 0;
        for (int i = 0; i < n; i++) {
            if (i == n - 1 || s[i] != s[i + 1]) {
                if (s[i] == '1') {
                    total1 += i - start + 1;
                } else {
                    a.add(new Pair(start, i + 1)); // 左闭右开
                }
                start = i + 1;
            }
        }
        a.add(new Pair(n + 1, n + 1)); // 哨兵

        SparseTable st = new SparseTable(a);
        List<Integer> ans = new ArrayList<>(queries.length);
        for (int[] query : queries) {
            int ql = query[0];
            int qr = query[1] + 1; // 左闭右开

            // a 中没有重复的区间左右端点，可以直接用库函数二分
            // 找第一个区间，左端点 >= ql
            int i = Collections.binarySearch(a, new Pair(ql, 0), (p, q) -> p.l - q.l);
            if (i < 0) i = ~i;
            // 找最后一个区间，右端点 <= qr
            int j = Collections.binarySearch(a, new Pair(0, qr + 1), (p, q) -> p.r - q.r);
            if (j < 0) j = ~j;
            j--;

            int mx = 0;
            if (i <= j) { // [ql,qr) 中有完整的区间
                int full = st.query(i, j); // 相邻完整区间的长度之和的最大值
                int sl = merge(a.get(i - 1).r - ql, a.get(i).r - a.get(i).l); // 残缺区间 i-1 + 完整区间 i
                int sr = merge(qr - a.get(j + 1).l, a.get(j).r - a.get(j).l); // 残缺区间 j+1 + 完整区间 j
                mx = Math.max(full, Math.max(sl, sr));
            } else if (i == j + 1) { // [ql,qr) 中有两个相邻的残缺区间
                mx = merge(a.get(i - 1).r - ql, qr - a.get(j + 1).l); // 残缺区间 i-1 + 残缺区间 j+1
            }
            ans.add(total1 + mx);
        }
        return ans;
    }

    private int merge(int x, int y) {
        return x > 0 && y > 0 ? x + y : 0;
    }
}
```

```cpp [sol-C++]
struct Pair { int l, r; }; // 左闭右开

class SparseTable {
    vector<vector<int>> st;

public:
    SparseTable(vector<Pair>& a) {
        int n = a.size() - 1;
        int w = bit_width(unsigned(n));
        st.resize(w, vector<int>(n));
        for (int j = 0; j < n; j++) {
            st[0][j] = a[j].r - a[j].l + a[j + 1].r - a[j + 1].l;
        }
        for (int i = 1; i < w; i++) {
            for (int j = 0; j + (1 << i) <= n; j++) {
                st[i][j] = max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))]);
            }
        }
    }

    // 查询区间最大值，[l,r) 左闭右开
    int query(int l, int r) const {
        if (l >= r) {
            return 0;
        }
        int k = bit_width(unsigned(r - l)) - 1;
        return max(st[k][l], st[k][r - (1 << k)]);
    }
};

class Solution {
public:
    vector<int> maxActiveSectionsAfterTrade(string s, vector<vector<int>>& queries) {
        int n = s.size();
        int total1 = 0;
        vector<Pair> a = {{-1, -1}}; // 哨兵
        int start = 0;
        for (int i = 0; i < n; i++) {
            if (i == n - 1 || s[i] != s[i + 1]) {
                if (s[i] == '1') {
                    total1 += i - start + 1;
                } else {
                    a.emplace_back(start, i + 1); // 左闭右开
                }
                start = i + 1;
            }
        }
        a.emplace_back(n + 1, n + 1); // 哨兵

        auto merge = [](int x, int y) {
            return x > 0 && y > 0 ? x + y : 0;
        };

        SparseTable st(a);
        vector<int> ans(queries.size());
        for (int qi = 0; qi < queries.size(); qi++) {
            int ql = queries[qi][0], qr = queries[qi][1] + 1; // 左闭右开
            int i = ranges::lower_bound(a, ql, {}, &Pair::l) - a.begin();
            int j = ranges::upper_bound(a, qr, {}, &Pair::r) - a.begin() - 1;
            int mx = 0;
            if (i <= j) { // [ql,qr) 中有完整的区间
                mx = max({
                     st.query(i, j), // 相邻完整区间的长度之和的最大值
                     merge(a[i - 1].r - ql, a[i].r - a[i].l), // 残缺区间 i-1 + 完整区间 i
                     merge(qr - a[j + 1].l, a[j].r - a[j].l), // 残缺区间 j+1 + 完整区间 j
                 });
            } else if (i == j + 1) { // [ql,qr) 中有两个相邻的残缺区间
                mx = merge(a[i - 1].r - ql, qr - a[j + 1].l); // 残缺区间 i-1 + 残缺区间 j+1
            }
            ans[qi] = total1 + mx;
        }
        return ans;
    }
};
```

```go [sol-Go]
type pair struct{ l, r int } // 左闭右开
type ST [][]int

func newST(a []pair) ST {
	n := len(a) - 1
	w := bits.Len(uint(n))
	st := make(ST, w)
	for i := range st {
		st[i] = make([]int, n)
	}
	for j, p := range a[:n] {
		st[0][j] = p.r - p.l + a[j+1].r - a[j+1].l
	}
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = max(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return st
}

// 查询区间最大值，[l,r) 左闭右开
func (st ST) query(l, r int) int {
	if l >= r {
		return 0
	}
	k := bits.Len(uint(r-l)) - 1
	return max(st[k][l], st[k][r-1<<k])
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)
	total1 := 0
	// 统计连续 0 段对应的区间（左闭右开）
	a := []pair{{-1, -1}} // 哨兵
	start := 0
	for i := range n {
		if i == n-1 || s[i] != s[i+1] {
			if s[i] == '1' {
				total1 += i - start + 1
			} else {
				a = append(a, pair{start, i + 1}) // 左闭右开
			}
			start = i + 1
		}
	}
	a = append(a, pair{n + 1, n + 1}) // 哨兵

	merge := func(x, y int) int {
		if x > 0 && y > 0 {
			return x + y
		}
		return 0
	}

	st := newST(a)
	m := len(a)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		ql, qr := q[0], q[1]+1 // 左闭右开
		i := sort.Search(m, func(i int) bool { return a[i].l >= ql })
		j := sort.Search(m, func(i int) bool { return a[i].r > qr }) - 1
		mx := 0
		if i <= j { // [ql,qr) 中有完整的区间
			mx = max(
				st.query(i, j),                    // 相邻完整区间的长度之和的最大值
				merge(a[i-1].r-ql, a[i].r-a[i].l), // 残缺区间 i-1 + 完整区间 i
				merge(qr-a[j+1].l, a[j].r-a[j].l), // 残缺区间 j+1 + 完整区间 j
			)
		} else if i == j+1 { // [ql,qr) 中有两个相邻的残缺区间
			mx = merge(a[i-1].r-ql, qr-a[j+1].l) // 残缺区间 i-1 + 残缺区间 j+1
		}
		ans[qi] = total1 + mx
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log n)$。

## 写法二：标记每个字符所属区间

```py [sol-Python3]
class SparseTable:
    def __init__(self, nums: List[int]):
        n = len(nums) - 1
        w = n.bit_length()
        st = [[0] * n for _ in range(w)]
        st[0] = [r1 - l1 + r2 - l2 for (l1, r1), (l2, r2) in pairwise(nums)]
        for i in range(1, w):
            for j in range(n - (1 << i) + 1):
                st[i][j] = max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))])
        self.st = st

    # 查询区间最大值，[l, r) 左闭右开
    def query(self, l: int, r: int) -> int:
        if l >= r:
            return 0
        k = (r - l).bit_length() - 1
        return max(self.st[k][l], self.st[k][r - (1 << k)])


class Solution:
    def maxActiveSectionsAfterTrade(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s)
        total1 = 0
        belong = [0] * n  # 每个 0 所属的区间下标，每个 1 右边最近的 0 区间下标
        a = [(-1, -1)]
        start = 0
        for i, b in enumerate(s):
            belong[i] = len(a)  # 标记
            if i == n - 1 or b != s[i + 1]:
                if b == '1':
                    total1 += i - start + 1
                else:
                    a.append((start, i + 1))
                start = i + 1
        a.append((n + 1, n + 1))

        def merge(x: int, y: int) -> int:
            return x + y if x > 0 and y > 0 else 0

        st = SparseTable(a)
        ans = []
        for ql, qr in queries:
            i = belong[ql]
            if ql and s[ql] == '0' == s[ql - 1]:
                i += 1  # i 在残缺区间中
            j = belong[qr] - 1
            if qr + 1 < n and s[qr] == '0' != s[qr + 1]:
                j += 1  # j 刚好在完整区间的右端点，无需减一
            qr += 1

            mx = 0
            if i <= j:
                mx = max(
                    st.query(i, j),
                    merge(a[i - 1][1] - ql, a[i][1] - a[i][0]),
                    merge(qr - a[j + 1][0], a[j][1] - a[j][0]),
                )
            elif i == j + 1:
                mx = merge(a[i - 1][1] - ql, qr - a[j + 1][0])
            ans.append(total1 + mx)
        return ans
```

```java [sol-Java]
class Solution {
    private record Pair(int l, int r) {
    }

    private static class SparseTable {
        private final int[][] st;

        SparseTable(List<Pair> a) {
            int n = a.size() - 1;
            int w = 32 - Integer.numberOfLeadingZeros(n);
            st = new int[w][n];
            for (int j = 0; j < n; j++) {
                st[0][j] = a.get(j).r - a.get(j).l + a.get(j + 1).r - a.get(j + 1).l;
            }
            for (int i = 1; i < w; i++) {
                for (int j = 0; j + (1 << i) <= n; j++) {
                    st[i][j] = Math.max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))]);
                }
            }
        }

        // 查询区间最大值，[l,r) 左闭右开
        int query(int l, int r) {
            if (l >= r) {
                return 0;
            }
            int k = 32 - Integer.numberOfLeadingZeros(r - l) - 1;
            return Math.max(st[k][l], st[k][r - (1 << k)]);
        }
    }

    public List<Integer> maxActiveSectionsAfterTrade(String S, int[][] queries) {
        char[] s = S.toCharArray();
        int n = s.length;
        int total1 = 0;
        int[] belong = new int[n]; // 每个 0 所属的区间下标，每个 1 右边最近的 0 区间下标
        List<Pair> a = new ArrayList<>();
        a.add(new Pair(-1, -1));
        int start = 0;
        for (int i = 0; i < n; i++) {
            belong[i] = a.size(); // 标记
            if (i == n - 1 || s[i] != s[i + 1]) {
                if (s[i] == '1') {
                    total1 += i - start + 1;
                } else {
                    a.add(new Pair(start, i + 1));
                }
                start = i + 1;
            }
        }
        a.add(new Pair(n + 1, n + 1));

        SparseTable st = new SparseTable(a);
        List<Integer> ans = new ArrayList<>(queries.length);
        for (int[] query : queries) {
            int ql = query[0];
            int qr = query[1];

            int i = belong[ql];
            if (ql > 0 && s[ql] == '0' && s[ql - 1] == '0') {
                i++; // i 在残缺区间中
            }
            int j = belong[qr] - 1;
            if (qr + 1 < n && s[qr] == '0' && s[qr + 1] == '1') {
                j++; // j 刚好在完整区间的右端点，无需减一
            }
            qr++;

            int mx = 0;
            if (i <= j) {
                int full = st.query(i, j);
                int sl = merge(a.get(i - 1).r - ql, a.get(i).r - a.get(i).l);
                int sr = merge(qr - a.get(j + 1).l, a.get(j).r - a.get(j).l);
                mx = Math.max(full, Math.max(sl, sr));
            } else if (i == j + 1) {
                mx = merge(a.get(i - 1).r - ql, qr - a.get(j + 1).l);
            }
            ans.add(total1 + mx);
        }
        return ans;
    }

    private int merge(int x, int y) {
        return x > 0 && y > 0 ? x + y : 0;
    }
}
```

```cpp [sol-C++]
struct Pair { int l, r; };

class SparseTable {
    vector<vector<int>> st;

public:
    SparseTable(vector<Pair>& a) {
        int n = a.size() - 1;
        int w = bit_width(unsigned(n));
        st.resize(w, vector<int>(n));
        for (int j = 0; j < n; j++) {
            st[0][j] = a[j].r - a[j].l + a[j + 1].r - a[j + 1].l;
        }
        for (int i = 1; i < w; i++) {
            for (int j = 0; j + (1 << i) <= n; j++) {
                st[i][j] = max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))]);
            }
        }
    }

    // 查询区间最大值，[l,r) 左闭右开
    int query(int l, int r) const {
        if (l >= r) {
            return 0;
        }
        int k = bit_width(unsigned(r - l)) - 1;
        return max(st[k][l], st[k][r - (1 << k)]);
    }
};

class Solution {
public:
    vector<int> maxActiveSectionsAfterTrade(string s, vector<vector<int>>& queries) {
        int n = s.size();
        int total1 = 0;
        vector<int> belong(n); // 每个 0 所属的区间下标，每个 1 右边最近的 0 区间下标
        vector<Pair> a = {{-1, -1}};
        int start = 0;
        for (int i = 0; i < n; i++) {
            belong[i] = a.size(); // 标记
            if (i == n - 1 || s[i] != s[i + 1]) {
                if (s[i] == '1') {
                    total1 += i - start + 1;
                } else {
                    a.emplace_back(start, i + 1);
                }
                start = i + 1;
            }
        }
        a.emplace_back(n + 1, n + 1);

        auto merge = [](int x, int y) {
            return x > 0 && y > 0 ? x + y : 0;
        };

        SparseTable st(a);
        vector<int> ans(queries.size());
        for (int qi = 0; qi < queries.size(); qi++) {
            int ql = queries[qi][0], qr = queries[qi][1];

            int i = belong[ql];
            if (ql > 0 && s[ql] == '0' && s[ql - 1] == '0') {
                i++; // i 在残缺区间中
            }
            int j = belong[qr] - 1;
            if (qr + 1 < n && s[qr] == '0' && s[qr + 1] == '1') {
                j++; // j 刚好在完整区间的右端点，无需减一
            }
            qr++;

            int mx = 0;
            if (i <= j) {
                mx = max({
                     st.query(i, j),
                     merge(a[i - 1].r - ql, a[i].r - a[i].l),
                     merge(qr - a[j + 1].l, a[j].r - a[j].l),
                 });
            } else if (i == j + 1) {
                mx = merge(a[i - 1].r - ql, qr - a[j + 1].l);
            }
            ans[qi] = total1 + mx;
        }
        return ans;
    }
};
```

```go [sol-Go]
type pair struct{ l, r int } // 左闭右开
type ST [][]int

func newST(a []pair) ST {
	n := len(a) - 1
	w := bits.Len(uint(n))
	st := make(ST, w)
	for i := range st {
		st[i] = make([]int, n)
	}
	for j, p := range a[:n] {
		st[0][j] = p.r - p.l + a[j+1].r - a[j+1].l
	}
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = max(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return st
}

// 查询区间最大值，[l,r) 左闭右开
func (st ST) query(l, r int) int {
	if l >= r {
		return 0
	}
	k := bits.Len(uint(r-l)) - 1
	return max(st[k][l], st[k][r-1<<k])
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)
	total1 := 0
	belong := make([]int, n) // 每个 0 所属的区间下标，每个 1 右边最近的 0 区间下标
	a := []pair{{-1, -1}}
	start := 0
	for i, b := range s {
		belong[i] = len(a) // 记录
		if i == n-1 || byte(b) != s[i+1] {
			if s[i] == '1' {
				total1 += i - start + 1
			} else {
				a = append(a, pair{start, i + 1})
			}
			start = i + 1
		}
	}
	a = append(a, pair{n + 1, n + 1})

	merge := func(x, y int) int {
		if x > 0 && y > 0 {
			return x + y
		}
		return 0
	}

	st := newST(a)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		ql, qr := q[0], q[1]

		i := belong[ql]
		if ql > 0 && s[ql] == '0' && s[ql-1] == '0' {
			i++ // i 在残缺区间中
		}
		j := belong[qr] - 1
		if qr+1 < n && s[qr] == '0' && s[qr+1] == '1' {
			j++ // j 刚好在完整区间的右端点，无需减一
		}
		qr++

		mx := 0
		if i <= j {
			mx = max(
				st.query(i, j),
				merge(a[i-1].r-ql, a[i].r-a[i].l),
				merge(qr-a[j+1].l, a[j].r-a[j].l),
			)
		} else if i == j+1 {
			mx = merge(a[i-1].r-ql, qr-a[j+1].l)
		}
		ans[qi] = total1 + mx
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + q)$，其中 $n$ 是 $s$ 的长度。预处理后，可以做到 $\mathcal{O}(1)$ 回答每个询问！
- 空间复杂度：$\mathcal{O}(n\log n)$。

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
