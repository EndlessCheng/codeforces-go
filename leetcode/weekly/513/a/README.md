枚举所有数对，计算强度，取最大值。

## 优化前

```py [sol-Python3]
class Solution:
    def maxPairStrength(self, nums: list[int]) -> int:
        ans = 0
        for i, x in enumerate(nums):
            for y in nums[:i]:  # 也可以用循环，避免切片
                g = gcd(x, y)
                ans = max(ans, x * y // (g * g))
        return ans
```

```java [sol-Java]
class Solution {
    public long maxPairStrength(int[] nums) {
        long ans = 0;
        for (int i = 0; i < nums.length; i++) {
            for (int j = 0; j < i; j++) {
                long g = gcd(nums[i], nums[j]);
                ans = Math.max(ans, (long) nums[i] * nums[j] / (g * g));
            }
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxPairStrength(vector<int>& nums) {
        long long ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            for (int j = 0; j < i; j++) {
                long long g = gcd(nums[i], nums[j]);
                ans = max(ans, 1LL * nums[i] * nums[j] / (g * g));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxPairStrength(nums []int) int64 {
	ans := 0
	for i, x := range nums {
		for _, y := range nums[:i] {
			g := gcd(x, y)
			ans = max(ans, x*y/(g*g))
		}
	}
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

## 最优性优化

设所有 $\textit{nums}[i]$ 的 GCD 为 $\textit{allGcd}$。

从大到小枚举数对，如果 $\dfrac{\textit{nums}[i]\cdot \textit{nums}[j]}{\textit{allGcd}^2}\le \textit{ans}$，则继续枚举不可能让答案变大，跳出循环。

```py [sol-Python3]
class Solution:
    def maxPairStrength(self, nums: list[int]) -> int:
        nums.sort(reverse=True)
        all_gcd2 = gcd(*nums) ** 2

        ans = 0
        for i, x in enumerate(nums):
            for j in range(i):
                y = nums[j]
                mul = x * y
                if mul // all_gcd2 <= ans:  # 最优性优化
                    break
                g = gcd(x, y)
                ans = max(ans, mul // (g * g))
        return ans
```

```java [sol-Java]
class Solution {
    public long maxPairStrength(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;

        int allGcd = 0;
        for (int x : nums) {
            allGcd = gcd(allGcd, x);
        }
        long allGcd2 = (long) allGcd * allGcd;

        long ans = 0;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j > i; j--) {
                long mul = (long) nums[i] * nums[j];
                if (mul / allGcd2 <= ans) { // 最优性优化
                    break;
                }
                long g = gcd(nums[i], nums[j]);
                ans = Math.max(ans, mul / (g * g));
            }
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxPairStrength(vector<int>& nums) {
        ranges::sort(nums, greater());

        int all_gcd = 0;
        for (int x : nums) {
            all_gcd = gcd(all_gcd, x);
        }
        long long all_gcd2 = 1LL * all_gcd * all_gcd;

        long long ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            for (int j = 0; j < i; j++) {
                long long mul = 1LL * nums[i] * nums[j];
                if (mul / all_gcd2 <= ans) { // 最优性优化
                    break;
                }
                long long g = gcd(nums[i], nums[j]);
                ans = max(ans, mul / (g * g));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxPairStrength(nums []int) int64 {
	slices.SortFunc(nums, func(a, b int) int { return b - a })

	allGcd := 0
	for _, x := range nums {
		allGcd = gcd(allGcd, x)
	}
	allGcd2 := allGcd * allGcd

	ans := 0
	for i, x := range nums {
		for _, y := range nums[:i] {
			mul := x * y
			if mul/allGcd2 <= ans {
				break
			}
			g := gcd(x, y)
			ans = max(ans, mul/(g*g))
		}
	}
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见下面数学题单的「**§1.6 最大公约数（GCD）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
