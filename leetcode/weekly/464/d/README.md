## 一、寻找子问题

先把机器人和墙壁从小到大排序。

考虑最右边的机器人。分类讨论：

- 如果它往左射击，那么需要解决的子问题为：对于前 $n-1$ 个机器人，在第 $n$ 个机器人往左射击的前提下，能摧毁的最大墙壁数量。
- 如果它往右射击，那么需要解决的子问题为：对于前 $n-1$ 个机器人，在第 $n$ 个机器人往右射击的前提下，能摧毁的最大墙壁数量。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

> 注：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。

## 二、状态定义与状态转移方程

根据上面的讨论，定义状态为 $\textit{dfs}(i,j)$，表示对于（排序后）下标在 $[0,i]$ 中的机器人，在机器人 $i+1$ 往左/右射击的前提下，能摧毁的最大墙壁数量。其中 $j=0$ 表示机器人 $i+1$ 往左射击，$j=1$ 表示机器人 $i+1$ 往右射击。

考虑机器人 $i$ 往哪个方向射击：

- 往左，那么接下来要解决的问题是，下标在 $[0,i-1]$ 中的机器人，在机器人 $i$ 往左射击的前提下，能摧毁的最大墙壁数量。即 $\textit{dfs}(i-1,0)$。然后加上机器人 $i$ 摧毁的墙壁数量。
   - 往左最远为 $\textit{leftX} = \max(x_i - d_i,x_{i-1}+1)$，其中 $x_i$ 和 $d_i$ 分别表示机器人 $i$ 的位置和射击距离。为避免重复计算，我们规定，往左不能到达机器人 $i-1$。
   - 在 $\textit{walls}$ 中二分查找 $\ge \textit{leftX}$ 的第一个数的下标，记作 $\textit{left}$。
   - 在 $\textit{walls}$ 中二分查找 $\le x_i$ 的最后一个数的下标**加一**。根据 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)，转化成二分查找 $\ge x_i + 1$ 的第一个数的下标，记作 $\textit{cur}_0$。
   - 那么 $[\textit{left},\textit{cur}_0-1]$ 中的墙都能摧毁，这有 $\textit{cur}_0- \textit{left}$ 个。
- 往右，那么接下来要解决的问题是，下标在 $[0,i-1]$ 中的机器人，在机器人 $i$ 往右射击的前提下，能摧毁的最大墙壁数量。即 $\textit{dfs}(i-1,1)$。
  - 往右最远为 $\textit{rightX} = \min(x_i + d_i,x_{i+1}-1)$ 或者 $\min(x_i + d_i,x_{i+1}-d_{i+1}-1)$，取决于右边那个机器人是往右还是往左射击。
  - 在 $\textit{walls}$ 中二分查找 $\le \textit{rightX}$ 的最后一个数的下标**加一**，即 $\ge \textit{rightX} + 1$ 的第一个数的下标，记作 $\textit{right}$。
  - 在 $\textit{walls}$ 中二分查找 $\ge x_i$ 的第一个数的下标，记作 $\textit{cur}_1$。
  - 那么 $[\textit{cur}_1,\textit{right}-1]$ 中的墙都能摧毁，这有 $\textit{right} - \textit{cur}_1$ 个。

这两种情况取最大值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \max(\textit{dfs}(i-1,0) + \textit{cur}_0- \textit{left}, \textit{dfs}(i-1,1) + \textit{right} - \textit{cur}_1)
$$

**递归边界**：$\textit{dfs}(-1,j)=0$。没有机器人，无法摧毁墙壁。

**递归入口**：$\textit{dfs}(n-1,1)$。机器人 $n-1$ 右边没有机器人，等价于右边那个机器人往右射击。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

⚠**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1X9eJz2EWE/?t=38m44s)，欢迎点赞关注~

### 优化前

```py [sol-Python3]
class Solution:
    def maxWalls(self, robots: List[int], distance: List[int], walls: List[int]) -> int:
        n = len(robots)
        a = sorted(zip(robots, distance), key=lambda p: p[0])
        walls.sort()

        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int) -> int:
            if i < 0:
                return 0

            x, d = a[i]
            # 往左射，墙的坐标范围为 [left_x, x]
            left_x = x - d
            if i > 0:
                left_x = max(left_x, a[i - 1][0] + 1)  # +1 表示不能射到左边那个机器人
            left = bisect_left(walls, left_x)
            cur = bisect_right(walls, x)
            res_left = dfs(i - 1, 0) + cur - left  # 下标在 [left, cur-1] 中的墙都能摧毁

            # 往右射，墙的坐标范围为 [x, right_x]
            right_x = x + d
            if i + 1 < n:
                x2, d2 = a[i + 1]
                if j == 0:  # 右边那个机器人往左射
                    x2 -= d2
                right_x = min(right_x, x2 - 1)  # -1 表示不能射到右边那个机器人（或者它往左射到的墙）
            right = bisect_right(walls, right_x)
            cur = bisect_left(walls, x)
            res_right = dfs(i - 1, 1) + right - cur  # 下标在 [cur, right-1] 中的墙都能摧毁

            return max(res_left, res_right)

        return dfs(n - 1, 1)
```

```java [sol-Java]
class Solution {
    public int maxWalls(int[] robots, int[] distance, int[] walls) {
        int n = robots.length;
        int[][] a = new int[n][2];
        for (int i = 0; i < n; i++) {
            a[i][0] = robots[i];
            a[i][1] = distance[i];
        }
        Arrays.sort(a, (p, q) -> p[0] - q[0]);
        Arrays.sort(walls);

        int[][] memo = new int[n][2];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, 1, a, walls, memo);
    }

    private int dfs(int i, int j, int[][] a, int[] walls, int[][] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
      
        int x = a[i][0], d = a[i][1];
        // 往左射，墙的坐标范围为 [leftX, x]
        int leftX = x - d;
        if (i > 0) {
            leftX = Math.max(leftX, a[i - 1][0] + 1); // +1 表示不能射到左边那个机器人
        }
        int left = lowerBound(walls, leftX);
        int cur = lowerBound(walls, x + 1);
        int resLeft = dfs(i - 1, 0, a, walls, memo) + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

        // 往右射，墙的坐标范围为 [x, rightX]
        int rightX = x + d;
        if (i + 1 < a.length) {
            int x2 = a[i + 1][0];
            if (j == 0) { // 右边那个机器人往左射
                x2 -= a[i + 1][1];
            }
            rightX = Math.min(rightX, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
        }
        int right = lowerBound(walls, rightX + 1);
        cur = lowerBound(walls, x);
        int resRight = dfs(i - 1, 1, a, walls, memo) + right - cur; // 下标在 [cur, right-1] 中的墙都能摧毁

        return memo[i][j] = Math.max(resLeft, resRight); // 记忆化
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
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
    int maxWalls(vector<int>& robots, vector<int>& distance, vector<int>& walls) {
        int n = robots.size();
        struct Pair { int x, d; };
        vector<Pair> a(n);
        for (int i = 0; i < n; i++) {
            a[i] = {robots[i], distance[i]};
        }
        ranges::sort(a, {}, &Pair::x);
        ranges::sort(walls);

        vector memo(n, array<int, 2>{-1, -1}); // -1 表示没有计算过
        auto dfs = [&](this auto&& dfs, int i, int j) -> int {
            if (i < 0) {
                return 0;
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            
            auto [x, d] = a[i];
            // 往左射，墙的坐标范围为 [left_x, x]
            int left_x = x - d;
            if (i > 0) {
                left_x = max(left_x, a[i - 1].x + 1); // +1 表示不能射到左边那个机器人
            }
            int left = ranges::lower_bound(walls, left_x) - walls.begin();
            int cur = ranges::upper_bound(walls, x) - walls.begin();
            res = dfs(i - 1, 0) + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            // 往右射，墙的坐标范围为 [x, right_x]
            int right_x = x + d;
            if (i + 1 < n) {
                auto [x2, d2] = a[i + 1];
                if (j == 0) { // 右边那个机器人往左射
                    x2 -= d2;
                }
                right_x = min(right_x, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
            }
            int right = ranges::upper_bound(walls, right_x) - walls.begin();
            cur = ranges::lower_bound(walls, x) - walls.begin();
            res = max(res, dfs(i - 1, 1) + right - cur); // 下标在 [cur, right-1] 中的墙都能摧毁
            return res;
        };

        return dfs(n - 1, 1);
    }
};
```

```go [sol-Go]
func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	type pair struct{ x, d int }
	a := make([]pair, n)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}

		// 往左射，墙的坐标范围为 [leftX, a[i].x]
		leftX := a[i].x - a[i].d
		if i > 0 {
			leftX = max(leftX, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		}
		left := sort.SearchInts(walls, leftX)
		cur := sort.SearchInts(walls, a[i].x+1)
		res := dfs(i-1, 0) + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		// 往右射，墙的坐标范围为 [a[i].x, rightX]
		rightX := a[i].x + a[i].d
		if i+1 < n {
			x2 := a[i+1].x
			if j == 0 { // 右边那个机器人往左射
				x2 -= a[i+1].d
			}
			rightX = min(rightX, x2-1) // 不能到达右边那个机器人（或者它往左射到的墙）
		}
		right := sort.SearchInts(walls, rightX+1)
		cur = sort.SearchInts(walls, a[i].x)
		res = max(res, dfs(i-1, 1)+right-cur) // 下标在 [cur, right-1] 中的墙都能摧毁

		*p = res
		return res
	}
	return dfs(n-1, 1)
}
```

### 优化

添加两个位置分别为 $0$ 和 $\infty$ 的机器人，当作哨兵，从而简化边界的判断。

```py [sol-Python3]
class Solution:
    def maxWalls(self, robots: List[int], distance: List[int], walls: List[int]) -> int:
        n = len(robots)
        a = [(0, 0)] + sorted(zip(robots, distance), key=lambda p: p[0]) + [(inf, 0)]
        walls.sort()

        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int) -> int:
            if i == 0:
                return 0

            x, d = a[i]
            # 往左射，墙的坐标范围为 [left_x, x]
            left_x = max(x - d, a[i - 1][0] + 1)  # +1 表示不能射到左边那个机器人
            left = bisect_left(walls, left_x)
            cur = bisect_right(walls, x)
            res_left = dfs(i - 1, 0) + cur - left  # 下标在 [left, cur-1] 中的墙都能摧毁

            # 往右射，墙的坐标范围为 [x, right_x]
            x2, d2 = a[i + 1]
            if j == 0:  # 右边那个机器人往左射
                x2 -= d2
            right_x = min(x + d, x2 - 1)  # -1 表示不能射到右边那个机器人（或者它往左射到的墙）
            right = bisect_right(walls, right_x)
            cur = bisect_left(walls, x)
            res_right = dfs(i - 1, 1) + right - cur  # 下标在 [cur, right-1] 中的墙都能摧毁

            return max(res_left, res_right)

        return dfs(n, 1)
```

```java [sol-Java]
class Solution {
    public int maxWalls(int[] robots, int[] distance, int[] walls) {
        int n = robots.length;
        int[][] a = new int[n + 2][2];
        for (int i = 0; i < n; i++) {
            a[i][0] = robots[i];
            a[i][1] = distance[i];
        }
        a[n + 1][0] = Integer.MAX_VALUE;
        Arrays.sort(a, (p, q) -> p[0] - q[0]);
        Arrays.sort(walls);

        int[][] memo = new int[n + 1][2];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n, 1, a, walls, memo);
    }

    private int dfs(int i, int j, int[][] a, int[] walls, int[][] memo) {
        if (i == 0) {
            return 0;
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }

        int x = a[i][0], d = a[i][1];
        // 往左射，墙的坐标范围为 [leftX, x]
        int leftX = Math.max(x - d, a[i - 1][0] + 1); // +1 表示不能射到左边那个机器人
        int left = lowerBound(walls, leftX);
        int cur = lowerBound(walls, x + 1);
        int resLeft = dfs(i - 1, 0, a, walls, memo) + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

        // 往右射，墙的坐标范围为 [x, rightX]
        int x2 = a[i + 1][0];
        if (j == 0) { // 右边那个机器人往左射
            x2 -= a[i + 1][1];
        }
        int rightX = Math.min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
        int right = lowerBound(walls, rightX + 1);
        cur = lowerBound(walls, x);
        int resRight = dfs(i - 1, 1, a, walls, memo) + right - cur; // 下标在 [cur, right-1] 中的墙都能摧毁

        return memo[i][j] = Math.max(resLeft, resRight); // 记忆化
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
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
    int maxWalls(vector<int>& robots, vector<int>& distance, vector<int>& walls) {
        int n = robots.size();
        struct Pair { int x, d; };
        vector<Pair> a(n + 2);
        for (int i = 0; i < n; i++) {
            a[i] = {robots[i], distance[i]};
        }
        a[n + 1].x = INT_MAX;
        ranges::sort(a, {}, &Pair::x);
        ranges::sort(walls);

        vector memo(n + 1, array<int, 2>{-1, -1}); // -1 表示没有计算过
        auto dfs = [&](this auto&& dfs, int i, int j) -> int {
            if (i == 0) {
                return 0;
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }

            auto [x, d] = a[i];
            // 往左射，墙的坐标范围为 [left_x, x]
            int left_x = max(x - d, a[i - 1].x + 1); // +1 表示不能射到左边那个机器人
            int left = ranges::lower_bound(walls, left_x) - walls.begin();
            int cur = ranges::upper_bound(walls, x) - walls.begin();
            res = dfs(i - 1, 0) + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            // 往右射，墙的坐标范围为 [x, right_x]
            auto [x2, d2] = a[i + 1];
            if (j == 0) { // 右边那个机器人往左射
                x2 -= d2;
            }
            int right_x = min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
            int right = ranges::upper_bound(walls, right_x) - walls.begin();
            cur = ranges::lower_bound(walls, x) - walls.begin();
            res = max(res, dfs(i - 1, 1) + right - cur); // 下标在 [cur, right-1] 中的墙都能摧毁
            return res;
        };

        return dfs(n, 1);
    }
};
```

```go [sol-Go]
func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	memo := make([][2]int, n+1)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}

		// 往左射，墙的坐标范围为 [leftX, a[i].x]
		leftX := max(a[i].x-a[i].d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		left := sort.SearchInts(walls, leftX)
		cur := sort.SearchInts(walls, a[i].x+1)
		res := dfs(i-1, 0) + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		// 往右射，墙的坐标范围为 [a[i].x, rightX]
		x2 := a[i+1].x
		if j == 0 { // 右边那个机器人往左射
			x2 -= a[i+1].d
		}
		rightX := min(a[i].x+a[i].d, x2-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		right := sort.SearchInts(walls, rightX+1)
		cur = sort.SearchInts(walls, a[i].x)
		res = max(res, dfs(i-1, 1)+right-cur) // 下标在 [cur, right-1] 中的墙都能摧毁

		*p = res
		return res
	}
	return dfs(n, 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m + n\log m)$，其中 $n$ 是 $\textit{robots}$ 的长度，$m$ 是 $\textit{walls}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(\log m)$，所以动态规划的时间复杂度为 $\mathcal{O}(n\log m)$。前面排序需要 $\mathcal{O}(n\log n + m\log m)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。保存多少状态，就需要多少空间。忽略排序的栈开销。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示对于（排序，添加哨兵后的）下标在 $[1,i]$ 中的机器人，在机器人 $i+1$ 往左/右射击的前提下，能摧毁的最大墙壁数量。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \max(f[i-1][0] + \textit{cur}_0- \textit{left}, f[i-1][1] + \textit{right} - \textit{cur}_1)
$$

初始值 $f[0][j]=0$，翻译自（添加哨兵后的）递归边界 $\textit{dfs}(0,j)=0$。

答案为 $f[n][1]$，翻译自（添加哨兵后的）递归入口 $\textit{dfs}(n,1)$。

```py [sol-Python3]
class Solution:
    def maxWalls(self, robots: List[int], distance: List[int], walls: List[int]) -> int:
        n = len(robots)
        a = [(0, 0)] + sorted(zip(robots, distance), key=lambda p: p[0]) + [(inf, 0)]
        walls.sort()

        f = [[0, 0] for _ in range(n + 1)]
        for i in range(1, n + 1):
            x, d = a[i]

            # 往左射，墙的坐标范围为 [left_x, x]
            left_x = max(x - d, a[i - 1][0] + 1)  # +1 表示不能射到左边那个机器人
            left = bisect_left(walls, left_x)
            cur = bisect_right(walls, x)
            left_res = f[i - 1][0] + cur - left  # 下标在 [left, cur-1] 中的墙都能摧毁

            cur = bisect_left(walls, x)
            for j in range(2):
                # 往右射，墙的坐标范围为 [x, right_x]
                x2, d2 = a[i + 1]
                if j == 0:  # 右边那个机器人往左射
                    x2 -= d2
                right_x = min(x + d, x2 - 1)  # -1 表示不能射到右边那个机器人（或者它往左射到的墙）
                right = bisect_right(walls, right_x)
                f[i][j] = max(left_res, f[i - 1][1] + right - cur)  # 下标在 [cur, right-1] 中的墙都能摧毁
        return f[n][1]
```

```java [sol-Java]
class Solution {
    public int maxWalls(int[] robots, int[] distance, int[] walls) {
        int n = robots.length;
        int[][] a = new int[n + 2][2];
        for (int i = 0; i < n; i++) {
            a[i][0] = robots[i];
            a[i][1] = distance[i];
        }
        a[n + 1][0] = Integer.MAX_VALUE;
        Arrays.sort(a, (p, q) -> p[0] - q[0]);
        Arrays.sort(walls);

        int[][] f = new int[n + 1][2];
        for (int i = 1; i <= n; i++) {
            int x = a[i][0], d = a[i][1];

            // 往左射，墙的坐标范围为 [leftX, x]
            int leftX = Math.max(x - d, a[i - 1][0] + 1); // +1 表示不能射到左边那个机器人
            int left = lowerBound(walls, leftX);
            int cur = lowerBound(walls, x + 1);
            int leftRes = f[i - 1][0] + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            cur = lowerBound(walls, x);
            for (int j = 0; j < 2; j++) {
                // 往右射，墙的坐标范围为 [x, rightX]
                int x2 = a[i + 1][0];
                if (j == 0) { // 右边那个机器人往左射
                    x2 -= a[i + 1][1];
                }
                int rightX = Math.min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
                int right = lowerBound(walls, rightX + 1);
                f[i][j] = Math.max(leftRes, f[i - 1][1] + right - cur); // 下标在 [cur, right-1] 中的墙都能摧毁
            }
        }
        return f[n][1];
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
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
    int maxWalls(vector<int>& robots, vector<int>& distance, vector<int>& walls) {
        int n = robots.size();
        struct Pair { int x, d; };
        vector<Pair> a(n + 2);
        for (int i = 0; i < n; i++) {
            a[i] = {robots[i], distance[i]};
        }
        a[n + 1].x = INT_MAX;
        ranges::sort(a, {}, &Pair::x);
        ranges::sort(walls);

        vector<array<int, 2>> f(n + 1);
        for (int i = 1; i <= n; i++) {
            auto [x, d] = a[i];

            // 往左射，墙的坐标范围为 [left_x, x]
            int left_x = max(x - d, a[i - 1].x + 1); // +1 表示不能射到左边那个机器人
            int left = ranges::lower_bound(walls, left_x) - walls.begin();
            int cur = ranges::upper_bound(walls, x) - walls.begin();
            int left_res = f[i - 1][0] + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            cur = ranges::lower_bound(walls, x) - walls.begin();
            for (int j = 0; j < 2; j++) {
                // 往右射，墙的坐标范围为 [x, right_x]
                auto [x2, d2] = a[i + 1];
                if (j == 0) { // 右边那个机器人往左射
                    x2 -= d2;
                }
                int right_x = min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
                int right = ranges::upper_bound(walls, right_x) - walls.begin();
                f[i][j] = max(left_res, f[i - 1][1] + right - cur); // 下标在 [cur, right-1] 中的墙都能摧毁
            }
        }
        return f[n][1];
    }
};
```

```go [sol-Go]
func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	f := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		p := a[i]

		// 往左射，墙的坐标范围为 [leftX, p.x]
		leftX := max(p.x-p.d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		left := sort.SearchInts(walls, leftX)
		cur := sort.SearchInts(walls, p.x+1)
		leftRes := f[i-1][0] + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		cur = sort.SearchInts(walls, p.x)
		for j := range 2 {
			// 往右射，墙的坐标范围为 [p.x, rightX]
			x2 := a[i+1].x
			if j == 0 { // 右边那个机器人往左射
				x2 -= a[i+1].d
			}
			rightX := min(p.x+p.d, x2-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
			right := sort.SearchInts(walls, rightX+1)
			f[i][j] = max(leftRes, f[i-1][1]+right-cur) // 下标在 [cur, right-1] 中的墙都能摧毁
		}
	}
	return f[n][1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m + n\log m)$，其中 $n$ 是 $\textit{robots}$ 的长度，$m$ 是 $\textit{walls}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。忽略排序的栈开销。

## 五、空间优化

观察上面的状态转移方程，在计算 $f[i+1]$ 时，只会用到 $f[i]$，不会用到比 $i$ 更早的状态。

类似 [背包问题](https://www.bilibili.com/video/BV16Y411v7Y6/)，去掉 $f$ 的第一个维度，把 $f[i+1]$ 和 $f[i]$ 保存到**同一个数组**中。

```py [sol-Python3]
# 手写 min max 更快
min = lambda a, b: b if b < a else a
max = lambda a, b: b if b > a else a

class Solution:
    def maxWalls(self, robots: List[int], distance: List[int], walls: List[int]) -> int:
        n = len(robots)
        a = [(0, 0)] + sorted(zip(robots, distance), key=lambda p: p[0]) + [(inf, 0)]
        walls.sort()

        f = [0, 0]
        for i in range(1, n + 1):
            x, d = a[i]

            # 往左射，墙的坐标范围为 [left_x, x]
            left_x = max(x - d, a[i - 1][0] + 1)  # +1 表示不能射到左边那个机器人
            left = bisect_left(walls, left_x)
            cur = bisect_right(walls, x)
            left_res = f[0] + cur - left  # 下标在 [left, cur-1] 中的墙都能摧毁

            cur = bisect_left(walls, x)
            for j in range(2):
                # 往右射，墙的坐标范围为 [x, right_x]
                x2, d2 = a[i + 1]
                if j == 0:  # 右边那个机器人往左射
                    x2 -= d2
                right_x = min(x + d, x2 - 1)  # -1 表示不能射到右边那个机器人（或者它往左射到的墙）
                right = bisect_right(walls, right_x)
                f[j] = max(left_res, f[1] + right - cur)  # 下标在 [cur, right-1] 中的墙都能摧毁
        return f[1]
```

```java [sol-Java]
class Solution {
    public int maxWalls(int[] robots, int[] distance, int[] walls) {
        int n = robots.length;
        int[][] a = new int[n + 2][2];
        for (int i = 0; i < n; i++) {
            a[i][0] = robots[i];
            a[i][1] = distance[i];
        }
        a[n + 1][0] = Integer.MAX_VALUE;
        Arrays.sort(a, (p, q) -> p[0] - q[0]);
        Arrays.sort(walls);

        int[] f = new int[2];
        for (int i = 1; i <= n; i++) {
            int x = a[i][0], d = a[i][1];

            // 往左射，墙的坐标范围为 [leftX, x]
            int leftX = Math.max(x - d, a[i - 1][0] + 1); // +1 表示不能射到左边那个机器人
            int left = lowerBound(walls, leftX);
            int cur = lowerBound(walls, x + 1);
            int leftRes = f[0] + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            cur = lowerBound(walls, x);
            for (int j = 0; j < 2; j++) {
                // 往右射，墙的坐标范围为 [x, rightX]
                int x2 = a[i + 1][0];
                if (j == 0) { // 右边那个机器人往左射
                    x2 -= a[i + 1][1];
                }
                int rightX = Math.min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
                int right = lowerBound(walls, rightX + 1);
                f[j] = Math.max(leftRes, f[1] + right - cur); // 下标在 [cur, right-1] 中的墙都能摧毁
            }
        }
        return f[1];
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
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
    int maxWalls(vector<int>& robots, vector<int>& distance, vector<int>& walls) {
        int n = robots.size();
        struct Pair { int x, d; };
        vector<Pair> a(n + 2);
        for (int i = 0; i < n; i++) {
            a[i] = {robots[i], distance[i]};
        }
        a[n + 1].x = INT_MAX;
        ranges::sort(a, {}, &Pair::x);
        ranges::sort(walls);

        int f[2]{};
        for (int i = 1; i <= n; i++) {
            auto [x, d] = a[i];

            // 往左射，墙的坐标范围为 [left_x, x]
            int left_x = max(x - d, a[i - 1].x + 1); // +1 表示不能射到左边那个机器人
            int left = ranges::lower_bound(walls, left_x) - walls.begin();
            int cur = ranges::upper_bound(walls, x) - walls.begin();
            int left_res = f[0] + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            cur = ranges::lower_bound(walls, x) - walls.begin();
            for (int j = 0; j < 2; j++) {
                // 往右射，墙的坐标范围为 [x, right_x]
                auto [x2, d2] = a[i + 1];
                if (j == 0) { // 右边那个机器人往左射
                    x2 -= d2;
                }
                int right_x = min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
                int right = ranges::upper_bound(walls, right_x) - walls.begin();
                f[j] = max(left_res, f[1] + right - cur); // 下标在 [cur, right-1] 中的墙都能摧毁
            }
        }
        return f[1];
    }
};
```

```go [sol-Go]
func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	f := [2]int{}
	for i := 1; i <= n; i++ {
		p := a[i]

		// 往左射，墙的坐标范围为 [leftX, p.x]
		leftX := max(p.x-p.d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		left := sort.SearchInts(walls, leftX)
		cur := sort.SearchInts(walls, p.x+1)
		leftRes := f[0] + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		cur = sort.SearchInts(walls, p.x)
		for j := range 2 {
			// 往右射，墙的坐标范围为 [p.x, rightX]
			x2 := a[i+1].x
			if j == 0 { // 右边那个机器人往左射
				x2 -= a[i+1].d
			}
			rightX := min(p.x+p.d, x2-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
			right := sort.SearchInts(walls, rightX+1)
			f[j] = max(leftRes, f[1]+right-cur) // 下标在 [cur, right-1] 中的墙都能摧毁
		}
	}
	return f[1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m + n\log m)$，其中 $n$ 是 $\textit{robots}$ 的长度，$m$ 是 $\textit{walls}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。忽略排序的栈开销。

## 六、双指针优化

由于随着 $i$ 变大，二分查找中的 $\textit{left},\textit{cur},\textit{right}$ 也随之变大，我们可以用双指针（多指针）优化。这样算法瓶颈就在排序上了。

```py [sol-Python3]
# 手写 min max 更快
min = lambda a, b: b if b < a else a
max = lambda a, b: b if b > a else a

class Solution:
    def maxWalls(self, robots: List[int], distance: List[int], walls: List[int]) -> int:
        n, m = len(robots), len(walls)
        a = [(0, 0)] + sorted(zip(robots, distance), key=lambda p: p[0]) + [(inf, 0)]
        walls.sort()

        f0 = f1 = left = cur = right0 = right1 = 0
        for i in range(1, n + 1):
            x, d = a[i]

            # 往左射，墙的坐标范围为 [left_x, x]
            left_x = max(x - d, a[i - 1][0] + 1)  # +1 表示不能射到左边那个机器人
            while left < m and walls[left] < left_x:
                left += 1
            while cur < m and walls[cur] < x:
                cur += 1
            cur1 = cur
            if cur < m and walls[cur] == x:
                cur += 1
            left_res = f0 + cur - left  # 下标在 [left, cur-1] 中的墙都能摧毁

            # 往右射，右边那个机器人往左射，墙的坐标范围为 [x, right_x]
            x2, d2 = a[i + 1]
            right_x = min(x + d, x2 - d2 - 1)  # -1 表示不能射到右边那个机器人
            while right0 < m and walls[right0] <= right_x:
                right0 += 1
            f0 = max(left_res, f1 + right0 - cur1)  # 下标在 [cur1, right0-1] 中的墙都能摧毁

            # 往右射，右边那个机器人往右射，墙的坐标范围为 [x, right_x]
            right_x = min(x + d, x2 - 1)  # -1 表示不能射到右边那个机器人
            while right1 < m and walls[right1] <= right_x:
                right1 += 1
            f1 = max(left_res, f1 + right1 - cur1)  # 下标在 [cur1, right1-1] 中的墙都能摧毁
        return f1
```

```java [sol-Java]
class Solution {
    public int maxWalls(int[] robots, int[] distance, int[] walls) {
        int n = robots.length, m = walls.length;
        int[][] a = new int[n + 2][2];
        for (int i = 0; i < n; i++) {
            a[i][0] = robots[i];
            a[i][1] = distance[i];
        }
        a[n + 1][0] = Integer.MAX_VALUE;
        Arrays.sort(a, (p, q) -> p[0] - q[0]);
        Arrays.sort(walls);

        int f0 = 0, f1 = 0, left = 0, cur = 0, right0 = 0, right1 = 0;
        for (int i = 1; i <= n; i++) {
            int x = a[i][0], d = a[i][1];

            // 往左射，墙的坐标范围为 [leftX, x]
            int leftX = Math.max(x - d, a[i - 1][0] + 1); // +1 表示不能射到左边那个机器人
            while (left < m && walls[left] < leftX) {
                left++;
            }
            while (cur < m && walls[cur] < x) {
                cur++;
            }
            int cur1 = cur;
            if (cur < m && walls[cur] == x) {
                cur++;
            }
            int leftRes = f0 + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            // 往右射，右边那个机器人往左射，墙的坐标范围为 [x, rightX]
            int x2 = a[i + 1][0], d2 = a[i + 1][1];
            int rightX = Math.min(x + d, x2 - d2 - 1); // -1 表示不能射到右边那个机器人
            while (right0 < m && walls[right0] <= rightX) {
                right0++;
            }
            f0 = Math.max(leftRes, f1 + right0 - cur1); // 下标在 [cur1, right0-1] 中的墙都能摧毁

            // 往右射，右边那个机器人往右射，墙的坐标范围为 [x, rightX]
            rightX = Math.min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人
            while (right1 < m && walls[right1] <= rightX) {
                right1++;
            }
            f1 = Math.max(leftRes, f1 + right1 - cur1); // 下标在 [cur1, right1-1] 中的墙都能摧毁
        }
        return f1;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
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
    int maxWalls(vector<int>& robots, vector<int>& distance, vector<int>& walls) {
        int n = robots.size(), m = walls.size();
        struct Pair { int x, d; };
        vector<Pair> a(n + 2);
        for (int i = 0; i < n; i++) {
            a[i] = {robots[i], distance[i]};
        }
        a[n + 1].x = INT_MAX;
        ranges::sort(a, {}, &Pair::x);
        ranges::sort(walls);

        int f0 = 0, f1 = 0, left = 0, cur = 0, right0 = 0, right1 = 0;
        for (int i = 1; i <= n; i++) {
            auto [x, d] = a[i];

            // 往左射，墙的坐标范围为 [left_x, x]
            int left_x = max(x - d, a[i - 1].x + 1); // +1 表示不能射到左边那个机器人
            while (left < m && walls[left] < left_x) {
                left++;
            }
            while (cur < m && walls[cur] < x) {
                cur++;
            }
            int cur1 = cur;
            if (cur < m && walls[cur] == x) {
                cur++;
            }
            int left_res = f0 + cur - left; // 下标在 [left, cur-1] 中的墙都能摧毁

            // 往右射，右边那个机器人往左射，墙的坐标范围为 [x, right_x]
            auto [x2, d2] = a[i + 1];
            int right_x = min(x + d, x2 - d2 - 1); // -1 表示不能射到右边那个机器人
            while (right0 < m && walls[right0] <= right_x) {
                right0++;
            }
            f0 = max(left_res, f1 + right0 - cur1); // 下标在 [cur1, right0-1] 中的墙都能摧毁

            // 往右射，右边那个机器人往右射，墙的坐标范围为 [x, right_x]
            right_x = min(x + d, x2 - 1); // -1 表示不能射到右边那个机器人
            while (right1 < m && walls[right1] <= right_x) {
                right1++;
            }
            f1 = max(left_res, f1 + right1 - cur1); // 下标在 [cur1, right1-1] 中的墙都能摧毁
        }
        return f1;
    }
};
```

```go [sol-Go]
func maxWalls(robots []int, distance []int, walls []int) int {
	n, m := len(robots), len(walls)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	var f0, f1, left, cur, right0, right1 int
	for i := 1; i <= n; i++ {
		p := a[i]

		// 往左射，墙的坐标范围为 [leftX, p.x]
		leftX := max(p.x-p.d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		for left < m && walls[left] < leftX {
			left++
		}
		for cur < m && walls[cur] < p.x {
			cur++
		}
		cur1 := cur
		if cur < m && walls[cur] == p.x {
			cur++
		}
		leftRes := f0 + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		// 往右射，右边那个机器人往左射，墙的坐标范围为 [p.x, rightX]
		q := a[i+1]
		rightX := min(p.x+p.d, q.x-q.d-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		for right0 < m && walls[right0] <= rightX {
			right0++
		}
		f0 = max(leftRes, f1+right0-cur1) // 下标在 [cur1, right0-1] 中的墙都能摧毁

		// 往右射，右边那个机器人往右射，墙的坐标范围为 [p.x, rightX]
		rightX = min(p.x+p.d, q.x-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		for right1 < m && walls[right1] <= rightX {
			right1++
		}
		f1 = max(leftRes, f1+right1-cur1) // 下标在 [cur1, right0-1] 中的墙都能摧毁
	}
	return f1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m)$，其中 $n$ 是 $\textit{robots}$ 的长度，$m$ 是 $\textit{walls}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。忽略排序的栈开销。

## 专题训练

见下面动态规划题单的「**六、状态机 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
