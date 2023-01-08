读题题，按照要求分类讨论即可。

可以用变量名来简化逻辑。

```py [sol1-Python3]
class Solution:
    def categorizeBox(self, length: int, width: int, height: int, mass: int) -> str:
        x = length >= 10000 or width >= 10000 or height >= 10000 or length * width * height >= 10 ** 9
        y = mass >= 100
        if x and y: return "Both"
        if x: return "Bulky"
        if y: return "Heavy"
        return "Neither"
```

```go [sol1-Go]
func categorizeBox(length, width, height, mass int) string {
	x := length >= 1e4 || width >= 1e4 || height >= 1e4 || length*width*height >= 1e9
	y := mass >= 100
	switch {
	case x && y: return "Both"
	case x: return "Bulky"
	case y: return "Heavy"
	default: return "Neither"
	}
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
