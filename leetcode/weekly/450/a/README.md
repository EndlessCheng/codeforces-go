核心是计算数位和。

我们可以不用转成字符串处理，而是不断取最低位（模 $10$），去掉最低位（除以 $10$），直到数字为 $0$。

例如 $\textit{num}=123$：

1. 初始化 $x=\textit{num}$。
2. 通过 $x\bmod 10$ 取到个位数 $3$，然后把 $x$ 除以 $10$（下取整），得到 $x=12$。
3. 再次 $x\bmod 10$ 取到十位数 $2$，然后把 $x$ 除以 $10$（下取整），得到 $x=1$。
4. 最后 $x\bmod 10$ 取到百位数 $1$，然后把 $x$ 除以 $10$（下取整），得到 $x=0$。此时完成了遍历 $\textit{num}$ 的每个数位，退出循环。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

## 优化前

```py [sol-Python3 用字符串]
class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        for i, x in enumerate(nums):
            if sum(map(int, str(x))) == i:
                return i
        return -1
```

```py [sol-Python3 不用字符串]
class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        for i, x in enumerate(nums):
            s = 0
            while x > 0:
                s += x % 10
                x //= 10
            if s == i:
                return i
        return -1
```

```java [sol-Java]
class Solution {
    public int smallestIndex(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            int s = 0;
            for (int x = nums[i]; x > 0; x /= 10) {
                s += x % 10;
            }
            if (s == i) {
                return i;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestIndex(vector<int>& nums) {
        for (int i = 0; i < nums.size(); i++) {
            int s = 0;
            for (int x = nums[i]; x > 0; x /= 10) {
                s += x % 10;
            }
            if (s == i) {
                return i;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func smallestIndex(nums []int) int {
	for i, x := range nums {
		s := 0
		for ; x > 0; x /= 10 {
			s += x % 10
		}
		if s == i {
			return i
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 优化

根据本题的数据范围，数位和至多为 $999$ 的数位和，即 $3\times 9=27$。

所以至多枚举到 $i=27$ 即可。

```py [sol-Python3]
class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        for i, x in enumerate(nums[:28]):
            s = 0
            while x > 0:
                s += x % 10
                x //= 10
            if s == i:
                return i
        return -1
```

```java [sol-Java]
class Solution {
    public int smallestIndex(int[] nums) {
        int n = Math.min(nums.length, 28);
        for (int i = 0; i < n; i++) {
            int s = 0;
            for (int x = nums[i]; x > 0; x /= 10) {
                s += x % 10;
            }
            if (s == i) {
                return i;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestIndex(vector<int>& nums) {
        int n = min((int) nums.size(), 28);
        for (int i = 0; i < n; i++) {
            int s = 0;
            for (int x = nums[i]; x > 0; x /= 10) {
                s += x % 10;
            }
            if (s == i) {
                return i;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func smallestIndex(nums []int) int {
	for i, x := range nums[:min(len(nums), 28)] {
		s := 0
		for ; x > 0; x /= 10 {
			s += x % 10
		}
		if s == i {
			return i
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\min(n,k)\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$k=28$，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

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
