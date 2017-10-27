# envflag [![Build Status](https://travis-ci.org/vvatanabe/envflag.svg?branch=master)](https://travis-ci.org/vvatanabe/envflag)
GO library for parsing environment variables and combining them with command-line flags.

## Requires

* Go 1.7+

## Installation

This package can be installed with the go get command:

``` sh
$ go get github.com/vvatanabe/envflag
```

## Usage

Please see also example code.

### Basic
``` go
func main() {

	var (
		envFlagInt  int64
		envFlagBool bool
		envFlagStr  string
	)

	flag.Int64Var(&envFlagInt, "env_flag_int", 8080, "is int value.")
	flag.BoolVar(&envFlagBool, "env_flag_bool", false, "is boolean value.")
	flag.StringVar(&envFlagStr, "env_flag_str", "hello", "is string value.")

	// You can set some options for parsing environment variables.
	args := envflag.Args(
		envflag.Prefix("ENV_FLAG"),
		envflag.Lowercase(),
		)

	flag.CommandLine.Parse(args)

	flag.Parse()

}
```

### Functional Options

It compares variable name without distinguishing uppercase and lowercase
``` go
func IgnoreCase() Option
```

It adds exact pattern for matching. And it can specify more than one.
``` go
func Exact(exact string) Option
```

It adds prefix pattern for matching. And it can specify more than one.
``` go
func Prefix(prefix string) Option
```

It converts all variable names to the lowercase.
``` go
func ToLowercase() Option
```

It adds a long hyphen to the head of 2 characters or more variable names.
``` go
func LongHyphen() Option
```

It connects them with `=` if the variables are booleans.
``` go
func BoolValue() Option
```

## Priority

`command-line flag` > `environment variable` > `flag default value`

## Acknowledgments
[monochromegane/conflag](https://github.com/monochromegane/conflag) really inspired me. I appreciate it.

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)