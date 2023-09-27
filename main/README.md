# Codeforces AC Codes

Compiler info: go1.19.5, windows, amd64

## 代码框架

编写一个 `run(io.Reader, io.Writer)` 函数来处理输入输出。这样写的理由是：

- 在 `main` 中调用 `run(os.Stdin, os.Stdout)` 来执行代码；
- 测试时，将测试数据转换成 `strings.Reader` 当作输入，并用一个 `strings.Builder` 来接收输出，将这二者传入 `run` 中，然后就能比较输出与答案了；
- 对拍时需要实现一个暴力算法 `runAC`，参数和 `run` 一样。通过[随机数据生成器](/main/testutil/rand.go)来生成数据，分别传入 `runAC` 和 `run`，通过比对各自的输出，来检查 `run` 中的问题。

例如：[1439C_test.go](./1400-1499/1439C_test.go)

交互题的写法要复杂一些，为方便测试，需要把涉及输入输出的地方抽象成接口，详见 [interactive_problem](/copypasta/template/interactive_problem)。
