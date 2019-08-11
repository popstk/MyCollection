package main

import (
	"errors"
	"github.com/thoas/go-funk"
	"strconv"
	"strings"
)

// SplitMultiChar split by chars
func SplitMultiChar(s string, delims string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return strings.Contains(delims, strings.ToLower(string(r)))
	})
}


func ValidName(name string) string {
	parts := funk.FilterString(SplitMultiChar(name, conf.Delims), func(s string) bool {
		return !funk.Contains(conf.Ignores, strings.ToLower(s))
	})
	name = strings.ToUpper(strings.Join(parts, ""))
	return name
}

// FormatName format names
func FormatName(name string) (string, error) {
	name = ValidName(name)

	prefix := ""
	for _, pre := range conf.Prefix {
		if strings.HasPrefix(name, strings.ToUpper(pre)) {
			prefix = pre
			break
		}
	}

	if prefix == "" {
		return "", errors.New("no match prefix")
	}

	rest := name[len(prefix):]
	n, err := strconv.Atoi(rest)
	if err != nil {
		return "", errors.New("not number:" + rest)
	}

	return prefix + "-" + strconv.Itoa(n), nil
}
