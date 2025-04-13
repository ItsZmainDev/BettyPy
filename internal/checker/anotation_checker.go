package checker

import (
    "fmt"
    "strings"
)

type AnotationChecker struct {
    Content  string
    Filename string
}

func NewAnotationChecker(content string, filename string) *AnotationChecker {
    return &AnotationChecker{Content: content, Filename: filename}
}

func (c *AnotationChecker) Run() []string {
    var errors []string
    lines := strings.Split(c.Content, "\n")

    for i, line := range lines {
        trimmedLine := line
        hashCount := 0

        for len(trimmedLine) > 0 && trimmedLine[0] == '#' {
            hashCount++
            trimmedLine = trimmedLine[1:]
        }

        if hashCount > 0 {
            trimmedLine = strings.TrimSpace(trimmedLine)

			switch {
				case strings.Contains(trimmedLine, "TODO"):
					errors = append(errors, fmt.Sprintf("%s:%d: TODO comment detected (%s)", c.Filename, i+1, trimmedLine))
				case strings.Contains(trimmedLine, "FIXME"):
					errors = append(errors, fmt.Sprintf("%s:%d: FIXME comment detected (%s)", c.Filename, i+1, trimmedLine))
			}
        }
    }

    return errors
}