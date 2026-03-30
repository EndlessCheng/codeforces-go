## 方法一：前缀和优化 DP

设 $x = \textit{arr}[n-1]$，那么 $y = \textit{arr}[n-2]$ 必须 $\le x$。枚举所有满足数位和等于 $\textit{digitSum}[n-2]$ 的 $y$，问题变成：

- 在 $\textit{digitSum}$ 的 $[0,n-2]$ 中，以 $y$ 结尾的有效数组的个数。

这是一个规模更小的子问题。

定义 $f[i+1][x]$ 表示在 $\textit{digitSum}$ 的 $[0,i]$ 中，以 $x$ 结尾的有效数组的个数。

枚举所有满足数位和等于 $\textit{digitSum}[i-1]$ 的 $y$，问题变成在 $\textit{digitSum}$ 的 $[0,i-1]$ 中，以 $y$ 结尾的有效数组的个数，即 $f[i][y]$。

累加得

$$
f[i+1][x] = \sum_{y\le x} f[i][y]
$$

上式可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化。

初始值：$f[0][0] = 1$。

答案：$\sum\limits_{x} f[n][x]$。

代码实现时，注意取模。为什么可以在**中途取模**？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1dxXSBAE6F/?t=14m42s)，欢迎点赞关注~

```py [sol-Python3]
MOD = 1_000_000_007
MX = 5001
MAX_DIGIT_SUM = 31  # 4999 的数位和最大
dig_sum = [0] * MX

# 预处理数位和
for x in range(MX):
    # 去掉 x 的个位，问题变成 x // 10 的数位和，即 dig_sum[x // 10]
    dig_sum[x] = dig_sum[x // 10] + x % 10

class Solution:
    def countArrays(self, digitSum: List[int]) -> int:
        s = [1] * MX  # f 的前缀和
        for ds in digitSum:
            if ds > MAX_DIGIT_SUM:
                return 0
            for x in range(MX):
                # 如果 dig_sum[x] != ds，那么 f[x] = 0，否则 f[x] = s[x]
                # 把 f[x] 的值填到 s[x] 中，那么只需要把 dig_sum[x] != ds 的 s[x] 置为 0
                if dig_sum[x] != ds:
                    s[x] = 0
                if x > 0:
                    s[x] = (s[x] + s[x - 1]) % MOD
        return s[-1]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 5001;
    private static final int MAX_DIGIT_SUM = 31; // 4999 的数位和最大
    private static final int[] digSum = new int[MX];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理数位和
        for (int x = 0; x < MX; x++) {
            // 去掉 x 的个位，问题变成 x / 10 的数位和，即 digSum[x / 10]
            digSum[x] = digSum[x / 10] + x % 10;
        }
    }

    public int countArrays(int[] digitSum) {
        int[] sum = new int[MX]; // f 的前缀和
        Arrays.fill(sum, 1);
        for (int ds : digitSum) {
            if (ds > MAX_DIGIT_SUM) {
                return 0;
            }
            for (int x = 0; x < MX; x++) {
                // 如果 digSum[x] != ds，那么 f[x] = 0，否则 f[x] = sum[x]
                // 把 f[x] 的值填到 sum[x] 中，那么只需要把 digSum[x] != ds 的 sum[x] 置为 0
                if (digSum[x] != ds) {
                    sum[x] = 0;
                }
                if (x > 0) {
                    sum[x] = (sum[x] + sum[x - 1]) % MOD;
                }
            }
        }
        return sum[MX - 1];
    }
}
```

```cpp [sol-C++]
static constexpr int MX = 5001;
static constexpr int MAX_DIGIT_SUM = 31; // 4999 的数位和最大
int dig_sum[MX];

int init = [] {
    // 预处理数位和
    for (int x = 0; x < MX; x++) {
        // 去掉 x 的个位，问题变成 x / 10 的数位和，即 dig_sum[x / 10]
        dig_sum[x] = dig_sum[x / 10] + x % 10;
    }
    return 0;
}();

class Solution {
public:
    int countArrays(vector<int>& digitSum) {
        constexpr int MOD = 1'000'000'007;
        vector<int> sum(MX, 1); // f 的前缀和
        for (int ds : digitSum) {
            if (ds > MAX_DIGIT_SUM) {
                return 0;
            }
            for (int x = 0; x < MX; x++) {
                // 如果 dig_sum[x] != ds，那么 f[x] = 0，否则 f[x] = sum[x]
                // 把 f[x] 的值填到 sum[x] 中，那么只需要把 dig_sum[x] != ds 的 sum[x] 置为 0
                if (dig_sum[x] != ds) {
                    sum[x] = 0;
                }
                if (x > 0) {
                    sum[x] = (sum[x] + sum[x - 1]) % MOD;
                }
            }
        }
        return sum[MX - 1];
    }
};
```

```go [sol-Go]
const mx = 5001
const maxDigitSum = 31 // 4999 的数位和最大
var digSum [mx]int

func init() {
	// 预处理数位和
	for x := range digSum {
		// 去掉 x 的个位，问题变成 x/10 的数位和，即 digSum[x/10]
		digSum[x] = digSum[x/10] + x%10
	}
}

func countArrays(digitSum []int) int {
	const mod = 1_000_000_007
	sum := [mx]int{} // f 的前缀和
	for i := range sum {
		sum[i] = 1
	}
	for _, ds := range digitSum {
		if ds > maxDigitSum {
			return 0
		}
		for x := range mx {
			// 如果 digSum[x] != ds，那么 f[x] = 0，否则 f[x] = sum[x]
			// 把 f[x] 的值填到 sum[x] 中，那么只需要把 digSum[x] != ds 的 sum[x] 置为 0
			if digSum[x] != ds {
				sum[x] = 0
			}
			if x > 0 {
				sum[x] = (sum[x] + sum[x-1]) % mod
			}
		}
	}
	return sum[mx-1]
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 是 $\textit{digitSum}$ 的长度，$U=5000$。
- 空间复杂度：$\mathcal{O}(U)$。

## 方法二：双指针优化 DP

示例 1 的 $\textit{digitSum} = [25,1]$：

- 数位和为 $25$ 的数字为 $a=[799,889,898,979,988,997]$。
- 数位和为 $1$ 的数字为 $b=[1,10,100,1000]$。

我们其实只需考虑数组 $a$ 和 $b$ 中的数字如何转移，无需考虑 $[0,5000]$ 中的所有数字，从而减少计算量。

具体地，枚举 $b$ 中的数字 $x$，对于 $a$ 中所有 $\le x$ 的数字 $y$，我们需要知道，以 $y$ 结尾的有效数组的个数 $f[y]$ 是多少，从而算出以 $x$ 结尾的有效数组的个数 $f[x]$，即

$$
f[x] = \sum_{y\in a\ \wedge\ y\le x} f[y]
$$

如果 $a$ 和 $b$ 都是递增的，那么上述过程可以用双指针优化，具体见代码。

```py [sol-Python3]
MOD = 1_000_000_007
MX = 5001
MAX_DIGIT_SUM = 31  # 4999 的数位和最大
sum_to_nums = [[] for _ in range(MAX_DIGIT_SUM + 1)]
dig_sum = [0] * MX
for x in range(MX):
    # 去掉 x 的个位，问题变成 x // 10 的数位和，即 dig_sum[x // 10]
    dig_sum[x] = dig_sum[x // 10] + x % 10
    sum_to_nums[dig_sum[x]].append(x)

class Solution:
    def countArrays(self, digitSum: List[int]) -> int:
        f = [0] * MX  # f[x] 表示以 x 结尾的有效数组的个数
        f[0] = 1
        pre = 0

        for cur in digitSum:
            if cur > MAX_DIGIT_SUM:
                return 0
            a = sum_to_nums[pre]
            j, m = 0, len(a)
            s = 0
            for x in sum_to_nums[cur]:
                # 有效数组的前一个数只要 <= x 就行
                while j < m and a[j] <= x:
                    s += f[a[j]]
                    j += 1
                # s 现在就是以 x 结尾的有效数组的个数
                f[x] = s % MOD
            pre = cur  # 记录上一个数位和

        return sum(f[x] for x in sum_to_nums[pre]) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 5001;
    private static final int MAX_DIGIT_SUM = 31; // 4999 的数位和最大
    private static final List<Integer>[] sumToNums = new ArrayList[MAX_DIGIT_SUM + 1];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.setAll(sumToNums, _ -> new ArrayList<>());
        int[] digSum = new int[MX];
        for (int x = 0; x < MX; x++) {
            // 去掉 x 的个位，问题变成 x / 10 的数位和，即 digSum[x / 10]
            digSum[x] = digSum[x / 10] + x % 10;
            sumToNums[digSum[x]].add(x);
        }
    }

    public int countArrays(int[] digitSum) {
        int[] f = new int[MX]; // f[x] 表示以 x 结尾的有效数组的个数
        f[0] = 1;
        int pre = 0;

        for (int cur : digitSum) {
            if (cur > MAX_DIGIT_SUM) {
                return 0;
            }
            List<Integer> a = sumToNums[pre];
            int j = 0, m = a.size();
            int sum = 0;
            for (int x : sumToNums[cur]) {
                // 有效数组的前一个数只要 <= x 就行
                for (; j < m && a.get(j) <= x; j++) {
                    sum = (sum + f[a.get(j)]) % MOD;
                }
                // sum 现在就是以 x 结尾的有效数组的个数
                f[x] = sum;
            }
            pre = cur; // 记录上一个数位和
        }

        long ans = 0;
        for (int x : sumToNums[pre]) {
            ans += f[x];
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
static constexpr int MX = 5001;
static constexpr int MAX_DIGIT_SUM = 31; // 4999 的数位和最大
vector<int> sum_to_nums[MAX_DIGIT_SUM + 1];

int init = [] {
    int dig_sum[MX]{};
    for (int x = 0; x < MX; x++) {
        // 去掉 x 的个位，问题变成 x / 10 的数位和，即 dig_sum[x / 10]
        dig_sum[x] = dig_sum[x / 10] + x % 10;
        sum_to_nums[dig_sum[x]].push_back(x);
    }
    return 0;
}();

class Solution {
public:
    int countArrays(vector<int>& digitSum) {
        constexpr int MOD = 1'000'000'007;
        // f[x] 表示以 x 结尾的有效数组的个数
        int f[MX] = {1}; // f[0] = 1，其余 f[x] = 0 
        int pre = 0;

        for (int cur : digitSum) {
            if (cur > MAX_DIGIT_SUM) {
                return 0;
            }
            auto& a = sum_to_nums[pre];
            int j = 0, m = a.size();
            int sum = 0;
            for (int x : sum_to_nums[cur]) {
                // 有效数组的前一个数只要 <= x 就行
                for (; j < m && a[j] <= x; j++) {
                    sum = (sum + f[a[j]]) % MOD;
                }
                // sum 现在就是以 x 结尾的有效数组的个数
                f[x] = sum;
            }
            pre = cur; // 记录上一个数位和
        }

        long long ans = 0;
        for (int x : sum_to_nums[pre]) {
            ans += f[x];
        }
        return ans % MOD;
    }
};
```

```go [sol-Go]
const mx = 5001
const maxDigitSum = 31 // 4999 的数位和最大
var sumToNums [maxDigitSum + 1][]int

func init() {
	digSum := [mx]int{}
	for x := range digSum {
		// 去掉 x 的个位，问题变成 x/10 的数位和，即 digSum[x/10]
		digSum[x] = digSum[x/10] + x%10
		sumToNums[digSum[x]] = append(sumToNums[digSum[x]], x)
	}
}

func countArrays(digitSum []int) (ans int) {
	const mod = 1_000_000_007
	f := [mx]int{1} // f[x] 表示以 x 结尾的有效数组的个数
	pre := 0

	for _, cur := range digitSum {
		if cur > maxDigitSum {
			return 0
		}
		a := sumToNums[pre]
		j, m := 0, len(a)
		sum := 0
		for _, x := range sumToNums[cur] {
			// 有效数组的前一个数只要 <= x 就行
			for ; j < m && a[j] <= x; j++ {
				sum += f[a[j]]
			}
			// sum 现在就是以 x 结尾的有效数组的个数
			f[x] = sum % mod
		}
		pre = cur // 记录上一个数位和
	}

	for _, x := range sumToNums[pre] {
		ans += f[x]
	}
	return ans % mod
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(nV)$，其中 $n$ 是 $\textit{digitSum}$ 的长度，最多有 $V=365$ 个数的数位和相同。
- 空间复杂度：$\mathcal{O}(U)$，其中 $U = 5000$。

## 专题训练

1. 动态规划题单的「**§11.1 前缀和优化 DP**」。
2. 双指针题单的「**四、双序列双指针**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
