//go:build mage

// Automated build tasks for development and release to production.
package main

import "github.com/magefile/mage/sh"

// Build front end application
func Build() error {
	return sh.Run("go", "build", "-o", "target/rexlists-fe", "cmd/rexlists-fe/main.go")
}
