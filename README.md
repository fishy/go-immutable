[![PkgGoDev](https://pkg.go.dev/badge/go.yhsif.com/immutable)](https://pkg.go.dev/go.yhsif.com/immutable)
[![Go Report Card](https://goreportcard.com/badge/go.yhsif.com/immutable)](https://goreportcard.com/report/go.yhsif.com/immutable)

# go-immutable

Immutable map/set/list for go.

## Benchmark

This library comes with benchmark test to compare against builtin types
(there's no builtin set so that's only list and map).
Here is an example of the benchmark result (with baseline is builtin):

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: go.yhsif.com/immutable
cpu: Intel(R) Core(TM) i5-7260U CPU @ 2.20GHz
BenchmarkListBuilder/literal-10/baseline-4              1000000000               0.3048 ns/op          0 B/op          0 allocs/op
BenchmarkListBuilder/literal-10/immutable-4              6821798               174.2 ns/op           288 B/op          5 allocs/op
BenchmarkListBuilder/10/baseline-4                      33455275                35.15 ns/op           80 B/op          1 allocs/op
BenchmarkListBuilder/10/immutable-4                      5512944               182.3 ns/op           288 B/op          5 allocs/op
BenchmarkListBuilder/1024/baseline-4                      637911              1738 ns/op            8192 B/op          1 allocs/op
BenchmarkListBuilder/1024/immutable-4                     289621              4133 ns/op           24624 B/op          5 allocs/op
BenchmarkListBuilder/131072/baseline-4                      6976            165604 ns/op         1048580 B/op          1 allocs/op
BenchmarkListBuilder/131072/immutable-4                     2708            409006 ns/op         3145788 B/op          5 allocs/op
BenchmarkListRange/10/baseline-4                        263224618                4.456 ns/op           0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-4                       51426008                21.99 ns/op            0 B/op          0 allocs/op
BenchmarkListRange/1024/baseline-4                       3754664               324.0 ns/op             0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-4                       618620              1836 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/baseline-4                       30866             41303 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-4                       5124            246614 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/baseline-4                17114613                67.12 ns/op            0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/immutable-4                1418350               842.0 ns/op           688 B/op         12 allocs/op
BenchmarkMapBuilder/10/baseline-4                        3125396               385.5 ns/op           292 B/op          1 allocs/op
BenchmarkMapBuilder/10/immutable-literal-4                548138              2074 ns/op            1562 B/op         15 allocs/op
BenchmarkMapBuilder/10/immutable-builder-4                855835              1302 ns/op            1031 B/op         10 allocs/op
BenchmarkMapBuilder/1024/baseline-4                        16005             75099 ns/op           86580 B/op         64 allocs/op
BenchmarkMapBuilder/1024/immutable-literal-4                4414            263092 ns/op          260299 B/op        201 allocs/op
BenchmarkMapBuilder/1024/immutable-builder-4                6236            170937 ns/op          173528 B/op        134 allocs/op
BenchmarkMapBuilder/131072/baseline-4                         78          13458018 ns/op        10926740 B/op       4766 allocs/op
BenchmarkMapBuilder/131072/immutable-literal-4                24          48260776 ns/op        32782071 B/op      14315 allocs/op
BenchmarkMapBuilder/131072/immutable-builder-4                37          31927129 ns/op        21850800 B/op       9530 allocs/op
BenchmarkMapRange/10/baseline-4                          9839715               121.8 ns/op             0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-4                         8229484               139.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapRange/1024/baseline-4                          86908             13736 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-4                         77984             15236 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/baseline-4                          662           1685172 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-4                         627           1978727 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/immutable  40.332s
```

## License

[BSD License](LICENSE).
