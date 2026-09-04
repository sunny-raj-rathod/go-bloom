# go-bloom

A simple Bloom filter implementation in Go, built as a learning project.

## What's here

The `bloom` package (`bloom/bloom.go`) implements a standard Bloom filter:

- Sized from a target false-positive rate and expected entry count, using the
  standard optimal-`m`/optimal-`k` formulas.
- Uses [murmur3](https://github.com/spaolacci/murmur3) 64-bit hashing, split
  into two 32-bit halves and combined via double hashing (Kirsch-Mitzenmacher)
  to derive `k` bit positions per key from a single hash computation.
- Backed by a `[]byte` bit array.

## Usage

```go
package main

import (
	"fmt"

	"github.com/sunny-raj-rathod/go-bloom/bloom"
)

func main() {
	bf := bloom.CreateBloomFilter(0.001, 100000) // 0.1% target false-positive rate, 100k entries
	bf.Insert("hello")

	fmt.Println(bf.Contains("hello")) // true
	fmt.Println(bf.Contains("bye"))   // false (with occasional false positives, by design)
}
```

## Running tests

```
go test ./bloom
```

## Running benchmarks

```
go test -bench=. ./bloom
```

`BenchmarkInsert` and `BenchmarkContains` measure per-operation cost across a
set of distinct generated keys.

### Results

Measured with `fp_rate=0.001`, `total_entries=100000` (Apple M3):

| Benchmark | ns/op | B/op | allocs/op |
|---|---|---|---|
| Insert | 85.91 | 95 | 2 |
| Contains | 91.52 | 95 | 2 |

Numbers will vary by machine — re-run `go test -bench=. -benchmem ./bloom` to
reproduce on your own hardware.
