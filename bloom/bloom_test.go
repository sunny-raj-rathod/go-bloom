package bloom

import (
	"fmt"
	"testing"
)

func TestInsertAndContains(t *testing.T) {
	bf := CreateBloomFilter(0.001, 100000)
	bf.Insert("hello")
	bf.Insert("world")

	if !bf.Contains("hello") {
		t.Errorf("expected Contains(\"hello\") to be true after insert, got false")
	}

	if !bf.Contains("world") {
		t.Errorf("expected Contains(\"world\") to be true after insert, got false")
	}

	if bf.Contains("bye") {
		t.Errorf("expected Contains(\"bye\") to be false, got true")
	}
}

func TestEmpty(t *testing.T) {
	bf := CreateBloomFilter(0.001, 100000)

	if bf.Contains("hello") {
		t.Errorf("expected Contains(\"hello\") to be false, got true")
	}
}

func TestSmallFilter(t *testing.T) {
	bf := CreateBloomFilter(0.001, 1)

	bf.Insert("hello")
	bf.Insert("world")

	if !bf.Contains("hello") {
		t.Errorf("expected Contains(\"hello\") to be true after insert, got false")
	}

	if !bf.Contains("world") {
		t.Errorf("expected Contains(\"world\") to be true after insert, got false")
	}

	if bf.Contains("bye") {
		t.Errorf("expected Contains(\"bye\") to be false, got true")
	}
}

func BenchmarkInsert(b *testing.B) {
	bf := CreateBloomFilter(0.001, 100000)
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Insert(keys[i])
	}
}

func BenchmarkContains(b *testing.B) {
	bf := CreateBloomFilter(0.001, 100000)
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
		bf.Insert(keys[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Contains(keys[i])
	}
}
