package main

import (
	"github.com/softwareplace/go-password/pkg/str"
)

func main() {
	generate := str.Default().
		Generate()
	println(generate)
}
