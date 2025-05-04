把 $n$ 的十进制表示视作字符串 $s$，题意是把 $s$ 中的两个**不同位置**上的数相乘，计算最大乘积。

贪心地，相乘的数越大越好，所以答案为 $s$ 最大值与次大值的乘积。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1avVwz5EbY/)，欢迎点赞关注~

## 写法一：排序

```py [sol-Python3]
class Solution:
    def maxProduct(self, n: int) -> int:
        s = sorted(str(n))
        return int(s[-1]) * int(s[-2])
```

```java [sol-Java]
class Solution {
    public int maxProduct(int n) {
        char[] s = String.valueOf(n).toCharArray();
        Arrays.sort(s);
        int m = s.length;
        return (s[m - 1] - '0') * (s[m - 2] - '0');
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxProduct(int n) {
        string s = to_string(n);
        ranges::sort(s);
        int m = s.size();
        return (s[m - 1] - '0') * (s[m - 2] - '0');
    }
};
```

```go [sol-Go]
func maxProduct(n int) int {
	s := []byte(strconv.Itoa(n))
	slices.Sort(s)
	m := len(s)
	return int(s[m-1]-'0') * int(s[m-2]-'0')
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m=\mathcal{O}(\log n)$ 是 $n$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(m)$。

## 写法二：维护最大次大

不用字符串也不排序的写法。

```py [sol-Python3]
class Solution:
    def maxProduct(self, n: int) -> int:
        mx = mx2 = 0
        while n > 0:
            n, d = divmod(n, 10)
            if d > mx:
                mx2 = mx
                mx = d
            elif d > mx2:
                mx2 = d
        return mx * mx2
```

```java [sol-Java]
class Solution {
    public int maxProduct(int n) {
        int mx = 0;
        int mx2 = 0;
        while (n > 0) {
            int d = n % 10;
            if (d > mx) {
                mx2 = mx;
                mx = d;
            } else if (d > mx2) {
                mx2 = d;
            }
            n /= 10;
        }
        return mx * mx2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxProduct(int n) {
        int mx = 0, mx2 = 0;
        while (n > 0) {
            int d = n % 10;
            if (d > mx) {
                mx2 = mx;
                mx = d;
            } else if (d > mx2) {
                mx2 = d;
            }
            n /= 10;
        }
        return mx * mx2;
    }
};
```

```go [sol-Go]
func maxProduct(n int) int {
	mx, mx2 := 0, 0
	for ; n > 0; n /= 10 {
		d := n % 10
		if d > mx {
			mx2 = mx
			mx = d
		} else if d > mx2 {
			mx2 = d
		}
	}
	return mx * mx2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m)$，其中 $m=\mathcal{O}(\log n)$ 是 $n$ 的十进制长度。
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
