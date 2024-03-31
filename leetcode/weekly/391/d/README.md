![w391-c.png](https://pic.leetcode.cn/1711856129-vydCEa-w391-c.png)

这两种投影长度，其中较大者为曼哈顿距离（较小者是两段折线的投影长度之差，不合法），即如下恒等式

$$
|x_1 - x_2| + |y_1 - y_2| = \max(|x_1'-x_2'|,|y_1'-y_2'|)
$$

其中等式左侧为 $(x_1,y_1)$ 和 $(x_2,y_2)$ 的**曼哈顿距离**，等式右侧 $(x',y') = (x+y,y-x)$，计算的是 $(x_1',y_1')$ 和 $(x_2',y_2')$ 两点的曼哈顿距离投影到 $x'$ 轴和 $y'$ 轴的线段长度的最大值，即**切比雪夫距离**。

所以要求任意两点曼哈顿距离的最大值，根据上面的恒等式，我们只需要计算任意两个 $(x',y')$ 切比雪夫距离的最大值，即横纵坐标差的最大值

$$
\max(\max(x') - \min(x'), \max(y') - \min(y'))
$$

附：[视频讲解](https://www.bilibili.com/video/BV1fq421A7CY/) 第四题，欢迎点赞关注！

## 方法一：有序集合

本题可以删除一个点，我们可以枚举要删除的点，并用有序集合维护其它 $n-1$ 个点的 $x'$ 和 $y'$ 的最大值和最小值。

```py [sol-Python3]
from sortedcontainers import SortedList

class Solution:
    def minimumDistance(self, points: List[List[int]]) -> int:
        xs = SortedList()
        ys = SortedList()
        for x, y in points:
            xs.add(x + y)
            ys.add(y - x)
        ans = inf
        for x, y in points:
            x, y = x + y, y - x
            xs.remove(x)
            ys.remove(y)
            ans = min(ans, max(xs[-1] - xs[0], ys[-1] - ys[0]))
            xs.add(x)
            ys.add(y)
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumDistance(int[][] points) {
        TreeMap<Integer, Integer> xs = new TreeMap<>();
        TreeMap<Integer, Integer> ys = new TreeMap<>();
        for (int[] p : points) {
            xs.merge(p[0] + p[1], 1, Integer::sum);
            ys.merge(p[1] - p[0], 1, Integer::sum);
        }
        int ans = Integer.MAX_VALUE;
        for (int[] p : points) {
            int x = p[0] + p[1], y = p[1] - p[0];
            if (xs.get(x) == 1) xs.remove(x);
            else xs.merge(x, -1, Integer::sum);
            if (ys.get(y) == 1) ys.remove(y);
            else ys.merge(y, -1, Integer::sum);
            ans = Math.min(ans, Math.max(xs.lastKey() - xs.firstKey(), ys.lastKey() - ys.firstKey()));
            xs.merge(x, 1, Integer::sum);
            ys.merge(y, 1, Integer::sum);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumDistance(vector<vector<int>> &points) {
        multiset<int> xs, ys;
        for (auto &p : points) {
            xs.insert(p[0] + p[1]);
            ys.insert(p[1] - p[0]);
        }
        int ans = INT_MAX;
        for (auto &p : points) {
            int x = p[0] + p[1], y = p[1] - p[0];
            xs.erase(xs.find(x));
            ys.erase(ys.find(y));
            ans = min(ans, max(*xs.rbegin() - *xs.begin(), *ys.rbegin() - *ys.begin()));
            xs.insert(x);
            ys.insert(y);
        }
        return ans;
    }
};
```

```go [sol-Go]
// https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha
func minimumDistance(points [][]int) int {
	xs := redblacktree.New[int, int]()
	ys := redblacktree.New[int, int]()
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		put(xs, x)
		put(ys, y)
	}
	ans := math.MaxInt
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		remove(xs, x)
		remove(ys, y)
		ans = min(ans, max(xs.Right().Key-xs.Left().Key, ys.Right().Key-ys.Left().Key))
		put(xs, x)
		put(ys, y)
	}
	return ans
}

func put(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	t.Put(v, c+1)
}

func remove(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	if c == 1 {
		t.Remove(v)
	} else {
		t.Put(v, c-1)
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：维护最大次大、最小次小

优化：如果把最大的 $x'$ 删掉，那么次大 $x'$ 就作为剩下 $n-1$ 个 $x'$ 中的最大值了，对于最小值也同理。

所以我们可以维护 $x'$ 和 $y'$ 的最大次大、最小次小，一共 $8$ 个数。

```py [sol-Python3]
class Solution:
    def minimumDistance(self, points: List[List[int]]) -> int:
        max_x1, max_x2 = nlargest(2, (x + y for x, y in points))   # x 最大次大
        min_x1, min_x2 = nsmallest(2, (x + y for x, y in points))  # x 最小次小
        max_y1, max_y2 = nlargest(2, (y - x for x, y in points))   # y 最大次大
        min_y1, min_y2 = nsmallest(2, (y - x for x, y in points))  # y 最小次小

        ans = inf
        for x, y in points:
            x, y = x + y, y - x
            dx = (max_x2 if x == max_x1 else max_x1) - (min_x2 if x == min_x1 else min_x1)
            dy = (max_y2 if y == max_y1 else max_y1) - (min_y2 if y == min_y1 else min_y1)
            ans = min(ans, max(dx, dy))
        return ans
```

```py [sol-Python3 手动维护]
class Solution:
    def minimumDistance(self, points: List[List[int]]) -> int:
        max_x1 = max_x2 = max_y1 = max_y2 = -inf
        min_x1 = min_x2 = min_y1 = min_y2 = inf

        for x, y in points:
            x, y = x + y, y - x

            # x 最大次大
            if x > max_x1:
                max_x2 = max_x1
                max_x1 = x
            elif x > max_x2:
                max_x2 = x

            # x 最小次小
            if x < min_x1:
                min_x2 = min_x1
                min_x1 = x
            elif x < min_x2:
                min_x2 = x

            # y 最大次大
            if y > max_y1:
                max_y2 = max_y1
                max_y1 = y
            elif y > max_y2:
                max_y2 = y

            # y 最小次小
            if y < min_y1:
                min_y2 = min_y1
                min_y1 = y
            elif y < min_y2:
                min_y2 = y

        ans = inf
        for x, y in points:
            x, y = x + y, y - x
            dx = (max_x2 if x == max_x1 else max_x1) - (min_x2 if x == min_x1 else min_x1)
            dy = (max_y2 if y == max_y1 else max_y1) - (min_y2 if y == min_y1 else min_y1)
            ans = min(ans, max(dx, dy))
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumDistance(int[][] points) {
        final int INF = Integer.MAX_VALUE;
        int maxX1 = -INF, maxX2 = -INF, maxY1 = -INF, maxY2 = -INF;
        int minX1 = INF, minX2 = INF, minY1 = INF, minY2 = INF;

        for (int[] p : points) {
            int x = p[0] + p[1];
            int y = p[1] - p[0];

            // x 最大次大
            if (x > maxX1) {
                maxX2 = maxX1;
                maxX1 = x;
            } else if (x > maxX2) {
                maxX2 = x;
            }

            // x 最小次小
            if (x < minX1) {
                minX2 = minX1;
                minX1 = x;
            } else if (x < minX2) {
                minX2 = x;
            }

            // y 最大次大
            if (y > maxY1) {
                maxY2 = maxY1;
                maxY1 = y;
            } else if (y > maxY2) {
                maxY2 = y;
            }

            // y 最小次小
            if (y < minY1) {
                minY2 = minY1;
                minY1 = y;
            } else if (y < minY2) {
                minY2 = y;
            }
        }

        int ans = INF;
        for (int[] p : points) {
            int x = p[0] + p[1];
            int y = p[1] - p[0];
            int dx = (x == maxX1 ? maxX2 : maxX1) - (x == minX1 ? minX2 : minX1);
            int dy = (y == maxY1 ? maxY2 : maxY1) - (y == minY1 ? minY2 : minY1);
            ans = Math.min(ans, Math.max(dx, dy));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 更新最大次大
    void update_max(int v, int &max1, int &max2) {
        if (v > max1) {
            max2 = max1;
            max1 = v;
        } else if (v > max2) {
            max2 = v;
        }
    }

    // 更新最小次小
    void update_min(int v, int &min1, int &min2) {
        if (v < min1) {
            min2 = min1;
            min1 = v;
        } else if (v < min2) {
            min2 = v;
        }
    }

public:
    int minimumDistance(vector<vector<int>> &points) {
        int max_x1 = INT_MIN, max_x2 = INT_MIN, max_y1 = INT_MIN, max_y2 = INT_MIN;
        int min_x1 = INT_MAX, min_x2 = INT_MAX, min_y1 = INT_MAX, min_y2 = INT_MAX;

        for (auto &p : points) {
            int x = p[0] + p[1];
            int y = p[1] - p[0];
            update_max(x, max_x1, max_x2);
            update_min(x, min_x1, min_x2);
            update_max(y, max_y1, max_y2);
            update_min(y, min_y1, min_y2);
        }

        int ans = INT_MAX;
        for (auto &p : points) {
            int x = p[0] + p[1];
            int y = p[1] - p[0];
            int dx = (x == max_x1 ? max_x2 : max_x1) - (x == min_x1 ? min_x2 : min_x1);
            int dy = (y == max_y1 ? max_y2 : max_y1) - (y == min_y1 ? min_y2 : min_y1);
            ans = min(ans, max(dx, dy));
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumDistance(points [][]int) int {
	const inf = math.MaxInt
	maxX1, maxX2, maxY1, maxY2 := -inf, -inf, -inf, -inf
	minX1, minX2, minY1, minY2 := inf, inf, inf, inf

	for _, p := range points {
		x := p[0] + p[1]
		y := p[1] - p[0]

		// x 最大次大
		if x > maxX1 {
			maxX2 = maxX1
			maxX1 = x
		} else if x > maxX2 {
			maxX2 = x
		}

		// x 最小次小
		if x < minX1 {
			minX2 = minX1
			minX1 = x
		} else if x < minX2 {
			minX2 = x
		}

		// y 最大次大
		if y > maxY1 {
			maxY2 = maxY1
			maxY1 = y
		} else if y > maxY2 {
			maxY2 = y
		}

		// y 最小次小
		if y < minY1 {
			minY2 = minY1
			minY1 = y
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := inf
	for _, p := range points {
		x := p[0] + p[1]
		y := p[1] - p[0]
		dx := f(x, maxX1, maxX2) - f(x, minX1, minX2)
		dy := f(y, maxY1, maxY2) - f(y, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(v, v1, v2 int) int {
	if v == v1 {
		return v2
	}
	return v1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 思考题

如果题目求的是**欧氏距离**，要怎么做？

> 关键词：凸包、旋转卡壳。

## 相似题目

- [1330. 翻转子数组得到最大的数组值](https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/)
- [1131. 绝对值表达式的最大值](https://leetcode.cn/problems/maximum-of-absolute-value-expression/)

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
