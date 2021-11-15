package strval

import "testing"

func TestEqual(t *testing.T) {
	if ok := Equal("This is a test.")("This is a test."); !ok {
		t.Fatalf(`Equal("This is a test.")("This is a test.") = %t, want "%t"`, ok, true)
	}
}

func BenchmarkEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal("This is a test.")("This is a test.")
	}
}

func TestNotEqual(t *testing.T) {
	if ok := NotEqual("This is a test.")("This is not a test."); !ok {
		t.Fatalf(`Equal("This is a test.")("This is not a test.") = %t, want "%t"`, ok, true)
	}
}

func BenchmarkNotEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqual("This is a test.")("This is not a test.")
	}
}

func TestIn(t *testing.T) {
	if ok := In("This", "That")("That"); !ok {
		t.Fatalf(`In("This", "That")("That") = %t, want "%t"`, ok, true)
	}
}

func BenchmarkIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		In("This", "That")("That")
	}
}

func TestNotIn(t *testing.T) {
	if ok := NotIn("This", "That")("These"); !ok {
		t.Fatalf(`NotIn("This", "That")("These") = %t, want "%t"`, ok, true)
	}
}

func BenchmarkNotIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotIn("This", "That")("That")
	}
}
