package strval

type AssertFunc func(s string) bool

// Equal asserts that string s is equal to string t.
func Equal(s string) AssertFunc {
	return func(t string) bool {
		return s == t
	}
}

// NotEqual asserts that string s is not equal to string t.
func NotEqual(s string) AssertFunc {
	return func(t string) bool {
		return s != t
	}
}

// In asserts that strings ss contain string t.
func In(ss ...string) AssertFunc {
	return func(t string) bool {
		for i := 0; i < len(ss); i++ {
			if ss[i] == t {
				return true
			}
		}
		return false
	}
}

// NotIn asserts that strings ss do not contain string t.
func NotIn(ss ...string) AssertFunc {
	return func(t string) bool {
		for i := 0; i < len(ss); i++ {
			if ss[i] != t {
				return true
			}
		}
		return false
	}
}
