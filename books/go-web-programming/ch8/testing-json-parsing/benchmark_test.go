package main

import "testing"

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}
