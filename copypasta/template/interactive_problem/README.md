# Interactive Problem Template

Use below interface and implement it both in `main.go` and `main_test.go`.

```go
type interaction interface {
	readInitData() initData
	query(request) response
	printAnswer(answer)
}
``` 

In this way we can mock the IO part.
