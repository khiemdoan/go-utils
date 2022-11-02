package cryptoutils

import (
	"testing"
)

func TestMd5(t *testing.T) {
	data := []byte("abc")
	expected := "900150983cd24fb0d6963f7d28e17f72"
	if Md5(data) != expected {
		t.Error("Md5 was incorrect")
	}
}

func TestSha1(t *testing.T) {
	data := []byte("abc")
	expected := "a9993e364706816aba3e25717850c26c9cd0d89d"
	if Sha1(data) != expected {
		t.Error("Sha1 was incorrect")
	}
}

func TestSha224(t *testing.T) {
	data := []byte("abc")
	expected := "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7"
	if Sha224(data) != expected {
		t.Error("Sha224 was incorrect")
	}
}

func TestSha256(t *testing.T) {
	data := []byte("abc")
	expected := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if Sha256(data) != expected {
		t.Error("Sha256 was incorrect")
	}
}

func TestSha384(t *testing.T) {
	data := []byte("abc")
	expected := "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7"
	if Sha384(data) != expected {
		t.Error("Sha384 was incorrect")
	}
}

func TestSha512(t *testing.T) {
	data := []byte("abc")
	expected := "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"
	if Sha512(data) != expected {
		t.Error("Sha512 was incorrect")
	}
}
