题目要求子序列的 AND 结果非零，也就是说，AND 结果的某个比特位一定是 $1$。

**枚举**这个比特位。

示例 1 的 $\textit{nums}=[5,4,7]$，二进制长度最多为 $3$。

- 如果 AND 结果的最低位是 $1$，那么只有 $\textit{nums}$ 中的 $[5,7]$ 能在子序列中，问题变成 $[5,7]$ 的 [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)。注意本题 $n\le 10^5$，必须用二分优化，见 [我的题解](https://leetcode.cn/problems/longest-increasing-subsequence/solutions/2147040/jiao-ni-yi-bu-bu-si-kao-dpfu-o1-kong-jia-4zma/)。
- 如果 AND 结果的次低位是 $1$，那么只有 $\textit{nums}$ 中的 $[7]$ 能在子序列中。
- 如果 AND 结果的最高位是 $1$，那么 $\textit{nums}$ 中的 $[5,4,7]$ 都能在子序列中。

设 $\max(\textit{nums})$ 的二进制长度为 $w$，枚举 $i=0,1,2,\ldots,w-1$，对 $\textit{nums}$ 中的二进制第 $i$ 位是 $1$ 的数，计算最长递增子序列。

> **注**：也可以只枚举 $\textit{nums}$ 所有元素的 OR 中的 $1$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        ans = 0
        w = max(nums).bit_length()
        for i in range(w):
            bit = 1 << i
            # 300. 最长递增子序列
            f = []
            for x in nums:
                if x & bit == 0:  # x 二进制的第 i 位是 0
                    continue
                j = bisect_left(f, x)
                if j < len(f):
                    f[j] = x
                else:
                    f.append(x)
            ans = max(ans, len(f))
        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubsequence(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int w = 32 - Integer.numberOfLeadingZeros(mx);
        int ans = 0;

        // 数组比 ArrayList 快，读者可以对比另一份代码【Java ArrayList】
        int[] f = new int[nums.length];

        for (int i = 0; i < w; i++) {
            // 300. 最长递增子序列
            int size = 0;
            for (int x : nums) {
                if ((x >> i & 1) == 0) { // x 二进制的第 i 位是 0
                    continue;
                }

                int j = Arrays.binarySearch(f, 0, size, x);
                if (j < 0) {
                    j = ~j;
                }

                f[j] = x;
                if (j == size) {
                    size++;
                }
            }
            ans = Math.max(ans, size);
        }

        return ans;
    }
}
```

```java [sol-Java ArrayList]
class Solution {
    public int longestSubsequence(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int w = 32 - Integer.numberOfLeadingZeros(mx);
        int ans = 0;

        for (int i = 0; i < w; i++) {
            // 300. 最长递增子序列
            List<Integer> f = new ArrayList<>();
            for (int x : nums) {
                if ((x >> i & 1) == 0) { // x 二进制的第 i 位是 0
                    continue;
                }

                int j = Collections.binarySearch(f, x);
                if (j < 0) {
                    j = ~j;
                }

                if (j < f.size()) {
                    f.set(j, x);
                } else {
                    f.add(x);
                }
            }
            ans = Math.max(ans, f.size());
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubsequence(vector<int>& nums) {
        int ans = 0;
        int w = bit_width((uint32_t) ranges::max(nums));
        for (int i = 0; i < w; i++) {
            // 300. 最长递增子序列
            vector<int> f;
            for (int x : nums) {
                if ((x >> i & 1) == 0) { // x 二进制的第 i 位是 0
                    continue;
                }
                auto it = ranges::lower_bound(f, x);
                if (it != f.end()) {
                    *it = x;
                } else {
                    f.push_back(x);
                }
            }
            ans = max(ans, (int) f.size());
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSubsequence(nums []int) (ans int) {
	w := bits.Len(uint(slices.Max(nums)))
	for i := range w {
		// 300. 最长递增子序列
		f := []int{}
		for _, x := range nums {
			if x>>i&1 == 0 { // x 二进制的第 i 位是 0
				continue
			}
			j := sort.SearchInts(f, x)
			if j < len(f) {
				f[j] = x
			} else {
				f = append(f, x)
			}
		}
		ans = max(ans, len(f))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n \log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§4.2 最长递增子序列**」。

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
