## 问题转化

最大化最小值，考虑二分答案，即二分距离的**下界** $\textit{low}$。为什么？因为 $\textit{low}$ 越大，可以选的点越少，有单调性。

![lc3464.png](https://pic.leetcode.cn/1740305801-CRFYtB-lc3464.png)

把正方形拉成一条线，示例 2 按照左边界、上边界、右边界、下边界的顺时针顺序，这 $5$ 个点在一维上的坐标为

$$
a=[0,3,4,5,6]
$$

现在问题变成：

- 能否在数组 $a$ 中选 $k$ 个数，要求任意两个相邻元素相差至少为 $\textit{low}$，且第一个数和最后一个数相差至多为 $\textit{side}\cdot 4 - \textit{low}$。
- $\textit{side}\cdot 4 - \textit{low}$ 是因为 $a$ 是个环形数组，设第一个点为 $x$，最后一个点为 $y$，那么 $y$ 可以视作负方向上的 $y-\textit{side}\cdot 4$，我们要求 $x-(y-\textit{side}\cdot 4) \ge \textit{low}$，解得 $y-x\le \textit{side}\cdot 4 - \textit{low}$。

## 方法一

枚举第一个数，不断向后二分找相距至少为 $\textit{low}$ 的最近元素，直到找到 $k$ 个数，或者第一个数和最后一个数相差超过 $\textit{side}\cdot 4 - \textit{low}$ 时停止。

⚠**注意**：本题保证 $k\ge 4$，所以答案不会超过 $\textit{side}$。这也保证了如果下一个点不在正方形的当前边或者下一条边上，那么距离是一定满足要求的，所以「二分找下一个点」的做法是正确的。而 $k\le 3$ 时，答案可能会超过 $\textit{side}$，此时「二分找下一个点」的做法是错误的。

⚠**注意**：不需要找一圈后又绕回到数组 $a$ 的开头继续找。设 $\textit{start}$ 是第一个点，$p$ 是二分找到的最后一个点（绕回到数组开头找到的 $p$）。因为我们要求首尾两个点相距 $\ge \textit{low}$，那么把 $p$ 当作第一个点开始往后搜，下一个点是 $\textit{start}$ 或者 $\textit{start}$ 前面的点，所以相比之下，绕回数组开头是没有意义的。这也同时意味着，无需把环形数组 $a$ 复制一份。

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$1$。一定可以满足要求。
- 开区间右端点初始值：$\textit{side} + 1$。一定无法满足要求。

```py [sol-Python3]
class Solution:
    def maxDistance(self, side: int, points: list[list[int]], k: int) -> int:
        # 正方形边上的点，按照顺时针映射到一维数轴上
        a = []
        for x, y in points:
            if x == 0:
                a.append(y)
            elif y == side:
                a.append(side + x)
            elif x == side:
                a.append(side * 3 - y)
            else:
                a.append(side * 4 - x)
        a.sort()

        def check(low: int) -> bool:
            for start in a:  # 枚举第一个点
                end = start + side * 4 - low
                cur = start
                for _ in range(k - 1):  # 还需要找 k-1 个点
                    j = bisect_left(a, cur + low)
                    if j == len(a) or a[j] > end:  # 不能离第一个点太近
                        break
                    cur = a[j]
                else:
                    return True
            return False

        # 本题保证 k >= 4，所以最远距离不会超过 side
        left, right = 1, side + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```py [sol-Python3 库函数]
class Solution:
    def maxDistance(self, side: int, points: list[list[int]], k: int) -> int:
        # 正方形边上的点，按照顺时针映射到一维数轴上
        a = []
        for x, y in points:
            if x == 0:
                a.append(y)
            elif y == side:
                a.append(side + x)
            elif x == side:
                a.append(side * 3 - y)
            else:
                a.append(side * 4 - x)
        a.sort()

        def check(low: int) -> bool:
            # 如果 low+1 不满足要求，但 low 满足要求，那么答案就是 low
            low += 1
            for start in a:  # 枚举第一个点
                end = start + side * 4 - low
                cur = start
                for _ in range(k - 1):  # 还需要找 k-1 个点
                    j = bisect_left(a, cur + low)
                    if j == len(a) or a[j] > end:  # 不能离第一个点太近
                        break
                    cur = a[j]
                else:
                    return False
            return True

        # 本题保证 k >= 4，所以最远距离不会超过 side
        return bisect_left(range(side), True, key=check)
```

```java [sol-Java]
class Solution {
    public int maxDistance(int side, int[][] points, int k) {
        // 正方形边上的点，按照顺时针映射到一维数轴上
        long[] a = new long[points.length];
        for (int i = 0; i < points.length; i++) {
            int x = points[i][0];
            int y = points[i][1];
            if (x == 0) {
                a[i] = y;
            } else if (y == side) {
                a[i] = side + x;
            } else if (x == side) {
                a[i] = side * 3L - y;
            } else {
                a[i] = side * 4L - x;
            }
        }
        Arrays.sort(a);

        // 本题保证 k >= 4，所以最远距离不会超过 side
        int left = 1;
        int right = side + 1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(a, side, k, mid)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(long[] a, int side, int k, int low) {
        next:
        for (long start : a) { // 枚举第一个点
            long end = start + side * 4L - low;
            long cur = start;
            for (int i = 0; i < k - 1; i++) { // 还需要找 k-1 个点
                int j = lowerBound(a, cur + low);
                if (j == a.length || a[j] > end) { // 不能离第一个点太近
                    continue next;
                }
                cur = a[j];
            }
            return true;
        }
        return false;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(long[] nums, long target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (nums[mid] >= target) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistance(int side, vector<vector<int>>& points, int k) {
        // 正方形边上的点，按照顺时针映射到一维数轴上
        vector<long long> a;
        for (auto& p : points) {
            int x = p[0], y = p[1];
            if (x == 0) {
                a.push_back(y);
            } else if (y == side) {
                a.push_back(side + x);
            } else if (x == side) {
                a.push_back(side * 3LL - y);
            } else {
                a.push_back(side * 4LL - x);
            }
        }
        ranges::sort(a);

        auto check = [&](int low) -> bool {
            for (long long start : a) { // 枚举第一个点
                long long end = start + side * 4LL - low;
                long long cur = start;
                for (int i = 0; i < k - 1; i++) { // 还需要找 k-1 个点
                    auto it = ranges::lower_bound(a, cur + low);
                    if (it == a.end() || *it > end) { // 不能离第一个点太近
                        cur = -1;
                        break;
                    }
                    cur = *it;
                }
                if (cur >= 0) {
                    return true;
                }
            }
            return false;
        };

        // 本题保证 k >= 4，所以最远距离不会超过 side
        int left = 1, right = side + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maxDistance(side int, points [][]int, k int) int {
	// 正方形边上的点，按照顺时针映射到一维数轴上
	a := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)

	// 本题保证 k >= 4，所以最远距离不会超过 side
	ans := sort.Search(side, func(low int) bool {
		// 如果 low+1 不满足要求，但 low 满足要求，那么答案就是 low
		low++
	next:
		for i, start := range a { // 枚举第一个点
			cur := start
			for range k - 1 { // 还需要找 k-1 个点
				i += sort.Search(len(a)-i, func(j int) bool { return a[i+j] >= cur+low })
				if i == len(a) || a[i] > start+side*4-low { // 不能离第一个点太近
					continue next
				}
				cur = a[i]
			}
			return false
		}
		return true
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk \log \textit{side}\log n)$，其中 $n$ 是 $\textit{points}$ 的长度。由于中途会退出循环，这个复杂度是跑不满的。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：k 个同向指针

把方法一最内层的二分查找，改用 $k$ 个指针维护。

一开始，初始化一个长为 $k$ 的 $\textit{idx}$ 数组，初始值 $\textit{idx}[j]=0$。

然后写个 $k$ 指针（双指针的推广）：

- 遍历 $j=1,2,3,\ldots,k-1$，如果发现 $a[\textit{idx}[j]] < a[\textit{idx}[j-1]] + \textit{low}$，就不断把 $\textit{idx}[j]$ 加一直到不满足条件。如果 $\textit{idx}[j]=n$ 则返回。
- 这些指针移动后，如果首尾两个指针指向的数相差不超过 $\textit{side}\cdot 4 - \textit{low}$，则返回。
- 否则把 $\textit{idx}[0]$ 加一，继续循环。

### 优化前

```py [sol-Python3]
class Solution:
    def maxDistance(self, side: int, points: list[list[int]], k: int) -> int:
        a = []
        for x, y in points:
            if x == 0:
                a.append(y)
            elif y == side:
                a.append(side + x)
            elif x == side:
                a.append(side * 3 - y)
            else:
                a.append(side * 4 - x)
        a.sort()

        def check(low: int) -> bool:
            idx = [0] * k
            while True:
                for j in range(1, k):
                    while a[idx[j]] < a[idx[j - 1]] + low:
                        idx[j] += 1
                        if idx[j] == len(a):
                            return False
                if a[idx[-1]] - a[idx[0]] <= side * 4 - low:
                    return True
                idx[0] += 1

        left, right = 1, side + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```java [sol-Java]
class Solution {
    public int maxDistance(int side, int[][] points, int k) {
        long[] a = new long[points.length];
        for (int i = 0; i < points.length; i++) {
            int x = points[i][0];
            int y = points[i][1];
            if (x == 0) {
                a[i] = y;
            } else if (y == side) {
                a[i] = side + x;
            } else if (x == side) {
                a[i] = side * 3L - y;
            } else {
                a[i] = side * 4L - x;
            }
        }
        Arrays.sort(a);

        int left = 1;
        int right = side + 1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(a, side, k, mid)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(long[] a, int side, int k, int low) {
        int[] idx = new int[k];
        while (true) {
            for (int j = 1; j < k; j++) {
                while (a[idx[j]] < a[idx[j - 1]] + low) {
                    idx[j]++;
                    if (idx[j] == a.length) {
                        return false;
                    }
                }
            }
            if (a[idx[k - 1]] - a[idx[0]] <= side * 4L - low) {
                return true;
            }
            idx[0]++;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistance(int side, vector<vector<int>>& points, int k) {
        // 正方形边上的点，按照顺时针映射到一维数轴上
        vector<long long> a;
        for (auto& p : points) {
            int x = p[0], y = p[1];
            if (x == 0) {
                a.push_back(y);
            } else if (y == side) {
                a.push_back(side + x);
            } else if (x == side) {
                a.push_back(side * 3LL - y);
            } else {
                a.push_back(side * 4LL - x);
            }
        }
        ranges::sort(a);

        auto check = [&](int low) -> bool {
            vector<int> idx(k);
            while (true) {
                for (int j = 1; j < k; j++) {
                    while (a[idx[j]] < a[idx[j - 1]] + low) {
                        idx[j]++;
                        if (idx[j] == a.size()) {
                            return false;
                        }
                    }
                }
                if (a[idx[k - 1]] - a[idx[0]] <= side * 4LL - low) {
                    return true;
                }
                idx[0]++;
            }
        };

        // 本题保证 k >= 4，所以最远距离不会超过 side
        int left = 1, right = side + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maxDistance(side int, points [][]int, k int) int {
	a := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)

	ans := sort.Search(side, func(low int) bool {
		low++
		idx := make([]int, k)
		for {
			for j := 1; j < k; j++ {
				for a[idx[j]] < a[idx[j-1]]+low {
					idx[j]++
					if idx[j] == len(a) {
						return true
					}
				}
			}
			if a[idx[k-1]]-a[idx[0]] <= side*4-low {
				return false
			}
			idx[0]++
		}
	})
	return ans
}
```

### 优化

把从 $\textit{start}=a[0]$ 开始向后二分得到的 $k$ 个下标，记到 $\textit{idx}$ 数组中。如果没有 $k$ 个下标，直接返回。

这样初始化比从 $0$ 开始一个一个地向后移动指针更快。

此外，第一个指针至多移动到第二个指针的初始位置，就不用继续枚举了，后面必然无法得到符合要求的结果。

```py [sol-Python3]
class Solution:
    def maxDistance(self, side: int, points: list[list[int]], k: int) -> int:
        a = []
        for x, y in points:
            if x == 0:
                a.append(y)
            elif y == side:
                a.append(side + x)
            elif x == side:
                a.append(side * 3 - y)
            else:
                a.append(side * 4 - x)
        a.sort()

        def check(low: int) -> bool:
            idx = [0] * k
            cur = a[0]
            for j in range(1, k):
                i = bisect_left(a, cur + low)
                if i == len(a):
                    return False
                idx[j] = i
                cur = a[i]
            if cur - a[0] <= side * 4 - low:
                return True

            # 第一个指针移动到第二个指针的位置，就不用继续枚举了
            for idx[0] in range(1, idx[1]):
                for j in range(1, k):
                    while a[idx[j]] < a[idx[j - 1]] + low:
                        idx[j] += 1
                        if idx[j] == len(a):
                            return False
                if a[idx[-1]] - a[idx[0]] <= side * 4 - low:
                    return True
            return False

        left, right = 1, side + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```java [sol-Java]
class Solution {
    public int maxDistance(int side, int[][] points, int k) {
        long[] a = new long[points.length];
        for (int i = 0; i < points.length; i++) {
            int x = points[i][0];
            int y = points[i][1];
            if (x == 0) {
                a[i] = y;
            } else if (y == side) {
                a[i] = side + x;
            } else if (x == side) {
                a[i] = side * 3L - y;
            } else {
                a[i] = side * 4L - x;
            }
        }
        Arrays.sort(a);

        int left = 1;
        int right = side + 1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(a, side, k, mid)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(long[] a, int side, int k, int low) {
        int[] idx = new int[k];
        long cur = a[0];
        for (int j = 1; j < k; j++) {
            int i = lowerBound(a, cur + low);
            if (i == a.length) {
                return false;
            }
            idx[j] = i;
            cur = a[i];
        }
        if (cur - a[0] <= side * 4L - low) {
            return true;
        }

        // 第一个指针移动到第二个指针的位置，就不用继续枚举了
        int end0 = idx[1];
        for (idx[0] = 1; idx[0] < end0; idx[0]++) {
            for (int j = 1; j < k; j++) {
                while (a[idx[j]] < a[idx[j - 1]] + low) {
                    idx[j]++;
                    if (idx[j] == a.length) {
                        return false;
                    }
                }
            }
            if (a[idx[k - 1]] - a[idx[0]] <= side * 4L - low) {
                return true;
            }
        }
        return false;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(long[] nums, long target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (nums[mid] >= target) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistance(int side, vector<vector<int>>& points, int k) {
        vector<long long> a;
        for (auto& p : points) {
            int x = p[0], y = p[1];
            if (x == 0) {
                a.push_back(y);
            } else if (y == side) {
                a.push_back(side + x);
            } else if (x == side) {
                a.push_back(side * 3LL - y);
            } else {
                a.push_back(side * 4LL - x);
            }
        }
        ranges::sort(a);

        auto check = [&](int low) -> bool {
            vector<int> idx(k);
            long long cur = a[0];
            for (int j = 1; j < k; j++) {
                int i = ranges::lower_bound(a, cur + low) - a.begin();
                if (i == a.size()) {
                    return false;
                }
                idx[j] = i;
                cur = a[i];
            }
            if (cur - a[0] <= side * 4LL - low) {
                return true;
            }

            // 第一个指针移动到第二个指针的位置，就不用继续枚举了
            int end0 = idx[1];
            for (idx[0]++; idx[0] < end0; idx[0]++) {
                for (int j = 1; j < k; j++) {
                    while (a[idx[j]] < a[idx[j - 1]] + low) {
                        idx[j]++;
                        if (idx[j] == a.size()) {
                            return false;
                        }
                    }
                }
                if (a[idx[k - 1]] - a[idx[0]] <= side * 4LL - low) {
                    return true;
                }
            }
            return false;
        };

        int left = 1, right = side + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func maxDistance(side int, points [][]int, k int) int {
	a := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)

	ans := sort.Search(side, func(low int) bool {
		low++
		idx := make([]int, k)
		cur := a[0]
		for j, i := 1, 0; j < k; j++ {
			i += sort.Search(len(a)-i, func(j int) bool { return a[i+j] >= cur+low })
			if i == len(a) {
				return true
			}
			idx[j] = i
			cur = a[i]
		}
		if cur-a[0] <= side*4-low {
			return false
		}

		// 第一个指针移动到第二个指针的位置，就不用继续枚举了
		end0 := idx[1]
		for idx[0]++; idx[0] < end0; idx[0]++ {
			for j := 1; j < k; j++ {
				for a[idx[j]] < a[idx[j-1]]+low {
					idx[j]++
					if idx[j] == len(a) {
						return true
					}
				}
			}
			if a[idx[k-1]]-a[idx[0]] <= side*4-low {
				return false
			}
		}
		return true
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + nk \log \textit{side})$，其中 $n$ 是 $\textit{points}$ 的长度。其中 $\mathcal{O}(n\log n)$ 是排序的时间复杂度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. 【本题相关】[二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
