package cryptoutils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Md5(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha1(data []byte) string {
	hasher := sha1.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha224(data []byte) string {
	hasher := sha256.New224()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha256(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha384(data []byte) string {
	hasher := sha512.New384()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha512(data []byte) string {
	hasher := sha512.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}
