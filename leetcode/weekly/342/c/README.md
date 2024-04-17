### 本题视频讲解

见[【周赛 342】](https://www.bilibili.com/video/BV1Bs4y1A7Wa/)第三题。

### 思路

滑动窗口。由于值域很小，所以借鉴**计数排序**，用一个 $\textit{cnt}$ 数组维护窗口内每个数的出现次数。然后遍历 $\textit{cnt}$ 去求第 $x$ 小的数。

什么是第 $x$ 小的数？

设它是 $\textit{num}$，那么 $<\textit{num}$ 的数有 $<x$ 个，$\le\textit{num}$ 的数有 $\ge x$ 个，就说明 $\textit{num}$ 是第 $x$ 小的数。

```Python [sol1-Python3]
class Solution:
    def getSubarrayBeauty(self, nums: List[int], k: int, x: int) -> List[int]:
        cnt = [0] * 101
        for num in nums[:k - 1]:  # 先往窗口内添加 k-1 个数
            cnt[num] += 1
        ans = [0] * (len(nums) - k + 1)
        for i, (in_, out) in enumerate(zip(nums[k - 1:], nums)):
            cnt[in_] += 1  # 进入窗口（保证窗口有恰好 k 个数）
            left = x
            for j in range(-50, 0):  # 暴力枚举负数范围 [-50,-1]
                left -= cnt[j]
                if left <= 0:  # 找到美丽值
                    ans[i] = j
                    break
            cnt[out] -= 1  # 离开窗口
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] getSubarrayBeauty(int[] nums, int k, int x) {
        final int BIAS = 50;
        var cnt = new int[BIAS * 2 + 1];
        int n = nums.length;
        for (int i = 0; i < k - 1; ++i) // 先往窗口内添加 k-1 个数
            ++cnt[nums[i] + BIAS];
        var ans = new int[n - k + 1];
        for (int i = k - 1; i < n; ++i) {
            ++cnt[nums[i] + BIAS]; // 进入窗口（保证窗口有恰好 k 个数）
            int left = x;
            for (int j = 0; j < BIAS; ++j) { // 暴力枚举负数范围 [-50,-1]
                left -= cnt[j];
                if (left <= 0) { // 找到美丽值
                    ans[i - k + 1] = j - BIAS;
                    break;
                }
            }
            --cnt[nums[i - k + 1] + BIAS]; // 离开窗口
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> getSubarrayBeauty(vector<int> &nums, int k, int x) {
        const int BIAS = 50;
        int cnt[BIAS * 2 + 1]{}, n = nums.size();
        for (int i = 0; i < k - 1; ++i) // 先往窗口内添加 k-1 个数
            ++cnt[nums[i] + BIAS];
        vector<int> ans(n - k + 1);
        for (int i = k - 1; i < n; ++i) {
            ++cnt[nums[i] + BIAS]; // 进入窗口（保证窗口有恰好 k 个数）
            int left = x;
            for (int j = 0; j < BIAS; ++j) { // 暴力枚举负数范围 [-50,-1]
                left -= cnt[j];
                if (left <= 0) { // 找到美丽值
                    ans[i - k + 1] = j - BIAS;
                    break;
                }
            }
            --cnt[nums[i - k + 1] + BIAS]; // 离开窗口
        }
        return ans;
    }
};
```

```go [sol1-Go]
func getSubarrayBeauty(nums []int, k, x int) []int {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] { // 先往窗口内添加 k-1 个数
		cnt[num+bias]++
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+bias]++ // 进入窗口（保证窗口有恰好 k 个数）
		left := x
		for j, c := range cnt[:bias] { // 暴力枚举负数范围 [-50,-1]
			left -= c
			if left <= 0 { // 找到美丽值
				ans[i] = j - bias
				break
			}
		}
		cnt[nums[i]+bias]-- // 离开窗口
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=50$。
- 空间复杂度：$\mathcal{O}(U)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
