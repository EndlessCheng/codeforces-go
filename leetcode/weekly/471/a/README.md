用哈希表（或者数组）统计每个元素的出现次数。

遍历哈希表中的键值对 $(x,c)$，如果 $c$ 是 $k$ 的倍数，那么把 $c$ 个 $x$ 加入答案。

[本题视频讲解](https://www.bilibili.com/video/BV1FJ4uz1EkN/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def sumDivisibleByK(self, nums: List[int], k: int) -> int:
        cnt = Counter(nums)
        ans = 0
        for x, c in cnt.items():
            if c % k == 0:
                ans += x * c
        return ans
```

```py [sol-Python3 一行]
class Solution:
    def sumDivisibleByK(self, nums: List[int], k: int) -> int:
        return sum(x * c for x, c in Counter(nums).items() if c % k == 0)
```

```java [sol-Java]
class Solution {
    public int sumDivisibleByK(int[] nums, int k) {
        Map<Integer, Integer> cnt = new HashMap<>(); // 更快的写法见【Java 数组】
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
        }

        int ans = 0;
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) {
            int x = e.getKey(), c = e.getValue();
            if (c % k == 0) {
                ans += x * c;
            }
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int sumDivisibleByK(int[] nums, int k) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] cnt = new int[mx + 1];
        for (int x : nums) {
            cnt[x]++;
        }

        int ans = 0;
        for (int x = 1; x <= mx; x++) {
            if (cnt[x] % k == 0) {
                ans += x * cnt[x];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumDivisibleByK(vector<int>& nums, int k) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }

        int ans = 0;
        for (auto& [x, c] : cnt) {
            if (c % k == 0) {
                ans += x * c;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumDivisibleByK(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	for x, c := range cnt {
		if c%k == 0 {
			ans += x * c
		}
	}
	return
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
