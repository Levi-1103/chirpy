package auth

import (
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "test_pass123"
	hashedPass, err := HashPassword(password)
	if err != nil {
		t.Fatalf("couldnt hash pass: %v", err)
	}

	err = CheckPasswordHash(password, hashedPass)
	if err != nil {
		t.Fatalf("hash doesnt match: %v", err)
	}
}
