设在不移除子串的情况下，最终坐标为 $(x,y)$。

假设我们去掉了子串 $UR$，相当于少算了一次「$y$ 增加一」和一次「$x$ 增加一」，所以最终坐标为 $(x-1,y-1)$。

横看成岭侧成峰，**站在最终坐标的位置看**，去掉子串后的偏移量为 $(-1,-1)$。

问题让我们计算有多少个不同的最终坐标。由于**偏移量不同，最终坐标也不同**，所以问题变成：

- 计算有多少个不同的偏移量。
  
也就是对 $s$ 的长为 $k$ 的子串，计算 $\texttt{UDLR}$ 偏移量之和。

这是一个标准的**定长滑动窗口**问题，原理讲解[【套路】教你解决定长滑窗！适用于所有定长滑窗题目！](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)

[本题视频讲解](https://www.bilibili.com/video/BV1AKnRz8Ejn/?t=5m31s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
DIRS = {
    'L': (-1, 0),
    'R': (1, 0),
    'D': (0, -1),
    'U': (0, 1),
}

class Solution:
    def distinctPoints(self, s: str, k: int) -> int:
        st = set()
        x = y = 0
        for i, c in enumerate(s):
            # 1. 入
            dx, dy = DIRS[c]
            x += dx
            y += dy

            left = i + 1 - k  # 窗口左端点
            if left < 0:  # 窗口大小不足 k
                continue

            # 2. 记录答案
            st.add((x, y))

            # 3. 出，为下一个窗口做准备
            dx, dy = DIRS[s[left]]
            x -= dx
            y -= dy

        return len(st)
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = new int[128][];

    static {
        DIRS['L'] = new int[]{-1, 0};
        DIRS['R'] = new int[]{1, 0};
        DIRS['D'] = new int[]{0, -1};
        DIRS['U'] = new int[]{0, 1};
    }

    public int distinctPoints(String s, int k) {
        int n = s.length();
        Set<Long> set = new HashSet<>();
        int x = 0, y = 0;
        for (int i = 0; i < n; i++) {
            // 1. 入
            char c = s.charAt(i);
            x += DIRS[c][0];
            y += DIRS[c][1];

            int left = i + 1 - k; // 窗口左端点
            if (left < 0) { // 窗口大小不足 k
                continue;
            }

            // 2. 记录答案
            // 把两个 int 压缩到一个 long 中，+n 避免负数
            set.add((long) (x + n) << 20 | (y + n)); // 测试发现 << 20 比 << 32 快很多

            // 3. 出，为下一个窗口做准备
            char out = s.charAt(left);
            x -= DIRS[out][0];
            y -= DIRS[out][1];
        }
        return set.size();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int distinctPoints(string s, int k) {
        int n = s.size();
        unordered_set<long long> st;
        int x = 0, y = 0;
        for (int i = 0; i < n; i++) {
            // 1. 入
            switch (s[i]) {
                case 'L': x--; break;
                case 'R': x++; break;
                case 'D': y--; break;
                case 'U': y++; break;
            }

            int left = i + 1 - k; // 窗口左端点
            if (left < 0) { // 窗口大小不足 k
                continue;
            }

            // 2. 记录答案
            // 把两个 int 压缩到一个 long long 中，+n 避免负数
            st.insert(1LL * (x + n) << 32 | (y + n));

            // 3. 出，为下一个窗口做准备
            switch (s[left]) {
                case 'L': x++; break;
                case 'R': x--; break;
                case 'D': y++; break;
                case 'U': y--; break;
            }
        }
        return st.size();
    }
};
```

```go [sol-Go]
type pair struct{ x, y int }
var dirs = []pair{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}

func distinctPoints(s string, k int) int {
	set := map[pair]struct{}{}
	p := pair{}
	for i, c := range s {
		// 1. 入
		p.x += dirs[c].x
		p.y += dirs[c].y

		left := i + 1 - k // 窗口左端点
		if left < 0 { // 窗口大小不足 k
			continue
		}

		// 2. 记录答案
		set[p] = struct{}{}

		// 3. 出，为下一个窗口做准备
		out := s[left]
		p.x -= dirs[out].x
		p.y -= dirs[out].y
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n-k)$。有 $n-k+1$ 个窗口，所以哈希表中至多保存 $n-k+1$ 个不同坐标。

## 优化

所有窗口都是基于第一个窗口计算的，我们只需要算相对第一个窗口的增减量。

增减量不同，最终坐标也就不同。

所以不需要算第一个窗口，只需要把后续窗口的增减量算出来。哈希集合保存的是增减量。

第一个窗口的增减量是 $(0,0)$。

```py [sol-Python3]
# 更快的写法见【Python3 写法二】
DIRS = {
    'L': (-1, 0),
    'R': (1, 0),
    'D': (0, -1),
    'U': (0, 1),
}

class Solution:
    def distinctPoints(self, s: str, k: int) -> int:
        st = {(0, 0)}  # 第一个窗口
        x = y = 0
        for i in range(k, len(s)):
            in_x, in_y = DIRS[s[i]]
            out_x, out_y = DIRS[s[i - k]]
            x += in_x - out_x
            y += in_y - out_y
            st.add((x, y))
        return len(st)
```

```py [sol-Python3 写法二]
class Solution:
    def distinctPoints(self, s: str, k: int) -> int:
        n = len(s)
        DIRS = {
            'L': -(n + 1),
            'R': n + 1,
            'D': -1,
            'U': 1,
        }
        st = {0}  # 第一个窗口
        x = 0  # 二维坐标映射成一维坐标
        for i in range(k, n):
            x += DIRS[s[i]] - DIRS[s[i - k]]
            st.add(x)
        return len(st)
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = new int[128][];

    static {
        DIRS['L'] = new int[]{-1, 0};
        DIRS['R'] = new int[]{1, 0};
        DIRS['D'] = new int[]{0, -1};
        DIRS['U'] = new int[]{0, 1};
    }

    public int distinctPoints(String s, int k) {
        int n = s.length();
        int x = 0, y = 0;
        Set<Long> set = new HashSet<>();
        set.add((long) n << 20 | n); // 第一个窗口
        for (int i = k; i < s.length(); i++) {
            char in = s.charAt(i);
            char out = s.charAt(i - k);
            x += DIRS[in][0] - DIRS[out][0];
            y += DIRS[in][1] - DIRS[out][1];
            set.add((long) (x + n) << 20 | (y + n));
        }
        return set.size();
    }
}
```

```cpp [sol-C++]
int DIRS[128][2];

int init = [] {
    DIRS['L'][0] = DIRS['D'][1] = -1;
    DIRS['R'][0] = DIRS['U'][1] = 1;
    return 0;
}();

class Solution {
public:
    int distinctPoints(string s, int k) {
        int n = s.size();
        unordered_set<long long> st;
        st.insert(1LL * n << 32 | n); // 第一个窗口
        int x = 0, y = 0;
        for (int i = k; i < n; i++) {
            char in = s[i], out = s[i - k];
            x += DIRS[in][0] - DIRS[out][0];
            y += DIRS[in][1] - DIRS[out][1];
            st.insert(1LL * (x + n) << 32 | (y + n));
        }
        return st.size();
    }
};
```

```go [sol-Go]
type pair struct{ x, y int }
var dirs = []pair{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}

func distinctPoints(s string, k int) int {
	p := pair{}
	set := map[pair]struct{}{p: {}} // 第一个窗口
	for i := k; i < len(s); i++ {
		in, out := s[i], s[i-k]
		p.x += dirs[in].x - dirs[out].x
		p.y += dirs[in].y - dirs[out].y
		set[p] = struct{}{}
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n-k)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n-k)$。有 $n-k+1$ 个窗口，所以哈希表中至多保存 $n-k+1$ 个不同坐标。

## 变形题

- [CF1296C](https://codeforces.com/problemset/problem/1296/C)
- [CF1902D](https://codeforces.com/problemset/problem/1902/D)

## 专题训练

见下面滑动窗口题单的「**一、定长滑动窗口**」。

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
