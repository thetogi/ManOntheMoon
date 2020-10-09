package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvParam struct {
	Key   string
	Value string
}

type Data struct {
	Environment string
	EnvParams   []EnvParam
	EnvShow     bool
}

var environment string
var envParams []EnvParam
var envShow bool

func EnvData() Data {
	return Data{environment, envParams, envShow}
}

func ToggleEnvShow() {
	envShow = !envShow
}

func EnvShowStatus() bool {
	return envShow
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envFile, err := godotenv.Read()
	for k, v := range envFile {
		envParams = append(envParams, EnvParam{k, v})
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	environment = os.Getenv("ENVIRONMENT")
	envShow = false
}
