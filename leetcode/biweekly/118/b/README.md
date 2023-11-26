[本题视频讲解](https://www.bilibili.com/video/BV1Rw411P72r/)

考虑最大矩形面积，然后再考虑正方形的面积。

矩形面积是长和宽的乘积，长、宽可以分别求出来。

以 $\textit{hBars}$ 为例：

- 如果不做任何移除，那么最长长度为 $1$。
- 如果移除一条线，那么最长长度为 $2$。
- 如果移除两条编号相邻的线，那么最长长度为 $3$。
- 如果移除三条编号连续的线（例如 $2,3,4$），那么最长长度为 $4$。
- 依此类推。

所以把数组排序后，求出最长连续递增长度即可。

这可以用分组循环求出，具体请看：[【简单题杀手】分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)

求出后，正方形的边长是长宽的最小值，其平方即为正方形的面积。

```py [sol-Python3]
class Solution:
    def maximizeSquareHoleArea(self, n: int, m: int, hBars: List[int], vBars: List[int]) -> int:
        def f(a: List[int]) -> int:
            a.sort()
            n = len(a)
            mx = i = 0
            while i < n:
                st = i
                i += 1
                while i < n and a[i] - a[i - 1] == 1:
                    i += 1
                mx = max(mx, i - st)
            return mx + 1
        return min(f(hBars), f(vBars)) ** 2
```

```java [sol-Java]
class Solution {
    public int maximizeSquareHoleArea(int n, int m, int[] hBars, int[] vBars) {
        int size = Math.min(f(hBars), f(vBars));
        return size * size;
    }

    private int f(int[] a) {
        Arrays.sort(a);
        int n = a.length;
        int mx = 0, i = 0;
        while (i < n) {
            int st = i;
            for (i++; i < n && a[i] - a[i - 1] == 1; i++) ;
            mx = Math.max(mx, i - st + 1);
        }
        return mx;
    }
}
```

```cpp [sol-C++]
class Solution {
    int f(vector<int> &a) {
        sort(a.begin(), a.end());
        int n = a.size();
        int mx = 0, i = 0;
        while (i < n) {
            int st = i;
            for (i++; i < n && a[i] - a[i - 1] == 1; i++);
            mx = max(mx, i - st + 1);
        }
        return mx;
    }

public:
    int maximizeSquareHoleArea(int, int, vector<int> &hBars, vector<int> &vBars) {
        int size = min(f(hBars), f(vBars));
        return size * size;
    }
};
```

```go [sol-Go]
func f(a []int) (mx int) {
	slices.Sort(a)
	for i, n := 0, len(a); i < n; {
		st := i
		for i++; i < n && a[i]-a[i-1] == 1; i++ {}
		mx = max(mx, i-st+1)
	}
	return
}

func maximizeSquareHoleArea(_, _ int, hBars, vBars []int) int {
	mn := min(f(hBars), f(vBars))
	return mn * mn
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m)$，其中 $n$ 为 $\textit{hBars}$ 的长度，$m$ 为 $\textit{vBars}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 相似题目

- [1465. 切割后面积最大的蛋糕](https://leetcode.cn/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/)
