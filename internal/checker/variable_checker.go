package checker

import (
    "fmt"
    "strings"
    "bettypy/internal/utils"
	"regexp"
)

type VariableChecker struct {
    Content  []string
    Filename string
}

func NewVariableChecker(content []string, filename string) *VariableChecker {
    return &VariableChecker{Content: content, Filename: filename}
}

func (c *VariableChecker) Run() []string {
    var errors []string
    definedVariables := make(map[string]int)
    usedVariables := make(map[string]bool)

    for lineC, line := range c.Content {
        if len(line) > 0 && strings.Contains(line, "=") {
            parts := strings.Split(line, "=")
            variableName := strings.TrimSpace(parts[0])
            if variableName != "" && variableName != "def" {
                if !utils.IsSnakeCase(variableName) {
                    errors = append(errors, fmt.Sprintf("%s:%d: Variable name '%s' should be in snake_case", c.Filename, lineC+1, variableName))
                }
                definedVariables[variableName] = lineC + 1
            }
        }
    }

    for lineC, line := range c.Content {
        for variableName, defLine := range definedVariables {
            if lineC+1 <= defLine {
                continue
            }

            pattern := fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(variableName))
            matched, err := regexp.MatchString(pattern, line)
            if err != nil {
                continue
            }
            if matched {
                usedVariables[variableName] = true
            }
        }
    }

    for variableName, line := range definedVariables {
        if !usedVariables[variableName] {
            errors = append(errors, fmt.Sprintf("%s:%d: Variable '%s' is defined but not used", c.Filename, line, variableName))
        }
    }

    return errors
}