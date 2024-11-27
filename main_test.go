package main

import (
	"fmt"
	"testing"
)

const execucoes = 20_000_000

func TestCreatePersonStack(t *testing.T) {
	for i := 0; i < execucoes; i++ {
		p := CreatePersonStack()
		if p.Age > 50 {
			fmt.Print(p)
		}
	}
}

func TestCreatePersonHeap(t *testing.T) {
	for i := 0; i < execucoes; i++ {
		p := CreatePersonHeap()
		if p.Age > 50 {
			fmt.Print(p)
		}
	}
}

func BenchmarkCreatePersonStack(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := CreatePersonStack()
		if p.Age > 50 {
			fmt.Print(p)
		}
	}
}

func BenchmarkCreatePersonHeap(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := CreatePersonHeap()
		if p.Age > 50 {
			fmt.Print(p)
		}
	}
}
