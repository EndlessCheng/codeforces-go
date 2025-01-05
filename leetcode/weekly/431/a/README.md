暴力枚举所有子数组，计算子数组的元素乘积、LCM 和 GCD。

注意乘积太大会溢出。设 $\textit{allLcm}$ 为所有元素的 LCM，我们可以枚举到子数组乘积超过 $\textit{allLcm}\cdot \max(\textit{nums})$ 为止，这二者分别是 LCM 的最大值和 GCD 的最大值。

注意题目保证 $n\ge 2$，又由于两个数一定满足题目要求，所以合法子数组是一定存在的，且长度至少是 $2$。

> 注：$\text{LCM}(1,2,\ldots,10)=2520$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV18srKYLEd8/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxLength(self, nums: List[int]) -> int:
        max_m = lcm(*nums) * max(nums)
        ans = 0
        for i in range(len(nums)):
            m, l, g = 1, 1, 0
            for j in range(i, len(nums)):
                x = nums[j]
                m *= x
                l = lcm(l, x)
                g = gcd(g, x)
                if m == l * g:
                    ans = max(ans, j - i + 1)
                if m > max_m:
                    break
        return ans
```

```java [sol-Java]
class Solution {
    public int maxLength(int[] nums) {
        int mx = Arrays.stream(nums).max().getAsInt();
        int allLcm = 1;
        for (int x : nums) {
            allLcm = lcm(allLcm, x);
        }

        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int m = 1;
            int l = 1;
            int g = 0;
            for (int j = i; j < nums.length && m <= allLcm * mx; j++) {
                int x = nums[j];
                m *= x;
                l = lcm(l, x);
                g = gcd(g, x);
                if (m == l * g) {
                    ans = Math.max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxLength(vector<int>& nums) {
        int mx = ranges::max(nums);
        int all_lcm = 1;
        for (int x : nums) {
            all_lcm = lcm(all_lcm, x);
        }

        int ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            int m = 1, l = 1, g = 0;
            for (int j = i; j < nums.size() && m <= all_lcm * mx; j++) {
                int x = nums[j];
                m *= x;
                l = lcm(l, x);
                g = gcd(g, x);
                if (m == l * g) {
                    ans = max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxLength(nums []int) (ans int) {
	mx := slices.Max(nums)
	allLcm := 1
	for _, x := range nums {
		allLcm = lcm(allLcm, x)
	}

	for i := range nums {
		m, l, g := 1, 1, 0
		for j := i; j < len(nums) && m <= allLcm*mx; j++ {
			x := nums[j]
			m *= x
			l = lcm(l, x)
			g = gcd(g, x)
			if m == l*g {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。计算 LCM 和 GCD 的时间视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
