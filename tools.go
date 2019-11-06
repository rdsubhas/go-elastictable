// +build tools
// Declaring build-time dependencies, these are ignored at compile-time
// Refer https://github.com/golang/go/issues/25922

package main

import (
	_ "golang.org/x/lint"
)
