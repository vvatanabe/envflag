package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vvatanabe/envflag"
)

const (
	version = "1.0.0"
)

var (
	showVersion bool

	appPort int64

	dbUsername  string
	dbPassword  string
	dbName      string
	dbHostname  string
	dbPort      int64
	dbParameter string
)

func init() {
	os.Setenv("APP_PORT", "5050")
	os.Setenv("DB_USERNAME", "vvatanabe")
	os.Setenv("DB_PASSWORD", "zxcvbn")
	os.Setenv("DB_NAME", "env-flag-2")
	os.Setenv("DB_HOSTNAME", "example.com")
}

func main() {

	flag.BoolVar(&showVersion, "version", false, "show version")

	flag.Int64Var(&appPort, "app_port", 8080, "database port")

	flag.StringVar(&dbUsername, "db_username", "root", "database username")
	flag.StringVar(&dbPassword, "db_password", "", "database password")
	flag.StringVar(&dbName, "db_name", "env-flag-1", "database name")
	flag.StringVar(&dbHostname, "db_hostname", "127.0.0.1", "database host")
	flag.Int64Var(&dbPort, "db_port", 3306, "database port")
	flag.StringVar(&dbParameter, "db_parameter", "?parseTime=true", "database parameter")

	args := envflag.Args(
		envflag.IgnoreCase(),
		envflag.Exact("app_port"),
		envflag.Prefix("db_"),
		envflag.ToLowercase(),
	)

	flag.CommandLine.Parse(args)
	flag.Parse()

	if showVersion {
		fmt.Println("version:", version)
		return
	}

	fmt.Println("[values]")
	fmt.Println("app_port:", appPort)
	fmt.Println("db_username:", dbUsername)
	fmt.Println("db_password:", dbPassword)
	fmt.Println("db_name:", dbName)
	fmt.Println("db_hostname:", dbHostname)
	fmt.Println("db_port:", dbPort)
	fmt.Println("db_parameter:", dbParameter)
}
