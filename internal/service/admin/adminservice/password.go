package adminservice

import (
	"github.com/matthewhartstonge/argon2"
)

func generatePasswordHash(secret string) ([]byte, error) {
	cfg := argon2.MemoryConstrainedDefaults()

	raw, err := cfg.Hash([]byte(secret), nil)
	if err != nil {
		return nil, err
	}

	encoded := raw.Encode()

	return encoded, nil
}

func validatePasswordHash(storedPassword, suppliedPassword []byte) (bool, error) {
	// Verify if the password hash match
	ok, err := argon2.VerifyEncoded(suppliedPassword, storedPassword)
	if err != nil {
		return false, err
	}

	// If ok the password  match
	if ok {
		return true, nil
	}

	return false, nil
}
