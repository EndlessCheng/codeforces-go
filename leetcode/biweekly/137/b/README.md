**核心思路**：找连续上升的段。如果段长至少是 $k$，那么这段中的所有长为 $k$ 的子数组都是符合要求的，子数组的最后一个元素是最大的。

具体来说，遍历数组的同时，用一个计数器 $\textit{cnt}$ 统计连续递增的元素个数：

- 初始化 $\textit{cnt}=0$。
- 如果 $i=0$ 或者 $\textit{nums}[i]= \textit{nums}[i-1]+1$，把 $\textit{cnt}$ 增加 $1$；否则，把 $\textit{cnt}$ 置为 $1$。
- 如果发现 $\textit{cnt}\ge k$，那么下标从 $i-k+1$ 到 $i$ 的这个子数组的能量值为 $\textit{nums}[i]$，即 $\textit{ans}[i-k+1]=\textit{nums}[i]$。

具体例子请看 [视频讲解](https://www.bilibili.com/video/BV1ZH4y1c7GA/)，欢迎关注~

```py [sol-Python3]
class Solution:
    def resultsArray(self, nums: List[int], k: int) -> List[int]:
        ans = [-1] * (len(nums) - k + 1)
        cnt = 0
        for i, x in enumerate(nums):
            cnt = cnt + 1 if i == 0 or x == nums[i - 1] + 1 else 1
            if cnt >= k:
                ans[i - k + 1] = x
        return ans
```

```java [sol-Java]
class Solution {
    public int[] resultsArray(int[] nums, int k) {
        int[] ans = new int[nums.length - k + 1];
        Arrays.fill(ans, -1);
        int cnt = 0;
        for (int i = 0; i < nums.length; i++) {
            cnt = i == 0 || nums[i] == nums[i - 1] + 1 ? cnt + 1 : 1;
            if (cnt >= k) {
                ans[i - k + 1] = nums[i];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> resultsArray(vector<int>& nums, int k) {
        vector<int> ans(nums.size() - k + 1, -1);
        int cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            cnt = i == 0 || nums[i] == nums[i - 1] + 1 ? cnt + 1 : 1;
            if (cnt >= k) {
                ans[i - k + 1] = nums[i];
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int* resultsArray(int* nums, int numsSize, int k, int* returnSize) {
    *returnSize = numsSize - k + 1;
    int* ans = malloc(*returnSize * sizeof(int));
    memset(ans, -1, *returnSize * sizeof(int));

    int cnt = 0;
    for (int i = 0; i < numsSize; i++) {
        cnt = i == 0 || nums[i] == nums[i - 1] + 1 ? cnt + 1 : 1;
        if (cnt >= k) {
            ans[i - k + 1] = nums[i];
        }
    }
    return ans;
}
```

```go [sol-Go]
func resultsArray(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	for i := range ans {
		ans[i] = -1
	}
	cnt := 0
	for i, x := range nums {
		if i == 0 || x == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			ans[i-k+1] = x
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var resultsArray = function(nums, k) {
    const ans = Array(nums.length - k + 1).fill(-1);
    let cnt = 0;
    for (let i = 0; i < nums.length; i++) {
        cnt = i === 0 || nums[i] === nums[i - 1] + 1 ? cnt + 1 : 1;
        if (cnt >= k) {
            ans[i - k + 1] = nums[i];
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn results_array(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let k = k as usize;
        let mut ans = vec![-1; nums.len() - k + 1];
        let mut cnt = 0;
        for i in 0..nums.len() {
            if i == 0 || nums[i] == nums[i - 1] + 1 {
                cnt += 1;
            } else {
                cnt = 1;
            }
            if cnt >= k {
                ans[i - k + 1] = nums[i];
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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
