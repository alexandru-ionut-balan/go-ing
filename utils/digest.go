package utils

import (
	"crypto"
	"encoding/base64"
	"net/url"
	"strings"
)

// Hash applies the given hash function to the bytes of the payload string,
// returning the resulting hashed byte array.
func Hash(hashFunction crypto.Hash, payload string) []byte {
	hash := hashFunction.New()
	bytePayload := []byte(payload)

	hash.Write(bytePayload)
	return hash.Sum(nil)
}

// Base64 applies standard base64 encoding to the byte array given as payload.
func Base64(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

// UrlEncode will use the standard Encode() function from url.Values, but
// it will replace "%20" characters with "+".
func UrlEncode(payload url.Values) string {
	return strings.ReplaceAll(payload.Encode(), "%20", "+")
}

// Digest applies the hashing function to the string payload and encodes the resulting
// bytes in base64 encoding.
//
// digest = base64(hashFunction(payload))
func Digest(hashFunction crypto.Hash, payload string) string {
	return Base64(Hash(hashFunction, payload))
}
