package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadEntries(filepath string) []Entry {
	jsonFile, err := os.Open(filepath)
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)
	var entries []Entry
	err1 := json.Unmarshal(byteValue, &entries)
	if err1 != nil {
		return nil
	}

	return entries
}

func saveEntries(entries []Entry) {
	file, _ := json.MarshalIndent(entries, "", " ")

	_ = os.WriteFile("./passwords.json", file, 0644)
}

func SaveEntry(entry Entry) {
	logger := Logger{Level: 1, Prefix: ""}
	logger.Log("Saved Password")
	var currentEntries []Entry = LoadEntries("./passwords.json")
	entries := append(currentEntries, entry)
	saveEntries(entries)
}

func SaveEntries(entries []Entry) {
	logger := Logger{Level: 1, Prefix: ""}
	logger.Log("Saved Password")

	saveEntries(entries)
}

func LoadTXTFile(filename string) string {
	content, _ := os.ReadFile(filename)
	return string(content[:])
}

func SaveTXTFile(filename string, content string) {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	file.WriteString(content)
	defer file.Close()
}
