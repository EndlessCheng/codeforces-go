[视频讲解](https://www.bilibili.com/video/BV1994y1A7oo/)

注意本题的子数组不要求是连续的。

既然元素的顺序并不重要，我们应当尽量把相近的数字都放在一起。

所以把数组排序后，从小到大三个三个地切分即可。

> 注：题目保证数组长度是 $3$ 的倍数。

```py [sol-Python3]
class Solution:
    def divideArray(self, nums: List[int], k: int) -> List[List[int]]:
        nums.sort()
        ans = []
        for i in range(2, len(nums), 3):
            if nums[i] - nums[i - 2] > k:
                return []
            ans.append(nums[i - 2: i + 1])
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] divideArray(int[] nums, int k) {
        Arrays.sort(nums);
        int n = nums.length;
        int[][] ans = new int[n / 3][3];
        for (int i = 2; i < n; i += 3) {
            if (nums[i] - nums[i - 2] > k) {
                return new int[][]{};
            }
            ans[i / 3] = new int[]{nums[i - 2], nums[i - 1], nums[i]};
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> divideArray(vector<int>& nums, int k) {
        ranges::sort(nums);
        vector<vector<int>> ans;
        for (int i = 2; i < nums.size(); i += 3) {
            if (nums[i] - nums[i - 2] > k) {
                return {};
            }
            ans.push_back({nums[i - 2], nums[i - 1], nums[i]});
        }
        return ans;
    }
};
```

```go [sol-Go]
func divideArray(nums []int, k int) (ans [][]int) {
	slices.Sort(nums)
	for i := 2; i < len(nums); i += 3 {
		if nums[i]-nums[i-2] > k {
			return nil
		}
		ans = append(ans, nums[i-2:i+1])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
