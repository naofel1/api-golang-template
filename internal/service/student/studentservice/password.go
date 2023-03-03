package studentservice

import (
	"context"

	"github.com/matthewhartstonge/argon2"
	"go.opentelemetry.io/otel/trace"
)

func generatePasswordHash(ctx context.Context, tracer trace.Tracer, secret string) ([]byte, error) {
	_, span := tracer.Start(ctx, "Password Hash Generation")
	defer span.End()

	cfg := argon2.MemoryConstrainedDefaults()

	raw, err := cfg.Hash([]byte(secret), nil)
	if err != nil {
		return nil, err
	}
	// Encode the raw secret byte
	encoded := raw.Encode()

	return encoded, nil
}

func validatePasswordHash(ctx context.Context, tracer trace.Tracer, storedPassword, suppliedPassword []byte) (bool, error) {
	_, span := tracer.Start(ctx, "Password Hash Validation")
	defer span.End()

	// Check if the password match
	ok, err := argon2.VerifyEncoded(suppliedPassword, storedPassword)
	if err != nil {
		span.RecordError(err)
		return false, err
	}
	// Check if password is valid
	if ok {
		return true, nil
	}

	return false, nil
}
