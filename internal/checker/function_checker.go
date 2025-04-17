package checker

import (
    "fmt"
    "strings"
    "bettypy/internal/utils"
)

type FunctionChecker struct {
    Content  []string
    Filename string
}

func NewFunctionChecker(content []string, filename string) *FunctionChecker {
    return &FunctionChecker{Content: content, Filename: filename}
}

func (c *FunctionChecker) Run() []string {
    var errors []string
    definedFunctions := make(map[string]int)
    usedFunctions := make(map[string]bool)

    for lineC, line := range c.Content {
        if len(line) > 0 && strings.HasPrefix(line, "def ") {
            start := 4
            for start < len(line) && line[start] == ' ' {
                start++
            }

            end := start
            for end < len(line) && line[end] != '(' {
                end++
            }

            functionName := line[start:end]
            if functionName != "" && functionName != "def" {
                if !utils.IsSnakeCase(functionName) {
                    errors = append(errors, fmt.Sprintf("%s:%d: Function name '%s' should be in snake_case", c.Filename, lineC+1, functionName))
                }
                definedFunctions[functionName] = lineC + 1
            }
        }
    }

	for i, line := range c.Content {
		trimmedLine := strings.TrimSpace(line)
		for functionName, defLine := range definedFunctions {
			if i+1 == defLine {
				continue
			}

			if strings.Contains(trimmedLine, functionName+"(") {
				usedFunctions[functionName] = true
			}
		}
	}

    for functionName, line := range definedFunctions {
        if !usedFunctions[functionName] {
            errors = append(errors, fmt.Sprintf("%s:%d: Function '%s' is defined but not used", c.Filename, line, functionName))
        }
    }

    return errors
}