package cryptor

import (
	"hash"
)

const (
	sm3BlockSize = 64
	sm3Size      = 32
	sm3T1        = 0x79cc4519
	sm3T2        = 0x7a879d8a
)

var sm3IV = [8]uint32{
	0x7380166f, 0x4914b2b9, 0x172442d7, 0xda8a0600,
	0xa96f30bc, 0x163138aa, 0xe38dee4d, 0xb0fb0e4e,
}

type sm3Digest struct {
	h   [8]uint32
	x   [sm3BlockSize]byte
	nx  int
	len uint64
}

func Sm3(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func newSm3() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (d *sm3Digest) Reset() { _ = "STUB: not implemented"; return }

func (d *sm3Digest) Size() int { _ = "STUB: not implemented"; return 0 }

func (d *sm3Digest) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (d *sm3Digest) Write(p []byte) (nn int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d *sm3Digest) Sum(in []byte) []byte { _ = "STUB: not implemented"; return nil }

func (d *sm3Digest) checkSum() [sm3Size]byte { _ = "STUB: not implemented"; return [sm3Size]byte{} }

func sm3Block(dig *sm3Digest, p []byte) { _ = "STUB: not implemented"; return }

func sm3RotateLeft(x, n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func sm3P0(x uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func sm3P1(x uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func sm3FF0(x, y, z uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func sm3FF1(x, y, z uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func sm3GG0(x, y, z uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func sm3GG1(x, y, z uint32) uint32 { _ = "STUB: not implemented"; return 0 }
