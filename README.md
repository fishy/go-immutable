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
cpu: 12th Gen Intel(R) Core(TM) i5-1235U
BenchmarkListBuilder/literal-10/baseline-12             1000000000               0.1226 ns/op          0 B/op          0 allocs/op
BenchmarkListBuilder/literal-10/immutable-12             8853948               138.0 ns/op           288 B/op          5 allocs/op
BenchmarkListBuilder/10/baseline-12                     35211993                30.00 ns/op           80 B/op          1 allocs/op
BenchmarkListBuilder/10/immutable-12                     7969965               141.7 ns/op           288 B/op          5 allocs/op
BenchmarkListBuilder/1024/baseline-12                     633338              1818 ns/op            8192 B/op          1 allocs/op
BenchmarkListBuilder/1024/immutable-12                    234624              5069 ns/op           24624 B/op          5 allocs/op
BenchmarkListBuilder/131072/baseline-12                     8010            145301 ns/op         1048578 B/op          1 allocs/op
BenchmarkListBuilder/131072/immutable-12                    1563            846922 ns/op         3145783 B/op          5 allocs/op
BenchmarkListRange/10/baseline-12                       861452865                1.387 ns/op           0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-12                      83046913                14.07 ns/op            0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-all-12                  14067778                77.32 ns/op           48 B/op          3 allocs/op
BenchmarkListRange/1024/baseline-12                      8330391               132.3 ns/op             0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-12                      969414              1187 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-all-12                  877918              1373 ns/op              48 B/op          3 allocs/op
BenchmarkListRange/131072/baseline-12                      79238             15261 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-12                      7893            154461 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-all-12                  6350            176498 ns/op              48 B/op          3 allocs/op
BenchmarkMapBuilder/literal-5/baseline-12               29004186                43.80 ns/op            0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/immutable-12               1931334               629.7 ns/op           688 B/op         12 allocs/op
BenchmarkMapBuilder/10/baseline-12                       4204045               282.5 ns/op           292 B/op          1 allocs/op
BenchmarkMapBuilder/10/immutable-literal-12               720141              1551 ns/op            1562 B/op         15 allocs/op
BenchmarkMapBuilder/10/immutable-builder-12              1233813               980.9 ns/op          1031 B/op         10 allocs/op
BenchmarkMapBuilder/1024/baseline-12                       19167             60340 ns/op           86568 B/op         64 allocs/op
BenchmarkMapBuilder/1024/immutable-literal-12               5065            220412 ns/op          260301 B/op        201 allocs/op
BenchmarkMapBuilder/1024/immutable-builder-12               7386            144569 ns/op          173537 B/op        134 allocs/op
BenchmarkMapBuilder/131072/baseline-12                       124          10252703 ns/op        10925729 B/op       4766 allocs/op
BenchmarkMapBuilder/131072/immutable-literal-12               32          36882045 ns/op        32773405 B/op      14277 allocs/op
BenchmarkMapBuilder/131072/immutable-builder-12               57          20891777 ns/op        21850429 B/op       9529 allocs/op
BenchmarkMapRange/10/baseline-12                        14706384                78.00 ns/op            0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-12                       13813046                89.78 ns/op            0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-all-12                    7471460               167.3 ns/op            48 B/op          3 allocs/op
BenchmarkMapRange/1024/baseline-12                        122810              8322 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-12                       134680              8906 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-all-12                   126817              9122 ns/op              48 B/op          3 allocs/op
BenchmarkMapRange/131072/baseline-12                        1198           1060454 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-12                       1054           1139414 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-all-12                   1010           1230834 ns/op              48 B/op          3 allocs/op
PASS
ok      go.yhsif.com/immutable  53.351s
```

## License

[BSD License](LICENSE).
