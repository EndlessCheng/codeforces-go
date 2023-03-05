思考：$4325$ 怎么分？

如果分成 $432$ 和 $5$，那么 $4$ 显然放在 $5$ 这边更优，所以要**均匀分**。

如果分成 $32$ 和 $45$，那么 $23$ 比 $32$ 更好；进一步地，分成 $24$ 和 $35$ 更好，所以**把小的排在高位更优，大的排在低位更优**。

设 $s$ 是 $\textit{num}$ 的字符串形式，这等价于把 $s$ 排序后，按照奇偶下标分组。

附：[视频讲解](https://www.bilibili.com/video/BV1dY4y1C77x/)

```py [sol1-Python3]
class Solution:
    def splitNum(self, num: int) -> int:
        s = sorted(str(num))
        return int(''.join(s[::2])) + int(''.join(s[1::2]))
```

```java [sol1-Java]
class Solution {
    public int splitNum(int num) {
        var s = Integer.toString(num).toCharArray();
        Arrays.sort(s);
        var a = new int[2];
        for (int i = 0; i < s.length; i++)
            a[i % 2] = a[i % 2] * 10 + s[i] - '0'; // 按照奇偶下标分组
        return a[0] + a[1];
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int splitNum(int num) {
        string s = to_string(num);
        sort(s.begin(), s.end());
        int a[2]{};
        for (int i = 0; i < s.length(); ++i)
            a[i % 2] = a[i % 2] * 10 + s[i] - '0'; // 按照奇偶下标分组
        return a[0] + a[1];
    }
};
```

```go [sol1-Go]
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

### 复杂度分析

- 时间复杂度：$O(m\log m)$，其中 $m$ 为 $\textit{num}$ 转成字符串后的长度。
- 空间复杂度：$O(m)$。
