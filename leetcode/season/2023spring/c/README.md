### 本题视频讲解

见[【力扣杯2023春·个人赛】](https://www.bilibili.com/video/BV1dg4y1j78A/)第三题。

### 思路

1. 统计所有左下和右上坐标，由于会出现 $0.5$，可以将坐标乘 $2$。
2. 离散化横纵坐标。
3. 二维差分，具体见视频讲解。
4. 用二维前缀和复原，计算最大值。

```py [sol1-Python3]
class Solution:
    def fieldOfGreatestBlessing(self, forceField: List[List[int]]) -> int:
        # 1. 统计所有左下和右上坐标
        x_set = set()
        y_set = set()
        for i, j, side in forceField:
            x_set.add(2 * i - side)
            x_set.add(2 * i + side)
            y_set.add(2 * j - side)
            y_set.add(2 * j + side)

        # 2. 离散化
        xs = sorted(x_set)
        ys = sorted(y_set)
        n, m = len(xs), len(ys)

        # 3. 二维差分：快速地把一个矩形范围内的数都 +1
        diff = [[0] * (m + 2) for _ in range(n + 2)]
        for i, j, side in forceField:
            r1 = bisect_left(xs, 2 * i - side)
            r2 = bisect_left(xs, 2 * i + side)
            c1 = bisect_left(ys, 2 * j - side)
            c2 = bisect_left(ys, 2 * j + side)
            # 将区域 r1<=r<=r2 && c1<=c<=c2 上的数都加上 x
            # 多 +1 是为了方便求后面用二维前缀和复原
            diff[r1 + 1][c1 + 1] += 1
            diff[r1 + 1][c2 + 2] -= 1
            diff[r2 + 2][c1 + 1] -= 1
            diff[r2 + 2][c2 + 2] += 1

        # 4. 直接在 diff 上复原（二维前缀和），计算最大值
        ans = 0
        for i in range(1, n + 1):
            for j in range(1, m + 1):
                diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1]
                ans = max(ans, diff[i][j])
        return ans
```

```java [sol1-Java]
class Solution {
    public int fieldOfGreatestBlessing(int[][] forceField) {
        // 1. 统计所有左下和右上坐标
        int nf = forceField.length, k = 0;
        long[] xs = new long[nf * 2], ys = new long[nf * 2];
        for (var f : forceField) {
            long i = f[0], j = f[1], side = f[2];
            xs[k] = 2 * i - side;
            xs[k + 1] = 2 * i + side;
            ys[k++] = 2 * j - side;
            ys[k++] = 2 * j + side;
        }

        // 2. 排序去重
        xs = unique(xs);
        ys = unique(ys);

        // 3. 二维差分
        int n = xs.length, m = ys.length;
        var diff = new int[n + 2][m + 2];
        for (var f : forceField) {
            long i = f[0], j = f[1], side = f[2];
            int r1 = Arrays.binarySearch(xs, 2 * i - side);
            int r2 = Arrays.binarySearch(xs, 2 * i + side);
            int c1 = Arrays.binarySearch(ys, 2 * j - side);
            int c2 = Arrays.binarySearch(ys, 2 * j + side);
            // 将区域 r1<=r<=r2 && c1<=c<=c2 上的数都加上 x
            // 多 +1 是为了方便求后面复原
            ++diff[r1 + 1][c1 + 1];
            --diff[r1 + 1][c2 + 2];
            --diff[r2 + 2][c1 + 1];
            ++diff[r2 + 2][c2 + 2];
        }

        // 4. 直接在 diff 上复原，计算最大值
        int ans = 0;
        for (int i = 1; i <= n; ++i) {
            for (int j = 1; j <= m; ++j) {
                diff[i][j] += diff[i - 1][j] + diff[i][j - 1] - diff[i - 1][j - 1];
                ans = Math.max(ans, diff[i][j]);
            }
        }
        return ans;
    }

    private long[] unique(long[] a) {
        Arrays.sort(a);
        int k = 0;
        for (int i = 1; i < a.length; i++)
            if (a[k] != a[i])
                a[++k] = a[i];
        return Arrays.copyOf(a, k + 1);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int fieldOfGreatestBlessing(vector<vector<int>> &forceField) {
        // 1. 统计所有左下和右上坐标
        vector<long long> xs, ys;
        for (auto &f: forceField) {
            long long i = f[0], j = f[1], side = f[2];
            xs.push_back(2 * i - side);
            xs.push_back(2 * i + side);
            ys.push_back(2 * j - side);
            ys.push_back(2 * j + side);
        }

        // 2. 排序去重
        sort(xs.begin(), xs.end());
        xs.erase(unique(xs.begin(), xs.end()), xs.end());
        sort(ys.begin(), ys.end());
        ys.erase(unique(ys.begin(), ys.end()), ys.end());

        // 3. 二维差分
        int n = xs.size(), m = ys.size(), diff[n + 2][m + 2];
        memset(diff, 0, sizeof(diff));
        for (auto &f: forceField) {
            long long i = f[0], j = f[1], side = f[2];
            int r1 = lower_bound(xs.begin(), xs.end(), 2 * i - side) - xs.begin();
            int r2 = lower_bound(xs.begin(), xs.end(), 2 * i + side) - xs.begin();
            int c1 = lower_bound(ys.begin(), ys.end(), 2 * j - side) - ys.begin();
            int c2 = lower_bound(ys.begin(), ys.end(), 2 * j + side) - ys.begin();
            // 将区域 r1<=r<=r2 && c1<=c<=c2 上的数都加上 x
            // 多 +1 是为了方便求后面复原
            ++diff[r1 + 1][c1 + 1];
            --diff[r1 + 1][c2 + 2];
            --diff[r2 + 2][c1 + 1];
            ++diff[r2 + 2][c2 + 2];
        }

        // 4. 直接在 diff 上复原，计算最大值
        int ans = 0;
        for (int i = 1; i <= n; ++i) {
            for (int j = 1; j <= m; ++j) {
                diff[i][j] += diff[i - 1][j] + diff[i][j - 1] - diff[i - 1][j - 1];
                ans = max(ans, diff[i][j]);
            }
        }
        return ans;
    }
};
```

```go [sol1-Golang]
func fieldOfGreatestBlessing(forceField [][]int) (ans int) {
	// 1. 统计所有左下和右上坐标
	var xs, ys []int
	for _, f := range forceField {
		i, j, side := f[0], f[1], f[2]
		xs = append(xs, 2*i-side, 2*i+side)
		ys = append(ys, 2*j-side, 2*j+side)
	}

	// 2. 排序去重
	unique := func(a []int) []int {
		sort.Ints(a)
		k := 0
		for _, x := range a[1:] {
			if a[k] != x {
				k++
				a[k] = x
			}
		}
		return a[:k+1]
	}
	xs = unique(xs)
	ys = unique(ys)

	// 3. 二维差分
	n, m := len(xs), len(ys)
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, m+2)
	}
	for _, f := range forceField {
		i, j, side := f[0], f[1], f[2]
		r1 := sort.SearchInts(xs, 2*i-side)
		r2 := sort.SearchInts(xs, 2*i+side)
		c1 := sort.SearchInts(ys, 2*j-side)
		c2 := sort.SearchInts(ys, 2*j+side)
		// 将区域 r1<=r<=r2 && c1<=c<=c2 上的数都加上 x
		// 多 +1 是为了方便求后面复原
		diff[r1+1][c1+1]++
		diff[r1+1][c2+2]--
		diff[r2+2][c1+1]--
		diff[r2+2][c2+2]++
	}

	// 4. 直接在 diff 上复原，计算最大值
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
			ans = max(ans, diff[i][j])
		}
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{forceField}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。
