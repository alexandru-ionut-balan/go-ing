package utils

import (
	"crypto"
	"net/http"
	"time"

	"github.com/alexandru-ionut-balan/go-ing/utils/log"
)

// DateHeader adds a new header containing the current date and time in GMT timezone and standard
// HTTP date header format.
func DateHeder(headers *http.Header) {
	location, err := time.LoadLocation("GMT")
	if err != nil {
		log.Warn.Println("Could not load location GMT. Not calculating date")
		return
	}

	date := time.Now().In(location).Format(time.RFC1123)

	headers.Add("Date", date)
}

// DigestHeader adds a new header containing base64(hashFunction(payload))
//
// See also: [Digest]
func DigestHeader(headers *http.Header, hashFunction crypto.Hash, payload string) {
	digest := Digest(hashFunction, payload)

	headers.Add(hashFunction.String(), digest)
}
