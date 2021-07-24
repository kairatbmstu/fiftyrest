package main

import "time"

type Cookie struct {
	Name     string
	value    string
	domain   string
	path     string
	httpOnly bool
	maxAge   int
	expires  time.Time
	secure   bool
	SameSite SameSite
}

type SameSite string

const (
	SameSiteNone   = iota
	SameSiteStrict = iota
	SameSiteLax    = iota
)

type Cookies []Cookie
