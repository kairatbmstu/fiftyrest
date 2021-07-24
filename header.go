package main

type Header interface {
	GetName() string
	GetValue() string
}
