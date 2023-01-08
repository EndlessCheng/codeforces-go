用一个计数器 $\textit{cnt}$ 记录 $\textit{value}$ 的出现次数。

遇到就加一，没遇到就重置为 $0$。

```py [sol1-Python3]
class DataStream:
    def __init__(self, value: int, k: int):
        self.value = value
        self.k = k
        self.cnt = 0

    def consec(self, num: int) -> bool:
        self.cnt = 0 if num != self.value else self.cnt + 1
        return self.cnt >= self.k
```

```go [sol1-Go]
type DataStream struct{ value, k, cnt int }

func Constructor(value, k int) DataStream {
	return DataStream{value, k, 0}
}

func (d *DataStream) Consec(num int) bool {
	if num == d.value {
		d.cnt++
	} else {
		d.cnt = 0
	}
	return d.cnt >= d.k
}
```

#### 复杂度分析

- 时间复杂度：所有函数均为 $O(1)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
