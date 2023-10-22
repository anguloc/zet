package main

import "testing"

// go test ./ -bench -v -run BenchmarkFunc1$ --count=1 -mem -n 10
func BenchmarkFunc1(b *testing.B) {
	f := &Foo{}
	data := make([]byte, 4*1024*1024)
	for i := 0; i < 4*1024*1024; i++ {
		data = append(data, 'a')
	}
	
	f.A.B.C = string(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Bar()
	}
}

func BenchmarkFunc2(b *testing.B) {
	f := &Foo{}
	data := make([]byte, 4*1024*1024)
	for i := 0; i < 4*1024*1024; i++ {
		data = append(data, 'a')
	}
	f.A.B.C = string(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Qux()
	}
}
