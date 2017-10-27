package envflag

import (
	"os"
	"strings"
)

func Args(opts ...Option) []string {
	o := &options{
		&matcher{
			false,
			[]string{},
			[]string{},
		},
		false,
		false,
		false,
	}
	for _, opt := range opts {
		opt(o)
	}
	return getEnviron().toArgs(o)
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
