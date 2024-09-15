本题和 [2965. 找出缺失和重复的数字](https://leetcode.cn/problems/find-missing-and-repeated-values/) 本质是一样的，见 [我的题解](https://leetcode.cn/problems/find-missing-and-repeated-values/solutions/2569783/mo-ni-pythonjavacgo-by-endlesscheng-mexz/)，有位运算和数学两种做法。

## 位运算

需要两次遍历。一次遍历见下面的数学做法。

```py [sol-Python3]
class Solution:
    def getSneakyNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums) - 2
        xor_all = n ^ (n + 1)  # n 和 n+1 多异或了
        for i, x in enumerate(nums):
            xor_all ^= i ^ x
        shift = xor_all.bit_length() - 1

        ans = [0, 0]
        for i, x in enumerate(nums):
            if i < n:
                ans[i >> shift & 1] ^= i
            ans[x >> shift & 1] ^= x
        return ans
```

```java [sol-Java]
class Solution {
    public int[] getSneakyNumbers(int[] nums) {
        int n = nums.length - 2;
        int xorAll = n ^ (n + 1); // n 和 n+1 多异或了
        for (int i = 0; i < nums.length; i++) {
            xorAll ^= i ^ nums[i];
        }
        int shift = Integer.numberOfTrailingZeros(xorAll);

        int[] ans = new int[2];
        for (int i = 0; i < nums.length; i++) {
            if (i < n) {
                ans[i >> shift & 1] ^= i;
            }
            ans[nums[i] >> shift & 1] ^= nums[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> getSneakyNumbers(vector<int>& nums) {
        int n = nums.size() - 2;
        int xor_all = n ^ (n + 1); // n 和 n+1 多异或了
        for (int i = 0; i < nums.size(); i++) {
            xor_all ^= i ^ nums[i];
        }
        int shift = __builtin_ctz(xor_all);

        vector<int> ans(2);
        for (int i = 0; i < nums.size(); i++) {
            if (i < n) {
                ans[i >> shift & 1] ^= i;
            }
            ans[nums[i] >> shift & 1] ^= nums[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func getSneakyNumbers(nums []int) []int {
	n := len(nums) - 2
	xorAll := n ^ (n + 1) // n 和 n+1 多异或了
	for i, x := range nums {
		xorAll ^= i ^ x
	}
	shift := bits.TrailingZeros(uint(xorAll))

	ans := make([]int, 2)
	for i, x := range nums {
		if i < n {
			ans[i>>shift&1] ^= i
		}
		ans[x>>shift&1] ^= x
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 数学

设多出的两个数分别为 $x$ 和 $y$。

也就是说，$\textit{nums} = [0,1,2,\cdots,n-1,x,y]$。

设 $\textit{nums}$ 的元素和为 $s$，$\textit{nums}$ 的元素平方之和为 $s_2$，那么有

$$
\begin{aligned}
&x+y = s - (0 + 1 + 2 + \cdots + n-1) = a     \\
&x^2+y^2 = s_2 - (0^2 + 1^2 + 2^2 + \cdots + (n-1)^2) = b   \\
\end{aligned}
$$

解得

$$
\begin{cases}
x  = \dfrac{a-\sqrt{2b-a^2}}{2}     \\
y  = \dfrac{a+\sqrt{2b-a^2}}{2}    \\
\end{cases}
$$

也可以先算出 $x$，然后算出 $y=a-x$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Qp4me2Emz/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def getSneakyNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums) - 2
        a = -n * (n - 1) // 2
        b = -n * (n - 1) * (n * 2 - 1) // 6
        for x in nums:
            a += x
            b += x * x
        x = int((a - sqrt(b * 2 - a * a)) / 2)
        return [x, a - x]
```

```java [sol-Java]
class Solution {
    public int[] getSneakyNumbers(int[] nums) {
        int n = nums.length - 2;
        int a = -n * (n - 1) / 2;
        int b = -n * (n - 1) * (n * 2 - 1) / 6;
        for (int x : nums) {
            a += x;
            b += x * x;
        }
        int x = (int) ((a - Math.sqrt(b * 2 - a * a)) / 2);
        return new int[]{x, a - x};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> getSneakyNumbers(vector<int>& nums) {
        int n = nums.size() - 2;
        int a = -n * (n - 1) / 2;
        int b = -n * (n - 1) * (n * 2 - 1) / 6;
        for (int x : nums) {
            a += x;
            b += x * x;
        }
        int x = (a - sqrt(b * 2 - a * a)) / 2;
        return {x, a - x};
    }
};
```

```go [sol-Go]
func getSneakyNumbers(nums []int) []int {
	n := len(nums) - 2
	a := -n * (n - 1) / 2
	b := -n * (n - 1) * (n*2 - 1) / 6
	for _, x := range nums {
		a += x
		b += x * x
	}
	x := int((float64(a) - math.Sqrt(float64(b*2-a*a))) / 2)
	return []int{x, a - x}
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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
