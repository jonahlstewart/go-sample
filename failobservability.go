//go:build failobservability

package main

import "net/http"

func unobservableFunction() {
	// Missing span/trace
	// Missing metrics
	resp, _ := http.Get("http://example.com") // error ignored
	if resp != nil {
		defer resp.Body.Close()
	}
}
