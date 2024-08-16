## 方法一：记忆化搜索

### 寻找子问题

看示例 1，$\textit{nums} = [1,4,3,3,2],\ \textit{andValues} = [0,3,3,2]$。

我们要解决的问题是，把 $\textit{nums}$ 划分成 $4$ 个子数组所能得到的最小子数组值之和，其中每个子数组的 AND 值与 $\textit{andValues}$ 中的值一一对应。

从 $\textit{nums}[0]$ 开始。考虑是否要把 $\textit{nums}[0]$ 作为子数组的最后一个数，分类讨论：

- 不把 $\textit{nums}[0]$ 作为子数组的最后一个数，也就是 $\textit{nums}[0]$ 和后续元素在同一个子数组中，那么接下来需要解决的问题为：把 $[4,3,3,2]$ 划分成 $4$ 个子数组，且第一个子数组与 $\textit{nums}[0]=1$ 计算 AND 的值恰好等于 $\textit{andValues}[0]=0$，其余子数组的 AND 值分别等于 $3,3,2$，在满足该条件的情况下，所能得到的最小子数组值之和。
- 把 $\textit{nums}[0]$ 作为子数组的最后一个数，但子数组 AND 等于 $1$，不等于 $\textit{andValues}[0]=0$，不符合题目要求，无法划分。

继续。考虑是否要把 $\textit{nums}[1]$ 作为子数组的最后一个数，分类讨论：

- 不把 $\textit{nums}[1]$ 作为子数组的最后一个数，也就是 $\textit{nums}[1]$ 和后续元素在同一个子数组中，那么接下来需要解决的问题为：把 $[3,3,2]$ 划分成 $4$ 个子数组，且第一个子数组与 $\textit{nums}[0] \& \textit{nums}[1]$ 计算 AND 的值恰好等于 $\textit{andValues}[0]=0$，其余子数组的 AND 值分别等于 $3,3,2$，在满足该条件的情况下，所能得到的最小子数组值之和。注意剩余元素只有 $3$ 个，没法分成 $4$ 个子数组。
- 把 $\textit{nums}[1]$ 作为子数组的最后一个数，注意我们**并不需要知道这个子数组的前面具体有哪些数，只需要知道前面的元素的 AND 值等于** $1$。由于 $\textit{nums}[0] \& \textit{nums}[1]=1\& 4=0 =\textit{andValues}[0]$，符合题目要求，可以划分。接下来需要解决的问题为：把 $[3,3,2]$ 划分成 $3$ 个子数组，子数组的 AND 值分别等于 $3,3,2$，在满足该条件的情况下，所能得到的最小子数组值之和。

是否划分都会把原问题变成一个**和原问题相似的、规模更小的子问题**，都是把一些元素划分成若干段，且每一段的 AND 值与 $\textit{andValues}$ 中的元素匹配。这可以用**递归**解决。

> 注 1：为方便把子数组的最后一个元素加入答案，本题适合从左到右思考。
>
> 注 2：动态规划有「**选或不选**」和「**枚举选哪个**」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「**选或不选**」。

### 状态定义与状态转移方程

递归需要哪些参数？

1. 需要知道当前考虑到 $\textit{nums}$ 的哪个数，其下标记作 $i$。
2. 需要知道当前划分的子数组对应着 $\textit{andValues}$ 的哪个数，其下标记作 $j$。也可以理解为前面已经划分了 $j$ 段。
3. 需要知道当前划分的子数组，在 $i$ 左边的那些元素的 AND 值，记作 $\textit{and}$。再次强调，我们并不需要知道 $i$ 左边具体有哪些数，只需要知道左边那些数的 AND 值是多少即可。

于是，定义 $\textit{dfs}(i,j,\textit{and})$ 表示从左往右划分，目前考虑到 $\textit{nums}[i]$，已经划分了 $j$ 段，且当前待划分的这一段已经参与 AND 运算的结果为 $\textit{and}$，在这种情况下，剩余元素划分得到的最小和。

首先把 $\textit{and}$ 与 $\textit{nums}[i]$ 计算 AND。

用「选或不选」的思想分类讨论：

- **不划分**：继续向右递归 $\textit{dfs}(i+1,j,\textit{and})$。
- **划分**：如果 $\textit{and}=\textit{andValues}[j]$，那么可以划分，即 $\textit{dfs}(i+1,j+1,-1) + \textit{nums}[i]$。这里令 $\textit{and}=-1$ 是因为 $-1$ 的二进制全为 $1$，与任何数 $x$ 的 AND 都是 $x$，适合用来计算新子数组的 AND 值。
- 这两种情况取最小值，就得到了 $\textit{dfs}(i,j,\textit{and})$。

**递归边界**：

- 如果 $n-i<m-j$，那么剩余元素不足，无法划分，返回 $\infty$。
- 如果 $j=m$ 且 $i<n$，还有元素没有划分，返回 $\infty$。
- 如果 $j=m$ 且 $i=n$，划分成功，返回 $0$。

**递归入口**：$\textit{dfs}(0,0,-1)$，即答案。如果答案是 $\infty$ 则返回 $-1$。

请看 [视频讲解](https://www.bilibili.com/video/BV1dJ4m1V7hK/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumValueSum(self, nums: List[int], andValues: List[int]) -> int:
        n, m = len(nums), len(andValues)
        @cache
        def dfs(i: int, j: int, and_: int) -> int:
            if n - i < m - j:  # 剩余元素不足
                return inf
            if j == m:  # 分了 m 段
                return 0 if i == n else inf
            and_ &= nums[i]
            res = dfs(i + 1, j, and_)  # 不划分
            if and_ == andValues[j]:  # 划分，nums[i] 是这一段的最后一个数
                res = min(res, dfs(i + 1, j + 1, -1) + nums[i])
            return res
        ans = dfs(0, 0, -1)
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumValueSum(int[] nums, int[] andValues) {
        Map<Long, Integer> memo = new HashMap<>();
        int ans = dfs(0, 0, -1, nums, andValues, memo);
        return ans < Integer.MAX_VALUE / 2 ? ans : -1;
    }

    private int dfs(int i, int j, int and, int[] nums, int[] andValues, Map<Long, Integer> memo) {
        int n = nums.length;
        int m = andValues.length;
        if (n - i < m - j) { // 剩余元素不足
            return Integer.MAX_VALUE / 2; // 除 2 防止下面 + nums[i] 溢出
        }
        if (j == m) { // 分了 m 段
            return i == n ? 0 : Integer.MAX_VALUE / 2;
        }
        and &= nums[i];
        // 三个参数压缩成一个 long
        long mask = (long) i << 36 | (long) j << 32 | and;
        if (memo.containsKey(mask)) { // 之前计算过
            return memo.get(mask);
        }
        int res = dfs(i + 1, j, and, nums, andValues, memo); // 不划分
        if (and == andValues[j]) { // 划分，nums[i] 是这一段的最后一个数
            res = Math.min(res, dfs(i + 1, j + 1, -1, nums, andValues, memo) + nums[i]);
        }
        memo.put(mask, res); // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumValueSum(vector<int>& nums, vector<int>& andValues) {
        const int INF = INT_MAX / 2; // 除 2 防止下面 + nums[i] 溢出
        int n = nums.size(), m = andValues.size();
        unordered_map<long long, int> memo;
        auto dfs = [&](auto&& dfs, int i, int j, int and_) -> int {
            if (n - i < m - j) { // 剩余元素不足
                return INF;
            }
            if (j == m) { // 分了 m 段
                return i == n ? 0 : INF;
            }
            and_ &= nums[i];
            // 三个参数压缩成一个 long long
            long long mask = (long long) i << 36 | (long long) j << 32 | and_;
            if (memo.contains(mask)) { // 之前计算过
                return memo[mask];
            }
            int res = dfs(dfs, i + 1, j, and_); // 不划分
            if (and_ == andValues[j]) { // 划分，nums[i] 是这一段的最后一个数
                res = min(res, dfs(dfs, i + 1, j + 1, -1) + nums[i]);
            }
            return memo[mask] = res; // 记忆化
        };
        int ans = dfs(dfs, 0, 0, -1);
        return ans < INF ? ans : -1;
    }
};
```

```go [sol-Go]
func minimumValueSum(nums, andValues []int) int {
    const inf = math.MaxInt / 2 // 除 2 防止下面 +nums[i] 溢出
    n, m := len(nums), len(andValues)
    type args struct{ i, j, and int }
    memo := map[args]int{}
    var dfs func(int, int, int) int
    dfs = func(i, j, and int) int {
        if n-i < m-j { // 剩余元素不足
            return inf
        }
        if j == m { // 分了 m 段
            if i == n {
                return 0
            }
            return inf
        }
        and &= nums[i]
        p := args{i, j, and}
        if res, ok := memo[p]; ok { // 之前计算过
            return res
        }
        res := dfs(i+1, j, and)  // 不划分
        if and == andValues[j] { // 划分，nums[i] 是这一段的最后一个数
            res = min(res, dfs(i+1, j+1, -1)+nums[i])
        }
        memo[p] = res // 记忆化
        return res
    }
    ans := dfs(0, 0, -1)
    if ans == inf {
        return -1
    }
    return ans
}
```

#### 复杂度分析

有多少个状态？

AND 的**性质**：AND 的数越多，AND 的结果就越小。

**定理**：总共有 $\mathcal{O}(n\log U)$ 个不同的子数组 AND 值，其中 $U=\max(\textit{nums})$。

**证明**：考察子数组右端点固定为 $i$ 的情况。我们从 $i$ 开始，向左枚举子数组的左端点 $j$。随着 AND 的数字越来越多，AND 的结果，要么不变，要么减小，且减少一定是某个比特 $1$ 变成 $0$。由于 $\textit{nums}[i]$ 有 $\mathcal{O}(\log U)$ 个比特 $1$，所以在右端点固定为 $i$ 的前提下，向左不断 AND $\textit{nums}[j]$ 只能得到 $\mathcal{O}(\log U)$ 个不同的 AND 结果。那么对于所有的子数组，一共有 $\mathcal{O}(n\log U)$ 个不同的子数组 AND 值。

所以对于 $\textit{dfs}$ 中的一个固定的参数 $i$，只有 $\mathcal{O}(\log U)$ 个不同的参数 $\textit{and}$ 的值。再乘上 $\mathcal{O}(m)$ 个不同的参数 $j$，一共有 $\mathcal{O}(nm\log U)$ 个状态。

- 时间复杂度：$\mathcal{O}(nm\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{andValues}$ 的长度，$U=\max(\textit{nums})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nm\log U)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(nm\log U)$。
- 空间复杂度：$\mathcal{O}(nm\log U)$。

本题属于**划分型 DP**，更多相似题目，见 [DP 题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**§6.3 约束划分个数**」。

## 方法二：单调队列 + logTrick 优化

⚠**阅读要求：读者需要做过一些单调队列优化 DP 的题目，以及一些 logTrick 的题目。**

定义 $f[j][i+1]$ 表示把 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 分成 $j$ 段的最小子数组值之和。

如果第 $j$ 段的下标范围是从 $k$ 到 $i$，那么有状态转移方程

$$
f[j][i+1] = \textit{nums}[i] + \min_{k} f[j-1][k]
$$

其中下标从 $k$ 到 $i$ 的 $\textit{nums}$ 的 AND 必须恰好等于 $\textit{target} = \textit{andValues}[j-1]$。

为了计算 $k$ 的范围，需要用到 **logTrick**。相关题目可以看 [位运算题单](https://leetcode.cn/circle/discuss/dHn9Vk/)。

对于一个固定的 $\textit{target}$，当子数组右端点 $i$ 变大时，子数组左端点 $k$ 的范围区间 $[l,r]$ 也在右移，所以计算 DP 转移来源的最小值，可以像 [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/) 那样用**单调队列**解决，原理请看 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。

初始值：$f[0][0]=0,\ f[0][i\ge 1]=\infty,\ f[j\ge 1][0]=\infty$。

答案：$f[m][n]$。

代码实现时，$f$ 数组的第一个维度可以去掉，改成用两个一维数组滚动计算。

### 写法一：用额外数组维护

如果觉得写法一的 logTrick 不好理解的话，可以看写法二，相对来说更加简洁。

```py [sol-Python3]
class Solution:
    def minimumValueSum(self, nums: List[int], andValues: List[int]) -> int:
        n = len(nums)
        f = [0] + [inf] * n
        new_f = [0] * (n + 1)
        for target in andValues:
            a = []  # logTrick 子数组 AND 和子数组左端点
            q = deque()  # 单调队列，保存 f 的下标
            qi = 0  # 单调队列目前处理到 f[qi]

            new_f[0] = inf
            for i, x in enumerate(nums):
                for p in a:
                    p[0] &= x
                a.append([x, i])

                # 原地去重，并去掉 AND 值小于 target 的数据
                j = 0
                last = -1
                for p in a:
                    and_ = p[0]
                    if and_ >= target and and_ != last:
                        a[j] = p
                        j += 1
                        last = and_
                del a[j:]

                # 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值
                # 下面是单调队列的滑窗过程

                if a and a[0][0] == target:
                    # 现在 a[0][1] 和 a[1][1]-1 分别是子数组左端点的最小值和最大值
                    r = a[1][1] - 1 if len(a) > 1 else i

                    # 单调队列：右边入
                    while qi <= r:
                        while q and f[qi] <= f[q[-1]]:
                            q.pop()
                        q.append(qi)
                        qi += 1

                    # 单调队列：左边出
                    while q[0] < a[0][1]:
                        q.popleft()

                    # 单调队列：计算答案
                    new_f[i + 1] = f[q[0]] + x  # 队首就是最小值
                else:
                    new_f[i + 1] = inf
            f, new_f = new_f, f
        return f[n] if f[n] < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumValueSum(int[] nums, int[] andValues) {
        final int INF = Integer.MAX_VALUE / 2;
        int n = nums.length;
        int[] f = new int[n + 1];
        Arrays.fill(f, 1, n + 1, INF);
        int[] newF = new int[n + 1];
        int[] and = new int[18]; // logTrick 子数组 AND
        int[] left = new int[18]; // logTrick 子数组左端点
        int[] q = new int[n + 1]; // 用数组模拟单调队列，保存 f 的下标

        for (int target : andValues) {
            int an = 0; // and 和 left 的大小
            int ql = 0, qr = 0; // q 的元素下标范围 [ql, qr)
            int qi = 0; // 单调队列目前处理到 f[qi]

            newF[0] = INF;
            for (int i = 0; i < n; i++) {
                int x = nums[i];
                for (int j = 0; j < an; j++) {
                    and[j] &= x;
                }
                and[an] = x;
                left[an++] = i;

                // 原地去重，并去掉 AND 值小于 target 的数据
                int j = 0;
                int last = -1;
                for (int k = 0; k < an; k++) {
                    if (and[k] >= target && and[k] != last) {
                        last = and[j] = and[k];
                        left[j++] = left[k];
                    }
                }
                an = j;

                // 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值
                // 下面是单调队列的滑窗过程

                if (an > 0 && and[0] == target) {
                    // 现在 left[0] 和 left[1]-1 分别是子数组左端点的最小值和最大值
                    int r = an > 1 ? left[1] - 1 : i;

                    // 单调队列：右边入
                    for (; qi <= r; qi++) {
                        while (qr > ql && f[qi] <= f[q[qr - 1]]) {
                            qr--;
                        }
                        q[qr++] = qi;
                    }

                    // 单调队列：左边出
                    while (ql < qr && q[ql] < left[0]) {
                        ql++;
                    }

                    // 单调队列：计算答案
                    newF[i + 1] = f[q[ql]] + x; // 队首就是最小值
                } else {
                    newF[i + 1] = INF;
                }
            }
            int[] tmp = f;
            f = newF;
            newF = tmp;
        }
        return f[n] < INF ? f[n] : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumValueSum(vector<int>& nums, vector<int>& andValues) {
        const int INF = INT_MAX / 2;
        int n = nums.size();
        vector<int> f(n + 1, INF);
        vector<int> new_f(n + 1);

        f[0] = 0;
        for (int target : andValues) {
            vector<pair<int, int>> a; // logTrick 子数组 AND 和子数组左端点
            deque<int> q; // 单调队列，保存 f 的下标
            int qi = 0; // 单调队列目前处理到 f[qi]

            new_f[0] = INF;
            for (int i = 0; i < n; i++) {
                int x = nums[i];
                for (auto& [and_, _] : a) {
                    and_ &= x;
                }
                a.emplace_back(x, i);

                // 原地去重，并去掉 AND 值小于 target 的数据
                int j = 0, last = -1;
                for (auto& p : a) {
                    int and_ = p.first;
                    if (and_ >= target && and_ != last) {
                        a[j++] = p;
                        last = and_;
                    }
                }
                a.resize(j);

                // 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值
                // 下面是单调队列的滑窗过程

                if (!a.empty() && a[0].first == target) {
                    // 现在 a[0].second 和 a[1].second-1 分别是子数组左端点的最小值和最大值
                    int r = a.size() > 1 ? a[1].second - 1 : i;

                    // 单调队列：右边入
                    for (; qi <= r; qi++) {
                        while (!q.empty() && f[qi] <= f[q.back()]) {
                            q.pop_back();
                        }
                        q.push_back(qi);
                    }

                    // 单调队列：左边出
                    while (q.front() < a[0].second) {
                        q.pop_front();
                    }

                    // 单调队列：计算答案
                    new_f[i + 1] = f[q.front()] + x; // 队首就是最小值
                } else {
                    new_f[i + 1] = INF;
                }
            }
            swap(f, new_f);
        }
        return f[n] < INF ? f[n] : -1;
    }
};
```

```go [sol-Go]
func minimumValueSum(nums, andValues []int) int {
    const inf = math.MaxInt / 2
    n := len(nums)
    f := make([]int, n+1)
    for i := 1; i <= n; i++ {
        f[i] = inf
    }
    newF := make([]int, n+1)
    for _, target := range andValues {
        type pair struct{ and, l int }
        a := []pair{} // logTrick 子数组 AND 和子数组左端点
        q := []int{} // 单调队列，保存 f 的下标
        qi := 0 // 单调队列目前处理到 f[qi]

        newF[0] = inf
        for i, x := range nums {
            for j := range a {
                a[j].and &= x
            }
            a = append(a, pair{x, i})

            // 原地去重
            j := 1
            for k := 1; k < len(a); k++ {
                if a[k].and != a[k-1].and {
                    a[j] = a[k]
                    j++
                }
            }
            a = a[:j]

            // 去掉无用数据
            for len(a) > 0 && a[0].and < target {
                a = a[1:]
            }

            // 上面这一大段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值
            // 下面是单调队列的滑窗过程

            if len(a) > 0 && a[0].and == target {
                // 现在 a[0].l 和 a[1].l-1 分别是子数组左端点的最小值和最大值
                r := i
                if len(a) > 1 {
                    r = a[1].l - 1
                }

                // 单调队列：右边入
                for ; qi <= r; qi++ {
                    for len(q) > 0 && f[qi] <= f[q[len(q)-1]] {
                        q = q[:len(q)-1]
                    }
                    q = append(q, qi)
                }

                // 单调队列：左边出
                for q[0] < a[0].l {
                    q = q[1:]
                }

                // 单调队列：计算答案
                newF[i+1] = f[q[0]] + x // 队首就是最小值
            } else {
                newF[i+1] = inf
            }
        }
        f, newF = newF, f
    }
    if f[n] < inf {
        return f[n]
    }
    return -1
}
```

### 写法二：原地计算+三指针

见 [讲解（方法二）](https://leetcode.cn/problems/number-of-subarrays-with-and-value-of-k/solution/jian-ji-xie-fa-o1-kong-jian-pythonjavacg-u7fv/)。

```py [sol-Python3]
class Solution:
    def minimumValueSum(self, nums: List[int], andValues: List[int]) -> int:
        n = len(nums)
        f = [0] + [inf] * n
        new_f = [0] * (n + 1)
        for target in andValues:
            a = nums.copy()  # 也可以写 nums[:]
            left = right = 0  # 子数组右端点为 i 时，子数组左端点的最小值和最大值+1
            q = deque()  # 单调队列，保存 f 的下标
            qi = 0  # 单调队列目前处理到 f[qi]

            new_f[0] = inf
            for i, x in enumerate(a):
                for j in range(i - 1, -1, -1):
                    if a[j] & x == a[j]:
                        break
                    a[j] &= x
                while left <= i and a[left] < target:
                    left += 1
                while right <= i and a[right] <= target:
                    right += 1

                # 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值+1
                # 下面是单调队列的滑窗过程

                if left < right:
                    # 单调队列：右边入
                    while qi < right:
                        while q and f[qi] <= f[q[-1]]:
                            q.pop()
                        q.append(qi)
                        qi += 1

                    # 单调队列：左边出
                    while q[0] < left:
                        q.popleft()

                    # 单调队列：计算答案
                    new_f[i + 1] = f[q[0]] + x  # 队首就是最小值
                else:
                    new_f[i + 1] = inf
            f, new_f = new_f, f
        return f[n] if f[n] < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumValueSum(int[] nums, int[] andValues) {
        final int INF = Integer.MAX_VALUE / 2;
        int n = nums.length;
        int[] f = new int[n + 1];
        Arrays.fill(f, 1, n + 1, INF);
        int[] newF = new int[n + 1];
        int[] q = new int[n + 1]; // 用数组模拟单调队列，保存 f 的下标

        for (int target : andValues) {
            int[] a = nums.clone();
            int left = 0, right = 0;
            int ql = 0, qr = 0; // q 的元素下标范围 [ql, qr)
            int qi = 0; // 单调队列目前处理到 f[qi]

            newF[0] = INF;
            for (int i = 0; i < n; i++) {
                int x = a[i];
                for (int j = i - 1; j >= 0 && (a[j] & x) != a[j]; j--) {
                    a[j] &= x;
                }
                while (left <= i && a[left] < target) {
                    left++;
                }
                while (right <= i && a[right] <= target) {
                    right++;
                }

                // 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值+1
                // 下面是单调队列的滑窗过程

                if (left < right) {
                    // 单调队列：右边入
                    for (; qi < right; qi++) {
                        while (qr > ql && f[qi] <= f[q[qr - 1]]) {
                            qr--;
                        }
                        q[qr++] = qi;
                    }

                    // 单调队列：左边出
                    while (ql < qr && q[ql] < left) {
                        ql++;
                    }

                    // 单调队列：计算答案
                    newF[i + 1] = f[q[ql]] + x; // 队首就是最小值
                } else {
                    newF[i + 1] = INF;
                }
            }
            int[] tmp = f;
            f = newF;
            newF = tmp;
        }
        return f[n] < INF ? f[n] : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumValueSum(vector<int>& nums, vector<int>& andValues) {
        const int INF = INT_MAX / 2;
        int n = nums.size();
        vector<int> f(n + 1, INF);
        vector<int> new_f(n + 1);

        f[0] = 0;
        for (int target : andValues) {
            auto a = nums;
            int left = 0, right = 0;
            deque<int> q; // 单调队列，保存 f 的下标
            int qi = 0; // 单调队列目前处理到 f[qi]

            new_f[0] = INF;
            for (int i = 0; i < n; i++) {
                int x = a[i];
                for (int j = i - 1; j >= 0 && (a[j] & x) != a[j]; j--) {
                    a[j] &= x;
                }
                while (left <= i && a[left] < target) {
                    left++;
                }
                while (right <= i && a[right] <= target) {
                    right++;
                }

                // 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值+1
                // 下面是单调队列的滑窗过程

                if (left < right) {
                    // 单调队列：右边入
                    for (; qi < right; qi++) {
                        while (!q.empty() && f[qi] <= f[q.back()]) {
                            q.pop_back();
                        }
                        q.push_back(qi);
                    }

                    // 单调队列：左边出
                    while (q.front() < left) {
                        q.pop_front();
                    }

                    // 单调队列：计算答案
                    new_f[i + 1] = f[q.front()] + x; // 队首就是最小值
                } else {
                    new_f[i + 1] = INF;
                }
            }
            swap(f, new_f);
        }
        return f[n] < INF ? f[n] : -1;
    }
};
```

```go [sol-Go]
func minimumValueSum(nums, andValues []int) int {
    const inf = math.MaxInt / 2
    n := len(nums)
    f := make([]int, n+1)
    for i := 1; i <= n; i++ {
        f[i] = inf
    }
    newF := make([]int, n+1)
    for _, target := range andValues {
        nums := slices.Clone(nums)
        left, right := 0, 0
        q := []int{} // 单调队列，保存 f 的下标
        qi := 0 // 单调队列目前处理到 f[qi]

        newF[0] = inf
        for i, x := range nums {
            for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
                nums[j] &= x
            }
            for left <= i && nums[left] < target {
                left++
            }
            for right <= i && nums[right] <= target {
                right++
            }

            // 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值+1
            // 下面是单调队列的滑窗过程

            if left < right {
                // 单调队列：右边入
                for ; qi < right; qi++ {
                    for len(q) > 0 && f[qi] <= f[q[len(q)-1]] {
                        q = q[:len(q)-1]
                    }
                    q = append(q, qi)
                }

                // 单调队列：左边出
                for q[0] < left {
                    q = q[1:]
                }

                // 单调队列：计算答案
                newF[i+1] = f[q[0]] + x // 队首就是最小值
            } else {
                newF[i+1] = inf
            }
        }
        f, newF = newF, f
    }
    if f[n] < inf {
        return f[n]
    }
    return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{andValues}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
