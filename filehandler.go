package main

import (
	"encoding/json"
	"log"
	"os"
)

//normal files handler

func getFileContent(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func getJsonContent(filepath string) map[string]interface{} {
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
