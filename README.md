[![PkgGoDev](https://pkg.go.dev/badge/go.yhsif.com/immutable)](https://pkg.go.dev/go.yhsif.com/immutable)
[![Go Report Card](https://goreportcard.com/badge/go.yhsif.com/immutable)](https://goreportcard.com/report/go.yhsif.com/immutable)

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
pkg: go.yhsif.com/immutable
BenchmarkListBuilder/literal-10/baseline-4              1000000000               0.296 ns/op           0 B/op          0 allocs/op
BenchmarkListBuilder/literal-10/immutable-4              3684753               324 ns/op             544 B/op          5 allocs/op
BenchmarkListBuilder/10/baseline-4                      11090810               108 ns/op             160 B/op          1 allocs/op
BenchmarkListBuilder/10/immutable-4                      3189986               376 ns/op             544 B/op          5 allocs/op
BenchmarkListBuilder/1024/baseline-4                       78103             15251 ns/op           22528 B/op        769 allocs/op
BenchmarkListBuilder/1024/immutable-4                      55538             21656 ns/op           55360 B/op        773 allocs/op
BenchmarkListBuilder/131072/baseline-4                       436           2739739 ns/op         3143688 B/op     130817 allocs/op
BenchmarkListBuilder/131072/immutable-4                      162           7346655 ns/op         7338070 B/op     130821 allocs/op
BenchmarkListRange/10/baseline-4                        336276412                3.56 ns/op            0 B/op          0 allocs/op
BenchmarkListRange/10/immutable-4                       42435681                28.3 ns/op             0 B/op          0 allocs/op
BenchmarkListRange/1024/baseline-4                       3885120               309 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/1024/immutable-4                       460292              2604 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/baseline-4                       30994             38736 ns/op               0 B/op          0 allocs/op
BenchmarkListRange/131072/immutable-4                       3529            336382 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/baseline-4                 8661763               138 ns/op               0 B/op          0 allocs/op
BenchmarkMapBuilder/literal-5/immutable-4                1000000              1041 ns/op            1024 B/op          8 allocs/op
BenchmarkMapBuilder/10/baseline-4                        1497246               804 ns/op             582 B/op          1 allocs/op
BenchmarkMapBuilder/10/immutable-literal-4                358076              3321 ns/op            2770 B/op         11 allocs/op
BenchmarkMapBuilder/10/immutable-builder-4                556004              2175 ns/op            1852 B/op          8 allocs/op
BenchmarkMapBuilder/1024/baseline-4                         6931            170979 ns/op          178829 B/op       1566 allocs/op
BenchmarkMapBuilder/1024/immutable-literal-4                2346            513132 ns/op          512781 B/op       1634 allocs/op
BenchmarkMapBuilder/1024/immutable-builder-4                3469            351918 ns/op          345981 B/op       1602 allocs/op
BenchmarkMapBuilder/131072/baseline-4                         30          38380055 ns/op        22409913 B/op     266298 allocs/op
BenchmarkMapBuilder/131072/immutable-literal-4                 8         131526178 ns/op        63049368 B/op     275657 allocs/op
BenchmarkMapBuilder/131072/immutable-builder-4                12          85273761 ns/op        42729248 B/op     270978 allocs/op
BenchmarkMapRange/10/baseline-4                         10055437               119 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/10/immutable-4                         8391579               144 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/baseline-4                          87032             13758 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/1024/immutable-4                         74726             16041 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/baseline-4                          657           1817246 ns/op               0 B/op          0 allocs/op
BenchmarkMapRange/131072/immutable-4                         578           2067425 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/immutable  42.700s
```

## License

[BSD License](LICENSE).
