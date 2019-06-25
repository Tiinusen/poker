package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	Username string
	Password string
	Key      string
	Secret   string
}

func main() {
	var config Config
	envconfig.Process("poker", &config)
	if data, err := ioutil.ReadFile("config.json"); err == nil {
		json.Unmarshal(data, &config)
	}
	// ctx := context.Background()
}

func panicIfError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err)
		os.Exit(1)
	}
}
