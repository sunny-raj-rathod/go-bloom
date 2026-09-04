package bloom

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	m    uint64
	k    uint64
	n    uint64
	data []byte
}

func createBloomFilter(fp_rate float64, total_entries uint64) BloomFilter {
	m := uint64(-2.081 * math.Log(fp_rate) * float64(total_entries))
	if m == 0 {
		m = 1
	}
	k := uint64(0.693 * (float64(m) / float64(total_entries)))
	n := (m + 7) / 8
	if n == 0 {
		n = 1
	}
	return BloomFilter{m: m, k: k, n: n, data: make([]byte, n)}
}

func (bf BloomFilter) getPositions(key string) []uint64 {
	h := murmur3.Sum64([]byte(key))
	h1 := h & 0xFFFFFFFF
	h2 := (h >> 32) | 1

	positions := make([]uint64, 0, bf.k)
	for i := uint64(0); i < bf.k; i++ {
		pos := (h1 + i*h2) % bf.m
		positions = append(positions, pos)
	}

	return positions
}

func (bf BloomFilter) contains(key string) bool {
	for _, bit := range bf.getPositions(key) {
		block := bit / 8
		if block >= uint64(len(bf.data)) {
			return false
		}
		if (bf.data[block]>>(bit%8))&1 == 0 {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) insert(key string) {
	for _, bit := range bf.getPositions(key) {
		block := bit / 8
		if block >= uint64(len(bf.data)) {
			continue
		}
		bf.data[block] = bf.data[block] | byte(1<<uint(bit%8))
	}
}

/*
func main() {
	// create bloom filter with .1% fp rate, total 100k entries
	bf := createBloomFilter(0.001, 100000)
	bf.insert("hello")
	bf.insert("world")
	fmt.Printf("check if hello exists : %t\n", bf.contains("hello"))
	fmt.Printf("check if bye exists  : %t\n", bf.contains("bye"))
	fmt.Printf("check if world exists  : %t\n", bf.contains("world"))
}
*/
