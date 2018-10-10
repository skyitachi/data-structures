package trie

import (
	"crypto/rand"
	"testing"
)
var stringKeys [1000]string // random string keys
const bytesPerKey = 30

func init() {
	// string keys
	for i := 0; i < len(stringKeys); i++ {
		key := make([]byte, bytesPerKey)
		if _, err := rand.Read(key); err != nil {
			panic("error generating random byte slice")
		}
		stringKeys[i] = string(key)
	}

}
func BenchmarkTriePutStringKey(b *testing.B) {
	trie := NewTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Add(stringKeys[i%len(stringKeys)], i)
	}
}

func BenchmarkTrieGetStringKey(b *testing.B) {
	trie := NewTrie()
	for i := 0; i < b.N; i++ {
		trie.Add(stringKeys[i%len(stringKeys)], i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(stringKeys[i%len(stringKeys)])
	}
}
