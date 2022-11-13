按要求模拟。

```py [sol1-Python3]
class Solution:
    def convertTemperature(self, celsius: float) -> List[float]:
        return [celsius + 273.15, celsius * 1.8 + 32]
```

```go [sol1-Go]
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32}
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。
