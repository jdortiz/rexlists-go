//go:build mage

// Automated build tasks for development and release to production.
package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

const appName = "rexlists"
const targetDir = "target"

// Build front end application
func BuildFrontEnd() error {
	fmt.Println("Building front end.")
	return buildVariant(appName, "fe")

}

// Build back end application
func BuildBackEnd() error {
	fmt.Println("Building back end.")
	return buildVariant(appName, "be")
}

func buildVariant(appName, variant string) error {
	var bin = fmt.Sprintf("target/%s-%s", appName, variant)
	var src = fmt.Sprintf("cmd/%s-%s/main.go", appName, variant)
	return sh.Run("go", "build", "-o", bin, src)
}
