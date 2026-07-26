## 方法一：字符串

分类讨论：

- 最多填 $n$ 个 $9$，所以如果 $s>9n$，则无解，返回 $-1$。
- 如果 $s=0$，那么答案为 $0$。
- 否则，整数的十进制长度可以是 $n$。从高到低填：
   - 如果 $s\le 9$，那么这一位填 $s$，剩余位填 $0$，跳出循环。
   - 如果 $s > 9$，那么这一位填 $9$，继续循环。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def largestInteger(self, n: int, s: int) -> int:
        if s > n * 9:
            return -1
        if s == 0:
            return 0

        res = ['0'] * n
        for i in range(n):
            if s <= 9:
                res[i] = str(s)
                break
            res[i] = '9'
            s -= 9
        return int(''.join(res))
```

```java [sol-Java]
class Solution {
    public int largestInteger(int n, int s) {
        if (s > n * 9) {
            return -1;
        }
        if (s == 0) {
            return 0;
        }

        char[] res = new char[n];
        Arrays.fill(res, '0');
        for (int i = 0; i < n; i++) {
            if (s <= 9) {
                res[i] += s;
                break;
            }
            res[i] = '9';
            s -= 9;
        }
        return Integer.parseInt(new String(res));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int largestInteger(int n, int s) {
        if (s > n * 9) {
            return -1;
        }
        if (s == 0) {
            return 0;
        }

        string res(n, '0');
        for (int i = 0; i < n; i++) {
            if (s <= 9) {
                res[i] += s;
                break;
            }
            res[i] = '9';
            s -= 9;
        }
        return stoi(res);
    }
};
```

```go [sol-Go]
func largestInteger(n, s int) int {
	if s > n*9 {
		return -1
	}
	if s == 0 {
		return 0
	}

	res := bytes.Repeat([]byte{'0'}, n)
	for i := range res {
		if s <= 9 {
			res[i] += byte(s)
			break
		}
		res[i] = '9'
		s -= 9
	}
	ans, _ := strconv.Atoi(string(res))
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：数学

先填 $L = \left\lfloor\dfrac{s}{9}\right\rfloor$ 个 $9$，即 $10^L - 1$。

然后，如果 $s\bmod 9 > 0$，再填一个 $s\bmod 9$。

最后填 $0$，直到十进制长度为 $n$。

```py [sol-Python3]
class Solution:
    def largestInteger(self, n: int, s: int) -> int:
        if s > n * 9:
            return -1

        ans = 10 ** (s // 9) - 1
        if s % 9:
            ans = ans * 10 + s % 9
            n -= 1
        return ans * 10 ** (n - s // 9)
```

```java [sol-Java]
class Solution {
    public int largestInteger(int n, int s) {
        if (s > n * 9) {
            return -1;
        }

        int ans = (int) Math.pow(10, s / 9) - 1;
        if (s % 9 > 0) {
            ans = ans * 10 + s % 9;
            n--;
        }
        return ans * (int) Math.pow(10, n - s / 9);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int largestInteger(int n, int s) {
        if (s > n * 9) {
            return -1;
        }

        int ans = (int) pow(10, s / 9) - 1;
        if (s % 9) {
            ans = ans * 10 + s % 9;
            n--;
        }
        return ans * (int) pow(10, n - s / 9);
    }
};
```

```go [sol-Go]
func largestInteger(n, s int) int {
	if s > n*9 {
		return -1
	}

	ans := int(math.Pow10(s/9)) - 1
	if s%9 > 0 {
		ans = ans*10 + s%9
		n--
	}
	return ans * int(math.Pow10(n-s/9))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$ 或 $\mathcal{O}(\log s)$，取决于实现。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心题单的「**§3.1 字典序最小/最大**」。

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
