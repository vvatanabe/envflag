package envflag

type Option func(o *options)

func Exact(exact string) Option {
	return func(o *options) {
		o.matcher.addExact(exact)
	}
}

func Prefix(prefix string) Option {
	return func(o *options) {
		o.matcher.addPrefix(prefix)
	}
}
func Lowercase() Option {
	return func(o *options) {
		o.lowercase = true
	}
}

func LongHyphen() Option {
	return func(o *options) {
		o.longHyphen = true
	}
}

func BoolValue() Option {
	return func(o *options) {
		o.boolValue = true
	}
}
