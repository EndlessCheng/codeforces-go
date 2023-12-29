读题题，按照要求分类讨论即可。

可以用变量名来简化逻辑。

```py [sol-Python3]
class Solution:
    def categorizeBox(self, length: int, width: int, height: int, mass: int) -> str:
        x = length >= 10000 or width >= 10000 or height >= 10000 or length * width * height >= 10 ** 9
        y = mass >= 100
        if x and y: return "Both"
        if x: return "Bulky"
        if y: return "Heavy"
        return "Neither"
```

```java [sol-Java]
class Solution {
    public String categorizeBox(int length, int width, int height, int mass) {
        boolean x = length >= 10000 || width >= 10000 || height >= 10000 || (long) length * width * height >= 1000000000;
        boolean y = mass >= 100;
        if (x && y) return "Both";
        if (x) return "Bulky";
        if (y) return "Heavy";
        return "Neither";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string categorizeBox(int length, int width, int height, int mass) {
        bool x = length >= 10000 || width >= 10000 || height >= 10000 || 1LL * length * width * height >= 1000000000;
        bool y = mass >= 100;
        if (x && y) return "Both";
        if (x) return "Bulky";
        if (y) return "Heavy";
        return "Neither";
    }
};
```

```go [sol-Go]
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

```js [sol-JavaScript]
var categorizeBox = function(length, width, height, mass) {
    const x = length >= 1e4 || width >= 1e4 || height >= 1e4 || length * width * height >= 1e9;
    const y = mass >= 100;
    if (x && y) return "Both";
    if (x) return "Bulky";
    if (y) return "Heavy";
    return "Neither";
};
```

```rust [sol-Rust]
impl Solution {
    pub fn categorize_box(length: i32, width: i32, height: i32, mass: i32) -> String {
        let x = length >= 10000 || width >= 10000 || height >= 10000 || length as i64 * width as i64 * height as i64 >= 1000000000;
        let y = mass >= 100;
        if x && y {
            return "Both".to_string();
        }
        if x {
            return "Bulky".to_string();
        }
        if y {
            return "Heavy".to_string();
        }
        "Neither".to_string()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
