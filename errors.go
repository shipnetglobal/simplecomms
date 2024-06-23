package main

import "errors"

var (
	ErrorMissingConfig = errors.New("application configuration is required")
	ErrorBadConfig     = errors.New("application configuration is bad")
)
