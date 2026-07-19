## 方法一：排序

如果 $x<y$，那么把 $s$ 从大到小排序，即可让所有 $y$ 都在所有 $x$ 的左边。

如果 $x>y$，那么把 $s$ 从小到大排序，即可让所有 $y$ 都在所有 $x$ 的左边。

> 注：题目保证 $x\ne y$。

[本题视频讲解](https://www.bilibili.com/video/BV1mJK66VEbN/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def rearrangeString(self, s: str, x: str, y: str) -> str:
        t = sorted(s, reverse = x < y)
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String rearrangeString(String s, char x, char y) {
        char[] t = s.toCharArray();
        Arrays.sort(t);
        if (x < y) {
            reverse(t);
        }
        return new String(t);
    }

    private void reverse(char[] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            char tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string rearrangeString(string s, char x, char y) {
        if (x < y) {
            ranges::sort(s, greater());
        } else {
            ranges::sort(s);
        }
        return s;
    }
};
```

```go [sol-Go]
func rearrangeString(s string, x, y byte) string {
	t := []byte(s)
	if x < y {
		slices.SortFunc(t, func(a, b byte) int { return int(b) - int(a) })
	} else {
		slices.Sort(t)
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 方法二：相向双指针

做法同 [905. 按奇偶排序数组](https://leetcode.cn/problems/sort-array-by-parity/)，请看 [我的题解](https://leetcode.cn/problems/sort-array-by-parity/solutions/3058524/xiang-xiang-shuang-zhi-zhen-pythonjavacc-bng5/)。

```py [sol-Python3]
class Solution:
    def rearrangeString(self, s: str, x: str, y: str) -> str:
        t = list(s)
        l, r = 0, len(t) - 1
        while l < r:  # 循环直到不足两个字母
            if t[l] != x:  # 寻找最左边的 x
                l += 1
            elif t[r] != y:  # 寻找最右边的 y
                r -= 1
            else:
                t[l], t[r] = t[r], t[l]
                # 交换后，[0,l] 都不含 x，[r,n-1] 都不含 y，问题变成 [l+1,r-1] 的子问题
                l += 1
                r -= 1
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String rearrangeString(String s, char x, char y) {
        char[] t = s.toCharArray();
        int l = 0;
        int r = t.length - 1;
        while (l < r) { // 循环直到不足两个字母
            if (t[l] != x) { // 寻找最左边的 x
                l++;
            } else if (t[r] != y) { // 寻找最右边的 y
                r--;
            } else {
                char tmp = t[l];
                t[l] = t[r];
                t[r] = tmp;
                // 交换后，[0,l] 都不含 x，[r,n-1] 都不含 y，问题变成 [l+1,r-1] 的子问题
                l++;
                r--;
            }
        }
        return new String(t);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string rearrangeString(string s, char x, char y) {
        int l = 0, r = s.size() - 1;
        while (l < r) { // 循环直到不足两个字母
            if (s[l] != x) { // 寻找最左边的 x
                l++;
            } else if (s[r] != y) { // 寻找最右边的 y
                r--;
            } else {
                swap(s[l], s[r]);
                // 交换后，[0,l] 都不含 x，[r,n-1] 都不含 y，问题变成 [l+1,r-1] 的子问题
                l++;
                r--;
            }
        }
        return s;
    }
};
```

```go [sol-Go]
func rearrangeString(s string, x, y byte) string {
	t := []byte(s)
	l, r := 0, len(t)-1
	for l < r { // 循环直到不足两个字母
		if t[l] != x { // 寻找最左边的 x
			l++
		} else if t[r] != y { // 寻找最右边的 y
			r--
		} else {
			t[l], t[r] = t[r], t[l]
			// 交换后，[0,l] 都不含 x，[r,n-1] 都不含 y，问题变成 [l+1,r-1] 的子问题
			l++
			r--
		}
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 专题训练

见下面双指针题单的「**§3.2 相向双指针**」和「**§3.5 原地修改**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
