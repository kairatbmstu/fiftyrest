package main

type ContentType string

const (
	APPLICATION_ATOM_XML        ContentType = "application/atom+xml"
	APPLICATION_FORM_URLENCODED ContentType = "application/x-www-form-urlencoded"
	APPLICATION_JSON            ContentType = "application/json"
	APPLICATION_JSON_PATCH      ContentType = "application/json-patch+json"
	APPLICATION_OCTET_STREAM    ContentType = "application/octet-stream"
	APPLICATION_SVG_XML         ContentType = "application/svg+xml"
	APPLICATION_XHTML_XML       ContentType = "application/xhtml+xml"
	APPLICATION_XML             ContentType = "application/xml"
	IMAGE_BMP                   ContentType = "image/bmp"
	IMAGE_GIF                   ContentType = "image/gif"
	IMAGE_JPEG                  ContentType = "image/jpeg"
	IMAGE_PNG                   ContentType = "image/png"
	IMAGE_SVG                   ContentType = "image/svg+xml"
	IMAGE_TIFF                  ContentType = "image/tiff"
	IMAGE_WEBP                  ContentType = "image/webp"
	MULTIPART_FORM_DATA         ContentType = "multipart/form-data"
	TEXT_HTML                   ContentType = "text/html"
	TEXT_PLAIN                  ContentType = "text/plain"
	TEXT_XML                    ContentType = "text/xml"
	WILDCARD                    ContentType = "*/*"
)
