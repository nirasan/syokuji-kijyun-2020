package main

//go:generate go build ./cmd/generator/
//go:generate ./generator data

//go:generate go build ./cmd/marshal/
//go:generate ./marshal json
