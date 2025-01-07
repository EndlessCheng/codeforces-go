去年十月的每日一题，出过一道类似的 [3171. 找到按位或最接近 K 的子数组](https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/)。

由于思路是一样的，本文只留代码，具体原理请看 [我的题解](https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/solutions/2798206/li-yong-and-de-xing-zhi-pythonjavacgo-by-gg4d/)。

## 方法一：LogTrick

```py [sol-Python3]
class Solution:
    def minimumSubarrayLength(self, nums: List[int], k: int) -> int:
        ans = inf
        for i, x in enumerate(nums):
            if x >= k:
                return 1
            j = i - 1
            while j >= 0 and nums[j] | x != nums[j]:
                nums[j] |= x
                if nums[j] >= k:
                    ans = min(ans, i - j + 1)
                j -= 1
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumSubarrayLength(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (x >= k) {
                return 1;
            }
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x;
                if (nums[j] >= k) {
                    ans = Math.min(ans, i - j + 1);
                }
            }
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSubarrayLength(vector<int>& nums, int k) {
        int ans = INT_MAX;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (x >= k) {
                return 1;
            }
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x;
                if (nums[j] >= k) {
                    ans = min(ans, i - j + 1);
                }
            }
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int minimumSubarrayLength(int* nums, int numsSize, int k) {
    int ans = INT_MAX;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        if (x >= k) {
            return 1;
        }
        for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
            nums[j] |= x;
            if (nums[j] >= k) {
                ans = MIN(ans, i - j + 1);
            }
        }
    }
    return ans == INT_MAX ? -1 : ans;
}
```

```go [sol-Go]
func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		if x >= k {
			return 1
		}
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			if nums[j] >= k {
				ans = min(ans, i-j+1)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

```js [sol-JavaScript]
var minimumSubarrayLength = function(nums, k) {
    let ans = Infinity;
    for (let i = 0; i < nums.length; i++) {
        const x = nums[i];
        if (x >= k) {
            return 1;
        }
        for (let j = i - 1; j >= 0 && (nums[j] | x) !== nums[j]; j--) {
            nums[j] |= x;
            if (nums[j] >= k) {
                ans = Math.min(ans, i - j + 1);
            }
        }
    }
    return ans === Infinity ? -1 : ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_subarray_length(mut nums: Vec<i32>, k: i32) -> i32 {
        let mut ans = usize::MAX;
        for i in 0..nums.len() {
            let x = nums[i];
            if x >= k {
                return 1;
            }
            let mut j = i - 1;
            while j < nums.len() && (nums[j] | x) != nums[j] {
                nums[j] |= x;
                if nums[j] >= k {
                    ans = ans.min(i - j + 1);
                }
                j -= 1;
            }
        }
        if ans == usize::MAX { -1 } else { ans as _ }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于 $2^{29}-1<10^9<2^{30}-1$，二进制数对应集合的大小不会超过 $29$，因此在 OR 运算下，每个数字至多可以增大 $29$ 次。**总体上看**，二重循环的总循环次数等于每个数字可以增大的次数之和，即 $O(n\log U)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：滑动窗口+栈

```py [sol-Python3]
class Solution:
    def minimumSubarrayLength(self, nums: List[int], k: int) -> int:
        ans = inf
        left = bottom = right_or = 0
        for right, x in enumerate(nums):
            right_or |= x
            while left <= right and nums[left] | right_or >= k:
                ans = min(ans, right - left + 1)
                left += 1
                if bottom < left:
                    # 重新构建一个栈
                    for i in range(right - 1, left - 1, -1):
                        nums[i] |= nums[i + 1]
                    bottom = right
                    right_or = 0
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumSubarrayLength(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;
        int left = 0;
        int bottom = 0;
        int rightOr = 0;
        for (int right = 0; right < nums.length; right++) {
            rightOr |= nums[right];
            while (left <= right && (nums[left] | rightOr) >= k) {
                ans = Math.min(ans, right - left + 1);
                left++;
                if (bottom < left) {
                    // 重新构建一个栈
                    for (int i = right - 1; i >= left; i--) {
                        nums[i] |= nums[i + 1];
                    }
                    bottom = right;
                    rightOr = 0;
                }
            }
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSubarrayLength(vector<int>& nums, int k) {
        int ans = INT_MAX, left = 0, bottom = 0, right_or = 0;
        for (int right = 0; right < nums.size(); right++) {
            right_or |= nums[right];
            while (left <= right && (nums[left] | right_or) >= k) {
                ans = min(ans, right - left + 1);
                left++;
                if (bottom < left) {
                    // 重新构建一个栈
                    for (int i = right - 1; i >= left; i--) {
                        nums[i] |= nums[i + 1];
                    }
                    bottom = right;
                    right_or = 0;
                }
            }
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int minimumSubarrayLength(int* nums, int numsSize, int k) {
    int ans = INT_MAX, left = 0, bottom = 0, right_or = 0;
    for (int right = 0; right < numsSize; right++) {
        right_or |= nums[right];
        while (left <= right && (nums[left] | right_or) >= k) {
            ans = MIN(ans, right - left + 1);
            left++;
            if (bottom < left) {
                // 重新构建一个栈
                for (int i = right - 1; i >= left; i--) {
                    nums[i] |= nums[i + 1];
                }
                bottom = right;
                right_or = 0;
            }
        }
    }
    return ans == INT_MAX ? -1 : ans;
}
```

```go [sol-Go]
func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, x := range nums {
		rightOr |= x
		for left <= right && nums[left]|rightOr >= k {
			ans = min(ans, right-left+1)
			left++
			if bottom < left {
				// 重新构建一个栈
				for i := right - 1; i >= left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

```js [sol-JavaScript]
var minimumSubarrayLength = function(nums, k) {
    let ans = Infinity, left = 0, bottom = 0, rightOr = 0;
    for (let right = 0; right < nums.length; right++) {
        rightOr |= nums[right];
        while (left <= right && (nums[left] | rightOr) >= k) {
            ans = Math.min(ans, right - left + 1);
            left++;
            if (bottom < left) {
                // 重新构建一个栈
                for (let i = right - 1; i >= left; i--) {
                    nums[i] |= nums[i + 1];
                }
                bottom = right;
                rightOr = 0;
            }
        }
    }
    return ans === Infinity ? -1 : ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_subarray_length(mut nums: Vec<i32>, k: i32) -> i32 {
        let mut ans = usize::MAX;
        let mut left = 0;
        let mut bottom = 0;
        let mut right_or = 0;
        for right in 0..nums.len() {
            right_or |= nums[right];
            while left <= right && (nums[left] | right_or) >= k {
                ans = ans.min(right - left + 1);
                left += 1;
                if bottom < left {
                    // 重新构建一个栈
                    for i in (left..right).rev() {
                        nums[i] |= nums[i + 1];
                    }
                    bottom = right;
                    right_or = 0;
                }
            }
        }
        if ans == usize::MAX { -1 } else { ans as _ }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然我们写了个三重循环，但每个元素至多入栈出栈各一次，所以三重循环的**总**循环次数是 $\mathcal{O}(n)$ 的，所以时间复杂度是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见位运算题单中的「**LogTrick**」。

## 思考题

如果把 OR 改成 XOR，要怎么做？

提示：前缀异或和 + [1803. 统计异或值在范围内的数对有多少](https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/)。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
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
