### Running benchmarks

```
go test -bench=.
```

### TLDR

#### String Manipulation

If inputs and outputs are strings, use `+`. Macbook air results:
```
BenchmarkAddStrings10	20000000	       121 ns/op
BenchmarkAddStringswJoin10	10000000	       189 ns/op
BenchmarkAddStringswFormat10	 5000000	       521 ns/op
BenchmarkAddStringswBytesBuffer10	10000000	       299 ns/op
```