# Interactive Problem Template

You should implement this interface and use it both in `main.go` and `main_test.go`.

```go
type interaction interface {
	readInitData() initData
	query(request) response
	printAnswer(answer)
}
``` 

In this way we can mock the IO part.
