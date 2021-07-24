package main

type HttpResponseSummary interface {
	GetStatus() int
	GetStatusText() string
}
