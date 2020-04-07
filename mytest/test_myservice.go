package test

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	expect := "hello func in package hello."

	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello()
	}
}