package checker

import (
    "fmt"
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

	for lineC, line := range c.Content {
		if len(line) > 0 && line[0] == 'd' && line[1] == 'e' && line[2] == 'f' {
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
				if !isSnakeCase(functionName) {
					errors = append(errors, fmt.Sprintf("%s:%d: Function name '%s' should be in snake_case", c.Filename, lineC+1, functionName))
				}
			}
		}
	}
    
    return errors
}

func isSnakeCase(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return false
		}
	}

	return true
}