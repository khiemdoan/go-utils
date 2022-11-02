package cryptoutils

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestBase32Encode(t *testing.T) {
	data := []byte("abc")
	expected := "MFRGG==="
	if Base32Encode(data) != expected {
		t.Error("Base32Encode was incorrect")
	}
}

func TestBase32Decode(t *testing.T) {
	encrypted := "MFRGG==="
	expected := []byte("abc")

	if data, _ := Base32Decode(encrypted); data == nil || slices.Compare(data, expected) != 0 {
		t.Error("Base32Decode was incorrect")
	}
}

func TestBase64Encode(t *testing.T) {
	data := []byte("abc")
	expected := "YWJj"
	if Base64Encode(data) != expected {
		t.Error("Base64Encode was incorrect")
	}
}

func TestBase64Decode(t *testing.T) {
	encrypted := "YWJj"
	expected := []byte("abc")

	if data, _ := Base64Decode(encrypted); data == nil || slices.Compare(data, expected) != 0 {
		t.Error("Base64Decode was incorrect")
	}
}

func TestBase64UrlEncode(t *testing.T) {
	data := []byte("abc")
	expected := "YWJj"
	if Base64UrlEncode(data) != expected {
		t.Error("Base64UrlEncode was incorrect")
	}
}

func TestBase64UrlDecode(t *testing.T) {
	encrypted := "YWJj"
	expected := []byte("abc")

	if data, _ := Base64UrlDecode(encrypted); data == nil || slices.Compare(data, expected) != 0 {
		t.Error("Base64UrlDecode was incorrect")
	}
}

func TestHtmlEncode(t *testing.T) {
	unescaped := `<script>alert(123);</script>`
	escaped := `&lt;script&gt;alert(123);&lt;/script&gt;`
	if HtmlEncode(unescaped) != escaped {
		t.Error("HtmlEncode was incorrect")
	}
}

func TestHtmlDecode(t *testing.T) {
	unescaped := `<script>alert(123);</script>`
	escaped := `&lt;script&gt;alert(123);&lt;/script&gt;`
	if HtmlDecode(escaped) != unescaped {
		t.Error("HtmlDecode was incorrect")
	}
}
