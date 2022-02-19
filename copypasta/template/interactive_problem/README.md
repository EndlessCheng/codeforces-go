# Interactive Problem Template

To mock the IO part, implement this interface and use it both in `main.go` and `main_test.go`.

```go
type interaction interface {
	readInitData() initData
	query(request) response
	printAnswer(answer)
}
``` 
