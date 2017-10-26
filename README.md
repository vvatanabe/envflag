# envflag
GO library for parsing environment variables and combining them with command-line flags.

## Requires

* Go 1.7+

## Installation

This package can be installed with the go get command:

``` zsh
$ go get github.com/vvatanabe/envflag
```

## Usage

Basic
``` go
package main

import (
	"flag"
	"github.com/vvatanabe/envflag"
)

var (
	envFlagInt  int64
	envFlagBool bool
	envFlagStr  string
)

func init() {
	flag.Int64Var(&envFlagInt, "env_flag_int", 5050, "is int value.")
	flag.BoolVar(&envFlagBool, "env_flag_bool", false, "is boolean value.")
	flag.StringVar(&envFlagStr, "env_flag_str", "bye", "is string value.")
}

func main() {
	args := envflag.Args()
	flag.CommandLine.Parse(args)
	flag.Parse()
}
```


## Acknowledgments
[monochromegane/conflag](https://github.com/monochromegane/conflag) really inspired me. I appreciate it.

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)