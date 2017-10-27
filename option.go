package envflag

type options struct {
	matcher     *matcher
	longHyphen  bool
	toLowercase bool
	boolValue   bool
}

type Option func(o *options)

// It compares variable name without distinguishing uppercase and lowercase
func IgnoreCase() Option {
	return func(o *options) {
		o.matcher.ignoreCase = true
	}
}

// It adds exact pattern for matching. And it can specify more than one.
func Exact(exact string) Option {
	return func(o *options) {
		o.matcher.addExact(exact)
	}
}

// It adds prefix pattern for matching. And it can specify more than one.
func Prefix(prefix string) Option {
	return func(o *options) {
		o.matcher.addPrefix(prefix)
	}
}

// It converts all variable names to the lowercase.
func ToLowercase() Option {
	return func(o *options) {
		o.toLowercase = true
	}
}

// It adds a long hyphen to the head of 2 characters or more variable names.
func LongHyphen() Option {
	return func(o *options) {
		o.longHyphen = true
	}
}

// It connects them with `=` if the variables are booleans.
func BoolValue() Option {
	return func(o *options) {
		o.boolValue = true
	}
}
