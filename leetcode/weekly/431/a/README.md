## 方法一：暴力

暴力枚举所有子数组（枚举左右端点），计算子数组乘积、子数组 $\text{lcm}$ 与子数组 $\gcd$。

子数组乘积太大会溢出，如何处理？由于子数组 $\text{lcm}$ 不超过 $\text{lcm}(\textit{nums})$，子数组 $\gcd$ 不超过 $\max(\textit{nums})$，我们可以枚举到子数组乘积超过 $\text{lcm}(\textit{nums})\cdot \max(\textit{nums})$ 为止。

> **注 1**：题目保证 $n\ge 2$，又由于两个数一定满足题目要求（$xy = \gcd(x,y)\cdot\text{lcm}(x,y)$ 恒成立），所以合法子数组是一定存在的，且长度至少是 $2$。
> 
> **注 2**：$\text{lcm}(1,2,\ldots,10)=2520$。

[本题视频讲解](https://www.bilibili.com/video/BV18srKYLEd8/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxLength(self, nums: List[int]) -> int:
        max_m = lcm(*nums) * max(nums)
        ans = 0
        for i in range(len(nums)):
            m, l, g = 1, 1, 0
            for j in range(i, len(nums)):
                x = nums[j]
                m *= x
                l = lcm(l, x)
                g = gcd(g, x)
                if m == l * g:
                    ans = max(ans, j - i + 1)
                if m > max_m:
                    break
        return ans
```

```java [sol-Java]
class Solution {
    public int maxLength(int[] nums) {
        int mx = Arrays.stream(nums).max().getAsInt();
        int allLcm = 1;
        for (int x : nums) {
            allLcm = lcm(allLcm, x);
        }

        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int m = 1;
            int l = 1;
            int g = 0;
            for (int j = i; j < nums.length && m <= allLcm * mx; j++) {
                int x = nums[j];
                m *= x;
                l = lcm(l, x);
                g = gcd(g, x);
                if (m == l * g) {
                    ans = Math.max(ans, j - i + 1);
                }
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

    private int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxLength(vector<int>& nums) {
        int mx = ranges::max(nums);
        int all_lcm = 1;
        for (int x : nums) {
            all_lcm = lcm(all_lcm, x);
        }

        int ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            int m = 1, l = 1, g = 0;
            for (int j = i; j < nums.size() && m <= all_lcm * mx; j++) {
                int x = nums[j];
                m *= x;
                l = lcm(l, x);
                g = gcd(g, x);
                if (m == l * g) {
                    ans = max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxLength(nums []int) (ans int) {
	mx := slices.Max(nums)
	allLcm := 1
	for _, x := range nums {
		allLcm = lcm(allLcm, x)
	}

	for i := range nums {
		m, l, g := 1, 1, 0
		for j := i; j < len(nums) && m <= allLcm*mx; j++ {
			x := nums[j]
			m *= x
			l = lcm(l, x)
			g = gcd(g, x)
			if m == l*g {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：滑动窗口

设子数组的长度为 $k$。

由于 $xy = \gcd(x,y)\cdot\text{lcm}(x,y)$ 恒成立，$k=2$ 的子数组一定满足要求，所以下面只讨论 $k\ge 3$ 的子数组。

考虑乘积、$\text{lcm}$ 和 $\gcd$ 的质因数分解中的一个质数 $p$ 的指数 $E$：

- 在子数组乘积中，$E$ 为各个元素的 $p$ 的指数 $e_i$ 之和，即 $E=e_1+e_2+\cdots+e_k$。
- 在子数组 $\gcd$ 中，$E$ 为各个元素的 $p$ 的指数 $e_i$ 的最小值，即 $E=\min(e_1,e_2,\ldots,e_k)$。
- 在子数组 $\text{lcm}$ 中，$E$ 为各个元素的 $p$ 的指数 $e_i$ 的最大值，即 $E=\max(e_1,e_2,\ldots,e_k)$。

所以本题的要求等价于，对于每个质数 $p$，都有

$$
e_1+e_2+\cdots+e_k = \min(e_1,e_2,\ldots,e_k) + \max(e_1,e_2,\ldots,e_k) 
$$

为方便分析，把这些 $e_i$ 从小到大排序，那么上式为

$$
e_1+e_2+\cdots+e_k = e_1 + e_k
$$

即

$$
e_2+e_3+\cdots+e_{k-1} = 0
$$

由于 $e_i\ge 0$，解得

$$
e_2 = e_3 = \cdots = e_{k-1} = 0
$$

由于 $e$ 已排序，所以 $e_1\le e_2 = 0$，所以 $e_1 = 0$。

所以只有 $e_k\ge 0$，即排序之前只有一个 $e_i\ge 0$，其余均为 $0$。这说明在 $k\ge 3$ 的情况下，子数组不可能有两个数都有质因子 $p$，所以子数组这 $k$ 个数必须**两两互质**。

两两互质是一个很强的性质，我们可以把子数组中的所有数「压缩」到子数组的乘积 $\textit{mul}$ 中，计算 $\textit{mul}$ 与即将加入子数组的数 $x$ 的 $\gcd$，如果 $\gcd(\textit{mul},x)>1$，则说明 $x$ 与子数组中的某个数不是互质的。

类似 [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) 的 [滑动窗口做法](https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/1959540/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-iaks/)，先去掉子数组中的与 $x$ 不互质的数，再把 $x$ 加到子数组中。

```py [sol-Python3]
class Solution:
    def maxLength(self, nums: List[int]) -> int:
        ans = 2
        mul = 1
        left = 0
        for right, x in enumerate(nums):
            while gcd(mul, x) > 1:
                mul //= nums[left]
                left += 1
            mul *= x
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxLength(int[] nums) {
        int ans = 2;
        int mul = 1;
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            while (gcd(mul, nums[right]) > 1) {
                mul /= nums[left];
                left++;
            }
            mul *= nums[right];
            ans = Math.max(ans, right - left + 1);
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
    int maxLength(vector<int>& nums) {
        int ans = 2, mul = 1, left = 0;
        for (int right = 0; right < nums.size(); right++) {
            while (gcd(mul, nums[right]) > 1) {
                mul /= nums[left];
                left++;
            }
            mul *= nums[right];
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxLength(nums []int) int {
	ans, mul, left := 2, 1, 0
	for right, x := range nums {
		for gcd(mul, x) > 1 {
			mul /= nums[left]
			left++
		}
		mul *= x
		ans = max(ans, right-left+1)
	}
	return ans
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

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
