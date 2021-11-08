package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T) {
	print1()
}

func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 * time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(1 * time.Millisecond)
}

func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}

func BenchmarkPrint3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print3()
	}
}

func BenchmarkGoPrint3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint3()
	}
}
