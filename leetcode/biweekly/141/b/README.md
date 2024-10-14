例如 $x=100111$，那么 $x\ |\ (x+1) = 100111\ |\ 101000 = 101111$。

可以发现，$x\ |\ (x+1)$ 的本质是把二进制最右边的 $0$ 置为 $1$。

反过来，如果我们知道了 $x\ |\ (x+1)$ 的结果 $101111$，那么对应的 $x$ 只能是这些：

- $100111$。
- $101011$。
- $101101$。
- $101110$。

其中最小的是 $100111$，也就是把 $101111$ 最右边的 $0$ 的右边的 $1$ 置为 $0$。

由于 $x\ |\ (x+1)$ 最低位一定是 $1$（因为 $x$ 和 $x+1$ 其中一定有一个奇数），所以如果 $\textit{nums}[i]$ 是偶数（质数中只有 $2$），那么无解，答案为 $-1$。

## 写法一

把 $101111$ 取反，得 $010000$，其 $\text{lowbit}=10000$，右移一位得 $1000$。把 $101111$ 与 $1000$ 异或，即可得到 $100111$。

关于 $\text{lowbit}$ 的原理，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

[本题视频讲解](https://www.bilibili.com/video/BV1iR2zYaESG/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        for i, x in enumerate(nums):
            if x == 2:
                nums[i] = -1
            else:
                t = ~x
                nums[i] ^= (t & -t) >> 1
        return nums
```

```java [sol-Java]
class Solution {
    public int[] minBitwiseArray(List<Integer> nums) {
        int n = nums.size();
        int[] ans = new int[n];
        for (int i = 0; i < n; i++) {
            int x = nums.get(i);
            if (x == 2) {
                ans[i] = -1;
            } else {
                int t = ~x;
                int invLb = t & -t;
                ans[i] = x ^ (invLb >> 1);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minBitwiseArray(vector<int>& nums) {
        for (int& x : nums) { // 注意这里是引用
            if (x == 2) {
                x = -1;
            } else {
                int t = ~x;
                x ^= (t & -t) >> 1;
            }
        }
        return nums;
    }
};
```

```go [sol-Go]
func minBitwiseArray(nums []int) []int {
	for i, x := range nums {
		if x == 2 {
			nums[i] = -1
		} else {
			t := ^x
			nums[i] ^= t & -t >> 1
		}
	}
	return nums
}
```

## 写法二

把 $101111$ 加一，得 $110000$，再 AND $101111$ 取反后的值 $010000$，可以得到方法一中的 $\text{lowbit}=10000$。

```py [sol-Python3]
class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        for i, x in enumerate(nums):
            if x == 2:
                nums[i] = -1
            else:
                inv_lb = (x + 1) & ~x
                nums[i] ^= inv_lb >> 1
        return nums
```

```java [sol-Java]
class Solution {
    public int[] minBitwiseArray(List<Integer> nums) {
        int n = nums.size();
        int[] ans = new int[n];
        for (int i = 0; i < n; i++) {
            int x = nums.get(i);
            if (x == 2) {
                ans[i] = -1;
            } else {
                int invLb = (x + 1) & ~x;
                ans[i] = x ^ (invLb >> 1);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minBitwiseArray(vector<int>& nums) {
        for (int& x : nums) { // 注意这里是引用
            if (x == 2) {
                x = -1;
            } else {
                int inv_lb = (x + 1) & ~x;
                x ^= inv_lb >> 1;
            }
        }
        return nums;
    }
};
```

```go [sol-Go]
func minBitwiseArray(nums []int) []int {
	for i, x := range nums {
		if x == 2 {
			nums[i] = -1
		} else {
			nums[i] ^= (x + 1) &^ x >> 1
		}
	}
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
