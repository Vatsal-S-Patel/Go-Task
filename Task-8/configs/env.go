package configs

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

// envFile is map, stores .env file
var envFile map[string]string

// ReadEnv function reads the .env file
func ReadEnv() error {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	envFile = envMap
	return nil
}

// GetEnv function return the value stored against key in .env file and return it with error if any
func GetEnv(detail string) (string, error) {
	data, ok := envFile[detail]

	if !ok {
		return "", errors.New("ERROR: Data Not Exist in .env File")
	} else {
		return data, nil
	}
}
