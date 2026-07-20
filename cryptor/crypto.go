package cryptor

import (
	"crypto"
	"crypto/rsa"
)

func AesEcbEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesEcbDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesCbcEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesCbcDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesCtrCrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesCtrEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesCtrDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesCfbEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesCfbDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesOfbEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesOfbDecrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesGcmEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func AesGcmDecrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesEcbEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesEcbDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesCbcEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesCbcDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesCtrCrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesCtrEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesCtrDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesCfbEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesCfbDecrypt(encrypted, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesOfbEncrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func DesOfbDecrypt(data, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func RsaEncrypt(data []byte, pubKeyFileName string) []byte { _ = "STUB: not implemented"; return nil }

func RsaDecrypt(data []byte, privateKeyFileName string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func GenerateRsaKeyPair(keySize int) (*rsa.PrivateKey, *rsa.PublicKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RsaEncryptOAEP(data []byte, label []byte, key rsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RsaDecryptOAEP(ciphertext []byte, label []byte, key rsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RsaSign(hash crypto.Hash, data []byte, privateKeyFileName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RsaVerifySign(hash crypto.Hash, data, signature []byte, pubKeyFileName string) error {
	_ = "STUB: not implemented"
	return nil
}
