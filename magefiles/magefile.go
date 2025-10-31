//go:build mage

// Automated build tasks for development and release to production.
package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const appName = "rexlists"
const targetDir = "target"

// Namespaces
type Build mg.Namespace
type Check mg.Namespace

// Build front end application
func (Build) FrontEnd() error {
	fmt.Println("Building front end.")
	return buildVariant(appName, "fe")

}

// Build back end application
func (Build) BackEnd() error {
	fmt.Println("Building back end.")
	return buildVariant(appName, "be")
}

func buildVariant(appName, variant string) error {
	var bin = fmt.Sprintf("target/%s-%s", appName, variant)
	var src = fmt.Sprintf("cmd/%s-%s/main.go", appName, variant)
	return sh.Run("go", "build", "-o", bin, src)
}

// Build both applications
func (Build) All() {
	mg.Deps(Build.FrontEnd, Build.BackEnd)
}

// Run linters.
func (Check) Lint() error {
	fmt.Println("Run Linters.")
	return sh.RunV("golangci-lint", "run")
}

// Check code quality.
func (Check) All() {
	mg.Deps(Check.Lint)
}
