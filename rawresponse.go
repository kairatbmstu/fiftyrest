package main

import "io"

type RawResponse interface {
	GetStatus() int
	GetStatusText() string
	GetHeaders() Headers
	GetContent() []byte // InputStream
	GetContentAsBytes() []byte
	GetContentAsString() string
	GetContentAsStringWithCharset(charset string) string
	GetContentReader() io.ByteReader //InputStreamReader
	HasContent() bool
	GetContentType() string
	GetEncoding() string
	GetConfig() *Config
	ToSummary() HttpResponseSummary
}
