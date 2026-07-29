## 方法一：前后缀分解

定义 $\textit{sum}[i]$ 表示 $[0,i-1]$ 的元素和（前缀和）。特别地，$\textit{sum}[0] = 0$。

定义 $\textit{mul}[i]$ 表示 $[i+1,n-1]$ 的元素积（后缀积）。特别地，$\textit{mul}[n-1] = 1$。

问题变成：

- 寻找满足 $\textit{sum}[i] = \textit{mul}[i]$ 的最小的下标 $i$。

由于 $\textit{nums}[i] > 0$，所以 $\textit{sum}$ 是严格递增的，$\textit{mul}$ 是（非严格）递减的。画出函数图像的话，至多有一个交点。

所以本题要么无解，要么恰好有一个解。

代码实现时，可以先把 $\textit{sum}$ 算出来，然后倒着遍历 $\textit{nums}$，同时计算 $\textit{mul}$。这样做，$\textit{mul}$ 可以简化成一个变量。

此外，当 $i=0$ 时，$\textit{sum}[0] = 0$，$\textit{mul}[0]>0$，所以 $\textit{sum}[0] \ne \textit{mul}[0]$。所以第二个循环只需循环到 $i=1$。

[本题视频讲解](https://www.bilibili.com/video/BV1H6NMzdEbo/)，欢迎点赞关注~

### 优化前

```py [sol-Python3]
class Solution:
    def smallestBalancedIndex(self, nums: list[int]) -> int:
        pre = list(accumulate(nums, initial=0))  # pre[i] 表示 [0,i-1] 之和
        mul = 1  # [i+1,n-1] 之积
        for i in range(len(nums) - 1, 0, -1):
            # 如果 pre[i] < mul，那么继续向左遍历，mul 越来越大（或者不变），pre 越来越小，不可能找到答案
            if pre[i] < mul:
                break
            if pre[i] == mul:  # [0,i-1] 之和等于 [i+1,n-1] 之积
                return i
            mul *= nums[i]
        return -1
```

```java [sol-Java]
class Solution {
    public int smallestBalancedIndex(int[] nums) {
        int n = nums.length;
        long[] sum = new long[n]; // sum[i] 表示 [0,i-1] 之和
        for (int i = 0; i < n - 1; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        long mul = 1; // [i+1,n-1] 之积
        for (int i = n - 1; i > 0; i--) {
            if (sum[i] == mul) { // [0,i-1] 之和等于 [i+1,n-1] 之积
                return i;
            }
            // 如果 mul * nums[i] > sum[i-1]，那么继续向左遍历，mul 越来越大（或者不变），sum 越来越小，不可能找到答案
            // 为避免乘法溢出，改成等价的除法
            if (mul > sum[i - 1] / nums[i]) {
                break;
            }
            mul *= nums[i];
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestBalancedIndex(vector<int>& nums) {
        int n = nums.size();
        vector<long long> sum(n); // sum[i] 表示 [0,i-1] 之和
        for (int i = 0; i < n - 1; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        long long mul = 1; // [i+1,n-1] 之积
        for (int i = n - 1; i > 0; i--) {
            if (sum[i] == mul) { // [0,i-1] 之和等于 [i+1,n-1] 之积
                return i;
            }
            // 如果 mul * nums[i] > sum[i-1]，那么继续向左遍历，mul 越来越大（或者不变），sum 越来越小，不可能找到答案
            // 为避免乘法溢出，改成等价的除法
            if (mul > sum[i - 1] / nums[i]) {
                break;
            }
            mul *= nums[i];
        }
        return -1;
    }
};
```

```go [sol-Go]
func smallestBalancedIndex(nums []int) int {
	n := len(nums)
	sum := make([]int, n) // sum[i] 表示 [0,i-1] 之和
	for i, x := range nums[:n-1] {
		sum[i+1] = sum[i] + x
	}

	mul := 1 // [i+1,n-1] 之积
	for i := n - 1; i > 0; i-- {
		if sum[i] == mul { // [0,i-1] 之和等于 [i+1,n-1] 之积
			return i
		}
		// 如果 mul*nums[i] > sum[i-1]，那么继续向左遍历，mul 越来越大（或者不变），sum 越来越小，不可能找到答案
		// 为避免乘法溢出，改成等价的除法
		if mul > sum[i-1]/nums[i] {
			break
		}
		mul *= nums[i]
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 优化

先计算 $\textit{nums}$ 的和，然后在倒序遍历的过程中减去遍历的数，也能求出前缀和。这样可以做到 $\mathcal{O}(1)$ 空间。

```py [sol-Python3]
class Solution:
    def smallestBalancedIndex(self, nums: list[int]) -> int:
        pre = sum(nums)
        mul = 1
        for i in range(len(nums) - 1, 0, -1):
            pre -= nums[i]
            if pre < mul:
                break
            if pre == mul:
                return i
            mul *= nums[i]
        return -1
```

```java [sol-Java]
class Solution {
    public int smallestBalancedIndex(int[] nums) {
        int n = nums.length;
        long sum = 0;
        for (int i = 0; i < n - 1; i++) {
            sum += nums[i];
        }

        long mul = 1;
        for (int i = n - 1; i > 0; i--) {
            if (sum == mul) {
                return i;
            }
            sum -= nums[i - 1];
            if (mul > sum / nums[i]) {
                break;
            }
            mul *= nums[i];
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestBalancedIndex(vector<int>& nums) {
        long long sum = reduce(nums.begin(), nums.end() - 1, 0LL);
        long long mul = 1;
        for (int i = nums.size() - 1; i > 0; i--) {
            if (sum == mul) {
                return i;
            }
            sum -= nums[i - 1];
            if (mul > sum / nums[i]) {
                break;
            }
            mul *= nums[i];
        }
        return -1;
    }
};
```

```go [sol-Go]
func smallestBalancedIndex(nums []int) int {
	n := len(nums)
	sum := 0
	for _, x := range nums[:n-1] {
		sum += x
	}

	mul := 1
	for i := n - 1; i > 0; i-- {
		if sum == mul {
			return i
		}
		sum -= nums[i-1]
		if mul > sum/nums[i] {
			break
		}
		mul *= nums[i]
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：相向双指针

左指针从 $0$ 开始，右指针从 $n-1$ 开始。

比较前缀和与后缀积的大小，谁小就移动哪边的指针（排除了一个错误答案）。

注意前缀和等于后缀积的情况。由于 $\textit{nums}$ 包含 $1$：

- 如果把前缀和加上 $1$，移动左指针，万一剩下的数都是 $1$，这会导致前缀和始终大于后缀积，最后返回的是 $-1$，计算错误。
- 右指针是可以移动的（排除了一个错误答案）。反证法：假设右指针是答案，由于 $\textit{nums}$ 中的数都是正数，前缀和加上中间剩余的数会大于后缀积，矛盾。

所以前缀和等于后缀积的情况，要移动右指针。

**优化**：由于 $n\le 10^5$ 以及 $\textit{nums}[i]\le 10^9$，所以前缀和不会超过 $10^{14}$。因此，一旦发现后缀积超过 $10^{14}$，那么后续不可能出现前缀和等于后缀积的情况，可以直接返回 $-1$。

```py [sol-Python3]
class Solution:
    def smallestBalancedIndex(self, nums: list[int]) -> int:
        pre, suf = 0, 1
        l, r = 0, len(nums) - 1
        while l < r and suf <= 10 ** 14:
            if pre < suf:
                pre += nums[l]
                l += 1
            else:
                suf *= nums[r]
                r -= 1
        return l if pre == suf else -1
```

```java [sol-Java]
class Solution {
    public int smallestBalancedIndex(int[] nums) {
        long sum = 0, mul = 1;
        int l = 0, r = nums.length - 1;
        while (l < r) {
            if (sum < mul) {
                sum += nums[l++];
            } else {
                if (mul > (long) 1e14 / nums[r]) {
                    return -1;
                }
                mul *= nums[r--];
            }
        }
        return sum == mul ? l : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestBalancedIndex(vector<int>& nums) {
        long long sum = 0, mul = 1;
        int l = 0, r = nums.size() - 1;
        while (l < r) {
            if (sum < mul) {
                sum += nums[l++];
            } else {
                if (mul > (long long) 1e14 / nums[r]) {
                    return -1;
                }
                mul *= nums[r--];
            }
        }
        return sum == mul ? l : -1;
    }
};
```

```go [sol-Go]
func smallestBalancedIndex(nums []int) int {
	sum, mul := 0, 1
	l, r := 0, len(nums)-1
	for l < r {
		if sum < mul {
			sum += nums[l]
			l++
		} else {
			if mul > 1e14/nums[r] {
				return -1
			}
			mul *= nums[r]
			r--
		}
	}
	if sum == mul {
		return l
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果 $\textit{nums}$ 包含 $0$ 和负数，怎么写？

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
