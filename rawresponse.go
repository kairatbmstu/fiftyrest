package main

type RawResponse interface {
	GetStatus() int
	GetStatusText() string
	GetHeaders() Headers
	GetContent() []byte // InputStream
	GetContentAsBytes() []byte
	GetContentAsString() string
	GetContentAsStringWithCharset(String charset) string
	GetContentReader() Reader //InputStreamReader
	HasContent() bool
	GetContentType() string
	GetEncoding() string
	GetConfig() *Config
	ToSummary() HttpResponseSummary
}
