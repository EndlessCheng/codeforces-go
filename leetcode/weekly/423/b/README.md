遍历 $\textit{nums}$，寻找严格递增段（子数组）。

示例 1 有四个严格递增段 $[2,5,7,8,9],[2,3,4],[3],[1]$。

- 在 $[2,5,7,8,9]$ 中，我们可以取出两个相邻的长度都是 $\left\lfloor\dfrac{5}{2}\right\rfloor=2$ 的严格递增段，比如 $[5,7]$ 和 $[8,9]$。
- 或者，取出 $[2,5,7,8,9]$ 的后缀 $[7,8,9]$，它与下一个递增段 $[2,3,4]$ 是相邻的，长度相同。

一般地，设当前严格递增段的长度为 $\textit{cnt}$，上一个严格递增段的长度为 $\textit{preCnt}$。

答案有两种情况：

- 两个子数组属于同一个严格递增段，那么 $k$ 最大是 $\left\lfloor\dfrac{\textit{cnt}}{2}\right\rfloor$。
- 两个子数组分别属于一对相邻的严格递增段，那么 $k$ 最大是 $\min(\textit{preCnt}, \textit{cnt})$。

[本题视频讲解](https://www.bilibili.com/video/BV1JVmBYvEnD/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxIncreasingSubarrays(self, nums: List[int]) -> int:
        ans = pre_cnt = cnt = 0
        for i, x in enumerate(nums):
            cnt += 1
            if i == len(nums) - 1 or x >= nums[i + 1]:  # i 是严格递增段的末尾
                ans = max(ans, cnt // 2, min(pre_cnt, cnt))
                pre_cnt = cnt
                cnt = 0
        return ans
```

```java [sol-Java]
class Solution {
    public int maxIncreasingSubarrays(List<Integer> nums) {
        int ans = 0;
        int preCnt = 0;
        int cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            cnt++;
            // i 是严格递增段的末尾
            if (i == nums.size() - 1 || nums.get(i) >= nums.get(i + 1)) {
                ans = Math.max(ans, Math.max(cnt / 2, Math.min(preCnt, cnt)));
                preCnt = cnt;
                cnt = 0;
            }
        }
        return ans;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public int maxIncreasingSubarrays(List<Integer> nums) {
        Integer[] a = nums.toArray(Integer[]::new); // 转成数组处理，更快
        int ans = 0;
        int preCnt = 0;
        int cnt = 0;
        for (int i = 0; i < a.length; i++) {
            cnt++;
            // i 是严格递增段的末尾
            if (i == a.length - 1 || a[i] >= a[i + 1]) {
                ans = Math.max(ans, Math.max(cnt / 2, Math.min(preCnt, cnt)));
                preCnt = cnt;
                cnt = 0;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxIncreasingSubarrays(vector<int>& nums) {
        int ans = 0, pre_cnt = 0, cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            cnt++;
            if (i == nums.size() - 1 || nums[i] >= nums[i + 1]) { // i 是严格递增段的末尾
                ans = max({ans, cnt / 2, min(pre_cnt, cnt)});
                pre_cnt = cnt;
                cnt = 0;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maxIncreasingSubarrays(int* nums, int numsSize) {
    int ans = 0, pre_cnt = 0, cnt = 0;
    for (int i = 0; i < numsSize; i++) {
        cnt++;
        if (i == numsSize - 1 || nums[i] >= nums[i + 1]) { // i 是严格递增段的末尾
            ans = MAX(ans, MAX(cnt / 2, MIN(pre_cnt, cnt)));
            pre_cnt = cnt;
            cnt = 0;
        }
    }
    return ans;
}
```

```go [sol-Go]
func maxIncreasingSubarrays(nums []int) (ans int) {
	preCnt, cnt := 0, 0
	for i, x := range nums {
		cnt++
		if i == len(nums)-1 || x >= nums[i+1] { // i 是严格递增段的末尾
			ans = max(ans, cnt/2, min(preCnt, cnt))
			preCnt = cnt
			cnt = 0
		}
	}
	return
}
```

```js [sol-JavaScript]
var maxIncreasingSubarrays = function(nums) {
    let ans = 0, preCnt = 0, cnt = 0;
    for (let i = 0; i < nums.length; i++) {
        cnt++;
        if (i === nums.length - 1 || nums[i] >= nums[i + 1]) { // i 是严格递增段的末尾
            ans = Math.max(ans, Math.floor(cnt / 2), Math.min(preCnt, cnt));
            preCnt = cnt;
            cnt = 0;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_increasing_subarrays(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut pre_cnt = 0;
        let mut cnt = 0;
        for i in 0..nums.len() {
            cnt += 1;
            if i == nums.len() - 1 || nums[i] >= nums[i + 1] { // i 是严格递增段的末尾
                ans = ans.max(cnt / 2).max(pre_cnt.min(cnt));
                pre_cnt = cnt;
                cnt = 0;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面双指针题单的「**六、分组循环**」。

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
