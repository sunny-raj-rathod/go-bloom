package bloom

import (
	"fmt"
	"testing"
)

func TestInsertAndContains(t *testing.T) {
	bf := createBloomFilter(0.001, 100000)
	bf.insert("hello")
	bf.insert("world")

	if !bf.contains("hello") {
		t.Errorf("expected contains(\"hello\") to be true after insert, got false")
	}

	if !bf.contains("world") {
		t.Errorf("expected contains(\"world\") to be true after insert, got false")
	}

	if bf.contains("bye") {
		t.Errorf("expected contains(\"bye\") to be false, got true")
	}
}

func TestEmpty(t *testing.T) {
	bf := createBloomFilter(0.001, 100000)

	if bf.contains("hello") {
		t.Errorf("expected contains(\"hello\") to be false, got true")
	}
}

func TestSmallFilter(t *testing.T) {
	bf := createBloomFilter(0.001, 1)

	bf.insert("hello")
	bf.insert("world")

	if !bf.contains("hello") {
		t.Errorf("expected contains(\"hello\") to be true after insert, got false")
	}

	if !bf.contains("world") {
		t.Errorf("expected contains(\"world\") to be true after insert, got false")
	}

	if bf.contains("bye") {
		t.Errorf("expected contains(\"bye\") to be false, got true")
	}
}

func BenchmarkInsert(b *testing.B) {
	bf := createBloomFilter(0.001, 100000)
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.insert(keys[i])
	}
}

func BenchmarkContains(b *testing.B) {
	bf := createBloomFilter(0.001, 100000)
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
		bf.insert(keys[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.contains(keys[i])
	}
}
