package main

type HttpMethod string

const (
	HttpMethodGet     = "GET"
	HttpMethodPost    = "POST"
	HttpMethodPut     = "PUT"
	HttpMethodDelete  = "DELETE"
	HttpMethodPatch   = "PATCH"
	HttpMethodHead    = "HEAD"
	HttpMethodOptions = "OPTIONS"
	HttpMethodTrace   = "TRACE"
)
