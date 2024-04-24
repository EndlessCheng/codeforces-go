## 方法一：模拟

按照每次消耗 $5$ 升燃料去模拟，直到 $\textit{mainTank}< 5$，退出循环。

别忘了最后还可以跑 $\textit{mainTank}\cdot 10$ km。

```py [sol-Python3]
class Solution:
    def distanceTraveled(self, mainTank: int, additionalTank: int) -> int:
        ans = 0
        while mainTank >= 5:
            mainTank -= 5
            ans += 50
            if additionalTank:
                additionalTank -= 1
                mainTank += 1
        return ans + mainTank * 10
```

```java [sol-Java]
class Solution {
    public int distanceTraveled(int mainTank, int additionalTank) {
        int ans = 0;
        while (mainTank >= 5) {
            mainTank -= 5;
            ans += 50;
            if (additionalTank > 0) {
                additionalTank--;
                mainTank++;
            }
        }
        return ans + mainTank * 10;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int distanceTraveled(int mainTank, int additionalTank) {
        int ans = 0;
        while (mainTank >= 5) {
            mainTank -= 5;
            ans += 50;
            if (additionalTank) {
                additionalTank--;
                mainTank++;
            }
        }
        return ans + mainTank * 10;
    }
};
```

```go [sol-Go]
func distanceTraveled(mainTank, additionalTank int) (ans int) {
	for mainTank >= 5 {
		mainTank -= 5
		ans += 50
		if additionalTank > 0 {
			additionalTank--
			mainTank++
		}
	}
	return ans + mainTank*10
}
```

```js [sol-JavaScript]
var distanceTraveled = function(mainTank, additionalTank) {
    let ans = 0;
    while (mainTank >= 5) {
        mainTank -= 5;
        ans += 50;
        if (additionalTank) {
            additionalTank--;
            mainTank++;
        }
    }
    return ans + mainTank * 10;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn distance_traveled(mut main_tank: i32, mut additional_tank: i32) -> i32 {
        let mut ans = 0;
        while main_tank >= 5 {
            main_tank -= 5;
            ans += 50;
            if additional_tank > 0 {
                additional_tank -= 1;
                main_tank += 1;
            }
        }
        ans + main_tank * 10
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{mainTank})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：快速模拟

如果 $\textit{mainTank}$ 有 $10^9$，那么方法一会超时。有没有更快的做法呢？

把方法一的减法改成除法，统计 `-= 5` 发生了 $t = \left\lfloor\dfrac{mainTank}{5}\right\rfloor$ 次。然后再一次性地把 $t$ 升燃料注入主油箱。注意 $t$ 不能超过 $\textit{additionalTank}$。

附：[视频讲解](https://www.bilibili.com/video/BV1Hj411D7Tr/)

```py [sol-Python3]
class Solution:
    def distanceTraveled(self, mainTank: int, additionalTank: int) -> int:
        ans = 0
        while mainTank >= 5:
            t = mainTank // 5
            ans += t * 50
            mainTank %= 5
            t = min(t, additionalTank)
            additionalTank -= t
            mainTank += t
        return ans + mainTank * 10
```

```java [sol-Java]
class Solution {
    public int distanceTraveled(int mainTank, int additionalTank) {
        int ans = 0;
        while (mainTank >= 5) {
            int t = mainTank / 5;
            ans += t * 50;
            mainTank %= 5;
            t = Math.min(t, additionalTank);
            additionalTank -= t;
            mainTank += t;
        }
        return ans + mainTank * 10;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int distanceTraveled(int mainTank, int additionalTank) {
        int ans = 0;
        while (mainTank >= 5) {
            int t = mainTank / 5;
            ans += t * 50;
            mainTank %= 5;
            t = min(t, additionalTank);
            additionalTank -= t;
            mainTank += t;
        }
        return ans + mainTank * 10;
    }
};
```

```go [sol-Go]
func distanceTraveled(mainTank, additionalTank int) (ans int) {
	for mainTank >= 5 {
		t := mainTank / 5
		ans += t * 50
		mainTank %= 5
		t = min(t, additionalTank)
		additionalTank -= t
		mainTank += t
	}
	return ans + mainTank*10
}
```

```js [sol-JavaScript]
var distanceTraveled = function(mainTank, additionalTank) {
    let ans = 0;
    while (mainTank >= 5) {
        let t = Math.floor(mainTank / 5);
        ans += t * 50;
        mainTank %= 5;
        t = Math.min(t, additionalTank);
        additionalTank -= t;
        mainTank += t;
    }
    return ans + mainTank * 10;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn distance_traveled(mut main_tank: i32, mut additional_tank: i32) -> i32 {
        let mut ans = 0;
        while main_tank >= 5 {
            let t = main_tank / 5;
            ans += t * 50;
            main_tank %= 5;
            let t = t.min(additional_tank);
            additional_tank -= t;
            main_tank += t;
        }
        ans + main_tank * 10
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log\textit{mainTank})$。每次循环 $\textit{mainTank}$ 至少减为原来的 $\dfrac{1}{4}$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法三：数学

考虑从副油箱能得到多少燃料。

主油箱消耗 $5$ 升燃料可以从副油箱得到 $1$ 升，这 $1$ 升又加回主油箱，相当于从主油箱的**初始燃料** $\textit{mainTank}$ 中每消耗 $4$ 升，就可以从副油箱中得到一升。所以看上去可以从副油箱得到 

$$
\min\left(\left\lfloor\dfrac{\textit{mainTank}}{4}\right\rfloor, \textit{additionalTank}\right)
$$

升燃料。但这是不对的，例如 $\textit{mainTank}=8$，只能从副油箱得到 $1$ 升燃料。因为 $8-5+1=4$，此时无法再从副油箱中得到燃料。

将上式分子减一就可以修复这个问题，也就是

$$
\min\left(\left\lfloor\dfrac{\textit{mainTank}-1}{4}\right\rfloor, \textit{additionalTank}\right)
$$

对于 $\textit{mainTank}=8$ 只能算出 $1$，而对于 $\textit{mainTank}=9$，则可以恰好从副油箱得到 $2$ 升燃料（如果有的话）。

所以答案为

$$
\left(\textit{mainTank} + \min\left(\left\lfloor\dfrac{\textit{mainTank}-1}{4}\right\rfloor, \textit{additionalTank}\right)\right)\cdot 10
$$

```py [sol-Python3]
class Solution:
    def distanceTraveled(self, mainTank: int, additionalTank: int) -> int:
        return (mainTank + min((mainTank - 1) // 4, additionalTank)) * 10
```

```java [sol-Java]
class Solution {
    public int distanceTraveled(int mainTank, int additionalTank) {
        return (mainTank + Math.min((mainTank - 1) / 4, additionalTank)) * 10;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int distanceTraveled(int mainTank, int additionalTank) {
        return (mainTank + min((mainTank - 1) / 4, additionalTank)) * 10;
    }
};
```

```go [sol-Go]
func distanceTraveled(mainTank, additionalTank int) int {
    return (mainTank + min((mainTank-1)/4, additionalTank)) * 10
}
```

```js [sol-JavaScript]
var distanceTraveled = function(mainTank, additionalTank) {
    return (mainTank + Math.min(Math.floor((mainTank - 1) / 4), additionalTank)) * 10;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn distance_traveled(main_tank: i32, additional_tank: i32) -> i32 {
        return (main_tank + additional_tank.min((main_tank - 1) / 4)) * 10;
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [1518. 换水问题](https://leetcode.cn/problems/water-bottles/)
- [3100. 换水问题 II](https://leetcode.cn/problems/water-bottles-ii/)

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
