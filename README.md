# 存储协程上下文的 RequestID

本工具可以在 go 协程中维持一个不变的 requestID，以便能在记录日志时获取到这个 ID，将整个请求的的日志串起来，方便追踪问题。

虽然使用 context 可以达到同样的目的，但是这会要求所有的方法增加 context 参数，如果只用来传递这个 ID，成本过高，在实际工程中很痛苦。

本工具并未提供一个具体的 requestID 生成方法，可以使用类似 uuid 的算法，或类似 snowflake 的算法。

## Usage
```go
import (
	"github.com/ncfwx/x/requestid"
)
```

一般可以在 http 的 middleware 中Set，在 logger 中 Get。如果请求中开了新的 goroutine，还可以继续 Set。

以下是在一个 goroutine 中使用示例
```go
go func() {
	requestid.Set("my-request-id")
	defer requestid.Delete()

	func() {
		requestid.Get()
	}()
}()
```

## Benchmark

因为 go test 的 benchmark 实际是在一个 goroutine 中运行，并没有并发，所以实际性能可能会有点差别。
```
goos: darwin
goarch: amd64
BenchmarkSet-4           20000000            92.6 ns/op
BenchmarkGet-4           20000000            75.8 ns/op
BenchmarkDelete-4        20000000            75.6 ns/op
BenchmarkGetGoID-4       2000000000          1.67 ns/op
```
