### 视频讲解

见[【周赛 350】](https://www.bilibili.com/video/BV1Hj411D7Tr/)第一题，欢迎点赞投币！

## 方法一：模拟

按照每次消耗 $5$ 升去模拟。如果不足 $5$ 升就退出循环，把答案加上 $10$ 乘主油箱剩余油量。

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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{mainTank})$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 方法二：快速模拟

如果 $\textit{mainTank}$ 有 $10^9$，那么方法一会超时。有没有更快的做法呢？

把方法一的减法改成除法，统计 `-=5` 发生了 $t$ 次。然后再一次性地把 $t$ 升注入主油箱。注意 $t$ 不能超过 $\textit{additionalTank}$。

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

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log\textit{mainTank})$。每次循环 $\textit{mainTank}$ 至少减为原来的 $\dfrac{1}{4}$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
