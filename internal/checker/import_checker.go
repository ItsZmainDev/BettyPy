package checker

import (
    "fmt"
    "regexp"
    "strings"
)

type UnusedImportChecker struct {
    Content  []string
    Filename string
}

func NewUnusedImportChecker(content []string, filename string) *UnusedImportChecker {
    return &UnusedImportChecker{Content: content, Filename: filename}
}

func (c *UnusedImportChecker) Run() []string {
    var errors []string
    imports := make(map[string]bool)

    importRegex := regexp.MustCompile(`^\s*(import\s+([\w\.]+)|from\s+([\w\.]+)\s+import\s+([\w,\s]+))`)
    usageRegexCache := make(map[string]*regexp.Regexp)

    for _, line := range c.Content {
        matches := importRegex.FindStringSubmatch(line)
        if matches != nil {
            if matches[2] != "" {
                imports[matches[2]] = false
            } else if matches[3] != "" && matches[4] != "" {
                for _, item := range strings.Split(matches[4], ",") {
                    imports[strings.TrimSpace(item)] = false
                }
            }
        }
    }

    for _, line := range c.Content {
        trimmedLine := strings.TrimSpace(line)

        if strings.HasPrefix(trimmedLine, "#") || importRegex.MatchString(trimmedLine) {
            continue
        }

        for imp := range imports {
            if _, exists := usageRegexCache[imp]; !exists {
                usageRegexCache[imp] = regexp.MustCompile(`\b` + regexp.QuoteMeta(imp) + `\b`)
            }
            if usageRegexCache[imp].MatchString(trimmedLine) {
                imports[imp] = true
            }
        }
    }

    for imp, used := range imports {
        if !used {
            errors = append(errors, fmt.Sprintf("%s: unused import '%s'", c.Filename, imp))
        }
    }

    return errors
}