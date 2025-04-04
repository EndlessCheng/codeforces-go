## 方法一：0-1 背包

由于题目让我们选的是范围 $[l_i, r_i]$ 内的一个下标**子集**，所以每个 $\textit{nums}[i]$ 是**互相独立**的，可以分别计算。

选出包含 $i$ 的询问，设这些询问的 $\textit{val}$ 组成了数组 $\textit{vals}$，问题变成：

- 从 $\textit{vals}$ 的前缀中选一些数，元素和能否恰好等于 $\textit{nums}[i]$？

这是 0-1 背包。[416. 分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/) 是本题的简单版本。原理见[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

每个 $\textit{nums}[i]$ 算出的答案取最大值，即为最终答案。

注意特判 $\textit{nums}[i]=0$ 的情况，此时无需操作。

[本题视频讲解](https://www.bilibili.com/video/BV1JYQ8YWEvD/?t=21m27s)。

### 写法一：布尔数组

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        ans = 0
        for i, x in enumerate(nums):  # 每个 nums[i] 单独计算 0-1 背包
            if x == 0:
                continue
            f = [True] + [False] * x
            for k, (l, r, val) in enumerate(queries):
                if not l <= i <= r:
                    continue
                for j in range(x, val - 1, -1):
                    f[j] = f[j] or f[j - val]
                if f[x]:  # 满足要求
                    ans = max(ans, k + 1)
                    break
            else:  # 没有中途 break，说明无法满足要求
                return -1
        return ans
```

```java [sol-Java]
class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        int ans = 0;
        for (int i = 0; i < nums.length; i++) { // 每个 nums[i] 单独计算 0-1 背包
            int x = nums[i];
            if (x == 0) {
                continue;
            }
            boolean[] f = new boolean[x + 1];
            f[0] = true;
            for (int k = 0; k < queries.length; k++) {
                int[] q = queries[k];
                if (i < q[0] || i > q[1]) {
                    continue;
                }
                int val = q[2];
                for (int j = x; j >= val; j--) {
                    f[j] = f[j] || f[j - val];
                }
                if (f[x]) { // 满足要求
                    ans = Math.max(ans, k + 1);
                    break;
                }
            }
            if (!f[x]) { // 所有操作都执行完了也无法满足
                return -1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        int ans = 0;
        for (int i = 0; i < nums.size(); i++) { // 每个 nums[i] 单独计算 0-1 背包
            int x = nums[i];
            if (x == 0) {
                continue;
            }
            vector<int> f(x + 1);
            f[0] = true;
            for (int k = 0; k < queries.size(); k++) {
                auto& q = queries[k];
                if (i < q[0] || i > q[1]) {
                    continue;
                }
                int val = q[2];
                for (int j = x; j >= val; j--) {
                    f[j] = f[j] || f[j - val];
                }
                if (f[x]) { // 满足要求
                    ans = max(ans, k + 1);
                    break;
                }
            }
            if (!f[x]) { // 所有操作都执行完了也无法满足
                return -1;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minZeroArray(nums []int, queries [][]int) (ans int) {
	for i, x := range nums { // 每个 nums[i] 单独计算 0-1 背包
		if x == 0 {
			continue
		}
		f := make([]bool, x+1)
		f[0] = true
		for k, q := range queries {
			if i < q[0] || i > q[1] {
				continue
			}
			val := q[2]
			for j := x; j >= val; j-- {
				f[j] = f[j] || f[j-val]
			}
			if f[x] { // 满足要求
				ans = max(ans, k+1)
				break
			}
		}
		if !f[x] { // 所有操作都执行完了也无法满足
			return -1
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nqU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(U)$。

### 写法二：bitset

用 bitset（视作一个二进制数）代替布尔数组。二进制数从低到高第 $i$ 位是 $0$，表示布尔数组的第 $i$ 个数是 $\texttt{false}$；从低到高第 $i$ 位是 $1$，表示布尔数组的第 $i$ 个数是 $\texttt{true}$。

转移方程等价于，把 $f$ 中的每个比特位增加 $\textit{val}$，即左移 $\textit{val}$ 位，然后跟原来 $f$ 计算 OR。前者对应选择一个值为 $\textit{val}$ 的物品，后者对应不选。

判断 $f[x]$ 是否为 $\texttt{true}$，等价于判断 $f$ 的第 $x$ 位是否为 $1$，即 `(f >> x & 1) == 1`。

更多位运算技巧，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        ans = 0
        for i, x in enumerate(nums):
            if x == 0:
                continue
            f = 1
            for k, (l, r, val) in enumerate(queries):
                if not l <= i <= r:
                    continue
                f |= f << val
                if f >> x & 1:
                    ans = max(ans, k + 1)
                    break
            else:
                return -1
        return ans
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (x == 0) {
                continue;
            }
            BigInteger f = BigInteger.ONE;
            for (int k = 0; k < queries.length; k++) {
                int[] q = queries[k];
                if (i < q[0] || i > q[1]) {
                    continue;
                }
                f = f.or(f.shiftLeft(q[2]));
                if (f.testBit(x)) {
                    ans = Math.max(ans, k + 1);
                    break;
                }
            }
            if (!f.testBit(x)) {
                return -1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        int ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (x == 0) {
                continue;
            }
            bitset<10001> f;
            f.set(0);
            for (int k = 0; k < queries.size(); k++) {
                auto& q = queries[k];
                if (i < q[0] || i > q[1]) {
                    continue;
                }
                f |= f << q[2];
                if (f[x]) {
                    ans = max(ans, k + 1);
                    break;
                }
            }
            if (!f[x]) {
                return -1;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minZeroArray(nums []int, queries [][]int) (ans int) {
	p := new(big.Int)
	for i, x := range nums {
		if x == 0 {
			continue
		}
		f := big.NewInt(1)
		for k, q := range queries {
			if i < q[0] || i > q[1] {
				continue
			}
			f.Or(f, p.Lsh(f, uint(q[2])))
			if f.Bit(x) > 0 {
				ans = max(ans, k+1)
				break
			}
		}
		if f.Bit(x) == 0 {
			return -1
		}
	}
	return
}
```

#### 复杂度分析

以下分析，不考虑超出 $\textit{nums}[i]$ 的比特位。

- 时间复杂度：$\mathcal{O}(nqU / w)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度，$U=\max(\textit{nums})$，$w=32$ 或 $64$。
- 空间复杂度：$\mathcal{O}(U / w)$。

## 方法二：二分答案 + 多重背包 + 二进制优化

由于本题 $\textit{val}_i$ 很小，可以二分答案，统计每个 $\textit{val}_i$ 的出现次数，这样变成多重背包问题，可以用二进制优化。

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$-1$。一定无法满足要求。
- 开区间左端点初始值（优化）：$\left\lceil\dfrac{\max(\textit{nums})}{\max(\textit{val}_i)}\right\rceil-1$。
- 开区间右端点初始值：$m+1$，其中 $m$ 是 $\textit{queries}$ 的长度。如果二分结果为 $m+1$，那么返回 $-1$。

注 1：也可以对每个 $\textit{nums}[i]$ 单独二分，这样可以用到 [2861. 最大合金数](https://leetcode.cn/problems/maximum-number-of-alloys/) 的技巧，把当前答案作为二分的下界。（减一后是开区间的左端点）

注 2：方法二理论时间复杂度更优，但实际运行时间不如方法一。

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        max_val = max(q[2] for q in queries)

        def check(mx: int) -> bool:
            for idx, x in enumerate(nums):
                if x == 0:
                    continue
                cnt = [0] * (max_val + 1)
                for l, r, val in queries[:mx]:
                    if l <= idx <= r:
                        cnt[val] += 1
                # 多重背包（二进制优化）
                f = 1
                for v in range(1, max_val + 1):
                    num = cnt[v]
                    k1 = 1
                    while num:
                        k = min(k1, num)
                        f |= f << (v * k)  # 视作一个大小为 v*k 的物品
                        num -= k
                        k1 *= 2
                    if f >> x & 1:
                        break
                else:
                    return False
            return True

        left = (max(nums) + max_val - 1) // max_val
        ans = bisect_left(range(len(queries) + 1), True, lo=left, key=check)
        return ans if ans <= len(queries) else -1
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        int left = -1;
        int right = queries.length + 1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(mid, nums, queries)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right <= queries.length ? right : -1;
    }

    private boolean check(int mx, int[] nums, int[][] queries) {
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (x == 0) {
                continue;
            }
            int[] cnt = new int[11];
            for (int k = 0; k < mx; k++) {
                int[] q = queries[k];
                if (q[0] <= i && i <= q[1]) {
                    cnt[q[2]]++;
                }
            }
            // 多重背包（二进制优化）
            BigInteger f = BigInteger.ONE;
            for (int v = 1; v <= 10 && !f.testBit(x); v++) {
                int num = cnt[v];
                for (int pow2 = 1; num > 0 && !f.testBit(x); pow2 *= 2) {
                    int k = Math.min(pow2, num);
                    f = f.or(f.shiftLeft(v * k)); // 视作一个大小为 v*k 的物品
                    num -= k;
                }
            }
            if (!f.testBit(x)) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        auto check = [&](int mx) -> bool {
            for (int i = 0; i < nums.size(); i++) {
                int x = nums[i];
                if (x == 0) {
                    continue;
                }
                int cnt[11]{};
                for (int k = 0; k < mx; k++) {
                    auto& q = queries[k];
                    if (q[0] <= i && i <= q[1]) {
                        cnt[q[2]]++;
                    }
                }
                // 多重背包（二进制优化）
                bitset<10001> f;
                f.set(0);
                for (int v = 1; v <= 10 && !f[x]; v++) {
                    int num = cnt[v];
                    for (int pow2 = 1; num && !f[x]; pow2 *= 2) {
                        int k = min(pow2, num);
                        f |= f << (v * k); // 视作一个大小为 v*k 的物品
                        num -= k;
                    }
                }
                if (!f[x]) {
                    return false;
                }
            }
            return true;
        };

        int left = -1, right = queries.size() + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right <= queries.size() ? right : -1;
    }
};
```

```go [sol-Go]
func minZeroArray(nums []int, queries [][]int) int {
	ans := sort.Search(len(queries)+1, func(mx int) bool {
		p := new(big.Int)
	next:
		for i, x := range nums {
			if x == 0 {
				continue
			}
			cnt := [11]int{}
			for _, q := range queries[:mx] {
				if q[0] <= i && i <= q[1] {
					cnt[q[2]]++
				}
			}
			// 多重背包（二进制优化）
			f := big.NewInt(1)
			for v, num := range cnt {
				for pow2 := 1; num > 0; pow2 *= 2 {
					k := min(pow2, num)
					f.Or(f, p.Lsh(f, uint(v*k))) // 视作一个大小为 v*k 的物品
					if f.Bit(x) > 0 {
						continue next
					}
					num -= k
				}
			}
			return false
		}
		return true
	})
	if ans <= len(queries) {
		return ans
	}
	return -1
}
```

#### 复杂度分析

以下分析，不考虑超出 $\textit{nums}[i]$ 的比特位。

- 时间复杂度：$\mathcal{O}(n(q + V(U/w)\log q)\log q)$，其中 $q$ 是 $\textit{queries}$ 的长度，$n$ 是 $\textit{nums}$ 的长度，$V=\max(\textit{val}_i)$，$U=\max(\textit{nums})$，$w=32$ 或 $64$。
- 空间复杂度：$\mathcal{O}(V + U/w)$。

更多相似题目，见下面动态规划题单中的「**§3.1 0-1 背包**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
