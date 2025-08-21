## 方法一：暴力枚举

问题相当于把 $\textit{grid}$ 划分成三个区域，每个区域暴力找 [3195. 包含所有 1 的最小矩形面积 I](https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i/)。

有如下六种划分方案。

![w403d.png](https://pic.leetcode.cn/1719114413-gJmraG-w403d.png){:width=500px}

暴力枚举分割线的位置，把矩形划分成三个区域，每个区域用 [3195. 包含所有 1 的最小矩形面积 I](https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i/) 的方法暴力求解。

只需实现图中上面三种方案，下面三种方案可以通过把 $\textit{grid}$ 顺时针旋转 $90^\circ$ 得到，从而复用同一段代码逻辑。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1MZ421M74P/) 第四题，欢迎点赞关注~

```py [sol-Python3]
# 把矩阵 a 顺时针旋转 90°
def rotate(a: List[List[int]]) -> List[List[int]]:
    return list(zip(*reversed(a)))

class Solution:
    def minimumSum(self, grid: List[List[int]]) -> int:
        return min(self.solve(grid), self.solve(rotate(grid)))

    def solve(self, a: List[List[int]]) -> int:
        # 3195. 包含所有 1 的最小矩形面积 I
        # 限定在 a 的 [l,r) 列中
        def minimumArea(a: List[List[int]], l: int, r: int) -> int:
            left = top = inf
            right = bottom = 0
            for i, row in enumerate(a):
                for j, x in enumerate(row[l:r]):
                    if x:
                        left = min(left, j)
                        right = max(right, j)
                        top = min(top, i)
                        bottom = i
            return (right - left + 1) * (bottom - top + 1)

        ans = inf
        m, n = len(a), len(a[0])

        if m >= 3:
            for i in range(1, m):
                for j in range(i + 1, m):
                    # 图片上左
                    area = minimumArea(a[:i], 0, n)
                    area += minimumArea(a[i:j], 0, n)
                    area += minimumArea(a[j:], 0, n)
                    ans = min(ans, area)

        if m >= 2 and n >= 2:
            for i in range(1, m):
                for j in range(1, n):
                    # 图片上中
                    area = minimumArea(a[:i], 0, n)
                    area += minimumArea(a[i:], 0, j)
                    area += minimumArea(a[i:], j, n)
                    ans = min(ans, area)

                    # 图片上右
                    area = minimumArea(a[:i], 0, j)
                    area += minimumArea(a[:i], j, n)
                    area += minimumArea(a[i:], 0, n)
                    ans = min(ans, area)
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumSum(int[][] grid) {
        return Math.min(solve(grid), solve(rotate(grid)));
    }

    private int solve(int[][] a) {
        int m = a.length;
        int n = a[0].length;
        int ans = Integer.MAX_VALUE;

        if (m >= 3) {
            for (int i = 1; i < m; i++) {
                for (int j = i + 1; j < m; j++) {
                    // 图片上左
                    int area = minimumArea(a, 0, i, 0, n);
                    area += minimumArea(a, i, j, 0, n);
                    area += minimumArea(a, j, m, 0, n);
                    ans = Math.min(ans, area);
                }
            }
        }

        if (m >= 2 && n >= 2) {
            for (int i = 1; i < m; i++) {
                for (int j = 1; j < n; j++) {
                    // 图片上中
                    int area = minimumArea(a, 0, i, 0, n);
                    area += minimumArea(a, i, m, 0, j);
                    area += minimumArea(a, i, m, j, n);
                    ans = Math.min(ans, area);

                    // 图片上右
                    area = minimumArea(a, 0, i, 0, j);
                    area += minimumArea(a, 0, i, j, n);
                    area += minimumArea(a, i, m, 0, n);
                    ans = Math.min(ans, area);
                }
            }
        }

        return ans;
    }

    // 3195. 包含所有 1 的最小矩形面积 I
    // 限定在 a 的 [u,d) 行，[l,r) 列中
    private int minimumArea(int[][] a, int u, int d, int l, int r) {
        int left = r;
        int right = 0;
        int top = d;
        int bottom = 0;
        for (int i = u; i < d; i++) {
            for (int j = l; j < r; j++) {
                if (a[i][j] == 1) {
                    left = Math.min(left, j);
                    right = Math.max(right, j);
                    top = Math.min(top, i);
                    bottom = i;
                }
            }
        }
        return (right - left + 1) * (bottom - top + 1);
    }

    // 把矩阵 a 顺时针旋转 90°
    private int[][] rotate(int[][] a) {
        int m = a.length;
        int n = a[0].length;
        int[][] b = new int[n][m];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                b[j][m - 1 - i] = a[i][j];
            }
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 把矩阵 a 顺时针旋转 90°
    vector<vector<int>> rotate(const vector<vector<int>>& a) {
        int m = a.size(), n = a[0].size();
        vector b(n, vector<int>(m));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                b[j][m - 1 - i] = a[i][j];
            }
        }
        return b;
    }

    // 3195. 包含所有 1 的最小矩形面积 I
    // 限定在 a 的 [u,d) 行，[l,r) 列中
    int minimumArea(const vector<vector<int>>& a, int u, int d, int l, int r) {
        int left = r, right = 0;
        int top = d, bottom = 0;
        for (int i = u; i < d; i++) {
            for (int j = l; j < r; j++) {
                if (a[i][j] == 1) {
                    left = min(left, j);
                    right = max(right, j);
                    top = min(top, i);
                    bottom = i;
                }
            }
        }
        return (right - left + 1) * (bottom - top + 1);
    }

    int solve(const vector<vector<int>>& a) {
        int m = a.size(), n = a[0].size();
        int ans = INT_MAX;

        if (m >= 3) {
            for (int i = 1; i < m; i++) {
                for (int j = i + 1; j < m; j++) {
                    // 图片上左
                    int area = minimumArea(a, 0, i, 0, n);
                    area += minimumArea(a, i, j, 0, n);
                    area += minimumArea(a, j, m, 0, n);
                    ans = min(ans, area);
                }
            }
        }

        if (m >= 2 && n >= 2) {
            for (int i = 1; i < m; i++) {
                for (int j = 1; j < n; j++) {
                    // 图片上中
                    int area = minimumArea(a, 0, i, 0, n);
                    area += minimumArea(a, i, m, 0, j);
                    area += minimumArea(a, i, m, j, n);
                    ans = min(ans, area);

                    // 图片上右
                    area = minimumArea(a, 0, i, 0, j);
                    area += minimumArea(a, 0, i, j, n);
                    area += minimumArea(a, i, m, 0, n);
                    ans = min(ans, area);
                }
            }
        }

        return ans;
    }

public:
    int minimumSum(vector<vector<int>>& grid) {
        return min(solve(grid), solve(rotate(grid)));
    }
};
```

```go [sol-Go]
// 3195. 包含所有 1 的最小矩形面积 I
// 限定在 a 的 [l,r) 列中
func minimumArea(a [][]int, l, r int) int {
	left, right := r, 0
	top, bottom := len(a), 0
	for i, row := range a {
		for j, x := range row[l:r] {
			if x == 1 {
				left = min(left, j)
				right = max(right, j)
				top = min(top, i)
				bottom = i
			}
		}
	}
	return (right - left + 1) * (bottom - top + 1)
}

func minimumSum(grid [][]int) int {
	ans := math.MaxInt

	solve := func(a [][]int) {
		m, n := len(a), len(a[0])
		if m >= 3 {
			for i := 1; i < m; i++ {
				for j := i + 1; j < m; j++ {
					// 图片上左
					area := minimumArea(a[:i], 0, n)
					area += minimumArea(a[i:j], 0, n)
					area += minimumArea(a[j:], 0, n)
					ans = min(ans, area)
				}
			}
		}

		if m >= 2 && n >= 2 {
			for i := 1; i < m; i++ {
				for j := 1; j < n; j++ {
					// 图片上中
					area := minimumArea(a[:i], 0, n)
					area += minimumArea(a[i:], 0, j)
					area += minimumArea(a[i:], j, n)
					ans = min(ans, area)

					// 图片上右
					area = minimumArea(a[:i], 0, j)
					area += minimumArea(a[:i], j, n)
					area += minimumArea(a[i:], 0, n)
					ans = min(ans, area)
				}
			}
		}
	}

	solve(grid)
	solve(rotate(grid))
	return ans
}

// 把矩阵 a 顺时针旋转 90°
func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((mn)^2)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

## 方法二：动态规划

方法一过于暴力了，大量的时间都花在遍历矩阵元素上，同一个元素会被反复遍历 $\mathcal{O}(mn)$ 次。

有重复计算的地方，就有优化。

回想一下你之前写过的一些二维 DP，比如网格图 DP、[二维前缀和](https://leetcode.cn/problems/range-sum-query-2d-immutable/solution/tu-jie-yi-zhang-tu-miao-dong-er-wei-qian-84qp/) 等。这类 DP 的特点是，可以从 $\textit{grid}$ 的一个角（比如左上角）开始，递推得到**包含这个角的所有子矩形的信息**。

在六种划分方案中，右边四种每个区域都包含 $\textit{grid}$ 的一个角。尝试 DP。

![w403d.png](https://pic.leetcode.cn/1719114413-gJmraG-w403d.png){:width=500px}

定义 $f[i+1][j+1]$ 表示考虑「左上角为 $(0,0)$ 右下角为 $(i,j)$ 的子矩形」，包含这个子矩形中的所有 $1$ 的最小矩形面积。（我们算的是这个子矩形中的子矩形）

定义 $\textit{border}[i+1][j+1]$ 包含三个数，分别表示上述最小矩形的上边界、左边界和右边界。（下边界无需记录）

从上到下，从左到右遍历 $\textit{grid}$。设当前遍历到 $\textit{grid}[i]$ 这一行。在 $\textit{grid}[i][0]$ 到 $\textit{grid}[i][j]$ 中，设最左边的 $1$ 的列号为 $\textit{left}$（不存在则为 $-1$），最右边的 $1$ 的列号为 $\textit{right}$。注意 $\textit{left}$ 和 $\textit{right}$ 会随着我们遍历 $\textit{grid}[i]$ 这一行而发生变化。

分类讨论：

- 如果 $\textit{grid}[i]$ 这一行包含 $1$，且 $0$ 到 $i-1$ 行全为 $0$，那么 $f[i+1][j+1] = \textit{right}-\textit{left}+1,\ \textit{border}[i+1][j+1]=(i,\textit{left},\textit{right})$。
- 如果 $\textit{grid}[i]$ 这一行全为 $0$，那么包含「左上角为 $(0,0)$ 右下角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积，等于包含「左上角为 $(0,0)$ 右下角为 $(i-1,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积，即 $f[i+1][j+1] = f[i][j+1]$。同理有 $\textit{border}[i+1][j+1]=\textit{border}[i][j+1]$。
- 如果 $\textit{grid}[i]$ 这一行包含 $1$，且 $0$ 到 $i-1$ 行中也有 $1$，那么最小矩形：
  - 上边界是包含「左上角为 $(0,0)$ 右下角为 $(i-1,j)$ 的子矩形」中的所有 $1$ 的最小矩形的上边界，即 $\textit{border}[i][j+1]$ 的上边界，记作 $t$。
  - 左边界是 $\textit{border}[i][j+1]$ 的左边界与 $\textit{left}$ 的最小值，记作 $l$。
  - 右边界是 $\textit{border}[i][j+1]$ 的右边界与 $\textit{right}$ 的最大值，记作 $r$。
  - 下边界是 $i$。
  - 面积 $f[i+1][j+1] = (r - l + 1) \cdot (i - t + 1)$。
  - 边界 $\textit{border}[i+1][j+1]=(t,l,r)$。

代码实现时，$\textit{border}$ 可以用一个长为 $n$ 的数组滚动计算。

按照上述方法，预处理：

- 包含「左上角为 $(0,0)$ 右下角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。
- 包含「左下角为 $(m-1,0)$ 右上角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。
- 包含「右下角为 $(m-1,n-1)$ 左上角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。
- 包含「右上角为 $(0,n-1)$ 左下角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。

这样就能快速计算六种划分方案中的右边四种。

对于左边两种，中间区域的最小矩形面积怎么算？

优化一下暴力的过程即可：

- 首先预处理每行最左最右的 $1$ 的列号。
- 然后在枚举两条分割线的同时，维护中间区域的最左最右 $1$ 的列号（最左取最小值，最右取最大值），以及最上最下的 $1$ 的行号。

```py [sol-Python3]
def rotate(a: List[List[int]]) -> List[List[int]]:
    return list(zip(*reversed(a)))

class Solution:
    def minimumSum(self, grid: List[List[int]]) -> int:
        return min(self.solve(grid), self.solve(rotate(grid)))

    def solve(self, a: List[List[int]]) -> int:
        m, n = len(a), len(a[0])

        def minimumArea(a: List[List[int]]) -> List[List[int]]:
            m, n = len(a), len(a[0])
            # f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
            f = [[0] * (n + 1) for _ in range(m + 1)]
            border = [(-1, 0, 0)] * n
            for i, row in enumerate(a):
                left, right = -1, 0
                for j, x in enumerate(row):
                    if x:
                        if left < 0:
                            left = j
                        right = j
                    pre_top, pre_left, pre_right = border[j]
                    if left < 0:  # 这一排目前全是 0
                        f[i + 1][j + 1] = f[i][j + 1]  # 等于上面的结果
                    elif pre_top < 0:  # 这一排有 1，上面全是 0
                        f[i + 1][j + 1] = right - left + 1
                        border[j] = (i, left, right)
                    else:  # 这一排有 1，上面也有 1
                        l = min(pre_left, left)
                        r = max(pre_right, right)
                        f[i + 1][j + 1] = (r - l + 1) * (i - pre_top + 1)
                        border[j] = (pre_top, l, r)
            return f

        # 预处理每一行最左最右 1 的列号，用于中间区域最小矩形面积的计算
        lr = [None] * m
        for i in range(m):
            l, r = -1, 0
            for j in range(n):
                if a[i][j] > 0:
                    if l < 0:
                        l = j
                    r = j
            lr[i] = (l, r)

        # lt[i+1][j+1] = 包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        lt = minimumArea(a)
        a = rotate(a)
        # lb[i][j+1] = 包含【左下角为 (m-1,0) 右上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        lb = rotate(rotate(rotate(minimumArea(a))))
        a = rotate(a)
        # rb[i][j] = 包含【右下角为 (m-1,n-1) 左上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        rb = rotate(rotate(minimumArea(a)))
        a = rotate(a)
        # rt[i+1][j] = 包含【右上角为 (0,n-1) 左下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        rt = rotate(minimumArea(a))

        ans = inf
        if m >= 3:
            for i in range(1, m):
                left, right, top, bottom = n, 0, m, 0
                for j in range(i + 1, m):
                    l, r = lr[j - 1]
                    if l >= 0:
                        left = min(left, l)
                        right = max(right, r)
                        top = min(top, j - 1)
                        bottom = j - 1
                    # 图片上左
                    ans = min(ans, lt[i][n] + (right - left + 1) * (bottom - top + 1) + lb[j][n])

        if m >= 2 and n >= 2:
            for i in range(1, m):
                for j in range(1, n):
                    # 图片上中
                    ans = min(ans, lt[i][n] + lb[i][j] + rb[i][j])
                    # 图片上右
                    ans = min(ans, lt[i][j] + rt[i][j] + lb[i][n])
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumSum(int[][] grid) {
        return Math.min(solve(grid), solve(rotate(grid)));
    }

    private int solve(int[][] a) {
        int m = a.length;
        int n = a[0].length;
        
        // 预处理每一行最左最右 1 的列号，用于中间区域最小矩形面积的计算
        int[][] lr = new int[m][2];
        for (int i = 0; i < m; i++) {
            int l = -1;
            int r = 0;
            for (int j = 0; j < n; j++) {
                if (a[i][j] > 0) {
                    if (l < 0) {
                        l = j;
                    }
                    r = j;
                }
            }
            lr[i][0] = l;
            lr[i][1] = r;
        }

        // lt[i+1][j+1] = 包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        int[][] lt = minimumArea(a);
        a = rotate(a);
        // lb[i][j+1] = 包含【左下角为 (m-1,0) 右上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        int[][] lb = rotate(rotate(rotate(minimumArea(a))));
        a = rotate(a);
        // rb[i][j] = 包含【右下角为 (m-1,n-1) 左上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        int[][] rb = rotate(rotate(minimumArea(a)));
        a = rotate(a);
        // rt[i+1][j] = 包含【右上角为 (0,n-1) 左下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        int[][] rt = rotate(minimumArea(a));

        int ans = Integer.MAX_VALUE;
        if (m >= 3) {
            for (int i = 1; i < m; i++) {
                int left = n;
                int right = 0;
                int top = m;
                int bottom = 0;
                for (int j = i + 1; j < m; j++) {
                    int l = lr[j - 1][0];
                    if (l >= 0) {
                        left = Math.min(left, l);
                        right = Math.max(right, lr[j - 1][1]);
                        top = Math.min(top, j - 1);
                        bottom = j - 1;
                    }
                    // 图片上左
                    ans = Math.min(ans, lt[i][n] + (right - left + 1) * (bottom - top + 1) + lb[j][n]);
                }
            }
        }

        if (m >= 2 && n >= 2) {
            for (int i = 1; i < m; i++) {
                for (int j = 1; j < n; j++) {
                    // 图片上中
                    ans = Math.min(ans, lt[i][n] + lb[i][j] + rb[i][j]);
                    // 图片上右
                    ans = Math.min(ans, lt[i][j] + rt[i][j] + lb[i][n]);
                }
            }
        }
        return ans;
    }

    private int[][] minimumArea(int[][] a) {
        int m = a.length;
        int n = a[0].length;
        // f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        int[][] f = new int[m + 1][n + 1];
        int[][] border = new int[n][3];
        for (int j = 0; j < n; j++) {
            border[j][0] = -1;
        }
        for (int i = 0; i < m; i++) {
            int left = -1;
            int right = 0;
            for (int j = 0; j < n; j++) {
                if (a[i][j] == 1) {
                    if (left < 0) {
                        left = j;
                    }
                    right = j;
                }
                int[] preB = border[j];
                if (left < 0) { // 这一排目前全是 0
                    f[i + 1][j + 1] = f[i][j + 1]; // 等于上面的结果
                } else if (preB[0] < 0) { // 这一排有 1，上面全是 0
                    f[i + 1][j + 1] = right - left + 1;
                    border[j][0] = i;
                    border[j][1] = left;
                    border[j][2] = right;
                } else { // 这一排有 1，上面也有 1
                    int l = Math.min(preB[1], left);
                    int r = Math.max(preB[2], right);
                    f[i + 1][j + 1] = (r - l + 1) * (i - preB[0] + 1);
                    border[j][1] = l;
                    border[j][2] = r;
                }
            }
        }
        return f;
    }

    private int[][] rotate(int[][] a) {
        int m = a.length;
        int n = a[0].length;
        int[][] b = new int[n][m];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                b[j][m - 1 - i] = a[i][j];
            }
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<int>> rotate(const vector<vector<int>> a) {
        int m = a.size(), n = a[0].size();
        vector b(n, vector<int>(m));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                b[j][m - 1 - i] = a[i][j];
            }
        }
        return b;
    }

    vector<vector<int>> minimumArea(const vector<vector<int>>& a) {
        int m = a.size(), n = a[0].size();
        // f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        vector f(m + 1, vector<int>(n + 1));
        vector<tuple<int, int, int>> border(n + 1, {-1, -1, -1});
        for (int i = 0; i < m; i++) {
            int left = -1, right = 0;
            for (int j = 0; j < n; j++) {
                if (a[i][j]) {
                    if (left < 0) {
                        left = j;
                    }
                    right = j;
                }
                auto& [pre_top, pre_left, pre_right] = border[j];
                if (left < 0) { // 这一排目前全是 0
                    f[i + 1][j + 1] = f[i][j + 1]; // 等于上面的结果
                } else if (pre_top < 0) { // 这一排有 1，上面全是 0
                    f[i + 1][j + 1] = right - left + 1;
                    border[j] = {i, left, right};
                } else { // 这一排有 1，上面也有 1
                    int l = min(pre_left, left), r = max(pre_right, right);
                    f[i + 1][j + 1] = (r - l + 1) * (i - pre_top + 1);
                    border[j] = {pre_top, l, r};
                }
            }
        }
        return f;
    }

    int solve(vector<vector<int>>& a) {
        int m = a.size(), n = a[0].size();
        
        // 预处理每一行最左最右 1 的列号，用于中间区域最小矩形面积的计算
        vector<pair<int, int>> lr(m);
        for (int i = 0; i < m; i++) {
            int l = -1, r = 0;
            for (int j = 0; j < n; j++) {
                if (a[i][j] > 0) {
                    if (l < 0) {
                        l = j;
                    }
                    r = j;
                }
            }
            lr[i] = {l, r};
        }

        // lt[i+1][j+1] = 包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        vector<vector<int>> lt = minimumArea(a);
        a = rotate(a);
        // lb[i][j+1] = 包含【左下角为 (m-1,0) 右上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        vector<vector<int>> lb = rotate(rotate(rotate(minimumArea(a))));
        a = rotate(a);
        // rb[i][j] = 包含【右下角为 (m-1,n-1) 左上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        vector<vector<int>> rb = rotate(rotate(minimumArea(a)));
        a = rotate(a);
        // rt[i+1][j] = 包含【右上角为 (0,n-1) 左下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
        vector<vector<int>> rt = rotate(minimumArea(a));

        int ans = INT_MAX;
        if (m >= 3) {
            for (int i = 1; i < m; i++) {
                int left = n, right = 0, top = m, bottom = 0;
                for (int j = i + 1; j < m; j++) {
                    if (auto& [l, r] = lr[j - 1]; l >= 0) {
                        left = min(left, l);
                        right = max(right, r);
                        top = min(top, j - 1);
                        bottom = j - 1;
                    }
                    // 图片上左
                    ans = min(ans, lt[i][n] + (right - left + 1) * (bottom - top + 1) + lb[j][n]);
                }
            }
        }

        if (m >= 2 && n >= 2) {
            for (int i = 1; i < m; i++) {
                for (int j = 1; j < n; j++) {
                    // 图片上中
                    ans = min(ans, lt[i][n] + lb[i][j] + rb[i][j]);
                    // 图片上右
                    ans = min(ans, lt[i][j] + rt[i][j] + lb[i][n]);
                }
            }
        }
        return ans;
    }

public:
    int minimumSum(vector<vector<int>>& grid) {
        auto a = rotate(grid);
        return min(solve(grid), solve(a));
    }
};
```

```go [sol-Go]
func minimumArea(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	// f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	type data struct{ top, left, right int }
	border := make([]data, n)
	for j := range border {
		border[j].top = -1 // 无
	}

	for i, row := range a {
		left, right := -1, 0
		for j, x := range row {
			if x > 0 {
				if left < 0 {
					left = j
				}
				right = j
			}
			preB := border[j]
			if left < 0 { // 这一排目前全是 0
				f[i+1][j+1] = f[i][j+1] // 等于上面的结果
			} else if preB.top < 0 { // 这一排有 1，上面全是 0
				f[i+1][j+1] = right - left + 1
				border[j] = data{i, left, right}
			} else { // 这一排有 1，上面也有 1
				l, r := min(preB.left, left), max(preB.right, right)
				f[i+1][j+1] = (r - l + 1) * (i - preB.top + 1)
				border[j] = data{preB.top, l, r}
			}
		}
	}
	return f
}

func minimumSum(grid [][]int) int {
	ans := math.MaxInt

	solve := func(a [][]int) {
		m, n := len(a), len(a[0])

		// 预处理每一行最左最右 1 的列号，用于中间区域最小矩形面积的计算
		type pair struct{ l, r int }
		lr := make([]pair, m)
		for i, row := range a {
			l, r := -1, 0
			for j, x := range row {
				if x > 0 {
					if l < 0 {
						l = j
					}
					r = j
				}
			}
			lr[i] = pair{l, r}
		}

		// lt[i+1][j+1] = 包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		lt := minimumArea(a)
		a = rotate(a)
		// lb[i][j+1] = 包含【左下角为 (m-1,0) 右上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		lb := rotate(rotate(rotate(minimumArea(a))))
		a = rotate(a)
		// rb[i][j] = 包含【右下角为 (m-1,n-1) 左上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		rb := rotate(rotate(minimumArea(a)))
		a = rotate(a)
		// rt[i+1][j] = 包含【右上角为 (0,n-1) 左下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		rt := rotate(minimumArea(a))

		if m >= 3 {
			for i := 1; i < m; i++ {
				left, right, top, bottom := n, 0, m, 0
				for j := i + 1; j < m; j++ {
					if p := lr[j-1]; p.l >= 0 {
						left = min(left, p.l)
						right = max(right, p.r)
						top = min(top, j-1)
						bottom = j - 1
					}
					// 图片上左
					area := lt[i][n] // minimumArea(a[:i], 0, n)
					area += (right - left + 1) * (bottom - top + 1) // minimumArea(a[i:j], 0, n)
					area += lb[j][n] // minimumArea(a[j:], 0, n)
					ans = min(ans, area)
				}
			}
		}

		if m >= 2 && n >= 2 {
			for i := 1; i < m; i++ {
				for j := 1; j < n; j++ {
					// 图片上中
					area := lt[i][n] // minimumArea(a[:i], 0, n)
					area += lb[i][j] // minimumArea(a[i:], 0, j)
					area += rb[i][j] // minimumArea(a[i:], j, n)
					ans = min(ans, area)
					// 图片上右
					area = lt[i][j]  // minimumArea(a[:i], 0, j)
					area += rt[i][j] // minimumArea(a[:i], j, n)
					area += lb[i][n] // minimumArea(a[i:], 0, n)
					ans = min(ans, area)
				}
			}
		}
	}

	solve(grid)
	solve(rotate(grid))
	return ans
}

func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

## 思考题

把题目中的 $3$ 改成 $4$ 呢？

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
