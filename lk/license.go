package lk

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// License represents a license with some data and a hash.
type License struct {
	Data []byte
	R    *big.Int
	S    *big.Int
}

// NewLicense create a new license. Make sure to sign it before use.
func NewLicense(data []byte) *License {
	return &License{
		Data: data,
	}
}

func (l *License) hash() ([]byte, error) {
	h256 := sha256.New()

	if _, err := h256.Write(l.Data); err != nil {
		return nil, err
	}
	return h256.Sum(nil), nil
}

// Sign signs a License with a PrivateKey.
func (l *License) Sign(k *PrivateKey) error {
	h, err := l.hash()
	if err != nil {
		return err
	}

	r, s, err := ecdsa.Sign(rand.Reader, k.toEcdsa(), h)
	if err != nil {
		return err
	}

	l.R = r
	l.S = s
	return nil
}

// Verify the Lisence with the public key
func (l *License) Verify(k *PublicKey) (bool, error) {
	h, err := l.hash()
	if err != nil {
		return false, err
	}

	return ecdsa.Verify(k.toEcdsa(), h, l.R, l.S), nil
}

// ToBytes transforms the public key to a base64 []byte.
func (l *License) ToBytes() ([]byte, error) {
	return toBytes(l)
}

// ToB64String transforms the public key to a base64 []byte.
func (l *License) ToB64String() (string, error) {
	return toB64String(l)
}

// LicenseFromBytes returns a License from a []byte.
func LicenseFromBytes(b []byte) (*License, error) {
	l := &License{}
	return l, fromBytes(l, b)
}

// LicenseFromB64String returns a License from a base64 encoded
// string.
func LicenseFromB64String(str string) (*License, error) {
	l := &License{}
	return l, fromB64String(l, str)
}