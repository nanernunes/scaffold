package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/nanernunes/scaffold/cmd/scaffold/cmd"
	"github.com/nanernunes/scaffold/pkg/helpers"
)

func main() {
	reqs := helpers.NewRequirements("make", "go", "git", "sqlite3")

	if missing := reqs.MissingPackages(); len(missing) != 0 {
		fmt.Printf("Error: missing the following packages: %s.\n", color.RedString(strings.Join(missing, ", ")))
		os.Exit(1)
	}

	cmd.Execute()
}
