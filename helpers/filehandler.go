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

func writeStringToFile(filepath string, content string) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	_, err = file.WriteString(content)
	if err != nil {
		return
	}
	err2 := file.Close()
	if err2 != nil {
		return
	}
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

func SavePassword(entry Entry) {
	logger := Logger{Level: 1, Prefix: "[GoPassword] | "}
	logger.Log("Saved Password")
	currentFileContent := GetFileContent("passwords.json")
	writeStringToFile("./passwords.json", currentFileContent+"dcccc")
}
