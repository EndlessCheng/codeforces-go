## 方法一

对于曼哈顿距离，由于水平方向的移动和垂直方向的移动互不影响，我们可以把横纵坐标分别计算。

设当前向西走了 $a$ 步，向东走了 $b$ 步。比如 $a=2,\ b=5$。

贪心地，把其中更小的 $a$ 改成向东走。

如果把 $a$ 减少 $d$，那么 $b$ 就能增大 $d$，所以修改后的当前位置的横坐标为

$$
(b+d) - (a-d) = b-a+2d
$$

如果 $b$ 更小，那么横坐标的绝对值为

$$
a-b+2d
$$

综合两种情况，修改后的当前位置的横坐标的绝对值为

$$
|a-b|+2d
$$

其中

$$
d = \min(a,b,k)
$$

然后把 $k$ 减少 $d$，按照同样的方法继续计算纵坐标。

用修改后的横纵坐标绝对值之和，更新答案的最大值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1D5F6eRECp/?t=1m43s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxDistance(self, s: str, k: int) -> int:
        ans = 0
        cnt = defaultdict(int)
        for ch in s:
            cnt[ch] += 1
            left = k
            def f(a: int, b: int) -> int:
                nonlocal left
                d = min(a, b, left)
                left -= d
                return abs(a - b) + d * 2
            ans = max(ans, f(cnt['N'], cnt['S']) + f(cnt['E'], cnt['W']))
        return ans
```

```java [sol-Java]
class Solution {
    private int left;

    public int maxDistance(String s, int k) {
        int ans = 0;
        int[] cnt = new int['X']; // 'W' + 1 = 'X'
        for (char ch : s.toCharArray()) {
            cnt[ch]++;
            left = k;
            ans = Math.max(ans, f(cnt['N'], cnt['S']) + f(cnt['E'], cnt['W']));
        }
        return ans;
    }

    private int f(int a, int b) {
        int d = Math.min(Math.min(a, b), left);
        left -= d;
        return Math.abs(a - b) + d * 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistance(string s, int k) {
        int ans = 0;
        int cnt['X']{}; // 'W' + 1 = 'X'
        for (char ch : s) {
            cnt[ch]++;
            int left = k;
            auto f = [&](int a, int b) -> int {
                int d = min({a, b, left});
                left -= d;
                return abs(a - b) + d * 2;
            };
            ans = max(ans, f(cnt['N'], cnt['S']) + f(cnt['E'], cnt['W']));
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDistance(s string, k int) (ans int) {
	cnt := ['X']int{} // 'W' + 1 = 'X'
	for _, ch := range s {
		cnt[ch]++
		left := k
		f := func(a, b int) int {
			d := min(a, b, left)
			left -= d
			return abs(a-b) + d*2
		}
		ans = max(ans, f(cnt['N'], cnt['S'])+f(cnt['E'], cnt['W']))
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 方法二

设当前位置为 $(x,y)$，那么到原点的曼哈顿距离为 $|x|+|y|$。

通过方法一可知，每操作一次，曼哈顿距离都会增大 $2$，但这不会超过移动的次数 $i+1$。

所以执行完 $s[i]$ 后的答案为

$$
\min(|x|+|y|+2k,i+1)
$$

```py [sol-Python3]
class Solution:
    def maxDistance(self, s: str, k: int) -> int:
        ans = x = y = 0
        for i, c in enumerate(s):
            if c == 'N': y += 1
            elif c == 'S': y -= 1
            elif c == 'E': x += 1
            else: x -= 1
            ans = max(ans, min(abs(x) + abs(y) + k * 2, i + 1))
        return ans
```

```java [sol-Java]
class Solution {
    public int maxDistance(String s, int k) {
        int ans = 0, x = 0, y = 0;
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == 'N') { y++; }
            else if (c == 'S') { y--; }
            else if (c == 'E') { x++; }
            else { x--; }
            ans = Math.max(ans, Math.min(Math.abs(x) + Math.abs(y) + k * 2, i + 1));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistance(string s, int k) {
        int ans = 0, x = 0, y = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == 'N') y++;
            else if (s[i] == 'S') y--;
            else if (s[i] == 'E') x++;
            else x--;
            ans = max(ans, min(abs(x) + abs(y) + k * 2, i + 1));
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDistance(s string, k int) int {
	ans, x, y := 0, 0, 0
	for i, c := range s {
		switch c {
		case 'N': y++
		case 'S': y--
		case 'E': x++
		default:  x--
		}
		ans = max(ans, min(abs(x)+abs(y)+k*2, i+1))
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
