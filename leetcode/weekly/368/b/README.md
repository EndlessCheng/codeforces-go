对于有三个（或者四个）变量的问题，枚举中间的变量往往更好算。

为什么？比如问题有三个下标，需要满足 $0\le i<j<k<n$，对比一下：

- 枚举 $i$，后续计算中还需保证 $j < k$。
- 枚举 $j$，那么 $i$ 和 $k$ 自动被 $j$ 隔开，互相独立，后续计算中无需关心 $i$ 和 $k$ 的位置关系。

所以枚举中间的变量更简单。

---

枚举中间的 $\textit{nums}[j]$，我们需要求出 $j$ **左边所有元素的最小值**和**右边所有元素的最小值**。

这可以**递推**计算。定义 $\textit{suf}[i]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 的最小值（后缀最小值），则有

$$
\textit{suf}[i] = \min(\textit{suf}[i+1], \textit{nums}[i])
$$

前缀最小值 $\textit{pre}$ 的计算方式同理，可以和答案一起算，所以只需要一个变量。

那么答案就是

$$
\textit{pre} + \textit{nums}[j] + \textit{suf}[j+1]
$$

的最小值。

[视频讲解](https://www.bilibili.com/video/BV12w411B7ia/) 第二题。

```py [sol-Python3]
class Solution:
    def minimumSum(self, nums: List[int]) -> int:
        n = len(nums)
        suf = [0] * n
        suf[-1] = nums[-1]  # 后缀最小值
        for i in range(n - 2, 1, -1):
            suf[i] = min(suf[i + 1], nums[i])

        ans = inf
        pre = nums[0]  # 前缀最小值
        for j in range(1, n - 1):
            if pre < nums[j] > suf[j + 1]:  # 山形
                ans = min(ans, pre + nums[j] + suf[j + 1])  # 更新答案
            pre = min(pre, nums[j])
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumSum(int[] nums) {
        int n = nums.length;
        int[] suf = new int[n]; // 后缀最小值
        suf[n - 1] = nums[n - 1];
        for (int i = n - 2; i > 1; i--) {
            suf[i] = Math.min(suf[i + 1], nums[i]);
        }

        int ans = Integer.MAX_VALUE;
        int pre = nums[0]; // 前缀最小值
        for (int j = 1; j < n - 1; j++) {
            if (pre < nums[j] && nums[j] > suf[j + 1]) { // 山形
                ans = Math.min(ans, pre + nums[j] + suf[j + 1]); // 更新答案
            }
            pre = Math.min(pre, nums[j]);
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSum(vector<int> &nums) {
        int n = nums.size();
        vector<int> suf(n); // 后缀最小值
        suf[n - 1] = nums[n - 1];
        for (int i = n - 2; i > 1; i--) {
            suf[i] = min(suf[i + 1], nums[i]);
        }

        int ans = INT_MAX;
        int pre = nums[0]; // 前缀最小值
        for (int j = 1; j < n - 1; j++) {
            if (pre < nums[j] && nums[j] > suf[j + 1]) { // 山形
                ans = min(ans, pre + nums[j] + suf[j + 1]); // 更新答案
            }
            pre = min(pre, nums[j]);
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
func minimumSum(nums []int) int {
	n := len(nums)
	suf := make([]int, n) // 后缀最小值
	suf[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans := math.MaxInt
	pre := nums[0] // 前缀最小值
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] { // 山形
			ans = min(ans, pre+nums[j]+suf[j+1]) // 更新答案
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

```js [sol-JavaScript]
var minimumSum = function(nums) {
    const n = nums.length;
    const suf = Array(n); // 后缀最小值
    suf[n - 1] = nums[n - 1];
    for (let i = n - 2; i > 1; i--) {
        suf[i] = Math.min(suf[i + 1], nums[i]);
    }

    let ans = Infinity;
    let pre = nums[0]; // 前缀最小值
    for (let j = 1; j < n - 1; j++) {
        if (pre < nums[j] && nums[j] > suf[j + 1]) { // 山形
            ans = Math.min(ans, pre + nums[j] + suf[j + 1]); // 更新答案
        }
        pre = Math.min(pre, nums[j]);
    }
    return ans === Infinity ? -1 : ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_sum(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut suf = vec![0; n]; // 后缀最小值
        suf[n - 1] = nums[n - 1];
        for i in (2..n - 1).rev() {
            suf[i] = suf[i + 1].min(nums[i]);
        }

        let mut ans = i32::MAX;
        let mut pre = nums[0]; // 前缀最小值
        for j in 1..n - 1 {
            if pre < nums[j] && nums[j] > suf[j + 1] { // 山形
                ans = ans.min(pre + nums[j] + suf[j + 1]); // 更新答案
            }
            pre = pre.min(nums[j]);
        }
        if ans == i32::MAX { -1 } else { ans }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 数据结构题单的「**§0.2 枚举中间**」。
2. 动态规划题单的「**专题：前后缀分解**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
