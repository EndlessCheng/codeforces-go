用哈希表统计 $\textit{nums}$ 每个数的出现次数，然后再用一个哈希表（或者数组）统计出现次数的出现次数。

最后从左到右遍历 $\textit{nums}$，第一个出现次数的出现次数是 $1$ 的 $\textit{nums}[i]$ 就是答案。

[本题视频讲解](https://www.bilibili.com/video/BV1VgZ4BCETj/?t=2m45s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def firstUniqueFreq(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        cc = Counter(cnt.values())

        for x in nums:
            if cc[cnt[x]] == 1:
                return x
        return -1
```

```java [sol-Java]
class Solution {
    public int firstUniqueFreq(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum);
        }

        int[] cc = new int[nums.length + 1];
        for (int c : cnt.values()) {
            cc[c]++;
        }

        for (int x : nums) {
            if (cc[cnt.get(x)] == 1) {
                return x;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int firstUniqueFreq(vector<int>& nums) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }

        vector<int> cc(nums.size() + 1);
        for (auto& [_, c] : cnt) {
            cc[c]++;
        }

        for (int x : nums) {
            if (cc[cnt[x]] == 1) {
                return x;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func firstUniqueFreq(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	cc := make([]int, len(nums)+1)
	for _, c := range cnt {
		cc[c]++
	}

	for _, x := range nums {
		if cc[cnt[x]] == 1 {
			return x
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
