package checker

import (
    "fmt"
)

type NewLineChecker struct {
    Content  string
    Filename string
}

func NewNewLineChecker(content string, filename string) *NewLineChecker {
    return &NewLineChecker{Content: content, Filename: filename}
}

func (c *NewLineChecker) Run() []string {
    var errors []string

	lines := []rune(c.Content)

	if len(lines) > 0 && lines[len(lines)-1] != '\n' {
		errors = append(errors, fmt.Sprintf("%s: file should end with a new line", c.Filename))
	}

	return errors
}