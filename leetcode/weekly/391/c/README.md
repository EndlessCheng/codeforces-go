请看 [视频讲解](https://www.bilibili.com/video/BV1fq421A7CY/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def countAlternatingSubarrays(self, nums: List[int]) -> int:
        ans = cnt = 0
        for i, x in enumerate(nums):
            cnt = 1 if i and x == nums[i - 1] else cnt + 1
            ans += cnt  # 有 cnt 个以 i 为右端点的交替子数组
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
        long ans = 0, cnt = 0;
        for (int i = 0; i < nums.length; i++) {
            cnt = i > 0 && nums[i] == nums[i - 1] ? 1 : cnt + 1;
            ans += cnt; // 有 cnt 个以 i 为右端点的交替子数组
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countAlternatingSubarrays(vector<int> &nums) {
        long long ans = 0, cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            i && nums[i] == nums[i - 1] ? cnt = 1 : cnt++;
            ans += cnt; // 有 cnt 个以 i 为右端点的交替子数组
        }
        return ans;
    }
};
```

```go [sol-Go]
func countAlternatingSubarrays(nums []int) (ans int64) {
	cnt := 0 // 连续交替长度
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
