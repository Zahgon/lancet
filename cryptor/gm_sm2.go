package cryptor

import (
	"crypto/elliptic"
	"io"
	"math/big"
)

var (
	sm2P256       *sm2Curve
	sm2P256Params = &elliptic.CurveParams{Name: "sm2p256v1"}
)

func init() {

	sm2P256Params.P, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	sm2P256Params.N, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	sm2P256Params.B, _ = new(big.Int).SetString("28E9FA9E9D9F5E344D5A9E4BCF6509A7F39789F515AB8F92DDBCBD414D940E93", 16)
	sm2P256Params.Gx, _ = new(big.Int).SetString("32C4AE2C1F1981195F9904466A39C9948FE30BBFF2660BE1715A4589334C74C7", 16)
	sm2P256Params.Gy, _ = new(big.Int).SetString("BC3736A2F4F6779C59BDCEE36B692153D0A9877CC62A474002DF32E52139F0A0", 16)
	sm2P256Params.BitSize = 256

	sm2P256 = &sm2Curve{sm2P256Params}
}

type sm2Curve struct {
	*elliptic.CurveParams
}

type Sm2PrivateKey struct {
	D         *big.Int
	PublicKey Sm2PublicKey
}

type Sm2PublicKey struct {
	X, Y *big.Int
}

func GenerateSm2Key() (*Sm2PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func Sm2Encrypt(pub *Sm2PublicKey, plaintext []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Sm2Decrypt(priv *Sm2PrivateKey, ciphertext []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sm2KDF(z []byte, klen int) []byte { _ = "STUB: not implemented"; return nil }

func toBytes(curve elliptic.Curve, value *big.Int) []byte { _ = "STUB: not implemented"; return nil }

func sm2MarshalUncompressed(curve *sm2Curve, x, y *big.Int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func sm2UnmarshalUncompressed(curve *sm2Curve, data []byte) (*big.Int, *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func randFieldElement(c elliptic.Curve, rand io.Reader) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
