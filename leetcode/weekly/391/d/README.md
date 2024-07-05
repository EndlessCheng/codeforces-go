![minimize-manhattan-distances.png](https://pic.leetcode.cn/1719906297-iZvmWo-minimize-manhattan-distances.png)

这两种投影长度，其中较大者为曼哈顿距离（较小者是两段折线的投影长度**之差**，不合法），即如下恒等式

$$
|x_1 - x_2| + |y_1 - y_2| = \max(|x_1'-x_2'|,|y_1'-y_2'|)
$$

其中等式左侧为 $(x_1,y_1)$ 和 $(x_2,y_2)$ 的**曼哈顿距离**，等式右侧 $(x',y') = (x+y,y-x)$，计算的是 $(x_1',y_1')$ 和 $(x_2',y_2')$ 两点的曼哈顿距离投影到 $x$ 轴和 $y$ 轴的线段长度的最大值，即**切比雪夫距离**。

所以要求任意两点曼哈顿距离的最大值，根据上面的恒等式，我们只需要计算任意两个 $(x',y')$ 切比雪夫距离的最大值，即横纵坐标差的最大值

$$
\max\{\max(x') - \min(x'), \max(y') - \min(y')\}
$$

请看 [视频讲解](https://www.bilibili.com/video/BV1fq421A7CY/) 第四题，欢迎点赞关注！

## 方法一：有序集合

枚举要移除的点，用两个**有序集合**维护其他 $n-1$ 个点的 $x'$ 和 $y'$，用 $\max\{\max(x') - \min(x'), \max(y') - \min(y')\}$ 更新答案的最大值。

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
            int x = p[0] + p[1];
            int y = p[1] - p[0];
            if (xs.get(x) == 1) xs.remove(x);
            else xs.merge(x, -1, Integer::sum); // 移除一个 x
            if (ys.get(y) == 1) ys.remove(y);
            else ys.merge(y, -1, Integer::sum); // 移除一个 y

            int dx = xs.lastKey() - xs.firstKey();
            int dy = ys.lastKey() - ys.firstKey();
            ans = Math.min(ans, Math.max(dx, dy));

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
    int minimumDistance(vector<vector<int>>& points) {
        multiset<int> xs, ys;
        for (auto& p : points) {
            xs.insert(p[0] + p[1]);
            ys.insert(p[1] - p[0]);
        }

        int ans = INT_MAX;
        for (auto& p : points) {
            int x = p[0] + p[1], y = p[1] - p[0];
            xs.erase(xs.find(x)); // 移除一个 x
            ys.erase(ys.find(y)); // 移除一个 y

            int dx = *xs.rbegin() - *xs.begin();
            int dy = *ys.rbegin() - *ys.begin();
            ans = min(ans, max(dx, dy));

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
		remove(xs, x) // 移除一个 x
		remove(ys, y) // 移除一个 y
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

优化：如果把最大的 $x'$ 移除，那么次大的 $x'$ 就是剩下 $n-1$ 个 $x'$ 中的最大值了，对于最小值也同理。

所以只需维护 $x'$ 和 $y'$ 的最大次大、最小次小，一共 $8$ 个数。

注意最大次大可能是相同的，最小次小可能是相同的。

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
    void update_max(int v, int& max1, int& max2) { // 注意这里是引用
        if (v > max1) {
            max2 = max1;
            max1 = v;
        } else if (v > max2) {
            max2 = v;
        }
    }

    // 更新最小次小
    void update_min(int v, int& min1, int& min2) { // 注意这里是引用
        if (v < min1) {
            min2 = min1;
            min1 = v;
        } else if (v < min2) {
            min2 = v;
        }
    }

public:
    int minimumDistance(vector<vector<int>>& points) {
        int max_x1 = INT_MIN, max_x2 = INT_MIN, max_y1 = INT_MIN, max_y2 = INT_MIN;
        int min_x1 = INT_MAX, min_x2 = INT_MAX, min_y1 = INT_MAX, min_y2 = INT_MAX;

        for (auto& p : points) {
            int x = p[0] + p[1];
            int y = p[1] - p[0];
            update_max(x, max_x1, max_x2);
            update_min(x, min_x1, min_x2);
            update_max(y, max_y1, max_y2);
            update_min(y, min_y1, min_y2);
        }

        int ans = INT_MAX;
        for (auto& p : points) {
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

进一步地，要移除的点只能是 $x'$ 或 $y'$ 最大最小的点（不然移除前后都一样），所以额外维护最大最小值的**下标**，一共 $4$ 个，最后只需遍历这 $4$ 个坐标，而不是遍历整个 $\textit{points}$ 数组。

```py [sol-Python3]
class Solution:
    def minimumDistance(self, points: List[List[int]]) -> int:
        max_x1 = max_x2 = max_y1 = max_y2 = -inf
        min_x1 = min_x2 = min_y1 = min_y2 = inf
        max_xi = min_xi = max_yi = min_yi = 0

        for i, (x, y) in enumerate(points):
            x, y = x + y, y - x

            # x 最大次大
            if x > max_x1:
                max_x2 = max_x1
                max_x1 = x
                max_xi = i
            elif x > max_x2:
                max_x2 = x

            # x 最小次小
            if x < min_x1:
                min_x2 = min_x1
                min_x1 = x
                min_xi = i
            elif x < min_x2:
                min_x2 = x

            # y 最大次大
            if y > max_y1:
                max_y2 = max_y1
                max_y1 = y
                max_yi = i
            elif y > max_y2:
                max_y2 = y

            # y 最小次小
            if y < min_y1:
                min_y2 = min_y1
                min_y1 = y
                min_yi = i
            elif y < min_y2:
                min_y2 = y

        ans = inf
        for i in max_xi, min_xi, max_yi, min_yi:
            dx = (max_x2 if i == max_xi else max_x1) - (min_x2 if i == min_xi else min_x1)
            dy = (max_y2 if i == max_yi else max_y1) - (min_y2 if i == min_yi else min_y1)
            ans = min(ans, max(dx, dy))
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumDistance(int[][] points) {
        final int INF = Integer.MAX_VALUE;
        int maxX1 = -INF, maxX2 = -INF, maxY1 = -INF, maxY2 = -INF;
        int minX1 = INF, minX2 = INF, minY1 = INF, minY2 = INF;
        int maxXi = 0, minXi = 0, maxYi = 0, minYi = 0;

        for (int i = 0; i < points.length; i++) {
            int[] p = points[i];
            int x = p[0] + p[1];
            int y = p[1] - p[0];

            // x 最大次大
            if (x > maxX1) {
                maxX2 = maxX1;
                maxX1 = x;
                maxXi = i;
            } else if (x > maxX2) {
                maxX2 = x;
            }

            // x 最小次小
            if (x < minX1) {
                minX2 = minX1;
                minX1 = x;
                minXi = i;
            } else if (x < minX2) {
                minX2 = x;
            }

            // y 最大次大
            if (y > maxY1) {
                maxY2 = maxY1;
                maxY1 = y;
                maxYi = i;
            } else if (y > maxY2) {
                maxY2 = y;
            }

            // y 最小次小
            if (y < minY1) {
                minY2 = minY1;
                minY1 = y;
                minYi = i;
            } else if (y < minY2) {
                minY2 = y;
            }
        }

        int ans = INF;
        for (int i : new int[]{maxXi, minXi, maxYi, minYi}) {
            int dx = (i == maxXi ? maxX2 : maxX1) - (i == minXi ? minX2 : minX1);
            int dy = (i == maxYi ? maxY2 : maxY1) - (i == minYi ? minY2 : minY1);
            ans = Math.min(ans, Math.max(dx, dy));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 更新最大次大
    void update_max(int i, int v, int& max_i, int& max1, int& max2) {
        if (v > max1) {
            max_i = i;
            max2 = max1;
            max1 = v;
        } else if (v > max2) {
            max2 = v;
        }
    }

    // 更新最小次小
    void update_min(int i, int v, int& min_i, int& min1, int& min2) {
        if (v < min1) {
            min_i = i;
            min2 = min1;
            min1 = v;
        } else if (v < min2) {
            min2 = v;
        }
    }

public:
    int minimumDistance(vector<vector<int>>& points) {
        int max_xi, min_xi, max_yi, min_yi;
        int max_x1 = INT_MIN, max_x2 = INT_MIN, max_y1 = INT_MIN, max_y2 = INT_MIN;
        int min_x1 = INT_MAX, min_x2 = INT_MAX, min_y1 = INT_MAX, min_y2 = INT_MAX;

        for (int i = 0; i < points.size(); i++) {
            auto& p = points[i];
            int x = p[0] + p[1];
            int y = p[1] - p[0];
            update_max(i, x, max_xi, max_x1, max_x2);
            update_min(i, x, min_xi, min_x1, min_x2);
            update_max(i, y, max_yi, max_y1, max_y2);
            update_min(i, y, min_yi, min_y1, min_y2);
        }

        int ans = INT_MAX;
        for (int i : {max_xi, min_xi, max_yi, min_yi}) {
            int dx = (i == max_xi ? max_x2 : max_x1) - (i == min_xi ? min_x2 : min_x1);
            int dy = (i == max_yi ? max_y2 : max_y1) - (i == min_yi ? min_y2 : min_y1);
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
	var maxXi, minXi, maxYi, minYi int

	for i, p := range points {
		x := p[0] + p[1]
		y := p[1] - p[0]

		// x 最大次大
		if x > maxX1 {
			maxX2 = maxX1
			maxX1 = x
			maxXi = i
		} else if x > maxX2 {
			maxX2 = x
		}

		// x 最小次小
		if x < minX1 {
			minX2 = minX1
			minX1 = x
			minXi = i
		} else if x < minX2 {
			minX2 = x
		}

		// y 最大次大
		if y > maxY1 {
			maxY2 = maxY1
			maxY1 = y
			maxYi = i
		} else if y > maxY2 {
			maxY2 = y
		}

		// y 最小次小
		if y < minY1 {
			minY2 = minY1
			minY1 = y
			minYi = i
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := inf
	for _, i := range []int{maxXi, minXi, maxYi, minYi} {
		dx := f(i != maxXi, maxX1, maxX2) - f(i != minXi, minX1, minX2)
		dy := f(i != maxYi, maxY1, maxY2) - f(i != minYi, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(b bool, v1, v2 int) int {
	if b {
		return v1
	}
	return v2
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

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
