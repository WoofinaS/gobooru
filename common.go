package gobooru

import (
	"encoding/xml"
	"io"
	"net/http"
)

// Request is an internal function, exported for conciseness.
func Request(url string, v interface{}) error {
	// TODO: is there a better way to write this?
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, v)
}
