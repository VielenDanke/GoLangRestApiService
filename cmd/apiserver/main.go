package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/vielendanke/restful-service/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}
	errStartServer := apiserver.Start(config)

	if errStartServer != nil {
		log.Fatal(errStartServer)
	}
}
