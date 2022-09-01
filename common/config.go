package common

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	ServerPort        string `env:"SV_PORT"`
	ServerEnvironment string `env:"SV_ENVIRONMENT"`
	DbName            string `env:"DB_NAME"`
	MongoURI          string `env:"MONGO_URI"`
}

var (
	Config *Configuration = LoadEnv()
)

func LoadEnv() *Configuration {
	var config Configuration
	_, file, _, _ := runtime.Caller(1)
	formatedPath := path.Join(filepath.Dir(file), "\\..", ".env")
	err := godotenv.Load(formatedPath)
	if err != nil {
		log.Fatal(err, formatedPath)
	}
	env.Set(&config)
	if os.Getenv("RUNNING_MAIN") == "" {
		config.DbName = os.Getenv("DB_TEST_NAME")
	}
	return &config
}
