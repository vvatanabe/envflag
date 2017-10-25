package envflag

import (
	"os"
	"strings"
)

type option struct {
	longHyphen bool
	prefix     string
}

func ArgsFromEnviron(longHyphen bool, prefix string) []string {
	return getEnviron().toArgs(&option{longHyphen, prefix})
}

func getEnviron() environ {
	environ := make(map[string]string)
	for _, s := range os.Environ() {
		splits := strings.SplitN(s, "=", 2)
		k := splits[0]
		v := splits[1]
		environ[k] = v
	}
	return environ
}

type environ map[string]string

func (environ environ) toArgs(opt *option) []string {
	var args []string
	for k, v := range environ {
		if opt.prefix != "" && !strings.HasPrefix(v, opt.prefix) {
			continue
		}
		hyphen := "-"
		if opt.longHyphen && len(k) > 1 {
			hyphen = "--"
		}
		name := hyphen + k
		args = append(args, name, v)
	}
	return args
}
