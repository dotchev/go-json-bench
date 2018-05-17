# go-json-bench
Benchmark JSON parsing and generation in go

Here are the results I get on my Mac:
```
$ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/dotchev/go-json-bench
BenchmarkUnmarshalUnstructuredJSON-8   	  300000	      4069 ns/op
BenchmarkUnmarshalStructuredJSON-8     	  300000	      4975 ns/op
BenchmarkMarshalUnstructuredJSON-8     	  300000	      5459 ns/op
BenchmarkMarshalStructuredJSON-8       	 1000000	      1284 ns/op
PASS
ok  	github.com/dotchev/go-json-bench	5.815s
```

```
$ go version
go version go1.10 darwin/amd64
```
