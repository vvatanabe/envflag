package envflag

import (
	"errors"
	"fmt"
	"strings"
)

type environ map[string]string

func (environ environ) toArgs(opts *options) []string {
	var args []string
	shouldBeMatch := opts.matcher.hasPattern()
	for k, v := range environ {
		if shouldBeMatch && !opts.matcher.match(k) {
			continue
		}
		hyphen := "-"
		if opts.longHyphen && len(k) > 1 {
			hyphen = "--"
		}
		name := hyphen + k
		if opts.toLowercase {
			name = strings.ToLower(name)
		}

		if boolean, err := strToBool(v); err == nil {
			if opts.boolValue {
				args = append(args, fmt.Sprintf("%s=%t", name, boolean))
			} else {
				args = append(args, name)
			}
		} else {
			args = append(args, name, v)
		}
	}
	return args
}

func strToBool(str string) (bool, error) {
	switch str {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, errors.New(fmt.Sprintf("ParseBool: %s", str))
}
