package helpers

import "github.com/fatih/color"

var (
	Green  = color.New(color.FgHiGreen).SprintfFunc()
	Red    = color.New(color.FgHiRed).SprintfFunc()
	Yellow = color.New(color.FgHiYellow).SprintfFunc()
)
