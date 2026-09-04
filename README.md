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

The current API (`createBloomFilter`, `insert`, `contains`) is unexported —
this is still a work-in-progress learning exercise, not a published library.

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
