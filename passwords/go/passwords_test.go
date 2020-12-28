package passwords_test

import (
	"passwords"
	"testing"
)

func Test_abcde_is_valid_for_a1(t *testing.T) {
	foo := passwords.PasswordCandidate{Min: 1, Max: 1, Letter: "a", Passowrd: "abcde"}
	valid := foo.IsValid()
	if !valid {
		t.Error("expected to be valid but was not")
	}
}

func Test_bcde_is_invalid_for_a1(t *testing.T) {
	foo := passwords.PasswordCandidate{Min: 1, Max: 1, Letter: "a", Passowrd: "bcde"}
	valid := foo.IsValid()
	if valid {
		t.Error("expected to be invalid but was not")
	}
}

func Test_cccccccc_is_valid_for_c1(t *testing.T) {
	foo := passwords.PasswordCandidate{Min: 1, Max: 9, Letter: "c", Passowrd: "cccccccc"}
	valid := foo.IsValid()
	if !valid {
		t.Error("expected to be valid but was not")
	}
}

func Test_abcde_is_invalid_for_a2(t *testing.T) {
	foo := passwords.PasswordCandidate{Min: 2, Max: 2, Letter: "a", Passowrd: "abcde"}
	valid := foo.IsValid()
	if valid {
		t.Error("expected to be invalid but was not")
	}
}

func Test_aaabcde_is_invalid_for_a12(t *testing.T) {
	foo := passwords.PasswordCandidate{Min: 1, Max: 2, Letter: "a", Passowrd: "aaabcde"}
	valid := foo.IsValid()
	if valid {
		t.Error("expected to be invalid but was not")
	}
}
