package cryptor

import (
	"crypto"
	"crypto/rsa"
)

func generateAesKey(key []byte, size int) []byte { _ = "STUB: not implemented"; return nil }

func generateDesKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func pkcs7Padding(src []byte, blockSize int) []byte { _ = "STUB: not implemented"; return nil }

func pkcs7UnPadding(src []byte) []byte { _ = "STUB: not implemented"; return nil }

func pkcs5Padding(data []byte, blockSize int) []byte { _ = "STUB: not implemented"; return nil }

func pkcs5UnPadding(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func isAesKeyLengthValid(n int) bool { _ = "STUB: not implemented"; return false }

func loadRsaPublicKey(filename string) (*rsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadRasPrivateKey(filename string) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hashData(hash crypto.Hash, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
