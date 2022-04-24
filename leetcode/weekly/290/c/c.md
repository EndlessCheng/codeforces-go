基本思路就是对每个点 $(x_i,y_i)$，统计横坐标不小于 $x_i$ 且纵坐标不小于 $y_i$ 的矩形个数。

为了加快统计速度，我们可以将 $\textit{rectangles}$ 和 $\textit{points}$ 都按照纵坐标从大到小排序。

然后遍历每个点 $(x_i,y_i)$，将所有纵坐标不小于 $y_i$ 的矩形的横坐标加入一个有序列表 $\textit{xs}$。

**由于纵坐标的范围只有 $[1,100]$，我们可以暴力地在每次插入完横坐标后对 $\textit{xs}$ 排序，排序的次数不会超过 $100$ 次。**

然后在 $\textit{xs}$ 中二分即可算出横坐标不小于 $x_i$ 的矩形个数，由于我们是按纵坐标从大到小遍历的，因此这些矩形的纵坐标均不小于 $y_i$，因此这些矩形均包含点 $(x_i,y_i)$。

```Python [sol1-Python3]
class Solution:
    def countRectangles(self, rectangles: List[List[int]], points: List[List[int]]) -> List[int]:
        rectangles.sort(key=lambda r: -r[1])
        n = len(points)
        ans = [0] * n
        i, xs = 0, []
        for (x, y), id in sorted(zip(points, range(n)), key=lambda x: -x[0][1]):
            start = i
            while i < len(rectangles) and y <= rectangles[i][1]:
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
            while (i < rectangles.length && points[id][1] <= rectangles[i][1])
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
            if (xs.get(mid) >= x)
                right = mid;
            else
                left = mid + 1;
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
            while (i < rectangles.size() && points[id][1] <= rectangles[i][1])
                xs.push_back(rectangles[i++][0]);
            if (start < i) sort(xs.begin(), xs.end()); // 只有在 xs 插入了新元素时才排序
            ans[id] = i - (lower_bound(xs.begin(), xs.end(), points[id][0]) - xs.begin());
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
		for ; i < n && p[1] <= rectangles[i][1]; i++ {
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

```python
from sortedcontainers import SortedList

class Solution:
    def countRectangles(self, rectangles: List[List[int]], points: List[List[int]]) -> List[int]:
        rectangles.sort(key=lambda r: -r[1])
        n = len(points)
        ans = [0] * n
        i, xs = 0, SortedList()
        for (x, y), id in sorted(zip(points, range(n)), key=lambda x: -x[0][1]):
            while i < len(rectangles) and y <= rectangles[i][1]:
                xs.add(rectangles[i][0])
                i += 1
            ans[id] = i - xs.bisect_left(x)
        return ans
```
