[![PkgGoDev](https://pkg.go.dev/badge/go.yhsif.com/immutable)](https://pkg.go.dev/go.yhsif.com/immutable)
[![Go Report Card](https://goreportcard.com/badge/go.yhsif.com/immutable)](https://goreportcard.com/report/go.yhsif.com/immutable)

# go-immutable

Immutable map/set/list for go.

## Benchmark

This library comes with benchmark test to compare against builtin types
(there's no builtin set so that's only list and map).
Here is an example of the benchmark result (with baseline is builtin):

```
$ go1.23rc1 test -bench .
goos: linux
goarch: amd64
pkg: go.yhsif.com/immutable
cpu: 12th Gen Intel(R) Core(TM) i5-1235U
BenchmarkListRangeOverFunc/10/baseline-12               845015775                1.426 ns/op           0 B/op          0 allocs/op
BenchmarkListRangeOverFunc/10/immutable-12              16573042                85.16 ns/op           48 B/op          3 allocs/op
BenchmarkListRangeOverFunc/1024/baseline-12              8870559               132.8 ns/op             0 B/op          0 allocs/op
BenchmarkListRangeOverFunc/1024/immutable-12              927172              1273 ns/op              48 B/op          3 allocs/op
BenchmarkListRangeOverFunc/131072/baseline-12              78766             15299 ns/op               0 B/op          0 allocs/op
BenchmarkListRangeOverFunc/131072/immutable-12              7062            170198 ns/op              48 B/op          3 allocs/op
BenchmarkListBuilder/literal-10/baseline-12             1000000000               0.1236 ns/op          0 B/op          0 allocs/op
BenchmarkListBuilder/literal-10/immutable-12             5638436               204.2 ns/op           288 B/op          5 allocs/op
BenchmarkListBuilder/10/baseline-12                     23636259                47.69 ns/op           80 B/op          1 allocs/op
BenchmarkListBuilder/10/immutable-12                     5019793               217.0 ns/op           288 B/op          5 allocs/op
BenchmarkListBuilder/1024/baseline-12                     610101              1972 ns/op            8192 B/op          1 allocs/op
BenchmarkListBuilder/1024/immutable-12                    223734              5064 ns/op           24624 B/op          5 allocs/op
BenchmarkListBuilder/131072/baseline-12                     5720            285895 ns/op         1048579 B/op          1 allocs/op
BenchmarkListBuilder/131072/immutable-12                    1530           1127645 ns/op         3145786 B/op          5 allocs/op
BenchmarkListRange/10/baseline-12                       816052042                1.428 ns/op           0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-12                      79910403                14.31 ns/op            0 B/op          0 allocs/op
BenchmarkListRange/1024/baseline-12                      8086268               140.5 ns/op             0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-12                      992265              1232 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/baseline-12                      74998             15303 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-12                      7922            157158 ns/op               0 B/op          0 allocs/op
BenchmarkMapRangeOverFunc/10/baseline-12                14847165                75.10 ns/op            0 B/op          0 allocs/op
BenchmarkMapRangeOverFunc/10/immutable-12                7036680               175.1 ns/op            48 B/op          3 allocs/op
BenchmarkMapRangeOverFunc/1024/baseline-12                154640              7538 ns/op               0 B/op          0 allocs/op
BenchmarkMapRangeOverFunc/1024/immutable-12               138348              9006 ns/op              48 B/op          3 allocs/op
BenchmarkMapRangeOverFunc/131072/baseline-12                1153            995477 ns/op               0 B/op          0 allocs/op
BenchmarkMapRangeOverFunc/131072/immutable-12               1082           1144126 ns/op              48 B/op          3 allocs/op
BenchmarkMapBuilder/literal-5/baseline-12               25528448                43.71 ns/op            0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/immutable-12               1249892               882.7 ns/op           688 B/op         12 allocs/op
BenchmarkMapBuilder/10/baseline-12                       2937368               369.7 ns/op           292 B/op          1 allocs/op
BenchmarkMapBuilder/10/immutable-literal-12               449425              2263 ns/op            1562 B/op         15 allocs/op
BenchmarkMapBuilder/10/immutable-builder-12               883972              1324 ns/op            1031 B/op         10 allocs/op
BenchmarkMapBuilder/1024/baseline-12                       13153             87346 ns/op           86562 B/op         64 allocs/op
BenchmarkMapBuilder/1024/immutable-literal-12               3673            301829 ns/op          260291 B/op        201 allocs/op
BenchmarkMapBuilder/1024/immutable-builder-12               5778            209962 ns/op          173520 B/op        134 allocs/op
BenchmarkMapBuilder/131072/baseline-12                       116          11546224 ns/op        10926761 B/op       4773 allocs/op
BenchmarkMapBuilder/131072/immutable-literal-12               32          39142880 ns/op        32773630 B/op      14279 allocs/op
BenchmarkMapBuilder/131072/immutable-builder-12               52          21827526 ns/op        21850592 B/op       9530 allocs/op
BenchmarkMapRange/10/baseline-12                        16441945                74.17 ns/op            0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-12                       13685006                92.67 ns/op            0 B/op          0 allocs/op
BenchmarkMapRange/1024/baseline-12                        144733              7872 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-12                       139111              9237 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/baseline-12                        1167           1008497 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-12                       1105           1183653 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/immutable  62.894s
```

## License

[BSD License](LICENSE).
