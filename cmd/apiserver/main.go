package main

import (
	"awesomeProject3/internal/app/apiserver"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "configs/apiserver.toml", "path to config file (toml)")
}

//@title Test App Statistics
//@version 1.0
//@host localhost:8081
func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(ConfigPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
