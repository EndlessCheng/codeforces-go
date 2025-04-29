下界分析：

- 把 $0$ 替换成最小的正整数 $1$。设 $s_1$ 为 $\textit{nums}_1$ 的元素和，$s_2$ 为 $\textit{nums}_2$ 的元素和。这是元素和的最小值。

分类讨论：

- 如果 $s_1 < s_2$ 且 $\textit{nums}_1$ 中没有 $0$，那么 $s_1$ 无法增大，$s_2$ 不能减小，所以无法让 $s_1=s_2$，返回 $-1$。
- 如果 $s_2 < s_1$ 且 $\textit{nums}_2$ 中没有 $0$，那么 $s_2$ 无法增大，$s_1$ 不能减小，所以无法让 $s_1=s_2$，返回 $-1$。
- 否则，可以把较小的元素和变成较大的元素和，答案为 $\max(s_1,s_2)$。

```py [sol-Python3]
class Solution:
    def minSum(self, nums1: List[int], nums2: List[int]) -> int:
        s1 = sum(max(x, 1) for x in nums1)
        s2 = sum(max(x, 1) for x in nums2)
        if s1 < s2 and 0 not in nums1 or s2 < s1 and 0 not in nums2:
            return -1
        return max(s1, s2)
```

```java [sol-Java]
class Solution {
    private record Pair(long sum, boolean zero) {}

    public long minSum(int[] nums1, int[] nums2) {
        Pair p1 = calc(nums1);
        Pair p2 = calc(nums2);
        if (!p1.zero && p1.sum < p2.sum || !p2.zero && p2.sum < p1.sum) {
            return -1;
        }
        return Math.max(p1.sum, p2.sum);
    }

    private Pair calc(int[] nums) {
        long sum = 0;
        boolean zero = false;
        for (int x : nums) {
            if (x == 0) {
                zero = true;
                sum++;
            } else {
                sum += x;
            }
        }
        return new Pair(sum, zero);
    }
}
```

```cpp [sol-C++]
class Solution {
    pair<long long, bool> calc(vector<int>& nums) {
        long long sum = 0;
        bool zero = false;
        for (int x : nums) {
            if (x == 0) {
                zero = true;
                sum++;
            } else {
                sum += x;
            }
        }
        return {sum, zero};
    }

public:
    long long minSum(vector<int>& nums1, vector<int>& nums2) {
        auto [s1, zero1] = calc(nums1);
        auto [s2, zero2] = calc(nums2);
        if (!zero1 && s1 < s2 || !zero2 && s2 < s1) {
            return -1;
        }
        return max(s1, s2);
    }
};
```

```go [sol-Go]
func calc(nums []int) (sum int64, zero bool) {
	for _, x := range nums {
		if x == 0 {
			zero = true
			sum++
		} else {
			sum += int64(x)
		}
	}
	return
}

func minSum(nums1, nums2 []int) int64 {
	s1, zero1 := calc(nums1)
	s2, zero2 := calc(nums2)
	if !zero1 && s1 < s2 || !zero2 && s2 < s1 {
		return -1
	}
	return max(s1, s2)
}
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

long long calc(int* nums, int numsSize, bool* zero) {
    long long sum = 0;
    *zero = false;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] == 0) {
            *zero = true;
            sum++;
        } else {
            sum += nums[i];
        }
    }
    return sum;
}

long long minSum(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    bool zero1, zero2;
    long long s1 = calc(nums1, nums1Size, &zero1);
    long long s2 = calc(nums2, nums2Size, &zero2);
    if (!zero1 && s1 < s2 || !zero2 && s2 < s1) {
        return -1;
    }
    return MAX(s1, s2);
}
```

```js [sol-JavaScript]
function calc(nums) {
    let sum = 0;
    let zero = false;
    for (const x of nums) {
        if (x === 0) {
            zero = true;
            sum++;
        } else {
            sum += x;
        }
    }
    return [sum, zero];
}

var minSum = function(nums1, nums2) {
    const [s1, zero1] = calc(nums1);
    const [s2, zero2] = calc(nums2);
    if (!zero1 && s1 < s2 || !zero2 && s2 < s1) {
        return -1;
    }
    return Math.max(s1, s2);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_sum(nums1: Vec<i32>, nums2: Vec<i32>) -> i64 {
        fn calc(nums: Vec<i32>) -> (i64, bool) {
            let mut sum = 0;
            let mut zero = false;
            for x in nums {
                if x == 0 {
                    zero = true;
                    sum += 1;
                } else {
                    sum += x as i64;
                }
            }
            (sum, zero)
        }

        let (s1, zero1) = calc(nums1);
        let (s2, zero2) = calc(nums2);
        if !zero1 && s1 < s2 || !zero2 && s2 < s1 {
            return -1;
        }
        s1.max(s2)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
