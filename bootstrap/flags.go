package bootstrap

import (
	"flag"
)

// InitFlags init and parse the flags passed in command
func InitFlags() string {
	confPath := flag.String("config", "./", "Path for the config file")
	flag.Parse()

	return *confPath
}
