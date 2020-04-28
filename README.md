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
BenchmarkListBuilder/literal-10/baseline-4              1000000000               0.297 ns/op           0 B/op          0 allocs/op
BenchmarkListBuilder/literal-10/immutable-4              3741436               323 ns/op             544 B/op          5 allocs/op
BenchmarkListBuilder/10/baseline-4                       5373658               224 ns/op             232 B/op         10 allocs/op
BenchmarkListBuilder/10/immutable-4                      2417404               497 ns/op             616 B/op         14 allocs/op
BenchmarkListBuilder/1024/baseline-4                       69098             17457 ns/op           24568 B/op       1024 allocs/op
BenchmarkListBuilder/1024/immutable-4                      50731             23809 ns/op           57400 B/op       1028 allocs/op
BenchmarkListBuilder/131072/baseline-4                       438           2626519 ns/op         3145729 B/op     131072 allocs/op
BenchmarkListBuilder/131072/immutable-4                      169           7430739 ns/op         7340110 B/op     131076 allocs/op
BenchmarkListRange/10/baseline-4                        335512334                3.56 ns/op            0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-4                       46481835                25.9 ns/op             0 B/op          0 allocs/op
BenchmarkListRange/1024/baseline-4                       3881248               309 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-4                       490236              2437 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/baseline-4                       30973             38782 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-4                       3625            330392 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/baseline-4                 8127631               147 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/immutable-4                1000000              1079 ns/op            1024 B/op          8 allocs/op
BenchmarkMapBuilder/10/baseline-4                        1000000              1093 ns/op             726 B/op         19 allocs/op
BenchmarkMapBuilder/10/immutable-literal-4                314469              3773 ns/op            2914 B/op         29 allocs/op
BenchmarkMapBuilder/10/immutable-builder-4                486196              2531 ns/op            1996 B/op         26 allocs/op
BenchmarkMapBuilder/1024/baseline-4                         6556            182510 ns/op          182884 B/op       2076 allocs/op
BenchmarkMapBuilder/1024/immutable-literal-4                2210            543145 ns/op          516892 B/op       2145 allocs/op
BenchmarkMapBuilder/1024/immutable-builder-4                3278            367344 ns/op          350118 B/op       2112 allocs/op
BenchmarkMapBuilder/131072/baseline-4                         26          39736612 ns/op        22418431 B/op     266824 allocs/op
BenchmarkMapBuilder/131072/immutable-literal-4                 8         133574731 ns/op        63065967 B/op     276210 allocs/op
BenchmarkMapBuilder/131072/immutable-builder-4                13          87654537 ns/op        42739040 B/op     271508 allocs/op
BenchmarkMapRange/10/baseline-4                         10257226               117 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-4                         8666278               138 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/baseline-4                          88386             13542 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-4                         79114             15431 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/baseline-4                          664           1790312 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-4                         592           1981791 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/fishy/go-immutable   41.328s
```
