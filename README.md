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
BenchmarkListBuilder/literal-10/immutable-4              3726493               325 ns/op             544 B/op          5 allocs/op
BenchmarkListBuilder/10/baseline-4                       5319204               223 ns/op             232 B/op         10 allocs/op
BenchmarkListBuilder/10/immutable-4                      2422670               494 ns/op             616 B/op         14 allocs/op
BenchmarkListBuilder/1024/baseline-4                       68940             17464 ns/op           24568 B/op       1024 allocs/op
BenchmarkListBuilder/1024/immutable-4                      50424             23989 ns/op           57400 B/op       1028 allocs/op
BenchmarkListBuilder/131072/baseline-4                       456           2638068 ns/op         3145728 B/op     131072 allocs/op
BenchmarkListBuilder/131072/immutable-4                      168           7152446 ns/op         7340110 B/op     131076 allocs/op
BenchmarkListRange/10/baseline-4                        58527200                20.5 ns/op             0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-4                       17346458                69.2 ns/op            16 B/op          1 allocs/op
BenchmarkListRange/1024/baseline-4                        583887              2063 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-4                       280560              4282 ns/op              16 B/op          1 allocs/op
BenchmarkListRange/131072/baseline-4                        4341            276180 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-4                       2152            568392 ns/op              16 B/op          1 allocs/op
BenchmarkMapBuilder/literal-5/baseline-4                 8151202               147 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/immutable-4                1000000              1068 ns/op            1024 B/op          8 allocs/op
BenchmarkMapBuilder/10/baseline-4                        1000000              1094 ns/op             726 B/op         19 allocs/op
BenchmarkMapBuilder/10/immutable-literal-4                319438              3752 ns/op            2914 B/op         29 allocs/op
BenchmarkMapBuilder/10/immutable-builder-4                471981              2527 ns/op            1996 B/op         26 allocs/op
BenchmarkMapBuilder/1024/baseline-4                         6572            183902 ns/op          182907 B/op       2076 allocs/op
BenchmarkMapBuilder/1024/immutable-literal-4                2151            545579 ns/op          516917 B/op       2145 allocs/op
BenchmarkMapBuilder/1024/immutable-builder-4                3237            369229 ns/op          350090 B/op       2112 allocs/op
BenchmarkMapBuilder/131072/baseline-4                         26          39188081 ns/op        22420148 B/op     266830 allocs/op
BenchmarkMapBuilder/131072/immutable-literal-4                 8         131487412 ns/op        63049328 B/op     276153 allocs/op
BenchmarkMapBuilder/131072/immutable-builder-4                13          86978147 ns/op        42736448 B/op     271499 allocs/op
BenchmarkMapRange/10/baseline-4                          9443821               126 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-4                         6012328               203 ns/op              16 B/op          1 allocs/op
BenchmarkMapRange/1024/baseline-4                          80563             14860 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-4                         65101             18576 ns/op              16 B/op          1 allocs/op
BenchmarkMapRange/131072/baseline-4                          631           1890723 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-4                         507           2345218 ns/op              16 B/op          1 allocs/op
PASS
ok      github.com/fishy/go-immutable   40.533s
```
