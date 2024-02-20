package utils

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

// Temporary map to get data from env whenever needed
var fileEnvMap map[string]string

// ReadEnvFile fucntion read env file and return error if any
func ReadEnvFile() (map[string]string, error) {

	// Reading env file and envMap is a map contains all details from env file
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	fileEnvMap = envMap

	return envMap, nil
}

// Get data from env file using key
func GetFromEnvFile(detail string) (string, error) {
	res, ok := fileEnvMap[detail]
	if ok {
		return res, nil
	} else {
		return "", errors.New("ERROR: Key not exist")
	}
}
