package checker

import (
	"fmt"
	"strings"
)

type DockstringChecker struct {
    Content  string
    Filename string
}

func NewDockstringChecker(content string, filename string) *DockstringChecker {
    return &DockstringChecker{Content: content, Filename: filename}
}

func (c *DockstringChecker) Run() []string {
    var errors []string

    lines := strings.Split(c.Content, "\n")
    for i, line := range lines {
        trimmedLine := strings.TrimSpace(line)
        if strings.HasPrefix(trimmedLine, "class ") || strings.HasPrefix(trimmedLine, "def ") {
            if i+1 < len(lines) {
                nextLine := strings.TrimSpace(lines[i+1])
                if !(strings.HasPrefix(nextLine, `"""`) || strings.HasPrefix(nextLine, `'''`)) {
                    entityType := "Class"
                    if strings.HasPrefix(trimmedLine, "def ") {
                        entityType = "Function"
                    }
                    errors = append(errors, fmt.Sprintf("%s:%d: %s '%s' is missing a docstring.", c.Filename, i+1, entityType, trimmedLine))
                }
            } else {
				errors = append(errors, fmt.Sprintf("%s:%d: %s '%s' is missing a docstring.", c.Filename, i+1, "Entity", trimmedLine))
            }
        }
    }

    return errors
}