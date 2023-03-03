package configs

import (
	"context"
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// JWTCertificate will hold certificate for JWT
type JWTCertificate struct {
	Student *CertPair
	Admin   *CertPair
}

// NewRSACerts get the content of the certificate file
func NewRSACerts(ctx context.Context, logger *otelzap.Logger, conf *Certs) *JWTCertificate {
	return &JWTCertificate{
		Student: newRSACert(ctx, logger, conf.PrivStudent, conf.PubStudent),
		Admin:   newRSACert(ctx, logger, conf.PrivAdmin, conf.PubAdmin),
	}
}

// CertPair will hold a pair of public and private key
type CertPair struct {
	Pub  *rsa.PublicKey
	Priv *rsa.PrivateKey
}

// newRSAStudent get the content of the certificate file
func newRSACert(ctx context.Context, logger *otelzap.Logger, priv, pub string) *CertPair {
	// Read the private key content
	prvKeyContent, err := os.ReadFile(priv)
	if err != nil {
		logger.Ctx(ctx).Fatal("error reading Private Key",
			zap.String("PrivateKey Directory: ", priv),
			zap.Error(err),
		)
	}

	// Read the public key content
	pubKeyContent, err := os.ReadFile(pub)
	if err != nil {
		logger.Ctx(ctx).Fatal("error reading Public Key",
			zap.String("PublicKey Directory: ", pub),
			zap.Error(err),
		)
	}

	// Parse private key content
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(prvKeyContent)
	if err != nil {
		logger.Ctx(ctx).Fatal("Error when parsing RSA Private key from PEM",
			zap.Error(err),
		)
	}

	// Parse public key content
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyContent)
	if err != nil {
		logger.Ctx(ctx).Fatal("Error when parsing RSA Public key from PEM",
			zap.Error(err),
		)
	}

	return &CertPair{
		Pub:  pubKey,
		Priv: privKey,
	}
}
