package helpers

import (
	"encoding/json"
	"log"
	"os"
)

//normal files handler

func GetFileContent(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func GetJsonContent(filepath string) map[string]interface{} {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}

	err_ := json.Unmarshal([]byte(content), &result)
	if err_ != nil {
		return nil
	}

	return result
}

func SavePassword(entry string) {
	logger := Logger{Level: 1, Prefix: "[GoPassword] | "}
	logger.Log("Saved Password")
}
