按题意模拟即可。

```py [sol-Python3]
class Solution:
    def lastVisitedIntegers(self, nums: List[int]) -> List[int]:
        ans = []
        seen = []
        k = 0
        for x in nums:
            if x > 0:
                seen.append(x)
                k = 0
            else:
                k += 1
                ans.append(-1 if k > len(seen) else seen[-k])  # 倒数第 k 个
        return ans
```

```java [sol-Java]
class Solution {
    public List<Integer> lastVisitedIntegers(int[] nums) {
        List<Integer> ans = new ArrayList<>();
        List<Integer> seen = new ArrayList<>();
        int k = 0;
        for (int x : nums) {
            if (x > 0) {
                seen.add(x);
                k = 0;
            } else {
                ans.add(++k > seen.size() ? -1 : seen.get(seen.size() - k)); // 倒数第 k 个
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> lastVisitedIntegers(vector<int>& nums) {
        vector<int> ans, seen;
        int k = 0;
        for (int x : nums) {
            if (x > 0) {
                seen.push_back(x);
                k = 0;
            } else {
                ans.push_back(++k > seen.size() ? -1 : seen[seen.size() - k]); // 倒数第 k 个
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func lastVisitedIntegers(nums []int) (ans []int) {
	seen := []int{}
	k := 0
	for _, x := range nums {
		if x > 0 {
			seen = append(seen, x)
			k = 0
		} else {
			k++
			if k > len(seen) {
				ans = append(ans, -1)
			} else {
				ans = append(ans, seen[len(seen)-k]) // 倒数第 k 个
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
