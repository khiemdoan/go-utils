package cryptoutils

import (
	"encoding/base32"
	"encoding/base64"
	"html"
)

func Base32Encode(data []byte) (encodedText string) {
	encodedText = base32.StdEncoding.EncodeToString(data)
	return
}

func Base32Decode(encodedText string) (data []byte, err error) {
	data, err = base32.StdEncoding.DecodeString(encodedText)
	return
}

func Base64Encode(data []byte) (encodedText string) {
	encodedText = base64.StdEncoding.EncodeToString(data)
	return
}

func Base64Decode(encodedText string) (data []byte, err error) {
	data, err = base64.StdEncoding.DecodeString(encodedText)
	return
}

func Base64UrlEncode(data []byte) (encodedText string) {
	encodedText = base64.URLEncoding.EncodeToString(data)
	return
}

func Base64UrlDecode(encodedText string) (data []byte, err error) {
	data, err = base64.URLEncoding.DecodeString(encodedText)
	return
}

func HtmlEncode(unescaped string) string {
	return html.EscapeString(unescaped)
}

func HtmlDecode(escaped string) string {
	return html.UnescapeString(escaped)
}
