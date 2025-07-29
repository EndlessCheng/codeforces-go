## 方法一：LogTrick

**详细讲解**：[LogTrick 入门教程](https://zhuanlan.zhihu.com/p/1933215367158830792)。

```py [sol-Python3]
class Solution:
    def smallestSubarrays(self, nums: List[int]) -> List[int]:
        ans = [1] * len(nums)  # 子数组的长度至少是 1
        for i, x in enumerate(nums):  # 计算右端点为 i 的子数组的或值
            for j in range(i - 1, -1, -1):
                if (nums[j] | x) == nums[j]:  # nums[j] 及其左边元素无法增大
                    break
                nums[j] |= x  # nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
                ans[j] = i - j + 1  # nums[j] 最后一次增大时的子数组长度就是答案
        return ans
```

```java [sol-Java]
class Solution {
    public int[] smallestSubarrays(int[] nums) {
        int n = nums.length;
        int[] ans = new int[n];
        for (int i = 0; i < n; i++) { // 计算右端点为 i 的子数组的或值
            int x = nums[i];
            ans[i] = 1; // 子数组的长度至少是 1
            // 循环直到 nums[j] 无法增大，其左侧元素也无法增大
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x; // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
                ans[j] = i - j + 1; // nums[j] 最后一次增大时的子数组长度就是答案
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> smallestSubarrays(vector<int>& nums) {
        int n = nums.size();
        vector<int> ans(n);
        for (int i = 0; i < n; i++) { // 计算右端点为 i 的子数组的或值
            int x = nums[i];
            ans[i] = 1; // 子数组的长度至少是 1
            // 循环直到 nums[j] 无法增大，其左侧元素也无法增大
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x; // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
                ans[j] = i - j + 1; // nums[j] 最后一次增大时的子数组长度就是答案
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int* smallestSubarrays(int* nums, int numsSize, int* returnSize) {
    int* ans = malloc(numsSize * sizeof(int));
    *returnSize = numsSize;

    for (int i = 0; i < numsSize; i++) { // 计算右端点为 i 的子数组的或值
        int x = nums[i];
        ans[i] = 1; // 子数组的长度至少是 1
        // 循环直到 nums[j] 无法增大，其左侧元素也无法增大
        for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
            nums[j] |= x; // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
            ans[j] = i - j + 1; // nums[j] 最后一次增大时的子数组长度就是答案
        }
    }

    return ans;
}
```

```go [sol-Go]
func smallestSubarrays(nums []int) []int {
	ans := make([]int, len(nums))
	for i, x := range nums { // 计算右端点为 i 的子数组的或值
		ans[i] = 1 // 子数组的长度至少是 1
		// 循环直到 nums[j] 无法增大，其左侧元素也无法增大
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
			ans[j] = i - j + 1 // nums[j] 最后一次增大时的子数组长度就是答案
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var smallestSubarrays = function(nums) {
    const n = nums.length;
    const ans = Array(n).fill(1); // 子数组的长度至少是 1
    for (let i = 0; i < n; i++) { // 计算右端点为 i 的子数组的或值
        let x = nums[i];
        // 循环直到 nums[j] 无法增大，其左侧元素也无法增大
        for (let j = i - 1; j >= 0 && (nums[j] | x) !== nums[j]; j--) {
            nums[j] |= x; // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
            ans[j] = i - j + 1; // nums[j] 最后一次增大时的子数组长度就是答案
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn smallest_subarrays(mut nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut ans = vec![0; n];
        for i in 0..n { // 计算右端点为 i 的子数组的或值
            let x = nums[i];
            ans[i] = 1; // 子数组的长度至少是 1
            for j in (0..i).rev() {
                if (nums[j] | x) == nums[j] { // nums[j] 及其左边元素无法增大
                    break;
                }
                nums[j] |= x; // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
                ans[j] = (i - j + 1) as i32; // nums[j] 最后一次增大时的子数组长度就是答案
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})\le 10^9$。由于 $10^9<2^{30}$，二进制数对应集合的大小不会超过 $30$，因此在或运算下，每个数字至多可以增大 $30$ 次（从空集增大到有 $30$ 个元素）。**总体上看**，二重循环的总循环次数等于每个数字可以增大的次数之和，即 $O(n\log U)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 方法二：滑动窗口+栈

[原理讲解](https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/solutions/2798206/li-yong-and-de-xing-zhi-pythonjavacgo-by-gg4d/)

本题由于要获知的信息都在 $\textit{nums}[i]$ 的右侧，所以要倒着滑窗。外层循环枚举左端点 $\textit{left}$，内层循环缩小右端点 $\textit{right}$。当我们发现子数组 $[\textit{left},\textit{right}]$ 的或值等于子数组 $[\textit{left},\textit{right}-1]$ 的或值时，说明窗口右端点可以缩小。

另外要保证栈中至少有两个数，方便判断窗口右端点是否可以缩小。

```py [sol-Python3]
class Solution:
    def smallestSubarrays(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n
        ans[-1] = 1
        if n == 1:
            return ans

        # 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
        nums[-1] |= nums[-2]
        left_or, right, bottom = 0, n - 1, n - 2
        for left in range(n - 2, -1, -1):
            left_or |= nums[left]
            # 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
            while right > left and (left_or | nums[right]) == (left_or | nums[right - 1]):
                right -= 1
                # 栈中只剩一个数
                if bottom >= right:
                    # 重新构建一个栈，栈底为 left，栈顶为 right
                    for i in range(left + 1, right + 1):
                        nums[i] |= nums[i - 1]
                    bottom = left
                    left_or = 0
            ans[left] = right - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public int[] smallestSubarrays(int[] nums) {
        int n = nums.length;
        int[] ans = new int[n];
        ans[n - 1] = 1;
        if (n == 1) {
            return ans;
        }

        // 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
        nums[n - 1] |= nums[n - 2];
        int leftOr = 0, right = n - 1, bottom = n - 2;
        for (int left = n - 2; left >= 0; left--) {
            leftOr |= nums[left];
            // 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
            while (right > left && (leftOr | nums[right]) == (leftOr | nums[right - 1])) {
                right--;
                // 栈中只剩一个数
                if (bottom >= right) {
                    // 重新构建一个栈，栈底为 left，栈顶为 right
                    for (int i = left + 1; i <= right; i++) {
                        nums[i] |= nums[i - 1];
                    }
                    bottom = left;
                    leftOr = 0;
                }
            }
            ans[left] = right - left + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> smallestSubarrays(vector<int>& nums) {
        int n = nums.size();
        vector<int> ans(n);
        ans[n - 1] = 1;
        if (n == 1) {
            return ans;
        }

        // 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
        nums[n - 1] |= nums[n - 2];
        int left_or = 0, right = n - 1, bottom = n - 2;
        for (int left = n - 2; left >= 0; left--) {
            left_or |= nums[left];
            // 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
            while (right > left && (left_or | nums[right]) == (left_or | nums[right - 1])) {
                right--;
                // 栈中只剩一个数
                if (bottom >= right) {
                    // 重新构建一个栈，栈底为 left，栈顶为 right
                    for (int i = left + 1; i <= right; i++) {
                        nums[i] |= nums[i - 1];
                    }
                    bottom = left;
                    left_or = 0;
                }
            }
            ans[left] = right - left + 1;
        }
        return ans;
    }
};
```

```c [sol-C]
int* smallestSubarrays(int* nums, int n, int* returnSize) {
    int* ans = malloc(n * sizeof(int));
    *returnSize = n;

    ans[n - 1] = 1;
    if (n == 1) {
        return ans;
    }

    // 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
    nums[n - 1] |= nums[n - 2];
    int left_or = 0, right = n - 1, bottom = n - 2;
    for (int left = n - 2; left >= 0; left--) {
        left_or |= nums[left];
        // 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
        while (right > left && (left_or | nums[right]) == (left_or | nums[right - 1])) {
            right--;
            // 栈中只剩一个数
            if (bottom >= right) {
                // 重新构建一个栈，栈底为 left，栈顶为 right
                for (int i = left + 1; i <= right; i++) {
                    nums[i] |= nums[i - 1];
                }
                bottom = left;
                left_or = 0;
            }
        }
        ans[left] = right - left + 1;
    }
    return ans;
}
```

```go [sol-Go]
func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[n-1] = 1
	if n == 1 {
		return ans
	}

	// 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
	nums[n-1] |= nums[n-2]
	leftOr, right, bottom := 0, n-1, n-2
	for left := n - 2; left >= 0; left-- {
		leftOr |= nums[left]
		// 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
		for right > left && leftOr|nums[right] == leftOr|nums[right-1] {
			right--
			// 栈中只剩一个数
			if bottom >= right {
				// 重新构建一个栈，栈底为 left，栈顶为 right
				for i := left + 1; i <= right; i++ {
					nums[i] |= nums[i-1]
				}
				bottom = left
				leftOr = 0
			}
		}
		ans[left] = right - left + 1
	}
	return ans
}
```

```js [sol-JavaScript]
var smallestSubarrays = function(nums) {
    const n = nums.length;
    const ans = Array(n).fill(0);
    ans[n - 1] = 1;
    if (n === 1) {
        return ans;
    }

    // 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
    nums[n - 1] |= nums[n - 2];
    let leftOr = 0, right = n - 1, bottom = n - 2;
    for (let left = n - 2; left >= 0; left--) {
        leftOr |= nums[left];
        // 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
        while (right > left && (leftOr | nums[right]) === (leftOr | nums[right - 1])) {
            right--;
            // 栈中只剩一个数
            if (bottom >= right) {
                // 重新构建一个栈，栈底为 left，栈顶为 right
                for (let i = left + 1; i <= right; i++) {
                    nums[i] |= nums[i - 1];
                }
                bottom = left;
                leftOr = 0;
            }
        }
        ans[left] = right - left + 1;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn smallest_subarrays(mut nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut ans = vec![0; n];
        ans[n - 1] = 1;
        if n == 1 {
            return ans;
        }

        // 保证栈中至少有两个数，方便判断窗口右端点是否要缩小
        nums[n - 1] |= nums[n - 2];
        let mut left_or = 0;
        let mut right = n - 1;
        let mut bottom = n - 2;
        for left in (0..=n - 2).rev() {
            left_or |= nums[left];
            // 子数组 [left,right] 的或值 = 子数组 [left,right-1] 的或值，说明窗口右端点可以缩小
            while right > left && (left_or | nums[right]) == (left_or | nums[right - 1]) {
                right -= 1;
                // 栈中只剩一个数
                if bottom >= right {
                    // 重新构建一个栈，栈底为 left，栈顶为 right
                    for i in left + 1..=right {
                        nums[i] |= nums[i - 1];
                    }
                    bottom = left;
                    left_or = 0;
                }
            }
            ans[left] = (right - left + 1) as i32;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然我们写了个三重循环，但每个元素至多入栈出栈各一次，所以三重循环的**总**循环次数是 $\mathcal{O}(n)$ 的，所以时间复杂度是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

把「或」改成「异或」，其余不变，要怎么做？

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面位运算题单的「**LogTrick**」。部分题目也可以用滑动窗口+栈解决。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
