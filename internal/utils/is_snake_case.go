package utils

import "regexp"

func IsSnakeCase(s string) bool {
    return regexp.MustCompile(`^[a-z]+(_[a-z0-9]+)*$`).MatchString(s)
}