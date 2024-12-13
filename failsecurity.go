//go:build securityfail

package main

import (
    "fmt"
    "log"
    "os"
)

// This file includes potential security vulnerabilities
func securityVulnerabilities() {
    // Potential security issues
    password := os.Getenv("SECRET_PASSWORD") // Avoid logging sensitive info
    fmt.Println(password) // Potential information disclosure

    // Potential command injection
    command := os.Getenv("USER_INPUT")
    exec.Command("sh", "-c", command) // Dangerous command execution
}