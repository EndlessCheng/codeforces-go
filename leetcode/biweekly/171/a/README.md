如何判断一个数 $n$ 是不是质数？

如果暴力枚举 $[2,n-1]$ 中是否存在能整除 $n$ 的数，会超时。

我们可以判断 $n$ 能否能被 $\sqrt n$ 以内的某个大于 $1$ 的整数整除，如果不能则说明 $n$ 是质数。为什么？

**反证法**：如果 $n$ 不能被 $\sqrt n$ 以内的（大于 $1$ 的）整数整除，但可以被大于 $\sqrt n$ 的整数 $d$ 整除，那么必然还有一个数 $\dfrac{n}{d}$ 也能整除 $n$。但是 $\dfrac{n}{d} < \dfrac{n}{\sqrt n} = \sqrt n$，说明存在一个 $\sqrt n$ 以内的（大于 $1$ 的）整数能整除 $n$，矛盾。

注意 $1$ 不是质数。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
def is_prime(n: int) -> bool:
    for i in range(2, isqrt(n) + 1):
        if n % i == 0:
            return False
    return n >= 2  # 1 不是质数

class Solution:
    def completePrime(self, num: int) -> bool:
        s = str(num)
        for i in range(len(s)):
            # 前缀
            x = int(s[:i + 1])
            if not is_prime(x):
                return False

            # 后缀
            x = int(s[i:])
            if not is_prime(x):
                return False

        return True
```

```java [sol-Java]
class Solution {
    public boolean completePrime(int num) {
        String s = Integer.toString(num);
        for (int i = 0; i < s.length(); i++) {
            // 前缀
            int x = Integer.parseInt(s.substring(0, i + 1));
            if (!isPrime(x)) {
                return false;
            }

            // 后缀
            x = Integer.parseInt(s.substring(i));
            if (!isPrime(x)) {
                return false;
            }
        }
        return true;
    }

    private boolean isPrime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2; // 1 不是质数
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_prime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2; // 1 不是质数
    }

public:
    bool completePrime(int num) {
        string s = to_string(num);
        for (int i = 0; i < s.size(); i++) {
            // 前缀
            int x = stoi(s.substr(0, i + 1));
            if (!is_prime(x)) {
                return false;
            }

            // 后缀
            x = stoi(s.substr(i));
            if (!is_prime(x)) {
                return false;
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2 // 1 不是质数
}

func completePrime(num int) bool {
	s := strconv.Itoa(num)
	for i := range len(s) {
		// 前缀
		x, _ := strconv.Atoi(s[:i+1])
		if !isPrime(x) {
			return false
		}

		// 后缀
		x, _ = strconv.Atoi(s[i:])
		if !isPrime(x) {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\sqrt \textit{num})$。虽然执行了多次 $\texttt{isPrime}$，但只有最大数的 $\texttt{isPrime}$ 的循环次数占主导。
- 空间复杂度：$\mathcal{O}(\log \textit{num})$。注：如果不用字符串，可以做到 $\mathcal{O}(1)$ 空间。

## 专题训练

见下面数学题单的「**§1.1 判断质数**」。

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
