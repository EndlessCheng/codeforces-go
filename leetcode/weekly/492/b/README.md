定义 $\textit{sum}[i]$ 表示 $[0,i-1]$ 的元素和（前缀和）。特别地，$\textit{sum}[0] = 0$。

定义 $\textit{mul}[i]$ 表示 $[i+1,n-1]$ 的元素积（后缀积）。特别地，$\textit{mul}[n-1] = 1$。

问题变成：

- 寻找满足 $\textit{sum}[i] = \textit{mul}[i]$ 的最小的下标 $i$。

由于 $\textit{nums}[i] > 0$，所以 $\textit{sum}$ 是严格递增的，$\textit{mul}$ 是（非严格）递减的。画出函数图像的话，至多有一个交点。

所以本题要么无解，要么恰好有一个解。

代码实现时，可以先把 $\textit{sum}$ 算出来，然后倒着遍历 $\textit{nums}$，同时计算 $\textit{mul}$。这样做，$\textit{mul}$ 可以简化成一个变量。

此外，当 $i=0$ 时，$\textit{sum}[0] = 0$，$\textit{mul}[0]>0$，所以 $\textit{sum}[0] \ne \textit{mul}[0]$。所以至多循环到 $i=1$ 为止。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def smallestBalancedIndex(self, nums: List[int]) -> int:
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

## 优化

先计算 $\textit{nums}$ 的和，然后在倒序遍历的过程中减去遍历的数，也能求出前缀和。这样可以做到 $\mathcal{O}(1)$ 空间。

```py [sol-Python3]
class Solution:
    def smallestBalancedIndex(self, nums: List[int]) -> int:
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

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
