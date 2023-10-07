思考：$4325$ 怎么分？

如果分成 $432$ 和 $5$ 这两组，不如分成 $32$ 和 $45$，因为 $4$ 在 $432$ 中是 $400$，在 $45$ 中是 $40$。这启发我们要尽量**均匀分**。如果有奇数个数，多出的那个数放在哪一组都可以。

如果均匀分成 $32$ 和 $45$ 这两组，那么 $32$ 这组调整一下顺序，$23$ 比 $32$ 更好。

进一步地，均匀分成 $24$ 和 $35$ 更好，也就是**把小的数字排在高位，大的数字排在低位**。

设 $s$ 是 $\textit{num}$ 的字符串形式，这等价于把 $s$ 排序后，按照奇偶下标分组。

附：[视频讲解](https://www.bilibili.com/video/BV1dY4y1C77x/)

```py [sol-Python3]
class Solution:
    def splitNum(self, num: int) -> int:
        s = sorted(str(num))
        return int(''.join(s[::2])) + int(''.join(s[1::2]))
```

```java [sol-Java]
class Solution {
    public int splitNum(int num) {
        char[] s = Integer.toString(num).toCharArray();
        Arrays.sort(s);
        int[] a = new int[2];
        for (int i = 0; i < s.length; i++)
            a[i % 2] = a[i % 2] * 10 + s[i] - '0'; // 按照奇偶下标分组
        return a[0] + a[1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int splitNum(int num) {
        string s = to_string(num);
        sort(s.begin(), s.end());
        int a[2]{};
        for (int i = 0; i < s.length(); i++)
            a[i % 2] = a[i % 2] * 10 + s[i] - '0'; // 按照奇偶下标分组
        return a[0] + a[1];
    }
};
```

```go [sol-Go]
func splitNum(num int) int {
	s := []byte(strconv.Itoa(num))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	a := [2]int{}
	for i, c := range s {
		a[i%2] = a[i%2]*10 + int(c-'0') // 按照奇偶下标分组
	}
	return a[0] + a[1]
}
```

```js [sol-JavaScript]
var splitNum = function(num) {
    const s = num.toString().split('').sort();
    const a = [0, 0];
    for (let i = 0; i < s.length; i++) {
        a[i % 2] = a[i % 2] * 10 + parseInt(s[i]);
    }
    return a[0] + a[1];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn split_num(num: i32) -> i32 {
        let mut s: Vec<u8> = num.to_string().bytes().collect();
        s.sort_unstable();
        let mut a = [0, 0];
        for (i, &c) in s.iter().enumerate() {
            a[i % 2] = a[i % 2] * 10 + c as i32 - '0' as i32;
        }
        return a[0] + a[1];
    }
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 为 $\textit{num}$ 转成字符串后的长度。
- 空间复杂度：$\mathcal{O}(m)$。
