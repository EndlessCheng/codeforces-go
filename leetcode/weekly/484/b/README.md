写个二重循环，枚举子数组的左端点和右端点。

在枚举右端点的同时，用一个哈希集合记录子数组中的元素，用一个变量 $s$ 维护子数组和。

如果 $s$ 在哈希集合中，那么把答案加一。

[本题视频讲解](https://www.bilibili.com/video/BV1tv6dBME7K/?t=5m10s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def centeredSubarrays(self, nums: List[int]) -> int:
        ans = 0
        for i in range(len(nums)):
            st = set()
            s = 0
            for x in nums[i:]:
                st.add(x)
                s += x
                if s in st:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int centeredSubarrays(int[] nums) {
        HashSet<Integer> set = new HashSet<>();
        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            set.clear();
            int s = 0;
            for (int j = i; j < nums.length; j++) {
                int x = nums[j];
                set.add(x);
                s += x;
                if (set.contains(s)) {
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int centeredSubarrays(vector<int>& nums) {
        unordered_set<int> st;
        int ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            st.clear();
            int s = 0;
            for (int j = i; j < nums.size(); j++) {
                int x = nums[j];
                st.insert(x);
                s += x;
                ans += st.contains(s);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func centeredSubarrays(nums []int) (ans int) {
	has := map[int]int{}
	for i := range nums {
		clear(has)
		s := 0
		for _, x := range nums[i:] {
			has[x] = 1
			s += x
			ans += has[s]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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
