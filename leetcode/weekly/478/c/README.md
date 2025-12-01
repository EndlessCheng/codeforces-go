枚举 $j$，同时用哈希表维护 $j$ 左边的 $\text{reverse}(\textit{nums}[i])$ 的最大下标，哈希表的 key 是 $\text{reverse}(\textit{nums}[i])$，value 是下标 $i$。

如果哈希表中有 $\textit{nums}[j]$，获取对应的下标 $i$，用 $j-i$ 更新答案的最小值。

[本题视频讲解](https://www.bilibili.com/video/BV1D4SiB5Ee3/?t=56m18s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minMirrorPairDistance(self, nums: List[int]) -> int:
        last_index = {}
        ans = inf

        for j, x in enumerate(nums):
            if x in last_index:
                ans = min(ans, j - last_index[x])
            rev = int(str(x)[::-1])
            last_index[rev] = j

        return ans if ans < inf else -1
```

```py [sol-Python3 不用字符串]
class Solution:
    def minMirrorPairDistance(self, nums: List[int]) -> int:
        last_index = {}
        ans = inf

        for j, x in enumerate(nums):
            if x in last_index:
                ans = min(ans, j - last_index[x])

            # 计算 reverse(x)，不用字符串
            rev = 0
            while x > 0:
                x, d = divmod(x, 10)
                rev = rev * 10 + d
            last_index[rev] = j

        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minMirrorPairDistance(int[] nums) {
        int n = nums.length;
        int ans = n;
        Map<Integer, Integer> lastIndex = new HashMap<>(n, 1); // 预分配空间

        for (int j = 0; j < n; j++) {
            int x = nums[j];
            Integer i = lastIndex.get(x);
            if (i != null) {
                ans = Math.min(ans, j - i);
            }

            // 计算 reverse(x)，不用字符串
            int rev = 0;
            for (; x > 0; x /= 10) {
                rev = rev * 10 + x % 10;
            }
            lastIndex.put(rev, j);
        }

        return ans < n ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minMirrorPairDistance(vector<int>& nums) {
        unordered_map<int, int> last_index;
        int n = nums.size(), ans = n;

        for (int j = 0; j < n; j++) {
            int x = nums[j];
            auto it = last_index.find(x);
            if (it != last_index.end()) {
                ans = min(ans, j - it->second);
            }

            // 计算 reverse(x)，不用字符串
            int rev = 0;
            for (; x > 0; x /= 10) {
                rev = rev * 10 + x % 10;
            }
            last_index[rev] = j;
        }

        return ans < n ? ans : -1;
    }
};
```

```go [sol-Go]
func minMirrorPairDistance(nums []int) int {
	n := len(nums)
	ans := n
	lastIndex := make(map[int]int, n) // 预分配空间

	for j, x := range nums {
		if i, ok := lastIndex[x]; ok {
			ans = min(ans, j-i)
		}

		// 计算 reverse(x)，不用字符串
		rev := 0
		for ; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		lastIndex[rev] = j
	}

	if ans == n {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。反转一个数字需要 $\mathcal{O}(\log U)$ 时间。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§0.1 枚举右，维护左**」。

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
