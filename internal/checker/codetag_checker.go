package checker

import (
    "fmt"
    "strings"
)

type CodetagChecker struct {
    Content  string
    Filename string
}

func NewCodetagChecker(content string, filename string) *CodetagChecker {
    return &CodetagChecker{Content: content, Filename: filename}
}

func (c *CodetagChecker) Run() []string {
    var errors []string
    lines := strings.Split(c.Content, "\n")

    codetags := map[string]bool{
        "TODO":  true,
        "FIXME": true,
        "HACK":  true,
        "NOTE":  true,
        "XXX":   true,
        "BUG":  true,
    }

    for i, line := range lines {
        trimmedLine := line
        hashCount := 0

        for len(trimmedLine) > 0 && trimmedLine[0] == '#' {
            hashCount++
            trimmedLine = trimmedLine[1:]
        }

        if hashCount > 0 {
            trimmedLine = strings.TrimSpace(trimmedLine)

            for tag := range codetags {
                if strings.Contains(trimmedLine, tag) {
                    errors = append(errors, fmt.Sprintf("%s:%d: %s codetag detected (%s)", c.Filename, i+1, tag, trimmedLine))
                    break
                }
            }
        }
    }

    return errors
}