怎么计算子数组的 AND？

首先，我们有如下 $\mathcal{O}(n^2)$ 的暴力算法：

从左到右正向遍历 $\textit{nums}$，对于 $x=\textit{nums}[i]$，从 $i-1$ 开始倒着遍历 $\textit{nums}[j]$，更新 $\textit{nums}[j]=\textit{nums}[j]\&x$。

- $i=1$ 时，我们会把 $\textit{nums}[0]$ 到 $\textit{nums}[1]$ 的 AND 记录在 $\textit{nums}[0]$ 中。
- $i=2$ 时，我们会把 $\textit{nums}[1]$ 到 $\textit{nums}[2]$ 的 AND 记录在 $\textit{nums}[1]$ 中，$\textit{nums}[0]$ 到 $\textit{nums}[2]$ 的 AND 记录在 $\textit{nums}[0]$ 中。
- $i=3$ 时，我们会把 $\textit{nums}[2]$ 到 $\textit{nums}[3]$ 的 AND 记录在 $\textit{nums}[2]$ 中；$\textit{nums}[1]$ 到 $\textit{nums}[3]$ 的 AND 记录在 $\textit{nums}[1]$ 中；$\textit{nums}[0]$ 到 $\textit{nums}[3]$ 的 AND 记录在 $\textit{nums}[0]$ 中。
- 按照该算法，可以计算出所有子数组的 AND。注意单个元素也算子数组。

下面来优化该算法。

前置知识：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

把二进制数看成集合，两个数的 AND 就是两个集合的**交集**。

对于两个二进制数 $a$ 和 $b$，如果 $a\&b = a$，从集合的角度上看，$a$ 对应的集合是 $b$ 对应的集合的子集。或者说，$b$ 对应的集合是 $a$ 对应的集合的**超集**。

据此我们可以提出如下优化：

仍然是从左到右正向遍历 $\textit{nums}$，对于 $x=\textit{nums}[i]$，从 $i-1$ 开始倒着遍历 $\textit{nums}[j]$：
- 如果 $\textit{nums}[j]\&x\ne\textit{nums}[j]$，说明 $\textit{nums}[j]$ 可以变小（求交集后，集合元素只会减少不会变多），更新 $\textit{nums}[j]=\textit{nums}[j]\&x$。
- 否则 $\textit{nums}[j]\&x=\textit{nums}[j]$，从集合的角度看，此时 $x$ 不仅是 $\textit{nums}[j]$ 的超集，同时也是 $\textit{nums}[k]\ (k<j)$ 的超集（因为前面的循环保证了每个集合都是其左侧相邻集合的超集），在 $A\subseteq B$ 的前提下，$A\cap B=A$，所以后续的循环都不会改变元素值，**退出内层循环**。

## 方法一：二分查找

由于每个元素都是其右侧元素的子集，所以从 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的元素值是非降的。既然是有序数组，我们可以在 $[0,i]$ 中**二分查找** $k$，做法见 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)（[视频讲解](https://www.bilibili.com/video/BV1AP41137w7/)）。

设左闭右开区间 $[\textit{left},\textit{right})$ 是 $\textit{nums}[j]=k$ 的 $j$ 的范围。把左闭右开区间的长度 $\textit{right}-\textit{left}$ 加入答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Yz421q7dD/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        ans = 0
        for i, x in enumerate(nums):
            for j in range(i - 1, -1, -1):
                if nums[j] & x == nums[j]:
                    break
                nums[j] &= x
            ans += bisect_right(nums, k, 0, i + 1) - bisect_left(nums, k, 0, i + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, int k) {
        long ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            for (int j = i - 1; j >= 0 && (nums[j] & x) != nums[j]; j--) {
                nums[j] &= x;
            }
            ans += lowerBound(nums, i + 1, k + 1) - lowerBound(nums, i + 1, k);
        }
        return ans;
    }

    // https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int right, int target) {
        int left = -1; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int>& nums, int k) {
        long long ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            for (int j = i - 1; j >= 0 && (nums[j] & x) != nums[j]; j--) {
                nums[j] &= x;
            }
            ans += upper_bound(nums.begin(), nums.begin() + i + 1, k) -
                   lower_bound(nums.begin(), nums.begin() + i + 1, k);
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, k int) (ans int64) {
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		a := nums[:i+1]
		ans += int64(sort.SearchInts(a, k+1) - sort.SearchInts(a, k))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U + n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于 $2^{29}-1<10^9<2^{30}-1$，二进制数对应集合的大小不会超过 $29$，因此在 AND 运算下，每个数字至多可以减小 $29$ 次。总体上看（除去二分），二重循环的总循环次数等于每个数字可以减小的次数之和，即 $O(n\log U)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：三指针

由于元素值只会减少，所以当 $i$ 增大时，$\textit{left}$ 和 $\textit{right}$ 不会减少，有了单调性的保证，上面的二分查找可以改成 [滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/)：

- 当 $\textit{left}\le i$ 且 $\textit{nums}[\textit{left}] < k$ 时，把 $\textit{left}$ 加一。
- 当 $\textit{right}\le i$ 且 $\textit{nums}[\textit{right}] \le k$ 时，把 $\textit{right}$ 加一。
- 把左闭右开区间的长度 $\textit{right}-\textit{left}$ 加入答案。

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        ans = left = right = 0
        for i, x in enumerate(nums):
            for j in range(i - 1, -1, -1):
                if nums[j] & x == nums[j]:
                    break
                nums[j] &= x
            while left <= i and nums[left] < k:
                left += 1
            while right <= i and nums[right] <= k:
                right += 1
            ans += right - left
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, int k) {
        long ans = 0;
        int left = 0;
        int right = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            for (int j = i - 1; j >= 0 && (nums[j] & x) != nums[j]; j--) {
                nums[j] &= x;
            }
            while (left <= i && nums[left] < k) {
                left++;
            }
            while (right <= i && nums[right] <= k) {
                right++;
            }
            ans += right - left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int>& nums, int k) {
        long long ans = 0;
        int left = 0, right = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            for (int j = i - 1; j >= 0 && (nums[j] & x) != nums[j]; j--) {
                nums[j] &= x;
            }
            while (left <= i && nums[left] < k) {
                left++;
            }
            while (right <= i && nums[right] <= k) {
                right++;
            }
            ans += right - left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, k int) (ans int64) {
	left, right := 0, 0
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		for left <= i && nums[left] < k {
			left++
		}
		for right <= i && nums[right] <= k {
			right++
		}
		ans += int64(right - left)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于 $\textit{left}$ 和 $\textit{right}$ 只会增大，不会减少，所以 $\textit{left}$ 和 $\textit{right}$ 的移动次数之和为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法三：维护等于 k 的子数组个数

在更新 $\textit{nums}[j]$ 的同时，维护值等于 $k$ 的元素个数，也就是 AND 值为 $k$ 的子数组个数。

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        ans = cnt = 0
        for i, x in enumerate(nums):
            if x == k:
                cnt += 1
            for j in range(i - 1, -1, -1):
                if nums[j] & x == nums[j]:
                    break
                if nums[j] == k:
                    cnt -= 1
                nums[j] &= x
                if nums[j] == k:
                    cnt += 1
            ans += cnt
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, int k) {
        long ans = 0;
        int cnt = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            cnt += x == k ? 1 : 0;
            for (int j = i - 1; j >= 0 && (nums[j] & x) != nums[j]; j--) {
                cnt -= nums[j] == k ? 1 : 0;
                nums[j] &= x;
                cnt += nums[j] == k ? 1 : 0;
            }
            ans += cnt;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int>& nums, int k) {
        long long ans = 0;
        int cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            cnt += x == k;
            for (int j = i - 1; j >= 0 && (nums[j] & x) != nums[j]; j--) {
                cnt -= nums[j] == k;
                nums[j] &= x;
                cnt += nums[j] == k;
            }
            ans += cnt;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, k int) (ans int64) {
	cnt := 0
	for i, x := range nums {
		if x == k {
			cnt++
		}
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			if nums[j] == k {
				cnt--
			}
			nums[j] &= x
			if nums[j] == k {
				cnt++
			}
		}
		ans += int64(cnt)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见位运算题单中的「**LogTrick**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. 【本题相关】[位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
