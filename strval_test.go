package strval

import "testing"

var assertMock = func(string) bool { return true }

func TestValidate(t *testing.T) {
	if isValid := Validate("The 語.", assertMock); !isValid {
		t.Fatalf(`Validate("The 語.") = %t, want "%t"`, isValid, true)
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validate("The 語.", assertMock)
	}
}

func TestValidateSubstr(t *testing.T) {
	if isValid := ValidateSubstr(4, 5, "The 語.", assertMock); !isValid {
		t.Fatalf(`Validate("The 語.") = %t, want "%t"`, isValid, true)
	}
}

func BenchmarkValidateSubstr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateSubstr(4, 5, "The 語.", assertMock)
	}
}
