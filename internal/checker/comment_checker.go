package checker

import "fmt"

type CommentChecker struct {
    Content  []string
    Filename string
}

func NewCommentChecker(content []string, filename string) *CommentChecker {
    return &CommentChecker{Content: content, Filename: filename}
}

func (c *CommentChecker) Run() []string {
    var errors []string
    for i, line := range c.Content {
        trimmedLine := line
        hashCount := 0

        for len(trimmedLine) > 0 && trimmedLine[0] == '#' {
            hashCount++
            trimmedLine = trimmedLine[1:]
        }

        if hashCount > 0 {
            if len(trimmedLine) > 0 && trimmedLine[0] != ' ' {
                errors = append(errors, fmt.Sprintf("%s:%d: comment should start with a space after #", c.Filename, i+1))
            }

            trimmedLine = trimLeadingSpaces(trimmedLine)
            if len(trimmedLine) < 10 {
                errors = append(errors, fmt.Sprintf("%s:%d: comment too short (%d characters)", c.Filename, i+1, len(trimmedLine)))
            }
        }
    }
    return errors
}

func trimLeadingSpaces(s string) string {
    for len(s) > 0 && s[0] == ' ' {
        s = s[1:]
    }
    return s
}