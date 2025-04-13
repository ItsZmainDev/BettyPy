package checker

import (
    "fmt"
    "strings"
)

type EmptyFileChecker struct {
    Content  string
    Filename string
}

func NewEmptyFileChecker(content string, filename string) *EmptyFileChecker {
    return &EmptyFileChecker{Content: content, Filename: filename}
}

func (c *EmptyFileChecker) Run() []string {
    var errors []string

    lines := strings.Split(c.Content, "\n")

    noEmptyLines := 0
    for _, line := range lines {
        if strings.TrimSpace(line) != "" {
            noEmptyLines++
        }
    }

    if noEmptyLines == 0 {
        errors = append(errors, fmt.Sprintf("File %s is empty or contains only empty lines", c.Filename))
    }

    return errors
}