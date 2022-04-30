### 方法一：按纵坐标排序

基本思路就是对每个点 $(x_i,y_i)$，统计横坐标不小于 $x_i$ 且纵坐标不小于 $y_i$ 的矩形个数。

为了加快统计速度，我们可以将 $\textit{rectangles}$ 和 $\textit{points}$ 都按照纵坐标从大到小排序。

然后遍历每个点 $(x_i,y_i)$，将所有纵坐标不小于 $y_i$ 的矩形的横坐标加入一个有序列表 $\textit{xs}$。

**由于纵坐标的范围只有 $[1,100]$，我们可以暴力地在每次插入完横坐标后对 $\textit{xs}$ 排序，排序的次数不会超过 $100$ 次。**

然后在 $\textit{xs}$ 中二分即可算出横坐标不小于 $x_i$ 的矩形个数，由于我们是按纵坐标从大到小遍历的，因此这些矩形的纵坐标均不小于 $y_i$，因此这些矩形均包含点 $(x_i,y_i)$。

- 时间复杂度：$O(Hn\log n + m\log m + m\log n)$，其中 $H=\max(h_i)$，$n$ 为 $\textit{rectangles}$ 的长度，$m$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$O(n+m)$。

```Python [sol1-Python3]
class Solution:
    def countRectangles(self, rectangles: List[List[int]], points: List[List[int]]) -> List[int]:
        rectangles.sort(key=lambda r: -r[1])
        n = len(points)
        ans = [0] * n
        i, xs = 0, []
        for (x, y), id in sorted(zip(points, range(n)), key=lambda x: -x[0][1]):
            start = i
            while i < len(rectangles) and rectangles[i][1] >= y:
                xs.append(rectangles[i][0])
                i += 1
            if start < i:
                xs.sort()  # 只有在 xs 插入了新元素时才排序
            ans[id] = i - bisect_left(xs, x)
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] countRectangles(int[][] rectangles, int[][] points) {
        Arrays.sort(rectangles, (a, b) -> b[1] - a[1]);

        var n = points.length;
        var ids = IntStream.range(0, n).boxed().toArray(Integer[]::new);
        Arrays.sort(ids, (i, j) -> points[j][1] - points[i][1]);

        var ans = new int[n];
        var xs = new ArrayList<Integer>();
        var i = 0;
        for (var id : ids) {
            var start = i;
            while (i < rectangles.length && rectangles[i][1] >= points[id][1])
                xs.add(rectangles[i++][0]);
            if (start < i) Collections.sort(xs); // 只有在 xs 插入了新元素时才排序
            ans[id] = i - lowerBound(xs, points[id][0]);
        }
        return ans;
    }

    int lowerBound(List<Integer> xs, int x) {
        int left = 0, right = xs.size();
        while (left < right) {
            var mid = (left + right) / 2;
            if (xs.get(mid) >= x) right = mid;
            else left = mid + 1;
        }
        return left;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    vector<int> countRectangles(vector<vector<int>> &rectangles, vector<vector<int>> &points) {
        sort(rectangles.begin(), rectangles.end(), [](auto &a, auto &b) { return a[1] > b[1]; });

        int n = points.size();
        vector<int> ids(n);
        iota(ids.begin(), ids.end(), 0);
        sort(ids.begin(), ids.end(), [&](int i, int j) { return points[i][1] > points[j][1]; });

        vector<int> ans(n), xs;
        int i = 0;
        for (int id : ids) {
            int start = i;
            while (i < rectangles.size() && rectangles[i][1] >= points[id][1])
                xs.push_back(rectangles[i++][0]);
            if (start < i) sort(xs.begin(), xs.end()); // 只有在 xs 插入了新元素时才排序
            ans[id] = xs.end() - lower_bound(xs.begin(), xs.end(), points[id][0]);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i][1] > rectangles[j][1] })
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] > points[j][1] })

	ans := make([]int, len(points))
	i, n, xs := 0, len(rectangles), []int{}
	for _, p := range points {
		start := i
		for ; i < n && rectangles[i][1] >= p[1]; i++ {
			xs = append(xs, rectangles[i][0])
		}
		if start < i { // 只有在 xs 插入了新元素时才排序
			sort.Ints(xs)
		}
		ans[p[2]] = i - sort.SearchInts(xs, p[0])
	}
	return ans
}
```

注：如果这题纵坐标的范围也是 $10^9$，我们还可以用名次树来做出此题（如 Python 的 `SortedList`）。

- 时间复杂度：$O(n\log n + m\log m+m\log n)$。这种做法就与 $H$ 无关了。
- 空间复杂度：$O(n+m)$。

```python
from sortedcontainers import SortedList

class Solution:
    def countRectangles(self, rectangles: List[List[int]], points: List[List[int]]) -> List[int]:
        rectangles.sort(key=lambda r: -r[1])
        n = len(points)
        ans = [0] * n
        i, xs = 0, SortedList()
        for (x, y), id in sorted(zip(points, range(n)), key=lambda x: -x[0][1]):
            while i < len(rectangles) and rectangles[i][1] >= y:
                xs.add(rectangles[i][0])
                i += 1
            ans[id] = i - xs.bisect_left(x)
        return ans
```

### 方法二：按横坐标排序

我们也可以按横坐标从大到小排序，对于点 $(x_i,y_i)$，统计横坐标不小于 $x_i$ 的矩形个数，由于高度不超过 $100$，可以用一个数组来存储每个高度有多少个矩形。

那么只要累加高度不小于 $y_i$ 的矩形个数即可。

实现时可以暴力累加，也可以用树状数组，由于这里高度很小，代码直接用的暴力累加的写法。

- 时间复杂度：$O(Hm + n\log n + m\log m)$。
- 空间复杂度：$O(H+m)$。

```Python [sol1-Python3]
class Solution:
    def countRectangles(self, rectangles: List[List[int]], points: List[List[int]]) -> List[int]:
        rectangles.sort(key=lambda r: -r[0])
        n = len(points)
        ans = [0] * n
        cnt = [0] * (max(y for _, y in rectangles) + 1)
        i = 0
        for (x, y), id in sorted(zip(points, range(n)), key=lambda x: -x[0][0]):
            while i < len(rectangles) and rectangles[i][0] >= x:
                cnt[rectangles[i][1]] += 1
                i += 1
            ans[id] = sum(cnt[y:])
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] countRectangles(int[][] rectangles, int[][] points) {
        Arrays.sort(rectangles, (a, b) -> b[0] - a[0]);

        var n = points.length;
        var ids = IntStream.range(0, n).boxed().toArray(Integer[]::new);
        Arrays.sort(ids, (i, j) -> points[j][0] - points[i][0]);

        var ans = new int[n];
        var cnt = new int[101];
        var i = 0;
        for (var id : ids) {
            while (i < rectangles.length && rectangles[i][0] >= points[id][0])
                ++cnt[rectangles[i++][1]];
            for (var j = points[id][1]; j < cnt.length; ++j)
                ans[id] += cnt[j];
        }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    vector<int> countRectangles(vector<vector<int>> &rectangles, vector<vector<int>> &points) {
        sort(rectangles.begin(), rectangles.end(), [](auto &a, auto &b) { return a[0] > b[0]; });

        int n = points.size();
        vector<int> ids(n);
        iota(ids.begin(), ids.end(), 0);
        sort(ids.begin(), ids.end(), [&](int i, int j) { return points[i][0] > points[j][0]; });

        vector<int> ans(n), cnt(101);
        int i = 0;
        for (int id : ids) {
            while (i < rectangles.size() && rectangles[i][0] >= points[id][0])
                ++cnt[rectangles[i++][1]];
            ans[id] = accumulate(cnt.begin() + points[id][1], cnt.end(), 0);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i][0] > rectangles[j][0] })
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool { return points[i][0] > points[j][0] })

	ans := make([]int, len(points))
	i, cnt := 0, [101]int{}
	for _, p := range points {
		for ; i < len(rectangles) && rectangles[i][0] >= p[0]; i++ {
			cnt[rectangles[i][1]]++
		}
		for _, c := range cnt[p[1]:] {
			ans[p[2]] += c
		}
	}
	return ans
}
```

### 方法三：按行统计 + 二分查找

由于至多有 $100$ 行数据，我们可以统计这 $100$ 行的矩阵的横坐标坐标，这样对于每个点 $(x_i,y_i)$，从第 $y_i$ 行遍历到第 $100$ 行，对于每一行，二分求出有多少个矩阵的横坐标不小于 $x_i$。累加即为答案。

- 时间复杂度：$O((n+Hm)\log n)$。
- 空间复杂度：$O(n+H)$。不计返回值的空间。

```Python [sol3-Python3]
class Solution:
    def countRectangles(self, rectangles: List[List[int]], points: List[List[int]]) -> List[int]:
        max_y = max(y for _, y in rectangles)
        xs = [[] for i in range(max_y + 1)]
        for x, y in rectangles:
            xs[y].append(x)
        for x in xs:
            x.sort()
        return [sum(len(x) - bisect_left(x, px) for x in xs[py:]) for px, py in points]
```

```java [sol3-Java]
class Solution {
    public int[] countRectangles(int[][] rectangles, int[][] points) {
        List<Integer>[] xs = new List[101];
        Arrays.setAll(xs, e -> new ArrayList<>());
        for (var r : rectangles) xs[r[1]].add(r[0]);
        for (var x : xs) Collections.sort(x);

        var n = points.length;
        var ans = new int[n];
        for (var i = 0; i < n; ++i)
            for (var j = points[i][1]; j <= 100; j++)
                ans[i] += xs[j].size() - lowerBound(xs[j], points[i][0]);
        return ans;
    }

    int lowerBound(List<Integer> xs, int x) {
        int left = 0, right = xs.size();
        while (left < right) {
            var mid = (left + right) / 2;
            if (xs.get(mid) >= x) right = mid;
            else left = mid + 1;
        }
        return left;
    }
}
```

```C++ [sol3-C++]
class Solution {
public:
    vector<int> countRectangles(vector<vector<int>> &rectangles, vector<vector<int>> &points) {
        vector<int> xs[101];
        for (auto &r: rectangles)
            xs[r[1]].push_back(r[0]);
        for (auto &x: xs)
            sort(x.begin(), x.end());

        int n = points.size();
        vector<int> ans(n);
        for (int i = 0; i < n; ++i)
            for (int y = points[i][1]; y <= 100; ++y) {
                auto &x = xs[y];
                ans[i] += x.end() - lower_bound(x.begin(), x.end(), points[i][0]);
            }
        return ans;
    }
};
```

```go [sol3-Go]
func countRectangles(rectangles [][]int, points [][]int) []int {
	xs := [101][]int{}
	for _, r := range rectangles {
		xs[r[1]] = append(xs[r[1]], r[0])
	}
	for _, x := range xs {
		sort.Ints(x)
	}

	ans := make([]int, len(points))
	for i, p := range points {
		for _, x := range xs[p[1]:] {
			ans[i] += len(x) - sort.SearchInts(x, p[0])
		}
	}
	return ans
}
```
