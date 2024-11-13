**核心思想**：枚举交替子数组的**右端点下标** $i$，统计**交替子数组的个数**，加入答案。

以 $\textit{nums}=[0,1,1,0,1,1]$ 为例，我们来看看有哪些交替子数组的**右端点下标**是 $i$：

| $i$  | $\textit{nums}_i$  | 交替子数组  | 个数 |
|---|---|---|---|
|  $0$ | $0$  | $[0]$  | $1$ |
|  $1$ | $1$  | $[1],[0,1]$  | $2$ |
|  $2$ | $1$  | $[1]$  | $1$ |
|  $3$ | $0$  | $[0],[1,0]$  | $2$ |
|  $4$ | $1$  | $[1],[0,1],[1,0,1]$  | $3$ |
|  $5$ | $1$  | $[1]$  | $1$ |

把 $\textit{nums}_4$ 加到以 $3$ 为右端点的交替子数组的末尾，可以得到 $2$ 个交替子数组 $[0,1]$ 和 $[1,0,1]$，再算上 $\textit{nums}_4$ 单独作为一个长为 $1$ 的交替子数组，因此以 $4$ 为右端点的交替子数组有 $3$ 个。

一般地，如果 $\textit{nums}_i \ne \textit{nums}_{i-1}$，我们可以把 $\textit{nums}_i$ 加到所有以 $i-1$ 为右端点的交替子数组的末尾，所以「以 $i$ 为右端点的交替子数组个数」比「以 $i-1$ 为右端点的交替子数组个数」多 $1$。

**算法**：遍历 $\textit{nums}$ 的同时，维护 $\textit{cnt}$，表示右端点下标为 $i$ 的交替子数组的个数。

- 如果 $i>0$ 且 $\textit{nums}_i \ne \textit{nums}_{i-1}$，根据上面的讨论，把 $\textit{cnt}$ 增加 $1$。
- 否则，把 $\textit{cnt}$ 重置为 $1$，表示 $\textit{nums}_i$ 单独组成一个长为 $1$ 的交替子数组。

累加遍历过程中的 $\textit{cnt}$ 值，即为答案。

请看 [视频讲解](https://www.bilibili.com/video/BV1fq421A7CY/) 第三题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countAlternatingSubarrays(self, nums: List[int]) -> int:
        ans = cnt = 0
        for i in range(len(nums)):
            if i > 0 and nums[i] != nums[i - 1]:
                cnt += 1
            else:
                cnt = 1
            ans += cnt  # 有 cnt 个右端点下标为 i 的交替子数组
        return ans
```

```py [sol-Python3 pairwise]
class Solution:
    def countAlternatingSubarrays(self, nums: List[int]) -> int:
        ans = cnt = 1
        for x, y in pairwise(nums):
            cnt = 1 if x == y else cnt + 1
            ans += cnt  # 有 cnt 个以 i 为右端点的交替子数组
        return ans
```

```java [sol-Java]
class Solution {
    public long countAlternatingSubarrays(int[] nums) {
        long ans = 0;
        int cnt = 0;
        for (int i = 0; i < nums.length; i++) {
            if (i > 0 && nums[i] != nums[i - 1]) {
                cnt++;
            } else {
                cnt = 1;
            }
            ans += cnt; // 有 cnt 个右端点下标为 i 的交替子数组
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countAlternatingSubarrays(vector<int>& nums) {
        long long ans = 0;
        int cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (i > 0 && nums[i] != nums[i - 1]) {
                cnt++;
            } else {
                cnt = 1;
            }
            ans += cnt; // 有 cnt 个右端点下标为 i 的交替子数组
        }
        return ans;
    }
};
```

```c [sol-C]
long long countAlternatingSubarrays(int* nums, int numsSize) {
    long long ans = 0;
    int cnt = 0;
    for (int i = 0; i < numsSize; i++) {
        if (i > 0 && nums[i] != nums[i - 1]) {
            cnt++;
        } else {
            cnt = 1;
        }
        ans += cnt; // 有 cnt 个右端点下标为 i 的交替子数组
    }
    return ans;
}
```

```go [sol-Go]
func countAlternatingSubarrays(nums []int) (ans int64) {
	cnt := 0
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			cnt = 1
		} else {
			cnt++
		}
		ans += int64(cnt) // 有 cnt 个以 i 为右端点的交替子数组
	}
	return
}
```

```js [sol-JS]
var countAlternatingSubarrays = function(nums) {
    let ans = 0;
    let cnt = 0;
    for (let i = 0; i < nums.length; i++) {
        if (i > 0 && nums[i] != nums[i - 1]) {
            cnt++;
        } else {
            cnt = 1;
        }
        ans += cnt; // 有 cnt 个右端点下标为 i 的交替子数组
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_alternating_subarrays(nums: Vec<i32>) -> i64 {
        let mut ans = 0;
        let mut cnt = 0;
        for i in 0..nums.len() {
            if i > 0 && nums[i] != nums[i - 1] {
                cnt += 1;
            } else {
                cnt = 1;
            }
            ans += cnt; // 有 cnt 个右端点下标为 i 的交替子数组
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
