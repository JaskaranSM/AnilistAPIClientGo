package AnilistAPIClientGo

import (
	"bytes"
	"io"
	"strings"
)

func GetBytesReader(data []byte) io.Reader {
	return bytes.NewReader(data)
}

func IsAPIErrorBytes(data []byte) bool {
	return strings.Contains(string(data), "errors")
}

func IsAPIError(err error) bool {
	return strings.Contains(err.Error(), "errors")
}
