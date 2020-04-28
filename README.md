# go-immutable

Immutable map/set/list for go.

## Benchmark

This library comes with benchmark test to compare against builtin types
(there's no builtin set so that's only list and map).
Here is an example of the benchmark result (with baseline is builtin):

```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/fishy/go-immutable
BenchmarkListBuilder/literal-10/baseline-4              1000000000               0.296 ns/op           0 B/op          0 allocs/op
BenchmarkListBuilder/literal-10/immutable-4              3756285               321 ns/op             544 B/op          5 allocs/op
BenchmarkListBuilder/10/baseline-4                       5379604               223 ns/op             232 B/op         10 allocs/op
BenchmarkListBuilder/10/immutable-4                      2432061               493 ns/op             616 B/op         14 allocs/op
BenchmarkListBuilder/1024/baseline-4                       68760             17429 ns/op           24568 B/op       1024 allocs/op
BenchmarkListBuilder/1024/immutable-4                      50372             23889 ns/op           57400 B/op       1028 allocs/op
BenchmarkListBuilder/131072/baseline-4                       433           2909630 ns/op         3145728 B/op     131072 allocs/op
BenchmarkListBuilder/131072/immutable-4                      159           7127343 ns/op         7340108 B/op     131076 allocs/op
BenchmarkListRange/10/baseline-4                        309889891                3.86 ns/op            0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-4                       46318929                25.8 ns/op             0 B/op          0 allocs/op
BenchmarkListRange/1024/baseline-4                       3883486               309 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-4                       492441              2435 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/baseline-4                       30990             38774 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-4                       3642            330401 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/baseline-4                 8207418               146 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/immutable-4                1000000              1067 ns/op            1024 B/op          8 allocs/op
BenchmarkMapBuilder/10/baseline-4                        1000000              1090 ns/op             726 B/op         19 allocs/op
BenchmarkMapBuilder/10/immutable-literal-4                311656              3767 ns/op            2914 B/op         29 allocs/op
BenchmarkMapBuilder/10/immutable-builder-4                485113              2534 ns/op            1996 B/op         26 allocs/op
BenchmarkMapBuilder/1024/baseline-4                         6588            182777 ns/op          182908 B/op       2076 allocs/op
BenchmarkMapBuilder/1024/immutable-literal-4                2216            543977 ns/op          516954 B/op       2145 allocs/op
BenchmarkMapBuilder/1024/immutable-builder-4                3259            366851 ns/op          350093 B/op       2112 allocs/op
BenchmarkMapBuilder/131072/baseline-4                         27          38594846 ns/op        22419431 B/op     266827 allocs/op
BenchmarkMapBuilder/131072/immutable-literal-4                 8         132091528 ns/op        63076548 B/op     276247 allocs/op
BenchmarkMapBuilder/131072/immutable-builder-4                12          86554600 ns/op        42746385 B/op     271533 allocs/op
BenchmarkMapRange/10/baseline-4                         10290792               117 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-4                         8087659               144 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/baseline-4                          88663             13530 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-4                         78784             15593 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/baseline-4                          670           1766900 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-4                         609           1983012 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/fishy/go-immutable   41.258s
```
