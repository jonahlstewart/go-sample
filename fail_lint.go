//go:build lintfail

package main

// This will trigger various golangci-lint warnings
var UnusedVar string     // unused var
var mixedCaps string     // wrong case
func () invalidSpacing() {} // formatting issues