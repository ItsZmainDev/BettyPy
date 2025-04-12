package checker

import "fmt"

type Checker interface {
	Run() []string
}

type LineLengthChecker struct {
    Content  []string
    Filename string
}

func NewLineLengthChecker(content []string, filename string) *LineLengthChecker {
    return &LineLengthChecker{Content: content, Filename: filename}
}

func (c *LineLengthChecker) Run() []string {
    var errors []string
    for i, line := range c.Content {
        if len(line) > 80 {
            errors = append(errors, fmt.Sprintf("%s:%d: line too long (%d characters)", c.Filename, i+1, len(line)))
        }
    }
    return errors
}