package main

import (
	"flag"
	"fmt"
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

	args := envflag.Args(
		envflag.Exact("ENV_FLAG_STR"),
		envflag.Prefix("ENV_FLAG"),
		envflag.Lowercase(),
		envflag.BoolValue(),
	)

	flag.CommandLine.Parse(args)
	flag.Parse()

	fmt.Printf("env_flag_int: %v\n", envFlagInt)
	fmt.Printf("env_flag_bool: %v\n", envFlagBool)
	fmt.Printf("env_flag_str: %v\n", envFlagStr)

}
