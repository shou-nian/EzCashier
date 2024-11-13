package utils

import "testing"

func TestHashPassword(t *testing.T) {
	pwds := []string{
		"123456aA.",
		".,./,3254A",
	}

	for _, pwd := range pwds {
		hashed, err := HashPassword(pwd)
		if err != nil {
			t.Error(err)
		}
		t.Log(hashed)
	}
}
