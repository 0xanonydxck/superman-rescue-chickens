# Problem: Superman's Chicken Rescue

## Setup & Execute
1. Download project used dependencies.
```bash
go mod download
```

2. Run `Go` command.
```bash
go run main.go
```
or (optional)
```bash
make run
```

## Test & Benchmark
1. **Time Complexity**: `O(n^2)`
2. **Memory Complexity**: `O(n)`

### Test Result
```bash
=== RUN   TestSupermanRescueChicken
=== RUN   TestSupermanRescueChicken/SupermanRescueChicken(5,_5,_[]uint{2,_5,_10,_12,_15})
=== RUN   TestSupermanRescueChicken/SupermanRescueChicken(6,_10,_[]uint{1,_11,_30,_34,_35,_37})
=== RUN   TestSupermanRescueChicken/SupermanRescueChicken(5,_5,_[]uint{2,_5,_10,_12,_15,_16,_17,_18,_18,_20})
--- PASS: TestSupermanRescueChicken (0.00s)
    --- PASS: TestSupermanRescueChicken/SupermanRescueChicken(5,_5,_[]uint{2,_5,_10,_12,_15}) (0.00s)
    --- PASS: TestSupermanRescueChicken/SupermanRescueChicken(6,_10,_[]uint{1,_11,_30,_34,_35,_37}) (0.00s)
    --- PASS: TestSupermanRescueChicken/SupermanRescueChicken(5,_5,_[]uint{2,_5,_10,_12,_15,_16,_17,_18,_18,_20}) (0.00s)
PASS
coverage: 0.0% of statements
ok      command-line-arguments  0.525s  coverage: 0.0% of statements
```

### Benchmark Result
```bash
goos: darwin
goarch: arm64
pkg: superman/module
BenchmarkSupermanRescueChicken-14         252180              4881 ns/op
BenchmarkSupermanRescueChicken-14         256515              4751 ns/op
BenchmarkSupermanRescueChicken-14         255108              4723 ns/op
BenchmarkSupermanRescueChicken-14         255540              4737 ns/op
BenchmarkSupermanRescueChicken-14         253846              4732 ns/op
PASS
ok      superman/module 7.658s
```