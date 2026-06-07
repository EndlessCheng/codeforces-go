枚举 $\textit{selectedValue}$。

例如 $\textit{nums}=[2,3,3,12,5]$，现在枚举到 $\textit{selectedValue}=6$。

$\textit{nums}$ 中的与 $6$ **不互质**的数有 $4$ 个：$2,3,3,12$。讨论：

- 其中一个数改成 $6$。
- 其余数改成 $1$。因为 $1$ 和任何数都互质。
- 一共改 $4$ 次。数组变成 $[6,1,1,1,5]$，其中的 $6$ 与其他数都互质，满足要求。
- 得分为 $6-4 = 2$。

一般地，我们需要计算 $\textit{nums}$ 中的与 $\textit{selectedValue}$ 不互质的数的个数。

例如 $\textit{selectedValue}=6=2\times 3$：

- $2$ 的倍数都与 $6$ 不互质，在上面的例子中，这有 $2$ 个数 $\{2,12\}$。
- $3$ 的倍数都与 $6$ 不互质，在上面的例子中，这有 $3$ 个数 $\{3,3,12\}$。
- 这两个集合的**并集**为 $\{2,3,3,12\}$。并集的大小怎么算？如果直接计算 $2+3=5$，会多算一次交集 $\{12\}$ 的大小，所以要减去 $1$，得到并集的大小为 $2+3-1=4$。这个计算过程叫做**容斥原理**。

一般地，设 $\textit{nums}$ 中的是 $x$ 的倍数的个数为 $\text{cntMulti}(x)$，设 $\textit{selectedValue}$ 的质因子集合为 $P$，由容斥原理可得，$\textit{nums}$ 中的与 $\textit{selectedValue}$ 不互质的数的个数为

$$
\sum_{S\subseteq P} (-1)^{|S|} \cdot \text{cntMulti}\left(\prod S\right)
$$

用莫比乌斯函数表示，即

$$
\sum_{d|\textit{selectedValue}} \mu(d) \cdot \text{cntMulti}(d)
$$

其中莫比乌斯函数为

$$
\mu(n)=
\begin{cases}
1, & n=1\\
(-1)^k, & n = p_1p_2\dots p_k\\
0, & n\ 有大于\ 1\ 的平方因子
\end{cases}
$$

> 预处理 $\mu$ 以及每个数的（不含平方因子的）因子列表，可以加速上述计算过程。

最后，我们要快速计算 $\text{cntMulti}(d)$。可以预处理：先统计 $\textit{nums}$ 的每个数的出现次数 $\textit{cnt}[x]$，然后枚举 $x$ 以及 $x$ 的倍数 $y$，累加 $\textit{cnt}[y]$，即为 $\text{cntMulti}(x)$。

**特殊情况**：

1. 如果 $\textit{nums}$ 包含 $\textit{selectedValue}$，可以少改一次。
2. 如果 $\textit{nums}$ 的每个数都与 $\textit{selectedValue}$ 互质，但 $\textit{nums}$ 中没有 $\textit{selectedValue}$，那么必须把一个数改成 $\textit{selectedValue}$。
3. 单独处理 $\textit{selectedValue}=1$ 的情况。如果 $\textit{nums}$ 中有 $1$，那么无需修改，得分为 $1-0=1$；否则要把一个数改成 $1$，得分为 $1-1=0$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
MX = 100_001

# 预处理莫比乌斯函数
# 当 n > 1 时，sum_{d|n} mu[d] = 0
# 所以 mu[n] = -sum_{d|n ∧ d<n} mu[d]
mu = [0] * MX
mu[1] = 1
for i in range(1, MX):
    for j in range(i * 2, MX, i):
        mu[j] -= mu[i]  # i 是 j 的真因子

# 预处理不含平方因子的因子列表
# 本题不需要因子 1
divisors = [[] for _ in range(MX)]
for i in range(2, MX):
    if mu[i]:
        for j in range(i, MX, i):
            divisors[j].append(i)  # i 是 j 的因子，且 mu[i] != 0


class Solution:
    def maxScore(self, nums: list[int], maxVal: int) -> int:
        max_num = max(nums)
        cnt = [0] * (max_num + 1)
        for x in nums:
            cnt[x] += 1

        cnt_multi = [0] * (max_num + 1)
        for i in range(2, max_num + 1):
            for j in range(i, max_num + 1, i):
                cnt_multi[i] += cnt[j]  # 统计 nums 中有多少个数是 i 的倍数

        # 单独计算 selected_value = 1 时的得分
        ans = 1 if cnt[1] > 0 else 0

        # 从大到小枚举 selected_value
        for selected_value in range(max(max_num, maxVal), 1, -1):
            # 优化：如果 selected_value <= ans，那么 ans 不会变大，跳出循环
            if selected_value <= ans:
                break

            if selected_value > maxVal and cnt[selected_value] == 0:
                continue  # 无法改成 selected_value

            # 与 selected_value 不互质的数，其中一个数改成 selected_value，其余数都改成 1
            cost = 0
            for d in divisors[selected_value]:
                if d > max_num:
                    break
                cost -= mu[d] * cnt_multi[d]

            if selected_value <= max_num and cnt[selected_value] > 0:
                cost -= 1  # 如果某个 nums[i] 恰好等于 selected_value，可以少改一次
            elif cost == 0:
                cost = 1  # 至少要有一个数改成 selected_value

            ans = max(ans, selected_value - cost)

        return ans
```

```java [sol-Java]
class Solution {
    private static final int MX = 100_001;
    private static final int[] mu = new int[MX];
    private static final List<Integer>[] divisors = new ArrayList[MX];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理莫比乌斯函数
        // 当 n > 1 时，sum_{d|n} mu[d] = 0
        // 所以 mu[n] = -sum_{d|n ∧ d<n} mu[d]
        mu[1] = 1;
        for (int i = 1; i < MX; i++) {
            for (int j = i * 2; j < MX; j += i) {
                mu[j] -= mu[i]; // i 是 j 的真因子
            }
        }

        // 预处理不含平方因子的因子列表
        // 本题不需要因子 1
        Arrays.setAll(divisors, _ -> new ArrayList<>());
        for (int i = 2; i < MX; i++) {
            if (mu[i] == 0) {
                continue;
            }
            for (int j = i; j < MX; j += i) {
                divisors[j].add(i); // i 是 j 的因子，且 mu[i] != 0
            }
        }
    }

    public int maxScore(int[] nums, int maxVal) {
        int maxNum = 0;
        for (int x : nums) {
            maxNum = Math.max(maxNum, x);
        }

        int[] cnt = new int[maxNum + 1];
        for (int x : nums) {
            cnt[x]++;
        }

        int[] cntMulti = new int[maxNum + 1];
        for (int i = 2; i <= maxNum; i++) {
            for (int j = i; j <= maxNum; j += i) {
                cntMulti[i] += cnt[j]; // 统计 nums 中有多少个数是 i 的倍数
            }
        }

        // 单独计算 selectedValue = 1 时的得分
        int ans = cnt[1] > 0 ? 1 : 0;

        // 从大到小枚举 selectedValue
        // 优化：如果 selectedValue <= ans，那么 ans 不会变大，跳出循环
        for (int selectedValue = Math.max(maxNum, maxVal); selectedValue > ans; selectedValue--) {
            if (selectedValue > maxVal && cnt[selectedValue] == 0) {
                continue; // 无法改成 selectedValue
            }

            // 与 selectedValue 不互质的数，其中一个数改成 selectedValue，其余数都改成 1
            int cost = 0;
            for (int d : divisors[selectedValue]) {
                if (d > maxNum) {
                    break;
                }
                cost -= mu[d] * cntMulti[d];
            }

            if (selectedValue <= maxNum && cnt[selectedValue] > 0) {
                cost--; // 如果某个 nums[i] 恰好等于 selectedValue，可以少改一次
            } else if (cost == 0) {
                cost = 1; // 至少要有一个数改成 selectedValue
            }

            ans = Math.max(ans, selectedValue - cost);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
constexpr int MX = 100'001;
int8_t mu[MX];
vector<int> divisors[MX];

int init = [] {
    // 预处理莫比乌斯函数
    // 当 n > 1 时，sum_{d|n} mu[d] = 0
    // 所以 mu[n] = -sum_{d|n ∧ d<n} mu[d]
    mu[1] = 1;
    for (int i = 1; i < MX; i++) {
        for (int j = i * 2; j < MX; j += i) {
            mu[j] -= mu[i]; // i 是 j 的真因子
        }
    }

    // 预处理不含平方因子的因子列表
    // 本题不需要因子 1
    for (int i = 2; i < MX; i++) {
        if (mu[i]) {
            for (int j = i; j < MX; j += i) {
                divisors[j].push_back(i); // i 是 j 的因子，且 mu[i] != 0
            }
        }
    }
    return 0;
}();

class Solution {
public:
    int maxScore(vector<int>& nums, int maxVal) {
        int max_num = ranges::max(nums);
        vector<int> cnt(max_num + 1);
        for (int x : nums) {
            cnt[x]++;
        }
        vector<int> cnt_multi(max_num + 1);
        for (int i = 2; i <= max_num; i++) {
            for (int j = i; j <= max_num; j += i) {
                cnt_multi[i] += cnt[j]; // 统计 nums 中有多少个数是 i 的倍数
            }
        }

        // 单独计算 selected_value = 1 时的得分
        int ans = cnt[1] > 0;

        // 从大到小枚举 selected_value
        // 优化：如果 selected_value <= ans，那么 ans 不会变大，跳出循环
        for (int selected_value = max(max_num, maxVal); selected_value > ans; selected_value--) {
            if (selected_value > maxVal && cnt[selected_value] == 0) {
                continue; // 无法改成 selected_value
            }

            // 与 selected_value 不互质的数，其中一个数改成 selected_value，其余数都改成 1
            int cost = 0;
            for (int d : divisors[selected_value]) {
                if (d > max_num) {
                    break;
                }
                cost -= mu[d] * cnt_multi[d];
            }

            if (selected_value <= max_num && cnt[selected_value] > 0) {
                cost--; // 如果某个 nums[i] 恰好等于 selected_value，可以少改一次
            } else if (cost == 0) {
                cost = 1; // 至少要有一个数改成 selected_value
            }

            ans = max(ans, selected_value - cost);
        }

        return ans;
    }
};
```

```go [sol-Go]
const MX = 100_001
var mu = [MX]int{1: 1}     // 莫比乌斯函数
var divisors = [MX][]int{} // 不含平方因子的因子列表，用于容斥

func init() {
	// 预处理莫比乌斯函数
	// 当 n > 1 时，sum_{d|n} mu[d] = 0
	// 所以 mu[n] = -sum_{d|n ∧ d<n} mu[d]
	for i := 1; i < MX; i++ {
		for j := i * 2; j < MX; j += i {
			mu[j] -= mu[i] // i 是 j 的真因子
		}
	}

	// 预处理不含平方因子的因子列表
	// 本题不需要因子 1
	for i := 2; i < MX; i++ {
		if mu[i] == 0 {
			continue
		}
		for j := i; j < MX; j += i {
			divisors[j] = append(divisors[j], i) // i 是 j 的因子，且 mu[i] != 0
		}
	}
}

func maxScore(nums []int, maxVal int) (ans int) {
	maxNum := slices.Max(nums)
	cnt := make([]int, maxNum+1)
	for _, x := range nums {
		cnt[x]++
	}
	cntMulti := make([]int, maxNum+1)
	for i := 2; i <= maxNum; i++ {
		for j := i; j <= maxNum; j += i {
			cntMulti[i] += cnt[j] // 统计 nums 中有多少个数是 i 的倍数
		}
	}

	if cnt[1] > 0 {
		ans = 1 // selectedValue = 1 时，无需修改，得分为 1
	}

	// 从大到小枚举 selectedValue
	// 优化：如果 selectedValue <= ans，那么 ans 不会变大，跳出循环
	for selectedValue := max(maxNum, maxVal); selectedValue > ans; selectedValue-- {
		if selectedValue > maxVal && cnt[selectedValue] == 0 {
			continue // 无法改成 selectedValue
		}

		// 与 selectedValue 不互质的数，其中一个数改成 selectedValue，其余数都改成 1
		cost := 0
		for _, d := range divisors[selectedValue] {
			if d > maxNum {
				break
			}
			cost -= mu[d] * cntMulti[d]
		}

		if selectedValue <= maxNum && cnt[selectedValue] > 0 {
			cost-- // 如果某个 nums[i] 恰好等于 selectedValue，可以少改一次
		} else if cost == 0 {
			cost = 1 // 至少要有一个数改成 selectedValue
		}

		ans = max(ans, selectedValue-cost)
	}

	return
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n + U\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\max(\textit{nums}), \textit{maxVal})$。由调和级数可知，预处理 $\textit{cntMulti}$ 数组的时间复杂度为 $\mathcal{O}(U\log U)$。由调和级数可知，不超过 $U$ 的因子总数为 $\mathcal{O}(U\log U)$。特别地，不含平方因子的因子总数也是 $\mathcal{O}(U\log U)$。
- 空间复杂度：$\mathcal{O}(U)$。

## 专题训练

见下面数学题单的「**§2.4 容斥原理**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
