## 式子变形

题目给出的式子有子数组和，我们先用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 将其简化。

定义 $\textit{sumNum}[i+1]$ 表示 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的元素和。

定义 $s[i+1]$ 表示 $\textit{cost}[0]$ 到 $\textit{cost}[i]$ 的元素和。

题目给出的式子转换成

$$
\begin{aligned}
    & (\textit{sumNum}[r+1] + k\cdot i) \cdot (s[r+1] - s[l])      \\
={} & \textit{sumNum}[r+1] \cdot (s[r+1] - s[l]) + k\cdot i \cdot (s[r+1] - s[l])       \\
\end{aligned}
$$

如果能把式子中的 $i$ 去掉，我们就可以写一个 $\mathcal{O}(n^2)$ 的划分型 DP。

横看成岭侧成峰，换一个角度看待 $i \cdot (s[r+1] - s[l])$。

假设划分成了三段，$\textit{cost}$ 的子数组和分别为 $A,B,C$。

这三段的 $i \cdot (s[r+1] - s[l])$ 分别为 $A,2B,3C$，累加得

$$
\begin{aligned}
    & A+2B+3C      \\
={} & (A+B+C) + (B+C) + C        \\
\end{aligned}
$$

如此变形后，我们可以把 $A+B+C$ 当作第一段的 $i \cdot (s[r+1] - s[l])$，把 $B+C$ 当作第二段的 $i \cdot (s[r+1] - s[l])$，把 $C$ 当作第三段的 $i \cdot (s[r+1] - s[l])$。详细证明见文末的 **Abel 求和公式**。

换句话说，我们可以跨越时空，把未来要计算的内容，放到现在计算！

式子中的 $i \cdot (s[r+1] - s[l])$ 可以替换成 $s[n] - s[l]$，因为 $A+B+C,B+C,C$ 都是 $\textit{cost}$ 的后缀和。

题目给出的式子替换成

$$
\textit{sumNum}[r+1] \cdot (s[r+1] - s[l]) + k\cdot (s[n] - s[l])
$$

> 注意上式和原式并不一定相等，但计算所有子数组的上式之和后，是相等的。

## 方法一：划分型 DP

根据 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/)「§5.2 最优划分」，定义 $f[i+1]$ 表示下标 $[0,i]$ 分割后的最小总代价。

枚举最后一个子数组的左端点 $j$，问题变成下标 $[0,j-1]$ 分割后的最小总代价，即 $f[j]$。其中 $j$ 最小是 $0$，最大是 $i$。

取最小值，有

$$
f[i+1] = \min_{j=0}^{i} f[j] + \textit{sumNum}[i+1] \cdot (s[i+1] - s[j]) + k\cdot (s[n] - s[j])
$$

初始值 $f[0]=0$。

答案为 $f[n]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JrZzYhEHt/?t=15m21s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumCost(self, nums: List[int], cost: List[int], k: int) -> int:
        n = len(nums)
        s = list(accumulate(cost, initial=0))  # cost 的前缀和
        f = [0] * (n + 1)
        for i, sum_num in enumerate(accumulate(nums), 1):  # 这里把 i 加一了，下面不用加一
            f[i] = min(f[j] + sum_num * (s[i] - s[j]) + k * (s[n] - s[j])
                       for j in range(i))
        return f[n]
```

```java [sol-Java]
class Solution {
    public long minimumCost(int[] nums, int[] cost, int k) {
        int n = nums.length;
        int[] s = new int[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + cost[i]; // cost 的前缀和
        }

        long[] f = new long[n + 1];
        int sumNum = 0;
        for (int i = 1; i <= n; i++) { // 注意这里 i 从 1 开始，下面不用把 i 加一
            sumNum += nums[i - 1];
            f[i] = Long.MAX_VALUE;
            for (int j = 0; j < i; j++) {
                f[i] = Math.min(f[i], f[j] + (long) sumNum * (s[i] - s[j]) + k * (s[n] - s[j]));
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(vector<int>& nums, vector<int>& cost, int k) {
        int n = nums.size();
        vector<int> s(n + 1);
        partial_sum(cost.begin(), cost.end(), s.begin() + 1); // cost 的前缀和

        vector<long long> f(n + 1, LLONG_MAX);
        f[0] = 0;
        int sum_num = 0;
        for (int i = 1; i <= n; i++) { // 注意这里 i 从 1 开始，下面不用把 i 加一
            sum_num += nums[i - 1];
            for (int j = 0; j < i; j++) {
                f[i] = min(f[i], f[j] + 1LL * sum_num * (s[i] - s[j]) + k * (s[n] - s[j]));
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func minimumCost(nums, cost []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, c := range cost {
		s[i+1] = s[i] + c // cost 的前缀和
	}

	f := make([]int, n+1)
	sumNum := 0
	for i, x := range nums {
		i++ // 这里把 i 加一了，下面不用加一
		sumNum += x
		res := math.MaxInt
		for j := range i {
			res = min(res, f[j]+sumNum*(s[i]-s[j])+k*(s[n]-s[j]))
		}
		f[i] = res
	}
	return int64(f[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：凸包 + 双指针 / 双端队列

**前置知识**：二维计算几何，凸包，Andrew 算法。

同 [3494. 酿造药水需要的最少总时间](https://leetcode.cn/problems/find-the-minimum-amount-of-time-to-brew-potions/)，方法一中的状态转移方程，也可以改成点积的形式，从而可以用凸包优化。

首先把递推式变形，把 $\min$ 中和 $j$ 无关的量提出来：

$$
\begin{aligned}
f[i+1] ={} & \min_{j=0}^{i} f[j] + \textit{sumNum}[i+1] \cdot (s[i+1] - s[j]) + k\cdot (s[n] - s[j])   \\
       ={} & \textit{sumNum}[i+1] \cdot s[i+1] + k\cdot s[n] + \min_{j=0}^{i} f[j] - (\textit{sumNum}[i+1] + k) \cdot s[j] \\
\end{aligned}
$$

把其中的

$$
f[j] - (\textit{sumNum}[i+1] + k) \cdot s[j]
$$

改成点积的形式，这样我们能得到来自几何意义上的观察。

设向量 $\mathbf{v}_j = (s[j], f[j])$。

设向量 $\mathbf{p} = (-\textit{sumNum}[i+1] - k, 1)$。

那么我们求的是

$$
\min_{j=0}^{i} \mathbf{p}\cdot \mathbf{v}_j
$$

根据点积的几何意义，我们求的是 $\mathbf{v}_j$ 在 $\mathbf{p}$ 方向上的投影长度，再乘以 $\mathbf{p}$ 的模长 $||\mathbf{p}||$。由于 $||\mathbf{p}||$ 是个定值，所以要最小化投影长度。

考虑 $\mathbf{v}_j$ 的**下凸包**（用 Andrew 算法计算），在凸包内的点，比凸包顶点的投影长度长。所以只需考虑凸包顶点。

> 由于 $s[j]$ 是单调递增的，求下凸包无需排序。

这样有一个很好的性质：从左到右遍历凸包顶点，$\mathbf{p}\cdot \mathbf{v}_j$ 会先变小再变大（单峰函数）。那么要计算最小值，就类似 [852. 山脉数组的峰顶索引](https://leetcode.cn/problems/peak-index-in-a-mountain-array/)，**二分**首个「上坡」的位置，具体见 [我的题解](https://leetcode.cn/problems/peak-index-in-a-mountain-array/solutions/2984800/er-fen-gen-ju-shang-po-huan-shi-xia-po-p-uoev/)。

实际上不需要二分。由于 $-\textit{sumNum}[i+1] - k$ 是单调递减的，也可以双指针（或者双端队列）。

```py [sol-Python3]
class Vec:
    __slots__ = 'x', 'y'

    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y

    def __sub__(self, b: "Vec") -> "Vec":
        return Vec(self.x - b.x, self.y - b.y)

    def det(self, b: "Vec") -> int:
        return self.x * b.y - self.y * b.x

    def dot(self, b: "Vec") -> int:
        return self.x * b.x + self.y * b.y

class Solution:
    def minimumCost(self, nums: List[int], cost: List[int], k: int) -> int:
        total_cost = sum(cost)

        q = deque([Vec(0, 0)])
        sum_num = sum_cost = 0
        for x, c in zip(nums, cost):
            sum_num += x
            sum_cost += c

            p = Vec(-sum_num - k, 1)
            while len(q) > 1 and p.dot(q[0]) >= p.dot(q[1]):
                q.popleft()

            # 一边算 DP 一边构建下凸包
            p = Vec(sum_cost, p.dot(q[0]) + sum_num * sum_cost + k * total_cost)
            while len(q) > 1 and (q[-1] - q[-2]).det(p - q[-1]) <= 0:
                q.pop()
            q.append(p)

        return q[-1].y
```

```java [sol-Java]
class Solution {
    private record Vec(long x, long y) {
        Vec sub(Vec b) {
            return new Vec(x - b.x, y - b.y);
        }

        long det(Vec b) {
            return x * b.y - y * b.x;
        }

        long dot(Vec b) {
            return x * b.x + y * b.y;
        }
    }

    public long minimumCost(int[] nums, int[] cost, int k) {
        int totalCost = 0;
        for (int c : cost) {
            totalCost += c;
        }

        List<Vec> q = new ArrayList<>();
        q.add(new Vec(0, 0));
        int sumNum = 0;
        int sumCost = 0;
        int j = 0;

        for (int i = 0; i < nums.length; i++) {
            sumNum += nums[i];
            sumCost += cost[i];

            Vec p = new Vec(-sumNum - k, 1);
            while (j + 1 < q.size() && p.dot(q.get(j)) >= p.dot(q.get(j + 1))) {
                j++;
            }

            // 一边算 DP 一边构建下凸包
            p = new Vec(sumCost, p.dot(q.get(j)) + (long) sumNum * sumCost + k * totalCost);
            while (q.size() - j > 1 && q.getLast().sub(q.get(q.size() - 2)).det(p.sub(q.getLast())) <= 0) {
                q.removeLast();
            }
            q.add(p);
        }

        return q.getLast().y;
    }
}
```

```java [sol-Java 数组]
class Solution {
    private record Vec(long x, long y) {
        Vec sub(Vec b) {
            return new Vec(x - b.x, y - b.y);
        }

        long det(Vec b) {
            return x * b.y - y * b.x;
        }

        long dot(Vec b) {
            return x * b.x + y * b.y;
        }
    }

    public long minimumCost(int[] nums, int[] cost, int k) {
        int totalCost = 0;
        for (int c : cost) {
            totalCost += c;
        }

        Vec[] q = new Vec[nums.length + 1];
        int n = 0;
        q[n++] = new Vec(0, 0);
        int sumNum = 0;
        int sumCost = 0;
        int j = 0;

        for (int i = 0; i < nums.length; i++) {
            sumNum += nums[i];
            sumCost += cost[i];

            Vec p = new Vec(-sumNum - k, 1);
            while (j + 1 < n && p.dot(q[j]) >= p.dot(q[j + 1])) {
                j++;
            }

            // 一边算 DP 一边构建下凸包
            p = new Vec(sumCost, p.dot(q[j]) + (long) sumNum * sumCost + k * totalCost);
            while (n - j > 1 && q[n - 1].sub(q[n - 2]).det(p.sub(q[n - 1])) <= 0) {
                n--;
            }
            q[n++] = p;
        }

        return q[n - 1].y;
    }
}
```

```cpp [sol-C++]
struct Vec {
    long long x, y;
    Vec operator-(const Vec& b) { return {x - b.x, y - b.y}; }
    long long det(const Vec& b) { return x * b.y - y * b.x; }
    long long dot(const Vec& b) { return x * b.x + y * b.y; }
};

class Solution {
public:
    long long minimumCost(vector<int>& nums, vector<int>& cost, int k) {
        int total_cost = reduce(cost.begin(), cost.end());

        deque<Vec> q = {{0, 0}};
        int sum_num = 0, sum_cost = 0;
        for (int i = 0; i < nums.size(); i++) {
            sum_num += nums[i];
            sum_cost += cost[i];

            Vec p = {-sum_num - k, 1};
            while (q.size() > 1 && p.dot(q[0]) >= p.dot(q[1])) {
                q.pop_front();
            }

            // 一边算 DP 一边构建下凸包
            p = {sum_cost, p.dot(q[0]) + 1LL * sum_num * sum_cost + k * total_cost};
            while (q.size() > 1 && (q.back() - q[q.size() - 2]).det(p - q.back()) <= 0) {
                q.pop_back();
            }
            q.push_back(p);
        }
        return q.back().y;
    }
};
```

```go [sol-Go]
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

func minimumCost(nums, cost []int, k int) int64 {
	totalCost := 0
	for _, c := range cost {
		totalCost += c
	}

	q := []vec{{}}
	sumNum, sumCost := 0, 0
	for i, x := range nums {
		sumNum += x
		sumCost += cost[i]

		p := vec{-sumNum - k, 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}

		// 一边算 DP 一边构建下凸包
		p = vec{sumCost, p.dot(q[0]) + sumNum*sumCost + k*totalCost}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	return int64(q[len(q)-1].y)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。每个点入队出队各至多一次。
- 空间复杂度：$\mathcal{O}(n)$。

## 附：Abel 求和公式

分部积分法

$$
\int_a^b f(x)g'(x)\,\mathrm{d}x = \left[f(x)g(x)\right]_a^b - \int_a^b f'(x)g(x)\,\mathrm{d}x
$$

的离散形式，叫做 Abel 求和公式

$$
\sum_{i=m}^{n} a_i b_i = a_nS_n - \sum_{i=m}^{n-1} (a_{i+1} - a_i) S_i
$$

其中 $S_i = \sum\limits_{k=m}^{i} b_k$。

**证明**：

$$
\begin{aligned}
    & \sum_{i=m}^{n} a_i b_i      \\
={} & \sum_{i=m}^{n} a_i (S_i - S_{i-1})        \\
={} & \sum_{i=m}^{n} a_nS_n - a_nS_{n-1} + a_{n-1}S_{n-1} - a_{n-1}S_{n-2} + \cdots + a_{m+1}S_{m+1} - a_{m+1}S_m + a_mS_m        \\
={} & a_nS_n - \sum_{i=m}^{n-1} (a_{i+1} - a_i) S_i        \\
\end{aligned}
$$

本题相当于 $a_i = i$，$b_i$ 为 $\textit{cost}$ 的第 $i$ 个子数组和。那么 $S_i$ 就是 $\textit{cost}$ 的前 $i$ 个子数组和的前缀和（注意不是 $\textit{cost}$ 的前缀和）。

假设分割成 $k$ 个子数组。代入 Abel 求和公式，得

$$
\begin{aligned}
\sum_{i=1}^{k} i b_i ={}& i S_k - \sum_{i=1}^{k-1} S_i \\
={} & S_k + (S_k - S_1) + (S_k - S_2) + \cdots + (S_k - S_{k-1})        \\
\end{aligned}
$$

其中 $S_k$ 是整个 $\textit{cost}$ 的元素和，$S_k - S_i$ 是 $\textit{cost}$ 的后缀和。上式相当于 $k$ 个 $\textit{cost}$ 的后缀和，这正是题解开头得出的结论。

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
