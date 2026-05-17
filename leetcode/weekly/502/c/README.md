## 方法一：二维 ST 表

[二维 ST 表视频讲解](https://www.bilibili.com/video/BV18gLE6VETZ/)，先讲一维原理，再过渡到二维。

[二维 ST 表文字版讲解](https://blog.nowcoder.net/n/3eccd1386a8846398d3bee62b485309b)。

预处理二维 ST 表，即可 $\mathcal{O}(1)$ 算出任意子矩阵的最大值。

需要注意的是，本题不能计入子矩阵的四个角，怎么办？

![w502c.png](https://pic.leetcode.cn/1778988029-AamMpU-w502c.png){:width=300px}

我们可以把蓝色区域视作**两个子矩阵的并集**。以上图（示例 1）为例说明：

- 计算以 $2$ 为中心的 $3$ 行 $5$ 列的子矩阵的最大值。
- 计算以 $2$ 为中心的 $5$ 行 $3$ 列的子矩阵的最大值。
- 这两个最大值的最大值，即为蓝色区域的最大值。

```py [sol-Python3]
class Solution:
    def countLocalMaximums(self, matrix: list[list[int]]) -> int:
        n, m = len(matrix), len(matrix[0])
        wn, wm = n.bit_length(), m.bit_length()

        # st[k1][k2][i][j] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
        st = [[[[0] * m for _ in range(n)] for _ in range(wm)] for _ in range(wn)]

        # 初始值
        st[0][0] = matrix

        # 单独计算 k1 = 0
        for k2 in range(1, wm):
            half = 1 << (k2 - 1)
            for i in range(n):
                for j in range(m - (1 << k2) + 1):
                    st[0][k2][i][j] = max(st[0][k2 - 1][i][j], st[0][k2 - 1][i][j + half])

        for k1 in range(1, wn):
            half = 1 << (k1 - 1)
            for k2 in range(wm):
                for i in range(n - (1 << k1) + 1):
                    for j in range(m - (1 << k2) + 1):
                        st[k1][k2][i][j] = max(st[k1 - 1][k2][i][j], st[k1 - 1][k2][i + half][j])

        # 返回子矩阵最大值
        # 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
        def query(r1: int, c1: int, r2: int, c2: int) -> int:
            r1 = max(r1, 0)
            c1 = max(c1, 0)
            r2 = min(r2, n)
            c2 = min(c2, m)
            k1 = (r2 - r1).bit_length() - 1
            k2 = (c2 - c1).bit_length() - 1
            # 视作四个子矩阵的并集
            return max(st[k1][k2][r1][c1],
                       st[k1][k2][r2 - (1 << k1)][c1],
                       st[k1][k2][r1][c2 - (1 << k2)],
                       st[k1][k2][r2 - (1 << k1)][c2 - (1 << k2)])

        ans = 0
        for i, row in enumerate(matrix):
            for j, x in enumerate(row):
                if x > 0 and max(query(i - x, j - x + 1, i + x + 1, j + x), query(i - x + 1, j - x, i + x, j + x + 1)) <= x:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countLocalMaximums(int[][] matrix) {
        int n = matrix.length;
        int m = matrix[0].length;
        int wn = 32 - Integer.numberOfLeadingZeros(n);
        int wm = 32 - Integer.numberOfLeadingZeros(m);

        // st[k1][k2][i][j] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
        int[][][][] st = new int[wn][wm][n][m];

        // 初始值
        st[0][0] = matrix;

        // 单独计算 k1 = 0
        for (int k2 = 1; k2 < wm; k2++) {
            int half = 1 << (k2 - 1);
            for (int i = 0; i < n; i++) {
                for (int j = 0; j <= m - (1 << k2); j++) {
                    st[0][k2][i][j] = Math.max(st[0][k2 - 1][i][j], st[0][k2 - 1][i][j + half]);
                }
            }
        }

        for (int k1 = 1; k1 < wn; k1++) {
            int half = 1 << (k1 - 1);
            for (int k2 = 0; k2 < wm; k2++) {
                for (int i = 0; i <= n - (1 << k1); i++) {
                    for (int j = 0; j <= m - (1 << k2); j++) {
                        st[k1][k2][i][j] = Math.max(st[k1 - 1][k2][i][j], st[k1 - 1][k2][i + half][j]);
                    }
                }
            }
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int x = matrix[i][j];
                if (x == 0) {
                    continue;
                }
                int max1 = query(st, i - x, j - x + 1, i + x + 1, j + x, n, m);
                int max2 = query(st, i - x + 1, j - x, i + x, j + x + 1, n, m);
                if (Math.max(max1, max2) <= x) {
                    ans++;
                }
            }
        }
        return ans;
    }

    // 返回子矩阵最大值
    // 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
    private int query(int[][][][] st, int r1, int c1, int r2, int c2, int n, int m) {
        r1 = Math.max(r1, 0);
        c1 = Math.max(c1, 0);
        r2 = Math.min(r2, n);
        c2 = Math.min(c2, m);
        int k1 = 31 - Integer.numberOfLeadingZeros(r2 - r1);
        int k2 = 31 - Integer.numberOfLeadingZeros(c2 - c1);
        // 视作四个子矩阵的并集
        return Math.max(
                Math.max(st[k1][k2][r1][c1], st[k1][k2][r2 - (1 << k1)][c1]),
                Math.max(st[k1][k2][r1][c2 - (1 << k2)], st[k1][k2][r2 - (1 << k1)][c2 - (1 << k2)])
        );
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countLocalMaximums(vector<vector<int>>& matrix) {
        int n = matrix.size(), m = matrix[0].size();
        int wn = bit_width(1u * n), wm = bit_width(1u * m);

        // st[i][j][k1][k2] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
        // 用 array 内存连续性更好。不用 array 的做法见【C++ 写法二】
        vector st(n, vector<array<array<int, 8>, 8>>(m));

        // 初始值
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                st[i][j][0][0] = matrix[i][j];
            }
        }

        // 单独计算 k1 = 0
        for (int k2 = 1; k2 < wm; k2++) {
            for (int i = 0; i < n; i++) {
                for (int j = 0; j <= m - (1 << k2); j++) {
                    st[i][j][0][k2] = max(st[i][j][0][k2 - 1], st[i][j + (1 << (k2 - 1))][0][k2 - 1]);
                }
            }
        }

        for (int k1 = 1; k1 < wn; k1++) {
            for (int k2 = 0; k2 < wm; k2++) {
                for (int i = 0; i <= n - (1 << k1); i++) {
                    for (int j = 0; j <= m - (1 << k2); j++) {
                        st[i][j][k1][k2] = max(st[i][j][k1 - 1][k2], st[i + (1 << (k1 - 1))][j][k1 - 1][k2]);
                    }
                }
            }
        }

        // 返回子矩阵最大值
        // 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
        auto query = [&](int r1, int c1, int r2, int c2) -> int {
            r1 = max(r1, 0);
            c1 = max(c1, 0);
            r2 = min(r2, n);
            c2 = min(c2, m);
            int k1 = bit_width(1u * (r2 - r1)) - 1;
            int k2 = bit_width(1u * (c2 - c1)) - 1;
            // 视作四个子矩阵的并集
            return max({
                st[r1][c1][k1][k2],
                st[r2 - (1 << k1)][c1][k1][k2],
                st[r1][c2 - (1 << k2)][k1][k2],
                st[r2 - (1 << k1)][c2 - (1 << k2)][k1][k2],
            });
        };

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int x = matrix[i][j];
                if (x > 0 && max(query(i - x, j - x + 1, i + x + 1, j + x), query(i - x + 1, j - x, i + x, j + x + 1)) <= x) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```cpp [sol-C++ 写法二]
class Solution {
public:
    int countLocalMaximums(vector<vector<int>>& matrix) {
        int n = matrix.size(), m = matrix[0].size();
        int wn = bit_width(1u * n), wm = bit_width(1u * m);

        // st[k1][k2][i][j] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
        vector st(wn, vector(wm, vector(n, vector<int>(m))));

        // 初始值
        st[0][0] = matrix;

        // 单独计算 k1 = 0
        for (int k2 = 1; k2 < wm; k2++) {
            for (int i = 0; i < n; i++) {
                for (int j = 0; j <= m - (1 << k2); j++) {
                    st[0][k2][i][j] = max(st[0][k2 - 1][i][j], st[0][k2 - 1][i][j + (1 << (k2 - 1))]);
                }
            }
        }

        for (int k1 = 1; k1 < wn; k1++) {
            for (int k2 = 0; k2 < wm; k2++) {
                for (int i = 0; i <= n - (1 << k1); i++) {
                    for (int j = 0; j <= m - (1 << k2); j++) {
                        st[k1][k2][i][j] = max(st[k1 - 1][k2][i][j], st[k1 - 1][k2][i + (1 << (k1 - 1))][j]);
                    }
                }
            }
        }

        // 返回子矩阵最大值
        // 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
        auto query = [&](int r1, int c1, int r2, int c2) -> int {
            r1 = max(r1, 0);
            c1 = max(c1, 0);
            r2 = min(r2, n);
            c2 = min(c2, m);
            int k1 = bit_width(1u * (r2 - r1)) - 1;
            int k2 = bit_width(1u * (c2 - c1)) - 1;
            // 视作四个子矩阵的并集
            return max({
                st[k1][k2][r1][c1],
                st[k1][k2][r2 - (1 << k1)][c1],
                st[k1][k2][r1][c2 - (1 << k2)],
                st[k1][k2][r2 - (1 << k1)][c2 - (1 << k2)],
            });
        };

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int x = matrix[i][j];
                if (x > 0 && max(query(i - x, j - x + 1, i + x + 1, j + x), query(i - x + 1, j - x, i + x, j + x + 1)) <= x) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countLocalMaximums(matrix [][]int) (ans int) {
	n, m := len(matrix), len(matrix[0])
	wn, wm := bits.Len(uint(n)), bits.Len(uint(m))

	// st[k1][k2][n][m] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
	st := make([][][][]int, wn)
	for k1 := range st {
		st[k1] = make([][][]int, wm)
		for k2 := range st[k1] {
			st[k1][k2] = make([][]int, n)
			for i := range st[k1][k2] {
				st[k1][k2][i] = make([]int, m)
			}
		}
	}

	// 初始值
	st[0][0] = matrix

	// 单独计算 k1 = 0
	for k2 := 1; k2 < wm; k2++ {
		for i := range n {
			for j := range m - 1<<k2 + 1 {
				st[0][k2][i][j] = max(st[0][k2-1][i][j], st[0][k2-1][i][j+1<<(k2-1)])
			}
		}
	}

	for k1 := 1; k1 < wn; k1++ {
		for k2 := range wm {
			for i := range n - 1<<k1 + 1 {
				for j := range m - 1<<k2 + 1 {
					st[k1][k2][i][j] = max(st[k1-1][k2][i][j], st[k1-1][k2][i+1<<(k1-1)][j])
				}
			}
		}
	}

	// 返回子矩阵最大值
	// 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
	query := func(r1, c1, r2, c2 int) int {
		r1 = max(r1, 0)
		c1 = max(c1, 0)
		r2 = min(r2, n)
		c2 = min(c2, m)
		k1 := bits.Len8(uint8(r2-r1)) - 1
		k2 := bits.Len8(uint8(c2-c1)) - 1
		// 视作四个子矩阵的并集
		return max(
			st[k1][k2][r1][c1],
			st[k1][k2][r2-1<<k1][c1],
			st[k1][k2][r1][c2-1<<k2],
			st[k1][k2][r2-1<<k1][c2-1<<k2],
		)
	}

	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && max(query(i-x, j-x+1, i+x+1, j+x), query(i-x+1, j-x, i+x, j+x+1)) <= x {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\log n\log m)$，其中 $n$ 和 $m$ 分别是 $\textit{matrix}$ 的行数和列数。瓶颈在预处理 ST 表上。
- 空间复杂度：$\mathcal{O}(nm\log n\log m)$。

## 方法二：线段树套 ST 表

当上下边界固定为 $\ell$ 和 $r$ 时，把每一列压缩成这一列的最大值，我们可以得到 $m$ 个数。

线段树每个节点（设对应区间为 $[\ell, r]$）保存的是这 $m$ 个数的一维 ST 表。

```py [sol-Python3]
class SparseTable:
    # 时间复杂度 O(n * log n)
    def __init__(self, a: list[int]):
        n = len(a)
        w = n.bit_length()
        st = [[0] * n for _ in range(w)]
        st[0] = a
        for i in range(1, w):
            for j in range(n - (1 << i) + 1):
                st[i][j] = max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))])
        self.st = st

    # [l, r) 左闭右开，下标从 0 开始
    # 返回 max(a[l:r])
    # 时间复杂度 O(1)
    def query(self, l: int, r: int) -> int:
        k = (r - l).bit_length() - 1
        return max(self.st[k][l], self.st[k][r - (1 << k)])


# 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree:
    def __init__(self, a: list[list[int]]):
        n = len(a)
        self.t = [None] * (2 << (n - 1).bit_length())
        self.build(a, 1, 0, n - 1)

    def build(self, a: list[list[int]], node: int, l: int, r: int):
        if l == r:  # 叶子
            self.t[node] = SparseTable(a[l])
            return

        m = (l + r) // 2
        self.build(a, node * 2, l, m)  # 初始化左子树
        self.build(a, node * 2 + 1, m + 1, r)  # 初始化右子树

        # 行号 [l, r] 中的每一列的最大值
        merged = [max(x, y) for x, y in zip(self.t[node * 2].st[0], self.t[node * 2 + 1].st[0])]
        self.t[node] = SparseTable(merged)

    def query(self, node: int, l: int, r: int, r1: int, r2: int, c1: int, c2: int) -> int:
        if r1 <= l and r <= r2:  # 当前子树完全在 [r1, r2] 内
            return self.t[node].query(c1, c2)
        m = (l + r) // 2
        if r2 <= m:  # [r1, r2] 在左子树
            return self.query(node * 2, l, m, r1, r2, c1, c2)
        if r1 > m:  # [r1, r2] 在右子树
            return self.query(node * 2 + 1, m + 1, r, r1, r2, c1, c2)
        return max(self.query(node * 2, l, m, r1, r2, c1, c2), self.query(node * 2 + 1, m + 1, r, r1, r2, c1, c2))


class Solution:
    def countLocalMaximums(self, matrix: list[list[int]]) -> int:
        n, m = len(matrix), len(matrix[0])
        # 线段树每个节点 [l, r] 保存的是，当上下边界固定为 l 和 r 时，把每一列的最大值视作一个 int，这 m 个数的一维 ST 表
        t = SegmentTree(matrix)

        ans = 0
        for i, row in enumerate(matrix):
            for j, x in enumerate(row):
                if x > 0 and max(t.query(1, 0, n - 1, max(i - x, 0), min(i + x, n - 1), max(j - x + 1, 0), min(j + x, m)),
                                 t.query(1, 0, n - 1, max(i - x + 1, 0), min(i + x - 1, n - 1), max(j - x, 0), min(j + x + 1, m))) <= x:
                    ans += 1
        return ans
```

```java [sol-Java]
class SparseTable {
    final int[][] st;

    // 时间复杂度 O(n * log n)
    public SparseTable(int[] nums) {
        int n = nums.length;
        int w = 32 - Integer.numberOfLeadingZeros(n);
        st = new int[w][n];
        st[0] = nums;
        for (int i = 1; i < w; i++) {
            for (int j = 0; j <= n - (1 << i); j++) {
                st[i][j] = Math.max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))]);
            }
        }
    }

    // [l, r) 左闭右开，下标从 0 开始
    // 时间复杂度 O(1)
    public int query(int l, int r) {
        int k = 31 - Integer.numberOfLeadingZeros(r - l);
        return Math.max(st[k][l], st[k][r - (1 << k)]);
    }
}

// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    private final SparseTable[] t;

    public SegmentTree(int[][] a) {
        int n = a.length;
        t = new SparseTable[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(a, 1, 0, n - 1);
    }

    private void build(int[][] a, int node, int l, int r) {
        if (l == r) { // 叶子
            t[node] = new SparseTable(a[l]);
            return;
        }

        int m = (l + r) / 2;
        build(a, node * 2, l, m);     // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树

        int[] merged = new int[a[0].length];
        for (int i = 0; i < merged.length; i++) {
            merged[i] = Math.max(t[node * 2].st[0][i], t[node * 2 + 1].st[0][i]); // 行号 [l, r] 中的第 i 列的最大值
        }
        t[node] = new SparseTable(merged);
    }

    // 行号闭区间 [r1, r2]，列号左闭右开 [c1, c2)
    public int query(int node, int l, int r, int r1, int r2, int c1, int c2) {
        if (r1 <= l && r <= r2) { // 当前子树完全在 [r1, r2] 内
            return t[node].query(c1, c2);
        }
        int m = (l + r) / 2;
        if (r2 <= m) { // [r1, r2] 在左子树
            return query(node * 2, l, m, r1, r2, c1, c2);
        }
        if (r1 > m) { // [r1, r2] 在右子树
            return query(node * 2 + 1, m + 1, r, r1, r2, c1, c2);
        }
        return Math.max(query(node * 2, l, m, r1, r2, c1, c2), query(node * 2 + 1, m + 1, r, r1, r2, c1, c2));
    }
}

class Solution {
    public int countLocalMaximums(int[][] matrix) {
        int n = matrix.length;
        int m = matrix[0].length;

        // 线段树每个节点 [l, r] 保存的是，当上下边界固定为 l 和 r 时，把每一列的最大值视作一个 int，这 m 个数的一维 ST 表
        SegmentTree t = new SegmentTree(matrix);

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int x = matrix[i][j];
                if (x == 0) {
                    continue;
                }
                int max1 = t.query(1, 0, n - 1, Math.max(i - x, 0), Math.min(i + x, n - 1), Math.max(j - x + 1, 0), Math.min(j + x, m));
                int max2 = t.query(1, 0, n - 1, Math.max(i - x + 1, 0), Math.min(i + x - 1, n - 1), Math.max(j - x, 0), Math.min(j + x + 1, m));
                if (Math.max(max1, max2) <= x) {
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class SparseTable {
public:
    vector<vector<int>> st;

    SparseTable() {}

    // 时间复杂度 O(n * log n)
    SparseTable(const vector<int>& a) {
        int n = a.size();
        int w = bit_width((uint32_t) n);
        st.resize(w, vector<int>(n));
        st[0] = a;
        for (int i = 1; i < w; i++) {
            for (int j = 0; j <= n - (1 << i); j++) {
                st[i][j] = max(st[i - 1][j], st[i - 1][j + (1 << (i - 1))]);
            }
        }
    }

    // [l, r) 左闭右开，下标从 0 开始
    // 时间复杂度 O(1)
    int query(int l, int r) const {
        int k = bit_width(1u * (r - l)) - 1;
        return max(st[k][l], st[k][r - (1 << k)]);
    }
};

// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
struct SegmentTree {
    vector<SparseTable> t;

    void build(const vector<vector<int>>& a, int node, int l, int r) {
        if (l == r) { // 叶子
            t[node] = SparseTable(a[l]);
            return;
        }

        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树

        vector<int> merged(a[0].size());
        for (int i = 0; i < merged.size(); i++) {
            merged[i] = max(t[node * 2].st[0][i], t[node * 2 + 1].st[0][i]); // 行号 [l, r] 中的第 i 列的最大值
        }
        t[node] = SparseTable(merged);
    }

public:
    SegmentTree(const vector<vector<int>>& a) : t(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, a.size() - 1);
    }

    // 行号闭区间 [r1, r2]，列号左闭右开 [c1, c2)
    int query(int node, int l, int r, int r1, int r2, int c1, int c2) const {
        if (r1 <= l && r <= r2) { // 当前子树完全在 [r1, r2] 内
            return t[node].query(c1, c2);
        }
        int m = (l + r) / 2;
        if (r2 <= m) { // [r1, r2] 在左子树
            return query(node * 2, l, m, r1, r2, c1, c2);
        }
        if (r1 > m) { // [r1, r2] 在右子树
            return query(node * 2 + 1, m + 1, r, r1, r2, c1, c2);
        }
        return max(query(node * 2, l, m, r1, r2, c1, c2), query(node * 2 + 1, m + 1, r, r1, r2, c1, c2));
    }
};

class Solution {
public:
    int countLocalMaximums(vector<vector<int>>& matrix) {
        int n = matrix.size(), m = matrix[0].size();
        // 线段树每个节点 [l, r] 保存的是，当上下边界固定为 l 和 r 时，把每一列的最大值视作一个 int，这 m 个数的一维 ST 表
        SegmentTree t(matrix);

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int x = matrix[i][j];
                if (x > 0 && max(t.query(1, 0, n - 1, max(i - x, 0), min(i + x, n - 1), max(j - x + 1, 0), min(j + x, m)),
                                 t.query(1, 0, n - 1, max(i - x + 1, 0), min(i + x - 1, n - 1), max(j - x, 0), min(j + x + 1, m))) <= x) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 一维 ST 表（泛型版本）
type sparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

func newSparseTable[T any](a []T, op func(T, T) T) sparseTable[T] {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	st[0] = a
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable[T]{st, op}
}

func (s sparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k])
}

// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
type seg []sparseTable[int]

func (t seg) build(a [][]int, node, l, r int) {
	if l == r { // 叶子
		t[node] = newSparseTable(a[l], func(a, b int) int { return max(a, b) })
		return
	}

	m := (l + r) / 2
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树

	merged := make([]int, len(a[0]))
	for i := range merged {
		merged[i] = max(t[node*2].st[0][i], t[node*2+1].st[0][i]) // 行号 [l, r] 中的第 i 列的最大值
	}
	t[node] = newSparseTable(merged, func(a, b int) int { return max(a, b) })
}

// 行号闭区间 [r1, r2]，列号左闭右开 [c1, c2)
func (t seg) query(node, l, r, r1, r2, c1, c2 int) int {
	if r1 <= l && r <= r2 { // 当前子树完全在 [r1, r2] 内
		return t[node].query(c1, c2)
	}
	m := (l + r) / 2
	if r2 <= m { // [r1, r2] 在左子树
		return t.query(node*2, l, m, r1, r2, c1, c2)
	}
	if r1 > m { // [r1, r2] 在右子树
		return t.query(node*2+1, m+1, r, r1, r2, c1, c2)
	}
	return max(t.query(node*2, l, m, r1, r2, c1, c2), t.query(node*2+1, m+1, r, r1, r2, c1, c2))
}

func countLocalMaximums(matrix [][]int) (ans int) {
	n, m := len(matrix), len(matrix[0])
	// 线段树每个节点 [l, r] 保存的是，当上下边界固定为 l 和 r 时，把每一列的最大值视作一个 int，这 m 个数的一维 ST 表
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(matrix, 1, 0, n-1)

	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && max(t.query(1, 0, n-1, max(i-x, 0), min(i+x, n-1), max(j-x+1, 0), min(j+x, m)),
				t.query(1, 0, n-1, max(i-x+1, 0), min(i+x-1, n-1), max(j-x, 0), min(j+x+1, m))) <= x {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm(\log n + \log m))$，其中 $n$ 和 $m$ 分别是 $\textit{matrix}$ 的行数和列数。线段树有 $\mathcal{O}(n)$ 个节点，每个节点用 $\mathcal{O}(m\log m)$ 的时间创建 ST 表。查询 $\mathcal{O}(nm)$ 次，每次 $\mathcal{O}(\log n)$ 时间。
- 空间复杂度：$\mathcal{O}(nm\log m)$。

## 专题训练

见下面数据结构题单的「**§8.7 ST 表**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
